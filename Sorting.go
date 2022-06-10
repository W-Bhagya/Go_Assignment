package main

import "fmt"

func bubbleSort(arr[] int, size int)  {
	
	for i := 0; i <= size-1; i++ {
		for j := 0; j < size-i-1; j++ {		
		
		if arr[j] > arr[j+1] {
			temp := arr[j]
			arr[j] = arr[j+1]
			arr[j+1] = temp
			
		}
	}
	
 }

}


func main() {

	fmt.Print("Enter size of array : ")
	var size int
	fmt.Scanf("%d",&size)
	
	var arr = make([]int, size)

	fmt.Print("Enter arry elements : ")

	for i := 0; i <= size-1; i++ {
		fmt.Scan(&arr[i])
	}
	
	fmt.Print("Array Before Sorting : ")
	for i := 0; i <= size-1; i++ {
		fmt.Print(" ",arr[i])
	}

	bubbleSort(arr,size)
	
	fmt.Print("\nArray After Sorting : ")
	
	for i := 0; i <= size-1; i++ {
		fmt.Print(" ",arr[i])
	}

}
