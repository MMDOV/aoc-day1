package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("numbers.input")
	if err != nil {
		fmt.Println("error opening file", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numbers [][]float64
	for scanner.Scan() {
		strNumbers := strings.Fields(scanner.Text())
		var intNumbers []float64
		for _, str := range strNumbers {
			num, err := strconv.ParseFloat(str, 64) // Convert string to int
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			intNumbers = append(intNumbers, num)
		}
		numbers = append(numbers, intNumbers)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error opening file", err)
	}
	list1 := make([]float64, len(numbers))
	list2 := make([]float64, len(numbers))

	for i, pair := range numbers {
		list1[i] = pair[0]
		list2[i] = pair[1]
	}
	sort.Float64s(list1)
	sort.Float64s(list2)
	//var diffs []float64
	//for i, num := range list1 {
	//	diffs = append(diffs, math.Abs(num-list2[i]))
	//}
	//var sum float64
	//for _, num := range diffs {
	//	sum += num
	//}
	//fmt.Printf("%.0f\n", sum)
	var allNums []float64
	for _, i := range list1 {
		var count float64
		for _, j := range list2 {
			if i == j {
				count += 1
			}
		}
		num := i * count
		allNums = append(allNums, num)
	}

	var sum float64
	for _, num := range allNums {
		sum += num
	}
	fmt.Printf("%.0f\n", sum)
}
