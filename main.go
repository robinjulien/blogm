package main

import (
	"os"

	"github.com/robinjulien/rblog/cli"
)

//go:generate go run scripts/includeAssets.go

func main() {
	cli.Execute(&os.Args)
}
