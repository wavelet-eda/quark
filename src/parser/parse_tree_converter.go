//Home of the parser and lexer implementations.
package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wavelet-eda/quark/src/quark"
)

//Visitor implementation which converts parse tree from Antlr into
//proper quark.AST. This class is hellish. Don't change it if you don't
//know what you're doing. The private methods are mostly added to
//consolidate type casting.
type ParseTreeConverter struct {
	BaseQuarkParserVisitor
	file *quark.QuarkFile
}

func NewParseTreeConverter(file *quark.QuarkFile) *ParseTreeConverter {
	return &ParseTreeConverter{file: file}
}

func (ptc *ParseTreeConverter) terminalPosition(node antlr.TerminalNode) quark.ObjectPosition {
	return ptc.tokenPosition(node.GetSymbol())
}

func (ptc *ParseTreeConverter) tokenPosition(token antlr.Token) quark.ObjectPosition {
	return quark.NewObjectPosition(ptc.file, token.GetTokenIndex())
}

func (ptc *ParseTreeConverter) visitExpr(ctx IExprContext) quark.Expr {
	return ptc.Visit(ctx).(quark.Expr)
}

func (ptc *ParseTreeConverter) visitTypeExpr(ctx ITypeexprContext) quark.TypeExpr {
	return ptc.Visit(ctx).(quark.TypeExpr)
}

func (ptc *ParseTreeConverter) visitParamArgList(ctx IParamarglistContext) []*quark.ParamArgument {
	realCtx := ctx.(*ParamarglistContext)
	result := make([]*quark.ParamArgument, len(realCtx.AllParamarg()))
	for i, node := range realCtx.AllParamarg() {
		result[i] = ptc.VisitParamarg(node.(*ParamargContext)).(*quark.ParamArgument)
	}
	return result
}

func (ptc *ParseTreeConverter) visitAssignment(ctx IAssignmentContext) quark.AssignmentOp {
	switch ctx.GetStart().GetTokenType() {
	case QuarkLexerOP_ASSIGN: 			  		return quark.OpAssign
	case QuarkLexerOP_ADD_ASSIGN: 		  		return quark.OpAddAssign
	case QuarkLexerOP_SUB_ASSIGN: 		  		return quark.OpSubAssign
	case QuarkLexerOP_MUL_ASSIGN: 		  		return quark.OpMulAssign
	case QuarkLexerOP_DIV_ASSIGN: 		  		return quark.OpDivAssign
	case QuarkLexerOP_BAND_ASSIGN: 		  		return quark.OpBandAssign
	case QuarkLexerOP_BOR_ASSIGN:   	  		return quark.OpBorAssign
	case QuarkLexerOP_XOR_ASSIGN: 		  		return quark.OpBorAssign
	case QuarkLexerOP_LEFT_SHIFT_ASSIGN:  		return quark.OpLeftShiftAssign
	case QuarkLexerOP_RIGHT_SHIFT_ASSIGN: 		return quark.OpRightShiftAssign
	case QuarkLexerOP_ARITH_LEFT_SHIFT_ASSIGN: 	return quark.OpArithLeftShiftAssign
	case QuarkLexerOP_ARITH_RIGHT_SHIFT_ASSIGN: return quark.OpArithRightShiftAssign
	}
	panic("bad assignment operator")
}

func (ptc *ParseTreeConverter) visitBinop(op antlr.Token) quark.BinaryOp {
	switch op.GetTokenType() {
	case QuarkLexerOP_MUL:					return quark.OpMul
	case QuarkLexerOP_DIV:					return quark.OpDiv
	case QuarkLexerOP_MOD:					return quark.OpMod
	case QuarkLexerOP_SUB:					return quark.OpSub
	case QuarkLexerOP_ADD:					return quark.OpAdd
	case QuarkLexerOP_LEFT_SHIFT:			return quark.OpLeftShift
	case QuarkLexerOP_RIGHT_SHIFT: 			return quark.OpRightShift
	case QuarkLexerOP_ARITH_LEFT_SHIFT: 	return quark.OpArithLeftShift
	case QuarkLexerOP_ARITH_RIGHT_SHFIT:	return quark.OpArithRightShift
	case QuarkLexerOP_BAND:					return quark.OpBinaryAnd
	case QuarkLexerOP_BOR:					return quark.OpBinaryOr
	case QuarkLexerOP_XOR:					return quark.OpBinaryXor
	case QuarkLexerOP_BNAND:				return quark.OpBinaryNand
	case QuarkLexerOP_BNOR:					return quark.OpBinaryNor
	case QuarkLexerOP_XNOR:					return quark.OpBinaryXnor
	case QuarkLexerOP_LAND:					return quark.OpLogicalAnd
	case QuarkLexerOP_LOR:					return quark.OpLogicalOr
	case QuarkLexerOP_IMPLICATION:			return quark.OpImplication
	case QuarkLexerOP_EQUIVALENCE:			return quark.OpEquivalence
	}
	panic("bad binary operator")
}

