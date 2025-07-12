package user

import (
	"fmt"
	"miniolearn/config"
	"miniolearn/internal/mcwrapper"
	"miniolearn/internal/prompt"
)

func UserDisable() {
	lists := GetUserList()
	prompt.PrintList("👤Username", lists)
	usernames := prompt.PromptCSV("Enter username(s) seperated by , : ")
	for _, u := range usernames {
		if UserPresence(u) {
			fmt.Printf("\nUser %s is present\n", u)
			fmt.Println()
			disableuser, err := mcwrapper.Admin("user", "disable", config.MinioAlias, u)
			if err != nil {
				fmt.Printf("\nCould not disable the %s user due to error: %s\n", disableuser, err)
			} else {
				fmt.Printf("\nUser: %s is now disabled.\n", disableuser)
				fmt.Println()
			}
		} else {
			fmt.Printf("\nUser %s you entered is not present\n", u)
		}
	}
}
