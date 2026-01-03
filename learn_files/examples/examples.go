package examples

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

// interview question

func Reading(dataCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	// Read the file line by line
	for Scanner.Scan() {
		line := Scanner.Text()
		fmt.Println(line)
		dataCh <- line
	}
	close(dataCh)

	if err := Scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Writing(dataCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Write the modified data back to the file
	file, err := os.OpenFile("write.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for ch := range dataCh {
		data := []byte(ch + "\n")
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
