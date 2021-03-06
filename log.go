package util

import (
	"log"
	"os"
)

var Log *log.Logger

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func LogInit(path string) {
	println("logging path = " + path)
	file, err := os.Create(path)
	check(err)
	logger = log.New(file, "LOG: ", log.LstdFlags|log.Lshortfile)
}
