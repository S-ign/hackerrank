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
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) (vallyCount int32) {
    // Write your code here
    seaLevel := 0
    current := 0
    surfaced := true

    for _, v := range path {
        switch string(v) {
        case "U":
            current++
            //fmt.Println("going up!", current)
        case "D":
            current--
            //fmt.Println("going down!", current)
        }
        if current == seaLevel {
            surfaced = true
            //fmt.Println("we have resurfaced")
        }
        if current == seaLevel-1 {
            if surfaced {
                if current < seaLevel {
                    vallyCount++
                    //fmt.Println("we in da valley!!!")
                }
                surfaced = false
            }
        }
    }
    return
}

func main() {
    //v := "UDDDUDUU"
    //fmt.Println(countingValleys(int32(len(v)),v))

    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    steps := int32(stepsTemp)

    path := readLine(reader)

    result := countingValleys(steps, path)

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

