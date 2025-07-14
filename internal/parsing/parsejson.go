package parsing

import (
	"encoding/json"
	"fmt"
	"miniolearn/config"
	"os"
)

// Structs to match your JSON structure
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

func ParseJson() {
	// Open the JSON file
	file, err := os.ReadFile(config.DirPath + "/config.json")
	if err != nil {
		fmt.Println("❌ Error reading file:", err)
		return
	}

	// Parse the JSON into the struct
	var config AliasConfig
	if err := json.Unmarshal(file, &config); err != nil {
		fmt.Println("❌ Error parsing JSON:", err)
		return
	}

	// Define default aliases to ignore
	defaults := map[string]bool{
		"gcs":   true,
		"local": true,
		"play":  true,
		"s3":    true,
	}

	// Print only non-default (custom) aliases
	fmt.Println("Custom Aliases:")
	for name, alias := range config.Aliases {
		if !defaults[name] {
			fmt.Printf("- %s => %s\n", name, alias.URL)
		}
	}
}
