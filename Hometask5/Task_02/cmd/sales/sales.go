package main

import (
	csvParse "Task_02/internal/csv_parse"
	"Task_02/internal/good"
	"Task_02/internal/md"
)

func main() {
	data, err := csvParse.Parse("./assets/sales.csv")
	checkErr(err)
	goods, err := good.Parse(data)
	checkErr(err)
	err = md.WriteGoodsByMdTemplate("./assets/report.md", goods)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