func (ptc *ParseTreeConverter) visitClockExpr(ctx IClockexprContext) quark.ClockExpr {
	return ptc.Visit(ctx).(quark.ClockExpr)
}

func (ptc *ParseTreeConverter) visitAssignable(ctx IAssignableContext) quark.Assignable {
	return ptc.Visit(ctx).(quark.Assignable)
}

func (ptc *ParseTreeConverter) visitBranch(ctx IBranchContext) quark.Branch {
	return ptc.Visit(ctx).(quark.Branch)
}

func (ptc *ParseTreeConverter) visitRealname(ctx IRealnameContext) *quark.RealName {
	return quark.NewRealName(ctx.GetText(), ptc.tokenPosition(ctx.GetStart()))
}

func (ptc *ParseTreeConverter) visitName(ctx INameContext) quark.Name {
	return ptc.Visit(ctx).(quark.Name)
}

func (ptc *ParseTreeConverter) visitArgumentList(ctx IArgumentlistContext) []*quark.ArgumentDef {
	return ptc.Visit(ctx).([]*quark.ArgumentDef)
}

func (ptc *ParseTreeConverter) visitParameterList(ctx IParameterlistContext) []*quark.ParameterDef {
	return ptc.Visit(ctx).([]*quark.ParameterDef)
}

func (ptc *ParseTreeConverter) visitParameterDef(ctx IParameterdefContext) *quark.ParameterDef {
	return ptc.Visit(ctx).(*quark.ParameterDef)
}

func (ptc *ParseTreeConverter) visitReturnList(ctx IReturnlistContext) quark.ReturnList {
	return ptc.Visit(ctx).(quark.ReturnList)
}

func (ptc *ParseTreeConverter) visitCallArgList(ctx ICallarglistContext) []*quark.CallArgument {
	actualList := ctx.(*CallarglistContext) //this is the only option
	callArgs := make([]*quark.CallArgument, len(actualList.AllCallarg()))

	for i, node := range actualList.AllCallarg() {
		callArg := ptc.Visit(node).(*quark.CallArgument)
		callArgs[i] = callArg
	}
	return callArgs
}

func (ptc *ParseTreeConverter) visitBlock(ctx IBlockContext) []quark.Stmt {
	return ptc.Visit(ctx).([]quark.Stmt)
}

func (ptc *ParseTreeConverter) visitInnerConcat(ctx IInnerconcatContext) quark.InnerConcat {
	return ptc.Visit(ctx).(quark.InnerConcat)
}

func (ptc *ParseTreeConverter) VisitQuarkpackage(ctx *QuarkpackageContext) interface{} {
	var rawImportDecls = ctx.AllImportdecl()
	var rawDecls = ctx.AllDecl()

	var importDecls = make([]quark.ImportDecl, len(rawImportDecls))
	var decls = make([]quark.Decl, len(rawDecls))

	for index, importDecl := range rawImportDecls {
		importDecls[index] = ptc.visitImportDecl(importDecl)
	}

	for index, decl := range rawDecls {
		decls[index] = ptc.visitDecl(decl)
	}

	return &quark.Package{
		Imports: importDecls,
		Symbols: decls,
	}
}

func (ptc *ParseTreeConverter) visitPattern(pattern IPatternContext) quark.Pattern {
	return ptc.Visit(pattern).(quark.Pattern)
}

func (ptc *ParseTreeConverter) visitImportDecl(rawImportDecl IImportdeclContext) quark.ImportDecl {
	return ptc.Visit(rawImportDecl).(quark.ImportDecl)
}

func (ptc *ParseTreeConverter) visitDecl(rawDecl IDeclContext) quark.Decl {
	return ptc.Visit(rawDecl.GetChild(0).(antlr.ParseTree)).(quark.Decl)
}

func (ptc *ParseTreeConverter) VisitSingleImport(ctx *SingleImportContext) interface{} {
	println("visiting single import")
	quarkName := ptc.Visit(ctx.Name()).(quark.Name)
	kwImport := ptc.terminalPosition(ctx.KW_IMPORT())
	return &quark.SingleImport{
		GenericImport: quark.GenericImport{
			PackageName:quarkName,
			KwImport: kwImport,
		},
	}
}

func (ptc *ParseTreeConverter) VisitWildcardImport(ctx *WildcardImportContext) interface{} {
	quarkName := ptc.Visit(ctx.Name()).(quark.Name)
	kwImport := ptc.terminalPosition(ctx.KW_IMPORT())
	return &quark.WildcardImport{
		GenericImport: quark.GenericImport{
			PackageName:quarkName,
			KwImport: kwImport,
		},
	}
}

