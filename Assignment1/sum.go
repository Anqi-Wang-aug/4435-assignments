//This file takes a txt file name in command line arguments and add all numbers in this file
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum float64
	sum = 0
	var file, err = os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error in opening file!")
		return
	}

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		add, _ := strconv.ParseFloat(line, 64)
		sum = sum + add
	}

	fmt.Println("The sum of all numbers in this file is: ", sum)
	file.Close()
}
