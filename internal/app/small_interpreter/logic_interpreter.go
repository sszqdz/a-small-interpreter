package small_interpreter

import (
	small_ast2 "a-small-interpreter/internal/app/small_ast"
	stderr "errors"
)

// TODO go缺少泛型/方法重载，因此定义访问者模式中公共Visit方法意义不大，各解析器自行实现即可
//  解耦Parser与Interpreter，双方互不引用，仅将 AST Root Node 作为联系依据

type LogicInterpreter struct {
}

func NewLogicInterpreter() *LogicInterpreter {
	return &LogicInterpreter{}
}

func (intpret *LogicInterpreter) Interpret(rootNode small_ast2.Node, condition map[string]bool) ([]string, error) {
	result, err := intpret.visit(rootNode, condition)
	if err != nil {
		return nil, err
	}
	return intpret.removeDuplicates(result), nil
}

func (intpret *LogicInterpreter) visit(node small_ast2.Node, condition map[string]bool) ([]string, error) {
	switch node.(type) {
	case *small_ast2.BinaryOperatorNode:
		return intpret.visitBinaryOperatorNode(node.(*small_ast2.BinaryOperatorNode), condition)
	case *small_ast2.StrNode:
		return intpret.visitStrNode(node.(*small_ast2.StrNode), condition)
	default:
		return nil, stderr.New("[visit] invalid node type")
	}
}

func (intpret *LogicInterpreter) visitBinaryOperatorNode(node *small_ast2.BinaryOperatorNode, condition map[string]bool) ([]string, error) {
	switch node.Kind {
	case small_ast2.TOKEN_KIND_COMMA:
		return intpret.handleCommaNode(node, condition)
	case small_ast2.TOKEN_KIND_OR:
		return intpret.handleOrNode(node, condition)
	default:
		return nil, stderr.New("[visitBinaryOperatorNode] invalid node kind")
	}
}

func (intpret *LogicInterpreter) handleCommaNode(node *small_ast2.BinaryOperatorNode, condition map[string]bool) ([]string, error) {
	left, err := intpret.visit(node.Left, condition)
	if err != nil {
		return nil, err
	}
	right, err := intpret.visit(node.Right, condition)
	if err != nil {
		return nil, err
	}
	merge := append(left, right...)

	return merge, nil
}

func (intpret *LogicInterpreter) handleOrNode(node *small_ast2.BinaryOperatorNode, condition map[string]bool) ([]string, error) {
	left, err := intpret.visit(node.Left, condition)
	if err != nil {
		return nil, err
	}
	if len(left) > 0 {
		return left, nil
	}

	right, err := intpret.visit(node.Right, condition)
	if err != nil {
		return nil, err
	}
	if len(right) > 0 {
		return right, nil
	}

	return []string{}, nil
}

func (intpret *LogicInterpreter) visitStrNode(node *small_ast2.StrNode, condition map[string]bool) ([]string, error) {
	sli := make([]string, 0, 1)

	if condition[node.Value] {
		sli = append(sli, node.Value)
	}

	return sli, nil
}

func (intpret *LogicInterpreter) removeDuplicates(sli []string) []string {
	size := len(sli)
	newSli := make([]string, 0, size)
	if size == 0 {
		return newSli
	}

	dic := make(map[string]bool, size)
	for _, s := range sli {
		if dic[s] {
			continue
		}
		dic[s] = true
		newSli = append(newSli, s)
	}

	return newSli
}