func (ptc *ParseTreeConverter) VisitMultiImport(ctx *MultiImportContext) interface{} {
	packageName := ptc.visitName(ctx.Name())
	kwImport := ptc.terminalPosition(ctx.KW_IMPORT())
	symbols := make([]*quark.RealName, len(ctx.AllRealname()))
	for i, node := range ctx.AllRealname() {
		symbols[i] = ptc.visitRealname(node)
	}
	return &quark.MultiImport{
		GenericImport: quark.GenericImport{
			PackageName:packageName,
			KwImport: kwImport,
		},
		Symbols:       nil,
		CloseParen:    quark.ObjectPosition{},
	}
}

func (ptc *ParseTreeConverter) VisitBlock(ctx *BlockContext) interface{} {
	parsedStmts := ctx.AllStmt()
	quarkStmts := make([]quark.Stmt, len(parsedStmts))
	for index, node := range parsedStmts {
		stmt := ptc.Visit(node).(quark.Stmt)
		quarkStmts[index] = stmt
	}

	return quarkStmts
}

func (ptc *ParseTreeConverter) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	assignable := ptc.visitAssignable(ctx.Assignable())
	assignment := ptc.visitAssignment(ctx.Assignment())
	expr := ptc.visitExpr(ctx.Expr())
	return quark.NewAssignStmt(assignable, assignment, expr, ptc.terminalPosition(ctx.SEMI()))
}

func (ptc *ParseTreeConverter) VisitRegAssignStmt(ctx *RegAssignStmtContext) interface{} {
	clock := ptc.visitClockExpr(ctx.GetClk())
	var reset quark.ClockExpr
	if ctx.GetRst() != nil {
		reset = ptc.visitClockExpr(ctx.GetClk())
	} else {
		reset = nil
	}
	assignable := ptc.visitAssignable(ctx.Assignable())
	assignment := ptc.visitAssignment(ctx.Assignment())
	expr := ptc.visitExpr(ctx.Expr())
	semi := ptc.terminalPosition(ctx.SEMI())
	reg := ptc.terminalPosition(ctx.KW_REG())

	properAssign := quark.NewAssignStmt(assignable, assignment, expr, semi)
	return quark.NewRegAssignStmt(clock, reset, properAssign, reg)
}

func (ptc *ParseTreeConverter) VisitDeclarationStmt(ctx *DeclarationStmtContext) interface{} {
	assignable := ptc.visitAssignable(ctx.Assignable())
	semi := ptc.terminalPosition(ctx.SEMI())
	return quark.NewDeclStmt(assignable, semi)
}

func (ptc *ParseTreeConverter) VisitBranchStmt(ctx *BranchStmtContext) interface{} {
	return &quark.BranchStmt{TheBranch: ptc.visitBranch(ctx.Branch())}
}

func (ptc *ParseTreeConverter) VisitFutureStmt(ctx *FutureStmtContext) interface{} {
	kwFuture := ptc.terminalPosition(ctx.KW_FUTURE())
	semi := ptc.terminalPosition(ctx.SEMI())
	typeExpr := ptc.visitTypeExpr(ctx.Typeexpr())
	name := ptc.visitRealname(ctx.Realname())

	return &quark.FutureStmt{
		FutureType: typeExpr,
		FutureName: name,
		KwFuture:   kwFuture,
		Semi:       semi,
	}
}

func (ptc *ParseTreeConverter) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	expr := ptc.visitExpr(ctx.Expr())
	returnPos := ptc.terminalPosition(ctx.KW_RETURN())
	semiPos := ptc.terminalPosition(ctx.SEMI())
	return quark.NewReturnStmt(expr, returnPos, semiPos)
}

func (ptc *ParseTreeConverter) VisitValueAssignment(ctx *ValueAssignmentContext) interface{} {
	return &quark.ValueAssignment{Variable: ptc.visitName(ctx.Name())}
}

func (ptc *ParseTreeConverter) VisitVariableDefinition(ctx *VariableDefinitionContext) interface{} {
	varType := ptc.visitTypeExpr(ctx.Typeexpr())
	name := ptc.visitRealname(ctx.Realname())
	var mutPos quark.ObjectPosition
	isMut := false
	if ctx.KW_MUT() != nil {
		mutPos = ptc.terminalPosition(ctx.KW_MUT())
		isMut = true
	}
	return quark.NewVariableDefAssignment(isMut, varType, name, mutPos)
}

func (ptc *ParseTreeConverter) VisitTupleDestructer(ctx *TupleDestructerContext) interface{} {
	assignables := make([]quark.Assignable, len(ctx.AllAssignable()))
	for index, node := range ctx.AllAssignable() {
		assignables[index] = ptc.visitAssignable(node)
	}
	return &quark.TupleDestructionAssignment{Assignables:assignables}
}

