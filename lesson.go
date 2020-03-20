package main

import "fmt"

func thirdPartyConnectDB() {
	panic("Unable to connect database!")
//	panic は書かないようにする。何をしたらいいのかわからない状態だから。自分のコードではそういう状況のコードは書かない
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	// panicのスタックトレースが出ずにコードを実行可能
	// ただし、recover は panic の前に書く
	thirdPartyConnectDB()
	//if err := thirdPartyConnectDB というような書き方が推奨されている。

}

func main() {
	save()
	fmt.Println("OK!")
}

/*
panic: Unable to connect database!

goroutine 1 [running]:
main.thirdPartyConnectDB(...)
        /Users/hitoshi/go/src/awesomeProject/lesson.go:6
main.save(...)
        /Users/hitoshi/go/src/awesomeProject/lesson.go:10
main.main()
        /Users/hitoshi/go/src/awesomeProject/lesson.go:14 +0x3a

*/
