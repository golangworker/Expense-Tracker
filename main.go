package main

import (
	"app/term"
	"fmt"
	"os"
)

func main() {
	err := term.RunningLoop(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

// version 2.11.2025 by Vladimir Aristarkhov
