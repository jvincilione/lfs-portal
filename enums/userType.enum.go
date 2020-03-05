package enums

type UserType int

const (
	Admin UserType = iota
	CompanyAdmin
	CompanyStaff
	Guest
)
