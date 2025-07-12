package user

import (
	"fmt"
	"miniolearn/config"
	"miniolearn/internal/mcwrapper"
	"miniolearn/internal/prompt"
)

func UserEnable() {
	lists := GetUserList()
	prompt.PrintList("👤Username", lists)
	usernames := prompt.PromptCSV("Enter username(s) seperated by , : ")
	for _, u := range usernames {
		if UserPresence(u) {
			fmt.Printf("\nUser %s is present\n", u)
			fmt.Println()
			enableuser, err := mcwrapper.Admin("user", "enable", config.MinioAlias, u)
			if err != nil {
				fmt.Printf("\nCould not enable the %s user due to error: %s\n", enableuser, err)
			} else {
				fmt.Printf("\nUser: %s is now enable.\n", enableuser)
				fmt.Println()
			}
		} else {
			fmt.Printf("\nUser %s you entered is not present\n", u)
		}
	}
}
