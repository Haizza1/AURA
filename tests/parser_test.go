package test_test

import (
	"fmt"
	"lpp/lpp"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseProgram(t *testing.T) {
	assert := assert.New(t)
	source := "var x = 5;"
	lexer := lpp.NewLexer(source)
	parser := lpp.NewParser(lexer)
	program := parser.ParseProgam()

	if !assert.NotNil(program) {
		t.Log("program is nil")
		t.Fail()
	}

	if !assert.IsType(lpp.Program{}, program) {
		t.Log("program type is not Program!")
		t.Fail()
	}

	if !assert.Implements((*lpp.ASTNode)(nil), program) {
		t.Log("program type does not implement ast node interface")
		t.Fail()
	}
}

func TestLetStatements(t *testing.T) {
	assert := assert.New(t)
	source := `
		var x = 5;
		var y = 10;
		var foo = 20;
	`
	lexer := lpp.NewLexer(source)
	parser := lpp.NewParser(lexer)
	program := parser.ParseProgam()

	if !assert.Equal(3, len(program.Staments)) {
		t.Log("len of program statements are not 3")
		t.Fail()
	}

	for _, statement := range program.Staments {
		if !assert.Equal("var", statement.TokenLiteral()) {
			t.Log("token are not a variable")
			t.Fail()
		}

		if !assert.IsType(&lpp.LetStatement{}, statement.(*lpp.LetStatement)) {
			t.Log("statement are not let statement type")
			t.Fail()
		}
	}
}

func TestNamesInLetStatements(t *testing.T) {
	assert := assert.New(t)
	source := `
		var x = 5;
		var y = 10;
		var foo = 20;
	`
	lexer := lpp.NewLexer(source)
	parser := lpp.NewParser(lexer)
	program := parser.ParseProgam()

	var names []string
	for _, stament := range program.Staments {
		stament := stament.(*lpp.LetStatement)
		if !assert.NotNil(stament.Name) {
			t.Fail()
		}

		if !assert.Implements((*lpp.Stmt)(nil), stament) {
			t.Fail()
		}

		names = append(names, stament.Name.Str())
	}

	expectedNames := []string{"x", "y", "foo"}
	if !assert.Equal(expectedNames, names) {
		t.Fail()
	}

}

func TestParseErrors(t *testing.T) {
	source := "var x 5;"
	lexer := lpp.NewLexer(source)
	parser := lpp.NewParser(lexer)
	parser.ParseProgam()

	if !assert.Equal(t, 1, len(parser.Errors())) {
		t.Fail()
	}
}

func TestReturnStatement(t *testing.T) {
	assert := assert.New(t)
	source := `
		regresa 5;
		regresa foo;
	`

	lexer := lpp.NewLexer(source)
	parser := lpp.NewParser(lexer)
	program := parser.ParseProgam()

	if !assert.Equal(2, len(program.Staments)) {
		t.Log("len of program statements are not 2")
		t.Fail()
	}

	for _, statement := range program.Staments {
		assert.Equal("regresa", statement.TokenLiteral())
		assert.IsType(&lpp.ReturnStament{}, statement.(*lpp.ReturnStament))
	}
}

func TestIdentifierExpression(t *testing.T) {
	source := "foobar;"
	lexer := lpp.NewLexer(source)
	parser := lpp.NewParser(lexer)
	program := parser.ParseProgam()

	testProgramStatements(t, parser, &program, 1)

	expressionStament := program.Staments[0].(*lpp.ExpressionStament)
	if !assert.NotNil(t, expressionStament.Expression) {
		t.Fail()
	}

	testLiteralExpression(t, expressionStament.Expression, "foobar")
}

func testProgramStatements(t *testing.T, parser *lpp.Parser, program *lpp.Program, expectedStamenetCount int) {
	assert := assert.New(t)
	assert.Equal(0, len(parser.Errors()))
	assert.Equal(expectedStamenetCount, len(program.Staments))
	assert.IsType(&lpp.ExpressionStament{}, program.Staments[0].(*lpp.ExpressionStament))
}

func testLiteralExpression(t *testing.T, expression lpp.Expression, expectedValue interface{}) {
	switch expectedValue := expectedValue.(type) {
	case string:
		testIdentifier(t, expression, expectedValue)
	default:
		t.Log(fmt.Sprintf("unhandled type of expression, Got=%s", reflect.TypeOf(expectedValue).String()))
		t.Fail()
	}
}

func testIdentifier(t *testing.T, expression lpp.Expression, expectedValue string) {
	assert := assert.New((t))
	assert.IsType(&lpp.Identifier{}, expression.(*lpp.Identifier))

	identifier := expression.(*lpp.Identifier)
	assert.Equal(expectedValue, identifier.Str())
	assert.Equal(expectedValue, identifier.TokenLiteral())
}
