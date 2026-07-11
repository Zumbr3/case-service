package migrations

// Embeds the .sql migration files into the Go binary using the embed package.
// This allows cmd/migrate/main.go to run migrations without relying on external .sql files being present on disk in production.

import "embed"

//go:embed *.sql
var FS embed.FS
