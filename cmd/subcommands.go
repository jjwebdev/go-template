package cmd

import (
	"fmt"
	"os"
	"strings"

	"log"
)

// Command is the interface a command must satisfy
type Command interface {
	Run(args []string)
	Usage() string
}

// Subcommands is the instance of commands
type Subcommands struct {
	name, desc string
	cmds       map[string]info
}

// New creates a new subcommand object
func New(name, desc string) *Subcommands {
	c := &Subcommands{
		name: name,
		desc: desc,
		cmds: make(map[string]info),
	}

	c.Register("help", "Describe a subcommand", help{c})

	return c
}

// Run begins the command
func (c *Subcommands) Run(args []string) {
	if len(args) == 0 {
		args = []string{"help"}
	}

	for cmd, f := range c.cmds {
		if args[0] == cmd {
			f.cmd.Run(args[1:])
			return
		}
	}

	if os.Args[1][0] == '-' {
		log.Fatalf("flag provided but not defined '%s'", os.Args[1])
	}

	log.Fatalf("Unknown subcommand '%s'", os.Args[1])
}

// Register adds a new subcommand
func (c *Subcommands) Register(name, desc string, cmd Command) {
	c.cmds[name] = info{desc: desc, cmd: cmd}
}

type info struct {
	desc string
	cmd  Command
}

type help struct {
	c *Subcommands
}

// Run will run the help command
func (h help) Run(args []string) {
	if len(args) == 1 {
		c, ok := h.c.cmds[args[0]]
		if !ok {
			log.Fatalln("unknown help topic:", args[0])
		}

		usage := c.cmd.Usage()
		if usage == "" {
			log.Fatalln("Unknown help topic:", args[0])
		}
		log.Println(usage)
		return
	}

	var commandsDesc string
	width := 0
	for cmd := range h.c.cmds {
		if len(cmd) > width {
			width = len(cmd)
		}
	}

	for cmd := range h.c.cmds {
		commandsDesc = fmt.Sprintf("%s\n    %s%s%s",
			commandsDesc,
			cmd,
			strings.Repeat(" ", width+4-len(cmd)),
			h.c.cmds[cmd].desc)
	}
	log.Fatalln(h.c.name, h.c.desc, commandsDesc)
}

func (h help) Usage() string {
	return "help describes a help topic"
}
