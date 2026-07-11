package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"github.com/Zumbr3/case-service/migrations"
)

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	if cmd == "" {
		cmd = "up"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		fmt.Fprintln(os.Stderr, "DATABASE_URL is not set")
		os.Exit(1)
	}

	src, err := iofs.New(migrations.FS, ".")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	m, err := migrate.NewWithSourceInstance("iofs", src, databaseURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch cmd {
	case "up":
		err = m.Up()
	case "down":
		err = m.Steps(-1)
	case "version":
		version, dirty, verr := m.Version()
		if verr != nil {
			err = verr
			break
		}
		fmt.Printf("version=%d dirty=%v\n", version, dirty)
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q (expected up|down|version)\n", cmd)
		os.Exit(1)
	}

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
