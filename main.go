package main

import (
	"github.com/rydelll/papermc/cmd"
	_ "github.com/rydelll/papermc/cmd/folia"
	_ "github.com/rydelll/papermc/cmd/paper"
	_ "github.com/rydelll/papermc/cmd/velocity"
	_ "github.com/rydelll/papermc/cmd/waterfall"
)

func main() {
	cmd.Execute()
}
