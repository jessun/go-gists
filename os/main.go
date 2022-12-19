package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	line, err := getOsInfo("./ky", '\n')
	if err != nil {
		panic(err)
	}
	newLine := strings.Fields(line)
	fmt.Printf("=====")
	fmt.Println(len(newLine))
	fmt.Printf("=====")
	if len(newLine) >= 7 {
		version := strings.Split(newLine[5], ".")
		versions := newLine[0] + version[0]
		fmt.Println(versions)
	}

}

func getOsInfo(filename string, delim byte) (string, error) {
	var newLine string
	f, err := os.OpenFile(filename, os.O_APPEND, 0666)
	if nil == err {
		buff := bufio.NewReader(f)
		line, err := buff.ReadString(delim)
		if err != nil || io.EOF == err {
			return "", err
		}
		newLine = line
	}
	return newLine, nil
}
