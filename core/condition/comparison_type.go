package condition

type ComparisonType int

const (
	ComparisonTypeEqual = ComparisonType(0)
	ComparisonTypeNotEqual = ComparisonType(1)
	ComparisonTypeIn = ComparisonType(2)
)
