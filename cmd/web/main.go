package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Define application struct that holds application-wide dependencies
type application struct {
	logger *slog.Logger
}

func main() {
	// Define new command-line flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Read in command line argument
	// Need to use this before using addr
	// If any errors occur, program will terminate
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info("Starting server", slog.String("addr", *addr))

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
