package firstrun

import (
	"fmt"
	"miniolearn/config"
	"os"
	"path/filepath"
	"runtime"
)

func Directories() {
	operatingsystem := runtime.GOOS
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

	case "linux":
		config.Basedir = filepath.Join(home, ".minioadmin")
		config.Confdir = filepath.Join(config.Basedir, "conf")
		config.Logdir = filepath.Join(config.Basedir, "logs")
		config.Bindir = filepath.Join(config.Basedir, "bin")
		os.MkdirAll(config.Confdir, 0755)
		os.MkdirAll(config.Logdir, 0755)
		os.MkdirAll(config.Bindir, 0755)

	case "darwin":
		config.Basedir = filepath.Join(home, ".minioadmin")
		config.Confdir = filepath.Join(config.Basedir, "conf")
		config.Logdir = filepath.Join(config.Basedir, "logs")
		config.Bindir = filepath.Join(config.Basedir, "bin")
		os.MkdirAll(config.Confdir, 0755)
		os.MkdirAll(config.Logdir, 0755)
		os.MkdirAll(config.Bindir, 0755)
	}
}
