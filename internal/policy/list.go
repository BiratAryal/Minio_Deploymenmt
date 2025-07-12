package policy

import (
	"fmt"
	"miniolearn/config"
	"miniolearn/internal/mcwrapper"
	"miniolearn/internal/utils"
	"strings"
)

func GetPolicyList() []string {
	utils.ClearScreen()
	output, err := mcwrapper.Admin("policy", "list", config.MinioAlias)
	if err != nil {
		fmt.Println("Error listing policies:", err)
		return nil
	}

	lines := strings.Split(string(output), "\n")
	var policies []string

	for _, line := range lines {
		policy := strings.TrimSpace(line)
		if policy != "" {
			policies = append(policies, policy)
		}
	}
	return policies
}
