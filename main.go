package main

import (
	"errors"
	"fmt"
	"os"
)

var TheBigBangError error = errors.New("The Big Bang Error")

func BigBang() error {
	return TheBigBangError
}

func CreateWormhole() error {
	if err := BigBang(); err != nil {
		return fmt.Errorf("An accident: %w. Wormhole unstable.", err)
	}
	return nil
}

func main() {
	err := CreateWormhole()
	if errors.Is(err, TheBigBangError) {
		fmt.Printf("Invalid Universe. %s\n", err)
		os.Exit(1)
	}
}
