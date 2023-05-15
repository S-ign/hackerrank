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
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) (largest int32) {
	// sort a
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	// create 2d slice to store list of sub arrays
	var matrix [][]int32

	// create 2d slice for storing sub arrays
	var arr []int32

	// itterate over slice a
	for i, v := range a {

		// at the start, place the first number in the sub array slice
		if i == 0 {
			arr = append(arr, v)
			continue
		}

		// if the first number in the sub array is the same as the current or only 1 above
		// add it to the sub array
		if v == arr[0] || v == arr[0]+1 || v == arr[0]-1 {
			arr = append(arr, v)

		// else add what we have currently to the matrix, reset the sub array,
		// and then add our current number into the sub array
		} else {
			matrix = append(matrix, arr)
			arr = []int32{}
			arr = append(arr, v)
		}

		// if we're at the end of the loop, add what we have in the sub array to the matrix
		if i == len(a)-1 {
			matrix = append(matrix, arr)
		}

	}

	// itterate through our matrix and find the largest length sub array
	for _, v := range matrix {
		if len(v) > int(largest) {
			largest = int32(len(v))
		}
	}

	return
}

func main() {
	// a := []int32{1, 1, 2, 2, 4, 4, 5, 5, 5}
	// fmt.Println(pickingNumbers((a))) // 11
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := pickingNumbers(a)

	fmt.Fprintf(writer, "%d\n", result)

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
