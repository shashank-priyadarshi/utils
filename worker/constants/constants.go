package constants

import "go.ssnk.in/utils/worker/types"

const (
	Active    types.Status = "active"
	Inactive  types.Status = "inactive"
	Queued    types.Status = "inactive"
	Completed types.Status = "completed"
	Cancelled types.Status = "cancelled"
	Timeout   types.Status = "timeout"
	Invalid   types.Status = "invalid"
)
