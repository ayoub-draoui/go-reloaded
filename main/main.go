package main

import (
	"bufio"
	"fmt"
	"os"
	"reloaded"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("you must enter two arguments:")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)

	}
	if os.Args[1] != "sample.txt" {
		fmt.Println("Error: the argument must be sample.txt")
		os.Exit(1)
	}
	if os.Args[2] != "result.txt" {
		fmt.Println("Error: the argument must be result.txt")
		os.Exit(1)
	}

	output, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		changedText := reloaded.EditText(line)
		_, err = output.WriteString(changedText + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			os.Exit(1)
		}
		defer output.Close()
	}
}
