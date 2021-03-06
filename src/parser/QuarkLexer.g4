lexer grammar QuarkLexer;

fragment UPPERCASE : [A-Z];
fragment LOWERCASE : [a-z];
fragment CHARACTER: [A-Za-z];
fragment NAME_FRAGMENT: [A-Za-z0-9_];
fragment DIGIT : [0-9];
fragment HEX_CHARS: [a-fA-F];
fragment TICK: '\'';
fragment LITERAL_TYPE: ('h' | 'o' | 'd' | 'b');

COLON: ':';
SEMI: ';';
COMMA: ',';
DOUBLE_DOT: '..';
DOT: '.';
UNDERSCORE: '_';
QUESTION_MARK: '?';

LPAREN: '(';
RPAREN: ')';

LBRACE: '[';
RBRACE: ']';

LCURLY: '{';
RCURLY: '}';

PARAM_OPEN: '!(';

OP_ARROW: '=>';
OP_ASSIGN: '=';

OP_ADD: '+';
OP_ADD_ASSIGN: '+=';
OP_SUB: '-';
OP_SUB_ASSIGN: '-=';
OP_MUL: '*';
OP_MUL_ASSIGN: '*=';
OP_DIV: '/';
OP_DIV_ASSIGN: '/=';
OP_MOD: '%';
OP_MOD_ASSIGN: '%=';

OP_BAND: '&';
OP_BAND_ASSIGN: '&=';
OP_BNAND: '~&';
OP_BOR: '|';
OP_BOR_ASSIGN: '|=';
OP_BNOR: '~|';
OP_XOR: '^';
OP_XOR_ASSIGN: '^=';
OP_XNOR: '~^';
OP_LEFT_SHIFT: '<<';
OP_LEFT_SHIFT_ASSIGN: '<<=';
OP_RIGHT_SHIFT: '>>';
OP_RIGHT_SHIFT_ASSIGN: '>>=';
OP_ARITH_LEFT_SHIFT: '<<<';
OP_ARITH_LEFT_SHIFT_ASSIGN: '<<<=';
OP_ARITH_RIGHT_SHFIT: '>>>';
OP_ARITH_RIGHT_SHIFT_ASSIGN: '>>>=';

LANGLE: '<';
RANGLE: '>';

OP_COMPLIMENT: '~';

OP_LAND: 'and';
OP_LOR: 'or';
OP_LNOT: 'not';

OP_IMPLICATION: 'implies';
OP_EQUIVALENCE: 'equates';

OP_LTE: '<=';
OP_GTE: '>=';

OP_EQ: '==';
OP_NEQ: '!=';

//reserved keywords
KW_MODULE: 'module';
KW_DEF: 'def';
KW_IF: 'if';
KW_ELSE: 'else';
KW_ELIF: 'elif';
KW_FOR: 'for';
KW_IN: 'in';
KW_HAS: 'has';
KW_STRUCT: 'struct';
KW_ABSTRACT: 'abstract';
KW_TYPE: 'type';
KW_ENUM: 'enum';
KW_RETURN: 'return';
KW_IMPORT: 'import';
KW_MATCH: 'match';
KW_CASE: 'case';
KW_FUTURE: 'future';
KW_LAMBDA: 'lambda';
KW_CLOCK: 'clock';
KW_RESET: 'reset';
KW_REG: 'reg';
KW_VAR: 'var';
KW_NEW: 'new';
KW_MUT: 'mut';
KW_SIGNAL: 'signal';
KW_TRAIT: 'trait';

//user constructs
COMMENT_START: '//' -> skip, mode(SINGLE_LINE_COMMENT);
BLOCK_COMMENT_START: '/*' -> skip, mode(BLOCK_COMMENT);

DECIMAL: DIGIT+;
INTEGRAL: (DIGIT+)? TICK 's'? LITERAL_TYPE (DIGIT | HEX_CHARS)+;
BIT_VECTOR_PATTERN_TOKEN: (DIGIT+)? TICK 's'? LITERAL_TYPE (DIGIT | HEX_CHARS | '?')+;

REAL_NAME: CHARACTER NAME_FRAGMENT*;
ANNOTATION_START: '@'; //Annotation names are @camelCase
//Type names: CamelCase
//Function names: camelCase
//Param names SNAKE_CASE
//value names: snake_case
//annotation names: camelCase


WS : [ \t\r\n]+ -> skip ;

mode SINGLE_LINE_COMMENT;

NEW_LINE: [\n\r] -> skip, mode(DEFAULT_MODE);
ANYCHAR: ~[\n\r] -> skip;

mode BLOCK_COMMENT;

BLOCK_COMMENT_END: '*/' -> skip, mode(DEFAULT_MODE);
BLOCK_COMMENT_CHAR: . -> skip;
