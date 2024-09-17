package examples

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func Reading(dataCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	Scanner := bufio.NewScanner(file)

	// Read the file line by line
	for Scanner.Scan() {
		line := Scanner.Text()
		fmt.Println(line)
		dataCh <- line
	}
	close(dataCh)
	file.Close()

	if err := Scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Writing(dataCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.OpenFile("write.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for ch := range dataCh {
		_, err := writer.WriteString(ch + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}
	// write buffered data
	if err := writer.Flush(); err != nil {
		log.Fatal(err)
	}
}
