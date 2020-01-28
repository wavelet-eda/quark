package quark

//A unary operator
type UnaryOp = uint16

const (
	OpComplement UnaryOp = iota
	OpNot
)

//A binary operator.
type BinaryOp = uint16

const (
	OpMul BinaryOp = iota
	OpDiv
	OpMod
	OpSub
	OpAdd
	OpLeftShift
	OpRightShift
	OpArithLeftShift
	OpArithRightShift
	OpLT
	OpGT
	OpLTE
	OpGTE
	OpBinaryAnd
	OpBinaryOr
	OpBinaryXor
	OpBinaryNand
	OpBinaryNor
	OpBinaryXnor
	OpLogicalAnd
	OpLogicalOr
	OpLogicalNor
	OpImplication
	OpEquivalence
)

//An assignment operator
type AssignmentOp = uint16

const (
	OpAssign = iota
	OpAddAssign
	OpSubAssign
	OpMulAssign
	OpDivAssign
	OpModAssign
	OpBandAssign
	OpBorAssign
	OpXorAssign
	OpLeftShiftAssign
	OpRightShiftAssign
	OpArithLeftShiftAssign
	OpArithRightShiftAssign
)