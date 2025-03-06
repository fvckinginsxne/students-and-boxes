package main

import (
	"fmt"
	"math/rand"
)

const (
	numberOfBoxes    = 50
	numberOfStudents = 50
	numberOfChoices  = 25
	sample           = 10000
)

func numberSearch(boxes, students []int) float64 {
	winCount := 0

	for range sample {
		allStudentsFound := true
		for _, student := range students {
			found := false
			for range numberOfChoices {
				if student == boxes[rand.Intn(numberOfBoxes)] {
					found = true
					break
				}
			}

			if !found {
				allStudentsFound = false
				break
			}
		}

		if allStudentsFound {
			winCount++
		}
	}

	return float64(winCount) / sample * 100
}

func numberSearchWithContract(boxes, students []int) float64 {
	winCount := 0

	for range sample {
		allStudentsFound := true
		for _, student := range students {
			found := false
			choice := student - 1
			for range numberOfChoices {
				if student == boxes[choice] {
					found = true
					break
				}

				choice = boxes[choice] - 1
			}

			if !found {
				allStudentsFound = false
				break
			}
		}

		if allStudentsFound {
			winCount++
		}
	}

	return float64(winCount) / sample * 100
}

func generateNumbers(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = i + 1
	}

	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	return numbers
}

func main() {
	boxes := generateNumbers(numberOfBoxes)
	students := generateNumbers(numberOfStudents)

	var probability float64

	probability = numberSearch(boxes, students)

	fmt.Println(probability)

	probability = numberSearchWithContract(boxes, students)

	fmt.Println(probability)
}
