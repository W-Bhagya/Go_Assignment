package main

import "fmt"

func fact(x int) int {
	if x == 1 {
		return 1
	}
	return x * fact(x-1)
}

func main() { 
	var num int

	fmt.Print("Enter the number : ")
	fmt.Scanf("%d",&num)

	res := fact(num)
	fmt.Println("Factorial = ", res)
}

