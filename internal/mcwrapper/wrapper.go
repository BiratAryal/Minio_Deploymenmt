package mcwrapper

import (
	"fmt"
	"os/exec"
	"strings"
)

// Run runs general mc commands like mc ls, mc mb, etc.
func Run(args ...string) (string, error) {
	cmd := exec.Command("mc", args...)
	output, err := cmd.CombinedOutput()
	outStr := strings.TrimSpace(string(output))
	if err != nil {
		return outStr, fmt.Errorf("mc %s failed: %w\nOutput:\n%s", strings.Join(args, " "), err, outStr)
	}
	return outStr, nil
}

// Admin runs mc admin commands like mc admin user ls, etc.
func Admin(args ...string) (string, error) {
	fullArgs := append([]string{"admin"}, args...)
	return Run(fullArgs...)
}
