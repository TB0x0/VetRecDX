package auth

// query db
func getUserId(username string) int {
	if username == "admin" {
		return 100
	} else {
		return 101
	}
}
