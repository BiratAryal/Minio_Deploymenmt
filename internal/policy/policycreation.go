package policy

import (
	"encoding/json"
	"fmt"
	"miniolearn/config"
	"os"
)

type Outline struct {
	Version   string    `json:"Version"`
	Statement []Details `json:"Statement"`
}

type Details struct {
	Effect   string   `json:"Effect"`
	Action   []string `json:"Action"`
	Resource string   `json:"Resource"`
}

func ReadWriteDeletePolicyCreate(buckname string) string {
	objectPolicy := Details{
		Effect:   "Allow",
		Action:   []string{"s3:GetObject", "s3:PutObject", "s3:DeleteObject"},
		Resource: fmt.Sprintf("arn:aws:s3:::%s/*", buckname),
	}

	bucketPolicy := Details{
		Effect:   "Allow",
		Action:   []string{"s3:ListBucket", "s3:GetBucketLocation"},
		Resource: fmt.Sprintf("arn:aws:s3:::%s", buckname),
	}

	// Base policy
	basePolicy := Outline{
		Version:   "2012-10-17",
		Statement: []Details{bucketPolicy, objectPolicy},
	}

	// Convert to JSON
	policyJson, err := json.MarshalIndent(basePolicy, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(policyJson)
}

func ReadWritePolicyCreate(buckname string) string {
	objectPolicy := Details{
		Effect:   "Allow",
		Action:   []string{"s3:GetObject", "s3:PutObject"},
		Resource: fmt.Sprintf("arn:aws:s3:::%s/*", buckname),
	}

	bucketPolicy := Details{
		Effect:   "Allow",
		Action:   []string{"s3:ListBucket", "s3:GetBucketLocation"},
		Resource: fmt.Sprintf("arn:aws:s3:::%s", buckname),
	}

	// Base policy
	basePolicy := Outline{
		Version:   "2012-10-17",
		Statement: []Details{bucketPolicy, objectPolicy},
	}

	// Convert to JSON
	policyJson, err := json.MarshalIndent(basePolicy, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(policyJson)
}

func ReadPolicyCreate(buckname string) string {
	objectPolicy := Details{
		Effect:   "Allow",
		Action:   []string{"s3:GetObject"},
		Resource: fmt.Sprintf("arn:aws:s3:::%s/*", buckname),
	}

	bucketPolicy := Details{
		Effect:   "Allow",
		Action:   []string{"s3:ListBucket", "s3:GetBucketLocation"},
		Resource: fmt.Sprintf("arn:aws:s3:::%s", buckname),
	}

	// Base policy
	basePolicy := Outline{
		Version:   "2012-10-17",
		Statement: []Details{bucketPolicy, objectPolicy},
	}

	// Convert to JSON
	policyJson, err := json.MarshalIndent(basePolicy, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(policyJson)
}

func CreatePolicyFile(bucketName, policyType string) {
	var policy string
	switch policyType {
	case "readonly":
		policy = ReadPolicyCreate(bucketName)
	case "readwrite":
		policy = ReadWritePolicyCreate(bucketName)
	case "readwritedelete":
		policy = ReadWriteDeletePolicyCreate(bucketName)
	default:
		fmt.Println("Unknown policy type:", policyType)
		return
	}

	filePath := fmt.Sprintf("%s/%s-%s.conf", config.Confdir, policyType, bucketName)
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	n, err := f.WriteString(policy)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		f.Close()
		return
	}

	fmt.Printf("%d bytes written successfully to %s\n", n, filePath)

	if err := f.Close(); err != nil {
		fmt.Println("Error closing file:", err)
	}
}
