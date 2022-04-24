package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	pointer uintptr // start
	length  int     // end
}

func main() {

	// 문자열이 같다면 -> 같은 곳을 쳐다보고있다.
	str := "hello"
	dump(str)      // same
	dump("hello")  // same
	dump("hello!") // different

	for i := range str {
		dump(str[0:i])
	}

	// byte는 다른 포인터 주소로 잡는다
	dump(string([]byte(str)))
	dump(string([]byte(str[0:2])))
	dump(string([]byte(str[0:4])))
}

func dump(str string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&str))
	fmt.Printf("%q : %+v\n", str, ptr)
}
