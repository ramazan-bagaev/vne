package main

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

type OS interface {
	UserManager
	GetShellPath(user string) string
}

func GetOS() OS {
	if runtime.GOOS == "darwin" {
		return mac{}
	}

	log.Fatal("this os is not implemented")
	panic(1)
}

type mac struct {
}

type windows struct {
}

type linux struct {
}

func (m mac) GetShellPath(user string) string {
	cmd := exec.Command("dscl", ".", "-read", "/Users/"+user, "UserShell")

	out, err := cmd.Output()

	Check(err)

	return strings.TrimSpace(string(out))
}