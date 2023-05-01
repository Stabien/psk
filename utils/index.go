package utils

func IsInMap(strMap map[string]string, str string) bool {
	for key, value := range strMap {
		if str == key || str == value {
			return true
		}
	}
	return false
}
