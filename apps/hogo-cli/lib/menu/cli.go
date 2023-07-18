package menu

import (
	"hogo/apps/hogo-cli/lib/actions"
	"hogo/apps/hogo-cli/lib/menu/projects"
	"hogo/apps/hogo-cli/lib/menu/users"
)

func Cli() {
	task := ""
	for {
		task = SelectTask()

		switch task {
		case "Project management":
			projects.Entrance()
			break

		case "User management":
			users.Entrance()
			break

		case "x | Exit":
			actions.Exit()
		}

		Cli()
	}
}
