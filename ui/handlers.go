package ui

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/MishaAC/go-string-utils-cli/service"
)

func countVowels(reader *bufio.Reader, svc service.StringService) {
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading input:", err)
		return
	}
	text = strings.TrimSpace(text)
	numberOfVowels, err := svc.CountVowels(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Number of vowels: %d\n", numberOfVowels)
}

func countConsonants(reader *bufio.Reader, svc service.StringService) {
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading input:", err)
		return
	}
	text = strings.TrimSpace(text)
	numberOfConsonants, err := svc.CountConsonants(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Number of consonants: %d\n", numberOfConsonants)
}

func reverse(reader *bufio.Reader, svc service.StringService) {
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading input:", err)
		return
	}
	text = strings.TrimSpace(text)
	reversed, err := svc.Reverse(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reverse: %s\n", reversed)
}
