package quark

type (
	//A branch using if, elif, and else
	IfBranch struct {
		IfPart CondBranchPart
		ElifParts []CondBranchPart
		ElsePart *BranchPart

		kwIf ObjectPosition
		closeCurly ObjectPosition
	}

	//The body of a branch. BodyReturn is an optional expression which is used
	//as the value of the branch in an expression context.
	BranchPart struct {
		Body []Stmt
		BodyReturn Expr
	}

	//A conditional branch part. Represents branch parts which contain a condition
	//expression like if and elif.
	CondBranchPart struct {
		Cond Expr
		BranchPart
	}

	//A match statement or expression.
	MatchBranch struct {
		MatchExpr Expr
		Cases []CaseBranchPart

		kwMatch ObjectPosition
		closeCurly ObjectPosition
	}

	//A case branch part. Used in match branches.
	CaseBranchPart struct {
		CasePattern Pattern
		BranchPart
	}
)

func NewIfBranch(ifPart CondBranchPart, elifPart []CondBranchPart, elsePart *BranchPart, kwIf ObjectPosition, closeCurly ObjectPosition) *IfBranch {
	return &IfBranch{
		IfPart:     ifPart,
		ElifParts:  elifPart,
		ElsePart:   elsePart,
		kwIf:       kwIf,
		closeCurly: closeCurly,
	}
}

func NewMatchBranch(expr Expr, cases []CaseBranchPart, kwMatch ObjectPosition, closeCurly ObjectPosition) *MatchBranch {
	return &MatchBranch{
		MatchExpr:  expr,
		Cases:      cases,
		kwMatch:    kwMatch,
		closeCurly: closeCurly,
	}
}


func (b *IfBranch) Start() *ObjectPosition {
	return &b.kwIf
}

func (b *MatchBranch) Start() *ObjectPosition {
	return &b.kwMatch
}

func (b *IfBranch) End() *ObjectPosition {
	return &b.closeCurly
}

func (b *MatchBranch) End() *ObjectPosition {
	return &b.closeCurly
}

func (b *IfBranch) branchNode() {}
func (b *MatchBranch) branchNode() {}
