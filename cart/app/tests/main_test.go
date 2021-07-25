package tests

import (
	"errors"
	configs "project/app/configs"
	functions "project/app/functions"
	structs "project/app/structs"
	"strings"
	"testing"
)

func TestAddProductScenario1(t *testing.T) {
	configs.Block{
		Try: func() {
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

func TestShowProduct(t *testing.T) {
	configs.Block{
		Try: func() {
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
