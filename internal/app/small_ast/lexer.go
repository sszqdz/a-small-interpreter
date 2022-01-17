package small_ast

import (
	"strings"
)

type Lexer interface {
	GetNextToken() *Token
}

func NewLexer(text string) Lexer {
	lex := &lexer{
		text: text,
		pos:  0,
		len:  len(text),
	}
	if lex.len > 0 {
		lex.currentChar = string(text[0])
	}

	return lex
}

type lexer struct {
	text        string
	pos         int
	len         int
	currentChar string
}

func (lex *lexer) GetNextToken() *Token {
	for lex.currentChar != EOF {
		switch lex.currentChar {
		case SPACE:
			lex.skipSpace()
			continue
		case LPAREN:
			lex.advance()
			return NewToken(TOKEN_KIND_LPAREN, LPAREN)
		case RPAREN:
			lex.advance()
			return NewToken(TOKEN_KIND_RPAREN, RPAREN)
		case OR:
			lex.advance()
			return NewToken(TOKEN_KIND_OR, OR)
		case COMMA:
			lex.advance()
			return NewToken(TOKEN_KIND_COMMA, COMMA)
		default:
			str := lex.extractStr()
			if len(str) > 0 {
				return NewToken(TOKEN_KIND_STRING, str)
			} else {
				return NewToken(TOKEN_KIND_EOF, EOF)
			}
		}
	}

	return NewToken(TOKEN_KIND_EOF, EOF)
}

func (lex *lexer) advance() {
	lex.pos += 1
	if lex.pos > lex.len-1 {
		lex.currentChar = EOF
	} else {
		lex.currentChar = string(lex.text[lex.pos])
	}
}

func (lex *lexer) skipSpace() {
	for lex.currentChar == SPACE {
		lex.advance()
	}
}

func (lex *lexer) extractStr() string {
	var sb strings.Builder
Loop:
	for lex.currentChar != EOF {
		switch {
		case lex.currentChar == SPACE:
			lex.skipSpace()
			continue
		case lex.isOperator():
			break Loop
		default:
			sb.WriteString(lex.currentChar)
			lex.advance()
		}
	}

	return sb.String()
}

func (lex *lexer) isOperator() bool {
	switch lex.currentChar {
	case LPAREN:
	case RPAREN:
	case OR:
	case COMMA:
	default:
		return false
	}

	return true
}
