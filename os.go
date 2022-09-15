package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type OS interface {
	UserManager
	GetShellPath(user string) string
	GetUsersDir() string
}

func GetOS() OS {
	if runtime.GOOS == "darwin" {
		return mac{}
	}

	if runtime.GOOS == "linux" {
		return linux{}
	}

	log.Fatal("this os is not implemented: " + runtime.GOOS)
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

func (m mac) GetUsersDir() string {
	return "/Users"
}

func (l linux) GetShellPath(user string) string {
	return os.Getenv("SHELL")
}

func (l linux) GetUsersDir() string {
	return "/home"
}
