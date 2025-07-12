package user

import (
	"slices"
	"sort"
)

func UserPresence(username string) bool {
	lists := GetUserList()
	sort.Strings(lists)
	return slices.Contains(lists, username)
}
