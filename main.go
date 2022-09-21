package main

import (
	"os"
)

func main() {
	cmd := ParseCommand(os.Args)

	switch cmd.Cmd {
	case "create":
		var conf *Config
		if cmd.ConfigPath != "" {
			conf = CreateConfig(cmd.ConfigPath)
		} else {
			conf = CreateEmptyConfig()
		}

		GetOS().Create(cmd.User, conf)
	case "delete":
		GetOS().Delete(cmd.User)
	case "load":
		env := RetrieveEnv(cmd.User)

		env.LoadToVNEConfig(cmd.ConfigPath)
	case "unload":
		env := RetrieveEnv(cmd.User)

		env.UnloadToUser(cmd.ConfigPath)
	}
}
