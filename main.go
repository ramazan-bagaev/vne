package main

import (
	"os"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "-create":
		GetOS().Create(os.Args[2])
	case "-delete":
		GetOS().Delete(os.Args[2])
	case "-load":
		env := CreateEnv(os.Args[2])
		env.LoadToVNEConfig()
		env.PrintEnvs()
		env.PrintTools()
	case "-unload":
		env := CreateEnv(os.Args[2])
		env.UnloadToUser()
		env.PrintEnvs()
	}
}
