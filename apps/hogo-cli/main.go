package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"log"
	"strings"
)

func main() {
	Greetings()
	Cli()
	Exit()
}

func Greetings() {
	fmt.Println("Welcome to HoGo Cli interface")
}

func Exit() {
	fmt.Println("Bye (:")
}

func Cli() {
	task := ""
	for {
		if IsValidTask(task) {
			break
		}
		task = SelectTask()
	}

	if task == "Exit" {
		return
	}

	survey.Title("--------")
	Cli()
}

func SelectTask() string {
	task := ""

	err := survey.AskOne(
		&survey.Select{
			Message: "Please select task:",
			Options: []string{
				"User management",
				"Project management",
				"Exit",
			},
		},
		&task,
	)
	if err != nil {
		log.Fatal("SelectTask:", err.Error())
	}

	return task
}

func IsValidTask(task string) bool {
	task = strings.TrimSpace(task)
	return task != ""
}
