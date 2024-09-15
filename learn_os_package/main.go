package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("--------------os---------------")

	// cmd := exec.Command("firefox")

	// err := cmd.Run()

	// if err != nil{
	// 	log.Fatal(err)
	// }

	cmd := exec.Command("tr", "a-z", "A-Z")

	cmd.Stdin = strings.NewReader("Subscribe to the channel")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translated phrase: %q\n", out.String())
}
