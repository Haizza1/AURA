package builtins

import (
	obj "aura/src/object"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var scanner = bufio.NewScanner(os.Stdin)

func wrongNumberofArgs(funcName string, found, actual int) *obj.Error {
	return &obj.Error{
		Message: fmt.Sprintf("numero incorrecto de argumentos para %s, se recibieron %d, se requieren %d", funcName, found, actual),
	}
}

func unsoportedArgumentType(funcname, objType string) *obj.Error {
	return &obj.Error{
		Message: fmt.Sprintf("argumento para %s no valido, se recibio %s", funcname, objType),
	}
}

func Longitud(args ...obj.Object) obj.Object {
	if len(args) != 1 {
		return wrongNumberofArgs("largo", len(args), 1)
	}

	switch arg := args[0].(type) {

	case *obj.String:
		return &obj.Number{Value: utf8.RuneCountInString(arg.Value)}

	case *obj.List:
		return &obj.Number{Value: len(arg.Values)}

	case *obj.Map:
		return &obj.Number{Value: len(arg.Store)}

	default:
		return unsoportedArgumentType("largo", obj.Types[args[0].Type()])
	}
}

func Escribir(args ...obj.Object) obj.Object {
	var buff strings.Builder

	for _, arg := range args {
		switch node := arg.(type) {

		case *obj.String:
			buff.WriteString(node.Inspect())

		case *obj.Number:
			buff.WriteString(node.Inspect())

		case *obj.List:
			buff.WriteString(node.Inspect())

		case *obj.Bool:
			buff.WriteString(node.Inspect())

		case *obj.Map:
			buff.WriteString(node.Inspect())

		default:
			return unsoportedArgumentType("escribir", obj.Types[node.Type()])
		}
	}

	fmt.Println(buff.String())
	return obj.SingletonNUll
}

func Recibir(args ...obj.Object) obj.Object {
	if len(args) > 1 {
		return wrongNumberofArgs("recibir", len(args), 1)
	}

	if len(args) == 0 {
		str := input(scanner)
		return &obj.String{Value: str}
	}

	if arg, isString := args[0].(*obj.String); isString {
		fmt.Print(arg.Inspect())
		str := input(scanner)
		return &obj.String{Value: str}
	}

	return unsoportedArgumentType("recibir", obj.Types[args[0].Type()])
}

func castInt(args ...obj.Object) obj.Object {
	if len(args) > 1 {
		return wrongNumberofArgs("entero", len(args), 1)
	}

	if arg, isString := args[0].(*obj.String); isString {
		return toInt(arg.Value)
	}

	return unsoportedArgumentType("entero", obj.Types[args[0].Type()])
}

func castString(args ...obj.Object) obj.Object {
	if len(args) > 1 {
		return wrongNumberofArgs("texto", len(args), 1)
	}

	if arg, isNumber := args[0].(*obj.Number); isNumber {
		strInt := strconv.Itoa(arg.Value)
		return &obj.String{Value: strInt}
	}

	return unsoportedArgumentType("recibir", obj.Types[args[0].Type()])
}

func RecibirEntero(args ...obj.Object) obj.Object {
	if len(args) > 1 {
		return wrongNumberofArgs("recibir_entero", len(args), 1)
	}

	if len(args) == 0 {
		strInt := input(scanner)
		return toInt(strInt)
	}

	if arg, isString := args[0].(*obj.String); isString {
		fmt.Print(arg.Inspect())
		strInt := input(scanner)
		return toInt(strInt)
	}

	return unsoportedArgumentType("recbir_entero", obj.Types[args[0].Type()])

}

func rango(args ...obj.Object) obj.Object {
	switch len(args) {
	case 1:
		return makeOneArgList(args[0])

	case 2:
		return makeTwoArgList(args[0], args[1])

	case 3:
		return makeTreArgList(args[0], args[1], args[2])

	default:
		return wrongNumberofArgs("rango", len(args), 2)
	}
}

func Tipo(args ...obj.Object) obj.Object {
	if len(args) > 1 || len(args) < 1 {
		return wrongNumberofArgs("tipo", len(args), 1)
	}

	return &obj.String{Value: obj.Types[args[0].Type()]}
}

func input(scan *bufio.Scanner) string {
	scan.Scan()
	str := scan.Text()
	return str
}

func toInt(str string) obj.Object {
	number, err := strconv.Atoi(str)
	if err != nil {
		return &obj.Error{Message: fmt.Sprintf("No se puede parsear como entero %s", str)}
	}

	return &obj.Number{Value: number}
}

var BUILTINS = map[string]*obj.Builtin{
	"largo":          obj.NewBuiltin(Longitud),
	"escribir":       obj.NewBuiltin(Escribir),
	"recibir":        obj.NewBuiltin(Recibir),
	"recibir_entero": obj.NewBuiltin(RecibirEntero),
	"tipo":           obj.NewBuiltin(Tipo),
	"entero":         obj.NewBuiltin(castInt),
	"texto":          obj.NewBuiltin(castString),
	"rango":          obj.NewBuiltin(rango),
	"agregar":        obj.NewBuiltin(add),
	"pop":            obj.NewBuiltin(pop),
	"popIndice":      obj.NewBuiltin(remove),
	"contiene":       obj.NewBuiltin(contains),
	"valores":        obj.NewBuiltin(values),
	"mayusculas":     obj.NewBuiltin(toUppper),
	"minusculas":     obj.NewBuiltin(toLower),
}