package main

import (
	"fmt"
	"os"

	"weather-app-tui/pkg"
)

func main() {
	if err := pkg.Program(); err != nil {
		fmt.Println("Program failed to Run: ", err)
		os.Exit(1)
	}
}
