package initial

import (
	"fmt"
	"miniolearn/internal/mcwrapper"
	"miniolearn/internal/prompt"
	"miniolearn/internal/styles"
	"miniolearn/validate"
	"strings"
)

var (
	ServerAlias []string
)

func AliasSetup() {
	ServerAlias = prompt.PromptCSV("Enter the minio server alias(es) separated by ',' : ")
	for _, a := range ServerAlias {
		if err := validate.CheckUsername(a); err != nil {
			fmt.Printf("Invalid Alias Name '%s': %v\n", a, err)
			continue
		}

		var ServerIP, ServerPort, ServerProtocol string

		for {
			ServerPort = prompt.PromptLine(fmt.Sprintf("Enter the port for '%s':", a))
			ServerIP = prompt.PromptLine(fmt.Sprintf("Enter IP for '%s':", a))
			ServerProtocol = prompt.PromptLine(fmt.Sprintf("Enter the protocol [HTTP/HTTPS] for '%s':", a))
			ServerProtocol = strings.ToLower(ServerProtocol)
			if ServerProtocol != "http" && ServerProtocol != "https" {
				fmt.Println("❌ Invalid protocol. Use 'http' or 'https'")
				continue
			}
			if err := validate.CheckServer(ServerIP, ServerPort); err != nil {
				fmt.Println("❌ Could not reach", ServerIP+":"+ServerPort, "-", err)
				continue // retry
			}
			break // connection successful
		}

		username := prompt.PromptLine(fmt.Sprintf("Enter Username for '%s':", a))
		password := prompt.PromptLine(fmt.Sprintf("Enter Password for '%s':", a))
		endpoint := fmt.Sprintf("%s://%s:%s", ServerProtocol, ServerIP, ServerPort)

		fmt.Println(" -", a)
		output, err := mcwrapper.Admin("alias", "set", a, endpoint, username, password)
		if err != nil {
			fmt.Printf("Failed to authenticate '%s': %v\n", a, err)
			fmt.Println("Command output:", output)
			continue
		}
		fmt.Println(styles.SuccessStyle.Render(fmt.Sprintf("New Alias %s is created \n", a)))
		fmt.Println()
	}
}
