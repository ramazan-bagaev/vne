package main

import (
	"os"
)

func main() {
	cmd := ParseCommand(os.Args)

	switch cmd.Cmd {
	case "create":
		GetOS().Create(cmd.User)
	case "delete":
		GetOS().Delete(cmd.User)
	case "load":
		env := CreateEnv(cmd.User, cmd.ConfigPath)
		env.LoadToVNEConfig()
	case "unload":
		env := CreateEnv(cmd.User, cmd.ConfigPath)
		env.UnloadToUser()
	}
}
