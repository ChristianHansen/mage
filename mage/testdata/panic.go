// +build mage

package main

import (
	"errors"

	"github.com/magefile/mage/mg"
)

// Function that panics.
func Panics() {
	panic("boom!")
}

// Error function that panics.
func PanicsErr() error {
	panic(errors.New("kaboom!"))
}

func PanicsInDep() {
	mg.Deps(Panics)
}
