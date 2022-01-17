package small_ast

type Node interface {
	ExtractKind() int
}

type BinaryOperatorNode struct {
	token *Token

	Kind  int
	Value string
	Left  Node
	Right Node
}

func NewBinaryOperatorNode(token *Token, left Node, right Node) *BinaryOperatorNode {
	node := &BinaryOperatorNode{
		token: token,
		Left:  left,
		Right: right,
	}
	if token != nil {
		node.Kind = token.kind
		node.Value = token.value
	}

	return node
}

func (node *BinaryOperatorNode) ExtractKind() int {
	return node.Kind
}

type StrNode struct {
	token *Token

	Kind  int
	Value string
}

func NewStrNode(token *Token) *StrNode {
	node := &StrNode{token: token}
	if token != nil {
		node.Kind = token.kind
		node.Value = token.value
	}

	return node
}

func (node *StrNode) ExtractKind() int {
	return node.Kind
}
