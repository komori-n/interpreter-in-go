package evaluator

import "monkey/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
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
}
