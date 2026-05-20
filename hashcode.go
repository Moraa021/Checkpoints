package main

import "fmt"

func HashCode(s string)string{
	result:=""
	n:=len(s)

	for i:=0;i<n;i++{
		code:=(int(s[i])+n)%127
		if code<32{
			code+=33
		}
		result=result+string(rune(code))
	}
	return result
}


func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}