package common

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"machine_info_gatherer/model"
	"os"
	"strings"

	"github.com/quay/claircore/osrelease"
)

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
