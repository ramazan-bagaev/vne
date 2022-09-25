package main

import (
	"os"
)

type CliCommand struct {
	Cmd        string
	User       string
	ConfigPath string
}

func ParseCommand(Args []string) *CliCommand {
	if len(Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	cliCmd := CliCommand{}

	cmd := Args[1]
	ensureCommand(cmd)

	cliCmd.Cmd = cmd

	if cmd == "create" || cmd == "delete" {
		cliCmd.User = "vne-user"
		cliCmd.ConfigPath = ""
	} else {
		cliCmd.User = os.Getenv("USER")
		cliCmd.ConfigPath = os.Getenv("HOME") + "/.vne"
	}

	skip := false

	for i, arg := range Args {
		if i == 0 || i == 1 {
			continue
		}

		if skip {
			skip = false
			continue
		}

		if arg == "-u" && i+1 < len(Args) {
			cliCmd.User = Args[i+1]
			skip = true
			continue
		}

		if arg == "-d" && i+1 < len(Args) {
			cliCmd.ConfigPath = Args[i+1]
			skip = true
			continue
		}

		printHelp()
	}

	return &cliCmd
}

func ensureCommand(cmd string) {
	if cmd == "create" {
		return
	}

	if cmd == "delete" {
		return
	}

	if cmd == "load" {
		return
	}

	if cmd == "unload" {
		return
	}

	printHelp()
	os.Exit(0)
}

func printHelp() {
	println("vne - user managment tool, to save most important user configurations and unpack on another machine")
	println("usage: vne command [-options]")
	println("commands:")
	println("	create		creates new os user									| possible options: -u,-d")
	println("	delete		deletes os user										| possible options: -u")
	println("	load		retrieves configuration from os user (env variables, path content)			| possible options: -d")
	println("	unload		configures os user from vne config file	(env variables, path content)			| possible options: -d")
	println("options:")
	println("	-u		user name (default \"vne-user\")")
	println("	-d		vne config path (default \"~/.vne\")")
}
