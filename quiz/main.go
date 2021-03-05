package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	limit := 30
	csvFile := "problems.csv"

	// 定义参数格式
	flag.StringVar(&csvFile, "csv", "problems.csv", "a csv file in format of 'question,answer'")
	flag.IntVar(&limit, "limit", limit, "the time limit for the quiz in seconds")
	flag.Parse()

	// 定义一个 timer
	t := time.NewTimer(time.Duration(limit) * time.Second)

	fileByte, err := os.ReadFile(csvFile)
	if err != nil {
		fmt.Println(err)
	}
	r := csv.NewReader(strings.NewReader(string(fileByte)))
	scanner := bufio.NewScanner(os.Stdin)
	
	go func ()  {
		for i := 1; ;i++ {
			record, err := r.Read()
	
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Problem #%d: %s = ", i, record[0])
			for scanner.Scan(){
				if scanner.Text() == record[1] {
					break
				} else {
					fmt.Println("cal err")
					return
				}
			}
		}
	}()

	<-t.C
	fmt.Println("")
	fmt.Println("Timeout")
}
