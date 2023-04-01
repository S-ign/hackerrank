package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type camelCaseString struct {
	operation string //split, combine
	typeOf    string //method, class, variable
	word      []rune
}

func newCamelCaseString(s string) (c camelCaseString) {
	cs := strings.Split(s, ";")
	c.operation = cs[0]
	c.typeOf = cs[1]
	c.word = []rune(cs[2])
	return
}

func (c *camelCaseString) split() {
	pos := []int{}
	c.word = []rune(strings.Trim(string(c.word), "()"))
	for i, v := range c.word {
		if unicode.IsUpper(v) {
			c.word[i] = unicode.ToLower(v)
			pos = append(pos, i)

			if i == 0 {
				continue
			} else {
				c.word = append(c.word[:i], append([]rune{' '}, c.word[i:]...)...)
			}
		}
	}
}

func (c *camelCaseString) join() {
	for i, v := range c.word {
		if v == ' ' {
			c.word[i+1] = unicode.ToUpper(c.word[i+1])
			c.word = append(c.word[:i], c.word[i+1:]...)
		}
	}
}

func (c camelCaseString) variable() {
	c.word[0] = unicode.ToLower(c.word[0])
}

func (c *camelCaseString) method() {
	c.word[0] = unicode.ToLower(c.word[0])
	c.word = append(c.word, []rune{'(', ')'}...)
}

func (c *camelCaseString) class() {
	c.word[0] = unicode.ToUpper(c.word[0])
}

func (c camelCaseString) toString() string {
	return string(c.word)
}

func camelCaseHandler(s string) string {
	var c camelCaseString
	c = newCamelCaseString(s)

	switch c.operation {
	case "S":
		c.split()
		return c.toString()
	case "C":
		c.join()
	}

	switch c.typeOf {
	case "M":
		c.method()
	case "C":
		c.class()
	case "V":
		c.variable()
	}

	return c.toString()
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	//fmt.Println(camelCaseHandler("S;M;plasticCup()"))
	//fmt.Println(camelCaseHandler("C;V;mobile phone"))
	//fmt.Println(camelCaseHandler("C;C;coffee machine"))
	//fmt.Println(camelCaseHandler("S;C;LargeSoftwareBook"))
	//fmt.Println(camelCaseHandler("C;M;white sheet of paper"))
	//fmt.Println(camelCaseHandler("S;V;pictureFrame"))

	//var input string
	//fmt.Scanf("%s", &input)
	//fmt.Printf(camelCaseHandler(input))

	//for i := 0; i < 4; i++ {
	//	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	//	fmt.Println(camelCaseHandler(readLine(reader)))
	//}

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	for true {
		s := readLine(reader)
		if s == "" {
			break
		}
		result := camelCaseHandler(s)
		fmt.Fprintf(writer, "%s\n", result)
	}

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
