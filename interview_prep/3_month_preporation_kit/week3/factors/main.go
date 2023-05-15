package main

import "fmt"

// getLargest returns the largest number in a slice
func getLargest(arr []int32) (largest int32) {
	for _, v := range arr {
		if v > largest {
			largest = v
		}
	}
	return
}

func getTotalX(a []int32, b []int32) (betweens int32) {
	start := getLargest(a)
	loopTo := getLargest(b)
	considered := map[int32]int32{}

	for _, v := range b {
		for i := start; i < loopTo; i++ {
			if v%i == 0 {
				considered[i]++
			}
		}
	}
	
	between := map[int32]int32{}
	for k, v := range considered {
		if int(v) == len(b) {
			for _, av := range a {
				if k%av == 0 {
					between[k]++
				}
			}
		}
	}

	for _, v := range between {
		if int(v) == len(a) {
			betweens++
		}
	}
	if len(a) == 1 && len(b) == 1 {
		if a[0] > b[0] {
			return 0
		}
		return int32(len(between)+1)
	}
	return
}

func main() {
	a := []int32{1}
	b := []int32{100}
	// a := []int32{51}
	// b := []int32{50}
	fmt.Println(getTotalX(a, b))
}
