// Code generated from src/parser/QuarkParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // QuarkParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQuarkParserListener is a complete listener for a parse tree produced by QuarkParser.
type BaseQuarkParserListener struct{}

var _ QuarkParserListener = &BaseQuarkParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQuarkParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQuarkParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQuarkParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQuarkParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuarkpackage is called when production quarkpackage is entered.
func (s *BaseQuarkParserListener) EnterQuarkpackage(ctx *QuarkpackageContext) {}

// ExitQuarkpackage is called when production quarkpackage is exited.
func (s *BaseQuarkParserListener) ExitQuarkpackage(ctx *QuarkpackageContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseQuarkParserListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseQuarkParserListener) ExitDecl(ctx *DeclContext) {}

// EnterSingleImport is called when production SingleImport is entered.
func (s *BaseQuarkParserListener) EnterSingleImport(ctx *SingleImportContext) {}

// ExitSingleImport is called when production SingleImport is exited.
func (s *BaseQuarkParserListener) ExitSingleImport(ctx *SingleImportContext) {}

// EnterWildcardImport is called when production WildcardImport is entered.
func (s *BaseQuarkParserListener) EnterWildcardImport(ctx *WildcardImportContext) {}

// ExitWildcardImport is called when production WildcardImport is exited.
func (s *BaseQuarkParserListener) ExitWildcardImport(ctx *WildcardImportContext) {}

// EnterMultiImport is called when production MultiImport is entered.
func (s *BaseQuarkParserListener) EnterMultiImport(ctx *MultiImportContext) {}

