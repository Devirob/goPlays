package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Shoot")

	for shoots := 10; shoots > 1; shoots-- {

		inputFromKeyboard, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputFromKeyboard = strings.TrimSpace(inputFromKeyboard)
		shoot, err := strconv.Atoi(inputFromKeyboard)
		if err != nil {
			log.Fatal(err)
		}

		if shoot < target {
			fmt.Println("zamalo")
		} else if shoot > target {
			fmt.Println("Za duzo")
		} else {
			fmt.Println("Huirra")
			break
		}
		fmt.Println("Left " + fmt.Sprint(shoots))

	}

}
