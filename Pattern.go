package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		var num = i
		for j := 1; j <= i; j++ {
			fmt.Print(num, " ")
			num = num + 5 - j
		}
		fmt.Println()
	}

}
