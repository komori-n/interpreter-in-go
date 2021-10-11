package object

import "fmt"

type ObjectKind int

const (
	INTEGER ObjectKind = iota
	BOOLEAN
	RETURN_VALUE
	ERROR
	NULL
)

func (ok ObjectKind) String() string {
	switch ok {
	case INTEGER:
		return "INTEGER"
	case BOOLEAN:
		return "BOOLEAN"
	case RETURN_VALUE:
		return "RETURN_VALUE"
	case ERROR:
		return "ERROR"
	case NULL:
		return "NULL"
	default:
		return "<error kind>"
	}
}

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

type Error struct {
	Message string
}

func (e *Error) Kind() ObjectKind { return ERROR }
func (e *Error) Inspect() string  { return "Error: " + e.Message }

type Null struct{}

func (n *Null) Kind() ObjectKind { return NULL }
func (n *Null) Inspect() string  { return "null" }
