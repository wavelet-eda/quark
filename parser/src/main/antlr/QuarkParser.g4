parser grammar QuarkParser;

options {tokenVocab = QuarkLexer;}

program: importdecl* (structdecl | funcdecl | traitimpl | moduledecl)* EOF;

importdecl
    : KW_IMPORT name SEMI # SingleImport
    | KW_IMPORT name DOT UNDERSCORE SEMI #WildcardImport
    ;

block: stmt*;

stmt
    : KW_VAR vardef (COMMA vardef)* OP_ASSIGN expr SEMI #AssignStmt
    | KW_REG vardef (COMMA vardef)* OP_ASSIGN expr SEMI #RegAssignStmt
    | KW_RETURN expr SEMI #ReturnStmt
    ;

vardef: VALUE_NAME (COLON type)?;

typename
    : (VALUE_NAME | TYPE_NAME) (DOT (VALUE_NAME | TYPE_NAME))* DOT TYPE_NAME #QualifiedTypeName
    | TYPE_NAME #TypeName
    ;

valuename
    : (VALUE_NAME | TYPE_NAME) (DOT (VALUE_NAME | TYPE_NAME))* DOT VALUE_NAME #QualifiedValueName
    | VALUE_NAME #ValueName
    ;

name
    : typename
    | valuename
    ;

expr
    : INTEGRAL #LiteralExpr
    | valuename #VarExpr
    | LPAREN expr RPAREN #ParensExpr
    | LPAREN expr (COMMA expr)+ RPAREN #TupleExpr
    | LCURLY (VALUE_NAME OP_ASSIGN expr COMMA?) RCURLY #ConstructorExpr
    | KW_NEW type LPAREN (expr (COMMA expr)*)? RPAREN #NewModuleExpr
    | KW_LAMBDA argumentlist OP_ARROW LCURLY stmt* expr? RCURLY #LambdaExpr
    ;

type
    : typename
    | valuename
    | type LBRACE (type | expr) (COMMA (type | expr))* RBRACE
    ;

parameterlist: LBRACE parameterdef (COMMA parameterdef)* RBRACE;

parameterdef
    : KW_TYPE type (KW_IMPLEMENTS typename)?
    | KW_TYPE type KW_IMPLEMENTS structdef
    | KW_VALUE expr (COLON type)?
    ;

argumentdef: VALUE_NAME (COLON type)?;

argumentlist: LPAREN (argumentdef? (COMMA argumentdef)*)? RPAREN;

structdecl: annotation* KW_STRUCT TYPE_NAME parameterlist? structdef;

structdef: LCURLY fielddecl* RCURLY;

fielddecl: annotation* VALUE_NAME COLON type SEMI;

traitimpl: annotation* KW_IMPLEMENTS TYPE_NAME KW_FOR TYPE_NAME LCURLY funcdecl* RCURLY;

funcdecl: annotation* KW_DEF VALUE_NAME parameterlist? argumentlist? (COLON type)? LCURLY block RCURLY;

moduledecl: annotation* KW_MODULE TYPE_NAME parameterlist? argumentlist? (COLON type)? LCURLY innermodule RCURLY;

innermodule: structdecl* block;

annotation: ANNOTATION_NAME (LBRACE valuename (COMMA valuename)* RBRACE)?;
