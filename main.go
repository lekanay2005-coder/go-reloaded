
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need both input and output file")
		return
	}

	inputFile := os.Args[1]
	output := os.Args[2]

	date, err := os.ReadFile(inputFile)
	if err != nil{
		fmt.Println("Error reading imput file", err)
		return
	}

	// result := string(date)
	result := processText(string(date))

    err = os.WriteFile(output, []byte(result), 0644)
	if err != nil{
		fmt.Println("Error writing file", err)
    }
}