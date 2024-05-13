package fileLogger

import (
	"fmt"
	"os"
)

func Log(data string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = fmt.Fprintf(file, "%s\n", data); err != nil {
		return err
	}
	return nil
}

func LogSlice(data []string, filePath string) error { //wanted to use fmt.Stringer, but []string and []Worker(implements) don't convert to []fmt.Stringer
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, v := range data {
		if _, err = fmt.Fprintf(file, "%s\n", v); err != nil {
			return err
		}
	}
	return nil
}
