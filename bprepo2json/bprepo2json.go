package main

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello World!")
}

func processFile(filename string, out io.Writer) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Warningln("processFile error:", err.Error())
		return err
	}
	defer f.Close()
	return processReader(filename, f, out)
}

func processReader(filename string, in io.Reader, out io.Writer) error {
	return nil
}
