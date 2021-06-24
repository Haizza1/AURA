package evaluator

import (
	obj "aura/src/object"
	"fmt"
)

// utils functions to return errors

// generates a new error instance
func newError(message string) *obj.Error {
	return &obj.Error{Message: message}
}

func typeMismatchError(left, operator, rigth string) *obj.Error {
	return &obj.Error{
		Message: fmt.Sprintf("Discrepancia de tipos: %s %s %s", left, operator, rigth),
	}
}

func unknownPrefixOperator(operator, rigth string) *obj.Error {
	return &obj.Error{
		Message: fmt.Sprintf("Operador desconocido: %s%s", operator, rigth),
	}
}

func unknownInfixOperator(left, operator, rigth string) *obj.Error {
	return &obj.Error{
		Message: fmt.Sprintf("Operador desconocido: %s %s %s", left, operator, rigth),
	}
}

func unknownIdentifier(identifier string) *obj.Error {
	return &obj.Error{
		Message: fmt.Sprintf("Identificador no encontrado: %s", identifier),
	}
}

func notAFunction(identifier string) *obj.Error {
	return &obj.Error{
		Message: fmt.Sprintf("No es una funcion: %s", identifier),
	}
}

func notAList(identifier string) *obj.Error {
	return &obj.Error{
		Message: fmt.Sprintf("No es una lista: %s", identifier),
	}
}

func notAVariable(identifier string) *obj.Error {
	return &obj.Error{
		Message: fmt.Sprintf("No es una variable: %s", identifier),
	}
}

func noSuchMethod(method, ident string) *obj.Error {
	return &obj.Error{
		Message: fmt.Sprintf("%s no tiene un metodo %s", ident, method),
	}
}

func notIterable(ident string) *obj.Error {
	return &obj.Error{Message: fmt.Sprintf("No es un iteralble: %s", ident)}
}