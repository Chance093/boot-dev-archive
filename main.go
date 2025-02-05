package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Chance093/gator/internal/config"
	"github.com/Chance093/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
  db *database.Queries
	cfg *config.Config
}

func main() {
  // Read config from root dir
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

  // Open pg db connection
  db, err := sql.Open("postgres", cfg.DB_URL)
  if err != nil {
    log.Fatalf("error connecting to db: %v", err)
  }
  defer db.Close()

  // Init application state
  s := &state{
    db: database.New(db),
    cfg: cfg,
  }

  // Register commands for cli
  cmds := commands{
    registeredCommands: make(map[string]func(*state, command) error),
  }
  cmds.register("login", handlerLogin)
  cmds.register("register", handlerRegister)
  cmds.register("reset", handlerReset)
  cmds.register("users", handlerListUsers)
  cmds.register("agg", handlerAgg)
  cmds.register("addfeed", handlerAddFeed)
  cmds.register("feeds", handlerListFeeds)
  cmds.register("follow", handlerFollow)
  cmds.register("following", handlerFollowing)

  // Run command
	rawArgs := os.Args[1:]
	if len(rawArgs) < 1 {
		log.Fatal(err)
	}

	cmdName, cmdArgs := rawArgs[0], rawArgs[1:]

	if err := cmds.run(s, command{cmdName, cmdArgs}); err != nil {
		log.Fatal(err)
	}
}
