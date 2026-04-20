package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/shafayetsadi/gator/internal/config"
	"github.com/shafayetsadi/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
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

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:     dbQueries,
		config: &cfg,
	}

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)

	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	err = commands.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
