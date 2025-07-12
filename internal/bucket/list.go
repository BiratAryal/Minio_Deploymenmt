package bucket

import (
	"fmt"
	"miniolearn/config"
	"miniolearn/internal/mcwrapper"
	"miniolearn/internal/utils"
	"strings"
)

func Bucketlists() []string {
	utils.ClearScreen()
	output, err := mcwrapper.Run("ls", config.MinioAlias)
	if err != nil {
		fmt.Println("Error while listing:", err)
	}

	lines := strings.Split(string(output), "\n")
	var buckets []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) > 0 {
			bucketWithSlash := fields[len(fields)-1]
			bucket := strings.TrimSuffix(bucketWithSlash, "/")
			buckets = append(buckets, bucket)
		}
	}
	return buckets
}
