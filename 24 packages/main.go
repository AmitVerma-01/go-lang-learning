package main

import (
	"amitverma-01/go-lang/learning/auth"
	"amitverma-01/go-lang/learning/user"
)

func main() {
	result := auth.Login("John", "123456")
	println(result)
	result = auth.GetSession()
	println(result)

	user := user.NewUser("John", 20)
	println(user.Name)
	// color.Green(user.Name)
}