func (ptc *ParseTreeConverter) VisitArrayIndexAssignment(ctx *ArrayIndexAssignmentContext) interface{} {
	allExprs := make([]quark.Expr, len(ctx.AllExpr()))
	for index, node := range ctx.AllExpr() {
		allExprs[index] = ptc.visitExpr(node)
	}

	arrayExpr := allExprs[0]
	indices := allExprs[1:]
	closeBracePos := ptc.terminalPosition(ctx.RBRACE())
	return quark.NewArrayIndex(arrayExpr, indices, closeBracePos)
}

func (ptc *ParseTreeConverter) VisitArraySliceAssignment(ctx *ArraySliceAssignmentContext) interface{} {
	arrayExpr := ptc.visitExpr(ctx.Expr(0))
	var msb quark.Expr = nil
	var lsb quark.Expr = nil
	var step quark.Expr = nil
	rBracePos := ptc.terminalPosition(ctx.RBRACE())

	if ctx.msb != nil {
		msb = ptc.visitExpr(ctx.msb)
	}
	if ctx.lsb != nil {
		lsb = ptc.visitExpr(ctx.lsb)
	}
	if ctx.step != nil {
		step = ptc.visitExpr(ctx.step)
	}

	return quark.NewArraySlice(arrayExpr, msb, lsb, step, rBracePos)
}

func (ptc *ParseTreeConverter) VisitRealname(ctx *RealnameContext) interface{} { //quark.Name
	pos := quark.NewObjectPosition(ptc.file, ctx.REAL_NAME().GetSymbol().GetTokenIndex())
	return quark.NewRealName(ctx.REAL_NAME().GetText(), pos)
}

func (ptc *ParseTreeConverter) VisitUnqualifiedName(ctx *UnqualifiedNameContext) interface{} {
	return ptc.visitRealname(ctx.Realname())
}

func (ptc *ParseTreeConverter) VisitQualifiedName(ctx *QualifiedNameContext) interface{} { //quark.QualifiedName
	text := ctx.AllREAL_NAME()
	nameParts := make([]*quark.RealName, len(text))
	for index, node := range text {
		pos := quark.NewObjectPosition(ptc.file, node.GetSymbol().GetTokenIndex())
		text := node.GetText()
		nameParts[index] = quark.NewRealName(text, pos)
	}
	return &quark.QualifiedName{
		Parts: nameParts,
	}
}

func (ptc *ParseTreeConverter) VisitAtomicClock(ctx *AtomicClockContext) interface{} {
	return &quark.AtomicClock{ClockName: ptc.visitName(ctx.Name())}
}

func (ptc *ParseTreeConverter) VisitArrayIndexExpr(ctx *ArrayIndexExprContext) interface{} {
	allExprs := make([]quark.Expr, len(ctx.AllExpr()))
	for index, node := range ctx.AllExpr() {
		allExprs[index] = ptc.visitExpr(node)
	}

	arrayExpr := allExprs[0]
	indices := allExprs[1:]
	closeBracePos := ptc.terminalPosition(ctx.RBRACE())
	return quark.NewArrayIndex(arrayExpr, indices, closeBracePos)
}

func (ptc *ParseTreeConverter) VisitBranchExpr(ctx *BranchExprContext) interface{} {
	return &quark.BranchExpr{X: ptc.visitBranch(ctx.Branch())}
}

func (ptc *ParseTreeConverter) VisitLambdaExpr(ctx *LambdaExprContext) interface{} {
	lambdaPos := ptc.terminalPosition(ctx.KW_LAMBDA())
	args := ptc.visitArgumentList(ctx.Argumentlist())
	body := ptc.visitBlock(ctx.Block())
	var endExpr quark.Expr = nil
	if ctx.Expr() != nil {
		endExpr = ptc.visitExpr(ctx.Expr())
	}
	rCurly := ptc.terminalPosition(ctx.RCURLY())

	return quark.NewLambdaExpr(args, body, endExpr, lambdaPos, rCurly)
}

func (ptc *ParseTreeConverter) VisitConstructorExpr(ctx *ConstructorExprContext) interface{} {
	openCurly := ptc.terminalPosition(ctx.LCURLY())
	closeCurly := ptc.terminalPosition(ctx.RCURLY())
	args := ptc.visitCallArgList(ctx.Callarglist())

	return quark.NewConstructorExpr(args, openCurly, closeCurly)
}

func (ptc *ParseTreeConverter) VisitTupleExpr(ctx *TupleExprContext) interface{} {
	exprs := make([]quark.Expr, len(ctx.AllExpr()))
	for index, node := range ctx.AllExpr() {
		exprs[index] = ptc.visitExpr(node)
	}
	openParen := ptc.terminalPosition(ctx.LPAREN())
	closeParen := ptc.terminalPosition(ctx.RPAREN())
	return quark.NewTupleExpr(exprs, openParen, closeParen)
}

