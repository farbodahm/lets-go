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

	lexer.skipWhiteSpaces()

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
	default:
		if isLetter(lexer.ch) {
			nextToken.Literal = lexer.readIdentifier()
			nextToken.Type = token.LookupIdent(nextToken.Literal)
			return nextToken
		} else if isDigit(lexer.ch) {
			nextToken.Literal = lexer.readNumber()
			nextToken.Type = token.INT
			return nextToken
		} else {
			nextToken = newToken(token.ILLEGAL, lexer.ch)
		}
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

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) skipWhiteSpaces() {
	for lexer.ch == ' ' || lexer.ch == '\n' || lexer.ch == '\t' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}
