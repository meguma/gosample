package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	/*
		// エラーを返す関数
		file, err := os.Open("hello.go")
		if err != nil {
			// エラー処理
			// returnなどで処理を別の場所に抜ける
		}
		// fileを用いた処理
	*/
	n, err := div(10, 0)
	if err != nil {
		// エラーを出力しプログラムを終了する
		log.Fatal(err)
	}
	fmt.Println(n)
}

// 自作のエラー
func div(i, j int) (int, error) {
	if j == 0 {
		// 自作のエラーを返す
		return 0, errors.New("divied by zero")
	}
	return i / j, nil
}
