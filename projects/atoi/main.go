package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for _, c := range s {
		c -= '0'
		fmt.Print(c * c)
	}
}
