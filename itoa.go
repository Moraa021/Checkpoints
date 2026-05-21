package main

import "fmt"

func Itoa(n int)string{
	if n==0{
		return "0"
	}

	sign:=""
	if n<0{
		sign ="-"
		n=-n
	}
	result:=""

	for n>0{
		result=string(rune(n%10+'0'))+result
		n/=10
	}
	return sign + result
}


func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}