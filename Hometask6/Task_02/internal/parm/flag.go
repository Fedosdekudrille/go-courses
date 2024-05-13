package parm

import (
	"flag"
	"strconv"
)

const (
	WRITE int = iota + 1
	READ
	COMPLETE
)

func GetFlags() (mode int, val string) {
	l := flag.Bool("list", false, "List all tasks")
	c := flag.Int("complete", -1, "Complete task")
	t := flag.String("task", "", "Add task")
	flag.Parse()
	if *l {
		mode = READ
	} else if *c != -1 {
		mode = COMPLETE
		val = strconv.Itoa(*c)
	} else if *t != "" {
		mode = WRITE
		val = *t
	}
	return
}
