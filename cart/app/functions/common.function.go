package functions

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	configs "project/app/configs"
	structs "project/app/structs"
	variables "project/app/variables"
	"runtime"
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
	)
	ClearScreen()
	fmt.Print("Insert your product that you want to add: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	carts, added := AddProduct(input, carts)
	if added {
		fmt.Println("\nYour input was successfully added")
	} else {
		fmt.Println("\nYour request unsuccessfully added")
	}
	fmt.Println("Enter to back....")
	fmt.Scanln(&input)
	return carts
}

func DeleteProductScreen(carts []*structs.Cart) []*structs.Cart {
	var (
		input string
	)
	ClearScreen()
	fmt.Print("Insert your product that you want to delete: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	carts, deleted := DeleteProduct(input, carts)
	if !deleted {
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
	carts, _ = ShowProduct(carts)
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
	fmt.Println(variables.MENU__EXIT + ". Exit")
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
	case variables.MENU__EXIT:
		break
	default:
		Dashboart(carts)
		break
	}
}
