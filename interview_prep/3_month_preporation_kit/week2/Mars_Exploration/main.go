package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
* Complete the 'marsExploration' function below.
*
* The function is expected to return an INTEGER.
* The function accepts STRING s as parameter.
*/

func marsExploration(s string) (curruption int32) {
    // Write your code here
    desired := "SOS"
    for i := 0; i < len(s); i += 3 {
	messageSlice := s[i:i+3]
	if messageSlice != desired {
	    //fmt.Println("we've found a bad transmision:", messageSlice)
	    for c := 0; c < 3; c++ {
		if messageSlice[c] != desired[c] {
		    //fmt.Printf("%s -> %s\n", string(messageSlice[c]), string(desired[c]))
		    //fmt.Println()
		    curruption++
		}
	    }
	}
    }
    return
}

func main() {
    //fmt.Println(marsExploration("SOSSPSSQSSOR"))
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := marsExploration(s)

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

