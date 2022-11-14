package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// https://syaning.github.io/go-pkgs/strconv/
func main() {
	if len(os.Args) == 1 {
		io.WriteString(os.Stdout, "请输入至少一个参数！")
		os.Exit(1)
	}
	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)
	for i := 1; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}

	}
	fmt.Println("min: ", min, "max: ", max)
}
