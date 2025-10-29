module github.com/radheymugdal/go

go 1.25.3

// this are all dependencies used in this project
// created after we installed new pacakge
// using "go get <package_name>"

// it gives indirect because this installed pacakge also uses other packages
// so since we are not using these packages directly in our code
// go marks them as indirect
require (
	github.com/fatih/color v1.18.0
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)
