package main

import "fmt"

func one(x *int) {
	*x = 1
}
func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)
	// アドレスは？
	fmt.Println(&n)
	// アドレスの中身は？
	fmt.Println(*&n)
	/*
		fmt.Println(n)

		fmt.Println(&n)

		var p *int = &n

		fmt.Println(p)
		// 参照メモリの中身を参照するときは *
		fmt.Println(*p)
	*/
}
