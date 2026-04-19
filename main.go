package main

import (
	"log"
	"os"

	"github.com/shafayetsadi/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cfg := config.Config{}
	if err := cfg.Read(); err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		config: &cfg,
	}

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	err := commands.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
