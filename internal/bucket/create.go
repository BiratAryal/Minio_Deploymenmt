package bucket

import (
	"fmt"
	"miniolearn/config"
	"miniolearn/internal/mcwrapper"
	"miniolearn/internal/prompt"
	"miniolearn/internal/styles"
	"miniolearn/validate"
	"strings"
)

func BucketCreate() {
	lists := Bucketlists()
	prompt.PrintList("🪣 Existing Buckets", lists)

	buckets := prompt.PromptCSV("Enter the bucket name:")
	for _, b := range buckets {
		b = strings.ToLower(b)
		if err := validate.BucketValidate(b, lists); err != nil {
			fmt.Println(styles.ErrorStyle.Render(fmt.Sprintf("❌ Invalid Bucket '%s': %v\n", b, err)))
			continue
		}

		// Create bucket using mc
		output, err := mcwrapper.Run("mb", config.MinioAlias+"/"+b)
		if err != nil {
			fmt.Printf("❌ Failed to create bucket '%s': %v\n", b, err)
			fmt.Println("Command output:", output)
			continue
		}
		fmt.Println(styles.SuccessStyle.Render(fmt.Sprintf("✅ Bucket '%s' created successfully.", b)))
		fmt.Println()
	}
}
