package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func loggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func main() {
	loggingSettings("test.log")
	_, err := os.Open("dsffhgjhkj")
	if err != nil {
		log.Fatalln("Exit")
	}
	log.Println("logging")
	log.Printf("%T %v", "test", "test")

	// Fatalを使うと処理が終了する！
	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!")

	fmt.Println("ok!")
}
