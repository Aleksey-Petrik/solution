package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {

	var lineBrackets string

	fmt.Print("Input line: ")
	_, err := fmt.Scan(&lineBrackets)
	if err != nil {
		log.Fatalln(err)
	}

	if isBalanced(strings.Split(lineBrackets, "")) {
		fmt.Println("Balanced")
	} else {
		fmt.Println("Unbalanced")
	}
}
