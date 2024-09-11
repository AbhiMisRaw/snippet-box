package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	fmt.Println("Snippet-Box Server is starting... ✨")
	// fetching the address
	address := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
	}

	logger.Info("Starting Server on 🖳: ", "address", *address)
	err := http.ListenAndServe(*address, app.routes())
	// logging error message
	logger.Error(err.Error())
	os.Exit(1)
}
