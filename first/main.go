package main

import (
	"fmt"
)

var message string = "hello world"

var foo, bar, buz string = "foo", "bar", "buz"

var (
	a string = "aaa"
	b        = "bbb"
	c        = "ccc"
)

func main() {

	// どちらの書き方も同じ意味
	// var message string = "hello world"
	message := "hello world"
	fmt.Println(message)

	//	const Hello string = "hello"
	//	Hello = "bye" // cannot assign to Hello

	var i int      // i はゼロ値で初期化
	fmt.Println(i) // 0
}
