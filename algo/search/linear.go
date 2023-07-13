package search

func Search(haystack []interface{}, needle interface{}) int {
	for i, h := range haystack {
		if h == needle {
			return i
		}
	}
	return -1
}
