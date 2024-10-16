package main

import (
	"fmt"
	"syscall"
)

func main() {
	var mod = syscall.NewLazyDLL("..\\makeDLL\\dist\\makedll.dll")
	var proc = mod.NewProc("HelloWorld")

	ret, _, _ := proc.Call(0)
	fmt.Printf("Return: %d\n", ret)

}
