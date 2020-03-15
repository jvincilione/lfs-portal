package enums

type JobStatus int
type MessageStatus int

const (
	Pending JobStatus = iota
	Accepted
	Scheduled
	AwaitingPayment
	Closed
	Rejected
)

const (
	Delivered MessageStatus = iota
	Viewed
)
