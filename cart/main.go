package main

import (
	"fmt"
	configs "project/app/configs"
)

func main() {
	configs.Block{
		Try: func() {

		},
		Catch: func(e error) {
			fmt.Println(e)
		},
	}.Do()
}
