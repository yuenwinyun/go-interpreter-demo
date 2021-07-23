package eval

import (
	"go-interpreter-demo/ast"
	"go-interpreter-demo/datastructure"
)

func Eval(node ast.Node) datastructure.DataStructure {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &datastructure.Integer{Value: node.Value}
	}
	return nil
}

func evalStatements(stmts []ast.Statement) datastructure.DataStructure {
	var result datastructure.DataStructure

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
