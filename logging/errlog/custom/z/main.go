package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world!")
	}
	http.HandleFunc("/test", handler)

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	server := &http.Server{
		Addr:     ":18888",
		ErrorLog: log.New(logger, "", 0),
	}

	logger.Fatal().Msgf("server: %v", server.ListenAndServe())
}
