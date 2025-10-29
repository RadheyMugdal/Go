package auth

import (
	"fmt"
)

// In go if we don't give function name starting with  ex. login()
// capital character than it will be only
// accessible within package
// it can't be accessible in other package after importing it

// We can make this function public and make it accessible outside
// Making first character of function name in capital like ex Login

// Same we need to do with struct and varibles we must need to make first char capital
// to access it publically
// see user/user.go

func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username, password)

}
