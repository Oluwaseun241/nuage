package main

import (
	"fmt"
	"nuage/internal/server"
)

func main() {
	server := server.NewServer()

	// test
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
