package main

import (
	"os"
	"github.com/juliecoding/svg-fun-go/cli"
)

func main() {
    os.Exit(cli.Run(os.Args[1:]))
}
