package main

import (
	"fmt"	
	"os"
	"bufio"
	"github.com/gomarkdown/markdown"
)

func main() {

	fileName := "smth.md"

	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)

	var markdownText string = ""

	for scanner.Scan() {
		markdownText += scanner.Text() + "\n"
	}

	markdown1 := markdown.ToHTML([]byte(markdownText), nil, nil)

	fmt.Println(string(markdown1))
}