package main

import (
	"log"
	"os"
	"os/exec"
)

type UserManager interface {
	Create(user string, conf *Config)
	Delete(user string)
	CheckUser(user string)
}

func (mc mac) Create(user string, conf *Config) {
	mc.CheckUser("root")
	log.Printf("start creating user: %s", user)

	s1 := "sysadminctl"
	s2 := "-addUser"
	s3 := "-password"
	s4 := "pass"
	s5 := "-admin"
	s6 := "-shell"

	cmd := exec.Command(s1, s2, user, s3, s4, s5, s6, conf.GetShell())
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

func (l linux) Create(user string, conf *Config) {
	l.CheckUser("root")
	log.Printf("start creating user: %s", user)
	s1 := "useradd"
	s2 := "-m"
	s3 := "-s"
	s4 := "/bin/bash"

	cmd := exec.Command(s1, s2, s3, s4, user)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	Check(err)
	log.Println(string(out))

	log.Println("setting password..")
	cmd = exec.Command("/bin/bash", "-c", "echo "+user+":pass | chpasswd")
	cmd.Stderr = os.Stderr

	out, err = cmd.Output()
	Check(err)
	log.Println(string(out))
}

func (l linux) Delete(user string) {
	l.CheckUser("root")
	log.Printf("start deleting user: %s", user)
	s1 := "userdel"
	s2 := "-r"

	cmd := exec.Command(s1, s2, user)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	Check(err)
	log.Println(string(out))
}

func (l linux) CheckUser(user string) {
	if os.Getenv("USER") != user {
		log.Fatalf("you should login as %s", user)
	}
}
