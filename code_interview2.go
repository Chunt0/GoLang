package main

import (
	"fmt"
	"math"
)

type Histo struct {
	ages      []float64
	histogram []float64
}

func get_histo(histo Histo) Histo {
	for _, v := range histo.ages {
		index := int(math.Floor(v / 10))
		histo.histogram[index]++
	}
	for k, v := range histo.histogram {
		histo.histogram[k] = (v / float64(len(histo.ages)))
	}
	return histo
}

func main() {

	ages := []float64{99, 88, 77, 66, 55, 44, 33, 22, 11}
	histogram := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	histo := Histo{
		ages:      ages,
		histogram: histogram}

	histo = get_histo(histo)

	for k, v := range histo.histogram {
		fmt.Printf("Block %v: %.2f\n", k, v)
	}
}
