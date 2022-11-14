package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument !"
	} else {
		for _, v := range arguments[1:] {
			myString += v + " "
		}
	}
	io.WriteString(os.Stdout, myString)
}
