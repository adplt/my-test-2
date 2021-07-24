package tests

import (
	configs "project/app/configs"
	"testing"
)

func TestAddProduct(t *testing.T) {
	configs.Block{
		Try: func() {
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestDeleteProduct(t *testing.T) {
	configs.Block{
		Try: func() {
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestShowProduct(t *testing.T) {
	configs.Block{
		Try: func() {
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}
