package main

import "fmt"

func Gcd(a,b int)int{
	if a<0||b<0{
		return 0
	}

	for b!=0{
		temp:=b
		b=a%b
		a=temp
	}
	return a
}

func main(){
	fmt.Println(Gcd(4,18))
}