func (ptc *ParseTreeConverter) VisitLogicalBinopExpr(ctx *LogicalBinopExprContext) interface{} {
	left := ptc.visitExpr(ctx.Expr(0))
	right := ptc.visitExpr(ctx.Expr(1))
	op := ptc.visitBinop(ctx.op)
	opPos := ptc.tokenPosition(ctx.op)
	return quark.NewBinOp(left, right, op, opPos)
}

func (ptc *ParseTreeConverter) VisitConcatExpr(ctx *ConcatExprContext) interface{} {
	concat := ctx.Concat().(*ConcatContext) //Note: only option
	inners := make([]quark.InnerConcat, len(concat.AllInnerconcat()))
	for index, node := range concat.AllInnerconcat() {
		inners[index] = ptc.visitInnerConcat(node)
	}
	lCurly := ptc.terminalPosition(concat.LCURLY())
	rCurly := ptc.terminalPosition(concat.RCURLY())
	return quark.NewConcatExpr(inners, lCurly, rCurly)
}

func (ptc *ParseTreeConverter) VisitMulDivModExpr(ctx *MulDivModExprContext) interface{} {
	left := ptc.visitExpr(ctx.Expr(0))
	right := ptc.visitExpr(ctx.Expr(1))
	op := ptc.visitBinop(ctx.op)
	opPos := ptc.tokenPosition(ctx.op)
	return quark.NewBinOp(left, right, op, opPos)
}

func (ptc *ParseTreeConverter) VisitComplimentExpr(ctx *ComplimentExprContext) interface{} {
	expr := ptc.visitExpr(ctx.Expr())
	op := quark.OpComplement
	opPos := ptc.terminalPosition(ctx.OP_COMPLIMENT())
	return quark.NewUnOp(expr, op, opPos)
}

func (ptc *ParseTreeConverter) VisitArrayLiteralExpr(ctx *ArrayLiteralExprContext) interface{} {
	rBrace := ptc.terminalPosition(ctx.RBRACE())
	lBrace := ptc.terminalPosition(ctx.LBRACE())

	exprs := make([]quark.Expr, len(ctx.AllExpr()))
	for i, node := range ctx.AllExpr() {
		exprs[i] = ptc.visitExpr(node)
	}
	return quark.NewArrayLiteralExpr(exprs, rBrace, lBrace)
}

func (ptc *ParseTreeConverter) VisitClockToExpr(ctx *ClockToExprContext) interface{} {
	kwSignal := ptc.terminalPosition(ctx.KW_SIGNAL())
	rParen := ptc.terminalPosition(ctx.RPAREN())
	clockExpr := ptc.visitClockExpr(ctx.Clockexpr())
	return quark.NewClockToExpr(clockExpr, kwSignal, rParen)
}

func (ptc *ParseTreeConverter) VisitLiteralExpr(ctx *LiteralExprContext) interface{} {
	literal := ctx.Literal().(*LiteralContext) //only option
	return &quark.LiteralExpr{Value:ptc.VisitLiteral(literal).(*quark.Literal)}
}

func (ptc *ParseTreeConverter) VisitVarExpr(ctx *VarExprContext) interface{} {
	name := ptc.visitRealname(ctx.Realname())
	return &quark.VarExpr{VarName:name}
}

func (ptc *ParseTreeConverter) VisitParensExpr(ctx *ParensExprContext) interface{} {
	expr := ptc.visitExpr(ctx.Expr())
	start := ptc.terminalPosition(ctx.LPAREN())
	end := ptc.terminalPosition(ctx.RPAREN())
	return quark.NewParensExpr(expr, start, end)
}

func (ptc *ParseTreeConverter) VisitBitwiseBinopExpr(ctx *BitwiseBinopExprContext) interface{} {
	left := ptc.visitExpr(ctx.Expr(0))
	right := ptc.visitExpr(ctx.Expr(1))
	op := ptc.visitBinop(ctx.op)
	opPos := ptc.tokenPosition(ctx.op)
	return quark.NewBinOp(left, right, op, opPos)
}

func (ptc *ParseTreeConverter) VisitSliceExpr(ctx *SliceExprContext) interface{} {
	arrayExpr := ptc.visitExpr(ctx.Expr(0))
	var msb quark.Expr = nil
	var lsb quark.Expr = nil
	var step quark.Expr = nil
	rBracePos := ptc.terminalPosition(ctx.RBRACE())

	if ctx.msb != nil {
		msb = ptc.visitExpr(ctx.msb)
	}
	if ctx.lsb != nil {
		lsb = ptc.visitExpr(ctx.lsb)
	}
	if ctx.step != nil {
		step = ptc.visitExpr(ctx.step)
	}

	return quark.NewArraySlice(arrayExpr, msb, lsb, step, rBracePos)
}

