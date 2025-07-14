package parsing

import (
	"encoding/json"
	"fmt"
	"miniolearn/config"
	"miniolearn/internal/firstrun"
	"miniolearn/internal/prompt"
	"os"
	"strconv"
	"strings"
)

// Structs...
type Alias struct {
	URL       string `json:"url"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	API       string `json:"api"`
	Path      string `json:"path"`
}
type AliasConfig struct {
	Version string           `json:"version"`
	Aliases map[string]Alias `json:"aliases"`
}

func ParseJson() string {
	file, err := os.ReadFile(config.DirPath + "/config.json")
	if err != nil {
		fmt.Println("❌ Error reading file:", err)
		return ""
	}

	var cfg AliasConfig
	if err := json.Unmarshal(file, &cfg); err != nil {
		fmt.Println("❌ Error parsing JSON:", err)
		return ""
	}

	defaults := map[string]bool{
		"gcs": true, "local": true, "play": true, "s3": true,
	}

	var availableAliases []string
	for name, alias := range cfg.Aliases {
		if defaults[name] {
			continue
		}

		if alias.URL == "" {
			fmt.Printf("⚠️  Alias '%s' has an empty URL.\n", name)
			fmt.Println("➡ Prompting login setup for this alias...")
			firstrun.InitialSetup()
			return ""
		}
		availableAliases = append(availableAliases, name)
	}

	if len(availableAliases) == 0 {
		fmt.Println("❌ No valid custom aliases found.")
		return ""
	}

	// Show list and prompt selection
	prompt.PrintList("Available MinIO Servers", availableAliases)

	var selectedIndex int
	for {
		input := prompt.PromptLine("Select a server number:")
		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || choice < 1 || choice > len(availableAliases) {
			fmt.Println("❌ Invalid selection. Please try again.")
			continue
		}
		selectedIndex = choice - 1
		break
	}

	selectedAlias := availableAliases[selectedIndex]
	fmt.Printf("✅ You selected: %s\n", selectedAlias)

	// Optional: set globally
	config.MinioAlias = selectedAlias

	return selectedAlias
}
