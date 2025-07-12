package bucket

import (
	"slices"
	"sort"
)

func BucketPresence(bucket string) bool {
	lists := Bucketlists()
	sort.Strings(lists)
	return slices.Contains(lists, bucket)
}
