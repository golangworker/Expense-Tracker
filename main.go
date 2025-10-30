package main

import (
	"app/logic"
	"fmt"
)

func main() {
	expence := logic.Init("Sosal", 6.66)
	fmt.Println(expence)
}
