package main

import (
	"fmt"
	"os"
)

func foo() {
	// funcの処理が終わってから実行してください
	defer fmt.Println("world foo")

	fmt.Println("Hello foo")

}

func main() {
	/*
		defer fmt.Println("world")

		foo()

		defer fmt.Println("world")

		fmt.Println("hello")
	*/

	/*
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")
	 */
	// defer は 降順 で 実行される
	/*
		run
		success
		3
		2
		1
	*/

	file, _ :=os.Open("./lesson.go")
	defer file.Close()
	// fileを開く際に、byte文字列を作る必要がある
	// 文字列を扱う配列をばいと配列と呼ぶ。
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
