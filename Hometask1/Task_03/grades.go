package main

import "fmt"

func main() {
	var maxTaskGrade, tasksNum int
	totalGrade := 0

	fmt.Println("Insert max grade from single task")
	fmt.Scanln(&maxTaskGrade)

	fmt.Println("Insert the number of tasks")
	fmt.Scanln(&tasksNum)

	var currentTaskGrade int
	for i := 0; i < tasksNum; i++ {
		fmt.Printf("Insert grade for the %d task\n", i+1)
		fmt.Scanln(&currentTaskGrade)
		totalGrade += currentTaskGrade
	}
	var (
		grade   byte
		percent float64
	)
	percent = float64(totalGrade) / float64(tasksNum*maxTaskGrade)
	switch {
	case percent >= 0.9:
		grade = 'A'
	case percent >= 0.8:
		grade = 'B'
	case percent >= 0.7:
		grade = 'C'
	case percent >= 0.65:
		grade = 'D'
	default:
		grade = 'F'
	}
	fmt.Printf("Student's grade is %c\n", grade)
}
