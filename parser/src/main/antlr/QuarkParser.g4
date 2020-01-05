parser grammar QuarkParser;

options {tokenVocab = QuarkLexer;}

program: EOF;

structdecl: KW_STRUCT;
