package main

import "fmt"

func FromTo(from int,to int)string{
	if from<0||from>99||to<0||to>99{
		return "Invalid\n"
	}
	var res []byte
	step := 1
	if from>to{
		step=-1
	}

	for i:=from;;i+=step{
		res=append(res,byte(i/10+'0'))
		res=append(res,byte(i%10+'0'))

		if i==to{
			break
		}
		res=append(res,',',' ')
	}
	res=append(res,'\n')
	return string(res)
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
