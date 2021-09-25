package lexer_test

import (
	"testing"

	"github.com/farbodahm/lets-go/monkeyLangInterpreter/lexer"
	"github.com/farbodahm/lets-go/monkeyLangInterpreter/token"
)

func TestNextTokenDelimiters(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := lexer.NewLexer(input)

	for i, test := range tests {
		nextToken := lexer.NextToken()

		if nextToken.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, test.expectedType, nextToken.Type)
		}

		if nextToken.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedLiteral, nextToken.Literal)
		}
	}
}
