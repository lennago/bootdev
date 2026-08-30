package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/lennago/bootdotdev_gator/internal/config"
	"github.com/lennago/bootdotdev_gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	programState := &state{
		cfg: &cfg,
		db:  database.New(db),
	}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	if err = cmds.run(programState, cmd); err != nil {
		log.Fatal(err)
	}
}
