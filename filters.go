package main

type FilterInterface interface {
	GetColumn() string
	GetValue() string
	Check(value string) bool
}

// Check if column matches value, % serves as a wildcard.
type EqualsFilter struct {
	Column string
	Value  string
}

func (f EqualsFilter) Check(value string) bool {
	return value == f.GetValue()
}

func (f EqualsFilter) GetColumn() string {
	return f.Column
}

func (f EqualsFilter) GetValue() string {
	return f.Value
}
