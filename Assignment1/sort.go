//This program takes the input file name on command line arguments sorts all numbers in the inut file, and puts the sorted numbers into a new file
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var file, err0 = os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	if err0 != nil {
		fmt.Println("Error in opening file!")
		return
	}

	fileScanner := bufio.NewScanner(file)
	var numbers []float64

	for fileScanner.Scan() {
		line := fileScanner.Text()
		add, _ := strconv.ParseFloat(line, 64)
		numbers = append(numbers, add)
	}

	sort.Float64s(numbers)

	output, err1 := os.Create("sortedNumbers.txt")

	if err1 != nil {
		fmt.Println("Error in creating files!")
		return
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Fprintln(output, numbers[i])
	}

	fmt.Println("Writing completed")
	file.Close()
}
