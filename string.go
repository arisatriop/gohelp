package gohelp

func EmptyInteger(count int) bool {
	return count == 0
}

func EmptyString(text string) bool {
	return len(text) == 0
}

func ContainsString(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
