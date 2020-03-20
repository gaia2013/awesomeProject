package main

import "fmt"

func main() {
	// Q1の回答
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	a := l[0]
	for i := 1; i < len(l); i++ {
		if a > l[i] {
			a = l[i]
		}
	}
	fmt.Println(a)

	// Q2の回答
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	sum := 0
	for _,v := range m{
		sum += v
	}
	fmt.Println(sum)
}
