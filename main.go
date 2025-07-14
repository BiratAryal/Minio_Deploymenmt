package main

import (
	"fmt"
	"miniolearn/bin"
	"miniolearn/bubbletea"
	"miniolearn/config"
	"miniolearn/internal/firstrun"
	"miniolearn/internal/parsing"
	"miniolearn/internal/prompt"
	"miniolearn/internal/system"
	"miniolearn/internal/utils"
)

func main() {
	firstrun.Directories()
	firstrun.McDirCheck()
	bin.ExtractMcBinary()
	system.OwnerBanner()
	for {
		choice := prompt.PromptLine("🔁 Print Main Meanu? [Y/N] ")
		// config.MinioAlias = "privateminio"

		if len(choice) == 0 {
			fmt.Println("Please enter a valid choice.")
			continue
		}
		switch choice[0] {
		case 'Y', 'y':
			utils.ClearScreen()
			if config.MinioAlias == "" {
				parsing.ParseJson()
			}
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
