package system

import (
	"fmt"
	"miniolearn/config"
	"os"
	"path/filepath"
	"runtime"
)

func Directories() {
	operatingsystem := runtime.GOOS
	architecture := runtime.GOARCH
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to create directory due to error", err)
	}
	switch operatingsystem {
	case "windows":
		config.Basedir = filepath.Join(home, "AppData", "Local", "Minioadmin")
		config.Confdir = filepath.Join(config.Basedir, "conf")
		config.Logdir = filepath.Join(config.Basedir, "logs")
		config.Bindir = filepath.Join(config.Basedir, "bin")
		os.MkdirAll(config.Confdir, 0755)
		os.MkdirAll(config.Logdir, 0755)
		os.MkdirAll(config.Bindir, 0755)
		switch architecture {
		case "amd64", "arm64":
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
			fmt.Println("Directories created are: ", config.Basedir, config.Confdir, config.Logdir, config.Bindir)
		case "386", "arm":
			fmt.Println("Directories created are: ", config.Basedir, config.Confdir, config.Logdir, config.Bindir)
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
		}

	case "linux":
		config.Basedir = filepath.Join(home, ".minioadmin")
		config.Confdir = filepath.Join(config.Basedir, "conf")
		config.Logdir = filepath.Join(config.Basedir, "logs")
		config.Bindir = filepath.Join(config.Basedir, "bin")
		os.MkdirAll(config.Confdir, 0755)
		os.MkdirAll(config.Logdir, 0755)
		os.MkdirAll(config.Bindir, 0755)
		switch architecture {
		case "amd64", "arm64":
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
			fmt.Println("Directories created are: ", config.Basedir, config.Confdir, config.Logdir, config.Bindir)
		case "arm", "386":
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
			fmt.Println("Directories created are: ", config.Basedir, config.Confdir, config.Logdir, config.Bindir)
		}

	case "darwin":
		config.Basedir = filepath.Join(home, ".minioadmin")
		config.Confdir = filepath.Join(config.Basedir, "conf")
		config.Logdir = filepath.Join(config.Basedir, "logs")
		config.Bindir = filepath.Join(config.Basedir, "bin")
		os.MkdirAll(config.Confdir, 0755)
		os.MkdirAll(config.Logdir, 0755)
		os.MkdirAll(config.Bindir, 0755)
		switch architecture {
		case "amd64":
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
			fmt.Println("Directories created are: ", config.Basedir, config.Confdir, config.Logdir, config.Bindir)
		case "arm64":
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
			fmt.Println("Directories created are: ", config.Basedir, config.Confdir, config.Logdir, config.Bindir)
		}
	}
}
