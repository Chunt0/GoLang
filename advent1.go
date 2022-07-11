package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Open advent1.txt file, check for errors
	file := "./advent1.txt"
	data, err := os.Open(file)
	check(err) // Check for errors

	// Turn data into a scanner struct then split into words
	adventInput := bufio.NewScanner(data)
	adventInput.Split(bufio.ScanWords)

	// Scan through words, convert to int, append to lines slice
	var lines []int
	for adventInput.Scan() {
		line := adventInput.Text()
		val, _ := strconv.Atoi(line)
		check(err)
		lines = append(lines, val)
	}

	// Part One
	counter := 0
	temp := 0
	for k, v := range lines {
		if k == 0 {
			temp = v
			continue
		}
		if temp < v {
			counter++
		}
		temp = v
	}
	fmt.Println(counter)

	// Part Two
	counter = 0
	prev, current, next, tempSum, currentSum := 0, 0, 0, 0, 0
	for i := 1; i < len(lines)-1; i++ {
		prev = lines[i-1]
		current = lines[i]
		next = lines[i+1]
		currentSum = prev + current + next

		if i == 1 {
			tempSum = currentSum
			continue
		}
		if tempSum < currentSum {
			counter++
		}
		tempSum = currentSum
	}
	fmt.Println(counter)
}
