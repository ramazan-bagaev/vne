package main

import (
	"log"
	"os/exec"
	"runtime"
)

type UserManager interface {
	Create(user string)
	Delete(user string)
}

func GetUserManager() UserManager {
	if runtime.GOOS == "darwin" {
		return macOS{}
	}

	log.Fatal("this os is not implemented")
	panic(1)
}

type macOS struct {
}

func (mc macOS) Create(user string) {
	log.Printf("start creating user: %s", user)

	s1 := "sysadminctl"
	s2 := "-addUser"
	s3 := "-password"
	s4 := "pass"

	cmd := exec.Command(s1, s2, user, s3, s4)
	out, err := cmd.Output()
	Check(err)
	log.Print(string(out))
}

func (mc macOS) Delete(user string) {
	log.Printf("start deleting user: %s", user)
	s1 := "sysadminctl"
	s2 := "-deleteUser"

	cmd := exec.Command(s1, s2, user)
	out, err := cmd.Output()
	Check(err)
	log.Print(string(out))
}
