package lexer

import "github.com/farbodahm/lets-go/monkeyLangInterpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar() // Initialize lexer
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var nextToken token.Token

	switch lexer.ch {
	case '=':
		nextToken = newToken(token.ASSIGN, lexer.ch)
	case ';':
		nextToken = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		nextToken = newToken(token.LPAREN, lexer.ch)
	case ')':
		nextToken = newToken(token.RPAREN, lexer.ch)
	case ',':
		nextToken = newToken(token.COMMA, lexer.ch)
	case '+':
		nextToken = newToken(token.PLUS, lexer.ch)
	case '{':
		nextToken = newToken(token.LBRACE, lexer.ch)
	case '}':
		nextToken = newToken(token.RBRACE, lexer.ch)
	case 0:
		nextToken.Type = token.EOF
		nextToken.Literal = ""
	}

	lexer.readChar()
	return nextToken
}

// Helper functions:

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
