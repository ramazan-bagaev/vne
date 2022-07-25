package main

import (
	"fmt"
	"os"
	"os/exec"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createUser(user string) {
	fmt.Println("start creating user: " + user)

	s1 := "sysadminctl"
	s2 := "-addUser"
	s3 := "-password"
	s4 := "pass"

	cmd := exec.Command(s1, s2, user, s3, s4)
	out, err := cmd.Output()
	check(err)
	fmt.Println(string(out))
}

func deleteUser(user string) {
	fmt.Println("start deleting user: " + user)
	s1 := "sysadminctl"
	s2 := "-deleteUser"

	cmd := exec.Command(s1, s2, user)
	out, err := cmd.Output()
	check(err)
	fmt.Println(string(out))
}

func main() {
	arg := os.Args[1]

	switch arg {
	case "-create":
		env := CreateEnv(os.Args[2])
		env.Retrieve()
		env.PrintEnvs()

	case "-delete":
		DeleteOSUser(os.Args[2])
	case "-load":
		env := CreateEnv(os.Args[2])
		env.LoadToConfig()
		env.PrintEnvs()
	}
}
