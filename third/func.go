package main

import (
	"fmt"
)

func main() {

	good() // good

	sum(1, 2) // 3

	n := sumReturn(1, 2)
	fmt.Println(n) // 3

	x, y := 3, 4
	x, y = swap(x, y)
	fmt.Println(x, y) // 4, 3

	/*
		x = swap(x, y) // コンパイルエラー
	*/

	x, _ = swap(x, y)
	fmt.Println(x) // 3

	swap(x, y) // コンパイル、実行ともに可能
}

// 関数：引数のない場合
func good() {
	fmt.Println("good")
}

// 関数：引数がある場合
func sum(i, j int) {
	fmt.Println(i + j)
}

// 関数：戻り値がある場合：引数の次に指定
func sumReturn(i, j int) int {
	return i + j
}

// Goの関数は複数の値を返せる
func swap(i, j int) (int, int) {
	return j, i
}
