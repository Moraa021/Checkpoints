package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args)!=2{
		fmt.Print("\n")
		return
	}
	s:=os.Args[1]
	word:=false
	firstword:=true

	for _,r:=range s{
		if r!=' '&&r!='\t'{
			if !word && !firstword{
				fmt.Print(" ")
			}
			fmt.Print(string(r))
			word=true
			firstword=false
		}else{
			word=false
		}
	}
	fmt.Print("\n")
}