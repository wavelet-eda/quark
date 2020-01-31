// Code generated from src/parser/QuarkLexer.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 101, 664,
	8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6,
	9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4,
	12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17,
	9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9,
	22, 4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27,
	4, 28, 9, 28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4,
	33, 9, 33, 4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38,
	9, 38, 4, 39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9,
	43, 4, 44, 9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48,
	4, 49, 9, 49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4,
	54, 9, 54, 4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59,
	9, 59, 4, 60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9,
	64, 4, 65, 9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69,
	4, 70, 9, 70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4,
	75, 9, 75, 4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80,
	9, 80, 4, 81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9,
	85, 4, 86, 9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90,
	4, 91, 9, 91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4,
	96, 9, 96, 4, 97, 9, 97, 4, 98, 9, 98, 4, 99, 9, 99, 4, 100, 9, 100, 4,
	101, 9, 101, 4, 102, 9, 102, 4, 103, 9, 103, 4, 104, 9, 104, 4, 105, 9,
	105, 4, 106, 9, 106, 4, 107, 9, 107, 4, 108, 9, 108, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3,
	40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44,
	3, 44, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3,
	47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3,
	52, 3, 52, 3, 53, 3, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56,
	3, 56, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3,
	59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60,
	3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 63, 3,
	63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65,
	3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3,
	68, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70,
	3, 70, 3, 71, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 72, 3, 73, 3, 73, 3,
	73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74,
	3, 74, 3, 74, 3, 74, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 76, 3, 76, 3,
	76, 3, 76, 3, 76, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 78, 3, 78,
	3, 78, 3, 78, 3, 78, 3, 78, 3, 78, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3,
	79, 3, 79, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 81, 3, 81,
	3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 82, 3, 82, 3, 82, 3, 82, 3,
	82, 3, 82, 3, 83, 3, 83, 3, 83, 3, 83, 3, 83, 3, 84, 3, 84, 3, 84, 3, 84,
	3, 84, 3, 84, 3, 84, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 3, 86, 3,
	86, 3, 86, 3, 86, 3, 86, 3, 86, 3, 86, 3, 87, 3, 87, 3, 87, 3, 87, 3, 87,
	3, 87, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 89, 3, 89, 3, 89, 3,
	89, 3, 90, 3, 90, 3, 90, 3, 90, 3, 91, 3, 91, 3, 91, 3, 91, 3, 92, 3, 92,
	3, 92, 3, 92, 3, 92, 3, 93, 3, 93, 3, 93, 3, 93, 3, 93, 3, 93, 3, 94, 3,
	94, 3, 94, 3, 94, 3, 95, 3, 95, 3, 95, 3, 95, 3, 95, 3, 95, 3, 95, 3, 96,
	3, 96, 3, 96, 3, 96, 3, 96, 3, 96, 3, 97, 3, 97, 3, 97, 3, 97, 3, 97, 3,
	97, 3, 98, 3, 98, 3, 98, 3, 98, 3, 98, 3, 98, 3, 99, 6, 99, 589, 10, 99,
	13, 99, 14, 99, 590, 3, 100, 6, 100, 594, 10, 100, 13, 100, 14, 100, 595,
	5, 100, 598, 10, 100, 3, 100, 3, 100, 5, 100, 602, 10, 100, 3, 100, 3,
	100, 3, 100, 6, 100, 607, 10, 100, 13, 100, 14, 100, 608, 3, 101, 6, 101,
	612, 10, 101, 13, 101, 14, 101, 613, 5, 101, 616, 10, 101, 3, 101, 3, 101,
	5, 101, 620, 10, 101, 3, 101, 3, 101, 3, 101, 3, 101, 6, 101, 626, 10,
	101, 13, 101, 14, 101, 627, 3, 102, 3, 102, 7, 102, 632, 10, 102, 12, 102,
	14, 102, 635, 11, 102, 3, 103, 3, 103, 3, 104, 6, 104, 640, 10, 104, 13,
	104, 14, 104, 641, 3, 104, 3, 104, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105,
	3, 106, 3, 106, 3, 106, 3, 106, 3, 107, 3, 107, 3, 107, 3, 107, 3, 107,
	3, 107, 3, 108, 3, 108, 3, 108, 3, 108, 2, 2, 109, 5, 2, 7, 2, 9, 2, 11,
	2, 13, 2, 15, 2, 17, 2, 19, 2, 21, 3, 23, 4, 25, 5, 27, 6, 29, 7, 31, 8,
	33, 9, 35, 10, 37, 11, 39, 12, 41, 13, 43, 14, 45, 15, 47, 16, 49, 17,
	51, 18, 53, 19, 55, 20, 57, 21, 59, 22, 61, 23, 63, 24, 65, 25, 67, 26,
	69, 27, 71, 28, 73, 29, 75, 30, 77, 31, 79, 32, 81, 33, 83, 34, 85, 35,
	87, 36, 89, 37, 91, 38, 93, 39, 95, 40, 97, 41, 99, 42, 101, 43, 103, 44,
	105, 45, 107, 46, 109, 47, 111, 48, 113, 49, 115, 50, 117, 51, 119, 52,
	121, 53, 123, 54, 125, 55, 127, 56, 129, 57, 131, 58, 133, 59, 135, 60,
	137, 61, 139, 62, 141, 63, 143, 64, 145, 65, 147, 66, 149, 67, 151, 68,
	153, 69, 155, 70, 157, 71, 159, 72, 161, 73, 163, 74, 165, 75, 167, 76,
	169, 77, 171, 78, 173, 79, 175, 80, 177, 81, 179, 82, 181, 83, 183, 84,
	185, 85, 187, 86, 189, 87, 191, 88, 193, 89, 195, 90, 197, 91, 199, 92,
	201, 93, 203, 94, 205, 95, 207, 96, 209, 97, 211, 98, 213, 99, 215, 100,
	217, 101, 5, 2, 3, 4, 11, 3, 2, 67, 92, 3, 2, 99, 124, 4, 2, 67, 92, 99,
	124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 67, 72,
	99, 104, 6, 2, 100, 100, 102, 102, 106, 106, 113, 113, 5, 2, 11, 12, 15,
	15, 34, 34, 4, 2, 12, 12, 15, 15, 2, 667, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31,
	3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2,
	39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2,
	2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2,
	2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2,
	2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3,
	2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77,
	3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2,
	85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2,
	2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2,
	2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107,
	3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2,
	2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3,
	2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2,
	129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2,
	2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143,
	3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2,
	2, 151, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157, 3,
	2, 2, 2, 2, 159, 3, 2, 2, 2, 2, 161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 2,
	165, 3, 2, 2, 2, 2, 167, 3, 2, 2, 2, 2, 169, 3, 2, 2, 2, 2, 171, 3, 2,
	2, 2, 2, 173, 3, 2, 2, 2, 2, 175, 3, 2, 2, 2, 2, 177, 3, 2, 2, 2, 2, 179,
	3, 2, 2, 2, 2, 181, 3, 2, 2, 2, 2, 183, 3, 2, 2, 2, 2, 185, 3, 2, 2, 2,
	2, 187, 3, 2, 2, 2, 2, 189, 3, 2, 2, 2, 2, 191, 3, 2, 2, 2, 2, 193, 3,
	2, 2, 2, 2, 195, 3, 2, 2, 2, 2, 197, 3, 2, 2, 2, 2, 199, 3, 2, 2, 2, 2,
	201, 3, 2, 2, 2, 2, 203, 3, 2, 2, 2, 2, 205, 3, 2, 2, 2, 2, 207, 3, 2,
	2, 2, 2, 209, 3, 2, 2, 2, 3, 211, 3, 2, 2, 2, 3, 213, 3, 2, 2, 2, 4, 215,
	3, 2, 2, 2, 4, 217, 3, 2, 2, 2, 5, 219, 3, 2, 2, 2, 7, 221, 3, 2, 2, 2,
	9, 223, 3, 2, 2, 2, 11, 225, 3, 2, 2, 2, 13, 227, 3, 2, 2, 2, 15, 229,
	3, 2, 2, 2, 17, 231, 3, 2, 2, 2, 19, 233, 3, 2, 2, 2, 21, 235, 3, 2, 2,
	2, 23, 237, 3, 2, 2, 2, 25, 239, 3, 2, 2, 2, 27, 241, 3, 2, 2, 2, 29, 244,
	3, 2, 2, 2, 31, 246, 3, 2, 2, 2, 33, 248, 3, 2, 2, 2, 35, 250, 3, 2, 2,
	2, 37, 252, 3, 2, 2, 2, 39, 254, 3, 2, 2, 2, 41, 256, 3, 2, 2, 2, 43, 258,
	3, 2, 2, 2, 45, 260, 3, 2, 2, 2, 47, 262, 3, 2, 2, 2, 49, 265, 3, 2, 2,
	2, 51, 268, 3, 2, 2, 2, 53, 270, 3, 2, 2, 2, 55, 272, 3, 2, 2, 2, 57, 275,
	3, 2, 2, 2, 59, 277, 3, 2, 2, 2, 61, 280, 3, 2, 2, 2, 63, 282, 3, 2, 2,
	2, 65, 285, 3, 2, 2, 2, 67, 287, 3, 2, 2, 2, 69, 290, 3, 2, 2, 2, 71, 292,
	3, 2, 2, 2, 73, 295, 3, 2, 2, 2, 75, 297, 3, 2, 2, 2, 77, 300, 3, 2, 2,
	2, 79, 303, 3, 2, 2, 2, 81, 305, 3, 2, 2, 2, 83, 308, 3, 2, 2, 2, 85, 311,
	3, 2, 2, 2, 87, 313, 3, 2, 2, 2, 89, 316, 3, 2, 2, 2, 91, 319, 3, 2, 2,
	2, 93, 322, 3, 2, 2, 2, 95, 326, 3, 2, 2, 2, 97, 329, 3, 2, 2, 2, 99, 333,
	3, 2, 2, 2, 101, 337, 3, 2, 2, 2, 103, 342, 3, 2, 2, 2, 105, 346, 3, 2,
	2, 2, 107, 351, 3, 2, 2, 2, 109, 353, 3, 2, 2, 2, 111, 355, 3, 2, 2, 2,
	113, 357, 3, 2, 2, 2, 115, 361, 3, 2, 2, 2, 117, 364, 3, 2, 2, 2, 119,
	368, 3, 2, 2, 2, 121, 376, 3, 2, 2, 2, 123, 384, 3, 2, 2, 2, 125, 387,
	3, 2, 2, 2, 127, 390, 3, 2, 2, 2, 129, 393, 3, 2, 2, 2, 131, 396, 3, 2,
	2, 2, 133, 403, 3, 2, 2, 2, 135, 407, 3, 2, 2, 2, 137, 410, 3, 2, 2, 2,
	139, 415, 3, 2, 2, 2, 141, 420, 3, 2, 2, 2, 143, 424, 3, 2, 2, 2, 145,
	427, 3, 2, 2, 2, 147, 431, 3, 2, 2, 2, 149, 438, 3, 2, 2, 2, 151, 447,
	3, 2, 2, 2, 153, 452, 3, 2, 2, 2, 155, 457, 3, 2, 2, 2, 157, 463, 3, 2,
	2, 2, 159, 470, 3, 2, 2, 2, 161, 477, 3, 2, 2, 2, 163, 484, 3, 2, 2, 2,
	165, 492, 3, 2, 2, 2, 167, 498, 3, 2, 2, 2, 169, 503, 3, 2, 2, 2, 171,
	510, 3, 2, 2, 2, 173, 516, 3, 2, 2, 2, 175, 523, 3, 2, 2, 2, 177, 529,
	3, 2, 2, 2, 179, 535, 3, 2, 2, 2, 181, 539, 3, 2, 2, 2, 183, 543, 3, 2,
	2, 2, 185, 547, 3, 2, 2, 2, 187, 552, 3, 2, 2, 2, 189, 558, 3, 2, 2, 2,
	191, 562, 3, 2, 2, 2, 193, 569, 3, 2, 2, 2, 195, 575, 3, 2, 2, 2, 197,
	581, 3, 2, 2, 2, 199, 588, 3, 2, 2, 2, 201, 597, 3, 2, 2, 2, 203, 615,
	3, 2, 2, 2, 205, 629, 3, 2, 2, 2, 207, 636, 3, 2, 2, 2, 209, 639, 3, 2,
	2, 2, 211, 645, 3, 2, 2, 2, 213, 650, 3, 2, 2, 2, 215, 654, 3, 2, 2, 2,
	217, 660, 3, 2, 2, 2, 219, 220, 9, 2, 2, 2, 220, 6, 3, 2, 2, 2, 221, 222,
	9, 3, 2, 2, 222, 8, 3, 2, 2, 2, 223, 224, 9, 4, 2, 2, 224, 10, 3, 2, 2,
	2, 225, 226, 9, 5, 2, 2, 226, 12, 3, 2, 2, 2, 227, 228, 9, 6, 2, 2, 228,
	14, 3, 2, 2, 2, 229, 230, 9, 7, 2, 2, 230, 16, 3, 2, 2, 2, 231, 232, 7,
	41, 2, 2, 232, 18, 3, 2, 2, 2, 233, 234, 9, 8, 2, 2, 234, 20, 3, 2, 2,
	2, 235, 236, 7, 60, 2, 2, 236, 22, 3, 2, 2, 2, 237, 238, 7, 61, 2, 2, 238,
	24, 3, 2, 2, 2, 239, 240, 7, 46, 2, 2, 240, 26, 3, 2, 2, 2, 241, 242, 7,
	48, 2, 2, 242, 243, 7, 48, 2, 2, 243, 28, 3, 2, 2, 2, 244, 245, 7, 48,
	2, 2, 245, 30, 3, 2, 2, 2, 246, 247, 7, 97, 2, 2, 247, 32, 3, 2, 2, 2,
	248, 249, 7, 65, 2, 2, 249, 34, 3, 2, 2, 2, 250, 251, 7, 42, 2, 2, 251,
	36, 3, 2, 2, 2, 252, 253, 7, 43, 2, 2, 253, 38, 3, 2, 2, 2, 254, 255, 7,
	93, 2, 2, 255, 40, 3, 2, 2, 2, 256, 257, 7, 95, 2, 2, 257, 42, 3, 2, 2,
	2, 258, 259, 7, 125, 2, 2, 259, 44, 3, 2, 2, 2, 260, 261, 7, 127, 2, 2,
	261, 46, 3, 2, 2, 2, 262, 263, 7, 35, 2, 2, 263, 264, 7, 42, 2, 2, 264,
	48, 3, 2, 2, 2, 265, 266, 7, 63, 2, 2, 266, 267, 7, 64, 2, 2, 267, 50,
	3, 2, 2, 2, 268, 269, 7, 63, 2, 2, 269, 52, 3, 2, 2, 2, 270, 271, 7, 45,
	2, 2, 271, 54, 3, 2, 2, 2, 272, 273, 7, 45, 2, 2, 273, 274, 7, 63, 2, 2,
	274, 56, 3, 2, 2, 2, 275, 276, 7, 47, 2, 2, 276, 58, 3, 2, 2, 2, 277, 278,
	7, 47, 2, 2, 278, 279, 7, 63, 2, 2, 279, 60, 3, 2, 2, 2, 280, 281, 7, 44,
	2, 2, 281, 62, 3, 2, 2, 2, 282, 283, 7, 44, 2, 2, 283, 284, 7, 63, 2, 2,
	284, 64, 3, 2, 2, 2, 285, 286, 7, 49, 2, 2, 286, 66, 3, 2, 2, 2, 287, 288,
	7, 49, 2, 2, 288, 289, 7, 63, 2, 2, 289, 68, 3, 2, 2, 2, 290, 291, 7, 39,
	2, 2, 291, 70, 3, 2, 2, 2, 292, 293, 7, 39, 2, 2, 293, 294, 7, 63, 2, 2,
	294, 72, 3, 2, 2, 2, 295, 296, 7, 40, 2, 2, 296, 74, 3, 2, 2, 2, 297, 298,
	7, 40, 2, 2, 298, 299, 7, 63, 2, 2, 299, 76, 3, 2, 2, 2, 300, 301, 7, 128,
	2, 2, 301, 302, 7, 40, 2, 2, 302, 78, 3, 2, 2, 2, 303, 304, 7, 126, 2,
	2, 304, 80, 3, 2, 2, 2, 305, 306, 7, 126, 2, 2, 306, 307, 7, 63, 2, 2,
	307, 82, 3, 2, 2, 2, 308, 309, 7, 128, 2, 2, 309, 310, 7, 126, 2, 2, 310,
	84, 3, 2, 2, 2, 311, 312, 7, 96, 2, 2, 312, 86, 3, 2, 2, 2, 313, 314, 7,
	96, 2, 2, 314, 315, 7, 63, 2, 2, 315, 88, 3, 2, 2, 2, 316, 317, 7, 128,
	2, 2, 317, 318, 7, 96, 2, 2, 318, 90, 3, 2, 2, 2, 319, 320, 7, 62, 2, 2,
	320, 321, 7, 62, 2, 2, 321, 92, 3, 2, 2, 2, 322, 323, 7, 62, 2, 2, 323,
	324, 7, 62, 2, 2, 324, 325, 7, 63, 2, 2, 325, 94, 3, 2, 2, 2, 326, 327,
	7, 64, 2, 2, 327, 328, 7, 64, 2, 2, 328, 96, 3, 2, 2, 2, 329, 330, 7, 64,
	2, 2, 330, 331, 7, 64, 2, 2, 331, 332, 7, 63, 2, 2, 332, 98, 3, 2, 2, 2,
	333, 334, 7, 62, 2, 2, 334, 335, 7, 62, 2, 2, 335, 336, 7, 62, 2, 2, 336,
	100, 3, 2, 2, 2, 337, 338, 7, 62, 2, 2, 338, 339, 7, 62, 2, 2, 339, 340,
	7, 62, 2, 2, 340, 341, 7, 63, 2, 2, 341, 102, 3, 2, 2, 2, 342, 343, 7,
	64, 2, 2, 343, 344, 7, 64, 2, 2, 344, 345, 7, 64, 2, 2, 345, 104, 3, 2,
	2, 2, 346, 347, 7, 64, 2, 2, 347, 348, 7, 64, 2, 2, 348, 349, 7, 64, 2,
	2, 349, 350, 7, 63, 2, 2, 350, 106, 3, 2, 2, 2, 351, 352, 7, 62, 2, 2,
	352, 108, 3, 2, 2, 2, 353, 354, 7, 64, 2, 2, 354, 110, 3, 2, 2, 2, 355,
	356, 7, 128, 2, 2, 356, 112, 3, 2, 2, 2, 357, 358, 7, 99, 2, 2, 358, 359,
	7, 112, 2, 2, 359, 360, 7, 102, 2, 2, 360, 114, 3, 2, 2, 2, 361, 362, 7,
	113, 2, 2, 362, 363, 7, 116, 2, 2, 363, 116, 3, 2, 2, 2, 364, 365, 7, 112,
	2, 2, 365, 366, 7, 113, 2, 2, 366, 367, 7, 118, 2, 2, 367, 118, 3, 2, 2,
	2, 368, 369, 7, 107, 2, 2, 369, 370, 7, 111, 2, 2, 370, 371, 7, 114, 2,
	2, 371, 372, 7, 110, 2, 2, 372, 373, 7, 107, 2, 2, 373, 374, 7, 103, 2,
	2, 374, 375, 7, 117, 2, 2, 375, 120, 3, 2, 2, 2, 376, 377, 7, 103, 2, 2,
	377, 378, 7, 115, 2, 2, 378, 379, 7, 119, 2, 2, 379, 380, 7, 99, 2, 2,
	380, 381, 7, 118, 2, 2, 381, 382, 7, 103, 2, 2, 382, 383, 7, 117, 2, 2,
	383, 122, 3, 2, 2, 2, 384, 385, 7, 62, 2, 2, 385, 386, 7, 63, 2, 2, 386,
	124, 3, 2, 2, 2, 387, 388, 7, 64, 2, 2, 388, 389, 7, 63, 2, 2, 389, 126,
	3, 2, 2, 2, 390, 391, 7, 63, 2, 2, 391, 392, 7, 63, 2, 2, 392, 128, 3,
	2, 2, 2, 393, 394, 7, 35, 2, 2, 394, 395, 7, 63, 2, 2, 395, 130, 3, 2,
	2, 2, 396, 397, 7, 111, 2, 2, 397, 398, 7, 113, 2, 2, 398, 399, 7, 102,
	2, 2, 399, 400, 7, 119, 2, 2, 400, 401, 7, 110, 2, 2, 401, 402, 7, 103,
	2, 2, 402, 132, 3, 2, 2, 2, 403, 404, 7, 102, 2, 2, 404, 405, 7, 103, 2,
	2, 405, 406, 7, 104, 2, 2, 406, 134, 3, 2, 2, 2, 407, 408, 7, 107, 2, 2,
	408, 409, 7, 104, 2, 2, 409, 136, 3, 2, 2, 2, 410, 411, 7, 103, 2, 2, 411,
	412, 7, 110, 2, 2, 412, 413, 7, 117, 2, 2, 413, 414, 7, 103, 2, 2, 414,
	138, 3, 2, 2, 2, 415, 416, 7, 103, 2, 2, 416, 417, 7, 110, 2, 2, 417, 418,
	7, 107, 2, 2, 418, 419, 7, 104, 2, 2, 419, 140, 3, 2, 2, 2, 420, 421, 7,
	104, 2, 2, 421, 422, 7, 113, 2, 2, 422, 423, 7, 116, 2, 2, 423, 142, 3,
	2, 2, 2, 424, 425, 7, 107, 2, 2, 425, 426, 7, 112, 2, 2, 426, 144, 3, 2,
	2, 2, 427, 428, 7, 106, 2, 2, 428, 429, 7, 99, 2, 2, 429, 430, 7, 117,
	2, 2, 430, 146, 3, 2, 2, 2, 431, 432, 7, 117, 2, 2, 432, 433, 7, 118, 2,
	2, 433, 434, 7, 116, 2, 2, 434, 435, 7, 119, 2, 2, 435, 436, 7, 101, 2,
	2, 436, 437, 7, 118, 2, 2, 437, 148, 3, 2, 2, 2, 438, 439, 7, 99, 2, 2,
	439, 440, 7, 100, 2, 2, 440, 441, 7, 117, 2, 2, 441, 442, 7, 118, 2, 2,
	442, 443, 7, 116, 2, 2, 443, 444, 7, 99, 2, 2, 444, 445, 7, 101, 2, 2,
	445, 446, 7, 118, 2, 2, 446, 150, 3, 2, 2, 2, 447, 448, 7, 118, 2, 2, 448,
	449, 7, 123, 2, 2, 449, 450, 7, 114, 2, 2, 450, 451, 7, 103, 2, 2, 451,
	152, 3, 2, 2, 2, 452, 453, 7, 103, 2, 2, 453, 454, 7, 112, 2, 2, 454, 455,
	7, 119, 2, 2, 455, 456, 7, 111, 2, 2, 456, 154, 3, 2, 2, 2, 457, 458, 7,
	120, 2, 2, 458, 459, 7, 99, 2, 2, 459, 460, 7, 110, 2, 2, 460, 461, 7,
	119, 2, 2, 461, 462, 7, 103, 2, 2, 462, 156, 3, 2, 2, 2, 463, 464, 7, 116,
	2, 2, 464, 465, 7, 103, 2, 2, 465, 466, 7, 118, 2, 2, 466, 467, 7, 119,
	2, 2, 467, 468, 7, 116, 2, 2, 468, 469, 7, 112, 2, 2, 469, 158, 3, 2, 2,
	2, 470, 471, 7, 107, 2, 2, 471, 472, 7, 111, 2, 2, 472, 473, 7, 114, 2,
	2, 473, 474, 7, 113, 2, 2, 474, 475, 7, 116, 2, 2, 475, 476, 7, 118, 2,
	2, 476, 160, 3, 2, 2, 2, 477, 478, 7, 114, 2, 2, 478, 479, 7, 119, 2, 2,
	479, 480, 7, 100, 2, 2, 480, 481, 7, 110, 2, 2, 481, 482, 7, 107, 2, 2,
	482, 483, 7, 101, 2, 2, 483, 162, 3, 2, 2, 2, 484, 485, 7, 114, 2, 2, 485,
	486, 7, 116, 2, 2, 486, 487, 7, 107, 2, 2, 487, 488, 7, 120, 2, 2, 488,
	489, 7, 99, 2, 2, 489, 490, 7, 118, 2, 2, 490, 491, 7, 103, 2, 2, 491,
	164, 3, 2, 2, 2, 492, 493, 7, 111, 2, 2, 493, 494, 7, 99, 2, 2, 494, 495,
	7, 118, 2, 2, 495, 496, 7, 101, 2, 2, 496, 497, 7, 106, 2, 2, 497, 166,
	3, 2, 2, 2, 498, 499, 7, 101, 2, 2, 499, 500, 7, 99, 2, 2, 500, 501, 7,
	117, 2, 2, 501, 502, 7, 103, 2, 2, 502, 168, 3, 2, 2, 2, 503, 504, 7, 104,
	2, 2, 504, 505, 7, 119, 2, 2, 505, 506, 7, 118, 2, 2, 506, 507, 7, 119,
	2, 2, 507, 508, 7, 116, 2, 2, 508, 509, 7, 103, 2, 2, 509, 170, 3, 2, 2,
	2, 510, 511, 7, 110, 2, 2, 511, 512, 7, 113, 2, 2, 512, 513, 7, 105, 2,
	2, 513, 514, 7, 107, 2, 2, 514, 515, 7, 101, 2, 2, 515, 172, 3, 2, 2, 2,
	516, 517, 7, 110, 2, 2, 517, 518, 7, 99, 2, 2, 518, 519, 7, 111, 2, 2,
	519, 520, 7, 100, 2, 2, 520, 521, 7, 102, 2, 2, 521, 522, 7, 99, 2, 2,
	522, 174, 3, 2, 2, 2, 523, 524, 7, 101, 2, 2, 524, 525, 7, 110, 2, 2, 525,
	526, 7, 113, 2, 2, 526, 527, 7, 101, 2, 2, 527, 528, 7, 109, 2, 2, 528,
	176, 3, 2, 2, 2, 529, 530, 7, 116, 2, 2, 530, 531, 7, 103, 2, 2, 531, 532,
	7, 117, 2, 2, 532, 533, 7, 103, 2, 2, 533, 534, 7, 118, 2, 2, 534, 178,
	3, 2, 2, 2, 535, 536, 7, 116, 2, 2, 536, 537, 7, 103, 2, 2, 537, 538, 7,
	105, 2, 2, 538, 180, 3, 2, 2, 2, 539, 540, 7, 120, 2, 2, 540, 541, 7, 99,
	2, 2, 541, 542, 7, 116, 2, 2, 542, 182, 3, 2, 2, 2, 543, 544, 7, 112, 2,
	2, 544, 545, 7, 103, 2, 2, 545, 546, 7, 121, 2, 2, 546, 184, 3, 2, 2, 2,
	547, 548, 7, 113, 2, 2, 548, 549, 7, 114, 2, 2, 549, 550, 7, 103, 2, 2,
	550, 551, 7, 112, 2, 2, 551, 186, 3, 2, 2, 2, 552, 553, 7, 101, 2, 2, 553,
	554, 7, 110, 2, 2, 554, 555, 7, 113, 2, 2, 555, 556, 7, 117, 2, 2, 556,
	557, 7, 103, 2, 2, 557, 188, 3, 2, 2, 2, 558, 559, 7, 111, 2, 2, 559, 560,
	7, 119, 2, 2, 560, 561, 7, 118, 2, 2, 561, 190, 3, 2, 2, 2, 562, 563, 7,
	117, 2, 2, 563, 564, 7, 107, 2, 2, 564, 565, 7, 105, 2, 2, 565, 566, 7,
	112, 2, 2, 566, 567, 7, 99, 2, 2, 567, 568, 7, 110, 2, 2, 568, 192, 3,
	2, 2, 2, 569, 570, 7, 118, 2, 2, 570, 571, 7, 116, 2, 2, 571, 572, 7, 99,
	2, 2, 572, 573, 7, 107, 2, 2, 573, 574, 7, 118, 2, 2, 574, 194, 3, 2, 2,
	2, 575, 576, 7, 49, 2, 2, 576, 577, 7, 49, 2, 2, 577, 578, 3, 2, 2, 2,
	578, 579, 8, 97, 2, 2, 579, 580, 8, 97, 3, 2, 580, 196, 3, 2, 2, 2, 581,
	582, 7, 49, 2, 2, 582, 583, 7, 44, 2, 2, 583, 584, 3, 2, 2, 2, 584, 585,
	8, 98, 2, 2, 585, 586, 8, 98, 4, 2, 586, 198, 3, 2, 2, 2, 587, 589, 5,
	13, 6, 2, 588, 587, 3, 2, 2, 2, 589, 590, 3, 2, 2, 2, 590, 588, 3, 2, 2,
	2, 590, 591, 3, 2, 2, 2, 591, 200, 3, 2, 2, 2, 592, 594, 5, 13, 6, 2, 593,
	592, 3, 2, 2, 2, 594, 595, 3, 2, 2, 2, 595, 593, 3, 2, 2, 2, 595, 596,
	3, 2, 2, 2, 596, 598, 3, 2, 2, 2, 597, 593, 3, 2, 2, 2, 597, 598, 3, 2,
	2, 2, 598, 599, 3, 2, 2, 2, 599, 601, 5, 17, 8, 2, 600, 602, 7, 117, 2,
	2, 601, 600, 3, 2, 2, 2, 601, 602, 3, 2, 2, 2, 602, 603, 3, 2, 2, 2, 603,
	606, 5, 19, 9, 2, 604, 607, 5, 13, 6, 2, 605, 607, 5, 15, 7, 2, 606, 604,
	3, 2, 2, 2, 606, 605, 3, 2, 2, 2, 607, 608, 3, 2, 2, 2, 608, 606, 3, 2,
	2, 2, 608, 609, 3, 2, 2, 2, 609, 202, 3, 2, 2, 2, 610, 612, 5, 13, 6, 2,
	611, 610, 3, 2, 2, 2, 612, 613, 3, 2, 2, 2, 613, 611, 3, 2, 2, 2, 613,
	614, 3, 2, 2, 2, 614, 616, 3, 2, 2, 2, 615, 611, 3, 2, 2, 2, 615, 616,
	3, 2, 2, 2, 616, 617, 3, 2, 2, 2, 617, 619, 5, 17, 8, 2, 618, 620, 7, 117,
	2, 2, 619, 618, 3, 2, 2, 2, 619, 620, 3, 2, 2, 2, 620, 621, 3, 2, 2, 2,
	621, 625, 5, 19, 9, 2, 622, 626, 5, 13, 6, 2, 623, 626, 5, 15, 7, 2, 624,
	626, 7, 65, 2, 2, 625, 622, 3, 2, 2, 2, 625, 623, 3, 2, 2, 2, 625, 624,
	3, 2, 2, 2, 626, 627, 3, 2, 2, 2, 627, 625, 3, 2, 2, 2, 627, 628, 3, 2,
	2, 2, 628, 204, 3, 2, 2, 2, 629, 633, 5, 9, 4, 2, 630, 632, 5, 11, 5, 2,
	631, 630, 3, 2, 2, 2, 632, 635, 3, 2, 2, 2, 633, 631, 3, 2, 2, 2, 633,
	634, 3, 2, 2, 2, 634, 206, 3, 2, 2, 2, 635, 633, 3, 2, 2, 2, 636, 637,
	7, 66, 2, 2, 637, 208, 3, 2, 2, 2, 638, 640, 9, 9, 2, 2, 639, 638, 3, 2,
	2, 2, 640, 641, 3, 2, 2, 2, 641, 639, 3, 2, 2, 2, 641, 642, 3, 2, 2, 2,
	642, 643, 3, 2, 2, 2, 643, 644, 8, 104, 2, 2, 644, 210, 3, 2, 2, 2, 645,
	646, 9, 10, 2, 2, 646, 647, 3, 2, 2, 2, 647, 648, 8, 105, 2, 2, 648, 649,
	8, 105, 5, 2, 649, 212, 3, 2, 2, 2, 650, 651, 10, 10, 2, 2, 651, 652, 3,
	2, 2, 2, 652, 653, 8, 106, 2, 2, 653, 214, 3, 2, 2, 2, 654, 655, 7, 44,
	2, 2, 655, 656, 7, 49, 2, 2, 656, 657, 3, 2, 2, 2, 657, 658, 8, 107, 2,
	2, 658, 659, 8, 107, 5, 2, 659, 216, 3, 2, 2, 2, 660, 661, 11, 2, 2, 2,
	661, 662, 3, 2, 2, 2, 662, 663, 8, 108, 2, 2, 663, 218, 3, 2, 2, 2, 18,
	2, 3, 4, 590, 595, 597, 601, 606, 608, 613, 615, 619, 625, 627, 633, 641,
	6, 8, 2, 2, 4, 3, 2, 4, 4, 2, 4, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "SINGLE_LINE_COMMENT", "BLOCK_COMMENT",
}

