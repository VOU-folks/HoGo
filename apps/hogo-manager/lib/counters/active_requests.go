package counters

import "hogo/lib/core/helpers"

var ActiveRequests *helpers.Counter

func init() {
	ActiveRequests = helpers.NewCounter()
	ActiveRequests.Reset()
}
