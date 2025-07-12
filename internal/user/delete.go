package user

import (
	"fmt"
	"miniolearn/config"
	"miniolearn/internal/mcwrapper"
	"miniolearn/internal/prompt"
	"miniolearn/internal/styles"
)

func UserDelete() {
	lists := GetUserList()
	prompt.PrintList("👤Username", lists)
	usernames := prompt.PromptCSV("Enter username(s) seperated by , : ")
	for _, u := range usernames {
		if UserPresence(u) {
			fmt.Printf("\nUser %s is present\n", u)
			fmt.Println()
			removeduser, err := mcwrapper.Admin("user", "remove", config.MinioAlias, u)
			if err != nil {
				fmt.Printf("\nCould not remove the %s user due to error: %s\n", removeduser, err)
			} else {
				fmt.Printf("\nUser: %s is now removed.\n", removeduser)
				fmt.Println()
			}
		} else {
			fmt.Println(styles.ErrorStyle.Render(fmt.Sprintln("User", u, "you entered is not present")))
		}
	}
	lists = GetUserList()
	prompt.PrintList("👤Username", lists)
}
