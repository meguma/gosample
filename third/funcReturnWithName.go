package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	n, err := div(10, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

// 関数：名前付き戻り値
func div(i, j int) (result int, err error) {
	if j == 0 {
		err = errors.New("divied by zero")
		return // return 0, err と同じ
	}
	result = i / j
	return // return result, nil と同じ
}
