package quark


//A unary operator.
type UnaryOp interface {

	unaryOp()
}

//A binary operator.
type BinaryOp interface {

	binaryOp()
}

//An assignment operator
type AssignmentOp interface {
	assignmentOp()
}