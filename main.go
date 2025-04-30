// Package main is the entrypoint for the CLI.
package main

import (
	"log"
	"os"

	"github.com/tx3stn/pkb/cmd"
)

func main() {
	code := 0
	defer func() {
		os.Exit(code)
	}()

	if err := cmd.Execute(); err != nil {
		code = 1

		log.Printf("%s\n", err.Error())
	}
}
