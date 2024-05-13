package md

import (
	"os"
	"strconv"
)

//IMPORTANT! Wrote this file before writing with the help of templates, liked it so much that leaved it, still it is UNUSED

type Md struct {
	file *os.File
}

func NewMd(filePath string) *Md {
	md := &Md{}
	md.OpenOrCreate(filePath)
	return md
}

func (md *Md) OpenOrCreate(filePath string) error {
	if md.file != nil {
		if err := md.file.Close(); err != nil {
			return err
		}
	}
	if file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666); err != nil {
		return err
	} else {
		md.file = file
	}
	return nil
}

func (md *Md) Close() error {
	if md.file != nil {
		return md.file.Close()
	}
	return nil
}

func (md *Md) WriteH1(str string) error {
	_, err := md.file.WriteString("# " + str + "\n")
	return err
}

func (md *Md) WriteH2(str string) error {
	_, err := md.file.WriteString("## " + str + "\n")
	return err
}

func (md *Md) Write(str string) error {
	_, err := md.file.WriteString(str)
	return err
}

func (md *Md) WriteLn(str string) error {
	_, err := md.file.WriteString(str + "\n")
	return err
}

func (md *Md) WriteNumList(list []string) error {
	for id, item := range list {
		_, err := md.file.WriteString(strconv.Itoa(id+1) + ". " + item + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
