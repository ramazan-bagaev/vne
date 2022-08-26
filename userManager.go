package main

import (
	"log"
	"os"
	"os/exec"
)

type UserManager interface {
	Create(user string)
	Delete(user string)
	CheckUser(user string)
}

func (mc mac) Create(user string) {
	mc.CheckUser("root")
	log.Printf("start creating user: %s", user)

	s1 := "sysadminctl"
	s2 := "-addUser"
	s3 := "-password"
	s4 := "pass"
	s5 := "-admin"

	cmd := exec.Command(s1, s2, user, s3, s4, s5)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	Check(err)
	log.Print(string(out))
}

func (mc mac) Delete(user string) {
	mc.CheckUser("root")
	log.Printf("start deleting user: %s", user)
	s1 := "sysadminctl"
	s2 := "-deleteUser"

	cmd := exec.Command(s1, s2, user)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	Check(err)
	log.Println(string(out))
}

func (mc mac) CheckUser(user string) {
	if os.Getenv("USER") != user {
		log.Fatalf("you should login as %s", user)
	}
}
