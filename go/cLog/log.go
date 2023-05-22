package cLog

import "fmt"

const (
	LOG_DEBUG = iota
	LOG_INFO
	LOG_WARM
	LOG_ERROR
)

const LOG_LEVEL = LOG_DEBUG

const SHOW_LOG = false

func Println(a ...interface{}) (n int, err error) {
	if SHOW_LOG {
		return fmt.Println(a...)
	}
	return 0, nil
}

func Printf(s string, a ...any) (n int, err error) {
	if SHOW_LOG {
		return fmt.Printf(s, a...)
	}
	return 0, nil
}

func Print(a ...interface{}) (n int, err error) {
	if SHOW_LOG {
		return fmt.Print(a...)
	}
	return 0, nil
}