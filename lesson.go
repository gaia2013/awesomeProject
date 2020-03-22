//package main
//
//import (
//	"fmt"
//)
//
//type UserNotFound struct {
//	Username string
//}
//
//func (e *UserNotFound) Error() string {
//	return fmt.Sprintf("User not found: %v", e.Username)
//}
//func myFunc() error {
//	//return nil.
//	ok := false
//	if ok {
//		return nil
//	}
//	return &UserNotFound{Username: "mike"}
//}
//
//
//func main() {
//	e1 := &UserNotFound{"mike"}
//	e2:=&UserNotFound{"mike"}
//	fmt.Println(e1==e2)
//	if err := myFunc(); err != nil {
//		fmt.Println(err)
//		if err == e1{
//
//		}
//		if err == e2{
//
//		}
//	}
//}


package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	e := errors.New("EOF")
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		//if err == io.EOF {
		// -> 読み込みものはないよ という場合に頻出
		if err == e {
			break
		}
	}
}
