package user

import (
	"fmt"
	"miniolearn/config"
	"miniolearn/internal/mcwrapper"
	"miniolearn/internal/prompt"
)

func UserDetails() {
	lists := GetUserList()

	prompt.PrintList("👤Username", lists)
	username := prompt.PromptCSV("Give the user name(s) for inspecting:")
	for _, u := range username {
		if UserPresence(u) {
			fmt.Println("User You selected", u)
			userdetail, err := mcwrapper.Admin("user", "info", config.MinioAlias, u)
			if err != nil {
				fmt.Printf("The User %s faced issues due to %s", userdetail, err)
			} else {
				fmt.Println(userdetail)
				fmt.Println()
			}
		} else {
			fmt.Printf("\nThe user %s is not present", u)
		}
	}
}
