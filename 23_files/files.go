package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read file
	f, err := os.Open("./hello.txt")
	// error hanndling in go is vey important
	// in go we return error instead of throwing it
	if err != nil {
		//log the error
		// we can make out code in panic mode  which will not execute anything after this error
		panic(err)
	}

	// f is File Object
	fileInfo, err := f.Stat()
	fmt.Println("file name:", fileInfo.Name()) // this will give file name
	fmt.Println("File or folder:", fileInfo.IsDir())
	fmt.Println("Size of file:", fileInfo.Size())
	fmt.Println("file permission:", fileInfo.Mode())
	fmt.Println("file modified at:", fileInfo.ModTime())

	// how to read it
	filer, errorr := os.Open("./hello.txt")

	if err != nil {
		panic(errorr)
	}
	defer f.Close() // Close closes the [File], rendering it unusable for I/O.

	// we read file using buffer
	buf := make([]byte, 10)   // second argument is size of bufffer it will read only this byte size data from file and ignores other
	d, err := filer.Read(buf) // we pass buffer where we want to read in read
	// it returns number of bytes read
	// and error
	if err != nil {
		panic(err)
	}
	println("data", d, string(buf)) // converting buffer to string

	// Read file simplest way to read

	fsimp, errsimp := os.ReadFile("./hello.txt") // this will return buffer and error if any
	// this read file loads all content at one time to memory
	// if we have big files this is not good way of doing it
	if errsimp != nil {
		panic(errsimp)
	}

	fmt.Println(string(fsimp))

	// How to read Folders
	dir, errdir := os.Open(".")
	if errdir != nil {
		panic(errdir)
	}
	defer dir.Close()
	fileInfodir, errfileInfo := dir.ReadDir(1) // it reads directory we meed to pass number of entry we want to see
	// if we put 1 it will read one file/folder
	// if we put 2 it will read two file/folder
	// if we put anything in nagetive it will give all the files folder available there
	if errfileInfo != nil {
		panic(errfileInfo)
	}

	for _, fi := range fileInfodir {
		fmt.Println(fi.Name()) //
	}

	// How to create file and write on it

	createdfile, errorCreatedFile := os.Create("example.txt")
	if errorCreatedFile != nil {
		panic(errorCreatedFile)
	}
	defer createdfile.Close()
	f.WriteString("hi go ") // this will write string if we do write() it needs to pass buffer

	bytes := []byte("hello go lang")
	bytefile, errorBytefile := os.Create("another.txt")
	if errorBytefile != nil {
		panic(errorBytefile)
	}
	defer bytefile.Close()
	bytefile.Write(bytes)

	// read and write to another file ( streaming fashion )
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	// how we use streaming fashion
	// that can be done using buffio

	// first we create reader

	reader := bufio.NewReader(sourceFile) // source file is also reder interface
	writer := bufio.NewWriter(destFile)

	for {
		// we are reading byte by byte
		b, err := reader.ReadByte()
		// this error also comes for end of file like if it doesn't have
		// anything to read then it will return error end of file
		// in that case we doesn't need to panic so we write login accordingly
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		errWrite := writer.WriteByte(b)
		if err != nil {
			panic(errWrite)
		}
	}

	// we than need to flush writer
	writer.Flush()
	fmt.Println("Written to new file successfully")

	// we also have copy function in file that can also be used

	// delete a file

	delerr := os.Remove("example2.txt")
	if err != nil {
		panic(delerr)
	}
	fmt.Println("File deleted successfully")
}
