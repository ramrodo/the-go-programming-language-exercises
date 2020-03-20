package main

import (
	"fmt"

	"github.com/ramrodo/the-go-programming-language-exercises/pkg/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)

	fmt.Printf("FreezingC into Kelvin is: %v\n", tempconv.CtoK(tempconv.FreezingC))
}
