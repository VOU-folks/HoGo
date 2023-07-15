package globals

import (
	"hogo/lib/core/helpers"
)

var Args helpers.Args

func ReadArgs() {
	Args = helpers.GetArgs()
}
