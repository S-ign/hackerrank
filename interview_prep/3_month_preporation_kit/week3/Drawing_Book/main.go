package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'pageCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
 */

func pageCount(n int32, p int32) (page int32) {
	// Write your code here
	page = 0

	// convert p to floating number
	f := float32(p)

	// if the desired page is greater than half of all pages in the book, then start flipping
	// pages from the begining
	if f <= float32(n/2) {
		fmt.Println("flipping from start")

		// iterate over the length n, 2 pages at a time ( since there are two pages in each "page flip" )
		for i := 1; i < int(n); i += 2 {

			// if the page number (i) is the correct page, stop flipping pages
			if i == int(p) || i-1 == int(p) {
				break
			}
			// else incrament the pages fliped
			page++
		}

		// else start fliping from the end of the book
	} else {

		fmt.Println("flipping from end", f, float32(n/2))

		// If the the page is 1 less than the legnth of the book and the last page is even return 1
		// This just handles the edge case if the book ends on an even page, similar to how the
		// beggining of the book starts with 1 page.
		if p == n-1 && n%2 == 0 {
			return int32(1)
		}

		// starting from the last page, iterate over the length n, 2 pages at a time
		// ( since there are two pages in each "page flip" )
		for i := int(n) - 1; i > 0; i -= 2 {

			// if the page number (i) is the correct page, stop flipping pages
			if i == int(p) || i+1 == int(p) {
				break
			}
			// else incrament the pages fliped
			page++
		}
	}
	return
}

func main() {
	fmt.Println(pageCount(6, 5))


	// This stuff gets input from hackerrank
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	// pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// checkError(err)
	// p := int32(pTemp)

	// result := pageCount(n, p)

	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
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
