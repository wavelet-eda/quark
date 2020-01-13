//Home of the parser and lexer implementations.
package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wwerst/wavelet/wavelet-go/src/quark"
)

//Visitor implementation which converts parse tree from Antlr into
//proper quark.AST.
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

func (ptc *ParseTreeConverter) visitCallArgList(ctx ICallarglistContext) []*quark.CallArgument {
	return ptc.Visit(ctx).([]*quark.CallArgument)
}

func (ptc *ParseTreeConverter) visitBlock(ctx IBlockContext) []quark.Stmt {
	return ptc.Visit(ctx).([]quark.Stmt)
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

	return quark.Package{
		Imports: importDecls,
		Symbols: decls,
	}
}



func (ptc *ParseTreeConverter) visitImportDecl(rawImportDecl IImportdeclContext) quark.ImportDecl {
	return ptc.Visit(rawImportDecl).(quark.ImportDecl)
}

func (ptc *ParseTreeConverter) visitDecl(rawDecl IDeclContext) quark.Decl {
	return ptc.Visit(rawDecl).(quark.Decl)
}

func (ptc *ParseTreeConverter) VisitSingleImport(ctx *SingleImportContext) interface{} {
	println("visiting single import")
	quarkName := ptc.Visit(ctx.Name()).(quark.Name)
	return &quark.SingleImport{
		GenericImport: quark.GenericImport{PackageName:quarkName},
	}
}

func (ptc *ParseTreeConverter) VisitWildcardImport(ctx *WildcardImportContext) interface{} {
	quarkName := ptc.Visit(ctx.Name()).(quark.Name)
	return &quark.WildcardImport{
		GenericImport: quark.GenericImport{PackageName:quarkName},
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
	future := ctx.Future().(*FutureContext)
	futures := ptc.visitArgumentList(future.Argumentlist())
	body := ptc.visitBlock(future.Block())
	assignments := ptc.visitCallArgList(future.Callarglist())

	futurePos := ptc.terminalPosition(future.KW_FUTURE())
	semiPos := ptc.terminalPosition(ctx.SEMI())
	return quark.NewFutureStmt(futures, body, assignments, futurePos, semiPos)
}

func (ptc *ParseTreeConverter) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	expr := ptc.visitExpr(ctx.Expr())
	returnPos := ptc.terminalPosition(ctx.KW_RETURN())
	semiPos := ptc.terminalPosition(ctx.SEMI())
	return quark.NewReturnStmt(expr, returnPos, semiPos)
}

func (ptc *ParseTreeConverter) VisitFuture(_ *FutureContext) interface{} {
	panic("unreachable code")
}

func (ptc *ParseTreeConverter) VisitValueAssignment(ctx *ValueAssignmentContext) interface{} {
	return ptc.visitName(ctx.Name())
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
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitConstructorExpr(ctx *ConstructorExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitTupleExpr(ctx *TupleExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitLogicalBinopExpr(ctx *LogicalBinopExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitConcatExpr(ctx *ConcatExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitMulDivModExpr(ctx *MulDivModExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitComplimentExpr(ctx *ComplimentExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitArrayLiteralExpr(ctx *ArrayLiteralExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitClockToExpr(ctx *ClockToExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitLiteralExpr(ctx *LiteralExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitVarExpr(ctx *VarExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitParensExpr(ctx *ParensExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitBitwiseBinopExpr(ctx *BitwiseBinopExprContext) interface{} {
	panic("implement me")
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
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitFieldExpr(ctx *FieldExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitShiftExpr(ctx *ShiftExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitTernaryExpr(ctx *TernaryExprContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitNewModuleExpr(ctx *NewModuleExprContext) interface{} {
	panic("implement me")
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

func (ptc *ParseTreeConverter) VisitConcat(ctx *ConcatContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitInnerconcat(ctx *InnerconcatContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitParameterizedType(ctx *ParameterizedTypeContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitCompleteType(ctx *CompleteTypeContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitIfBranch(ctx *IfBranchContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitMatchBranch(ctx *MatchBranchContext) interface{} {
	panic("implement me")
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
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitValueParameter(ctx *ValueParameterContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitSingleReturn(ctx *SingleReturnContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitNamedReturn(ctx *NamedReturnContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitArgumentdef(ctx *ArgumentdefContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitArgumentlist(ctx *ArgumentlistContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitStructdecl(ctx *StructdeclContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitStructdef(ctx *StructdefContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitFielddecl(ctx *FielddeclContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitFuncdecl(ctx *FuncdeclContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitModuledecl(ctx *ModuledeclContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitInnermodule(ctx *InnermoduleContext) interface{} {
	panic("implement me")
}

func (ptc *ParseTreeConverter) VisitAnnotation(ctx *AnnotationContext) interface{} {
	panic("implement me")
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

