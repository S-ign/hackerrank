package main

import "fmt"

func rotateLeft(d int32, arr []int32) []int32 {
	// rotate/loop as many times as dictated by varibale d
	for ; d > 0; d-- {
		// itterate through loop 1 less than length of arr
		for i := 0; i < len(arr)-1; i++ {
			// flip the current number with the number above it
			// also why we stopped one less than the size of the arr, so we don't overflow
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	return arr
}

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	fmt.Println(rotateLeft(5, arr))
}
