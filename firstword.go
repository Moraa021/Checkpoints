package main

import "fmt"

func FirstWord(s string) string {
	if len(s)==0{
		return ""
	}

	result:=""+"\n"
	i:=0

	for i<len(s) && s[i] ==' '{
		i++
	}

	for i<len(s) && s[i] !=' '{
		result=result+string(s[i])
		i++
	}
	return result+"\n"
}

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}