package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func pangrams(s string) string {
    // Write your code here
    alpha := map[rune]bool{
	'a':false, 'b':false, 'c':false, 'd':false, 'e':false, 'f':false, 'g':false, 'h':false, 'i':false, 'j':false,
	'k':false, 'l':false, 'm':false, 'n':false, 'o':false, 'p':false, 'q':false, 'r':false, 's':false, 't':false,
	'u':false, 'v':false, 'w':false, 'x':false, 'y':false, 'z':false,
    }

    s = strings.TrimSpace(strings.ToLower(s))

    for _,v := range s {
	alpha[v] = true
    }
    for _, v := range alpha {
	if !v {
	    return "not pangram"
	}
    }
    return "pangram"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := pangrams(s)

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


