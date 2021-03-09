package main

import "fmt"

func NbYear(p0 int, percent float64, aug int, p int) int {
	currentPopulation := float64(p0)
	var yearsTaken int
	for int(currentPopulation) < p {
		currentPopulation += currentPopulation*(percent/100) + float64(aug)

		yearsTaken++
	}

	return yearsTaken
}

func main() {
	fmt.Println(
		NbYear(1500, 5, 100, 5000),
		NbYear(1500000, 2.5, 10000, 2000000),
		NbYear(1500000, 0.25, 1000, 2000000),
		NbYear(1500000, 0.25, -1000, 2000000),
	)
}
