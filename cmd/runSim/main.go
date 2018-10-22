package main

import (
	_ "github.com/StStep/go-test-simulation/internal/configuration"
	_ "github.com/StStep/go-test-simulation/internal/factory"
	_ "github.com/StStep/go-test-simulation/internal/id"
	_ "github.com/StStep/go-test-simulation/internal/physics"
	_ "github.com/StStep/go-test-simulation/internal/state"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// our main function
func main() {
	print("Running simulation\r\n")

	// Make Conf, Ledger, Id gens

	// Make Physics

	// Make Constructor

	// Make Unit
}
