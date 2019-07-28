package evaluator

import (
	"bufio"
	"fmt"
	"goRamble/object"
	"os"
	"strings"
)

// todo - add exit to close the interpreter
// todo - add import
// todo add builtins
var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1",
				len(args))
		}

		switch arg := args[0].(type) {
		case *object.Array:
			return &object.Integer{Value: int64(len(arg.Elements))}
		case *object.String:
			return &object.Integer{Value: int64(len(arg.Value))}
		default:
			return newError("argument to `len` not supported, got %s",
				args[0].Type())
		}
	},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
	"str": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			out := args[0].Inspect()
			return &object.String{Value: out}
		},
	},
	"type": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			out := strings.ToLower(string(args[0].Type()))
			return &object.String{Value: out}
		},
	},
	"open": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			path := ""
			if len(args) < 1 {
				return newError("wrong number of arguments. got=%d, want=1+",
					len(args))
			}

			// Get the filename
			switch args[0].(type) {
			case *object.String:
				path = args[0].(*object.String).Value
			default:
				return newError("argument to `file` not supported, got=%s",
					args[0].Type())

			}

			var lines []string
			file, err := os.Open(path)
			if err != nil {
				newError("the following error occured: %s", err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			newLines := strings.Join(lines, " ")
			return &object.String{Value: newLines}
		},
	},
	"exit": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			code := 0
			if len(args) > 0 {
				switch arg := args[0].(type) {
				case *object.Integer:
					code = int(arg.Value)
				}
			}
			os.Exit(code)
			return NULL
		},
	},
}

// todo page 230 - Test Arrays

//},
//"byte": &object.Builtin{
//	Fn: func(args ...object.Object) object.Object {
//		if len(args) != 1 {
//			return newError("wrong number of arguments. got=%d, want=1",
//				len(args))
//		}
//		//for _, arg := range args {
//		//	b := byte(arg.Type())
//		//}
//		out := bytes.Buffer{args[0].Inspect()}
//		return &object.Byte{Value: out}
//	},
//},
//"eval": &object.Builtin{
//	Fn: func(args ...object.Object) object.Object {
//		if len(args) != 1 {
//			return newError("wrong number of arguments. got=%d, want=1",
//				len(args))
//		}
//		//var env *object.Environment
//
//		switch args[0].(type) {
//		case *object.String:
//			txt := args[0].(*object.String).Value
//
//			// Lex the input
//			l := lexer.New(txt)
//
//			// parse it.
//			p := parser.New(l)
//
//			// If there are no errors
//			program := p.ParseProgram()
//			//if len(p.Errors()) == 0 {
//			//	// evaluate it, and return the output.
//			//	return (Eval(program, env))
//			//}
//
//			// Otherwise abort.  We should have try { } catch
//			// to allow this kind of error to be caught in the future!
//			//fmt.Printf("Error parsing eval-string: %s", txt)
//			//for _, msg := range p.Errors() {
//			//	fmt.Printf("\t%s\n", msg)
//			//}
//			//os.Exit(1)
//			return program
//		}
//		return newError("argument to `eval` not supported, got=%s",
//			args[0].Type())
//		},
