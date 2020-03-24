// Cf converts its numeric argument to Celsius and Fahrenheit
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/ramrodo/the-go-programming-language-exercises/pkg/tempconv"
)

func convertNumberToTemperatures(number string) {
	t, err := strconv.ParseFloat(number, 64)

	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)

	fmt.Printf("%s = %s, %s = %s\n\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		for _, arg := range os.Args[1:] {
			convertNumberToTemperatures(arg)
		}
	} else {
		for text := ""; text == ""; {
			fmt.Print("Enter a number: ")
			scanner := bufio.NewScanner(os.Stdin)
			// Scans a line from Stdin(Console)
			scanner.Scan()
			// Holds the string that scanned
			text := scanner.Text()

			if text == "q" || text == "exit" {
				os.Exit(1)
			} else if _, err := strconv.ParseFloat(text, 64); text != "" && err == nil {
				convertNumberToTemperatures(text)
			}
		}
	}
}
