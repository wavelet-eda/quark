#pragma once

#include <string>
#include <stdint.h>
#include <gmpxx.h>

namespace Quark {

    enum class NodeType : uint8_t {
        QUARK,
        IMPORT,
        MODULE,
        INTERFACE,
        FUNCTION,
        STRUCT,
        TYPEDEF,
        ENUM,

        // Assignment operators are explicitly expanded here and
        //  
        ASSIGNMENT,
        BINARY_OPERATOR,
        UNARY_OPERATOR,
        PARAMETERIZATION, // Anything modifier using []
        IDENTIFIER,

    };

    class Node {

    public:
        virtual NodeType get_type() = 0;
    };

}