var lexerLiteralNames = []string{
	"", "':'", "';'", "','", "'..'", "'.'", "'_'", "'?'", "'('", "')'", "'['",
	"']'", "'{'", "'}'", "'!('", "'=>'", "'='", "'+'", "'+='", "'-'", "'-='",
	"'*'", "'*='", "'/'", "'/='", "'%'", "'%='", "'&'", "'&='", "'~&'", "'|'",
	"'|='", "'~|'", "'^'", "'^='", "'~^'", "'<<'", "'<<='", "'>>'", "'>>='",
	"'<<<'", "'<<<='", "'>>>'", "'>>>='", "'<'", "'>'", "'~'", "'and'", "'or'",
	"'not'", "'implies'", "'equates'", "'<='", "'>='", "'=='", "'!='", "'module'",
	"'def'", "'if'", "'else'", "'elif'", "'for'", "'in'", "'has'", "'struct'",
	"'abstract'", "'type'", "'enum'", "'value'", "'return'", "'import'", "'public'",
	"'private'", "'match'", "'case'", "'future'", "'logic'", "'lambda'", "'clock'",
	"'reset'", "'reg'", "'var'", "'new'", "'open'", "'close'", "'mut'", "'signal'",
	"'trait'", "'//'", "'/*'", "", "", "", "", "'@'", "", "", "", "'*/'",
}

