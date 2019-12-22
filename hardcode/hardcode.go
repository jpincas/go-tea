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
	// JS
	data, err := ioutil.ReadFile("js/dist/gotea.js")
	if err != nil {
		log.Fatalln("File reading error", err)
	}

	js := string(data)

	// // Map
	// data, err = ioutil.ReadFile("js/dist/gotea.map")
	// if err != nil {
	// 	log.Fatalln("File reading error", err)
	// }

	// mapp := string(data)

	ioutil.WriteFile("js.go", []byte(fmt.Sprintf(fileContents, js)), os.ModePerm)
}
