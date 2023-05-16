// https://www.hackerrank.com/challenges/three-month-preparation-kit-kangaroo

package main

import (
	"fmt"
) 

func kangaroo(x1, v1, x2, v2 int32) string {
	// loop until we're beyond the upper bounds of x1
	for i:= 0; i <= 10000; i++ {
		// check if the kangaroos are in the same position
		if x1 == x2 {
			return "YES"
		} 
		// move kangaroo
		x1 += v1
		x2 += v2
	} 
	return "NO"
} 

func main() {
	fmt.Println(kangaroo(1,2,2,1))
} 
