package common

// CheckListExisting for this function to check a list of friends that is existing a friend or not
func CheckListExisting(lst []string, s string) bool {
	for _, v := range lst {
		if s == v {
			return true
		}
	}
	return false
}
