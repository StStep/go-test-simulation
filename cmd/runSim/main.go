package main

import (
	"bufio"
	"github.com/StStep/go-test-simulation/internal/state"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO Consider state and physics releationship

// our main function
func main() {
	print("Running simulation\r\n")

	// Make State
	st := state.New("config.json")

	// Setup logfile
	f, err := os.Create("space.logfmt")
	check(err)
	defer f.Close()
	st.Physics.SetLogOutput(bufio.NewWriter(f))

	// Make Unit
	st.NewUnit("swords", [2]float64{0, 0})

	// Update Physics
	for i := 1; i < 100; i++ {
		st.Physics.Step(0.01)
	}

	print("Finished simulation\r\n")
}
