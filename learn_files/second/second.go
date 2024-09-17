package second

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {
	fmt.Println("Writing to a file in golang")

	file, err := os.Create("file.txt")

	if err != nil {
		log.Fatalf("failed creating file %s", err)
	}

	defer file.Close()

	len, err := file.WriteString("Welcome to the channel go guruji. :)")

	if err != nil {
		log.Fatalf("failed writing file %s", err)
	}

	fmt.Println("File name: ", file.Name())
	fmt.Printf("Length: %d bytes \n", len)
}

func ReadFile(){
	fmt.Println("Reading from the file")

	fileName := "file.txt"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Panicf("failed reading data from file %s", err)
	}

	fmt.Printf("\n File name: %s \n", fileName)
	fmt.Printf("\n Size: %d bytes \n", len(data))
	fmt.Printf("\n Data: %s \n", data)
	

}	