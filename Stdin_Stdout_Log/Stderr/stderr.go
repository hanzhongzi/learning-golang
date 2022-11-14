package main

import (
	"io"
	"os"
)

func main() {
	argument := os.Args
	myString := ""
	if len(argument) == 1 {
		myString = "Please give me a argument"
		io.WriteString(os.Stderr, "这里是标准错误"+myString)
	} else {
		myString = argument[1]
		io.WriteString(os.Stdout, "这是标准输出 "+myString)
	}

}
