// Code generated from src/parser/QuarkParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // QuarkParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// QuarkParserListener is a complete listener for a parse tree produced by QuarkParser.
type QuarkParserListener interface {
	antlr.ParseTreeListener

	// EnterQuarkpackage is called when entering the quarkpackage production.
	EnterQuarkpackage(c *QuarkpackageContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterSingleImport is called when entering the SingleImport production.
	EnterSingleImport(c *SingleImportContext)

	// EnterWildcardImport is called when entering the WildcardImport production.
	EnterWildcardImport(c *WildcardImportContext)

	// EnterMultiImport is called when entering the MultiImport production.
	EnterMultiImport(c *MultiImportContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterAssignStmt is called when entering the AssignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterRegAssignStmt is called when entering the RegAssignStmt production.
	EnterRegAssignStmt(c *RegAssignStmtContext)

	// EnterFutureStmt is called when entering the FutureStmt production.
	EnterFutureStmt(c *FutureStmtContext)

	// EnterDeclarationStmt is called when entering the DeclarationStmt production.
	EnterDeclarationStmt(c *DeclarationStmtContext)

	// EnterBranchStmt is called when entering the BranchStmt production.
	EnterBranchStmt(c *BranchStmtContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterValueAssignment is called when entering the ValueAssignment production.
	EnterValueAssignment(c *ValueAssignmentContext)

	// EnterVariableDefinition is called when entering the VariableDefinition production.
	EnterVariableDefinition(c *VariableDefinitionContext)

	// EnterTupleDestructer is called when entering the TupleDestructer production.
	EnterTupleDestructer(c *TupleDestructerContext)

	// EnterArrayIndexAssignment is called when entering the ArrayIndexAssignment production.
	EnterArrayIndexAssignment(c *ArrayIndexAssignmentContext)

	// EnterArraySliceAssignment is called when entering the ArraySliceAssignment production.
	EnterArraySliceAssignment(c *ArraySliceAssignmentContext)

	// EnterRealname is called when entering the realname production.
	EnterRealname(c *RealnameContext)

	// EnterUnqualifiedName is called when entering the UnqualifiedName production.
	EnterUnqualifiedName(c *UnqualifiedNameContext)

	// EnterQualifiedName is called when entering the QualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterAtomicClock is called when entering the AtomicClock production.
	EnterAtomicClock(c *AtomicClockContext)

	// EnterArrayIndexExpr is called when entering the ArrayIndexExpr production.
	EnterArrayIndexExpr(c *ArrayIndexExprContext)

	// EnterBranchExpr is called when entering the BranchExpr production.
	EnterBranchExpr(c *BranchExprContext)

	// EnterLambdaExpr is called when entering the LambdaExpr production.
	EnterLambdaExpr(c *LambdaExprContext)

	// EnterConstructorExpr is called when entering the ConstructorExpr production.
	EnterConstructorExpr(c *ConstructorExprContext)

	// EnterTupleExpr is called when entering the TupleExpr production.
	EnterTupleExpr(c *TupleExprContext)

	// EnterLogicalBinopExpr is called when entering the LogicalBinopExpr production.
	EnterLogicalBinopExpr(c *LogicalBinopExprContext)

	// EnterConcatExpr is called when entering the ConcatExpr production.
	EnterConcatExpr(c *ConcatExprContext)

	// EnterSelectorExpr is called when entering the SelectorExpr production.
	EnterSelectorExpr(c *SelectorExprContext)

	// EnterMulDivModExpr is called when entering the MulDivModExpr production.
	EnterMulDivModExpr(c *MulDivModExprContext)

	// EnterComplimentExpr is called when entering the ComplimentExpr production.
	EnterComplimentExpr(c *ComplimentExprContext)

	// EnterArrayLiteralExpr is called when entering the ArrayLiteralExpr production.
	EnterArrayLiteralExpr(c *ArrayLiteralExprContext)

	// EnterClockToExpr is called when entering the ClockToExpr production.
	EnterClockToExpr(c *ClockToExprContext)

	// EnterLiteralExpr is called when entering the LiteralExpr production.
	EnterLiteralExpr(c *LiteralExprContext)

	// EnterVarExpr is called when entering the VarExpr production.
	EnterVarExpr(c *VarExprContext)

	// EnterParensExpr is called when entering the ParensExpr production.
	EnterParensExpr(c *ParensExprContext)

	// EnterBitwiseBinopExpr is called when entering the BitwiseBinopExpr production.
	EnterBitwiseBinopExpr(c *BitwiseBinopExprContext)

	// EnterSliceExpr is called when entering the SliceExpr production.
	EnterSliceExpr(c *SliceExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterAddSubExpr is called when entering the AddSubExpr production.
	EnterAddSubExpr(c *AddSubExprContext)

	// EnterShiftExpr is called when entering the ShiftExpr production.
	EnterShiftExpr(c *ShiftExprContext)

	// EnterTernaryExpr is called when entering the TernaryExpr production.
	EnterTernaryExpr(c *TernaryExprContext)

	// EnterNewModuleExpr is called when entering the NewModuleExpr production.
	EnterNewModuleExpr(c *NewModuleExprContext)

	// EnterCallarglist is called when entering the callarglist production.
	EnterCallarglist(c *CallarglistContext)

	// EnterNamedCallArg is called when entering the NamedCallArg production.
	EnterNamedCallArg(c *NamedCallArgContext)

	// EnterUnamedCallArg is called when entering the UnamedCallArg production.
	EnterUnamedCallArg(c *UnamedCallArgContext)

	// EnterParamarglist is called when entering the paramarglist production.
	EnterParamarglist(c *ParamarglistContext)

	// EnterConcat is called when entering the concat production.
	EnterConcat(c *ConcatContext)

	// EnterInnerconcat is called when entering the innerconcat production.
	EnterInnerconcat(c *InnerconcatContext)

	// EnterParameterizedType is called when entering the ParameterizedType production.
	EnterParameterizedType(c *ParameterizedTypeContext)

	// EnterCompleteType is called when entering the CompleteType production.
	EnterCompleteType(c *CompleteTypeContext)

	// EnterParamarg is called when entering the paramarg production.
	EnterParamarg(c *ParamargContext)

	// EnterIfBranch is called when entering the IfBranch production.
	EnterIfBranch(c *IfBranchContext)

	// EnterMatchBranch is called when entering the MatchBranch production.
	EnterMatchBranch(c *MatchBranchContext)

	// EnterAtomicPattern is called when entering the AtomicPattern production.
	EnterAtomicPattern(c *AtomicPatternContext)

	// EnterParamerterizedTypePattern is called when entering the ParamerterizedTypePattern production.
	EnterParamerterizedTypePattern(c *ParamerterizedTypePatternContext)

	// EnterArrayPattern is called when entering the ArrayPattern production.
	EnterArrayPattern(c *ArrayPatternContext)

	// EnterLiteralPattern is called when entering the LiteralPattern production.
	EnterLiteralPattern(c *LiteralPatternContext)

	// EnterStructPattern is called when entering the StructPattern production.
	EnterStructPattern(c *StructPatternContext)

	// EnterParameterlist is called when entering the parameterlist production.
	EnterParameterlist(c *ParameterlistContext)

	// EnterTypeParameter is called when entering the TypeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterValueParameter is called when entering the ValueParameter production.
	EnterValueParameter(c *ValueParameterContext)

	// EnterSingleReturn is called when entering the SingleReturn production.
	EnterSingleReturn(c *SingleReturnContext)

	// EnterNamedReturn is called when entering the NamedReturn production.
	EnterNamedReturn(c *NamedReturnContext)

	// EnterArgumentdef is called when entering the argumentdef production.
	EnterArgumentdef(c *ArgumentdefContext)

	// EnterArgumentlist is called when entering the argumentlist production.
	EnterArgumentlist(c *ArgumentlistContext)

	// EnterStructdecl is called when entering the structdecl production.
	EnterStructdecl(c *StructdeclContext)

	// EnterFielddecl is called when entering the fielddecl production.
	EnterFielddecl(c *FielddeclContext)

	// EnterFuncdecl is called when entering the funcdecl production.
	EnterFuncdecl(c *FuncdeclContext)

	// EnterModuledecl is called when entering the moduledecl production.
	EnterModuledecl(c *ModuledeclContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitQuarkpackage is called when exiting the quarkpackage production.
	ExitQuarkpackage(c *QuarkpackageContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitSingleImport is called when exiting the SingleImport production.
	ExitSingleImport(c *SingleImportContext)

	// ExitWildcardImport is called when exiting the WildcardImport production.
	ExitWildcardImport(c *WildcardImportContext)

	// ExitMultiImport is called when exiting the MultiImport production.
	ExitMultiImport(c *MultiImportContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitAssignStmt is called when exiting the AssignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitRegAssignStmt is called when exiting the RegAssignStmt production.
	ExitRegAssignStmt(c *RegAssignStmtContext)

	// ExitFutureStmt is called when exiting the FutureStmt production.
	ExitFutureStmt(c *FutureStmtContext)

	// ExitDeclarationStmt is called when exiting the DeclarationStmt production.
	ExitDeclarationStmt(c *DeclarationStmtContext)

	// ExitBranchStmt is called when exiting the BranchStmt production.
	ExitBranchStmt(c *BranchStmtContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitValueAssignment is called when exiting the ValueAssignment production.
	ExitValueAssignment(c *ValueAssignmentContext)

	// ExitVariableDefinition is called when exiting the VariableDefinition production.
	ExitVariableDefinition(c *VariableDefinitionContext)

	// ExitTupleDestructer is called when exiting the TupleDestructer production.
	ExitTupleDestructer(c *TupleDestructerContext)

	// ExitArrayIndexAssignment is called when exiting the ArrayIndexAssignment production.
	ExitArrayIndexAssignment(c *ArrayIndexAssignmentContext)

	// ExitArraySliceAssignment is called when exiting the ArraySliceAssignment production.
	ExitArraySliceAssignment(c *ArraySliceAssignmentContext)

	// ExitRealname is called when exiting the realname production.
	ExitRealname(c *RealnameContext)

	// ExitUnqualifiedName is called when exiting the UnqualifiedName production.
	ExitUnqualifiedName(c *UnqualifiedNameContext)

	// ExitQualifiedName is called when exiting the QualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitAtomicClock is called when exiting the AtomicClock production.
	ExitAtomicClock(c *AtomicClockContext)

	// ExitArrayIndexExpr is called when exiting the ArrayIndexExpr production.
	ExitArrayIndexExpr(c *ArrayIndexExprContext)

	// ExitBranchExpr is called when exiting the BranchExpr production.
	ExitBranchExpr(c *BranchExprContext)

	// ExitLambdaExpr is called when exiting the LambdaExpr production.
	ExitLambdaExpr(c *LambdaExprContext)

	// ExitConstructorExpr is called when exiting the ConstructorExpr production.
	ExitConstructorExpr(c *ConstructorExprContext)

	// ExitTupleExpr is called when exiting the TupleExpr production.
	ExitTupleExpr(c *TupleExprContext)

	// ExitLogicalBinopExpr is called when exiting the LogicalBinopExpr production.
	ExitLogicalBinopExpr(c *LogicalBinopExprContext)

	// ExitConcatExpr is called when exiting the ConcatExpr production.
	ExitConcatExpr(c *ConcatExprContext)

	// ExitSelectorExpr is called when exiting the SelectorExpr production.
	ExitSelectorExpr(c *SelectorExprContext)

	// ExitMulDivModExpr is called when exiting the MulDivModExpr production.
	ExitMulDivModExpr(c *MulDivModExprContext)

	// ExitComplimentExpr is called when exiting the ComplimentExpr production.
	ExitComplimentExpr(c *ComplimentExprContext)

	// ExitArrayLiteralExpr is called when exiting the ArrayLiteralExpr production.
	ExitArrayLiteralExpr(c *ArrayLiteralExprContext)

	// ExitClockToExpr is called when exiting the ClockToExpr production.
	ExitClockToExpr(c *ClockToExprContext)

	// ExitLiteralExpr is called when exiting the LiteralExpr production.
	ExitLiteralExpr(c *LiteralExprContext)

	// ExitVarExpr is called when exiting the VarExpr production.
	ExitVarExpr(c *VarExprContext)

	// ExitParensExpr is called when exiting the ParensExpr production.
	ExitParensExpr(c *ParensExprContext)

	// ExitBitwiseBinopExpr is called when exiting the BitwiseBinopExpr production.
	ExitBitwiseBinopExpr(c *BitwiseBinopExprContext)

	// ExitSliceExpr is called when exiting the SliceExpr production.
	ExitSliceExpr(c *SliceExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitAddSubExpr is called when exiting the AddSubExpr production.
	ExitAddSubExpr(c *AddSubExprContext)

	// ExitShiftExpr is called when exiting the ShiftExpr production.
	ExitShiftExpr(c *ShiftExprContext)

	// ExitTernaryExpr is called when exiting the TernaryExpr production.
	ExitTernaryExpr(c *TernaryExprContext)

	// ExitNewModuleExpr is called when exiting the NewModuleExpr production.
	ExitNewModuleExpr(c *NewModuleExprContext)

	// ExitCallarglist is called when exiting the callarglist production.
	ExitCallarglist(c *CallarglistContext)

	// ExitNamedCallArg is called when exiting the NamedCallArg production.
	ExitNamedCallArg(c *NamedCallArgContext)

	// ExitUnamedCallArg is called when exiting the UnamedCallArg production.
	ExitUnamedCallArg(c *UnamedCallArgContext)

	// ExitParamarglist is called when exiting the paramarglist production.
	ExitParamarglist(c *ParamarglistContext)

	// ExitConcat is called when exiting the concat production.
	ExitConcat(c *ConcatContext)

	// ExitInnerconcat is called when exiting the innerconcat production.
	ExitInnerconcat(c *InnerconcatContext)

	// ExitParameterizedType is called when exiting the ParameterizedType production.
	ExitParameterizedType(c *ParameterizedTypeContext)

	// ExitCompleteType is called when exiting the CompleteType production.
	ExitCompleteType(c *CompleteTypeContext)

	// ExitParamarg is called when exiting the paramarg production.
	ExitParamarg(c *ParamargContext)

	// ExitIfBranch is called when exiting the IfBranch production.
	ExitIfBranch(c *IfBranchContext)

	// ExitMatchBranch is called when exiting the MatchBranch production.
	ExitMatchBranch(c *MatchBranchContext)

	// ExitAtomicPattern is called when exiting the AtomicPattern production.
	ExitAtomicPattern(c *AtomicPatternContext)

	// ExitParamerterizedTypePattern is called when exiting the ParamerterizedTypePattern production.
	ExitParamerterizedTypePattern(c *ParamerterizedTypePatternContext)

	// ExitArrayPattern is called when exiting the ArrayPattern production.
	ExitArrayPattern(c *ArrayPatternContext)

	// ExitLiteralPattern is called when exiting the LiteralPattern production.
	ExitLiteralPattern(c *LiteralPatternContext)

	// ExitStructPattern is called when exiting the StructPattern production.
	ExitStructPattern(c *StructPatternContext)

	// ExitParameterlist is called when exiting the parameterlist production.
	ExitParameterlist(c *ParameterlistContext)

	// ExitTypeParameter is called when exiting the TypeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitValueParameter is called when exiting the ValueParameter production.
	ExitValueParameter(c *ValueParameterContext)

	// ExitSingleReturn is called when exiting the SingleReturn production.
	ExitSingleReturn(c *SingleReturnContext)

	// ExitNamedReturn is called when exiting the NamedReturn production.
	ExitNamedReturn(c *NamedReturnContext)

	// ExitArgumentdef is called when exiting the argumentdef production.
	ExitArgumentdef(c *ArgumentdefContext)

	// ExitArgumentlist is called when exiting the argumentlist production.
	ExitArgumentlist(c *ArgumentlistContext)

	// ExitStructdecl is called when exiting the structdecl production.
	ExitStructdecl(c *StructdeclContext)

	// ExitFielddecl is called when exiting the fielddecl production.
	ExitFielddecl(c *FielddeclContext)

	// ExitFuncdecl is called when exiting the funcdecl production.
	ExitFuncdecl(c *FuncdeclContext)

	// ExitModuledecl is called when exiting the moduledecl production.
	ExitModuledecl(c *ModuledeclContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
