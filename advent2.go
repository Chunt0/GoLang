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
	file := "./advent2.txt"
	data, err := os.Open(file)
	check(err)

	adventInput := bufio.NewScanner(data)
	adventInput.Split(bufio.ScanWords)

	move := ""
	x, y := 0, 0
	counter := 0
	for adventInput.Scan() {
		line := adventInput.Text()
		fmt.Println(line)
		switch line {
		case "forward", "down", "up":
			move = line

		default:
			counter++
			val, err := strconv.Atoi(line)
			check(err)
			switch move {
			case "forward":
				x = x + val

			case "down":
				y = y + val

			case "up":
				y = y - val

			}
		}
	}
	x = x + 7

	fmt.Println(x * y)
}
