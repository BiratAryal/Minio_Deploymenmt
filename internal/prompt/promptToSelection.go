package prompt

import (
	"bufio"
	"fmt"
	"miniolearn/internal/styles"
	"os"
	"strconv"
	"strings"
)

// Display list and return selected item
func PromptSelectFromList(title string, lists []string) string {
	PrintList(title, lists)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(styles.PromptStyle.Render("👉 Select a number: "))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		index, err := strconv.Atoi(input)
		if err != nil || index < 1 || index > len(lists) {
			fmt.Println("❌ Invalid selection. Try again.")
			continue
		}

		return lists[index-1]
	}
}
