#pragma once

#include <string>
#include <stdint.h>
#include <gmpxx.h>

#include "Quark.h"

namespace Quark {

    enum class TokenType : uint8_t {
        KEYWORD,
        IDENTIFIER,
        SYMBOL,
        STRING,
        COMMENT,
        LITERAL_INTEGER,
        LITERAL_WILDCARD,
        LITERAL_FLOAT,
        LITERAL_DOUBLE
    };

    // Base class for all tokens
    class Token {
    private:
        size_t line, col;

    protected:
        Token(size_t line, size_t col) : line(line), col(col) {}

    public:
        virtual TokenType get_type() = 0;

        size_t get_line() { return line; }
        size_t get_col() { return col; }
    };

    class TokenKeyword : public Token {
    private:
        size_t keyword_id;

    public:
        TokenKeyword(size_t line, size_t col, size_t keyword_id) : 
                Token(line, col), keyword_id(keyword_id) {}

        TokenType get_type() { return TokenType::KEYWORD; }
        size_t get_keyword_id() { return keyword_id; }
    }

    class TokenIdentifier : public Token {
    private:
        size_t identifier_id;

    public:
        TokenIdentifier(size_t line, size_t col, size_t identifier_id) : 
                Token(line, col), identifier_id(identifier_id) {}

        TokenType get_type() { return TokenType::IDENTIFIER; }
        size_t get_identifier_id() { return identifier_id; }
    }

    class TokenSymbol : public Token {
    private:
        Symbol sym;

    public:
        TokenSymbol(size_t line, size_t col, Symbol sym) :
            Token(line, col), sym(sym) {}

        TokenType get_type() { return TokenType::SYMBOL; }
        Symbol get_symbol() { return sym; }
    }

    class TokenString : public Token {
    private:
        std::string str;

    public:
        TokenString(size_t line, size_t col, std::string str) : 
                Token(line, col), str(str) {}

        TokenType get_type() { return TokenType::STRING; }
        std::string get_string() { return str; }
    }

    class TokenComment : public Token {
    private:
        std::string str;

    public:
        TokenComment(size_t line, size_t col, std::string str) : 
                Token(line, col), str(str) {}

        TokenType get_type() { return TokenType::COMMENT; }
        std::string get_comment() { return str; }
    }

    class TokenLiteralInteger : public Token {
    private:
        bool is_signed_num, is_ambiguous_width_num;
        unsigned int fixed_width;
        mpz_class num;

    public:

        TokenLiteralInteger(size_t line, size_t col, 
                mpz_class num, bool is_signed_num, unsigned int width) :
                Token(line, col), num(num), is_signed_num(is_signed_num),
                width(width), is_ambiguous_width_num(false) {

            if (mpz_sizeinbase(num.get_mpz_t(), 2) > width)
                throw QuarkException("Number larger than defined bit-width!");
        }

        TokenLiteralInteger(size_t line, size_t col, 
                mpz_class num, bool is_signed_num) :
                Token(line, col), num(num), is_signed_num(is_signed_num), 
                width(0), is_ambiguous_width_num(true) {}

        bool is_signed() { return is_signed_num; }
        bool is_ambiguous_width() { return is_ambiguous_width_num; }
        unsigned int get_width() { return is_ambiguous_width_num ? mpz_sizeinbase(num.get_mpz_t(), 2) : fixed_width; }

        mpz_class get_num() { return num; }
    }

    class TokenLiteralWildcard : public Token {
    private:
        bool is_signed_num, is_ambiguous_width_num;
        unsigned int fixed_width;
        mpz_class num, mask;

    public:

        TokenLiteralWildcard(size_t line, size_t col, 
                mpz_class num, mpz_class mask, 
                bool is_signed_num, unsigned int width) :
                Token(line, col), num(num), mask(mask), is_signed_num(is_signed_num),
                width(width), is_ambiguous_width_num(false) {

            if (mpz_sizeinbase(num.get_mpz_t(), 2) > width)
                throw QuarkException("Number larger than defined bit-width!");

            if (mpz_sizeinbase(mask.get_mpz_t(), 2) > width)
                throw QuarkException("Number wildcard mask larger than defined bit-width!");
        }

        TokenLiteralWildcard(size_t line, size_t col, 
                mpz_class num, mpz_class mask, bool is_signed_num) :
                Token(line, col), num(num), mask(mask), is_signed_num(is_signed_num), 
                width(0), is_ambiguous_width_num(true) {}

        bool is_signed() { return is_signed_num; }
        bool is_ambiguous_width() { return is_ambiguous_width_num; }
        unsigned int get_width() { return is_ambiguous_width_num ? mpz_sizeinbase(num.get_mpz_t(), 2) : fixed_width; }
    
        mpz_class get_num() { return num; }
        mpz_class get_mask() { return mask; }
    }

    class TokenLiteralFloat : public Token {
    private:
        float f;

    public:
        TokenLiteralFloat(size_t line, size_t col, float f) : 
                Token(line, col), f(f) {}

        TokenType get_type() { return TokenType::LITERAL_FLOAT; }
        float get_float() { return f; }
    }

    class TokenLiteralDouble : public Token {
    private:
        double d;

    public:
        TokenLiteralDouble(size_t line, size_t col, double d) : 
                Token(line, col), d(d) {}

        TokenType get_type() { return TokenType::LITERAL_DOUBLE; }
        double get_double() { return d; }
    }

        // QuarkNumber num;
        // QuarkNumber num_mask;
        // QuarkSymbol symbol;
        // size_t id;
        // std::string str;
        // float f;
        // double d;

        // QuarkToken(size_t line, size_t col, QuarkNumber num) :
        //         line(line), col(col),
        //         type(QuarkTokenType::NUMBER), num(num) {}

        // QuarkToken(size_t line, size_t col, QuarkNumber num, QuarkNumber nm) :
        //         line(line), col(col), num(num) {
        //     if (nm.num.get_ui() != 0) {
        //         num_mask = nm;
        //         type = QuarkTokenType::NUMBER_WILDCARD;
        //     } else {
        //         num_mask = 0;
        //         type = QuarkTokenType::NUMBER;
        //     }
        // }

        // QuarkToken(size_t line, size_t col, QuarkSymbol symbol) :
        //         line(line), col(col),
        //         type(QuarkTokenType::SYMBOL), symbol(symbol) {}


}
