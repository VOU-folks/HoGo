package users

import (
	"fmt"
	"hogo/apps/hogo-cli/lib/actions"
)

func Entrance() {
	fmt.Println("User management")
	fmt.Println("--------")

	for {
		action := SelectAction()
		switch action {
		case "< | Return":
			return
		case "x | Exit":
			actions.Exit()
			return
		}
	}
}
