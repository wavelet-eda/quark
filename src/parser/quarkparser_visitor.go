// Code generated from src/parser/QuarkParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // QuarkParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by QuarkParser.
type QuarkParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QuarkParser#quarkpackage.
	VisitQuarkpackage(ctx *QuarkpackageContext) interface{}

	// Visit a parse tree produced by QuarkParser#decl.
	VisitDecl(ctx *DeclContext) interface{}

	// Visit a parse tree produced by QuarkParser#SingleImport.
	VisitSingleImport(ctx *SingleImportContext) interface{}

	// Visit a parse tree produced by QuarkParser#WildcardImport.
	VisitWildcardImport(ctx *WildcardImportContext) interface{}

	// Visit a parse tree produced by QuarkParser#MultiImport.
	VisitMultiImport(ctx *MultiImportContext) interface{}

	// Visit a parse tree produced by QuarkParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by QuarkParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by QuarkParser#AssignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by QuarkParser#RegAssignStmt.
	VisitRegAssignStmt(ctx *RegAssignStmtContext) interface{}

	// Visit a parse tree produced by QuarkParser#FutureStmt.
	VisitFutureStmt(ctx *FutureStmtContext) interface{}

	// Visit a parse tree produced by QuarkParser#DeclarationStmt.
	VisitDeclarationStmt(ctx *DeclarationStmtContext) interface{}

	// Visit a parse tree produced by QuarkParser#BranchStmt.
	VisitBranchStmt(ctx *BranchStmtContext) interface{}

	// Visit a parse tree produced by QuarkParser#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by QuarkParser#ValueAssignment.
	VisitValueAssignment(ctx *ValueAssignmentContext) interface{}

	// Visit a parse tree produced by QuarkParser#VariableDefinition.
	VisitVariableDefinition(ctx *VariableDefinitionContext) interface{}

	// Visit a parse tree produced by QuarkParser#TupleDestructer.
	VisitTupleDestructer(ctx *TupleDestructerContext) interface{}

	// Visit a parse tree produced by QuarkParser#ArrayIndexAssignment.
	VisitArrayIndexAssignment(ctx *ArrayIndexAssignmentContext) interface{}

	// Visit a parse tree produced by QuarkParser#ArraySliceAssignment.
	VisitArraySliceAssignment(ctx *ArraySliceAssignmentContext) interface{}

	// Visit a parse tree produced by QuarkParser#realname.
	VisitRealname(ctx *RealnameContext) interface{}

	// Visit a parse tree produced by QuarkParser#UnqualifiedName.
	VisitUnqualifiedName(ctx *UnqualifiedNameContext) interface{}

	// Visit a parse tree produced by QuarkParser#QualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by QuarkParser#AtomicClock.
	VisitAtomicClock(ctx *AtomicClockContext) interface{}

	// Visit a parse tree produced by QuarkParser#ArrayIndexExpr.
	VisitArrayIndexExpr(ctx *ArrayIndexExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#BranchExpr.
	VisitBranchExpr(ctx *BranchExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#LambdaExpr.
	VisitLambdaExpr(ctx *LambdaExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#ConstructorExpr.
	VisitConstructorExpr(ctx *ConstructorExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#TupleExpr.
	VisitTupleExpr(ctx *TupleExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#LogicalBinopExpr.
	VisitLogicalBinopExpr(ctx *LogicalBinopExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#ConcatExpr.
	VisitConcatExpr(ctx *ConcatExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#SelectorExpr.
	VisitSelectorExpr(ctx *SelectorExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#MulDivModExpr.
	VisitMulDivModExpr(ctx *MulDivModExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#ComplimentExpr.
	VisitComplimentExpr(ctx *ComplimentExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#ArrayLiteralExpr.
	VisitArrayLiteralExpr(ctx *ArrayLiteralExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#ClockToExpr.
	VisitClockToExpr(ctx *ClockToExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#LiteralExpr.
	VisitLiteralExpr(ctx *LiteralExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#VarExpr.
	VisitVarExpr(ctx *VarExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#ParensExpr.
	VisitParensExpr(ctx *ParensExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#BitwiseBinopExpr.
	VisitBitwiseBinopExpr(ctx *BitwiseBinopExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#SliceExpr.
	VisitSliceExpr(ctx *SliceExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#FunctionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by QuarkParser#AddSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#ShiftExpr.
	VisitShiftExpr(ctx *ShiftExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#TernaryExpr.
	VisitTernaryExpr(ctx *TernaryExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#NewModuleExpr.
	VisitNewModuleExpr(ctx *NewModuleExprContext) interface{}

	// Visit a parse tree produced by QuarkParser#callarglist.
	VisitCallarglist(ctx *CallarglistContext) interface{}

	// Visit a parse tree produced by QuarkParser#NamedCallArg.
	VisitNamedCallArg(ctx *NamedCallArgContext) interface{}

	// Visit a parse tree produced by QuarkParser#UnamedCallArg.
	VisitUnamedCallArg(ctx *UnamedCallArgContext) interface{}

	// Visit a parse tree produced by QuarkParser#paramarglist.
	VisitParamarglist(ctx *ParamarglistContext) interface{}

	// Visit a parse tree produced by QuarkParser#concat.
	VisitConcat(ctx *ConcatContext) interface{}

	// Visit a parse tree produced by QuarkParser#innerconcat.
	VisitInnerconcat(ctx *InnerconcatContext) interface{}

	// Visit a parse tree produced by QuarkParser#ParameterizedType.
	VisitParameterizedType(ctx *ParameterizedTypeContext) interface{}

	// Visit a parse tree produced by QuarkParser#ArrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by QuarkParser#CompleteType.
	VisitCompleteType(ctx *CompleteTypeContext) interface{}

	// Visit a parse tree produced by QuarkParser#paramarg.
	VisitParamarg(ctx *ParamargContext) interface{}

	// Visit a parse tree produced by QuarkParser#IfBranch.
	VisitIfBranch(ctx *IfBranchContext) interface{}

	// Visit a parse tree produced by QuarkParser#MatchBranch.
	VisitMatchBranch(ctx *MatchBranchContext) interface{}

	// Visit a parse tree produced by QuarkParser#AtomicPattern.
	VisitAtomicPattern(ctx *AtomicPatternContext) interface{}

	// Visit a parse tree produced by QuarkParser#ParamerterizedTypePattern.
	VisitParamerterizedTypePattern(ctx *ParamerterizedTypePatternContext) interface{}

	// Visit a parse tree produced by QuarkParser#ArrayPattern.
	VisitArrayPattern(ctx *ArrayPatternContext) interface{}

	// Visit a parse tree produced by QuarkParser#LiteralPattern.
	VisitLiteralPattern(ctx *LiteralPatternContext) interface{}

	// Visit a parse tree produced by QuarkParser#StructPattern.
	VisitStructPattern(ctx *StructPatternContext) interface{}

	// Visit a parse tree produced by QuarkParser#parameterlist.
	VisitParameterlist(ctx *ParameterlistContext) interface{}

	// Visit a parse tree produced by QuarkParser#TypeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) interface{}

	// Visit a parse tree produced by QuarkParser#ValueParameter.
	VisitValueParameter(ctx *ValueParameterContext) interface{}

	// Visit a parse tree produced by QuarkParser#SingleReturn.
	VisitSingleReturn(ctx *SingleReturnContext) interface{}

	// Visit a parse tree produced by QuarkParser#NamedReturn.
	VisitNamedReturn(ctx *NamedReturnContext) interface{}

	// Visit a parse tree produced by QuarkParser#argumentdef.
	VisitArgumentdef(ctx *ArgumentdefContext) interface{}

	// Visit a parse tree produced by QuarkParser#argumentlist.
	VisitArgumentlist(ctx *ArgumentlistContext) interface{}

	// Visit a parse tree produced by QuarkParser#structdecl.
	VisitStructdecl(ctx *StructdeclContext) interface{}

	// Visit a parse tree produced by QuarkParser#fielddecl.
	VisitFielddecl(ctx *FielddeclContext) interface{}

	// Visit a parse tree produced by QuarkParser#funcsig.
	VisitFuncsig(ctx *FuncsigContext) interface{}

	// Visit a parse tree produced by QuarkParser#funcdecl.
	VisitFuncdecl(ctx *FuncdeclContext) interface{}

	// Visit a parse tree produced by QuarkParser#moduledecl.
	VisitModuledecl(ctx *ModuledeclContext) interface{}

	// Visit a parse tree produced by QuarkParser#traitdecl.
	VisitTraitdecl(ctx *TraitdeclContext) interface{}

	// Visit a parse tree produced by QuarkParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by QuarkParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
