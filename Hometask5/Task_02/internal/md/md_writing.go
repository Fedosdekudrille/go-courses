package md

import (
	"Task_02/internal/good"
	"math"
	"os"
	"text/template"
	"time"
)

func WriteGoodsByMdTemplate(filepath string, goods good.AllGoods) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	t := template.New("template.md").Funcs(
		template.FuncMap{
			"date": func() string {
				return time.Now().Format("1 January, 2006")
			},
			"sum": func(a, b int) int {
				return a + b
			},
			"round": func(float float64) float64 {
				return math.Round(float*100) / 100
			},
		},
	)
	t, err = t.ParseFiles("./internal/md/template.md")
	if err != nil {
		return err
	}
	return t.Execute(file, goods)
}
