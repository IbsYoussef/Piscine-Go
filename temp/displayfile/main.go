package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("File name missing")
		return
	}

	if len(os.Args[1:]) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	if len(os.Args[1:]) == 1 {
		filename := os.Args[1]

		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		fmt.Print(string(content))

	}
}
