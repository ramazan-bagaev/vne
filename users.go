package main

import (
	"fmt"
	"os/exec"
)

func CreateOSUser(user string) {
	fmt.Println("start creating user: " + user)

	s1 := "sysadminctl"
	s2 := "-addUser"
	s3 := "-password"
	s4 := "pass"
	
	cmd := exec.Command(s1, s2, user, s3, s4)
	out, err := cmd.Output()
	Check(err)
	fmt.Println(string(out))
}

func DeleteOSUser(user string) {
	fmt.Println("start deleting user: " + user)
	s1 := "sysadminctl"
	s2 := "-deleteUser"

	cmd := exec.Command(s1, s2, user)
	out, err := cmd.Output()
	check(err)
	fmt.Println(string(out))
}