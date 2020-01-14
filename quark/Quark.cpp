#include "Quark.h"

using namespace Quark;

const std::string QUARK_KEYWORDS [34] = {
    "quark", "import", "module", "parameter", "fixed", "input", "output", 
    "reg", "latch", "future", "function", "return", 
    "bit", "signed", "unsigned", "safe", 
    "overflow", "concat", "clog2", "width", "length", "resize", "cast", 
    "struct", "typedef", "enum", "strict",
    "channel", "interface", 
    "if", "else", "for", "break", "continue"
};

std::set<std::string> Quark::get_quark_keywords() {
    std::set<std::string> keyword_lookup;
    for (size_t i = 0; i < (sizeof(QUARK_KEYWORDS)/sizeof(QUARK_KEYWORDS[0])); i++) {
        keyword_lookup.insert(QUARK_KEYWORDS[i]);
    }
    return keyword_lookup;
}

int Quark::get_operator_precedence(QuarkSymbol s) {
    switch (s) {
        case QuarkSymbol::MULTIPLICATION:
        case QuarkSymbol::DIVISION:
        case QuarkSymbol::MODULUS:
            return 6;
        case QuarkSymbol::ADDITION: 
        case QuarkSymbol::SUBTRACTION:
            return 7;
        case QuarkSymbol::LSHIFT:
        case QuarkSymbol::RSHIFT:
            return 8;
        case QuarkSymbol::LESS:
        case QuarkSymbol::GREATER:
        case QuarkSymbol::LESS_EQUAL:
        case QuarkSymbol::GREATER_EQUAL:
            return 9;
        case QuarkSymbol::EQUAL:
        case QuarkSymbol::NOT_EQUAL:
            return 10;
        case QuarkSymbol::BITWISE_AND: 
        case QuarkSymbol::BITWISE_NAND:
        case QuarkSymbol::BITWISE_OR:
        case QuarkSymbol::BITWISE_NOR:
        case QuarkSymbol::BITWISE_XOR: 
        case QuarkSymbol::BITWISE_XNOR:
            return 11;
        case QuarkSymbol::COND_AND:
        case QuarkSymbol::COND_OR:
        case QuarkSymbol::COND_IMPLICATION:
        case QuarkSymbol::COND_EQUIVALENCE:
            return 12;
        default:
            return -1;
    }
}

bool Quark::is_symbol_assignment(QuarkSymbol s) {
    switch (s) {
        case QuarkSymbol::ASSIGN:
        case QuarkSymbol::ADDITION_ASSIGN:
        case QuarkSymbol::SUBTRACTION_ASSIGN:
        case QuarkSymbol::MULTIPLICATION_ASSIGN:
        case QuarkSymbol::DIVISION_ASSIGN:
        case QuarkSymbol::MODULUS_ASSIGN:
        case QuarkSymbol::BITWISE_AND_ASSIGN:
        case QuarkSymbol::BITWISE_OR_ASSIGN:
        case QuarkSymbol::BITWISE_XOR_ASSIGN:
        case QuarkSymbol::LSHIFT_ASSIGN:
        case QuarkSymbol::RSHIFT_ASSIGN:
        case QuarkSymbol::INCREMENT:
        case QuarkSymbol::DECREMENT:
            return true;
        default:
            return false;
    }

}

bool Quark::is_operator_unary(QuarkSymbol s) {
    switch (s) {
        case QuarkSymbol::ADDITION: // + 
        case QuarkSymbol::SUBTRACTION: // -
        case QuarkSymbol::BITWISE_AND: // &
        case QuarkSymbol::BITWISE_NAND: // ~&
        case QuarkSymbol::BITWISE_OR: // |
        case QuarkSymbol::BITWISE_NOR: // ~|
        case QuarkSymbol::BITWISE_XOR: // ^
        case QuarkSymbol::BITWISE_XNOR: // ~^
        case QuarkSymbol::INVERT: // ~
        case QuarkSymbol::BANG: // !
            return true;
        default:
            return false;
    }
}

char QuarkCharacter::get_escaped_char() {
    switch (c) {
        case 'a' : return '\a';
        case 'b' : return '\b';
        case 'e' : return '\e';
        case 'f' : return '\f';
        case 'n' : return '\n';
        case 'r' : return '\r';
        case 't' : return '\t';
        case 'v' : return '\v';
        case '\'' : return '\'';
        case '\"' : return '\"';
        case '\\' : return '\\';
        default: throw QuarkException("Invalid escaped character in string!");
    }
    return 0;
}

int QuarkCharacter::get_numeric_value(int base) {
    int num;
    if (c >= '0' and c <= '9') {
        num = c - '0';
    } else if (c >= 'a' and c <= 'f') {
        num = c - 'a' + 10;
    } else if (c >= 'A' and c <= 'F') {
        num = c - 'A' + 10;
    } else {
        return -1;
    }

    return (num < base) ? num : -1;
}
