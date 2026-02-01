package ui

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/MishaAC/go-string-utils-cli/service"
)

func RunMenu(reader *bufio.Reader, svc service.StringService) {
	for {
		fmt.Println("===== STRING UTILITIES =====")
		fmt.Println("1) Count vowels")
		fmt.Println("2) Count consonants")
		fmt.Println("3) Reverse")
		fmt.Println("4) Is palindrome?")
		fmt.Println("5) Word count")
		fmt.Println("6) Capitalize")
		fmt.Println("7) To snake case")
		fmt.Println("8) Exit")

		fmt.Print("Choose an option: ")
		option, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			countVowels(reader, svc)
		case "2":
			fmt.Println(2)
		case "3":
			fmt.Println(3)
		case "4":
			fmt.Println(4)
		case "5":
			fmt.Println(5)
		case "6":
			fmt.Println(6)
		case "7":
			fmt.Println(7)
		case "8":
			fmt.Println("Bye 👋")
			return
		default:
			fmt.Println("Invalid option")
		}

		fmt.Println()
	}
}
