package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

type ObjectKind int

const (
	INTEGER ObjectKind = iota
	BOOLEAN
	RETURN_VALUE
	ERROR
	FUNCTION
	BUILTIN
	STRING
	EXIT
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
	case FUNCTION:
		return "FUNCTION"
	case BUILTIN:
		return "BUILTIN"
	case STRING:
		return "STRING"
	case EXIT:
		return "EXIT"
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

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Kind() ObjectKind { return FUNCTION }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type BuiltinFunction func(args ...Object) Object
type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Kind() ObjectKind { return BUILTIN }
func (b *Builtin) Inspect() string  { return "builtin function" }

type String struct {
	Value string
}

func (s *String) Kind() ObjectKind { return STRING }
func (s *String) Inspect() string  { return s.Value }

type Exit struct {
	Status int
}

func (ex *Exit) Kind() ObjectKind { return EXIT }
func (ex *Exit) Inspect() string  { return fmt.Sprintf("exit(%d)", ex.Status) }

type Null struct{}

func (n *Null) Kind() ObjectKind { return NULL }
func (n *Null) Inspect() string  { return "null" }
