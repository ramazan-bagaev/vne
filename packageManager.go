package main

import (
	"log"
	"os"
	"os/exec"
)

func PackageManagers() []string {
	return []string{"apt-get", "yam", "pacman", "brew", "apk"}
}

func UploadLackingTools(env *Env) {
	pms := GetAvailableToolsFromList(env, PackageManagers())

	var pm string
	if len(pms) == 0 {
		pm = LoadPM(env)
	} else {
		pm = pms[0]
	}

	userTools := GetUserTools(env)

	for _, vneTool := range env.Tools {
		isThere := false
		for _, userTool := range userTools {
			if vneTool == userTool {
				isThere = true
			}
		}

		if !isThere {
			loadTool(vneTool, pm)
		}
	}
}

func LoadPM(env *Env) string {
	cmd := exec.Command("mkdir", env.Home()+"/.vne-tmp")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	Check(cmd.Run())

	cmd = exec.Command("curl", "-o", env.Home()+"/.vne-tmp/install.sh", "https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	exec.Command("chmod", "+x", env.Home()+"/.vne-tmp/install.sh").Run()

	cmd = exec.Command(env.Home() + "/.vne-tmp/./install.sh")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	Check(cmd.Run())

	return "brew"
}

func loadTool(tool string, pm string) {
	log.Println(os.Getenv("PATH"))
	log.Println(os.Getenv("SHELL"))
	cmd := exec.Command(pm, "install", tool)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	Check(err)
}
