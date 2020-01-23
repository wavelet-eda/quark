package quark

import (
	"testing"
)

type TestASTVisitor struct {
	count int
}

func (v *TestASTVisitor) Visit(node AST) Visitor {
	v.count += 1
	return v
}

func TestVisitorVisitCount(t *testing.T) {
	var tests = []struct{
		name string
		in AST
		out int
	}{
		{
			"variable-decl",
			&AssignStmt{
				AssignTo: &VariableDefinitionAssignable{
					IsMut:   false,
					VarType: &CompleteType{X:&RealName{Text:"Bit"}},
					VarName: &RealName{Text:"x"},
					kwMut:   ObjectPosition{},
				},
				AssignmentType: OpAssign,
				TheExpr:        &LiteralExpr{Value:&Literal{Text:"0"}},
				semi:           ObjectPosition{},
			},
			7,
		},
	}

	for _, test := range tests {
		name := test.name
		t.Run(name, func(t *testing.T) {
			visitor := &TestASTVisitor{count: 0}
			test.in.Accept(visitor)

			if test.out != visitor.count {
				t.Errorf("Expected %d nodes but visited %d", test.out, visitor.count)
			}
		})
	}
}