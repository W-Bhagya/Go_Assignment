package main

import "fmt"

func rev(s string) string {
	arr := []rune(s)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return string(arr)

}
func main() {
	str := "Bhagya"

	result := rev(str)
	fmt.Println("String Befor Reverse : ", str)
	fmt.Println("String After Reverse : ", result)
}
