package user

import (
	"fmt"
	"miniolearn/config"
	"miniolearn/internal/mcwrapper"
	"miniolearn/internal/prompt"
	"miniolearn/internal/styles"
	"miniolearn/validate"
)

func CreateUser() {
	lists := GetUserList()
	prompt.PrintList("👤Username", lists)
	usernames := prompt.PromptCSV("Enter comma-separated usernames: ")
	for _, u := range usernames {
		if err := validate.CheckUsername(u); err != nil {
			fmt.Printf("Invalid username '%s': %v\n", u, err)
			continue
		}

		var password string
		for {
			password = prompt.PromptLine(fmt.Sprintf("Enter password for '%s': ", u))
			if err := validate.CheckPassword(password); err != nil {
				fmt.Println("Error while creating password", err)
			} else {
				break
			}
			confirmPassword := prompt.PromptLine(fmt.Sprintf("Confirm password for '%s': ", u))
			if confirmPassword != password {
				fmt.Println("Passwords do not match.")
				continue
			}
			break // success
		}
		fmt.Println(" -", u)
		output, err := mcwrapper.Admin("user", "add", config.MinioAlias, u, password)
		if err != nil {
			fmt.Printf("Failed to create user '%s': %v\n", u, err)
			fmt.Println("Command output:", output)
			continue
		}
		fmt.Println(styles.SuccessStyle.Render(fmt.Sprintf("New user %s is created \n", u)))
		fmt.Println()
	}
}