var lexerSymbolicNames = []string{
	"", "COLON", "SEMI", "COMMA", "DOUBLE_DOT", "DOT", "UNDERSCORE", "QUESTION_MARK",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LCURLY", "RCURLY", "PARAM_OPEN",
	"OP_ARROW", "OP_ASSIGN", "OP_ADD", "OP_ADD_ASSIGN", "OP_SUB", "OP_SUB_ASSIGN",
	"OP_MUL", "OP_MUL_ASSIGN", "OP_DIV", "OP_DIV_ASSIGN", "OP_MOD", "OP_MOD_ASSIGN",
	"OP_BAND", "OP_BAND_ASSIGN", "OP_BNAND", "OP_BOR", "OP_BOR_ASSIGN", "OP_BNOR",
	"OP_XOR", "OP_XOR_ASSIGN", "OP_XNOR", "OP_LEFT_SHIFT", "OP_LEFT_SHIFT_ASSIGN",
	"OP_RIGHT_SHIFT", "OP_RIGHT_SHIFT_ASSIGN", "OP_ARITH_LEFT_SHIFT", "OP_ARITH_LEFT_SHIFT_ASSIGN",
	"OP_ARITH_RIGHT_SHFIT", "OP_ARITH_RIGHT_SHIFT_ASSIGN", "LANGLE", "RANGLE",
	"OP_COMPLIMENT", "OP_LAND", "OP_LOR", "OP_LNOT", "OP_IMPLICATION", "OP_EQUIVALENCE",
	"OP_LTE", "OP_GTE", "OP_EQ", "OP_NEQ", "KW_MODULE", "KW_DEF", "KW_IF",
	"KW_ELSE", "KW_ELIF", "KW_FOR", "KW_IN", "KW_HAS", "KW_STRUCT", "KW_ABSTRACT",
	"KW_TYPE", "KW_ENUM", "KW_VALUE", "KW_RETURN", "KW_IMPORT", "KW_PUBLIC",
	"KW_PRIVATE", "KW_MATCH", "KW_CASE", "KW_FUTURE", "KW_LOGIC", "KW_LAMBDA",
	"KW_CLOCK", "KW_RESET", "KW_REG", "KW_VAR", "KW_NEW", "KW_OPEN", "KW_CLOSE",
	"KW_MUT", "KW_SIGNAL", "KW_TRAIT", "COMMENT_START", "BLOCK_COMMENT_START",
	"DECIMAL", "INTEGRAL", "BIT_VECTOR_PATTERN_TOKEN", "REAL_NAME", "ANNOTATION_START",
	"WS", "NEW_LINE", "ANYCHAR", "BLOCK_COMMENT_END", "BLOCK_COMMENT_CHAR",
}

