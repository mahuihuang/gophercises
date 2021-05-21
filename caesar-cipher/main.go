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
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	k = k%26
	r := ""
	// temp string
	var AInt32 int32 = 65
	var ZInt32 int32 = 90
	var aInt32 int32 = 97
	var zInt32 int32 = 122

	for _, v := range s {
		// A-Z
		if AInt32 <= v && v <= ZInt32 {
			if t := v + k; t > ZInt32 {
				v = t - ZInt32 + AInt32 - 1
			} else {
				v += k
			}
			// a-z
		} else if aInt32 <= v && v <= zInt32 {
			if t := v + k; t > zInt32 {
				v = t - zInt32 + aInt32 -1
			} else {
				v += k
			}
		}
		c := string(v)
		r += c
	}
	fmt.Println(r)
	return r
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	_, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

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
