package main

import (
	"fmt"
)

// 関数を変数に代入できる
var sum func(i, j int) = func(i, j int) {
	fmt.Println(i + j)
}

func main() {
	// 関数リテラル
	func(i, j int) {
		fmt.Println(i + j)
	}(2, 4)

	sum(3, 5)
}
