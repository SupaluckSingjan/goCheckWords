package main

import (
	"fmt"

	"github.com/SupaluckSingjan/goCheckWords/pkg"
)

func main() {
	var nameDay string

	fmt.Println("Input Name of day")
	fmt.Scanf("%s", nameDay)

	sayword.SayWord(nameDay)
}