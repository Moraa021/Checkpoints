package main

import "fmt"

func Contains(s string, c rune) bool {
	for _, r := range s {
		if r == c {
			return true
		}
	}
	return false
}

func WeareUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	unique := ""

	for _, v := range str1 {
		if !Contains(str2, v) && !Contains(unique, v) {
			unique = unique + string(v)
		}
	}
	for _, v := range str2 {
		if !Contains(str1, v) && !Contains(unique, v) {
			unique = unique + string(v)
		}
	}
	return len(unique)
}

func main() {
	fmt.Println(WeareUnique("hello", "mercy"))
	fmt.Println(WeareUnique("return", "remain"))
	fmt.Println(WeareUnique("hi", "hit"))
	fmt.Println(WeareUnique("cool", "kind"))
}
