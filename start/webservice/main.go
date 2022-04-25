package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	// ①
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
		log.Info().Msg("receive hello world request")
	})

	// ②
	fmt.Println("Start listening at :8080")
	// ③
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// ④
		fmt.Fprintf(os.Stderr, "")
		io.WriteString(os.Stderr, err.Error())
		os.Exit(1)
	}
}
