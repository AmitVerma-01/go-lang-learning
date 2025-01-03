package auth

import "fmt"

func GetSession() string {
	username := "John"
	password := "123456"
	return fmt.Sprintf("Username: %s, Password: %s", username, password)
}