func (ptc *ParseTreeConverter) VisitNotExpr(ctx *NotExprContext) interface{} {
	expr := ptc.visitExpr(ctx.Expr())
	op := quark.OpNot
	opPos := ptc.terminalPosition(ctx.OP_LNOT())
	return quark.NewUnOp(expr, op, opPos)
}

func (ptc *ParseTreeConverter) VisitSelectorExpr(ctx *SelectorExprContext) interface{} {
	name := ptc.visitExpr(ctx.Expr())
	fieldName := ptc.visitRealname(ctx.Realname())
	return &quark.SelectorExpr{
		Selectable: name,
		FieldName:  fieldName,
	}
}

func (ptc *ParseTreeConverter) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	left := ptc.visitExpr(ctx.Expr(0))
	right := ptc.visitExpr(ctx.Expr(1))
	op := ptc.visitBinop(ctx.op)
	opPos := ptc.tokenPosition(ctx.op)
	return quark.NewBinOp(left, right, op, opPos)
}

func (ptc *ParseTreeConverter) VisitShiftExpr(ctx *ShiftExprContext) interface{} {
	left := ptc.visitExpr(ctx.Expr(0))
	right := ptc.visitExpr(ctx.Expr(1))
	op := ptc.visitBinop(ctx.op)
	opPos := ptc.tokenPosition(ctx.op)
	return quark.NewBinOp(left, right, op, opPos)
}

func (ptc *ParseTreeConverter) VisitTernaryExpr(ctx *TernaryExprContext) interface{} {
	valueExpr := ptc.visitExpr(ctx.Expr(0))
	cond := ptc.visitExpr(ctx.Expr(1))
	elseExpr := ptc.visitExpr(ctx.Expr(2))
	
	return &quark.TernaryExpr{
		IfExpr:   valueExpr,
		Cond:     cond,
		ElseExpr: elseExpr,
	}
}

func (ptc *ParseTreeConverter) VisitNewModuleExpr(ctx *NewModuleExprContext) interface{} {
	moduleType := ptc.visitTypeExpr(ctx.Typeexpr())
	args := ptc.visitCallArgList(ctx.Callarglist())
	newKw := ptc.terminalPosition(ctx.KW_NEW())
	closeParen := ptc.terminalPosition(ctx.RPAREN())
	return quark.NewNewModuleExpr(moduleType, args, newKw, closeParen)
}

func (ptc *ParseTreeConverter) VisitNamedCallArg(ctx *NamedCallArgContext) interface{} {
	name := ptc.visitRealname(ctx.Realname())
	expr := ptc.visitExpr(ctx.Expr())

	return &quark.CallArgument{
		FieldName: name,
		ValueExpr: expr,
	}
}

func (ptc *ParseTreeConverter) VisitUnamedCallArg(ctx *UnamedCallArgContext) interface{} {
	expr := ptc.visitExpr(ctx.Expr())

	return &quark.CallArgument{
		FieldName: nil,
		ValueExpr: expr,
	}
}

func (ptc *ParseTreeConverter) VisitConcat(_ *ConcatContext) interface{} {
	panic("not reachable")
}

func (ptc *ParseTreeConverter) VisitInnerconcat(_ *InnerconcatContext) interface{} {
	panic("not reachable")
}

func (ptc *ParseTreeConverter) VisitParameterizedType(ctx *ParameterizedTypeContext) interface{} {
	mainType := ptc.visitTypeExpr(ctx.Typeexpr())
	typeParams := ptc.visitParamArgList(ctx.Paramarglist())

	closeBrace := ptc.terminalPosition(ctx.Paramarglist().(*ParamarglistContext).RANGLE())
	return quark.NewParameterizedType(mainType, typeParams, closeBrace)
}

func (ptc *ParseTreeConverter) VisitParamarg(ctx *ParamargContext) interface{} {
	var expr quark.Expr = nil
	var typeExpr quark.TypeExpr = nil
	var kwType quark.ObjectPosition
	var name *quark.RealName = nil

	if ctx.Expr() != nil {
		expr = ptc.visitExpr(ctx.Expr())
	} else {
		kwType = ptc.terminalPosition(ctx.KW_TYPE())
		typeExpr = ptc.visitTypeExpr(ctx.Typeexpr())
	}

	if ctx.Realname() != nil {
		name = ptc.visitRealname(ctx.Realname())
	}
	
	return &quark.ParamArgument{
		XExpr:     expr,
		XType:     typeExpr,
		ParamName: name,
		KwType:    kwType,
	}
}

func (ptc *ParseTreeConverter) VisitCompleteType(ctx *CompleteTypeContext) interface{} {
	name := ptc.visitName(ctx.Name())
	return &quark.CompleteType{X:name}
}

