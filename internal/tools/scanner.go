package tools

import "fmt"

type Scan struct{}

func (s *Scan) Scanf(format string, a ...interface{}) (n int, err error) {
	return fmt.Scanf(format, a...)
}

type Scanner interface {
	Scanf(format string, a ...interface{}) (n int, err error)
}
