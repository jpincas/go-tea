package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	fileContents = "package gotea\n\nconst goteaJS = `%s`"
)

func main() {
	data, err := ioutil.ReadFile("js/dist/gotea.js")
	if err != nil {
		log.Fatalln("File reading error", err)
	}

	toWrite := fmt.Sprintf(fileContents, string(data))

	ioutil.WriteFile("js.go", []byte(toWrite), os.ModePerm)
}
