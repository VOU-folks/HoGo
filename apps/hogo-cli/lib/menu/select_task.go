package menu

import (
	"github.com/AlecAivazis/survey/v2"
)

func SelectTask() string {
	task := ""

	err := survey.AskOne(
		&survey.Select{
			Message: "Please select task:",
			Options: []string{
				"User management",
				"Project management",
				"x | Exit",
			},
		},
		&task,
	)
	if err != nil {
		panic(err)
	}

	return task
}
