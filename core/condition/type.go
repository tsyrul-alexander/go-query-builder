package condition

type Type int

const (
	Group     = Type(0)
	Column    = Type(1)
	Parameter = Type(2)
	Binary    = Type(3)
)
