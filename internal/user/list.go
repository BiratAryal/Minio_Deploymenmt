package user

import (
	"fmt"
	"miniolearn/config"
	"miniolearn/internal/mcwrapper"
	"miniolearn/internal/utils"
	"strings"
)

func GetUserList() []string {
	utils.ClearScreen()
	output, err := mcwrapper.Admin("user", "ls", config.MinioAlias)
	if err != nil {
		fmt.Println("Error listing users:", err)
		return nil
	}

	lines := strings.Split(string(output), "\n")
	var users []string

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			username := fields[1] // Assuming username is the 2nd column
			users = append(users, username)
		}
	}
	return users
}
