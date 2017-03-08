package cmd

import (
	"log"
)

// Command is the interface a command must satisfy
type Command interface {
	Run()
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
	return c
}

// Run begins the command
func (c *Subcommands) Run(cmd string) {
	if comm, ok := c.cmds[cmd]; ok {
		comm.cmd.Run()
		return
	}

	log.Fatalf("Unknown subcommand '%s'", cmd)
}

// Register adds a new subcommand
func (c *Subcommands) Register(name, desc string, cmd Command) {
	c.cmds[name] = info{desc: desc, cmd: cmd}
}

type info struct {
	desc string
	cmd  Command
}
