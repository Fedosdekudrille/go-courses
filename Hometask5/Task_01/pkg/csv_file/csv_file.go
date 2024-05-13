package csvFile

import (
	"bufio"
	"os"
	"strings"
)

const avgStrByteSize = 35 //prediction to count size of result slice

func ParseFile(filePath string) (res [][]string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		if err = scanner.Err(); err != nil {
			return nil, err
		}
		res = append(res, strings.Split(data, ","))
	}
	return
}
