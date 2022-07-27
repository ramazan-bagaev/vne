package main

import (
	"os"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "-create":
		env := CreateEnv(os.Args[2])
		env.PrintEnvs()

	case "-delete":
		GetUserManager().Delete(os.Args[2])
	case "-load":
		env := CreateEnv(os.Args[2])
		env.LoadToVNEConfig()
		env.PrintEnvs()
	case "-unload":
		env := CreateEnv(os.Args[2])
		env.UnloadToUser()
		env.PrintEnvs()
	}
}
