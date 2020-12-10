package primes

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/jorgenhanssen/primes/cli"
	"github.com/jorgenhanssen/primes/utils"
)

var primes = []int{}

func (p *Instance) Find(program *cli.Instance) {
	program.Write(cli.ColorWhite("Find primes:\n"))

	start := program.InputInt("From")
	end := program.InputInt("To")

	startedAt := time.Now()
	primes = p.findPrimes(start, end)
	program.Write("\nFound %s primes in %s", 
		cli.ColorGreen(len(primes)), 
		cli.ColorYellow(utils.PrettyDuration(time.Since(startedAt))))
}

func (p *Instance) Print(program *cli.Instance) {
	program.Write("%v", primes)
}

func (p *Instance) Save(program *cli.Instance) {
	fileName := "primes.txt"
	program.Write("Saving primes to %s...", fileName)
	data := []byte(fmt.Sprint(primes))
	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		program.WriteError(err.Error())
	}
}