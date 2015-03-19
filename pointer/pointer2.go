package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1}  // Y:0 is implicit
	s = Vertex{}      // X:0 and Y:0
)

func main() {
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)

	// 値の代入
	p1 := p
	p1.X = 3
	fmt.Println(p1)
	fmt.Println(p)

	// ポインタを代入
	p2 := &p
	p2.X = 4
	fmt.Println(p2)
	fmt.Println(p)

	// ポインタを代入
	q1 := q
	q1.X = 5
	fmt.Println(q1)
	fmt.Println(q)

	// ポインタ指す値を代入
	q2 := *q
	q2.X = 6
	fmt.Println(q2)
	fmt.Println(q)

	/* error Goはポインタのポインタは取得できない。ポインタ演算もできない。
	q3 := &q
	q3.X = 7
	fmt.Println(q3)
	fmt.Println(q)
	*/
}
