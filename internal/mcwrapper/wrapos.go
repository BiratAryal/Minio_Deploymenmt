package mcwrapper

import "runtime"

func GetOS() string {
	os := runtime.GOOS
	return os
}
func GetArch() string {
	arch := runtime.GOARCH
	return arch
}
