package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strings"
	"strconv"
)

func sxord(s1, s2 string) string {
	si1, err := strconv.Atoi(s1)
	if err != nil {
		return ""
	}
	si2, err := strconv.Atoi(s2)
	if err != nil {
		return ""
	}
	return strconv.Itoa(int(uint(si1) ^ uint(si2)))
}

func main() {
	//fmt.Println(sxord("10101", "00101"))

	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	xorStrings := make([]string, 2)
	for i := 0; i < 2; i++ {
		stringTemp := strings.TrimSpace(readLine(reader))
		xorStrings = append(xorStrings, stringTemp)
	}

    result := sxord(xorStrings[0], xorStrings[1])

	fmt.Println(result)
    fmt.Fprintf(writer, "%s\n", result)

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
