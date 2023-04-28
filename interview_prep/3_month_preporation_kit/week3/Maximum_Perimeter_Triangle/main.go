// Given an array of stick lengths, use 3 of them to construct a non-degenerate triangle with the maximum possible perimeter. Return an array of the lengths of its sides as 3 integers in non-decreasing order.
// 
// If there are several valid triangles having the maximum perimeter:
// 
// Choose the one with the longest maximum side.
// If more than one has that maximum, choose from them the one with the longest minimum side.
// If more than one has that maximum as well, print any one them.
// If no non-degenerate triangle exists, return [-1].

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'maximumPerimeterTriangle' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY sticks as parameter.
 */

func maximumPerimeterTriangle(sticks []int32) []int32 {
	// sort slice in decending order, because we only care about the largest 3 numbers.
	sort.Slice(sticks, func(i, j int) bool {
		return sticks[i] > sticks[j]
	})

	// slide through the slice, to assess each 3 next largest contiguous numbers in the slice.
	for i := 0; i < len(sticks)-2; i++ {
		a := sticks[i]
		b := sticks[i+1]
		c := sticks[i+2]

		// check if it is a non-degenerate triangle (a triangle that does NOT form a staight line)
		// **wild I know.. but apparently thats a thing**
		// anyway, since we need to find the largest possible non-degenerate triangle and we
		// already sorted the slice in decending order, we will return as soon as we find a
		// match, reordering the numbers from smallest to largest
		if a < b+c {
			return []int32{c, b, a}
		}
	}
	return []int32{-1}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	sticksTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var sticks []int32

	for i := 0; i < int(n); i++ {
		sticksItemTemp, err := strconv.ParseInt(sticksTemp[i], 10, 64)
		checkError(err)
		sticksItem := int32(sticksItemTemp)
		sticks = append(sticks, sticksItem)
	}

	result := maximumPerimeterTriangle(sticks)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
