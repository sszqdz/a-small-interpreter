package test

import (
	"a-small-interpreter/internal/app/small_ast"
	"a-small-interpreter/internal/app/small_interpreter"
	"testing"
)

func TestInterpreter(t *testing.T) {
	intpret := small_interpreter.NewLogicInterpreter()
	for _, param := range multiParams() {
		t.Logf("%-10s %s", "Rule:", param.rule)
		lex := small_ast.NewLexer(param.rule)
		par := small_ast.NewParser(lex)
		rootNode, _ := par.Parse()

		for _, value := range param.vals {
			out, _ := intpret.Interpret(rootNode, value.inDic)
			t.Logf("%-10s %v", "In:", value.inDic)
			t.Logf("%-10s %v", "Expect:", value.expect)
			t.Logf("%-10s %v", "Output:", out)

			if compareSlice(out, value.expect) {
				t.Logf("%s", "Wow!")
			} else {
				t.Error("Output inconsistent with expectation...")
				t.Fatal("Please check your code carefully")
			}
		}
		t.Log("\n\n")
	}
}

const (
	NAME   = "name"
	AGE    = "age"
	GENDER = "gender"
)

type paraObj struct {
	rule string
	vals []*valObj
}

type valObj struct {
	inDic  map[string]bool
	expect []string
}

func compareSlice(s1 []string, s2 []string) bool {
	size1 := len(s1)
	size2 := len(s2)

	if size1 != size2 {
		return false
	}

	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func multiParams() []*paraObj {
	return []*paraObj{
		{
			"name",
			[]*valObj{
				{
					map[string]bool{NAME: true},
					[]string{NAME},
				},
				{
					map[string]bool{NAME: false},
					[]string{},
				},
			},
		},
		{
			"name|age",
			[]*valObj{
				{
					map[string]bool{NAME: true, AGE: true},
					[]string{NAME},
				},
				{
					map[string]bool{NAME: true, AGE: false},
					[]string{NAME},
				},
				{
					map[string]bool{NAME: false, AGE: true},
					[]string{AGE},
				},
				{
					map[string]bool{NAME: false, AGE: false},
					[]string{},
				},
			},
		},
		{
			"name,age",
			[]*valObj{
				{
					map[string]bool{NAME: true, AGE: true},
					[]string{NAME, AGE},
				},
				{
					map[string]bool{NAME: true, AGE: false},
					[]string{NAME},
				},
				{
					map[string]bool{NAME: false, AGE: true},
					[]string{AGE},
				},
				{
					map[string]bool{NAME: false, AGE: false},
					[]string{},
				},
			},
		},
		{
			"name|age,gender",
			[]*valObj{
				{
					map[string]bool{NAME: true, AGE: true, GENDER: true},
					[]string{NAME, GENDER},
				},
				{
					map[string]bool{NAME: true, AGE: true, GENDER: false},
					[]string{NAME},
				},
				{
					map[string]bool{NAME: true, AGE: false, GENDER: true},
					[]string{NAME, GENDER},
				},
				{
					map[string]bool{NAME: true, AGE: false, GENDER: false},
					[]string{NAME},
				},
				{
					map[string]bool{NAME: false, AGE: true, GENDER: true},
					[]string{AGE, GENDER},
				},
				{
					map[string]bool{NAME: false, AGE: true, GENDER: false},
					[]string{AGE},
				},
				{
					map[string]bool{NAME: false, AGE: false, GENDER: true},
					[]string{GENDER},
				},
				{
					map[string]bool{NAME: false, AGE: false, GENDER: false},
					[]string{},
				},
			},
		},
		{
			"name|(age,gender)",
			[]*valObj{
				{
					map[string]bool{NAME: true, AGE: true, GENDER: true},
					[]string{NAME},
				},
				{
					map[string]bool{NAME: true, AGE: true, GENDER: false},
					[]string{NAME},
				},
				{
					map[string]bool{NAME: true, AGE: false, GENDER: true},
					[]string{NAME},
				},
				{
					map[string]bool{NAME: true, AGE: false, GENDER: false},
					[]string{NAME},
				},
				{
					map[string]bool{NAME: false, AGE: true, GENDER: true},
					[]string{AGE, GENDER},
				},
				{
					map[string]bool{NAME: false, AGE: true, GENDER: false},
					[]string{AGE},
				},
				{
					map[string]bool{NAME: false, AGE: false, GENDER: true},
					[]string{GENDER},
				},
				{
					map[string]bool{NAME: false, AGE: false, GENDER: false},
					[]string{},
				},
			},
		},
	}
}
