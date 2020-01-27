package enums

type UserType int

const (
	Admin UserType = iota
	CustomerAdmin
	CustomerStaff
	Guest
)
