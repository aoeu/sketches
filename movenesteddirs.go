package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	rootdir := fmt.Sprintf("/tmp/%v", os.Getpid())
	dir := "up/down/left/right"
	trash := "trash"
	source := fmt.Sprintf("%v/%v", rootdir, dir)
	dest := fmt.Sprintf("%v/%v/%v", rootdir, trash, dir)
	exec(os.MkdirAll, source)
	exec(os.Mkdir, fmt.Sprintf("%v/%v", rootdir, trash))
	exec(os.MkdirAll, dest)
	if err := os.Rename(source, dest); err != nil {
		log.Fatal(err)
	}
}

func exec(mkdirFunc func (string, os.FileMode) error, dirname string) {
	if err := mkdirFunc(dirname, 0644); err != nil {
		log.Fatal(err)
	}
}