var lexerRuleNames = []string{
	"UPPERCASE", "LOWERCASE", "CHARACTER", "NAME_FRAGMENT", "DIGIT", "HEX_CHARS",
	"TICK", "LITERAL_TYPE", "COLON", "SEMI", "COMMA", "DOUBLE_DOT", "DOT",
	"UNDERSCORE", "QUESTION_MARK", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
	"LCURLY", "RCURLY", "PARAM_OPEN", "OP_ARROW", "OP_ASSIGN", "OP_ADD", "OP_ADD_ASSIGN",
	"OP_SUB", "OP_SUB_ASSIGN", "OP_MUL", "OP_MUL_ASSIGN", "OP_DIV", "OP_DIV_ASSIGN",
	"OP_MOD", "OP_MOD_ASSIGN", "OP_BAND", "OP_BAND_ASSIGN", "OP_BNAND", "OP_BOR",
	"OP_BOR_ASSIGN", "OP_BNOR", "OP_XOR", "OP_XOR_ASSIGN", "OP_XNOR", "OP_LEFT_SHIFT",
	"OP_LEFT_SHIFT_ASSIGN", "OP_RIGHT_SHIFT", "OP_RIGHT_SHIFT_ASSIGN", "OP_ARITH_LEFT_SHIFT",
	"OP_ARITH_LEFT_SHIFT_ASSIGN", "OP_ARITH_RIGHT_SHFIT", "OP_ARITH_RIGHT_SHIFT_ASSIGN",
	"LANGLE", "RANGLE", "OP_COMPLIMENT", "OP_LAND", "OP_LOR", "OP_LNOT", "OP_IMPLICATION",
	"OP_EQUIVALENCE", "OP_LTE", "OP_GTE", "OP_EQ", "OP_NEQ", "KW_MODULE", "KW_DEF",
	"KW_IF", "KW_ELSE", "KW_ELIF", "KW_FOR", "KW_IN", "KW_HAS", "KW_STRUCT",
	"KW_ABSTRACT", "KW_TYPE", "KW_ENUM", "KW_VALUE", "KW_RETURN", "KW_IMPORT",
	"KW_PUBLIC", "KW_PRIVATE", "KW_MATCH", "KW_CASE", "KW_FUTURE", "KW_LOGIC",
	"KW_LAMBDA", "KW_CLOCK", "KW_RESET", "KW_REG", "KW_VAR", "KW_NEW", "KW_OPEN",
	"KW_CLOSE", "KW_MUT", "KW_SIGNAL", "KW_TRAIT", "COMMENT_START", "BLOCK_COMMENT_START",
	"DECIMAL", "INTEGRAL", "BIT_VECTOR_PATTERN_TOKEN", "REAL_NAME", "ANNOTATION_START",
	"WS", "NEW_LINE", "ANYCHAR", "BLOCK_COMMENT_END", "BLOCK_COMMENT_CHAR",
}

type QuarkLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewQuarkLexer(input antlr.CharStream) *QuarkLexer {

	l := new(QuarkLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "QuarkLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QuarkLexer tokens.
const (
	QuarkLexerCOLON                       = 1
	QuarkLexerSEMI                        = 2
	QuarkLexerCOMMA                       = 3
	QuarkLexerDOUBLE_DOT                  = 4
	QuarkLexerDOT                         = 5
	QuarkLexerUNDERSCORE                  = 6
	QuarkLexerQUESTION_MARK               = 7
	QuarkLexerLPAREN                      = 8
	QuarkLexerRPAREN                      = 9
	QuarkLexerLBRACE                      = 10
	QuarkLexerRBRACE                      = 11
	QuarkLexerLCURLY                      = 12
	QuarkLexerRCURLY                      = 13
	QuarkLexerPARAM_OPEN                  = 14
	QuarkLexerOP_ARROW                    = 15
	QuarkLexerOP_ASSIGN                   = 16
	QuarkLexerOP_ADD                      = 17
	QuarkLexerOP_ADD_ASSIGN               = 18
	QuarkLexerOP_SUB                      = 19
	QuarkLexerOP_SUB_ASSIGN               = 20
	QuarkLexerOP_MUL                      = 21
	QuarkLexerOP_MUL_ASSIGN               = 22
	QuarkLexerOP_DIV                      = 23
	QuarkLexerOP_DIV_ASSIGN               = 24
	QuarkLexerOP_MOD                      = 25
	QuarkLexerOP_MOD_ASSIGN               = 26
	QuarkLexerOP_BAND                     = 27
	QuarkLexerOP_BAND_ASSIGN              = 28
	QuarkLexerOP_BNAND                    = 29
	QuarkLexerOP_BOR                      = 30
	QuarkLexerOP_BOR_ASSIGN               = 31
	QuarkLexerOP_BNOR                     = 32
	QuarkLexerOP_XOR                      = 33
	QuarkLexerOP_XOR_ASSIGN               = 34
	QuarkLexerOP_XNOR                     = 35
	QuarkLexerOP_LEFT_SHIFT               = 36
	QuarkLexerOP_LEFT_SHIFT_ASSIGN        = 37
	QuarkLexerOP_RIGHT_SHIFT              = 38
	QuarkLexerOP_RIGHT_SHIFT_ASSIGN       = 39
	QuarkLexerOP_ARITH_LEFT_SHIFT         = 40
	QuarkLexerOP_ARITH_LEFT_SHIFT_ASSIGN  = 41
	QuarkLexerOP_ARITH_RIGHT_SHFIT        = 42
	QuarkLexerOP_ARITH_RIGHT_SHIFT_ASSIGN = 43
	QuarkLexerLANGLE                      = 44
	QuarkLexerRANGLE                      = 45
	QuarkLexerOP_COMPLIMENT               = 46
	QuarkLexerOP_LAND                     = 47
	QuarkLexerOP_LOR                      = 48
	QuarkLexerOP_LNOT                     = 49
	QuarkLexerOP_IMPLICATION              = 50
	QuarkLexerOP_EQUIVALENCE              = 51
	QuarkLexerOP_LTE                      = 52
	QuarkLexerOP_GTE                      = 53
	QuarkLexerOP_EQ                       = 54
	QuarkLexerOP_NEQ                      = 55
	QuarkLexerKW_MODULE                   = 56
	QuarkLexerKW_DEF                      = 57
	QuarkLexerKW_IF                       = 58
	QuarkLexerKW_ELSE                     = 59
	QuarkLexerKW_ELIF                     = 60
	QuarkLexerKW_FOR                      = 61
	QuarkLexerKW_IN                       = 62
	QuarkLexerKW_HAS                      = 63
	QuarkLexerKW_STRUCT                   = 64
	QuarkLexerKW_ABSTRACT                 = 65
	QuarkLexerKW_TYPE                     = 66
	QuarkLexerKW_ENUM                     = 67
	QuarkLexerKW_VALUE                    = 68
	QuarkLexerKW_RETURN                   = 69
	QuarkLexerKW_IMPORT                   = 70
	QuarkLexerKW_PUBLIC                   = 71
	QuarkLexerKW_PRIVATE                  = 72
	QuarkLexerKW_MATCH                    = 73
	QuarkLexerKW_CASE                     = 74
	QuarkLexerKW_FUTURE                   = 75
	QuarkLexerKW_LOGIC                    = 76
	QuarkLexerKW_LAMBDA                   = 77
	QuarkLexerKW_CLOCK                    = 78
	QuarkLexerKW_RESET                    = 79
	QuarkLexerKW_REG                      = 80
	QuarkLexerKW_VAR                      = 81
	QuarkLexerKW_NEW                      = 82
	QuarkLexerKW_OPEN                     = 83
	QuarkLexerKW_CLOSE                    = 84
	QuarkLexerKW_MUT                      = 85
	QuarkLexerKW_SIGNAL                   = 86
	QuarkLexerKW_TRAIT                    = 87
	QuarkLexerCOMMENT_START               = 88
	QuarkLexerBLOCK_COMMENT_START         = 89
	QuarkLexerDECIMAL                     = 90
	QuarkLexerINTEGRAL                    = 91
	QuarkLexerBIT_VECTOR_PATTERN_TOKEN    = 92
	QuarkLexerREAL_NAME                   = 93
	QuarkLexerANNOTATION_START            = 94
	QuarkLexerWS                          = 95
	QuarkLexerNEW_LINE                    = 96
	QuarkLexerANYCHAR                     = 97
	QuarkLexerBLOCK_COMMENT_END           = 98
	QuarkLexerBLOCK_COMMENT_CHAR          = 99
)

// QuarkLexer modes.
const (
	QuarkLexerSINGLE_LINE_COMMENT = iota + 1
	QuarkLexerBLOCK_COMMENT
)
