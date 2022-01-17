package small_ast

import stderr "errors"

// expr   : term (COMMA term)*
// term   : factor (OR factor)*
// factor : STRING | LPAREN expr RPAREN
// 注：COMMA与AND含义相同，但优先级不同

type Parser interface {
	Parse() (Node, error)
}

func NewParser(lex Lexer) Parser {
	par := &parser{lex: lex}
	par.currentToken = par.lex.GetNextToken()

	return par
}

type parser struct {
	lex          Lexer
	currentToken *Token
}

func (par *parser) Parse() (Node, error) {
	return par.expr()
}

func (par *parser) expr() (Node, error) {
	node, err := par.term()
	if err != nil {
		return nil, err
	}

	for par.currentToken.kind == TOKEN_KIND_COMMA {
		tok := par.currentToken
		err = par.consume(TOKEN_KIND_COMMA)
		if err != nil {
			return nil, err
		}

		left := node
		right, err := par.term()
		if err != nil {
			return nil, err
		}

		node = NewBinaryOperatorNode(tok, left, right)
	}

	return node, nil
}

func (par *parser) term() (Node, error) {
	node, err := par.factor()
	if err != nil {
		return nil, err
	}

	for par.currentToken.kind == TOKEN_KIND_OR {
		tok := par.currentToken
		err = par.consume(TOKEN_KIND_OR)
		if err != nil {
			return nil, err
		}

		left := node
		right, err := par.factor()
		if err != nil {
			return nil, err
		}

		node = NewBinaryOperatorNode(tok, left, right)
	}

	return node, nil
}

func (par *parser) factor() (Node, error) {
	tok := par.currentToken

	switch tok.kind {
	case TOKEN_KIND_STRING:
		err := par.consume(TOKEN_KIND_STRING)
		if err != nil {
			return nil, err
		}

		return NewStrNode(tok), nil

	case TOKEN_KIND_LPAREN:
		err := par.consume(TOKEN_KIND_LPAREN)
		if err != nil {
			return nil, err
		}
		node, err := par.expr()
		if err != nil {
			return nil, err
		}
		err = par.consume(TOKEN_KIND_RPAREN)
		if err != nil {
			return nil, err
		}

		return node, nil

	default:
		return nil, stderr.New("invalid factor token kind")
	}
}

func (par *parser) consume(tokenKind int) error {
	if par.currentToken.kind == tokenKind {
		par.currentToken = par.lex.GetNextToken()
	} else {
		return stderr.New("invalid consume token kind")
	}

	return nil
}

//func inSlice(slice []int, obj int) bool {
//	for _, item := range slice {
//		if item == obj {
//			return true
//		}
//	}
//
//	return false
//}
