package main

import (
	"fmt"
	"miniolearn/bubbletea"
	"miniolearn/config"
	"miniolearn/internal/prompt"
	"miniolearn/internal/system"
	"miniolearn/internal/utils"
)

func main() {
	config.MinioAlias = "privateminio"
	system.OwnerBanner()
	for {
		choice := prompt.PromptLine("🔁 Print Main Meanu? [Y/N] ")

		if len(choice) == 0 {
			fmt.Println("Please enter a valid choice.")
			continue
		}
		switch choice[0] {
		case 'Y', 'y':
			utils.ClearScreen()
			bubbletea.BubbleCall()()
			// cmd.PrintMainMenu()
			fmt.Println()
			continue
		case 'N', 'n':
			fmt.Println("OK Enjoy your day!!")
			return
		default:
			fmt.Println("Enter only designated Characaters")
		}
	}
}
