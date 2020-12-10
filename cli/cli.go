package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Errors
var ErrNotValidCommand = errors.New("Not a valid command");

type Routine struct {
	Command string
	Description string
	Function func(cli *Instance)
}

type Verbosity string
const (
	Normal Verbosity = "Normal"
	Debug Verbosity = "Debug"
)

type Instance struct {
	verbosity Verbosity
	routines []Routine
	running bool
}

func New(verbosity Verbosity, routines []Routine) (c *Instance) {
	c = &Instance{
		verbosity,
		routines,
		false,
	}

	c.routines = append(c.routines, Routine{
		Command: "C",
		Description: "Clear the screen",
		Function: clearTerminal,
	})
	c.routines = append(c.routines, Routine{
		Command: "E",
		Description: "Exit program",
		Function: stop,
	})

	return
}

func (c *Instance) Run() {
	clearTerminal(c)
	c.running = true
	
	for c.running {
		c.printMenu()
		if err := c.runNext(); err != nil {
			c.WriteError("%s\n", err.Error())
		}
	}
}

func (c *Instance) runNext() error {
	command := strings.ToUpper(c.Input(""))

	for _, r := range c.routines {
		if command == r.Command {
			c.Write("\n\n")
			r.Function(c)
			c.Write("\n\n")
			return nil
		}
	}

	return ErrNotValidCommand
}

func (c *Instance) printMenu() {
	c.Write(ColorWhite("\nMenu\n\n"))
	for _, r := range c.routines {
		c.Write("  %s%s%s %s\n", 
			ColorYellow("["),
			ColorBlue(r.Command),
			ColorYellow("]"),
			r.Description,
		)
	}
	c.Write("\n")
}

func (c *Instance) Input(prefix string) string {
	reader := bufio.NewReader(os.Stdin)
	c.Write(ColorBlue(fmt.Sprintf("%s > ", prefix)))
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}

func (c *Instance) InputInt(prefix string) int {
	for {
		text := c.Input(prefix)
		if num, err := strconv.Atoi(text); err == nil {
			return num
		}

		c.WriteError("Input must be integer\n\n")
	}
}

func (c *Instance) Write(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (c *Instance) WriteError(format string, args ...interface{}) {
	c.Write(ColorRed(fmt.Sprintf(format, args...)))
}

func (c *Instance) Debug(format string, args ...interface{}) {
	if c.verbosity == Debug {
		c.Write(ColorYellow(fmt.Sprintf(format, args...)))
	}
}

func clearTerminal(program *Instance) {
	for i := 0; i < 100; i++ {
		program.Write("\n")
	}
}

func stop(program *Instance) {
	program.running = false
}