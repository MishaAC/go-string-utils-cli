package main

import (
	"bufio"
	"os"

	"github.com/MishaAC/go-string-utils-cli/service"
	"github.com/MishaAC/go-string-utils-cli/ui"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	service := service.NewStringService()

	ui.RunMenu(reader, service)
}
