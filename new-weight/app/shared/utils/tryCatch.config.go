package utils

import (
	"errors"
	"fmt"
)

// Block struct
type Block struct {
	Try     func()
	Catch   func(error)
	Finally func()
}

// Throw function
func Throw(up error) {
	panic(up)
}

// Do function
func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(errors.New(fmt.Sprint(r)))
			}
		}()
	}
	tcf.Try()
}
