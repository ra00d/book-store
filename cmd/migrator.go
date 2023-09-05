// This is custom goose binary with sqlite3 support only.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	_ "github.com/ra00d/book_store/database/migrations"
	_ "modernc.org/sqlite"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./database/migrations/", "directory with migration files")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	cwd, _ := os.Getwd()
	command := args[0]
	db, err := goose.OpenDBWithDriver("sqlite", fmt.Sprintf("%s/%s", cwd, "./database/test.db"))
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()
	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
