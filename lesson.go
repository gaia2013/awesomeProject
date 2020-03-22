package main

import "fmt"

func do(i interface{}) {
	/*
		ii := i.(int) // interface型のiをinteger型にtype_assertion
		//ii := i * 2 // interfaceを乗算できない.
		ii *= 2
		fmt.Println(ii)
	*/

	/*
		ss := i.(string)
		fmt.Println(ss + "!")
	*/

	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know. %T\n", v)

	}
}

func main() {
	//var i interface{} = 10
	//do(i)
	//do("Mike")
	//do(true)
	do(10)
	do("Mike")
	do(true)

	var i int = 10
	ii := float64(10) // type conversion
	fmt.Printf("%T\n",i)
	fmt.Printf("%T\n", ii)
}
