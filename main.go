package main

import (
	"os"

	"github.com/rydelll/papermc/cmd"
)

func main() {
	ret := cmd.Execute()
	os.Exit(ret)
}
