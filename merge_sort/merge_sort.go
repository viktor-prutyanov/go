package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Enter size of numbers:\n")

	var size int = 0
	fmt.Scanf("%d", &size)
	var numbers []int = make([]int, size)
	var aux_arr []int = make([]int, size)

	fmt.Printf("Please enter %d numbers:\n", size)

	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &numbers[i])
	}

	fmt.Printf("Your array is:\n")
	fmt.Print(numbers)
	fmt.Printf("\nSorted array is:")
	sort(numbers, aux_arr, 0, size)
	fmt.Print(numbers)
	fmt.Printf("\n")
}

func sort(arr []int, aux []int, start int, end int) {
	if end <= start+1 {
		return
	}
	var middle int = (end-start+1)/2 + start
	sort(arr, aux, start, middle)
	sort(arr, aux, middle, end)
	merge(arr, aux, start, end, middle)
}

func merge(arr []int, aux []int, start int, end int, middle int) {
	copy(aux[0:(end-start)], arr[start:end])

	var j, k int = 0, (middle - start)
	for i := start; i < end; i++ {
		if k == end-start {
			arr[i] = aux[j]
			j++
		} else if j == (middle - start) {
			arr[i] = aux[k]
			k++
		} else {
			if aux[j] < aux[k] {
				arr[i] = aux[j]
				j++
			} else {
				arr[i] = aux[k]
				k++
			}
		}
	}
}
