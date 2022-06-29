package common

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"machine_info_gatherer/errorslist"
	"machine_info_gatherer/model"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/quay/claircore/osrelease"
)

var (
	SudoUserPassword = ""
)

func SetSudoPassword(pass string) {
	SudoUserPassword = pass
}

func RootNeeded(arg string) string {
	if arg == "None" || arg == "unknown" {
		return arg + " (need root access)"
	} else {
		return arg
	}
}

func IsService(svc string) bool {
	if strings.Contains(svc, "loaded act") || strings.Contains(svc, "loaded fail") {
		return true
	} else {
		return false
	}
}

func IsServiceRunning(svc string) bool {
	if strings.Contains(svc, "running") {
		return true
	} else {
		return false
	}
}

func RunFullCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)

	// Use a bytes.Buffer to get the output
	var buf bytes.Buffer
	cmd.Stdout = &buf

	cmd.Start()

	// Use a channel to signal completion so we can use a select statement
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	// Start a timer
	timeout := time.After(5 * time.Second)

	// The select statement allows us to execute based on which channel
	// we get a message from first.
	var err error
	select {
	case <-timeout:
		// Timeout happened first, kill the process and print a message.
		cmd.Process.Kill()
		err = errors.New(errorslist.ErrCommandTimeOut)
	case err = <-done:
		// Command completed before timeout. Print output and error if it exists.
		if err != nil {
			fmt.Println("Non-zero exit code:", err)
		}
	}

	return buf.String(), err
}

func RunFullCommandWithSudo(cmd string) (string, error) {
	return RunFullCommand("sudo -S <<< " + SudoUserPassword + " " + cmd)
}

func ParseService(svc string) *model.Service {
	service := model.Service{}

	svc = strings.ReplaceAll(svc, "●", " ")

	svc = strings.Join(strings.Fields(svc), " ")

	splitedSvc := strings.Split(svc, " ")

	service.Display_name = strings.TrimSpace(splitedSvc[0])
	service.State = strings.TrimSpace(splitedSvc[2])
	service.Status = strings.TrimSpace(splitedSvc[3])

	return &service
}

func ReadOSRelease() *map[string]string {
	ctx := context.Background()
	var b []byte

	sys := os.DirFS("/")

	// Look for an os-release file.
	b, err := fs.ReadFile(sys, osrelease.Path)
	if err != nil {
		fmt.Printf("error:%#v", err)
	}
	m, err := osrelease.Parse(ctx, bytes.NewReader(b))
	if err != nil {
		fmt.Printf("error:%#v", err)
	}
	return &m
}
