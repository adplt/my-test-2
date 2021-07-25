package functions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	configs "project/app/configs"
	structs "project/app/structs"
	variables "project/app/variables"
	"runtime"
	"strings"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			configs.Throw(err)
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			configs.Throw(err)
		}
	}
}

func AddProductScreen(carts []*structs.Cart) []*structs.Cart {
	var (
		input string
		add   = false
	)
	ClearScreen()
	fmt.Print("Insert your product that you want to add: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	for i, cart := range carts {
		if strings.ToUpper(cart.Name) == strings.ToUpper(input) {
			carts[i].Qty += 1
			add = true
			break
		}
	}
	if !add {
		cart := &structs.Cart{
			Name: input,
			Qty:  1,
		}
		carts = append(carts, cart)
	}
	fmt.Println("\nYour input was successfully added")
	fmt.Println("Enter to back....")
	fmt.Scanln(&input)
	return carts
}

func DeleteProductScreen(carts []*structs.Cart) []*structs.Cart {
	var (
		input  string
		delete = false
		index  = -1
	)
	ClearScreen()
	fmt.Print("Insert your product that you want to delete: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	for i, cart := range carts {
		if strings.ToUpper(cart.Name) == strings.ToUpper(input) {
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
	if !delete {
		fmt.Println("\nNo one on the list was deleted, your request not found on the list")
	} else {
		fmt.Println("\nYour input was successfully delete")
	}
	fmt.Println("Enter to back....")
	fmt.Scanln(&input)
	return carts
}

func GetProductScreen(carts []*structs.Cart) {
	var (
		input string
	)
	ClearScreen()
	if len(carts) == 0 {
		fmt.Println("Cart is empty")
	} else {
		cartJson, err := json.MarshalIndent(carts, "", "  ")
		if err != nil {
			configs.Throw(err)
		}
		fmt.Println(string(cartJson))
	}
	fmt.Print("\n\nBack [yes/no]: ")
	fmt.Scanln(&input)
	switch input {
	case "yes":
		Dashboart(carts)
		break
	case "no":
		break
	default:
		GetProductScreen(carts)
		break
	}
}

func Dashboart(carts []*structs.Cart) {
	var (
		input string
	)
	ClearScreen()
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("******************************* Software Library *******************************\n\n")
	fmt.Println("Choose your options: ")
	fmt.Println(variables.MENU__ADD_PRODUCT + ". Add Product")
	fmt.Println(variables.MENU__DELETE_PRODUCT + ". Delete Product")
	fmt.Println(variables.MENU__SHOW_PRODUCT + ". Show Product")
	fmt.Print("Answer: ")
	fmt.Scanln(&input)
	switch input {
	case variables.MENU__ADD_PRODUCT:
		carts = AddProductScreen(carts)
		break
	case variables.MENU__DELETE_PRODUCT:
		carts = DeleteProductScreen(carts)
		break
	case variables.MENU__SHOW_PRODUCT:
		GetProductScreen(carts)
		break
	default:
		break
	}
	Dashboart(carts)
}
