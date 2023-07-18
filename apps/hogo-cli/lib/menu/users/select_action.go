package users

import (
	"github.com/AlecAivazis/survey/v2"
)

func SelectAction() string {
	action := ""

	err := survey.AskOne(
		&survey.Select{
			Message: "Please select action:",
			Options: []string{
				"Create new User",
				"Find by Id",
				"Find by Username",
				"Find by Email",
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
