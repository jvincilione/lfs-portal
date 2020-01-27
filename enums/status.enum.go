package enums

type JobStatus int
type MessageStatus int

const (
	Pending JobStatus = iota
	Scheduled
	AwaitingPayment
	Closed
)

const (
	Delivered MessageStatus = iota
	Viewed
)
