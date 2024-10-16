package main

import "C"
import (
	"fmt"

)

//export HelloWorld
func HelloWorld() {
	fmt.Println("Hello World!")
}

func main() {
	fmt.Println("Hello World!")
}

/*
// Instead of DllMain, Golang calls it “init”. Unlike other exports, you don’t add the export comment above it.
func init() {
	fmt.Println("Hello World!")
}
*/
