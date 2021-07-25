package tests

import (
	"errors"
	"fmt"
	configs "project/app/configs"
	functions "project/app/functions"
	structs "project/app/structs"
	"strings"
	"testing"
)

/* unit testing */

func TestAddProductScenario1(t *testing.T) {
	configs.Block{
		Try: func() {
			fmt.Println("Add Product - Scenario 1 [Positive]: Add product from existing list -> expect: update qty only on the existing item from the list")
			var (
				inputName       = "nangka"
				inputQty  int64 = 1
			)
			carts := []*structs.Cart{
				{
					Name: inputName,
					Qty:  inputQty,
				},
			}
			carts, added := functions.AddProduct(inputName, carts)
			if !added {
				configs.Throw(errors.New("Adding is unsuccessfully"))
			}
			if len(carts) != 1 {
				configs.Throw(errors.New("Wrong cart length"))
			}
			var (
				exist = false
			)
			for _, cart := range carts {
				if strings.ToUpper(cart.Name) == strings.ToUpper(inputName) {
					exist = true
					if cart.Qty != int64(inputQty+1) {
						configs.Throw(errors.New("Qty on existing cart list (based on input) is wrong"))
					}
					break
				}
			}
			if !exist {
				configs.Throw(errors.New("Existing cart based on input is not found"))
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestAddProductScenario2(t *testing.T) {
	configs.Block{
		Try: func() {
			fmt.Println("Add Product - Scenario 2 [Positive]: Add product that doesn't exist on the list -> expect: add new item to the list")
			var (
				inputName       = "jambu"
				inputQty  int64 = 2
			)
			carts := []*structs.Cart{
				{
					Name: "nangka",
					Qty:  inputQty,
				},
			}
			carts, added := functions.AddProduct(inputName, carts)
			if !added {
				configs.Throw(errors.New("Adding is unsuccessfully"))
			}
			if len(carts) != 2 {
				configs.Throw(errors.New("Wrong cart length"))
			}
			for _, cart := range carts {
				if strings.ToUpper(cart.Name) == strings.ToUpper(inputName) {
					if cart.Qty != inputQty-1 {
						configs.Throw(errors.New("Qty on existing cart list (based on input) is wrong"))
					}
					break
				}
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestAddProductScenario3(t *testing.T) {
	configs.Block{
		Try: func() {
			fmt.Println("Add Product - Scenario 3 [Negative]: Add product that product name is empty -> expect: nothing change on the list")
			var (
				inputName = ""
			)
			carts := []*structs.Cart{
				{
					Name: "nangka",
					Qty:  2,
				},
			}
			carts, added := functions.AddProduct(inputName, carts)
			if added {
				configs.Throw(errors.New("Adding is successfully"))
			}
			if len(carts) != 1 {
				configs.Throw(errors.New("Wrong cart length"))
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestDeleteProductScenario1(t *testing.T) {
	configs.Block{
		Try: func() {
			fmt.Println("Delete Product - Scenario 1 [Positive]: Remove product from the list that the qty is more than 1 -> expect: the item still be exist but the qty is reduced")
			var (
				inputName = "nangka"
			)
			carts := []*structs.Cart{
				{
					Name: inputName,
					Qty:  2,
				},
				{
					Name: "jambu",
					Qty:  1,
				},
			}
			carts, deleted := functions.DeleteProduct(inputName, carts)
			if !deleted {
				configs.Throw(errors.New("Deleting is unsuccessfully"))
			}
			if len(carts) != 2 {
				configs.Throw(errors.New("Wrong cart length"))
			}
			var (
				exist = false
			)
			for _, cart := range carts {
				if strings.ToUpper(cart.Name) == strings.ToUpper(inputName) {
					exist = true
					if cart.Qty != 1 {
						configs.Throw(errors.New("Qty on existing cart list (based on input) is wrong"))
					}
					break
				}
			}
			if !exist {
				configs.Throw(errors.New("Existing cart based on input is not found"))
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestDeleteProductScenario2(t *testing.T) {
	configs.Block{
		Try: func() {
			fmt.Println("Delete Product - Scenario 2 [Positive]: Remove product from the list that the qty is the only 1 -> expect: the item should be removed from the list")
			var (
				inputName = "nangka"
			)
			carts := []*structs.Cart{
				{
					Name: inputName,
					Qty:  1,
				},
			}
			carts, deleted := functions.DeleteProduct(inputName, carts)
			if !deleted {
				configs.Throw(errors.New("Deleting is unsuccessfully"))
			}
			if len(carts) != 0 {
				configs.Throw(errors.New("Wrong cart length"))
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestDeleteProductScenario3(t *testing.T) {
	configs.Block{
		Try: func() {
			fmt.Println("Delete Product - Scenario 3 [Negative]: Remove product that doesn't exist on the list -> expect: nothing change on the list")
			var (
				inputName = "nangka"
			)
			carts := []*structs.Cart{
				{
					Name: "jambu",
					Qty:  1,
				},
			}
			carts, deleted := functions.DeleteProduct(inputName, carts)
			if deleted {
				configs.Throw(errors.New("Deleting is successfully"))
			}
			if len(carts) != 1 {
				configs.Throw(errors.New("Wrong cart length"))
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestShowProductScenario1(t *testing.T) {
	configs.Block{
		Try: func() {
			fmt.Println("Show Product - Scenario 1 [Positive]: Show existing item on the list -> expect: the item are exist on the list")
			carts := []*structs.Cart{
				{
					Name: "nangka",
					Qty:  1,
				},
			}
			carts, exist := functions.ShowProduct(carts)
			if !exist {
				configs.Throw(errors.New("The card is not exist"))
			}
			if len(carts) <= 0 {
				configs.Throw(errors.New("Wrong cart length"))
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestShowProductScenario2(t *testing.T) {
	configs.Block{
		Try: func() {
			fmt.Println("Show Product - Scenario 2 [Positive]: Show empty list -> expect: the item aren't exist on the list")
			carts := []*structs.Cart{}
			carts, exist := functions.ShowProduct(carts)
			if exist {
				configs.Throw(errors.New("The list is not empty"))
			}
			if len(carts) != 0 {
				configs.Throw(errors.New("Wrong cart length"))
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}
