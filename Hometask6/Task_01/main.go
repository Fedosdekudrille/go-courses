package main

import (
	"Task_01/internal/cntword"
	"Task_01/internal/inffile"
	"fmt"
)

func main() {
	fi, err := inffile.GetProcParamInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := cntword.CountWords(fi.FilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := cntword.SortWords(data)
	for i := 0; i < fi.WordCount && i < len(w); i++ {
		fmt.Printf("%s: %d\n", w[i].Word, w[i].Count)
	}
}
