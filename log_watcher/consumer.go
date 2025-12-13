package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func startConsumerForTextFile() {
	f, err := os.Open("/tmp/dat1")
	check(err)
	defer f.Close()

	scanner := bufio.NewReader(f)
	fmt.Println("starting the consumer")
	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				// fmt.Println("reached end of file, sleeping for 1 second")
				continue
			} else {
				fmt.Println("read error: ", err)
				break
			}
		}

		fmt.Println(line)
		PublishMessage(line)
		// if line == "close_connection\n" {
		// 	fmt.Println("close connection read")
		// 	break
		// }
	}
}
