package utils

func CheckElementExistingInArray(e interface{}, slice []interface{}) (ok bool) {
	ok = false
	for _, value := range slice {
		if e == value {
			ok = true
			return
		}
	}
	return
}