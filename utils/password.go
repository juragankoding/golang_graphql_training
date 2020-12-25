package utils

func HashPassword(password string) (string, error) {
	return password, nil
}

func ComparePassword(password string, hash string) bool {
	return password == hash
}
