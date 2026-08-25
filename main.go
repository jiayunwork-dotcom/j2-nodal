package main

import (
	"os"

	"j2-nodal/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
