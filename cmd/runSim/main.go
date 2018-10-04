package main

import (
	"bufio"
	"github.com/StStep/go-test-simulation/internal/space"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// our main function
func main() {
	print("Running simulation\r\n")
	f, err := os.Create("space.logfmt")
	check(err)
	defer f.Close()

	s := space.NewSpace()
	s.SetLogOutput(bufio.NewWriter(f))
	e1 := s.RegisterEntity([2]float64{0, 0}, 0.25)
	e2 := s.RegisterEntity([2]float64{1, 0}, 0.25)

	s.UpdateEntity(e1, [2]float64{0, 4})
	s.UpdateEntity(e2, [2]float64{0, 4})
	for k := 0; k < 100; k++ {
		s.Step(0.01)
	}
}
