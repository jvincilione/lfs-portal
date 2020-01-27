package enums

type Action int

const (
	Read Action = iota
	Write
	Delete
	Execute
)
