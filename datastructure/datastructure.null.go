package datastructure

type Null struct{}

func (n *Null) Inspect() string { return "null" }
func (n *Null) Type() DataType  { return NULL }
