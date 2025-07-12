package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FlushInputBuffer() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
func PromptLine(label string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func PromptCSV(label string) []string {
	raw := PromptLine(label)
	parts := strings.Split(raw, ",") // splits only on comma
	var result []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
