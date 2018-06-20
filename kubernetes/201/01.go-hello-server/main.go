package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "Hello, you've requested: %s (%s)\n", r.URL.Path, hostname)
	})

	http.ListenAndServe(":9000", nil)
}
