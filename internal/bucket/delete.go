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

//	func BucketDelete() {
//		lists := Bucketlists()
//		prompt.PrintList("🪣  Available Buckets:", lists)
//		buckets := prompt.PromptCSV("Enter bucket(s) seperated by , : ")
//		for _, b := range buckets {
//			b = strings.ToLower(b)
//			if err := validate.BucketValidate(b, lists); err != nil {
//				fmt.Println(styles.ErrorStyle.Render(fmt.Sprintf("Bucket %s is not present", b)))
//				continue
//			}
//			output, err := mcwrapper.Run("rm", config.MinioAlias+"/"+b)
//			if err != nil {
//				fmt.Println(styles.ErrorStyle.Render(fmt.Sprintf("Failed to Remove bucket %s : %v", b, err)))
//				fmt.Println(styles.ErrorStyle.Render(fmt.Sprintln("Command Output", output)))
//				continue
//			}
//			fmt.Println(styles.SuccessStyle.Render(fmt.Sprintf("✅ Bucket '%s' Removed successfully.", b)))
//			fmt.Println()
//		}
//	}
func BucketDelete() {
	lists := Bucketlists()
	prompt.PrintList("🪣  Available Buckets:", lists)

	buckets := prompt.PromptCSV("Enter bucket(s) separated by comma:")
	for _, b := range buckets {
		b = strings.ToLower(strings.TrimSpace(b))

		// Validate existence
		if err := validate.BucketDeleteValidate(b, lists); err != nil {
			fmt.Println(styles.ErrorStyle.Render(fmt.Sprintf("❌ %v", err)))
			continue
		}

		// Ask for confirmation
		confirm := prompt.PromptLine(fmt.Sprintf("⚠️ Are you sure you want to delete '%s'? (y/n): ", b))
		confirm = strings.TrimSpace(strings.ToLower(confirm))

		if confirm != "y" {
			fmt.Println(styles.ErrorStyle.Render(fmt.Sprintf("❌ Skipped deleting '%s'.", b)))
			continue
		}

		// Delete bucket
		output, err := mcwrapper.Run("rb", config.MinioAlias+"/"+b)
		if err != nil {
			fmt.Println(styles.ErrorStyle.Render(fmt.Sprintf("❌ Failed to remove bucket '%s': %v", b, err)))
			fmt.Println(styles.ErrorStyle.Render(fmt.Sprintf("Command Output: %s", output)))
			continue
		}

		fmt.Println(styles.SuccessStyle.Render(fmt.Sprintf("✅ Bucket '%s' removed successfully.", b)))
	}
	lists = Bucketlists()
	prompt.PrintList("🪣  Available Buckets:", lists)
}