func (ptc *ParseTreeConverter) VisitIfBranch(ctx *IfBranchContext) interface{} {
	kwIf := ptc.terminalPosition(ctx.KW_IF())
	closeCurly := ptc.terminalPosition(ctx.RCURLY(len(ctx.AllRCURLY()) - 1))

	var ifPart quark.CondBranchPart
	ifPart.Cond = ptc.visitExpr(ctx.Expr(0))
	ifPart.Body = ptc.visitBlock(ctx.Block(0))

	elifParts := make([]quark.CondBranchPart, len(ctx.AllKW_ELIF()))
	for i := 0; i < len(ctx.AllKW_ELIF()); i++ {
		part := quark.CondBranchPart{
			Cond:       ptc.visitExpr(ctx.Expr(i + 1)),
			BranchPart: quark.BranchPart {
				Body: ptc.visitBlock(ctx.Block(i + 1)),
			},
		}
		elifParts[i] = part
	}

	var elsePart *quark.BranchPart = nil
	if ctx.KW_ELSE() != nil {
		elsePart = &quark.BranchPart{
			Body:       ptc.visitBlock(ctx.Block(len(ctx.AllBlock()) - 1)),
			BodyReturn: nil,
		}
	}

	return quark.NewIfBranch(ifPart, elifParts, elsePart, kwIf, closeCurly)
}

func (ptc *ParseTreeConverter) VisitMatchBranch(ctx *MatchBranchContext) interface{} {
	match := ptc.terminalPosition(ctx.KW_MATCH())
	end := ptc.terminalPosition(ctx.RCURLY(len(ctx.AllRCURLY()) - 1))

	matchExpr := ptc.visitExpr(ctx.Expr())

	var branchPart = make([]quark.CaseBranchPart, len(ctx.AllBlock()))
	for i := 0; i < len(ctx.AllBlock()); i++ {
		pat := ptc.visitPattern(ctx.Pattern(i))
		body := ptc.visitBlock(ctx.Block(i))
		branchPart[i] = quark.CaseBranchPart{
			CasePattern: pat,
			BranchPart:  quark.BranchPart{
				Body:       body,
				BodyReturn: nil,
			},
		}
	}

	return quark.NewMatchBranch(matchExpr, branchPart, match, end)
}

