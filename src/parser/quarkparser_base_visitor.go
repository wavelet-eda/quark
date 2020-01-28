// Code generated from src/parser/QuarkParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // QuarkParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseQuarkParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseQuarkParserVisitor) VisitQuarkpackage(ctx *QuarkpackageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitDecl(ctx *DeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitSingleImport(ctx *SingleImportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitWildcardImport(ctx *WildcardImportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitMultiImport(ctx *MultiImportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitRegAssignStmt(ctx *RegAssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitFutureStmt(ctx *FutureStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitDeclarationStmt(ctx *DeclarationStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitBranchStmt(ctx *BranchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitValueAssignment(ctx *ValueAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitVariableDefinition(ctx *VariableDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitTupleDestructer(ctx *TupleDestructerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitArrayIndexAssignment(ctx *ArrayIndexAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitArraySliceAssignment(ctx *ArraySliceAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitRealname(ctx *RealnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitUnqualifiedName(ctx *UnqualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitAtomicClock(ctx *AtomicClockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitArrayIndexExpr(ctx *ArrayIndexExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitBranchExpr(ctx *BranchExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitLambdaExpr(ctx *LambdaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitConstructorExpr(ctx *ConstructorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitTupleExpr(ctx *TupleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitLogicalBinopExpr(ctx *LogicalBinopExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitConcatExpr(ctx *ConcatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitSelectorExpr(ctx *SelectorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitMulDivModExpr(ctx *MulDivModExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitComplimentExpr(ctx *ComplimentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitArrayLiteralExpr(ctx *ArrayLiteralExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitClockToExpr(ctx *ClockToExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitLiteralExpr(ctx *LiteralExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitVarExpr(ctx *VarExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitParensExpr(ctx *ParensExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitBitwiseBinopExpr(ctx *BitwiseBinopExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitSliceExpr(ctx *SliceExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitShiftExpr(ctx *ShiftExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitTernaryExpr(ctx *TernaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitNewModuleExpr(ctx *NewModuleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitCallarglist(ctx *CallarglistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitNamedCallArg(ctx *NamedCallArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitUnamedCallArg(ctx *UnamedCallArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitParamarglist(ctx *ParamarglistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitConcat(ctx *ConcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitInnerconcat(ctx *InnerconcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitParameterizedType(ctx *ParameterizedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitCompleteType(ctx *CompleteTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitParamarg(ctx *ParamargContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitIfBranch(ctx *IfBranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitMatchBranch(ctx *MatchBranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitAtomicPattern(ctx *AtomicPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitParamerterizedTypePattern(ctx *ParamerterizedTypePatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitArrayPattern(ctx *ArrayPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitLiteralPattern(ctx *LiteralPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitStructPattern(ctx *StructPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitParameterlist(ctx *ParameterlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitValueParameter(ctx *ValueParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitSingleReturn(ctx *SingleReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitNamedReturn(ctx *NamedReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitArgumentdef(ctx *ArgumentdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitArgumentlist(ctx *ArgumentlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitStructdecl(ctx *StructdeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitFielddecl(ctx *FielddeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitFuncsig(ctx *FuncsigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitFuncdecl(ctx *FuncdeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitModuledecl(ctx *ModuledeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitTraitdecl(ctx *TraitdeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQuarkParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
