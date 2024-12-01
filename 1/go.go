package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}

	list1 := []int{}
	list2 := []int{}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Failed to convert ", parts[0], " to integer.")
		}
		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Failed to convert ", parts[1], " to integer.")
		}
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from the file:", err)
	}
	
	numberCount := make(map[int]int)
	sort.Ints(list1)
	sort.Ints(list2)

	diffSum := 0

	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}
		diffSum += diff
		n := list2[i]
		if count, exists := numberCount[n]; exists {
			numberCount[n] = count + 1
		} else {
			numberCount[n] = 1
		}
	}
	fmt.Println("Sum of differences: ", diffSum)

	similarity := 0
	for i := 0; i < len(list1); i++ {
		similarity += list1[i] * numberCount[list1[i]]
	}

	fmt.Println("Similarity: ", similarity)
}
