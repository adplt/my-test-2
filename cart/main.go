package main

import (
	"fmt"
	configs "project/app/configs"
	functions "project/app/functions"
	structs "project/app/structs"
)

func main() {
	configs.Block{
		Try: func() {
			carts := []*structs.Cart{}
			functions.Dashboart(carts)
		},
		Catch: func(e error) {
			fmt.Println(e)
		},
	}.Do()
}
