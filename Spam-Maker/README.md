## Spam-Maker

### Desc 
- String, Bytes, Runes in Go

- Index
	- [ ] Byte And Character
	- [ ] Rune
	- [ ] Golang in Unicode, UTF-8
	- [ ] []String, String

## What's a String?

```
	Go의 문자열은 읽기전용 Byte조각 -> Byte를 여러개로 만든것 -> 문자열
```

- Go 소스코드는 항상 UTF-8이다.
- 문자열은 임의의 Bytes를 포함한다
- byte-level escapes가 없는 문자열 리터럴은 항상 유효한 UTF-8 시퀀스를 보유한다
- 이러한 시퀀스는 rune이라 불린다

```
	str := "🏓 🏓 🏓 🏓"
	bytes := []byte(str)
	str = string(bytes)

	fmt.Printf("%s, %d, %d\n", str, len(str), utf8.RuneCountInString(str))
	fmt.Printf("%x, %d, %d\n", bytes, len(bytes), utf8.RuneCount(bytes))

	for index, r := range str {
		fmt.Printf("%d, %x, %q\n", index, string(r), r)
	}

	// solution
	runes := []rune(str)
	fmt.Println("%s, %d\n", string(runes), len(runes))
```

### Conclusion

```
	Golang에서의 String은 UTF-8을 사용한다. 각각의 1~4 byte영역을 차지한다 (차지하는 영역이 다양한다)
	Rune에서는 모든 문자열을 4byte로 통일한다 -> rune -> int32
	=> 문자열 Byte를 통일한다
```

### Rune Example

```
	// Decode

	text := `
		slkdfnalsdknflkasdnflaksjdflkasdjflkasdf
		asdfl;kansdflkansdlfna;sldkfnalsdf
		🏓 🏓 🏓 🏓lfkalsdkfnalksdnflkadsf
		asdflaksdflkansdlfaksdfstr := "🏓 🏓 🏓 🏓
		asdlfkjasdlfksd
		🏓 🏓 🏓 🏓
	`
 
	// for i := 0; i < len(text); i++ {
	// 	fmt.Printf("%c", text[i])
	// }

	/*
				output
				slkdfnalsdknflkasdnflaksjdflkasdjflkasdf
		        asdfl;kansdflkansdlfna;sldkfnalsdf
		        ð ð ð ðlfkalsdkfnalksdnflkadsf
		        asdflaksdflkansdlfaksdfstr := "ð ð ð ð
		        asdlfkjasdlfksd
		        ð ð ð ð
	*/

	// for i := 0; i < len(text); {
	// 	// r, _ := utf8.DecodeRune([]byte(text))
	// 	r, size := utf8.DecodeLastRuneInString(text[i:])
	// 	fmt.Printf("%c", r)

	// 	i += size
	// }

	for _, r := range text {
		fmt.Printf("%c", r)
	}
```

### Immutable String

```
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

```