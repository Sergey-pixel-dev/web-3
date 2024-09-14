package main

import "fmt"

func main() {
	str := ""
	fmt.Scan(&str)
	for i := 0; i < len(str)-1; i++ {
		fmt.Print(string(str[i]))
		fmt.Print("*")
	}
	fmt.Print(string(str[len(str)-1]))
}
