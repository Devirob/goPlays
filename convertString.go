package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var status string

	reader := bufio.NewReader(os.Stdin)
	inputFromKeyboard, err := reader.ReadString('\n')
	inputFromKeyboard = strings.TrimSpace(inputFromKeyboard)

	grade, err := strconv.ParseFloat(inputFromKeyboard, 64)
	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		status = "passing"
	} else {
		status = "fail"
	}

	fmt.Println(status)
}
