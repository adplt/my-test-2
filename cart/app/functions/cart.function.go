package functions

import (
	"encoding/json"
	"fmt"
	configs "project/app/configs"
	structs "project/app/structs"
	"strings"
)

func AddProduct(product string, carts []*structs.Cart) ([]*structs.Cart, bool) {
	var (
		add = false
	)
	if product == "" {
		return carts, add
	}
	for i, cart := range carts {
		if strings.ToUpper(cart.Name) == strings.ToUpper(product) {
			carts[i].Qty += 1
			add = true
			break
		}
	}
	if !add {
		cart := &structs.Cart{
			Name: product,
			Qty:  1,
		}
		carts = append(carts, cart)
		add = true
	}
	return carts, add
}

func DeleteProduct(product string, carts []*structs.Cart) ([]*structs.Cart, bool) {
	var (
		delete = false
		index  = -1
	)
	for i, cart := range carts {
		if strings.ToUpper(cart.Name) == strings.ToUpper(product) {
			if carts[i].Qty == 1 {
				index = i
			} else {
				carts[i].Qty -= 1
				delete = true
			}
			break
		}
	}
	if index != -1 {
		if len(carts) == 1 {
			carts = []*structs.Cart{}
		} else {
			carts = append(carts[:index], carts[index+1:]...)
		}
		delete = true
	}
	return carts, delete
}

func ShowProduct(carts []*structs.Cart) ([]*structs.Cart, bool) {
	var (
		exist bool
	)
	if len(carts) == 0 {
		fmt.Println("Cart is empty")
		exist = false
	} else {
		cartJson, err := json.MarshalIndent(carts, "", "  ")
		if err != nil {
			configs.Throw(err)
		}
		fmt.Println(string(cartJson))
		exist = true
	}
	return carts, exist
}