// ExitMultiImport is called when production MultiImport is exited.
func (s *BaseQuarkParserListener) ExitMultiImport(ctx *MultiImportContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseQuarkParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseQuarkParserListener) ExitBlock(ctx *BlockContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseQuarkParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseQuarkParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssignStmt is called when production AssignStmt is entered.
func (s *BaseQuarkParserListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production AssignStmt is exited.
func (s *BaseQuarkParserListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterRegAssignStmt is called when production RegAssignStmt is entered.
func (s *BaseQuarkParserListener) EnterRegAssignStmt(ctx *RegAssignStmtContext) {}

// ExitRegAssignStmt is called when production RegAssignStmt is exited.
func (s *BaseQuarkParserListener) ExitRegAssignStmt(ctx *RegAssignStmtContext) {}

// EnterFutureStmt is called when production FutureStmt is entered.
func (s *BaseQuarkParserListener) EnterFutureStmt(ctx *FutureStmtContext) {}

// ExitFutureStmt is called when production FutureStmt is exited.
func (s *BaseQuarkParserListener) ExitFutureStmt(ctx *FutureStmtContext) {}

// EnterDeclarationStmt is called when production DeclarationStmt is entered.
func (s *BaseQuarkParserListener) EnterDeclarationStmt(ctx *DeclarationStmtContext) {}

// ExitDeclarationStmt is called when production DeclarationStmt is exited.
func (s *BaseQuarkParserListener) ExitDeclarationStmt(ctx *DeclarationStmtContext) {}

// EnterBranchStmt is called when production BranchStmt is entered.
func (s *BaseQuarkParserListener) EnterBranchStmt(ctx *BranchStmtContext) {}

// ExitBranchStmt is called when production BranchStmt is exited.
func (s *BaseQuarkParserListener) ExitBranchStmt(ctx *BranchStmtContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BaseQuarkParserListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BaseQuarkParserListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterValueAssignment is called when production ValueAssignment is entered.
func (s *BaseQuarkParserListener) EnterValueAssignment(ctx *ValueAssignmentContext) {}

// ExitValueAssignment is called when production ValueAssignment is exited.
func (s *BaseQuarkParserListener) ExitValueAssignment(ctx *ValueAssignmentContext) {}

// EnterVariableDefinition is called when production VariableDefinition is entered.
func (s *BaseQuarkParserListener) EnterVariableDefinition(ctx *VariableDefinitionContext) {}

// ExitVariableDefinition is called when production VariableDefinition is exited.
func (s *BaseQuarkParserListener) ExitVariableDefinition(ctx *VariableDefinitionContext) {}

// EnterTupleDestructer is called when production TupleDestructer is entered.
func (s *BaseQuarkParserListener) EnterTupleDestructer(ctx *TupleDestructerContext) {}

// ExitTupleDestructer is called when production TupleDestructer is exited.
func (s *BaseQuarkParserListener) ExitTupleDestructer(ctx *TupleDestructerContext) {}

// EnterArrayIndexAssignment is called when production ArrayIndexAssignment is entered.
func (s *BaseQuarkParserListener) EnterArrayIndexAssignment(ctx *ArrayIndexAssignmentContext) {}

// ExitArrayIndexAssignment is called when production ArrayIndexAssignment is exited.
func (s *BaseQuarkParserListener) ExitArrayIndexAssignment(ctx *ArrayIndexAssignmentContext) {}

// EnterArraySliceAssignment is called when production ArraySliceAssignment is entered.
func (s *BaseQuarkParserListener) EnterArraySliceAssignment(ctx *ArraySliceAssignmentContext) {}

// ExitArraySliceAssignment is called when production ArraySliceAssignment is exited.
func (s *BaseQuarkParserListener) ExitArraySliceAssignment(ctx *ArraySliceAssignmentContext) {}

// EnterRealname is called when production realname is entered.
func (s *BaseQuarkParserListener) EnterRealname(ctx *RealnameContext) {}

// ExitRealname is called when production realname is exited.
func (s *BaseQuarkParserListener) ExitRealname(ctx *RealnameContext) {}

// EnterUnqualifiedName is called when production UnqualifiedName is entered.
func (s *BaseQuarkParserListener) EnterUnqualifiedName(ctx *UnqualifiedNameContext) {}

// ExitUnqualifiedName is called when production UnqualifiedName is exited.
func (s *BaseQuarkParserListener) ExitUnqualifiedName(ctx *UnqualifiedNameContext) {}

// EnterQualifiedName is called when production QualifiedName is entered.
func (s *BaseQuarkParserListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production QualifiedName is exited.
func (s *BaseQuarkParserListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterAtomicClock is called when production AtomicClock is entered.
func (s *BaseQuarkParserListener) EnterAtomicClock(ctx *AtomicClockContext) {}

// ExitAtomicClock is called when production AtomicClock is exited.
func (s *BaseQuarkParserListener) ExitAtomicClock(ctx *AtomicClockContext) {}

// EnterArrayIndexExpr is called when production ArrayIndexExpr is entered.
func (s *BaseQuarkParserListener) EnterArrayIndexExpr(ctx *ArrayIndexExprContext) {}

// ExitArrayIndexExpr is called when production ArrayIndexExpr is exited.
func (s *BaseQuarkParserListener) ExitArrayIndexExpr(ctx *ArrayIndexExprContext) {}

// EnterBranchExpr is called when production BranchExpr is entered.
func (s *BaseQuarkParserListener) EnterBranchExpr(ctx *BranchExprContext) {}

// ExitBranchExpr is called when production BranchExpr is exited.
func (s *BaseQuarkParserListener) ExitBranchExpr(ctx *BranchExprContext) {}

// EnterLambdaExpr is called when production LambdaExpr is entered.
func (s *BaseQuarkParserListener) EnterLambdaExpr(ctx *LambdaExprContext) {}

// ExitLambdaExpr is called when production LambdaExpr is exited.
func (s *BaseQuarkParserListener) ExitLambdaExpr(ctx *LambdaExprContext) {}

// EnterConstructorExpr is called when production ConstructorExpr is entered.
func (s *BaseQuarkParserListener) EnterConstructorExpr(ctx *ConstructorExprContext) {}

// ExitConstructorExpr is called when production ConstructorExpr is exited.
func (s *BaseQuarkParserListener) ExitConstructorExpr(ctx *ConstructorExprContext) {}

// EnterTupleExpr is called when production TupleExpr is entered.
func (s *BaseQuarkParserListener) EnterTupleExpr(ctx *TupleExprContext) {}

// ExitTupleExpr is called when production TupleExpr is exited.
func (s *BaseQuarkParserListener) ExitTupleExpr(ctx *TupleExprContext) {}

// EnterLogicalBinopExpr is called when production LogicalBinopExpr is entered.
func (s *BaseQuarkParserListener) EnterLogicalBinopExpr(ctx *LogicalBinopExprContext) {}

// ExitLogicalBinopExpr is called when production LogicalBinopExpr is exited.
func (s *BaseQuarkParserListener) ExitLogicalBinopExpr(ctx *LogicalBinopExprContext) {}

// EnterConcatExpr is called when production ConcatExpr is entered.
func (s *BaseQuarkParserListener) EnterConcatExpr(ctx *ConcatExprContext) {}

// ExitConcatExpr is called when production ConcatExpr is exited.
func (s *BaseQuarkParserListener) ExitConcatExpr(ctx *ConcatExprContext) {}

// EnterSelectorExpr is called when production SelectorExpr is entered.
func (s *BaseQuarkParserListener) EnterSelectorExpr(ctx *SelectorExprContext) {}

// ExitSelectorExpr is called when production SelectorExpr is exited.
func (s *BaseQuarkParserListener) ExitSelectorExpr(ctx *SelectorExprContext) {}

// EnterMulDivModExpr is called when production MulDivModExpr is entered.
func (s *BaseQuarkParserListener) EnterMulDivModExpr(ctx *MulDivModExprContext) {}

// ExitMulDivModExpr is called when production MulDivModExpr is exited.
func (s *BaseQuarkParserListener) ExitMulDivModExpr(ctx *MulDivModExprContext) {}

// EnterCompareExpr is called when production CompareExpr is entered.
func (s *BaseQuarkParserListener) EnterCompareExpr(ctx *CompareExprContext) {}

// ExitCompareExpr is called when production CompareExpr is exited.
func (s *BaseQuarkParserListener) ExitCompareExpr(ctx *CompareExprContext) {}

// EnterComplimentExpr is called when production ComplimentExpr is entered.
func (s *BaseQuarkParserListener) EnterComplimentExpr(ctx *ComplimentExprContext) {}

// ExitComplimentExpr is called when production ComplimentExpr is exited.
func (s *BaseQuarkParserListener) ExitComplimentExpr(ctx *ComplimentExprContext) {}

// EnterArrayLiteralExpr is called when production ArrayLiteralExpr is entered.
func (s *BaseQuarkParserListener) EnterArrayLiteralExpr(ctx *ArrayLiteralExprContext) {}

// ExitArrayLiteralExpr is called when production ArrayLiteralExpr is exited.
func (s *BaseQuarkParserListener) ExitArrayLiteralExpr(ctx *ArrayLiteralExprContext) {}

// EnterClockToExpr is called when production ClockToExpr is entered.
func (s *BaseQuarkParserListener) EnterClockToExpr(ctx *ClockToExprContext) {}

// ExitClockToExpr is called when production ClockToExpr is exited.
func (s *BaseQuarkParserListener) ExitClockToExpr(ctx *ClockToExprContext) {}

// EnterLiteralExpr is called when production LiteralExpr is entered.
func (s *BaseQuarkParserListener) EnterLiteralExpr(ctx *LiteralExprContext) {}

// ExitLiteralExpr is called when production LiteralExpr is exited.
func (s *BaseQuarkParserListener) ExitLiteralExpr(ctx *LiteralExprContext) {}

// EnterVarExpr is called when production VarExpr is entered.
func (s *BaseQuarkParserListener) EnterVarExpr(ctx *VarExprContext) {}

// ExitVarExpr is called when production VarExpr is exited.
func (s *BaseQuarkParserListener) ExitVarExpr(ctx *VarExprContext) {}

// EnterParensExpr is called when production ParensExpr is entered.
func (s *BaseQuarkParserListener) EnterParensExpr(ctx *ParensExprContext) {}

// ExitParensExpr is called when production ParensExpr is exited.
func (s *BaseQuarkParserListener) ExitParensExpr(ctx *ParensExprContext) {}

// EnterBitwiseBinopExpr is called when production BitwiseBinopExpr is entered.
func (s *BaseQuarkParserListener) EnterBitwiseBinopExpr(ctx *BitwiseBinopExprContext) {}

// ExitBitwiseBinopExpr is called when production BitwiseBinopExpr is exited.
func (s *BaseQuarkParserListener) ExitBitwiseBinopExpr(ctx *BitwiseBinopExprContext) {}

// EnterSliceExpr is called when production SliceExpr is entered.
func (s *BaseQuarkParserListener) EnterSliceExpr(ctx *SliceExprContext) {}

// ExitSliceExpr is called when production SliceExpr is exited.
func (s *BaseQuarkParserListener) ExitSliceExpr(ctx *SliceExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseQuarkParserListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseQuarkParserListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BaseQuarkParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BaseQuarkParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAddSubExpr is called when production AddSubExpr is entered.
func (s *BaseQuarkParserListener) EnterAddSubExpr(ctx *AddSubExprContext) {}

// ExitAddSubExpr is called when production AddSubExpr is exited.
func (s *BaseQuarkParserListener) ExitAddSubExpr(ctx *AddSubExprContext) {}

// EnterShiftExpr is called when production ShiftExpr is entered.
func (s *BaseQuarkParserListener) EnterShiftExpr(ctx *ShiftExprContext) {}

// ExitShiftExpr is called when production ShiftExpr is exited.
func (s *BaseQuarkParserListener) ExitShiftExpr(ctx *ShiftExprContext) {}

// EnterTernaryExpr is called when production TernaryExpr is entered.
func (s *BaseQuarkParserListener) EnterTernaryExpr(ctx *TernaryExprContext) {}

// ExitTernaryExpr is called when production TernaryExpr is exited.
func (s *BaseQuarkParserListener) ExitTernaryExpr(ctx *TernaryExprContext) {}

// EnterNewModuleExpr is called when production NewModuleExpr is entered.
func (s *BaseQuarkParserListener) EnterNewModuleExpr(ctx *NewModuleExprContext) {}

// ExitNewModuleExpr is called when production NewModuleExpr is exited.
func (s *BaseQuarkParserListener) ExitNewModuleExpr(ctx *NewModuleExprContext) {}

// EnterCallarglist is called when production callarglist is entered.
func (s *BaseQuarkParserListener) EnterCallarglist(ctx *CallarglistContext) {}

// ExitCallarglist is called when production callarglist is exited.
func (s *BaseQuarkParserListener) ExitCallarglist(ctx *CallarglistContext) {}

// EnterNamedCallArg is called when production NamedCallArg is entered.
func (s *BaseQuarkParserListener) EnterNamedCallArg(ctx *NamedCallArgContext) {}

// ExitNamedCallArg is called when production NamedCallArg is exited.
func (s *BaseQuarkParserListener) ExitNamedCallArg(ctx *NamedCallArgContext) {}

// EnterUnamedCallArg is called when production UnamedCallArg is entered.
func (s *BaseQuarkParserListener) EnterUnamedCallArg(ctx *UnamedCallArgContext) {}

// ExitUnamedCallArg is called when production UnamedCallArg is exited.
func (s *BaseQuarkParserListener) ExitUnamedCallArg(ctx *UnamedCallArgContext) {}

// EnterParamarglist is called when production paramarglist is entered.
func (s *BaseQuarkParserListener) EnterParamarglist(ctx *ParamarglistContext) {}

// ExitParamarglist is called when production paramarglist is exited.
func (s *BaseQuarkParserListener) ExitParamarglist(ctx *ParamarglistContext) {}

// EnterConcat is called when production concat is entered.
func (s *BaseQuarkParserListener) EnterConcat(ctx *ConcatContext) {}

// ExitConcat is called when production concat is exited.
func (s *BaseQuarkParserListener) ExitConcat(ctx *ConcatContext) {}

// EnterInnerconcat is called when production innerconcat is entered.
func (s *BaseQuarkParserListener) EnterInnerconcat(ctx *InnerconcatContext) {}

// ExitInnerconcat is called when production innerconcat is exited.
func (s *BaseQuarkParserListener) ExitInnerconcat(ctx *InnerconcatContext) {}

// EnterParameterizedType is called when production ParameterizedType is entered.
func (s *BaseQuarkParserListener) EnterParameterizedType(ctx *ParameterizedTypeContext) {}

// ExitParameterizedType is called when production ParameterizedType is exited.
func (s *BaseQuarkParserListener) ExitParameterizedType(ctx *ParameterizedTypeContext) {}

// EnterArrayType is called when production ArrayType is entered.
func (s *BaseQuarkParserListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production ArrayType is exited.
func (s *BaseQuarkParserListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterCompleteType is called when production CompleteType is entered.
func (s *BaseQuarkParserListener) EnterCompleteType(ctx *CompleteTypeContext) {}

// ExitCompleteType is called when production CompleteType is exited.
func (s *BaseQuarkParserListener) ExitCompleteType(ctx *CompleteTypeContext) {}

// EnterParamarg is called when production paramarg is entered.
func (s *BaseQuarkParserListener) EnterParamarg(ctx *ParamargContext) {}

// ExitParamarg is called when production paramarg is exited.
func (s *BaseQuarkParserListener) ExitParamarg(ctx *ParamargContext) {}

// EnterIfBranch is called when production IfBranch is entered.
func (s *BaseQuarkParserListener) EnterIfBranch(ctx *IfBranchContext) {}

// ExitIfBranch is called when production IfBranch is exited.
func (s *BaseQuarkParserListener) ExitIfBranch(ctx *IfBranchContext) {}

// EnterMatchBranch is called when production MatchBranch is entered.
func (s *BaseQuarkParserListener) EnterMatchBranch(ctx *MatchBranchContext) {}

// ExitMatchBranch is called when production MatchBranch is exited.
func (s *BaseQuarkParserListener) ExitMatchBranch(ctx *MatchBranchContext) {}

// EnterAtomicPattern is called when production AtomicPattern is entered.
func (s *BaseQuarkParserListener) EnterAtomicPattern(ctx *AtomicPatternContext) {}

// ExitAtomicPattern is called when production AtomicPattern is exited.
func (s *BaseQuarkParserListener) ExitAtomicPattern(ctx *AtomicPatternContext) {}

// EnterParamerterizedTypePattern is called when production ParamerterizedTypePattern is entered.
func (s *BaseQuarkParserListener) EnterParamerterizedTypePattern(ctx *ParamerterizedTypePatternContext) {
}

// ExitParamerterizedTypePattern is called when production ParamerterizedTypePattern is exited.
func (s *BaseQuarkParserListener) ExitParamerterizedTypePattern(ctx *ParamerterizedTypePatternContext) {
}

// EnterArrayPattern is called when production ArrayPattern is entered.
func (s *BaseQuarkParserListener) EnterArrayPattern(ctx *ArrayPatternContext) {}

// ExitArrayPattern is called when production ArrayPattern is exited.
func (s *BaseQuarkParserListener) ExitArrayPattern(ctx *ArrayPatternContext) {}

// EnterLiteralPattern is called when production LiteralPattern is entered.
func (s *BaseQuarkParserListener) EnterLiteralPattern(ctx *LiteralPatternContext) {}

// ExitLiteralPattern is called when production LiteralPattern is exited.
func (s *BaseQuarkParserListener) ExitLiteralPattern(ctx *LiteralPatternContext) {}

// EnterStructPattern is called when production StructPattern is entered.
func (s *BaseQuarkParserListener) EnterStructPattern(ctx *StructPatternContext) {}

// ExitStructPattern is called when production StructPattern is exited.
func (s *BaseQuarkParserListener) ExitStructPattern(ctx *StructPatternContext) {}

// EnterParameterlist is called when production parameterlist is entered.
func (s *BaseQuarkParserListener) EnterParameterlist(ctx *ParameterlistContext) {}

// ExitParameterlist is called when production parameterlist is exited.
func (s *BaseQuarkParserListener) ExitParameterlist(ctx *ParameterlistContext) {}

// EnterTypeParameter is called when production TypeParameter is entered.
func (s *BaseQuarkParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production TypeParameter is exited.
func (s *BaseQuarkParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterValueParameter is called when production ValueParameter is entered.
func (s *BaseQuarkParserListener) EnterValueParameter(ctx *ValueParameterContext) {}

// ExitValueParameter is called when production ValueParameter is exited.
func (s *BaseQuarkParserListener) ExitValueParameter(ctx *ValueParameterContext) {}

// EnterSingleReturn is called when production SingleReturn is entered.
func (s *BaseQuarkParserListener) EnterSingleReturn(ctx *SingleReturnContext) {}

// ExitSingleReturn is called when production SingleReturn is exited.
func (s *BaseQuarkParserListener) ExitSingleReturn(ctx *SingleReturnContext) {}

// EnterNamedReturn is called when production NamedReturn is entered.
func (s *BaseQuarkParserListener) EnterNamedReturn(ctx *NamedReturnContext) {}

// ExitNamedReturn is called when production NamedReturn is exited.
func (s *BaseQuarkParserListener) ExitNamedReturn(ctx *NamedReturnContext) {}

// EnterArgumentdef is called when production argumentdef is entered.
func (s *BaseQuarkParserListener) EnterArgumentdef(ctx *ArgumentdefContext) {}

// ExitArgumentdef is called when production argumentdef is exited.
func (s *BaseQuarkParserListener) ExitArgumentdef(ctx *ArgumentdefContext) {}

// EnterArgumentlist is called when production argumentlist is entered.
func (s *BaseQuarkParserListener) EnterArgumentlist(ctx *ArgumentlistContext) {}

// ExitArgumentlist is called when production argumentlist is exited.
func (s *BaseQuarkParserListener) ExitArgumentlist(ctx *ArgumentlistContext) {}

// EnterStructdecl is called when production structdecl is entered.
func (s *BaseQuarkParserListener) EnterStructdecl(ctx *StructdeclContext) {}

// ExitStructdecl is called when production structdecl is exited.
func (s *BaseQuarkParserListener) ExitStructdecl(ctx *StructdeclContext) {}

// EnterFielddecl is called when production fielddecl is entered.
func (s *BaseQuarkParserListener) EnterFielddecl(ctx *FielddeclContext) {}

// ExitFielddecl is called when production fielddecl is exited.
func (s *BaseQuarkParserListener) ExitFielddecl(ctx *FielddeclContext) {}

// EnterFuncsig is called when production funcsig is entered.
func (s *BaseQuarkParserListener) EnterFuncsig(ctx *FuncsigContext) {}

// ExitFuncsig is called when production funcsig is exited.
func (s *BaseQuarkParserListener) ExitFuncsig(ctx *FuncsigContext) {}

// EnterFuncdecl is called when production funcdecl is entered.
func (s *BaseQuarkParserListener) EnterFuncdecl(ctx *FuncdeclContext) {}

// ExitFuncdecl is called when production funcdecl is exited.
func (s *BaseQuarkParserListener) ExitFuncdecl(ctx *FuncdeclContext) {}

// EnterModuledecl is called when production moduledecl is entered.
func (s *BaseQuarkParserListener) EnterModuledecl(ctx *ModuledeclContext) {}

// ExitModuledecl is called when production moduledecl is exited.
func (s *BaseQuarkParserListener) ExitModuledecl(ctx *ModuledeclContext) {}

// EnterTraitdecl is called when production traitdecl is entered.
func (s *BaseQuarkParserListener) EnterTraitdecl(ctx *TraitdeclContext) {}

// ExitTraitdecl is called when production traitdecl is exited.
func (s *BaseQuarkParserListener) ExitTraitdecl(ctx *TraitdeclContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseQuarkParserListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseQuarkParserListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseQuarkParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseQuarkParserListener) ExitLiteral(ctx *LiteralContext) {}
