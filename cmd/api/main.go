package main

import (
	"fmt"
	"nuage/internal/http"
)

func main() {
	server := http.NewServer()

	// test
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
