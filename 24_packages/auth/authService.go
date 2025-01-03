package auth

import "fmt"

func Login(username, password string) string {
	return fmt.Sprintf("Username: %s, Password: %s", username, password)
}