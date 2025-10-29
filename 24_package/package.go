package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/radheymugdal/go/auth"
	"github.com/radheymugdal/go/user"
)

// Package = Packages are the most powerful part of the Go language.
// The purpose of a package is to design and maintain a large number of programs
// by grouping related features together into single units so that they can be easy to maintain and understand and independent of the other package programs.
// This modularity allows them to share and reuse. In Go language, every package is defined with a different name and that name is close to their functionality like "strings" package and it contains methods and functions that only related to strings.
// fmt that we use to print is also pacakge
// it is used using import

// To create package we need to initialize module
// to create module in command line we run
// go mod init  module_name
// generally we name it as github.com/username/application_name
// this will create go.mod file in current directory
// you can see here

// than we create package like we created in auth/credential.go

// in go we can also get external packages like we do in js using npm
// search go package in google
// and go to first website and search any package you want to
// install it will have cmd like go get package_name
// it will add it in go.mod file that we created

func main() {
	auth.LoginWithCredentials("radhey", "123")
	auth.Session()
	user := user.User{Name: "radhey", Email: "email"}
	fmt.Println(user.Name)
	session := auth.GetSession()
	fmt.Println(session)
	color.Red(user.Email) // this will print email in red color which is functionlity of our thrid party package

}

// use full command
// go mod tidy
// this will fix pacakge related issues
// if we are using package without doint go get package name
// and than we do go mod tidy it will autometically add that package here
