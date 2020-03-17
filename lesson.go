package main

import "fmt"

const Pi = 3.14

const (
	Username=  "test_user"
	Password =  "test_pass"
)

const big = 9999999999999999998 + 1

func main()  {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big-1)
}