package main

import "fmt"

func Has(s string,c rune)bool{
	for _,r:=range s{
		if r==c{
			return true
		}
	}
	return false
}

func WeAreUnique(str1,str2 string)int{
	if str1==""&& str2==""{
		return -1
	}
	unique:=""

	for _,v:=range str1{
		if !Has(str2,v)&& !Has(unique,v){
			unique= unique+string(v)
		}
	}
	for _,v:=range str2{
		if !Has(str1,v) && !Has(unique,v){
			unique= unique+string(v)
		}
	}
	return len(unique)
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
