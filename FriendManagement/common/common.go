package common

// for this function to check a list of friends that is exsting a friend or not
func CheckExsiting(lst []string, s string) bool {
	for _, v := range lst {
		if s == v {
			return true
		}
	}
	return false
}
