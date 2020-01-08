parser grammar QuarkParser;

options {tokenVocab = QuarkLexer;}

program: importdecl* (structdecl | funcdecl | traitimpl | moduledecl)* EOF;

importdecl
    : KW_IMPORT name SEMI # SingleImport
    | KW_IMPORT name DOT UNDERSCORE SEMI #WildcardImport
    ;

block: stmt*;

assignment
    : OP_ASSIGN
    | OP_ADD_ASSIGN
    | OP_SUB_ASSIGN
    | OP_MUL_ASSIGN
    | OP_DIV_ASSIGN
    | OP_MOD_ASSIGN
    | OP_BAND_ASSIGN
    | OP_BOR_ASSIGN
    | OP_XOR_ASSIGN
    | OP_LEFT_SHIFT_ASSIGN
    | OP_RIGHT_SHIFT_ASSIGN
    | OP_ARITH_LEFT_SHIFT_ASSIGN
    | OP_ARITH_RIGHT_SHIFT_ASSIGN
    ;

stmt
    : assignable assignment expr SEMI #AssignStmt
    | KW_REG LPAREN clk=clockexpr (COMMA rst=clockexpr)? RPAREN assignable assignment expr SEMI #RegAssignStmt
    | assignable SEMI #DeclarationStmt
    | branch #BranchStmt
    | KW_RETURN expr SEMI #ReturnStmt
    ;

assignable
    : name LBRACE expr (COMMA expr)* RBRACE #ArrayIndexAssignment
    | name LBRACE msb=expr? COLON lsb=expr? (COLON step=expr?)? BRACE #ArraySliceAssignment
    | name #ValueAssignment
    | KW_MUT? typeexpr realname #VariableDefinition
    | assignable (COMMA assignable)+ #TupleDestructer
    ;

realname : REAL_NAME;

name
    : realname # RealName
    | REAL_NAME (DOT REAL_NAME)+ #QualifiedName
    ;


clockexpr
    : name #AtomicClock
    ;

expr
    : INTEGRAL #LiteralExpr
    | name #VarExpr
    | LPAREN expr RPAREN #ParensExpr
    | LPAREN expr (COMMA expr)+ RPAREN #TupleExpr
    | LCURLY (realname OP_ASSIGN expr COMMA?) RCURLY #ConstructorExpr
    | KW_NEW typeexpr LPAREN (expr (COMMA expr)*)? RPAREN #NewModuleExpr
    | KW_LAMBDA argumentlist OP_ARROW LCURLY stmt* expr? RCURLY #LambdaExpr
    | OP_COMPLIMENT expr # ComplimentExpr
    | OP_LNOT expr # NotExpr
    | LCURLY expr (COMMA expr)+ RCURLY #ConcatExpr
    | LCURLY expr LCURLY expr RCURLY RCURLY #ReplicateExpr
    | expr (OP_MUL | OP_DIV | OP_MOD) expr #MulDivModExpr
    | expr (OP_SUB | OP_ADD) expr #AddSubExpr
    | expr (OP_LEFT_SHIFT | OP_RIGHT_SHIFT | OP_ARITH_LEFT_SHIFT | OP_ARITH_RIGHT_SHFIT) expr #ShiftExpr
    | expr (OP_BAND | OP_BOR | OP_XOR | OP_BNAND | OP_BNOR | OP_XNOR) expr #BitwiseBinopExpr
    | expr (OP_LAND | OP_LOR | OP_IMPLICATION | OP_EQUIVALENCE) expr #LogicalBinopExpr
    | expr KW_IF expr KW_ELSE expr #TernaryExpr
    | branch #BranchExpr
    | arrayslice #SliceExpr
    | LBRACE (expr (COMMA expr)*)? RBRACE #ArrayConstructorExpr
    | KW_SIGNAL LPAREN clockexpr RPAREN #ClockToExpr
    ;

arrayslice
    : name LBRACE msb=expr? COLON lsb=expr? (COLON? step=expr?) RBRACE
    | name LBRACE expr (COMMA expr)* RBRACE
    ;

typeexpr
    : name #CompleteType
    | typeexpr LBRACE (typeexpr | expr) (COMMA (typeexpr | expr))* RBRACE #ParameterizedType
    ;

branch
    : KW_IF expr LCURLY block expr? RCURLY (KW_ELIF expr LCURLY block expr? RCURLY)* (KW_ELSE LCURLY block expr? RCURLY)? #IfBranch
    | KW_MATCH expr LCURLY (KW_CASE pattern LCURLY block expr? RCURLY)+ RCURLY #MatchBranch
    ;

pattern
    : realname # AtomicPattern
    | realname LBRACE pattern (COMMA pattern)* RBRACE #ParamerterizedTypePattern
    | LBRACE (pattern (COMMA pattern)*)? RBRACE #ArrayPattern
    | INTEGRAL #LiteralPattern
    | LCURLY (typeexpr? pattern) (COMMA typeexpr? pattern)* RCURLY #StructPattern
    ;

parameterlist: LBRACE parameterdef (COMMA parameterdef)* RBRACE;

parameterdef
    : KW_TYPE typeexpr (KW_IMPLEMENTS name)? #TypeParameter
    | KW_TYPE typeexpr KW_IMPLEMENTS structdef #AdhocTypeParameter
    | typeexpr? expr #ValueParameter
    ;

returnlist
    : typeexpr #SingleReturn
    | typeexpr realname (COMMA typeexpr realname)* #NamedReturn
    ;

argumentdef: typeexpr? realname;

argumentlist: LPAREN (argumentdef (COMMA argumentdef)*)? RPAREN;

structdecl: annotation* KW_STRUCT realname parameterlist? structdef;

structdef: LCURLY fielddecl* RCURLY;

fielddecl: annotation* realname COLON typeexpr SEMI;

traitimpl: annotation* KW_IMPLEMENTS TYPE_NAME KW_FOR TYPE_NAME LCURLY funcdecl* RCURLY;

funcdecl: annotation* KW_DEF realname parameterlist? argumentlist? (COLON returnlist)? LCURLY block RCURLY;

moduledecl: annotation* KW_MODULE realname parameterlist? argumentlist? (COLON returnlist)? LCURLY innermodule RCURLY;

innermodule: structdecl* block;

annotation: ANNOTATION_NAME (LBRACE name (COMMA name)* RBRACE)?;
