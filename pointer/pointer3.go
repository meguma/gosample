package main

import (
	"fmt"
)

func Swap(x, y *int) {
	buf := *x
	*x = *y
	*y = buf
}

func main() {
	// 値を代入
	var myHouse int = 123
	// ポインタを代入
	var myHouse_ptr *int = &myHouse

	// 値を出力
	fmt.Println(myHouse)
	// ポインタ（メモリのアドレス（番地））を出力
	fmt.Println(myHouse_ptr)
	// ポインタの指す値を出力
	fmt.Println(*myHouse_ptr)
	// 値からポインタを取得して出力
	fmt.Println(&myHouse)

	// ポインタ経由で値の代入
	*myHouse_ptr = 456
	fmt.Println(myHouse)

	x, y := 1, 2
	Swap(&x, &y)
	fmt.Println(x, y)
}