func (ptc *ParseTreeConverter) VisitAtomicPattern(ctx *AtomicPatternContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitParamerterizedTypePattern(ctx *ParamerterizedTypePatternContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitArrayPattern(ctx *ArrayPatternContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitLiteralPattern(ctx *LiteralPatternContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitStructPattern(ctx *StructPatternContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitParameterlist(ctx *ParameterlistContext) interface{} {
	params := make([]*quark.ParameterDef, len(ctx.AllParameterdef()))
	for i, node := range ctx.AllParameterdef() {
		params[i] = ptc.visitParameterDef(node)
	}

	return params
}

func (ptc *ParseTreeConverter) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	typeExpr := ptc.visitTypeExpr(ctx.Typeexpr())
	typePos := ptc.terminalPosition(ctx.KW_TYPE())
	
	return &quark.ParameterDef{
		IsType:  true,
		TypeVal: typeExpr,
		ParamName: nil,
		KwType:  typePos,
	}
}

func (ptc *ParseTreeConverter) VisitValueParameter(ctx *ValueParameterContext) interface{} {
	typeExpr := ptc.visitTypeExpr(ctx.Typeexpr())
	name := ptc.visitRealname(ctx.Realname())
	return &quark.ParameterDef{
		IsType:  false,
		TypeVal: typeExpr,
		ParamName: name,
		KwType:  quark.ObjectPosition{},
	}
}

func (ptc *ParseTreeConverter) VisitSingleReturn(ctx *SingleReturnContext) interface{} {
	return &quark.SingleReturn{ReturnType:ptc.visitTypeExpr(ctx.Typeexpr())}
}

func (ptc *ParseTreeConverter) VisitNamedReturn(ctx *NamedReturnContext) interface{} {
	returns := make([]quark.OneNamedReturn, len(ctx.AllTypeexpr()))

	for i := 0; i < len(ctx.AllTypeexpr()); i++ {
		types := ptc.visitTypeExpr(ctx.Typeexpr(i))
		names := ptc.visitRealname(ctx.Realname(i))

		returns[i] = quark.OneNamedReturn{
			ReturnName: names,
			ReturnType: types,
		}
	}

	start := ptc.terminalPosition(ctx.LPAREN())
	end := ptc.terminalPosition(ctx.RPAREN())

	return quark.NewNamedReturn(returns, start, end)
}

func (ptc *ParseTreeConverter) VisitArgumentdef(ctx *ArgumentdefContext) interface{} {
	return &quark.ArgumentDef{
		ArgType: ptc.visitTypeExpr(ctx.Typeexpr()),
		ArgName: ptc.visitRealname(ctx.Realname()),
	}
}

func (ptc *ParseTreeConverter) VisitArgumentlist(ctx *ArgumentlistContext) interface{} {
	args := make([]*quark.ArgumentDef, len(ctx.AllArgumentdef()))
	for i, node := range ctx.AllArgumentdef() {
		args[i] = ptc.VisitArgumentdef(node.(*ArgumentdefContext)).(*quark.ArgumentDef)
	}
	return args
}

func (ptc *ParseTreeConverter) VisitStructdecl(ctx *StructdeclContext) interface{} {
	kwStruct := ptc.terminalPosition(ctx.KW_STRUCT())

	rCurly := ptc.terminalPosition(ctx.RCURLY())

	name := ptc.visitRealname(ctx.Realname())
	var params []*quark.ParameterDef
	if ctx.Parameterlist() != nil {
		params = ptc.visitParameterList(ctx.Parameterlist())
	} else {
		params = make([]*quark.ParameterDef, 0)
	}

	traits := make([]quark.Name, len(ctx.AllName()))
	for i, node := range ctx.AllName() {
		traits[i] = ptc.visitName(node)
	}
	
	fields := make([]*quark.Field, len(ctx.AllFielddecl()))
	for i, node := range ctx.AllFielddecl() {
		fields[i] = ptc.VisitFielddecl(node.(*FielddeclContext)).(*quark.Field)
	}

	return &quark.StructDecl{
		Annotations: nil,
		StructName:  name,
		Parameters:  params,
		Fields:      fields,
		TraitImpls:  traits,
		KwStruct:    kwStruct,
		CloseCurly:  rCurly,
	}

}

func (ptc *ParseTreeConverter) VisitFielddecl(ctx *FielddeclContext) interface{} {
	fieldType := ptc.visitTypeExpr(ctx.Typeexpr())
	fieldName := ptc.visitRealname(ctx.Realname())
	var kwFuture quark.ObjectPosition
	isFuture := false
	if ctx.KW_FUTURE() != nil {
		kwFuture = ptc.terminalPosition(ctx.KW_FUTURE())
		isFuture = true
	}

	return &quark.Field{
		FieldType: fieldType,
		FieldName: fieldName,
		IsFuture: isFuture,
		KwFuture: kwFuture,
	}
}

func (ptc *ParseTreeConverter) VisitFuncdecl(ctx *FuncdeclContext) interface{} {
	println("visiting funcdecl")
	name := ptc.visitRealname(ctx.Realname())
	var params []*quark.ParameterDef
	if ctx.Parameterlist() != nil {
		params = ptc.visitParameterList(ctx.Parameterlist())
	} else {
		params = make([]*quark.ParameterDef, 0)
	}
	args := ptc.visitArgumentList(ctx.Argumentlist())
	returns := ptc.visitReturnList(ctx.Returnlist())
	body := ptc.visitBlock(ctx.Block())

	kwFunction := ptc.terminalPosition(ctx.KW_DEF())
	end := ptc.terminalPosition(ctx.RCURLY())

	return &quark.FunctionDecl{
		Annotations: nil,
		SymbolName:  name,
		Parameters:  params,
		Arguments:   args,
		Returns:     returns,
		Body:        body,
		KwFunction:  kwFunction,
		CloseCurly:  end,
	}
}

func (ptc *ParseTreeConverter) VisitModuledecl(ctx *ModuledeclContext) interface{} {
	name := ptc.visitRealname(ctx.Realname())
	var params []*quark.ParameterDef
	if ctx.Parameterlist() != nil {
		params = ptc.visitParameterList(ctx.Parameterlist())
	} else {
		params = make([]*quark.ParameterDef, 0)
	}
	args := ptc.visitArgumentList(ctx.Argumentlist())
	returns := ptc.visitReturnList(ctx.Returnlist())
	body := ptc.visitBlock(ctx.Block())

	kwModule := ptc.terminalPosition(ctx.KW_MODULE())
	end := ptc.terminalPosition(ctx.RCURLY())

	return &quark.ModuleDecl{
		Annotations: nil,
		SymbolName:  name,
		Parameters:  params,
		Arguments:   args,
		Returns:     returns,
		Body:        body,
		KwModule:  kwModule,
		CloseCurly:  end,
	}
}

func (ptc *ParseTreeConverter) VisitAnnotation(ctx *AnnotationContext) interface{} {
	name := ptc.visitRealname(ctx.Realname())
	start := ptc.terminalPosition(ctx.ANNOTATION_START())

	return &quark.Annotation{AnnotationName: name, Pos: start}
}

func (ptc *ParseTreeConverter) VisitLiteral(ctx *LiteralContext) interface{} {
	pos := quark.NewObjectPosition(ptc.file, ctx.GetStart().GetTokenIndex())
	text := ctx.INTEGRAL().GetText()
	return quark.NewLiteral(text, pos)
}

func (ptc *ParseTreeConverter) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(ptc)
}



//func (ptc *ParseTreeConverter) VisitChildren(tree antlr.RuleNode) interface{} {
//	return tree.Accept(ptc)
//}

