package evaluator

import (
	"fmt"
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Kind())
			}
		},
	},
	"exit": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				return newError("wrong number of arguments. got=%d, want<=1", len(args))
			}

			status := int(0)
			if len(args) == 1 {
				integ, ok := args[0].(*object.Integer)
				if !ok {
					return newError("argument to `exit` not supported, got %s", args[0].Kind())
				}
				status = int(integ.Value)
			}

			return &object.Exit{Status: status}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) > 0 {
					return arg.Elements[0]
				}
				return NULL
			default:
				return newError("arguments to `first` must be ARRAY, got %s", args[0].Kind())
			}
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				if length := len(arg.Elements); length > 0 {
					return arg.Elements[length-1]
				}
				return NULL
			default:
				return newError("arguments to `last` must be ARRAY, got %s", args[0].Kind())
			}
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				if length := len(arg.Elements); length > 0 {
					newElements := make([]object.Object, length-1, length-1)
					copy(newElements, arg.Elements[1:length])
					return &object.Array{Elements: newElements}
				}
				return NULL
			default:
				return newError("arguments to `rest` must be ARRAY, got %s", args[0].Kind())
			}
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				length := len(arg.Elements)
				newElements := make([]object.Object, length+1, length+1)
				copy(newElements, arg.Elements)
				newElements[length] = args[1]
				return &object.Array{Elements: newElements}
			default:
				return newError("arguments to `push` must be ARRAY, got %s", args[0].Kind())
			}
		},
	},
	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
	"kind": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			fmt.Println(args[0].Kind())

			return NULL
		},
	},
}
