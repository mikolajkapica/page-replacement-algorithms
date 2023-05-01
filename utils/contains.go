package utils

func Contains(array []int, element int) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}
