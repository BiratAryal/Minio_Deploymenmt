package bin

import (
	_ "embed"
	"fmt"
	"miniolearn/internal/system"
	"os"
	"path/filepath"
	"runtime"
)

// Variable declaration for importing the executable binaries.
// Could not use parent traversal (../../) so moved to main file.
var (
	// Windows OS
	//go:embed windows/64bit/mc.exe
	McWindows64 []byte
	//go:embed windows/32bit/mc.exe
	McWindows32 []byte

	// Linux OS
	//go:embed linux/64bit/mc
	McLinux64 []byte
	//go:embed linux/32bit/mc
	McLinux32 []byte

	// Mac OS
	//go:embed darwin/amd64/mc
	McDarwinAmd64 []byte
	//go:embed darwin/arm64/mc
	McDarwinArm64 []byte
)

func BinaryFiles() ([]byte, error) {

	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64", "arm64":
			return McWindows64, nil
		case "386", "arm":
			return McWindows32, nil
		}

	case "linux":
		switch runtime.GOARCH {
		case "amd64", "arm64":
			return McLinux64, nil
		case "arm", "386":
			return McLinux32, nil
		}

	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			return McDarwinAmd64, nil
		case "arm64":
			return McDarwinArm64, nil
		}
	}
	return nil, fmt.Errorf("unsupported platform: %s/%s", runtime.GOOS, runtime.GOARCH)
}
func ExtractMcBinary() (string, error) {
	binaryData, err := BinaryFiles()
	if err != nil {
		return "", err
	}

	// Use .exe on Windows
	filename := "mc"
	if runtime.GOOS == "windows" {
		filename = "mc.exe"
	}

	// Temp file location
	tempDir := system.Bindir
	fullPath := filepath.Join(tempDir, filename)

	// Write binary
	if err := os.WriteFile(fullPath, binaryData, 0755); err != nil {
		return "", fmt.Errorf("failed to write binary: %w", err)
	}

	return fullPath, nil
}
