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

	n, m := 1, 2

	// 変数に代入されているポインタを入れ替えても、
	// 中の値が書き換わっているわけではないので、
	// 元の値(n, m)が書き換わることはない。
	fmt.Println(n, m)
	Swap2(&n, &m)
	fmt.Println(n, m)
	fmt.Println(&n, &m)

	fmt.Println("------------------")

	// 変数自体にポインタを入れてあげる
	n_ptr, m_ptr := &n, &m
	fmt.Println(n, m)
	Swap2(n_ptr, m_ptr)
	fmt.Println(*n_ptr, *m_ptr)
}

func Swap2(n_ptr, m_ptr *int) {
	buf := n_ptr
	//	fmt.Println(n_ptr, buf)
	//	fmt.Println(*n_ptr, *buf)
	n_ptr = m_ptr
	//	fmt.Println(n_ptr, m_ptr)
	//	fmt.Println(*n_ptr, *m_ptr)
	m_ptr = buf
	//	fmt.Println(n_ptr, m_ptr)
	//	fmt.Println(*n_ptr, *m_ptr)
}
