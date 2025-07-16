package helper

import (
	"fmt"
	"miniolearn/internal/policy"
	"miniolearn/internal/styles"
)

func PolicyToBucket(bucketName, policyType string) {
	switch policyType {
	case "readonly":
		policy.CreatePolicyFile(bucketName, "readonly")
	case "readwrite":
		policy.CreatePolicyFile(bucketName, "readwrite")
	case "readwritedelete":
		policy.CreatePolicyFile(bucketName, "readwritedelete")
	case "all":
		policy.CreatePolicyFile(bucketName, "readonly")
		policy.CreatePolicyFile(bucketName, "readwrite")
		policy.CreatePolicyFile(bucketName, "readwritedelete")
	default:
		message := fmt.Sprintln("❌ Unknown policy type.")
		styles.ErrorStyle.Render(message)
	}
}

func PolicyToUser(policyName string, userName string) {
	fmt.Println("Assign policy to user.")
	message := fmt.Sprintf("The Policy Selected is: %s User selected is: %s", policyName, userName)
	fmt.Println(styles.SuccessStyle.Render(message))
}
