package main

import (
	"fmt"
)

func main() {
	//	PracticeIf()
	//	PracticeFor()
	PracticeSwitch()
	//	PracticeFallthrough()
}

func PracticeIf() {
	a, b := 10, 100
	if a > b {
		fmt.Println("a is larger than b")
	} else if a < b {
		fmt.Println("a is smaller than b")
	} else {
		fmt.Println("a equals b")
	}
	// 1行による波カッコ省略不可！
	// 三項演算子もなし！
}

func PracticeFor() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// While文もなし！
	// すべてfor文でかくべし！

	/********** while ***********/
	// C
	//	int n = 0;
	//	while (n < 10) {
	//		printf("n = %d\n", n);
	//		n++;
	//	}

	// Go
	n := 0
	for n < 10 {
		fmt.Printf("n = %d\n", n)
		n++
	}

	/********** 無限ループ ***********/

	// C
	//	for (;;) {
	//		doSomething();
	//	}

	// Go
	for {
		doSomething()
		break
	}
	// for文の条件部を省略することで表現できる！

	/********** break, continue ***********/

	m := 0
	for {
		m++
		if m > 10 {
			break // ループを抜ける
		}
		if m%2 == 0 {
			continue // 偶数なら次の繰り返しに移る
		}
		fmt.Println(m) // 奇数のみ表示
	}
}

func doSomething() {
	fmt.Println("doSomething")
}

func PracticeSwitch() {
	// 値での分岐
	n := 10
	switch n {
	case 15:
		fmt.Println("FizzBuzz")
	case 5, 10:
		fmt.Println("Buzz")
	case 3, 6, 9:
		fmt.Println("Fizz")
	default:
		fmt.Println(n)
	}
	// caseには単一の値だけでなく、カンマで区切った複数の値も指定できる！

	// 式での分岐
	m := 10
	switch {
	case m%15 == 0:
		fmt.Println("FizzBuzz")
	case m%5 == 0:
		fmt.Println("Buzz")
	case m%3 == 0:
		fmt.Println("Fizz")
	default:
		fmt.Println(m)
	}
}

func PracticeFizzBuzz() {
	n := 1
	for n <= 15 {
		switch n {
		case 15:
			fmt.Println("FizzBuzz")
		case 5, 10:
			fmt.Println("Buzz")
		case 3, 6, 9:
			fmt.Println("Fizz")
		default:
			fmt.Println(n)
		}
		n++
	}
}

func PracticeFallthrough() {
	n := 3
	switch n {
	case 3:
		n = n - 1
		fallthrough
	case 2:
		n = n - 1
		fallthrough
	case 1:
		n = n - 1
		fmt.Println(n) // 0
	}
}
