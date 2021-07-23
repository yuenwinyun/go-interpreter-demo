package datastructure

type DataType string

type DataStructure interface {
	Type() DataType
	Inspect() string
}

const (
	INTEGER  = "INTEGER"
	BOOLEAN  = "BOOLEAN"
	NULL     = "NULL"
	FUNCTION = "FUNCTION"
)
