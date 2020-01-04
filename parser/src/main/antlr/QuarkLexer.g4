lexer grammar QuarkLexer;

fragment UPPERCASE : [A-Z];
fragment LOWERCASE : [a-z];
fragment DIGIT : [0-9];
fragment TICK: '\'';
fragment LITERAL_TYPE: ('h' | 'o' | 'd' | 'b');

COLON: ':';
SEMI: ';';
COMMA: ',';

LPAREN: '(';
RPAREN: ')';

LBRACE: '[';
RBRACE: ']';

LCURLY: '{';
RCURLY: '}';

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

OP_LAND: 'and';
OP_LOR: 'or';
OP_LNOT: 'not';

OP_IMPLICATION: 'implies';
OP_EQUIVALENCE: 'equates';

OP_LT: '<';
OP_GT: '>';
OP_LTE: '<=';
OP_GTE: '>=';

OP_EQ: '==';
OP_NEQ: '!=';

KW_MODULE: 'module';
KW_DEF: 'def';
KW_IF: 'if';
KW_ELSE: 'else';
KW_ELIF: 'elif';
KW_FOR: 'for';
KW_IN: 'in';
KW_IMPLEMENTS: 'implements';
KW_STRUCT: 'struct';
KW_ABSTRACT: 'abstract';
KW_TYPE: 'type';
KW_ENUM: 'enum';
KW_VALUE: 'value';
KW_RETURN: 'return';
KW_IMPORT: 'import';
KW_PUBLIC: 'public';
KW_PRIVATE: 'private';
KW_MATCH: 'match';
KW_CASE: 'case';
KW_FUTURE: 'future';
KW_LOGIC: 'logic';
KW_CLOCK: 'clock';
KW_RESET: 'reset';
KW_REG: 'reg';

COMMENT_START: '//';
BLOCK_COMMENT_START: '/*';
BLOCK_COMMENT_END: '*/';

INTEGRAL: (DIGIT+)? TICK LITERAL_TYPE DIGIT+;

WHITESPACE: '\n\t\r ' -> skip;