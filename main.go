package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"educationalsp/rpc"
)

func main() {
	logger := getLogger("/home/vmubuntu/projects/lsp/log.txt")
	logger.Println("Hey, I started")
	fmt.Println("hi")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Hey, you didn't give me a good file")
	}

	return log.New(logfile, "[educationalsp]", log.Ldate|log.LUTC|log.Lshortfile)
}
