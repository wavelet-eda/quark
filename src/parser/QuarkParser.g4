parser grammar QuarkParser;

options {tokenVocab = QuarkLexer;}

quarkpackage: importdecl* decl* EOF;

decl
    : structdecl
    | funcdecl
    | moduledecl
    | traitdecl
    | enumdecl
    ;

importdecl
    : KW_IMPORT name SEMI # SingleImport
    | KW_IMPORT name DOT OP_MUL SEMI #WildcardImport
    | KW_IMPORT name DOT LPAREN realname (COMMA realname)+ RPAREN SEMI #MultiImport
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
    | KW_REG paramarglist? LPAREN clk=clockexpr (COMMA rst=clockexpr)? RPAREN assignable assignment expr SEMI #RegAssignStmt
    | KW_FUTURE typeexpr realname SEMI #FutureStmt
    | assignable SEMI #DeclarationStmt
    | branch #BranchStmt
    | KW_RETURN expr SEMI #ReturnStmt
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
    | realname #VarExpr
    | expr DOT realname #SelectorExpr
    | LPAREN expr RPAREN #ParensExpr
    | LPAREN expr (COMMA expr)+ RPAREN #TupleExpr
    | typeexpr LCURLY callarglist RCURLY #ConstructorExpr
    | KW_NEW typeexpr LPAREN callarglist RPAREN #NewModuleExpr
    | expr paramarglist? LPAREN callarglist RPAREN #FunctionCall
    | KW_LAMBDA parameterlist? argumentlist OP_ARROW LCURLY block expr? RCURLY #LambdaExpr
    | OP_COMPLIMENT expr # ComplimentExpr
    | OP_LNOT expr # NotExpr
    | concat #ConcatExpr
    | expr op=(OP_MUL | OP_DIV | OP_MOD) expr #MulDivModExpr
    | expr op=(OP_SUB | OP_ADD) expr #AddSubExpr
    | expr op=(OP_LEFT_SHIFT | OP_RIGHT_SHIFT | OP_ARITH_LEFT_SHIFT | OP_ARITH_RIGHT_SHFIT) expr #ShiftExpr
    | expr op=(LANGLE | RANGLE | OP_LTE | OP_GTE) expr #CompareExpr
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

paramarglist: PARAM_OPEN paramarg (COMMA paramarg)* RPAREN;

concat : LCURLY innerconcat (COMMA innerconcat)+ RCURLY;

innerconcat
    : expr LCURLY expr RCURLY
    | expr
    ;

typeexpr
    : name #CompleteType
    | typeexpr paramarglist #ParameterizedType
    | typeexpr LBRACE (expr (COMMA expr)*)? RBRACE #ArrayType
    ;

paramarg
    : KW_TYPE typeexpr
    | expr
    | realname OP_ASSIGN (expr | KW_TYPE typeexpr)
    ;

branch
    : KW_IF expr LCURLY block RCURLY (KW_ELIF expr LCURLY block RCURLY)* (KW_ELSE LCURLY block RCURLY)? #IfBranch
    | KW_MATCH expr LCURLY (KW_CASE pattern LCURLY block RCURLY)+ RCURLY #MatchBranch
    ;

pattern
    : literal #LiteralPattern
    | realname #NamedWildcardPattern
    | QUESTION_MARK #WildcardPattern
    | BIT_VECTOR_PATTERN_TOKEN #BitVectorPattern
    | LPAREN pattern (COMMA pattern)* RPAREN #TuplePattern
    | LBRACE (inner_array_pattern (COMMA inner_array_pattern)*) RBRACE #ArrayPattern
    | name param_pattern? LPAREN (pattern (COMMA pattern)*)? RPAREN #EnumPattern
    ;

param_pattern : PARAM_OPEN pattern (COMMA pattern)* RPAREN;

inner_array_pattern
    : pattern
    | pattern DOUBLE_DOT
    ;


parameterlist: PARAM_OPEN parameterdef (COMMA parameterdef)* RPAREN;

parameterdef
    : KW_TYPE typeexpr #TypeParameter
    | typeexpr realname #ValueParameter
    ;

returnlist
    : typeexpr #SingleReturn
    | LPAREN typeexpr realname (COMMA typeexpr realname)* RPAREN #NamedReturn
    ;

argumentdef: KW_FUTURE? typeexpr realname;

argumentlist: LPAREN (argumentdef (COMMA argumentdef)*)? RPAREN;

enumdecl: annotation* KW_ENUM realname parameterlist? (KW_HAS name (COMMA name)*)? LCURLY (enumconstructordecl SEMI | funcdecl)* RCURLY;

enumconstructordecl: realname (LPAREN enumargdef (COMMA enumargdef)* RPAREN)?;

enumargdef
    : KW_FUTURE? typeexpr
    | KW_FUTURE? typeexpr realname
    ;

structdecl: annotation* KW_STRUCT realname parameterlist? (KW_HAS name (COMMA name)*)? LCURLY (fielddecl | funcdecl)* RCURLY;

fielddecl: annotation* KW_FUTURE? typeexpr realname SEMI;

funcsig: annotation* KW_DEF realname parameterlist? argumentlist (COLON returnlist)?;

funcdecl: funcsig LCURLY block RCURLY;

moduledecl: annotation* KW_MODULE realname parameterlist? argumentlist? (COLON returnlist)? LCURLY block RCURLY;

traitdecl: annotation* KW_TRAIT realname parameterlist? LCURLY (funcsig SEMI)* RCURLY;

annotation: ANNOTATION_START realname;

literal
    : INTEGRAL
    | DECIMAL
    ;
