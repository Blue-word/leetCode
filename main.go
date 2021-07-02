package main

import (
	"fmt"
	"unsafe"
)

//go:noinline
func newArray() *[4]int {
	a := [4]int{1, 2, 3, 4}
	return &a
}

func main() {
	asd := asd()
	fmt.Println(asd)
	unsafe.Pointer(a, v)
}

func asd() int {
	var a int
	a = 1
	defer func() {
		a = 2
		fmt.Printf("a:%p", &a)
	}()
	fmt.Printf("a:%p", &a)
	return a
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println(x, y)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
