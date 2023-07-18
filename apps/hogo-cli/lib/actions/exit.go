package actions

import (
	"fmt"
	"os"
)

func Exit() {
	fmt.Println("Bye (:")
	os.Exit(0)
}
