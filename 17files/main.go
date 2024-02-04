package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("this is a class on files")
	content := "this text will go into some file by end of this class (I am not clear though but keep trust)"

	file, err := os.Create("./myFile.txt")
	checkNilErr(err)
	// if err != nil {
	// 	panic(err)
	// }

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", length)
	file.Close()

	readFile("./myFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file is: \n", databyte)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
