package inffile

import (
	"errors"
	"flag"
)

type FlagInfo struct {
	FilePath  string
	WordCount int
}

func GetProcParamInfo() (FlagInfo, error) {
	fi := FlagInfo{}
	flag.StringVar(&fi.FilePath, "file", "", "File path")
	flag.StringVar(&fi.FilePath, "f", "", "File path")
	flag.IntVar(&fi.WordCount, "top", 10, "Top word count")
	flag.IntVar(&fi.WordCount, "t", 10, "Top word count")
	flag.Parse()
	if fi.FilePath == "" {
		return fi, errors.New("File path is empty(specify -file or -f)")
	}
	return fi, nil
}
