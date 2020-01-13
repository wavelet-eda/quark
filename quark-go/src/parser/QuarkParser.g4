parser grammar QuarkParser;

options {tokenVocab = QuarkLexer;}

quarkpackage: importdecl* decl* EOF;

decl
    : structdecl
    | funcdecl
    | moduledecl
    ;

importdecl
    : KW_IMPORT name SEMI # SingleImport
    | KW_IMPORT name DOT OP_MUL SEMI #WildcardImport
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
    | future SEMI #FutureStmt
    | KW_RETURN expr SEMI #ReturnStmt
    ;

future
    : KW_FUTURE argumentlist RPAREN LBRACE block RBRACE LPAREN callarglist RPAREN
    ;

assignable
    : expr LBRACE expr (COMMA expr)* RBRACE #ArrayIndexAssignment
    | expr LBRACE msb=expr? COLON lsb=expr? (COLON step=expr?)? RBRACE #ArraySliceAssignment
    | name #ValueAssignment
    | KW_MUT? typeexpr realname #VariableDefinition
    | assignable (COMMA assignable)+ #TupleDestructer
    ;

realname : REAL_NAME;

name
    : realname # UnqualifiedName
    | REAL_NAME (DOT REAL_NAME)+ #QualifiedName
    ;


clockexpr
    : name #AtomicClock
    ;

expr
    : literal #LiteralExpr
    | name #VarExpr
    | expr DOT realname #FieldExpr
    | LPAREN expr RPAREN #ParensExpr
    | LPAREN expr (COMMA expr)+ RPAREN #TupleExpr
    | LCURLY callarglist RCURLY #ConstructorExpr
    | KW_NEW typeexpr LPAREN callarglist RPAREN #NewModuleExpr
    | KW_OPEN typeexpr LPAREN callarglist RPAREN #OpenExpr
    | KW_CLOSE expr LPAREN callarglist RPAREN #CloseExpr
    | expr LPAREN callarglist RPAREN #FunctionCall
    | KW_LAMBDA argumentlist OP_ARROW LCURLY block expr? RCURLY #LambdaExpr
    | OP_COMPLIMENT expr # ComplimentExpr
    | OP_LNOT expr # NotExpr
    | concat #ConcatExpr
    | expr op=(OP_MUL | OP_DIV | OP_MOD) expr #MulDivModExpr
    | expr op=(OP_SUB | OP_ADD) expr #AddSubExpr
    | expr op=(OP_LEFT_SHIFT | OP_RIGHT_SHIFT | OP_ARITH_LEFT_SHIFT | OP_ARITH_RIGHT_SHFIT) expr #ShiftExpr
    | expr op=(OP_BAND | OP_BOR | OP_XOR | OP_BNAND | OP_BNOR | OP_XNOR) expr #BitwiseBinopExpr
    | expr op=(OP_LAND | OP_LOR | OP_IMPLICATION | OP_EQUIVALENCE) expr #LogicalBinopExpr
    | expr KW_IF expr KW_ELSE expr #TernaryExpr
    | branch #BranchExpr
    | expr LBRACE msb=expr? COLON lsb=expr? (COLON? step=expr?) RBRACE #SliceExpr
    | expr LBRACE expr (COMMA expr)* RBRACE #ArrayIndexExpr
    | LBRACE (expr (COMMA expr)*)? RBRACE #ArrayLiteralExpr
    | KW_SIGNAL LPAREN clockexpr RPAREN #ClockToExpr
    ;

callarglist: callarg (COMMA callarg)*;

callarg
    : realname OP_ASSIGN expr #NamedCallArg
    | expr #UnamedCallArg
    ;

concat : LCURLY innerconcat (COMMA innerconcat)+ RCURLY;

innerconcat
    : expr LCURLY expr RCURLY
    | expr
    ;

typeexpr
    : name #CompleteType
    | typeexpr LBRACE typeparam (COMMA typeparam)* RBRACE #ParameterizedType
    ;

typeparam
    : typeexpr
    | expr
    ;

branch
    : KW_IF expr LCURLY block RCURLY (KW_ELIF expr LCURLY block RCURLY)* (KW_ELSE LCURLY block RCURLY)? #IfBranch
    | KW_MATCH expr LCURLY (KW_CASE pattern LCURLY block RCURLY)+ RCURLY #MatchBranch
    ;

pattern
    : realname # AtomicPattern
    | realname LBRACE pattern (COMMA pattern)* RBRACE #ParamerterizedTypePattern
    | LBRACE (pattern (COMMA pattern)*)? RBRACE #ArrayPattern
    | literal #LiteralPattern
    | LCURLY (typeexpr? pattern) (COMMA typeexpr? pattern)* RCURLY #StructPattern
    ;


parameterlist: LBRACE parameterdef (COMMA parameterdef)* RBRACE;

parameterdef
    : KW_TYPE typeexpr #TypeParameter
    | typeexpr expr #ValueParameter
    ;

returnlist
    : typeexpr #SingleReturn
    | LPAREN typeexpr realname (COMMA typeexpr realname)* RPAREN #NamedReturn
    ;

argumentdef: typeexpr realname;

argumentlist: LPAREN (argumentdef (COMMA argumentdef)*)? RPAREN;

structdecl: annotation* KW_STRUCT realname parameterlist? (KW_HAS name (COMMA name)*)? structdef;

structdef: LCURLY fielddecl* RCURLY;

fielddecl: annotation* realname COLON typeexpr SEMI;

funcdecl: annotation* KW_DEF realname parameterlist? argumentlist? (COLON returnlist)? LCURLY block RCURLY;

moduledecl: annotation* KW_MODULE realname parameterlist? argumentlist? (COLON returnlist)? LCURLY innermodule RCURLY;

innermodule: structdecl* block;

annotation: ANNOTATION_NAME (LBRACE name (COMMA name)* RBRACE)?;

literal: INTEGRAL;