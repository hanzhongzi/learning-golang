package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.SetOutput(sysLog)
		fmt.Println(syslog.LOG_INFO, syslog.LOG_LOCAL7)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!! ")
	sysLog, err = syslog.New(syslog.LOG_MAIL, "wahahah")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_MAIL: logging in go")
	fmt.Println("LOG_MAIL: logging in go")
	// log.Panic() 程序的报错，包含底层信息

}
