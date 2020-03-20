package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//file と err を := で イニシャライズ
	file, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatal("Exit")
	}
	defer file.Close()
	data := make([]byte, 100)
	//ここでも err を イニシャライズ & 上書き
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	//err := os.Chdir("test")
	// := すると エラーとなる
	//-> ./lesson.go:26:6: no new variables on left side of :=
	err = os.Chdir("test") //overrideで書き換えている
	if err != nil {
		log.Fatalln("Error")
	}
	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}
}
