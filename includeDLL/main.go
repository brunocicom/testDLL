package main

// #cgo LDFLAGS: -L. -lmakedll
// #include "../makeDLL/dist/makedll.h"
import "C"

func main() {
	C.HelloWorld()
}
