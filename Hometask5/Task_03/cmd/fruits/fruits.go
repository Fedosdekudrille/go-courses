package main

import (
	csvParse "Task_03/internal/csv_parse"
	"Task_03/pkg/knn"
	"fmt"
)

func main() {
	data, err := csvParse.Parse("assets/data.txt")
	if err != nil {
		panic(err)
	}
	marks, dataPoints, err := csvParse.ParseKnn(data)
	if err != nil {
		panic(err)
	}
	knn, err := knn.NewKnn(3)
	if err != nil {
		panic(err)
	}
	knn.Train(marks, dataPoints)
	for {
		var fruit []float64 = make([]float64, 2)
		fmt.Print("Enter fruit vector: ")
		_, err := fmt.Scanf("%f %f\n", &fruit[0], &fruit[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		mark, err := knn.Predict(fruit)
		if err != nil {
			panic(err)
		}
		fmt.Println(mark.Name)
	}
}
