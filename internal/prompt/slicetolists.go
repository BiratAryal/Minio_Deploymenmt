package prompt

import (
	"fmt"
	"miniolearn/internal/styles"
)

func PrintList(title string, lists []string) {
	// Print numbered user list
	fmt.Println(styles.BoxStyle.Render(title))
	for i, list := range lists {
		line := fmt.Sprintf("%2d) %s", i+1, list)
		fmt.Println(styles.OrderedItemStyle.Render(line))
	}
}
