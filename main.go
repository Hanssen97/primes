package main

import (
	"github.com/jorgenhanssen/primes/cli"
	p "github.com/jorgenhanssen/primes/primes"
)

func main() {
	primes := p.New()

	program := cli.New(cli.Debug, []cli.Routine{
		{
			Command: "F",
			Description: "Find primes in an interval",
			Function: primes.Find,
		},
		{
			Command: "P",
			Description: "Print primes",
			Function: primes.Print,
		},
		{
			Command: "S",
			Description: "Save primes to file",
			Function: primes.Save,
		},
	})
	
	primes.Program = program

	program.Run()
}