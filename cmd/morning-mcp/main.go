// Command morning-mcp is the stdio MCP server for morning30.com.
// Agents use it to write custom exercises and upload a 30-day cycle
// into one account's personal JSON. It opens the same SQLite file as the app.
package main

import (
	"log"
	"os"

	"morning30/internal/app"
	"morning30/internal/db"
	"morning30/internal/mcp"
)

func main() {
	dbPath := os.Getenv("MORNING_DB")
	if dbPath == "" {
		dbPath = os.Getenv("DANI_DB")
	}
	if dbPath == "" {
		dbPath = "morning.db"
	}
	database, err := db.Open(dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer database.Close()
	app.DB = database
	app.SeedBuiltins()
	if err := mcp.ServeStdio(); err != nil {
		log.Fatal(err)
	}
}
