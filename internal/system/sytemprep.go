package system

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	Basedir string
	Confdir string
	Logdir  string
	Bindir  string
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
		Basedir = filepath.Join(home, "AppData", "Local", "Minioadmin")
		Confdir = filepath.Join(Basedir, "conf")
		Logdir = filepath.Join(Basedir, "logs")
		Bindir = filepath.Join(Basedir, "bin")
		os.MkdirAll(Confdir, 0755)
		os.MkdirAll(Logdir, 0755)
		os.MkdirAll(Bindir, 0755)
		switch architecture {
		case "amd64", "arm64":
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
			fmt.Println("Directories created are: ", Basedir, Confdir, Logdir, Bindir)
		case "386", "arm":
			fmt.Println("Directories created are: ", Basedir, Confdir, Logdir, Bindir)
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
		}

	case "linux":
		Basedir = filepath.Join(home, ".minioadmin")
		Confdir = filepath.Join(Basedir, "conf")
		Logdir = filepath.Join(Basedir, "logs")
		Bindir = filepath.Join(Basedir, "bin")
		os.MkdirAll(Confdir, 0755)
		os.MkdirAll(Logdir, 0755)
		os.MkdirAll(Bindir, 0755)
		switch architecture {
		case "amd64", "arm64":
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
			fmt.Println("Directories created are: ", Basedir, Confdir, Logdir, Bindir)
		case "arm", "386":
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
			fmt.Println("Directories created are: ", Basedir, Confdir, Logdir, Bindir)
		}

	case "darwin":
		Basedir = filepath.Join(home, ".minioadmin")
		Confdir = filepath.Join(Basedir, "conf")
		Logdir = filepath.Join(Basedir, "logs")
		Bindir = filepath.Join(Basedir, "bin")
		os.MkdirAll(Confdir, 0755)
		os.MkdirAll(Logdir, 0755)
		os.MkdirAll(Bindir, 0755)
		switch architecture {
		case "amd64":
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
			fmt.Println("Directories created are: ", Basedir, Confdir, Logdir, Bindir)
		case "arm64":
			fmt.Println("You are using ", operatingsystem, "with Architecture", architecture)
			fmt.Println("Directories created are: ", Basedir, Confdir, Logdir, Bindir)
		}
	}
}
