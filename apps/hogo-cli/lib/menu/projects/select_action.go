package projects

import (
	"github.com/AlecAivazis/survey/v2"
)

func SelectAction() string {
	action := ""

	err := survey.AskOne(
		&survey.Select{
			Message: "Please select action:",
			Options: []string{
				"Create new Project",
				"Find by Id",
				"Find by Name",
				"Find by DNS Zone",
				"< | Return",
				"x | Exit",
			},
		},
		&action,
	)
	if err != nil {
		panic(err)
	}

	return action
}
