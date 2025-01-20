package object

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

const (
	BuiltinFuncNameLen   = "len"
	BuiltinFuncNamePuts  = "puts"
	BuiltinFuncNameFirst = "first"
	BuiltinFuncNameLast  = "last"
	BuiltinFuncNameRest  = "rest"
	BuiltinFuncNamePush  = "push"
)

func GetBuiltinByName(name string) *Builtin {

	switch name {

	case BuiltinFuncNamePuts:

		return &Builtin{
			Fn: func(args ...Object) Object {
				return NULL
			},
		}

	case BuiltinFuncNameFirst:

		return &Builtin{
			Fn: func(args ...Object) Object {
				return NULL
			},
		}

	case BuiltinFuncNameLast:

		return &Builtin{
			Fn: func(args ...Object) Object {
				return NULL
			},
		}

	case BuiltinFuncNameRest:

		return &Builtin{
			Fn: func(args ...Object) Object {
				return NULL
			},
		}

	case BuiltinFuncNamePush:

		return &Builtin{
			Fn: func(args ...Object) Object {
				return NULL
			},
		}

	default:

		return nil

	}

}
