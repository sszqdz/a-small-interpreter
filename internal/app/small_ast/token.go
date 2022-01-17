package small_ast

// 运算符优先级
// ()
// |
// ,
const (
	LPAREN = "("
	RPAREN = ")"
	OR     = "|"
	COMMA  = ","

	EOF = ""

	SPACE = " "
)

const (
	TOKEN_KIND_LPAREN = iota
	TOKEN_KIND_RPAREN
	TOKEN_KIND_OR
	TOKEN_KIND_COMMA

	TOKEN_KIND_EOF

	TOKEN_KIND_STRING
)

type Token struct {
	kind  int
	value string
}

func NewToken(kind int, value string) *Token {
	return &Token{kind, value}
}
