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
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func lowestInArr(arr []int32) (lowest int32) {
	lowest = 20000001
	fmt.Println(arr)
	for _, v := range arr {
		if lowest > v {
			lowest = v
		}
	}
	return
}

func migratoryBirds(arr []int32) int32 {
	mapCount := make(map[int32]int32)
	var highestCountId []int32
	var highestBirdCount int32

	for _, v := range arr {
		mapCount[v]++
	}

	fmt.Println("mapCount: ", mapCount)

	for _, birdCount := range mapCount {
		if highestBirdCount < birdCount {
			highestBirdCount = birdCount
		}
	}
	for id, birdCount := range mapCount {
		if birdCount == highestBirdCount {
			highestCountId = append(highestCountId, id)
		}
	}
	return lowestInArr(highestCountId)
}

func main() {
	//fmt.Println(migratoryBirds([]int32{1, 1, 2, 2, 3}))
	//fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(arrCount); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := migratoryBirds(arr)

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
