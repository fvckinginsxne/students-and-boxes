package main

import (
	"fmt"
	"math/rand"
)

const (
	numStudents    = 50
	maxAttempts    = 25
	numSimulations = 10000
)

func simulateSimulation(openBoxFunc func(int, []int) int) float64 {
	successCount := 0

	for range numSimulations {
		boxes := rand.Perm(numStudents)
		studentsSuccess := 0

		for student := range numStudents {
			attempts := 0
			currentBox := student

			for attempts < maxAttempts {
				currentBox = openBoxFunc(currentBox, boxes)
				attempts++

				if currentBox == student {
					studentsSuccess++
					break
				}
			}
		}

		if studentsSuccess == numStudents {
			successCount++
		}
	}

	return float64(successCount) / numSimulations
}

func openBox(currentBox int, boxes []int) int {
	return boxes[rand.Intn(numStudents)]
}

func openBoxWithAgreement(currentBox int, boxes []int) int {
	return boxes[currentBox]
}

func main() {
	fmt.Printf("Open boxes without agreement: %.1f\n", simulateSimulation(openBox))
	fmt.Printf("Open boxes with agreement: %.1f\n", simulateSimulation(openBoxWithAgreement))
}
