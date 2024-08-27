package main

import (
    "nyiso-electricity-data-analysis/pkg/api"
    "nyiso-electricity-data-analysis/pkg/data"
    "nyiso-electricity-data-analysis/config"
    "log"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    // Initialize database connection
    db, err := data.NewDBConnection(cfg)
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    defer db.Close()

    // Set up and start the API
    api.StartServer(db)
}
