package main

import (
	"bufio"
	"fmt"
	"os"
	configs "project/app/configs"
)

func main() {
	configs.Block{
		Try: func() {
			var (
				input string
			)
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				input = scanner.Text()
				fmt.Printf("Input was: %s\n", input)
			}
		},
		Catch: func(e error) {
			fmt.Println(e)
		},
	}.Do()
}
