package main

import (
	"hogo/apps/hogo-cli/lib/actions"
	"hogo/apps/hogo-cli/lib/menu"
)

func main() {
	menu.Greetings()
	menu.Cli()
	actions.Exit()
}
