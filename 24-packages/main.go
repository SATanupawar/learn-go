package main

import (
	"example.com/myapp/auth"
	"fmt"
	"example.com/myapp/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("admin", "password")

	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Username: "satyam pawar ",
		Password: "john doe ",
	}
	fmt.Println(user.Username , user.Password)

	color.Red("user.Username")

}
