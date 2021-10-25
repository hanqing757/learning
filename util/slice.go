package util

func InStringSlice(stack []string, needle string) bool {
	for _, s := range stack {
		if s == needle {
			return true
		}
	}

	return false
}
