package object

import "fmt"

type ObjectKind int

const (
	INTEGER ObjectKind = iota
	BOOLEAN
	RETURN_VALUE
	NULL
)

type Object interface {
	Kind() ObjectKind
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Kind() ObjectKind { return INTEGER }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Kind() ObjectKind { return BOOLEAN }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Kind() ObjectKind { return RETURN_VALUE }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Null struct{}

func (n *Null) Kind() ObjectKind { return NULL }
func (n *Null) Inspect() string  { return "null" }
