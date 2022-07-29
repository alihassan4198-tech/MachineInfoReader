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

const (
	rootAccessNeeded = " (need root previliges)"
)

func SetSudoPassword(pass string) {
	SudoUserPassword = pass
}

func RootNeeded(arg string) string {
	if arg == "None" || arg == "unknown" {
		return arg + rootAccessNeeded
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

func RunFullCommandNoTimeOut(command string) (string, error) {
	result, err := exec.Command(SudoUserPassword, "&", "sudo", "-S", command).CombinedOutput()

	return string(result), err
}

func RunFullCommand(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)

	// Use a bytes.Buffer to get the output
	var buf bytes.Buffer
	var bufErr bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &bufErr

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
		err = errors.New(bufErr.String())
	case err = <-done:
		// Command completed before timeout. Print output and error if it exists.
		if err != nil {
			err = errors.New(bufErr.String())
		}
	}

	res := buf.String()
	return res, err
}

func NeedSudoPreviliges(err error) string {
	if err.Error() == errorslist.ErrCommandTimeOut {
		return err.Error() + rootAccessNeeded
	}
	return err.Error()
}

func RunFullCommandWithSudo(cmd string) (string, error) {
	// 786 | sudo -S  dmidecode -s baseboard-manufacturer
	return RunFullCommandNoTimeOut(cmd)
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
	var isAnyError bool = false

	sys := os.DirFS("/")

	// Look for an os-release file.
	b, err := fs.ReadFile(sys, osrelease.Path)
	if err != nil {
		isAnyError = true
		fmt.Printf("error:%#v", err)
	}
	m, err := osrelease.Parse(ctx, bytes.NewReader(b))
	if err != nil {
		isAnyError = true
		fmt.Printf("error:%#v", err)
	}

	if isAnyError {
		// Mac Detection Method
		result, err := exec.Command("sw_vers").CombinedOutput()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if strings.Contains(strings.ToLower(string(result)), strings.ToLower("Mac")) {
			m["ID_LIKE"] = "darwin"
			return &m
		}
	}

	return &m
}

func MakeMapOfLines(input string) map[string]string {
	lines := strings.Split(input, "\n")

	var infoMap = make(map[string]string)

	for _, line := range lines {
		if strings.Contains(line, ":") {
			mapKeyVal := strings.Split(line, ":")
			infoMap[strings.TrimSpace(mapKeyVal[0])] = strings.TrimSpace(mapKeyVal[1])
		}
	}

	return infoMap
}

func RemoveCSVExtras(s string) string {
	s = strings.ReplaceAll(s, "[", "")
	s = strings.ReplaceAll(s, "]", "")
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.ReplaceAll(s, ",", "")
	// s = strings.TrimSpace(s)
	return s
}

func RemoveCurlyBraces(s string) string {
	s = strings.ReplaceAll(s, "{", "")
	// s = strings.ReplaceAll(s, "}", "")
	// s = strings.ReplaceAll(s, ",", "")
	return s
}

func SkipThisStr(s string) bool {
	if strings.Contains(s, "}") {
		return true
	} else {
		return false
	}
}

func DoesStringContainAlphaNumeric(str string) bool {
	for _, charVariable := range str {
		if charVariable == '*' || (charVariable >= 'a' && charVariable <= 'z') || (charVariable >= 'A' && charVariable <= 'Z') || (charVariable >= '0' && '9' >= charVariable) {
			return true
		}
	}
	return false
}

func IsStringInSlice(ss *[]string, substr string) (int, string) {

	for i, s := range *ss {
		if strings.Contains(s, substr) {
			return i, s
		}
	}

	return -1, ""
}
