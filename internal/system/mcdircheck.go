package system

import (
	"fmt"
	"miniolearn/config"
	initial "miniolearn/internal/initialsetup"
	"os"
	"path/filepath"
	"runtime"
)

func InitialSetup() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}
	switch runtime.GOOS {
	case "windows":
		config.DirPath = filepath.Join(home, "mc")
	case "linux":
		config.DirPath = filepath.Join(home, ".mc")
	case "darwin":
		config.DirPath = filepath.Join(home, ".mc")
	}
	info, err := os.Stat(config.DirPath)
	if os.IsNotExist(err) {
		initial.AliasSetup()
	} else if err != nil {
		fmt.Println("Error checking directory:", err)
	} else if info.IsDir() {
		fmt.Println("Directory exists:", config.DirPath)
	} else {
		fmt.Println("A file with the same name exists but it's not a directory:", config.DirPath)
	}
}
