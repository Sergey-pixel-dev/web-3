package main

import "fmt"

func main() {
	Str := ""
	var MaxVal rune
	fmt.Scan(&Str)
	for _, val := range Str {
		if MaxVal < val {
			MaxVal = val
		}
	}
	fmt.Print(MaxVal - 48)
}
