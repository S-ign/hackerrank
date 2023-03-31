package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func getSmallestNumbers(arr []int32) int64 {
	var largestIndex int
	l := arr[0]

	for i, v := range arr {
		if v > l {
			l = v
			largestIndex = i
		}
	}
	arr = append(arr[:largestIndex], arr[largestIndex+1:]...)
	return sum(arr)
}

func getLargestNumbers(arr []int32) int64 {
	var smallestIndex int
	s := arr[0]

	for i, v := range arr {
		if v < s {
			s = v
			smallestIndex = i
		}
	}
	largestNumbers := append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	return sum(largestNumbers)
}

func sum(arr []int32) (sum int64) {
	for _, v := range arr {
		sum += int64(v)
	}
	return
}

func miniMaxSum(arr []int32) {
	a := make([]int32, len(arr))
	copy(a, arr)
	b := make([]int32, len(arr))
	copy(b, arr)
	s := getSmallestNumbers(b)
	l := getLargestNumbers(a)
	fmt.Println(s, l)
}

func main() {
	//miniMaxSum([]int32{1, 2, 3, 4, 5})
	//miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
	//miniMaxSum([]int32{396285104, 573261094, 759641832, 819230764, 364801279})
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
