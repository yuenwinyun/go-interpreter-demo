package eval

import (
	"go-interpreter-demo/datastructure"
	"go-interpreter-demo/parser"
	"testing"
)

func testEval(input string) datastructure.DataStructure {
	p := parser.New(input)
	program := p.ParseProgram()
	return Eval(program)
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testIntegerObject(t *testing.T, obj datastructure.DataStructure, expected int64) bool {
	result, ok := obj.(*datastructure.Integer)
	if !ok {
		t.Errorf("datastructure is not Integer")
		return false
	}
	if result.Value != expected {
		t.Errorf("datastructure has wrong value")
		return false
	}
	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj datastructure.DataStructure, expected bool) bool {
	result, ok := obj.(*datastructure.Boolean)
	if !ok {
		t.Errorf("datastructure is not Boolean")
		return false
	}
	if result.Value != expected {
		t.Errorf("datastructure has wrong value")
		return false
	}
	return true
}
