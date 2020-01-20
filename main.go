package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("Untitled.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "<") {
			text := scanner.Text()
			fmt.Println(text)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
