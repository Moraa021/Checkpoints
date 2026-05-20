package main

import "fmt"

func IsCapitalized(s string)bool{
	if s==""{
		return false
	}
	result:=true
	for _,r:=range s{
		if r==' '{
			result= true
		}else if result{
			if r>='a'&&r<='z'{
				return false
			}
			result=false
		}
	}
	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
