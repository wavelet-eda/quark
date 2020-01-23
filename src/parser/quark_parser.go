// Code generated from src/parser/QuarkParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // QuarkParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 98, 656,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 7, 2, 66, 10,
	2, 12, 2, 14, 2, 69, 11, 2, 3, 2, 7, 2, 72, 10, 2, 12, 2, 14, 2, 75, 11,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 82, 10, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 6, 4, 101, 10, 4, 13, 4, 14, 4, 102, 3, 4, 3, 4, 3, 4, 5, 4, 108,
	10, 4, 3, 5, 7, 5, 111, 10, 5, 12, 5, 14, 5, 114, 11, 5, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 128, 10,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 149, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 157, 10, 8, 12, 8, 14, 8, 160, 11, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 167, 10, 8, 3, 8, 3, 8, 5, 8, 171,
	10, 8, 3, 8, 3, 8, 5, 8, 175, 10, 8, 5, 8, 177, 10, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 5, 8, 183, 10, 8, 3, 8, 3, 8, 3, 8, 5, 8, 188, 10, 8, 3, 8, 3, 8,
	3, 8, 6, 8, 193, 10, 8, 13, 8, 14, 8, 194, 7, 8, 197, 10, 8, 12, 8, 14,
	8, 200, 11, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 208, 10,
	10, 13, 10, 14, 10, 209, 5, 10, 212, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 227,
	10, 12, 13, 12, 14, 12, 228, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 246,
	10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 253, 10, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	7, 12, 267, 10, 12, 12, 12, 14, 12, 270, 11, 12, 5, 12, 272, 10, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 280, 10, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 5, 12, 308, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 5, 12, 317, 10, 12, 3, 12, 3, 12, 5, 12, 321, 10, 12,
	3, 12, 5, 12, 324, 10, 12, 3, 12, 5, 12, 327, 10, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 7, 12, 335, 10, 12, 12, 12, 14, 12, 338, 11, 12,
	3, 12, 3, 12, 7, 12, 342, 10, 12, 12, 12, 14, 12, 345, 11, 12, 3, 13, 3,
	13, 3, 13, 7, 13, 350, 10, 13, 12, 13, 14, 13, 353, 11, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 5, 14, 360, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7,
	15, 366, 10, 15, 12, 15, 14, 15, 369, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 6, 16, 377, 10, 16, 13, 16, 14, 16, 378, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 389, 10, 17, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 401, 10,
	18, 12, 18, 14, 18, 404, 11, 18, 5, 18, 406, 10, 18, 3, 18, 7, 18, 409,
	10, 18, 12, 18, 14, 18, 412, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 5, 19, 422, 10, 19, 5, 19, 424, 10, 19, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20,
	437, 10, 20, 12, 20, 14, 20, 440, 11, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 5, 20, 447, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 6, 20, 458, 10, 20, 13, 20, 14, 20, 459, 3, 20, 3, 20, 5,
	20, 464, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 472,
	10, 21, 12, 21, 14, 21, 475, 11, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 7, 21, 483, 10, 21, 12, 21, 14, 21, 486, 11, 21, 5, 21, 488, 10,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 494, 10, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 5, 21, 500, 10, 21, 3, 21, 7, 21, 503, 10, 21, 12, 21, 14, 21, 506,
	11, 21, 3, 21, 3, 21, 5, 21, 510, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 7,
	22, 516, 10, 22, 12, 22, 14, 22, 519, 11, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 5, 23, 528, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 7, 24, 538, 10, 24, 12, 24, 14, 24, 541, 11, 24,
	3, 24, 3, 24, 5, 24, 545, 10, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	26, 3, 26, 7, 26, 554, 10, 26, 12, 26, 14, 26, 557, 11, 26, 5, 26, 559,
	10, 26, 3, 26, 3, 26, 3, 27, 7, 27, 564, 10, 27, 12, 27, 14, 27, 567, 11,
	27, 3, 27, 3, 27, 3, 27, 5, 27, 572, 10, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	7, 27, 578, 10, 27, 12, 27, 14, 27, 581, 11, 27, 5, 27, 583, 10, 27, 3,
	27, 3, 27, 7, 27, 587, 10, 27, 12, 27, 14, 27, 590, 11, 27, 3, 27, 3, 27,
	3, 28, 7, 28, 595, 10, 28, 12, 28, 14, 28, 598, 11, 28, 3, 28, 5, 28, 601,
	10, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 7, 29, 608, 10, 29, 12, 29,
	14, 29, 611, 11, 29, 3, 29, 3, 29, 3, 29, 5, 29, 616, 10, 29, 3, 29, 5,
	29, 619, 10, 29, 3, 29, 3, 29, 5, 29, 623, 10, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 30, 7, 30, 630, 10, 30, 12, 30, 14, 30, 633, 11, 30, 3, 30, 3,
	30, 3, 30, 5, 30, 638, 10, 30, 3, 30, 5, 30, 641, 10, 30, 3, 30, 3, 30,
	5, 30, 645, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3,
	32, 3, 32, 3, 32, 2, 5, 14, 22, 34, 33, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 60, 62, 2, 8, 15, 2, 16, 16, 19, 19, 21, 21, 23, 23, 25, 25, 27,
	27, 29, 29, 32, 32, 35, 35, 38, 38, 40, 40, 42, 42, 44, 44, 5, 2, 22, 22,
	24, 24, 26, 26, 4, 2, 18, 18, 20, 20, 6, 2, 37, 37, 39, 39, 41, 41, 43,
	43, 6, 2, 28, 28, 30, 31, 33, 34, 36, 36, 4, 2, 46, 47, 49, 50, 2, 731,
	2, 67, 3, 2, 2, 2, 4, 81, 3, 2, 2, 2, 6, 107, 3, 2, 2, 2, 8, 112, 3, 2,
	2, 2, 10, 115, 3, 2, 2, 2, 12, 148, 3, 2, 2, 2, 14, 187, 3, 2, 2, 2, 16,
	201, 3, 2, 2, 2, 18, 211, 3, 2, 2, 2, 20, 213, 3, 2, 2, 2, 22, 279, 3,
	2, 2, 2, 24, 346, 3, 2, 2, 2, 26, 359, 3, 2, 2, 2, 28, 361, 3, 2, 2, 2,
	30, 372, 3, 2, 2, 2, 32, 388, 3, 2, 2, 2, 34, 390, 3, 2, 2, 2, 36, 423,
	3, 2, 2, 2, 38, 463, 3, 2, 2, 2, 40, 509, 3, 2, 2, 2, 42, 511, 3, 2, 2,
	2, 44, 527, 3, 2, 2, 2, 46, 544, 3, 2, 2, 2, 48, 546, 3, 2, 2, 2, 50, 549,
	3, 2, 2, 2, 52, 565, 3, 2, 2, 2, 54, 596, 3, 2, 2, 2, 56, 609, 3, 2, 2,
	2, 58, 631, 3, 2, 2, 2, 60, 650, 3, 2, 2, 2, 62, 653, 3, 2, 2, 2, 64, 66,
	5, 6, 4, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2,
	67, 68, 3, 2, 2, 2, 68, 73, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 72, 5,
	4, 3, 2, 71, 70, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73,
	74, 3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 77, 7, 2, 2,
	3, 77, 3, 3, 2, 2, 2, 78, 82, 5, 52, 27, 2, 79, 82, 5, 56, 29, 2, 80, 82,
	5, 58, 30, 2, 81, 78, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2,
	2, 82, 5, 3, 2, 2, 2, 83, 84, 7, 69, 2, 2, 84, 85, 5, 18, 10, 2, 85, 86,
	7, 4, 2, 2, 86, 108, 3, 2, 2, 2, 87, 88, 7, 69, 2, 2, 88, 89, 5, 18, 10,
	2, 89, 90, 7, 6, 2, 2, 90, 91, 7, 22, 2, 2, 91, 92, 7, 4, 2, 2, 92, 108,
	3, 2, 2, 2, 93, 94, 7, 69, 2, 2, 94, 95, 5, 18, 10, 2, 95, 96, 7, 6, 2,
	2, 96, 97, 7, 8, 2, 2, 97, 100, 5, 16, 9, 2, 98, 99, 7, 5, 2, 2, 99, 101,
	5, 16, 9, 2, 100, 98, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 100, 3, 2,
	2, 2, 102, 103, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105, 7, 9, 2, 2,
	105, 106, 7, 4, 2, 2, 106, 108, 3, 2, 2, 2, 107, 83, 3, 2, 2, 2, 107, 87,
	3, 2, 2, 2, 107, 93, 3, 2, 2, 2, 108, 7, 3, 2, 2, 2, 109, 111, 5, 12, 7,
	2, 110, 109, 3, 2, 2, 2, 111, 114, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112,
	113, 3, 2, 2, 2, 113, 9, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 116, 9,
	2, 2, 2, 116, 11, 3, 2, 2, 2, 117, 118, 5, 14, 8, 2, 118, 119, 5, 10, 6,
	2, 119, 120, 5, 22, 12, 2, 120, 121, 7, 4, 2, 2, 121, 149, 3, 2, 2, 2,
	122, 123, 7, 79, 2, 2, 123, 124, 7, 8, 2, 2, 124, 127, 5, 20, 11, 2, 125,
	126, 7, 5, 2, 2, 126, 128, 5, 20, 11, 2, 127, 125, 3, 2, 2, 2, 127, 128,
	3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 130, 7, 9, 2, 2, 130, 131, 5, 14,
	8, 2, 131, 132, 5, 10, 6, 2, 132, 133, 5, 22, 12, 2, 133, 134, 7, 4, 2,
	2, 134, 149, 3, 2, 2, 2, 135, 136, 7, 74, 2, 2, 136, 137, 5, 34, 18, 2,
	137, 138, 5, 16, 9, 2, 138, 139, 7, 4, 2, 2, 139, 149, 3, 2, 2, 2, 140,
	141, 5, 14, 8, 2, 141, 142, 7, 4, 2, 2, 142, 149, 3, 2, 2, 2, 143, 149,
	5, 38, 20, 2, 144, 145, 7, 68, 2, 2, 145, 146, 5, 22, 12, 2, 146, 147,
	7, 4, 2, 2, 147, 149, 3, 2, 2, 2, 148, 117, 3, 2, 2, 2, 148, 122, 3, 2,
	2, 2, 148, 135, 3, 2, 2, 2, 148, 140, 3, 2, 2, 2, 148, 143, 3, 2, 2, 2,
	148, 144, 3, 2, 2, 2, 149, 13, 3, 2, 2, 2, 150, 151, 8, 8, 1, 2, 151, 152,
	5, 22, 12, 2, 152, 153, 7, 10, 2, 2, 153, 158, 5, 22, 12, 2, 154, 155,
	7, 5, 2, 2, 155, 157, 5, 22, 12, 2, 156, 154, 3, 2, 2, 2, 157, 160, 3,
	2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 161, 3, 2, 2,
	2, 160, 158, 3, 2, 2, 2, 161, 162, 7, 11, 2, 2, 162, 188, 3, 2, 2, 2, 163,
	164, 5, 22, 12, 2, 164, 166, 7, 10, 2, 2, 165, 167, 5, 22, 12, 2, 166,
	165, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 170,
	7, 3, 2, 2, 169, 171, 5, 22, 12, 2, 170, 169, 3, 2, 2, 2, 170, 171, 3,
	2, 2, 2, 171, 176, 3, 2, 2, 2, 172, 174, 7, 3, 2, 2, 173, 175, 5, 22, 12,
	2, 174, 173, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 177, 3, 2, 2, 2, 176,
	172, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179,
	7, 11, 2, 2, 179, 188, 3, 2, 2, 2, 180, 188, 5, 18, 10, 2, 181, 183, 7,
	84, 2, 2, 182, 181, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 184, 3, 2, 2,
	2, 184, 185, 5, 34, 18, 2, 185, 186, 5, 16, 9, 2, 186, 188, 3, 2, 2, 2,
	187, 150, 3, 2, 2, 2, 187, 163, 3, 2, 2, 2, 187, 180, 3, 2, 2, 2, 187,
	182, 3, 2, 2, 2, 188, 198, 3, 2, 2, 2, 189, 192, 12, 3, 2, 2, 190, 191,
	7, 5, 2, 2, 191, 193, 5, 14, 8, 2, 192, 190, 3, 2, 2, 2, 193, 194, 3, 2,
	2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 197, 3, 2, 2, 2,
	196, 189, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 198,
	199, 3, 2, 2, 2, 199, 15, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 201, 202, 7,
	92, 2, 2, 202, 17, 3, 2, 2, 2, 203, 212, 5, 16, 9, 2, 204, 207, 7, 92,
	2, 2, 205, 206, 7, 6, 2, 2, 206, 208, 7, 92, 2, 2, 207, 205, 3, 2, 2, 2,
	208, 209, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210,
	212, 3, 2, 2, 2, 211, 203, 3, 2, 2, 2, 211, 204, 3, 2, 2, 2, 212, 19, 3,
	2, 2, 2, 213, 214, 5, 18, 10, 2, 214, 21, 3, 2, 2, 2, 215, 216, 8, 12,
	1, 2, 216, 280, 5, 62, 32, 2, 217, 280, 5, 16, 9, 2, 218, 219, 7, 8, 2,
	2, 219, 220, 5, 22, 12, 2, 220, 221, 7, 9, 2, 2, 221, 280, 3, 2, 2, 2,
	222, 223, 7, 8, 2, 2, 223, 226, 5, 22, 12, 2, 224, 225, 7, 5, 2, 2, 225,
	227, 5, 22, 12, 2, 226, 224, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 226,
	3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 231, 7, 9,
	2, 2, 231, 280, 3, 2, 2, 2, 232, 233, 5, 34, 18, 2, 233, 234, 7, 14, 2,
	2, 234, 235, 5, 24, 13, 2, 235, 236, 7, 15, 2, 2, 236, 280, 3, 2, 2, 2,
	237, 238, 7, 81, 2, 2, 238, 239, 5, 34, 18, 2, 239, 240, 7, 8, 2, 2, 240,
	241, 5, 24, 13, 2, 241, 242, 7, 9, 2, 2, 242, 280, 3, 2, 2, 2, 243, 245,
	7, 76, 2, 2, 244, 246, 5, 42, 22, 2, 245, 244, 3, 2, 2, 2, 245, 246, 3,
	2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 248, 5, 50, 26, 2, 248, 249, 7, 17,
	2, 2, 249, 250, 7, 14, 2, 2, 250, 252, 5, 8, 5, 2, 251, 253, 5, 22, 12,
	2, 252, 251, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254,
	255, 7, 15, 2, 2, 255, 280, 3, 2, 2, 2, 256, 257, 7, 45, 2, 2, 257, 280,
	5, 22, 12, 16, 258, 259, 7, 48, 2, 2, 259, 280, 5, 22, 12, 15, 260, 280,
	5, 30, 16, 2, 261, 280, 5, 38, 20, 2, 262, 271, 7, 10, 2, 2, 263, 268,
	5, 22, 12, 2, 264, 265, 7, 5, 2, 2, 265, 267, 5, 22, 12, 2, 266, 264, 3,
	2, 2, 2, 267, 270, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2,
	2, 269, 272, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 271, 263, 3, 2, 2, 2, 271,
	272, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 280, 7, 11, 2, 2, 274, 275,
	7, 85, 2, 2, 275, 276, 7, 8, 2, 2, 276, 277, 5, 20, 11, 2, 277, 278, 7,
	9, 2, 2, 278, 280, 3, 2, 2, 2, 279, 215, 3, 2, 2, 2, 279, 217, 3, 2, 2,
	2, 279, 218, 3, 2, 2, 2, 279, 222, 3, 2, 2, 2, 279, 232, 3, 2, 2, 2, 279,
	237, 3, 2, 2, 2, 279, 243, 3, 2, 2, 2, 279, 256, 3, 2, 2, 2, 279, 258,
	3, 2, 2, 2, 279, 260, 3, 2, 2, 2, 279, 261, 3, 2, 2, 2, 279, 262, 3, 2,
	2, 2, 279, 274, 3, 2, 2, 2, 280, 343, 3, 2, 2, 2, 281, 282, 12, 13, 2,
	2, 282, 283, 9, 3, 2, 2, 283, 342, 5, 22, 12, 14, 284, 285, 12, 12, 2,
	2, 285, 286, 9, 4, 2, 2, 286, 342, 5, 22, 12, 13, 287, 288, 12, 11, 2,
	2, 288, 289, 9, 5, 2, 2, 289, 342, 5, 22, 12, 12, 290, 291, 12, 10, 2,
	2, 291, 292, 9, 6, 2, 2, 292, 342, 5, 22, 12, 11, 293, 294, 12, 9, 2, 2,
	294, 295, 9, 7, 2, 2, 295, 342, 5, 22, 12, 10, 296, 297, 12, 8, 2, 2, 297,
	298, 7, 57, 2, 2, 298, 299, 5, 22, 12, 2, 299, 300, 7, 58, 2, 2, 300, 301,
	5, 22, 12, 9, 301, 342, 3, 2, 2, 2, 302, 303, 12, 23, 2, 2, 303, 304, 7,
	6, 2, 2, 304, 342, 5, 16, 9, 2, 305, 307, 12, 18, 2, 2, 306, 308, 5, 28,
	15, 2, 307, 306, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2,
	309, 310, 7, 8, 2, 2, 310, 311, 5, 24, 13, 2, 311, 312, 7, 9, 2, 2, 312,
	342, 3, 2, 2, 2, 313, 314, 12, 6, 2, 2, 314, 316, 7, 10, 2, 2, 315, 317,
	5, 22, 12, 2, 316, 315, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 318, 3,
	2, 2, 2, 318, 320, 7, 3, 2, 2, 319, 321, 5, 22, 12, 2, 320, 319, 3, 2,
	2, 2, 320, 321, 3, 2, 2, 2, 321, 323, 3, 2, 2, 2, 322, 324, 7, 3, 2, 2,
	323, 322, 3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324, 326, 3, 2, 2, 2, 325,
	327, 5, 22, 12, 2, 326, 325, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327, 328,
	3, 2, 2, 2, 328, 342, 7, 11, 2, 2, 329, 330, 12, 5, 2, 2, 330, 331, 7,
	10, 2, 2, 331, 336, 5, 22, 12, 2, 332, 333, 7, 5, 2, 2, 333, 335, 5, 22,
	12, 2, 334, 332, 3, 2, 2, 2, 335, 338, 3, 2, 2, 2, 336, 334, 3, 2, 2, 2,
	336, 337, 3, 2, 2, 2, 337, 339, 3, 2, 2, 2, 338, 336, 3, 2, 2, 2, 339,
	340, 7, 11, 2, 2, 340, 342, 3, 2, 2, 2, 341, 281, 3, 2, 2, 2, 341, 284,
	3, 2, 2, 2, 341, 287, 3, 2, 2, 2, 341, 290, 3, 2, 2, 2, 341, 293, 3, 2,
	2, 2, 341, 296, 3, 2, 2, 2, 341, 302, 3, 2, 2, 2, 341, 305, 3, 2, 2, 2,
	341, 313, 3, 2, 2, 2, 341, 329, 3, 2, 2, 2, 342, 345, 3, 2, 2, 2, 343,
	341, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 23, 3, 2, 2, 2, 345, 343, 3,
	2, 2, 2, 346, 351, 5, 26, 14, 2, 347, 348, 7, 5, 2, 2, 348, 350, 5, 26,
	14, 2, 349, 347, 3, 2, 2, 2, 350, 353, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2,
	351, 352, 3, 2, 2, 2, 352, 25, 3, 2, 2, 2, 353, 351, 3, 2, 2, 2, 354, 355,
	5, 16, 9, 2, 355, 356, 7, 16, 2, 2, 356, 357, 5, 22, 12, 2, 357, 360, 3,
	2, 2, 2, 358, 360, 5, 22, 12, 2, 359, 354, 3, 2, 2, 2, 359, 358, 3, 2,
	2, 2, 360, 27, 3, 2, 2, 2, 361, 362, 7, 12, 2, 2, 362, 367, 5, 36, 19,
	2, 363, 364, 7, 5, 2, 2, 364, 366, 5, 36, 19, 2, 365, 363, 3, 2, 2, 2,
	366, 369, 3, 2, 2, 2, 367, 365, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368,
	370, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 370, 371, 7, 13, 2, 2, 371, 29,
	3, 2, 2, 2, 372, 373, 7, 14, 2, 2, 373, 376, 5, 32, 17, 2, 374, 375, 7,
	5, 2, 2, 375, 377, 5, 32, 17, 2, 376, 374, 3, 2, 2, 2, 377, 378, 3, 2,
	2, 2, 378, 376, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2,
	380, 381, 7, 15, 2, 2, 381, 31, 3, 2, 2, 2, 382, 383, 5, 22, 12, 2, 383,
	384, 7, 14, 2, 2, 384, 385, 5, 22, 12, 2, 385, 386, 7, 15, 2, 2, 386, 389,
	3, 2, 2, 2, 387, 389, 5, 22, 12, 2, 388, 382, 3, 2, 2, 2, 388, 387, 3,
	2, 2, 2, 389, 33, 3, 2, 2, 2, 390, 391, 8, 18, 1, 2, 391, 392, 5, 18, 10,
	2, 392, 410, 3, 2, 2, 2, 393, 394, 12, 4, 2, 2, 394, 409, 5, 28, 15, 2,
	395, 396, 12, 3, 2, 2, 396, 405, 7, 10, 2, 2, 397, 402, 5, 22, 12, 2, 398,
	399, 7, 5, 2, 2, 399, 401, 5, 22, 12, 2, 400, 398, 3, 2, 2, 2, 401, 404,
	3, 2, 2, 2, 402, 400, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403, 406, 3, 2,
	2, 2, 404, 402, 3, 2, 2, 2, 405, 397, 3, 2, 2, 2, 405, 406, 3, 2, 2, 2,
	406, 407, 3, 2, 2, 2, 407, 409, 7, 11, 2, 2, 408, 393, 3, 2, 2, 2, 408,
	395, 3, 2, 2, 2, 409, 412, 3, 2, 2, 2, 410, 408, 3, 2, 2, 2, 410, 411,
	3, 2, 2, 2, 411, 35, 3, 2, 2, 2, 412, 410, 3, 2, 2, 2, 413, 414, 7, 65,
	2, 2, 414, 424, 5, 34, 18, 2, 415, 424, 5, 22, 12, 2, 416, 417, 5, 16,
	9, 2, 417, 421, 7, 16, 2, 2, 418, 422, 5, 22, 12, 2, 419, 420, 7, 65, 2,
	2, 420, 422, 5, 34, 18, 2, 421, 418, 3, 2, 2, 2, 421, 419, 3, 2, 2, 2,
	422, 424, 3, 2, 2, 2, 423, 413, 3, 2, 2, 2, 423, 415, 3, 2, 2, 2, 423,
	416, 3, 2, 2, 2, 424, 37, 3, 2, 2, 2, 425, 426, 7, 57, 2, 2, 426, 427,
	5, 22, 12, 2, 427, 428, 7, 14, 2, 2, 428, 429, 5, 8, 5, 2, 429, 438, 7,
	15, 2, 2, 430, 431, 7, 59, 2, 2, 431, 432, 5, 22, 12, 2, 432, 433, 7, 14,
	2, 2, 433, 434, 5, 8, 5, 2, 434, 435, 7, 15, 2, 2, 435, 437, 3, 2, 2, 2,
	436, 430, 3, 2, 2, 2, 437, 440, 3, 2, 2, 2, 438, 436, 3, 2, 2, 2, 438,
	439, 3, 2, 2, 2, 439, 446, 3, 2, 2, 2, 440, 438, 3, 2, 2, 2, 441, 442,
	7, 58, 2, 2, 442, 443, 7, 14, 2, 2, 443, 444, 5, 8, 5, 2, 444, 445, 7,
	15, 2, 2, 445, 447, 3, 2, 2, 2, 446, 441, 3, 2, 2, 2, 446, 447, 3, 2, 2,
	2, 447, 464, 3, 2, 2, 2, 448, 449, 7, 72, 2, 2, 449, 450, 5, 22, 12, 2,
	450, 457, 7, 14, 2, 2, 451, 452, 7, 73, 2, 2, 452, 453, 5, 40, 21, 2, 453,
	454, 7, 14, 2, 2, 454, 455, 5, 8, 5, 2, 455, 456, 7, 15, 2, 2, 456, 458,
	3, 2, 2, 2, 457, 451, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 457, 3, 2,
	2, 2, 459, 460, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461, 462, 7, 15, 2, 2,
	462, 464, 3, 2, 2, 2, 463, 425, 3, 2, 2, 2, 463, 448, 3, 2, 2, 2, 464,
	39, 3, 2, 2, 2, 465, 510, 5, 16, 9, 2, 466, 467, 5, 16, 9, 2, 467, 468,
	7, 10, 2, 2, 468, 473, 5, 40, 21, 2, 469, 470, 7, 5, 2, 2, 470, 472, 5,
	40, 21, 2, 471, 469, 3, 2, 2, 2, 472, 475, 3, 2, 2, 2, 473, 471, 3, 2,
	2, 2, 473, 474, 3, 2, 2, 2, 474, 476, 3, 2, 2, 2, 475, 473, 3, 2, 2, 2,
	476, 477, 7, 11, 2, 2, 477, 510, 3, 2, 2, 2, 478, 487, 7, 10, 2, 2, 479,
	484, 5, 40, 21, 2, 480, 481, 7, 5, 2, 2, 481, 483, 5, 40, 21, 2, 482, 480,
	3, 2, 2, 2, 483, 486, 3, 2, 2, 2, 484, 482, 3, 2, 2, 2, 484, 485, 3, 2,
	2, 2, 485, 488, 3, 2, 2, 2, 486, 484, 3, 2, 2, 2, 487, 479, 3, 2, 2, 2,
	487, 488, 3, 2, 2, 2, 488, 489, 3, 2, 2, 2, 489, 510, 7, 11, 2, 2, 490,
	510, 5, 62, 32, 2, 491, 493, 7, 14, 2, 2, 492, 494, 5, 34, 18, 2, 493,
	492, 3, 2, 2, 2, 493, 494, 3, 2, 2, 2, 494, 495, 3, 2, 2, 2, 495, 496,
	5, 40, 21, 2, 496, 504, 3, 2, 2, 2, 497, 499, 7, 5, 2, 2, 498, 500, 5,
	34, 18, 2, 499, 498, 3, 2, 2, 2, 499, 500, 3, 2, 2, 2, 500, 501, 3, 2,
	2, 2, 501, 503, 5, 40, 21, 2, 502, 497, 3, 2, 2, 2, 503, 506, 3, 2, 2,
	2, 504, 502, 3, 2, 2, 2, 504, 505, 3, 2, 2, 2, 505, 507, 3, 2, 2, 2, 506,
	504, 3, 2, 2, 2, 507, 508, 7, 15, 2, 2, 508, 510, 3, 2, 2, 2, 509, 465,
	3, 2, 2, 2, 509, 466, 3, 2, 2, 2, 509, 478, 3, 2, 2, 2, 509, 490, 3, 2,
	2, 2, 509, 491, 3, 2, 2, 2, 510, 41, 3, 2, 2, 2, 511, 512, 7, 12, 2, 2,
	512, 517, 5, 44, 23, 2, 513, 514, 7, 5, 2, 2, 514, 516, 5, 44, 23, 2, 515,
	513, 3, 2, 2, 2, 516, 519, 3, 2, 2, 2, 517, 515, 3, 2, 2, 2, 517, 518,
	3, 2, 2, 2, 518, 520, 3, 2, 2, 2, 519, 517, 3, 2, 2, 2, 520, 521, 7, 13,
	2, 2, 521, 43, 3, 2, 2, 2, 522, 523, 7, 65, 2, 2, 523, 528, 5, 34, 18,
	2, 524, 525, 5, 34, 18, 2, 525, 526, 5, 16, 9, 2, 526, 528, 3, 2, 2, 2,
	527, 522, 3, 2, 2, 2, 527, 524, 3, 2, 2, 2, 528, 45, 3, 2, 2, 2, 529, 545,
	5, 34, 18, 2, 530, 531, 7, 8, 2, 2, 531, 532, 5, 34, 18, 2, 532, 539, 5,
	16, 9, 2, 533, 534, 7, 5, 2, 2, 534, 535, 5, 34, 18, 2, 535, 536, 5, 16,
	9, 2, 536, 538, 3, 2, 2, 2, 537, 533, 3, 2, 2, 2, 538, 541, 3, 2, 2, 2,
	539, 537, 3, 2, 2, 2, 539, 540, 3, 2, 2, 2, 540, 542, 3, 2, 2, 2, 541,
	539, 3, 2, 2, 2, 542, 543, 7, 9, 2, 2, 543, 545, 3, 2, 2, 2, 544, 529,
	3, 2, 2, 2, 544, 530, 3, 2, 2, 2, 545, 47, 3, 2, 2, 2, 546, 547, 5, 34,
	18, 2, 547, 548, 5, 16, 9, 2, 548, 49, 3, 2, 2, 2, 549, 558, 7, 8, 2, 2,
	550, 555, 5, 48, 25, 2, 551, 552, 7, 5, 2, 2, 552, 554, 5, 48, 25, 2, 553,
	551, 3, 2, 2, 2, 554, 557, 3, 2, 2, 2, 555, 553, 3, 2, 2, 2, 555, 556,
	3, 2, 2, 2, 556, 559, 3, 2, 2, 2, 557, 555, 3, 2, 2, 2, 558, 550, 3, 2,
	2, 2, 558, 559, 3, 2, 2, 2, 559, 560, 3, 2, 2, 2, 560, 561, 7, 9, 2, 2,
	561, 51, 3, 2, 2, 2, 562, 564, 5, 60, 31, 2, 563, 562, 3, 2, 2, 2, 564,
	567, 3, 2, 2, 2, 565, 563, 3, 2, 2, 2, 565, 566, 3, 2, 2, 2, 566, 568,
	3, 2, 2, 2, 567, 565, 3, 2, 2, 2, 568, 569, 7, 63, 2, 2, 569, 571, 5, 16,
	9, 2, 570, 572, 5, 42, 22, 2, 571, 570, 3, 2, 2, 2, 571, 572, 3, 2, 2,
	2, 572, 582, 3, 2, 2, 2, 573, 574, 7, 62, 2, 2, 574, 579, 5, 18, 10, 2,
	575, 576, 7, 5, 2, 2, 576, 578, 5, 18, 10, 2, 577, 575, 3, 2, 2, 2, 578,
	581, 3, 2, 2, 2, 579, 577, 3, 2, 2, 2, 579, 580, 3, 2, 2, 2, 580, 583,
	3, 2, 2, 2, 581, 579, 3, 2, 2, 2, 582, 573, 3, 2, 2, 2, 582, 583, 3, 2,
	2, 2, 583, 584, 3, 2, 2, 2, 584, 588, 7, 14, 2, 2, 585, 587, 5, 54, 28,
	2, 586, 585, 3, 2, 2, 2, 587, 590, 3, 2, 2, 2, 588, 586, 3, 2, 2, 2, 588,
	589, 3, 2, 2, 2, 589, 591, 3, 2, 2, 2, 590, 588, 3, 2, 2, 2, 591, 592,
	7, 15, 2, 2, 592, 53, 3, 2, 2, 2, 593, 595, 5, 60, 31, 2, 594, 593, 3,
	2, 2, 2, 595, 598, 3, 2, 2, 2, 596, 594, 3, 2, 2, 2, 596, 597, 3, 2, 2,
	2, 597, 600, 3, 2, 2, 2, 598, 596, 3, 2, 2, 2, 599, 601, 7, 74, 2, 2, 600,
	599, 3, 2, 2, 2, 600, 601, 3, 2, 2, 2, 601, 602, 3, 2, 2, 2, 602, 603,
	5, 34, 18, 2, 603, 604, 5, 16, 9, 2, 604, 605, 7, 4, 2, 2, 605, 55, 3,
	2, 2, 2, 606, 608, 5, 60, 31, 2, 607, 606, 3, 2, 2, 2, 608, 611, 3, 2,
	2, 2, 609, 607, 3, 2, 2, 2, 609, 610, 3, 2, 2, 2, 610, 612, 3, 2, 2, 2,
	611, 609, 3, 2, 2, 2, 612, 613, 7, 56, 2, 2, 613, 615, 5, 16, 9, 2, 614,
	616, 5, 42, 22, 2, 615, 614, 3, 2, 2, 2, 615, 616, 3, 2, 2, 2, 616, 618,
	3, 2, 2, 2, 617, 619, 5, 50, 26, 2, 618, 617, 3, 2, 2, 2, 618, 619, 3,
	2, 2, 2, 619, 622, 3, 2, 2, 2, 620, 621, 7, 3, 2, 2, 621, 623, 5, 46, 24,
	2, 622, 620, 3, 2, 2, 2, 622, 623, 3, 2, 2, 2, 623, 624, 3, 2, 2, 2, 624,
	625, 7, 14, 2, 2, 625, 626, 5, 8, 5, 2, 626, 627, 7, 15, 2, 2, 627, 57,
	3, 2, 2, 2, 628, 630, 5, 60, 31, 2, 629, 628, 3, 2, 2, 2, 630, 633, 3,
	2, 2, 2, 631, 629, 3, 2, 2, 2, 631, 632, 3, 2, 2, 2, 632, 634, 3, 2, 2,
	2, 633, 631, 3, 2, 2, 2, 634, 635, 7, 55, 2, 2, 635, 637, 5, 16, 9, 2,
	636, 638, 5, 42, 22, 2, 637, 636, 3, 2, 2, 2, 637, 638, 3, 2, 2, 2, 638,
	640, 3, 2, 2, 2, 639, 641, 5, 50, 26, 2, 640, 639, 3, 2, 2, 2, 640, 641,
	3, 2, 2, 2, 641, 644, 3, 2, 2, 2, 642, 643, 7, 3, 2, 2, 643, 645, 5, 46,
	24, 2, 644, 642, 3, 2, 2, 2, 644, 645, 3, 2, 2, 2, 645, 646, 3, 2, 2, 2,
	646, 647, 7, 14, 2, 2, 647, 648, 5, 8, 5, 2, 648, 649, 7, 15, 2, 2, 649,
	59, 3, 2, 2, 2, 650, 651, 7, 93, 2, 2, 651, 652, 5, 16, 9, 2, 652, 61,
	3, 2, 2, 2, 653, 654, 7, 91, 2, 2, 654, 63, 3, 2, 2, 2, 78, 67, 73, 81,
	102, 107, 112, 127, 148, 158, 166, 170, 174, 176, 182, 187, 194, 198, 209,
	211, 228, 245, 252, 268, 271, 279, 307, 316, 320, 323, 326, 336, 341, 343,
	351, 359, 367, 378, 388, 402, 405, 408, 410, 421, 423, 438, 446, 459, 463,
	473, 484, 487, 493, 499, 504, 509, 517, 527, 539, 544, 555, 558, 565, 571,
	579, 582, 588, 596, 600, 609, 615, 618, 622, 631, 637, 640, 644,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "';'", "','", "'.'", "'_'", "'('", "')'", "'['", "']'", "'<'",
	"'>'", "'{'", "'}'", "'='", "'=>'", "'+'", "'+='", "'-'", "'-='", "'*'",
	"'*='", "'/'", "'/='", "'%'", "'%='", "'&'", "'&='", "'~&'", "'|'", "'|='",
	"'~|'", "'^'", "'^='", "'~^'", "'<<'", "'<<='", "'>>'", "'>>='", "'<<<'",
	"'<<<='", "'>>>'", "'>>>='", "'~'", "'and'", "'or'", "'not'", "'implies'",
	"'equates'", "'<='", "'>='", "'=='", "'!='", "'module'", "'def'", "'if'",
	"'else'", "'elif'", "'for'", "'in'", "'has'", "'struct'", "'abstract'",
	"'type'", "'enum'", "'value'", "'return'", "'import'", "'public'", "'private'",
	"'match'", "'case'", "'future'", "'logic'", "'lambda'", "'clock'", "'reset'",
	"'reg'", "'var'", "'new'", "'open'", "'close'", "'mut'", "'signal'", "'interface'",
	"'forward'", "'reverse'", "'//'", "'/*'", "", "", "'@'", "", "", "", "'*/'",
}
var symbolicNames = []string{
	"", "COLON", "SEMI", "COMMA", "DOT", "UNDERSCORE", "LPAREN", "RPAREN",
	"LBRACE", "RBRACE", "LANGLE", "RANGLE", "LCURLY", "RCURLY", "OP_ASSIGN",
	"OP_ARROW", "OP_ADD", "OP_ADD_ASSIGN", "OP_SUB", "OP_SUB_ASSIGN", "OP_MUL",
	"OP_MUL_ASSIGN", "OP_DIV", "OP_DIV_ASSIGN", "OP_MOD", "OP_MOD_ASSIGN",
	"OP_BAND", "OP_BAND_ASSIGN", "OP_BNAND", "OP_BOR", "OP_BOR_ASSIGN", "OP_BNOR",
	"OP_XOR", "OP_XOR_ASSIGN", "OP_XNOR", "OP_LEFT_SHIFT", "OP_LEFT_SHIFT_ASSIGN",
	"OP_RIGHT_SHIFT", "OP_RIGHT_SHIFT_ASSIGN", "OP_ARITH_LEFT_SHIFT", "OP_ARITH_LEFT_SHIFT_ASSIGN",
	"OP_ARITH_RIGHT_SHFIT", "OP_ARITH_RIGHT_SHIFT_ASSIGN", "OP_COMPLIMENT",
	"OP_LAND", "OP_LOR", "OP_LNOT", "OP_IMPLICATION", "OP_EQUIVALENCE", "OP_LTE",
	"OP_GTE", "OP_EQ", "OP_NEQ", "KW_MODULE", "KW_DEF", "KW_IF", "KW_ELSE",
	"KW_ELIF", "KW_FOR", "KW_IN", "KW_HAS", "KW_STRUCT", "KW_ABSTRACT", "KW_TYPE",
	"KW_ENUM", "KW_VALUE", "KW_RETURN", "KW_IMPORT", "KW_PUBLIC", "KW_PRIVATE",
	"KW_MATCH", "KW_CASE", "KW_FUTURE", "KW_LOGIC", "KW_LAMBDA", "KW_CLOCK",
	"KW_RESET", "KW_REG", "KW_VAR", "KW_NEW", "KW_OPEN", "KW_CLOSE", "KW_MUT",
	"KW_SIGNAL", "KW_INTERFACE", "KW_FORWARD", "KW_REVERSE", "COMMENT_START",
	"BLOCK_COMMENT_START", "INTEGRAL", "REAL_NAME", "ANNOTATION_START", "WS",
	"NEW_LINE", "ANYCHAR", "BLOCK_COMMENT_END", "BLOCK_COMMENT_CHAR",
}

var ruleNames = []string{
	"quarkpackage", "decl", "importdecl", "block", "assignment", "stmt", "assignable",
	"realname", "name", "clockexpr", "expr", "callarglist", "callarg", "paramarglist",
	"concat", "innerconcat", "typeexpr", "paramarg", "branch", "pattern", "parameterlist",
	"parameterdef", "returnlist", "argumentdef", "argumentlist", "structdecl",
	"fielddecl", "funcdecl", "moduledecl", "annotation", "literal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type QuarkParser struct {
	*antlr.BaseParser
}

func NewQuarkParser(input antlr.TokenStream) *QuarkParser {
	this := new(QuarkParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "QuarkParser.g4"

	return this
}

// QuarkParser tokens.
const (
	QuarkParserEOF                         = antlr.TokenEOF
	QuarkParserCOLON                       = 1
	QuarkParserSEMI                        = 2
	QuarkParserCOMMA                       = 3
	QuarkParserDOT                         = 4
	QuarkParserUNDERSCORE                  = 5
	QuarkParserLPAREN                      = 6
	QuarkParserRPAREN                      = 7
	QuarkParserLBRACE                      = 8
	QuarkParserRBRACE                      = 9
	QuarkParserLANGLE                      = 10
	QuarkParserRANGLE                      = 11
	QuarkParserLCURLY                      = 12
	QuarkParserRCURLY                      = 13
	QuarkParserOP_ASSIGN                   = 14
	QuarkParserOP_ARROW                    = 15
	QuarkParserOP_ADD                      = 16
	QuarkParserOP_ADD_ASSIGN               = 17
	QuarkParserOP_SUB                      = 18
	QuarkParserOP_SUB_ASSIGN               = 19
	QuarkParserOP_MUL                      = 20
	QuarkParserOP_MUL_ASSIGN               = 21
	QuarkParserOP_DIV                      = 22
	QuarkParserOP_DIV_ASSIGN               = 23
	QuarkParserOP_MOD                      = 24
	QuarkParserOP_MOD_ASSIGN               = 25
	QuarkParserOP_BAND                     = 26
	QuarkParserOP_BAND_ASSIGN              = 27
	QuarkParserOP_BNAND                    = 28
	QuarkParserOP_BOR                      = 29
	QuarkParserOP_BOR_ASSIGN               = 30
	QuarkParserOP_BNOR                     = 31
	QuarkParserOP_XOR                      = 32
	QuarkParserOP_XOR_ASSIGN               = 33
	QuarkParserOP_XNOR                     = 34
	QuarkParserOP_LEFT_SHIFT               = 35
	QuarkParserOP_LEFT_SHIFT_ASSIGN        = 36
	QuarkParserOP_RIGHT_SHIFT              = 37
	QuarkParserOP_RIGHT_SHIFT_ASSIGN       = 38
	QuarkParserOP_ARITH_LEFT_SHIFT         = 39
	QuarkParserOP_ARITH_LEFT_SHIFT_ASSIGN  = 40
	QuarkParserOP_ARITH_RIGHT_SHFIT        = 41
	QuarkParserOP_ARITH_RIGHT_SHIFT_ASSIGN = 42
	QuarkParserOP_COMPLIMENT               = 43
	QuarkParserOP_LAND                     = 44
	QuarkParserOP_LOR                      = 45
	QuarkParserOP_LNOT                     = 46
	QuarkParserOP_IMPLICATION              = 47
	QuarkParserOP_EQUIVALENCE              = 48
	QuarkParserOP_LTE                      = 49
	QuarkParserOP_GTE                      = 50
	QuarkParserOP_EQ                       = 51
	QuarkParserOP_NEQ                      = 52
	QuarkParserKW_MODULE                   = 53
	QuarkParserKW_DEF                      = 54
	QuarkParserKW_IF                       = 55
	QuarkParserKW_ELSE                     = 56
	QuarkParserKW_ELIF                     = 57
	QuarkParserKW_FOR                      = 58
	QuarkParserKW_IN                       = 59
	QuarkParserKW_HAS                      = 60
	QuarkParserKW_STRUCT                   = 61
	QuarkParserKW_ABSTRACT                 = 62
	QuarkParserKW_TYPE                     = 63
	QuarkParserKW_ENUM                     = 64
	QuarkParserKW_VALUE                    = 65
	QuarkParserKW_RETURN                   = 66
	QuarkParserKW_IMPORT                   = 67
	QuarkParserKW_PUBLIC                   = 68
	QuarkParserKW_PRIVATE                  = 69
	QuarkParserKW_MATCH                    = 70
	QuarkParserKW_CASE                     = 71
	QuarkParserKW_FUTURE                   = 72
	QuarkParserKW_LOGIC                    = 73
	QuarkParserKW_LAMBDA                   = 74
	QuarkParserKW_CLOCK                    = 75
	QuarkParserKW_RESET                    = 76
	QuarkParserKW_REG                      = 77
	QuarkParserKW_VAR                      = 78
	QuarkParserKW_NEW                      = 79
	QuarkParserKW_OPEN                     = 80
	QuarkParserKW_CLOSE                    = 81
	QuarkParserKW_MUT                      = 82
	QuarkParserKW_SIGNAL                   = 83
	QuarkParserKW_INTERFACE                = 84
	QuarkParserKW_FORWARD                  = 85
	QuarkParserKW_REVERSE                  = 86
	QuarkParserCOMMENT_START               = 87
	QuarkParserBLOCK_COMMENT_START         = 88
	QuarkParserINTEGRAL                    = 89
	QuarkParserREAL_NAME                   = 90
	QuarkParserANNOTATION_START            = 91
	QuarkParserWS                          = 92
	QuarkParserNEW_LINE                    = 93
	QuarkParserANYCHAR                     = 94
	QuarkParserBLOCK_COMMENT_END           = 95
	QuarkParserBLOCK_COMMENT_CHAR          = 96
)

// QuarkParser rules.
const (
	QuarkParserRULE_quarkpackage  = 0
	QuarkParserRULE_decl          = 1
	QuarkParserRULE_importdecl    = 2
	QuarkParserRULE_block         = 3
	QuarkParserRULE_assignment    = 4
	QuarkParserRULE_stmt          = 5
	QuarkParserRULE_assignable    = 6
	QuarkParserRULE_realname      = 7
	QuarkParserRULE_name          = 8
	QuarkParserRULE_clockexpr     = 9
	QuarkParserRULE_expr          = 10
	QuarkParserRULE_callarglist   = 11
	QuarkParserRULE_callarg       = 12
	QuarkParserRULE_paramarglist  = 13
	QuarkParserRULE_concat        = 14
	QuarkParserRULE_innerconcat   = 15
	QuarkParserRULE_typeexpr      = 16
	QuarkParserRULE_paramarg      = 17
	QuarkParserRULE_branch        = 18
	QuarkParserRULE_pattern       = 19
	QuarkParserRULE_parameterlist = 20
	QuarkParserRULE_parameterdef  = 21
	QuarkParserRULE_returnlist    = 22
	QuarkParserRULE_argumentdef   = 23
	QuarkParserRULE_argumentlist  = 24
	QuarkParserRULE_structdecl    = 25
	QuarkParserRULE_fielddecl     = 26
	QuarkParserRULE_funcdecl      = 27
	QuarkParserRULE_moduledecl    = 28
	QuarkParserRULE_annotation    = 29
	QuarkParserRULE_literal       = 30
)

// IQuarkpackageContext is an interface to support dynamic dispatch.
type IQuarkpackageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuarkpackageContext differentiates from other interfaces.
	IsQuarkpackageContext()
}

type QuarkpackageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuarkpackageContext() *QuarkpackageContext {
	var p = new(QuarkpackageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_quarkpackage
	return p
}

func (*QuarkpackageContext) IsQuarkpackageContext() {}

func NewQuarkpackageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuarkpackageContext {
	var p = new(QuarkpackageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_quarkpackage

	return p
}

func (s *QuarkpackageContext) GetParser() antlr.Parser { return s.parser }

func (s *QuarkpackageContext) EOF() antlr.TerminalNode {
	return s.GetToken(QuarkParserEOF, 0)
}

func (s *QuarkpackageContext) AllImportdecl() []IImportdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportdeclContext)(nil)).Elem())
	var tst = make([]IImportdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportdeclContext)
		}
	}

	return tst
}

func (s *QuarkpackageContext) Importdecl(i int) IImportdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportdeclContext)
}

func (s *QuarkpackageContext) AllDecl() []IDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclContext)(nil)).Elem())
	var tst = make([]IDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclContext)
		}
	}

	return tst
}

func (s *QuarkpackageContext) Decl(i int) IDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *QuarkpackageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuarkpackageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuarkpackageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterQuarkpackage(s)
	}
}

func (s *QuarkpackageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitQuarkpackage(s)
	}
}

func (s *QuarkpackageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitQuarkpackage(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Quarkpackage() (localctx IQuarkpackageContext) {
	localctx = NewQuarkpackageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QuarkParserRULE_quarkpackage)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserKW_IMPORT {
		{
			p.SetState(62)
			p.Importdecl()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(QuarkParserKW_MODULE-53))|(1<<(QuarkParserKW_DEF-53))|(1<<(QuarkParserKW_STRUCT-53)))) != 0) || _la == QuarkParserANNOTATION_START {
		{
			p.SetState(68)
			p.Decl()
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(74)
		p.Match(QuarkParserEOF)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) Structdecl() IStructdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructdeclContext)
}

func (s *DeclContext) Funcdecl() IFuncdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncdeclContext)
}

func (s *DeclContext) Moduledecl() IModuledeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuledeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuledeclContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (s *DeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Decl() (localctx IDeclContext) {
	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QuarkParserRULE_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Structdecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Funcdecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.Moduledecl()
		}

	}

	return localctx
}

// IImportdeclContext is an interface to support dynamic dispatch.
type IImportdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportdeclContext differentiates from other interfaces.
	IsImportdeclContext()
}

type ImportdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportdeclContext() *ImportdeclContext {
	var p = new(ImportdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_importdecl
	return p
}

func (*ImportdeclContext) IsImportdeclContext() {}

func NewImportdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportdeclContext {
	var p = new(ImportdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_importdecl

	return p
}

func (s *ImportdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportdeclContext) CopyFrom(ctx *ImportdeclContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ImportdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultiImportContext struct {
	*ImportdeclContext
}

func NewMultiImportContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiImportContext {
	var p = new(MultiImportContext)

	p.ImportdeclContext = NewEmptyImportdeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ImportdeclContext))

	return p
}

func (s *MultiImportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiImportContext) KW_IMPORT() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_IMPORT, 0)
}

func (s *MultiImportContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *MultiImportContext) DOT() antlr.TerminalNode {
	return s.GetToken(QuarkParserDOT, 0)
}

func (s *MultiImportContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserLPAREN, 0)
}

func (s *MultiImportContext) AllRealname() []IRealnameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRealnameContext)(nil)).Elem())
	var tst = make([]IRealnameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRealnameContext)
		}
	}

	return tst
}

func (s *MultiImportContext) Realname(i int) IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *MultiImportContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
}

func (s *MultiImportContext) SEMI() antlr.TerminalNode {
	return s.GetToken(QuarkParserSEMI, 0)
}

func (s *MultiImportContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *MultiImportContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *MultiImportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterMultiImport(s)
	}
}

func (s *MultiImportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitMultiImport(s)
	}
}

func (s *MultiImportContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitMultiImport(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleImportContext struct {
	*ImportdeclContext
}

func NewSingleImportContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleImportContext {
	var p = new(SingleImportContext)

	p.ImportdeclContext = NewEmptyImportdeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ImportdeclContext))

	return p
}

func (s *SingleImportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleImportContext) KW_IMPORT() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_IMPORT, 0)
}

func (s *SingleImportContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *SingleImportContext) SEMI() antlr.TerminalNode {
	return s.GetToken(QuarkParserSEMI, 0)
}

func (s *SingleImportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterSingleImport(s)
	}
}

func (s *SingleImportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitSingleImport(s)
	}
}

func (s *SingleImportContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitSingleImport(s)

	default:
		return t.VisitChildren(s)
	}
}

type WildcardImportContext struct {
	*ImportdeclContext
}

func NewWildcardImportContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WildcardImportContext {
	var p = new(WildcardImportContext)

	p.ImportdeclContext = NewEmptyImportdeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ImportdeclContext))

	return p
}

func (s *WildcardImportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WildcardImportContext) KW_IMPORT() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_IMPORT, 0)
}

func (s *WildcardImportContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *WildcardImportContext) DOT() antlr.TerminalNode {
	return s.GetToken(QuarkParserDOT, 0)
}

func (s *WildcardImportContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_MUL, 0)
}

func (s *WildcardImportContext) SEMI() antlr.TerminalNode {
	return s.GetToken(QuarkParserSEMI, 0)
}

func (s *WildcardImportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterWildcardImport(s)
	}
}

func (s *WildcardImportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitWildcardImport(s)
	}
}

func (s *WildcardImportContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitWildcardImport(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Importdecl() (localctx IImportdeclContext) {
	localctx = NewImportdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QuarkParserRULE_importdecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.Match(QuarkParserKW_IMPORT)
		}
		{
			p.SetState(82)
			p.Name()
		}
		{
			p.SetState(83)
			p.Match(QuarkParserSEMI)
		}

	case 2:
		localctx = NewWildcardImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(QuarkParserKW_IMPORT)
		}
		{
			p.SetState(86)
			p.Name()
		}
		{
			p.SetState(87)
			p.Match(QuarkParserDOT)
		}
		{
			p.SetState(88)
			p.Match(QuarkParserOP_MUL)
		}
		{
			p.SetState(89)
			p.Match(QuarkParserSEMI)
		}

	case 3:
		localctx = NewMultiImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.Match(QuarkParserKW_IMPORT)
		}
		{
			p.SetState(92)
			p.Name()
		}
		{
			p.SetState(93)
			p.Match(QuarkParserDOT)
		}
		{
			p.SetState(94)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(95)
			p.Realname()
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QuarkParserCOMMA {
			{
				p.SetState(96)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(97)
				p.Realname()
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(102)
			p.Match(QuarkParserRPAREN)
		}
		{
			p.SetState(103)
			p.Match(QuarkParserSEMI)
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QuarkParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(107)
				p.Stmt()
			}

		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_ASSIGN, 0)
}

func (s *AssignmentContext) OP_ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_ADD_ASSIGN, 0)
}

func (s *AssignmentContext) OP_SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_SUB_ASSIGN, 0)
}

func (s *AssignmentContext) OP_MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_MUL_ASSIGN, 0)
}

func (s *AssignmentContext) OP_DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_DIV_ASSIGN, 0)
}

func (s *AssignmentContext) OP_MOD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_MOD_ASSIGN, 0)
}

func (s *AssignmentContext) OP_BAND_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_BAND_ASSIGN, 0)
}

func (s *AssignmentContext) OP_BOR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_BOR_ASSIGN, 0)
}

func (s *AssignmentContext) OP_XOR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_XOR_ASSIGN, 0)
}

func (s *AssignmentContext) OP_LEFT_SHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_LEFT_SHIFT_ASSIGN, 0)
}

func (s *AssignmentContext) OP_RIGHT_SHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_RIGHT_SHIFT_ASSIGN, 0)
}

func (s *AssignmentContext) OP_ARITH_LEFT_SHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_ARITH_LEFT_SHIFT_ASSIGN, 0)
}

func (s *AssignmentContext) OP_ARITH_RIGHT_SHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_ARITH_RIGHT_SHIFT_ASSIGN, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QuarkParserRULE_assignment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-14)&-(0x1f+1)) == 0 && ((1<<uint((_la-14)))&((1<<(QuarkParserOP_ASSIGN-14))|(1<<(QuarkParserOP_ADD_ASSIGN-14))|(1<<(QuarkParserOP_SUB_ASSIGN-14))|(1<<(QuarkParserOP_MUL_ASSIGN-14))|(1<<(QuarkParserOP_DIV_ASSIGN-14))|(1<<(QuarkParserOP_MOD_ASSIGN-14))|(1<<(QuarkParserOP_BAND_ASSIGN-14))|(1<<(QuarkParserOP_BOR_ASSIGN-14))|(1<<(QuarkParserOP_XOR_ASSIGN-14))|(1<<(QuarkParserOP_LEFT_SHIFT_ASSIGN-14))|(1<<(QuarkParserOP_RIGHT_SHIFT_ASSIGN-14))|(1<<(QuarkParserOP_ARITH_LEFT_SHIFT_ASSIGN-14))|(1<<(QuarkParserOP_ARITH_RIGHT_SHIFT_ASSIGN-14)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyFrom(ctx *StmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BranchStmtContext struct {
	*StmtContext
}

func NewBranchStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BranchStmtContext {
	var p = new(BranchStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *BranchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchStmtContext) Branch() IBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *BranchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterBranchStmt(s)
	}
}

func (s *BranchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitBranchStmt(s)
	}
}

func (s *BranchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitBranchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclarationStmtContext struct {
	*StmtContext
}

func NewDeclarationStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclarationStmtContext {
	var p = new(DeclarationStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *DeclarationStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationStmtContext) Assignable() IAssignableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignableContext)
}

func (s *DeclarationStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(QuarkParserSEMI, 0)
}

func (s *DeclarationStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterDeclarationStmt(s)
	}
}

func (s *DeclarationStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitDeclarationStmt(s)
	}
}

func (s *DeclarationStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitDeclarationStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignStmtContext struct {
	*StmtContext
}

func NewAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStmtContext {
	var p = new(AssignStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) Assignable() IAssignableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignableContext)
}

func (s *AssignStmtContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *AssignStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(QuarkParserSEMI, 0)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type FutureStmtContext struct {
	*StmtContext
}

func NewFutureStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FutureStmtContext {
	var p = new(FutureStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *FutureStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FutureStmtContext) KW_FUTURE() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_FUTURE, 0)
}

func (s *FutureStmtContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *FutureStmtContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *FutureStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(QuarkParserSEMI, 0)
}

func (s *FutureStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterFutureStmt(s)
	}
}

func (s *FutureStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitFutureStmt(s)
	}
}

func (s *FutureStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitFutureStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	*StmtContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) KW_RETURN() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_RETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(QuarkParserSEMI, 0)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type RegAssignStmtContext struct {
	*StmtContext
	clk IClockexprContext
	rst IClockexprContext
}

func NewRegAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegAssignStmtContext {
	var p = new(RegAssignStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *RegAssignStmtContext) GetClk() IClockexprContext { return s.clk }

func (s *RegAssignStmtContext) GetRst() IClockexprContext { return s.rst }

func (s *RegAssignStmtContext) SetClk(v IClockexprContext) { s.clk = v }

func (s *RegAssignStmtContext) SetRst(v IClockexprContext) { s.rst = v }

func (s *RegAssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegAssignStmtContext) KW_REG() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_REG, 0)
}

func (s *RegAssignStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserLPAREN, 0)
}

func (s *RegAssignStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
}

func (s *RegAssignStmtContext) Assignable() IAssignableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignableContext)
}

func (s *RegAssignStmtContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *RegAssignStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RegAssignStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(QuarkParserSEMI, 0)
}

func (s *RegAssignStmtContext) AllClockexpr() []IClockexprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClockexprContext)(nil)).Elem())
	var tst = make([]IClockexprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClockexprContext)
		}
	}

	return tst
}

func (s *RegAssignStmtContext) Clockexpr(i int) IClockexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClockexprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClockexprContext)
}

func (s *RegAssignStmtContext) COMMA() antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, 0)
}

func (s *RegAssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterRegAssignStmt(s)
	}
}

func (s *RegAssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitRegAssignStmt(s)
	}
}

func (s *RegAssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitRegAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QuarkParserRULE_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.assignable(0)
		}
		{
			p.SetState(116)
			p.Assignment()
		}
		{
			p.SetState(117)
			p.expr(0)
		}
		{
			p.SetState(118)
			p.Match(QuarkParserSEMI)
		}

	case 2:
		localctx = NewRegAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(QuarkParserKW_REG)
		}
		{
			p.SetState(121)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(122)

			var _x = p.Clockexpr()

			localctx.(*RegAssignStmtContext).clk = _x
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserCOMMA {
			{
				p.SetState(123)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(124)

				var _x = p.Clockexpr()

				localctx.(*RegAssignStmtContext).rst = _x
			}

		}
		{
			p.SetState(127)
			p.Match(QuarkParserRPAREN)
		}
		{
			p.SetState(128)
			p.assignable(0)
		}
		{
			p.SetState(129)
			p.Assignment()
		}
		{
			p.SetState(130)
			p.expr(0)
		}
		{
			p.SetState(131)
			p.Match(QuarkParserSEMI)
		}

	case 3:
		localctx = NewFutureStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.Match(QuarkParserKW_FUTURE)
		}
		{
			p.SetState(134)
			p.typeexpr(0)
		}
		{
			p.SetState(135)
			p.Realname()
		}
		{
			p.SetState(136)
			p.Match(QuarkParserSEMI)
		}

	case 4:
		localctx = NewDeclarationStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.assignable(0)
		}
		{
			p.SetState(139)
			p.Match(QuarkParserSEMI)
		}

	case 5:
		localctx = NewBranchStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(141)
			p.Branch()
		}

	case 6:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(142)
			p.Match(QuarkParserKW_RETURN)
		}
		{
			p.SetState(143)
			p.expr(0)
		}
		{
			p.SetState(144)
			p.Match(QuarkParserSEMI)
		}

	}

	return localctx
}

// IAssignableContext is an interface to support dynamic dispatch.
type IAssignableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignableContext differentiates from other interfaces.
	IsAssignableContext()
}

type AssignableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignableContext() *AssignableContext {
	var p = new(AssignableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_assignable
	return p
}

func (*AssignableContext) IsAssignableContext() {}

func NewAssignableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignableContext {
	var p = new(AssignableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_assignable

	return p
}

func (s *AssignableContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignableContext) CopyFrom(ctx *AssignableContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AssignableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValueAssignmentContext struct {
	*AssignableContext
}

func NewValueAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueAssignmentContext {
	var p = new(ValueAssignmentContext)

	p.AssignableContext = NewEmptyAssignableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignableContext))

	return p
}

func (s *ValueAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueAssignmentContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ValueAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterValueAssignment(s)
	}
}

func (s *ValueAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitValueAssignment(s)
	}
}

func (s *ValueAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitValueAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableDefinitionContext struct {
	*AssignableContext
}

func NewVariableDefinitionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)

	p.AssignableContext = NewEmptyAssignableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignableContext))

	return p
}

func (s *VariableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *VariableDefinitionContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *VariableDefinitionContext) KW_MUT() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_MUT, 0)
}

func (s *VariableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitVariableDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleDestructerContext struct {
	*AssignableContext
}

func NewTupleDestructerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleDestructerContext {
	var p = new(TupleDestructerContext)

	p.AssignableContext = NewEmptyAssignableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignableContext))

	return p
}

func (s *TupleDestructerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleDestructerContext) AllAssignable() []IAssignableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignableContext)(nil)).Elem())
	var tst = make([]IAssignableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignableContext)
		}
	}

	return tst
}

func (s *TupleDestructerContext) Assignable(i int) IAssignableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignableContext)
}

func (s *TupleDestructerContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *TupleDestructerContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *TupleDestructerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterTupleDestructer(s)
	}
}

func (s *TupleDestructerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitTupleDestructer(s)
	}
}

func (s *TupleDestructerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitTupleDestructer(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayIndexAssignmentContext struct {
	*AssignableContext
}

func NewArrayIndexAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayIndexAssignmentContext {
	var p = new(ArrayIndexAssignmentContext)

	p.AssignableContext = NewEmptyAssignableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignableContext))

	return p
}

func (s *ArrayIndexAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexAssignmentContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ArrayIndexAssignmentContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayIndexAssignmentContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLBRACE, 0)
}

func (s *ArrayIndexAssignmentContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRBRACE, 0)
}

func (s *ArrayIndexAssignmentContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *ArrayIndexAssignmentContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *ArrayIndexAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterArrayIndexAssignment(s)
	}
}

func (s *ArrayIndexAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitArrayIndexAssignment(s)
	}
}

func (s *ArrayIndexAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitArrayIndexAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArraySliceAssignmentContext struct {
	*AssignableContext
	msb  IExprContext
	lsb  IExprContext
	step IExprContext
}

func NewArraySliceAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArraySliceAssignmentContext {
	var p = new(ArraySliceAssignmentContext)

	p.AssignableContext = NewEmptyAssignableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignableContext))

	return p
}

func (s *ArraySliceAssignmentContext) GetMsb() IExprContext { return s.msb }

func (s *ArraySliceAssignmentContext) GetLsb() IExprContext { return s.lsb }

func (s *ArraySliceAssignmentContext) GetStep() IExprContext { return s.step }

func (s *ArraySliceAssignmentContext) SetMsb(v IExprContext) { s.msb = v }

func (s *ArraySliceAssignmentContext) SetLsb(v IExprContext) { s.lsb = v }

func (s *ArraySliceAssignmentContext) SetStep(v IExprContext) { s.step = v }

func (s *ArraySliceAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraySliceAssignmentContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ArraySliceAssignmentContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArraySliceAssignmentContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLBRACE, 0)
}

func (s *ArraySliceAssignmentContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOLON)
}

func (s *ArraySliceAssignmentContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOLON, i)
}

func (s *ArraySliceAssignmentContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRBRACE, 0)
}

func (s *ArraySliceAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterArraySliceAssignment(s)
	}
}

func (s *ArraySliceAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitArraySliceAssignment(s)
	}
}

func (s *ArraySliceAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitArraySliceAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Assignable() (localctx IAssignableContext) {
	return p.assignable(0)
}

func (p *QuarkParser) assignable(_p int) (localctx IAssignableContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAssignableContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAssignableContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, QuarkParserRULE_assignable, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArrayIndexAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(149)
			p.expr(0)
		}
		{
			p.SetState(150)
			p.Match(QuarkParserLBRACE)
		}
		{
			p.SetState(151)
			p.expr(0)
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(152)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(153)
				p.expr(0)
			}

			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(159)
			p.Match(QuarkParserRBRACE)
		}

	case 2:
		localctx = NewArraySliceAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(161)
			p.expr(0)
		}
		{
			p.SetState(162)
			p.Match(QuarkParserLBRACE)
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
			{
				p.SetState(163)

				var _x = p.expr(0)

				localctx.(*ArraySliceAssignmentContext).msb = _x
			}

		}
		{
			p.SetState(166)
			p.Match(QuarkParserCOLON)
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
			{
				p.SetState(167)

				var _x = p.expr(0)

				localctx.(*ArraySliceAssignmentContext).lsb = _x
			}

		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserCOLON {
			{
				p.SetState(170)
				p.Match(QuarkParserCOLON)
			}
			p.SetState(172)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
				{
					p.SetState(171)

					var _x = p.expr(0)

					localctx.(*ArraySliceAssignmentContext).step = _x
				}

			}

		}
		{
			p.SetState(176)
			p.Match(QuarkParserRBRACE)
		}

	case 3:
		localctx = NewValueAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(178)
			p.Name()
		}

	case 4:
		localctx = NewVariableDefinitionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserKW_MUT {
			{
				p.SetState(179)
				p.Match(QuarkParserKW_MUT)
			}

		}
		{
			p.SetState(182)
			p.typeexpr(0)
		}
		{
			p.SetState(183)
			p.Realname()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTupleDestructerContext(p, NewAssignableContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_assignable)
			p.SetState(187)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(188)
						p.Match(QuarkParserCOMMA)
					}
					{
						p.SetState(189)
						p.assignable(0)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(192)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
			}

		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IRealnameContext is an interface to support dynamic dispatch.
type IRealnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealnameContext differentiates from other interfaces.
	IsRealnameContext()
}

type RealnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealnameContext() *RealnameContext {
	var p = new(RealnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_realname
	return p
}

func (*RealnameContext) IsRealnameContext() {}

func NewRealnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealnameContext {
	var p = new(RealnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_realname

	return p
}

func (s *RealnameContext) GetParser() antlr.Parser { return s.parser }

func (s *RealnameContext) REAL_NAME() antlr.TerminalNode {
	return s.GetToken(QuarkParserREAL_NAME, 0)
}

func (s *RealnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterRealname(s)
	}
}

func (s *RealnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitRealname(s)
	}
}

func (s *RealnameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitRealname(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Realname() (localctx IRealnameContext) {
	localctx = NewRealnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QuarkParserRULE_realname)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(QuarkParserREAL_NAME)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) CopyFrom(ctx *NameContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnqualifiedNameContext struct {
	*NameContext
}

func NewUnqualifiedNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnqualifiedNameContext {
	var p = new(UnqualifiedNameContext)

	p.NameContext = NewEmptyNameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NameContext))

	return p
}

func (s *UnqualifiedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnqualifiedNameContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *UnqualifiedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterUnqualifiedName(s)
	}
}

func (s *UnqualifiedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitUnqualifiedName(s)
	}
}

func (s *UnqualifiedNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitUnqualifiedName(s)

	default:
		return t.VisitChildren(s)
	}
}

type QualifiedNameContext struct {
	*NameContext
}

func NewQualifiedNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QualifiedNameContext {
	var p = new(QualifiedNameContext)

	p.NameContext = NewEmptyNameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NameContext))

	return p
}

func (s *QualifiedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameContext) AllREAL_NAME() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserREAL_NAME)
}

func (s *QualifiedNameContext) REAL_NAME(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserREAL_NAME, i)
}

func (s *QualifiedNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserDOT)
}

func (s *QualifiedNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserDOT, i)
}

func (s *QualifiedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterQualifiedName(s)
	}
}

func (s *QualifiedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitQualifiedName(s)
	}
}

func (s *QualifiedNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitQualifiedName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, QuarkParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnqualifiedNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.Realname()
		}

	case 2:
		localctx = NewQualifiedNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Match(QuarkParserREAL_NAME)
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(203)
					p.Match(QuarkParserDOT)
				}
				{
					p.SetState(204)
					p.Match(QuarkParserREAL_NAME)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IClockexprContext is an interface to support dynamic dispatch.
type IClockexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClockexprContext differentiates from other interfaces.
	IsClockexprContext()
}

type ClockexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClockexprContext() *ClockexprContext {
	var p = new(ClockexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_clockexpr
	return p
}

func (*ClockexprContext) IsClockexprContext() {}

func NewClockexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClockexprContext {
	var p = new(ClockexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_clockexpr

	return p
}

func (s *ClockexprContext) GetParser() antlr.Parser { return s.parser }

func (s *ClockexprContext) CopyFrom(ctx *ClockexprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ClockexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClockexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AtomicClockContext struct {
	*ClockexprContext
}

func NewAtomicClockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomicClockContext {
	var p = new(AtomicClockContext)

	p.ClockexprContext = NewEmptyClockexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ClockexprContext))

	return p
}

func (s *AtomicClockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomicClockContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AtomicClockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterAtomicClock(s)
	}
}

func (s *AtomicClockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitAtomicClock(s)
	}
}

func (s *AtomicClockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitAtomicClock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Clockexpr() (localctx IClockexprContext) {
	localctx = NewClockexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, QuarkParserRULE_clockexpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewAtomicClockContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Name()
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayIndexExprContext struct {
	*ExprContext
}

func NewArrayIndexExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayIndexExprContext {
	var p = new(ArrayIndexExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ArrayIndexExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ArrayIndexExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayIndexExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLBRACE, 0)
}

func (s *ArrayIndexExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRBRACE, 0)
}

func (s *ArrayIndexExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *ArrayIndexExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *ArrayIndexExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterArrayIndexExpr(s)
	}
}

func (s *ArrayIndexExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitArrayIndexExpr(s)
	}
}

func (s *ArrayIndexExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitArrayIndexExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BranchExprContext struct {
	*ExprContext
}

func NewBranchExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BranchExprContext {
	var p = new(BranchExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BranchExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchExprContext) Branch() IBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *BranchExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterBranchExpr(s)
	}
}

func (s *BranchExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitBranchExpr(s)
	}
}

func (s *BranchExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitBranchExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LambdaExprContext struct {
	*ExprContext
}

func NewLambdaExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LambdaExprContext {
	var p = new(LambdaExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LambdaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaExprContext) KW_LAMBDA() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_LAMBDA, 0)
}

func (s *LambdaExprContext) Argumentlist() IArgumentlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentlistContext)
}

func (s *LambdaExprContext) OP_ARROW() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_ARROW, 0)
}

func (s *LambdaExprContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserLCURLY, 0)
}

func (s *LambdaExprContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LambdaExprContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserRCURLY, 0)
}

func (s *LambdaExprContext) Parameterlist() IParameterlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterlistContext)
}

func (s *LambdaExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LambdaExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterLambdaExpr(s)
	}
}

func (s *LambdaExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitLambdaExpr(s)
	}
}

func (s *LambdaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitLambdaExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstructorExprContext struct {
	*ExprContext
}

func NewConstructorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstructorExprContext {
	var p = new(ConstructorExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConstructorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructorExprContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *ConstructorExprContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserLCURLY, 0)
}

func (s *ConstructorExprContext) Callarglist() ICallarglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallarglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallarglistContext)
}

func (s *ConstructorExprContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserRCURLY, 0)
}

func (s *ConstructorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterConstructorExpr(s)
	}
}

func (s *ConstructorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitConstructorExpr(s)
	}
}

func (s *ConstructorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitConstructorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleExprContext struct {
	*ExprContext
}

func NewTupleExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleExprContext {
	var p = new(TupleExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TupleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserLPAREN, 0)
}

func (s *TupleExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TupleExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TupleExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
}

func (s *TupleExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *TupleExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *TupleExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterTupleExpr(s)
	}
}

func (s *TupleExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitTupleExpr(s)
	}
}

func (s *TupleExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitTupleExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalBinopExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewLogicalBinopExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalBinopExprContext {
	var p = new(LogicalBinopExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LogicalBinopExprContext) GetOp() antlr.Token { return s.op }

func (s *LogicalBinopExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalBinopExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalBinopExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LogicalBinopExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogicalBinopExprContext) OP_LAND() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_LAND, 0)
}

func (s *LogicalBinopExprContext) OP_LOR() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_LOR, 0)
}

func (s *LogicalBinopExprContext) OP_IMPLICATION() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_IMPLICATION, 0)
}

func (s *LogicalBinopExprContext) OP_EQUIVALENCE() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_EQUIVALENCE, 0)
}

func (s *LogicalBinopExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterLogicalBinopExpr(s)
	}
}

func (s *LogicalBinopExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitLogicalBinopExpr(s)
	}
}

func (s *LogicalBinopExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitLogicalBinopExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConcatExprContext struct {
	*ExprContext
}

func NewConcatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConcatExprContext {
	var p = new(ConcatExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConcatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatExprContext) Concat() IConcatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcatContext)
}

func (s *ConcatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterConcatExpr(s)
	}
}

func (s *ConcatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitConcatExpr(s)
	}
}

func (s *ConcatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitConcatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectorExprContext struct {
	*ExprContext
}

func NewSelectorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorExprContext {
	var p = new(SelectorExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SelectorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SelectorExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(QuarkParserDOT, 0)
}

func (s *SelectorExprContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *SelectorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterSelectorExpr(s)
	}
}

func (s *SelectorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitSelectorExpr(s)
	}
}

func (s *SelectorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitSelectorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivModExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewMulDivModExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModExprContext {
	var p = new(MulDivModExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivModExprContext) GetOp() antlr.Token { return s.op }

func (s *MulDivModExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivModExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MulDivModExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivModExprContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_MUL, 0)
}

func (s *MulDivModExprContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_DIV, 0)
}

func (s *MulDivModExprContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_MOD, 0)
}

func (s *MulDivModExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterMulDivModExpr(s)
	}
}

func (s *MulDivModExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitMulDivModExpr(s)
	}
}

func (s *MulDivModExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitMulDivModExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComplimentExprContext struct {
	*ExprContext
}

func NewComplimentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComplimentExprContext {
	var p = new(ComplimentExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ComplimentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComplimentExprContext) OP_COMPLIMENT() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_COMPLIMENT, 0)
}

func (s *ComplimentExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ComplimentExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterComplimentExpr(s)
	}
}

func (s *ComplimentExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitComplimentExpr(s)
	}
}

func (s *ComplimentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitComplimentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayLiteralExprContext struct {
	*ExprContext
}

func NewArrayLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralExprContext {
	var p = new(ArrayLiteralExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ArrayLiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLBRACE, 0)
}

func (s *ArrayLiteralExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRBRACE, 0)
}

func (s *ArrayLiteralExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ArrayLiteralExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayLiteralExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *ArrayLiteralExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *ArrayLiteralExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterArrayLiteralExpr(s)
	}
}

func (s *ArrayLiteralExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitArrayLiteralExpr(s)
	}
}

func (s *ArrayLiteralExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitArrayLiteralExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ClockToExprContext struct {
	*ExprContext
}

func NewClockToExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClockToExprContext {
	var p = new(ClockToExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ClockToExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClockToExprContext) KW_SIGNAL() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_SIGNAL, 0)
}

func (s *ClockToExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserLPAREN, 0)
}

func (s *ClockToExprContext) Clockexpr() IClockexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClockexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClockexprContext)
}

func (s *ClockToExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
}

func (s *ClockToExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterClockToExpr(s)
	}
}

func (s *ClockToExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitClockToExpr(s)
	}
}

func (s *ClockToExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitClockToExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExprContext struct {
	*ExprContext
}

func NewLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExprContext {
	var p = new(LiteralExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExprContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterLiteralExpr(s)
	}
}

func (s *LiteralExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitLiteralExpr(s)
	}
}

func (s *LiteralExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitLiteralExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarExprContext struct {
	*ExprContext
}

func NewVarExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarExprContext {
	var p = new(VarExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VarExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarExprContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *VarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterVarExpr(s)
	}
}

func (s *VarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitVarExpr(s)
	}
}

func (s *VarExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitVarExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensExprContext struct {
	*ExprContext
}

func NewParensExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensExprContext {
	var p = new(ParensExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParensExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserLPAREN, 0)
}

func (s *ParensExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
}

func (s *ParensExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterParensExpr(s)
	}
}

func (s *ParensExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitParensExpr(s)
	}
}

func (s *ParensExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitParensExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitwiseBinopExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewBitwiseBinopExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitwiseBinopExprContext {
	var p = new(BitwiseBinopExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BitwiseBinopExprContext) GetOp() antlr.Token { return s.op }

func (s *BitwiseBinopExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *BitwiseBinopExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseBinopExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *BitwiseBinopExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BitwiseBinopExprContext) OP_BAND() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_BAND, 0)
}

func (s *BitwiseBinopExprContext) OP_BOR() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_BOR, 0)
}

func (s *BitwiseBinopExprContext) OP_XOR() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_XOR, 0)
}

func (s *BitwiseBinopExprContext) OP_BNAND() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_BNAND, 0)
}

func (s *BitwiseBinopExprContext) OP_BNOR() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_BNOR, 0)
}

func (s *BitwiseBinopExprContext) OP_XNOR() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_XNOR, 0)
}

func (s *BitwiseBinopExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterBitwiseBinopExpr(s)
	}
}

func (s *BitwiseBinopExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitBitwiseBinopExpr(s)
	}
}

func (s *BitwiseBinopExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitBitwiseBinopExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceExprContext struct {
	*ExprContext
	msb  IExprContext
	lsb  IExprContext
	step IExprContext
}

func NewSliceExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceExprContext {
	var p = new(SliceExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SliceExprContext) GetMsb() IExprContext { return s.msb }

func (s *SliceExprContext) GetLsb() IExprContext { return s.lsb }

func (s *SliceExprContext) GetStep() IExprContext { return s.step }

func (s *SliceExprContext) SetMsb(v IExprContext) { s.msb = v }

func (s *SliceExprContext) SetLsb(v IExprContext) { s.lsb = v }

func (s *SliceExprContext) SetStep(v IExprContext) { s.step = v }

func (s *SliceExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SliceExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SliceExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLBRACE, 0)
}

func (s *SliceExprContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOLON)
}

func (s *SliceExprContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOLON, i)
}

func (s *SliceExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRBRACE, 0)
}

func (s *SliceExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterSliceExpr(s)
	}
}

func (s *SliceExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitSliceExpr(s)
	}
}

func (s *SliceExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitSliceExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	*ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) OP_LNOT() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_LNOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	*ExprContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserLPAREN, 0)
}

func (s *FunctionCallContext) Callarglist() ICallarglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallarglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallarglistContext)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
}

func (s *FunctionCallContext) Paramarglist() IParamarglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamarglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamarglistContext)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext {
	var p = new(AddSubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubExprContext) GetOp() antlr.Token { return s.op }

func (s *AddSubExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddSubExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubExprContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_SUB, 0)
}

func (s *AddSubExprContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_ADD, 0)
}

func (s *AddSubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterAddSubExpr(s)
	}
}

func (s *AddSubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitAddSubExpr(s)
	}
}

func (s *AddSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitAddSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ShiftExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewShiftExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShiftExprContext {
	var p = new(ShiftExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ShiftExprContext) GetOp() antlr.Token { return s.op }

func (s *ShiftExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ShiftExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ShiftExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ShiftExprContext) OP_LEFT_SHIFT() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_LEFT_SHIFT, 0)
}

func (s *ShiftExprContext) OP_RIGHT_SHIFT() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_RIGHT_SHIFT, 0)
}

func (s *ShiftExprContext) OP_ARITH_LEFT_SHIFT() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_ARITH_LEFT_SHIFT, 0)
}

func (s *ShiftExprContext) OP_ARITH_RIGHT_SHFIT() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_ARITH_RIGHT_SHFIT, 0)
}

func (s *ShiftExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterShiftExpr(s)
	}
}

func (s *ShiftExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitShiftExpr(s)
	}
}

func (s *ShiftExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitShiftExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TernaryExprContext struct {
	*ExprContext
}

func NewTernaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExprContext {
	var p = new(TernaryExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TernaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TernaryExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TernaryExprContext) KW_IF() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_IF, 0)
}

func (s *TernaryExprContext) KW_ELSE() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_ELSE, 0)
}

func (s *TernaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterTernaryExpr(s)
	}
}

func (s *TernaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitTernaryExpr(s)
	}
}

func (s *TernaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitTernaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NewModuleExprContext struct {
	*ExprContext
}

func NewNewModuleExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewModuleExprContext {
	var p = new(NewModuleExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NewModuleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewModuleExprContext) KW_NEW() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_NEW, 0)
}

func (s *NewModuleExprContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *NewModuleExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserLPAREN, 0)
}

func (s *NewModuleExprContext) Callarglist() ICallarglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallarglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallarglistContext)
}

func (s *NewModuleExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
}

func (s *NewModuleExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterNewModuleExpr(s)
	}
}

func (s *NewModuleExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitNewModuleExpr(s)
	}
}

func (s *NewModuleExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitNewModuleExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *QuarkParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, QuarkParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(214)
			p.Literal()
		}

	case 2:
		localctx = NewVarExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(215)
			p.Realname()
		}

	case 3:
		localctx = NewParensExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(216)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(217)
			p.expr(0)
		}
		{
			p.SetState(218)
			p.Match(QuarkParserRPAREN)
		}

	case 4:
		localctx = NewTupleExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(220)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(221)
			p.expr(0)
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QuarkParserCOMMA {
			{
				p.SetState(222)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(223)
				p.expr(0)
			}

			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(228)
			p.Match(QuarkParserRPAREN)
		}

	case 5:
		localctx = NewConstructorExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(230)
			p.typeexpr(0)
		}
		{
			p.SetState(231)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(232)
			p.Callarglist()
		}
		{
			p.SetState(233)
			p.Match(QuarkParserRCURLY)
		}

	case 6:
		localctx = NewNewModuleExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(235)
			p.Match(QuarkParserKW_NEW)
		}
		{
			p.SetState(236)
			p.typeexpr(0)
		}
		{
			p.SetState(237)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(238)
			p.Callarglist()
		}
		{
			p.SetState(239)
			p.Match(QuarkParserRPAREN)
		}

	case 7:
		localctx = NewLambdaExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(241)
			p.Match(QuarkParserKW_LAMBDA)
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserLANGLE {
			{
				p.SetState(242)
				p.Parameterlist()
			}

		}
		{
			p.SetState(245)
			p.Argumentlist()
		}
		{
			p.SetState(246)
			p.Match(QuarkParserOP_ARROW)
		}
		{
			p.SetState(247)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(248)
			p.Block()
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
			{
				p.SetState(249)
				p.expr(0)
			}

		}
		{
			p.SetState(252)
			p.Match(QuarkParserRCURLY)
		}

	case 8:
		localctx = NewComplimentExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(254)
			p.Match(QuarkParserOP_COMPLIMENT)
		}
		{
			p.SetState(255)
			p.expr(14)
		}

	case 9:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(256)
			p.Match(QuarkParserOP_LNOT)
		}
		{
			p.SetState(257)
			p.expr(13)
		}

	case 10:
		localctx = NewConcatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(258)
			p.Concat()
		}

	case 11:
		localctx = NewBranchExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(259)
			p.Branch()
		}

	case 12:
		localctx = NewArrayLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(260)
			p.Match(QuarkParserLBRACE)
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
			{
				p.SetState(261)
				p.expr(0)
			}
			p.SetState(266)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == QuarkParserCOMMA {
				{
					p.SetState(262)
					p.Match(QuarkParserCOMMA)
				}
				{
					p.SetState(263)
					p.expr(0)
				}

				p.SetState(268)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(271)
			p.Match(QuarkParserRBRACE)
		}

	case 13:
		localctx = NewClockToExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(272)
			p.Match(QuarkParserKW_SIGNAL)
		}
		{
			p.SetState(273)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(274)
			p.Clockexpr()
		}
		{
			p.SetState(275)
			p.Match(QuarkParserRPAREN)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(339)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivModExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(279)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(280)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserOP_MUL)|(1<<QuarkParserOP_DIV)|(1<<QuarkParserOP_MOD))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(281)
					p.expr(12)
				}

			case 2:
				localctx = NewAddSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(283)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == QuarkParserOP_ADD || _la == QuarkParserOP_SUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(284)
					p.expr(11)
				}

			case 3:
				localctx = NewShiftExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(285)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(286)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ShiftExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(QuarkParserOP_LEFT_SHIFT-35))|(1<<(QuarkParserOP_RIGHT_SHIFT-35))|(1<<(QuarkParserOP_ARITH_LEFT_SHIFT-35))|(1<<(QuarkParserOP_ARITH_RIGHT_SHFIT-35)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ShiftExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(287)
					p.expr(10)
				}

			case 4:
				localctx = NewBitwiseBinopExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(289)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BitwiseBinopExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(QuarkParserOP_BAND-26))|(1<<(QuarkParserOP_BNAND-26))|(1<<(QuarkParserOP_BOR-26))|(1<<(QuarkParserOP_BNOR-26))|(1<<(QuarkParserOP_XOR-26))|(1<<(QuarkParserOP_XNOR-26)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BitwiseBinopExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(290)
					p.expr(9)
				}

			case 5:
				localctx = NewLogicalBinopExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(291)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(292)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalBinopExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(QuarkParserOP_LAND-44))|(1<<(QuarkParserOP_LOR-44))|(1<<(QuarkParserOP_IMPLICATION-44))|(1<<(QuarkParserOP_EQUIVALENCE-44)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalBinopExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(293)
					p.expr(8)
				}

			case 6:
				localctx = NewTernaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(294)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(295)
					p.Match(QuarkParserKW_IF)
				}
				{
					p.SetState(296)
					p.expr(0)
				}
				{
					p.SetState(297)
					p.Match(QuarkParserKW_ELSE)
				}
				{
					p.SetState(298)
					p.expr(7)
				}

			case 7:
				localctx = NewSelectorExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(300)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(301)
					p.Match(QuarkParserDOT)
				}
				{
					p.SetState(302)
					p.Realname()
				}

			case 8:
				localctx = NewFunctionCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(303)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				p.SetState(305)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == QuarkParserLANGLE {
					{
						p.SetState(304)
						p.Paramarglist()
					}

				}
				{
					p.SetState(307)
					p.Match(QuarkParserLPAREN)
				}
				{
					p.SetState(308)
					p.Callarglist()
				}
				{
					p.SetState(309)
					p.Match(QuarkParserRPAREN)
				}

			case 9:
				localctx = NewSliceExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(311)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(312)
					p.Match(QuarkParserLBRACE)
				}
				p.SetState(314)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
					{
						p.SetState(313)

						var _x = p.expr(0)

						localctx.(*SliceExprContext).msb = _x
					}

				}
				{
					p.SetState(316)
					p.Match(QuarkParserCOLON)
				}
				p.SetState(318)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(317)

						var _x = p.expr(0)

						localctx.(*SliceExprContext).lsb = _x
					}

				}

				p.SetState(321)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == QuarkParserCOLON {
					{
						p.SetState(320)
						p.Match(QuarkParserCOLON)
					}

				}
				p.SetState(324)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
					{
						p.SetState(323)

						var _x = p.expr(0)

						localctx.(*SliceExprContext).step = _x
					}

				}

				{
					p.SetState(326)
					p.Match(QuarkParserRBRACE)
				}

			case 10:
				localctx = NewArrayIndexExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(327)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(328)
					p.Match(QuarkParserLBRACE)
				}
				{
					p.SetState(329)
					p.expr(0)
				}
				p.SetState(334)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == QuarkParserCOMMA {
					{
						p.SetState(330)
						p.Match(QuarkParserCOMMA)
					}
					{
						p.SetState(331)
						p.expr(0)
					}

					p.SetState(336)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(337)
					p.Match(QuarkParserRBRACE)
				}

			}

		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
	}

	return localctx
}

// ICallarglistContext is an interface to support dynamic dispatch.
type ICallarglistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallarglistContext differentiates from other interfaces.
	IsCallarglistContext()
}

type CallarglistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallarglistContext() *CallarglistContext {
	var p = new(CallarglistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_callarglist
	return p
}

func (*CallarglistContext) IsCallarglistContext() {}

func NewCallarglistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallarglistContext {
	var p = new(CallarglistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_callarglist

	return p
}

func (s *CallarglistContext) GetParser() antlr.Parser { return s.parser }

func (s *CallarglistContext) AllCallarg() []ICallargContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICallargContext)(nil)).Elem())
	var tst = make([]ICallargContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICallargContext)
		}
	}

	return tst
}

func (s *CallarglistContext) Callarg(i int) ICallargContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallargContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICallargContext)
}

func (s *CallarglistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *CallarglistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *CallarglistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallarglistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallarglistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterCallarglist(s)
	}
}

func (s *CallarglistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitCallarglist(s)
	}
}

func (s *CallarglistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitCallarglist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Callarglist() (localctx ICallarglistContext) {
	localctx = NewCallarglistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, QuarkParserRULE_callarglist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Callarg()
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserCOMMA {
		{
			p.SetState(345)
			p.Match(QuarkParserCOMMA)
		}
		{
			p.SetState(346)
			p.Callarg()
		}

		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICallargContext is an interface to support dynamic dispatch.
type ICallargContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallargContext differentiates from other interfaces.
	IsCallargContext()
}

type CallargContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallargContext() *CallargContext {
	var p = new(CallargContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_callarg
	return p
}

func (*CallargContext) IsCallargContext() {}

func NewCallargContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallargContext {
	var p = new(CallargContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_callarg

	return p
}

func (s *CallargContext) GetParser() antlr.Parser { return s.parser }

func (s *CallargContext) CopyFrom(ctx *CallargContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CallargContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallargContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnamedCallArgContext struct {
	*CallargContext
}

func NewUnamedCallArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnamedCallArgContext {
	var p = new(UnamedCallArgContext)

	p.CallargContext = NewEmptyCallargContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CallargContext))

	return p
}

func (s *UnamedCallArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnamedCallArgContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnamedCallArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterUnamedCallArg(s)
	}
}

func (s *UnamedCallArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitUnamedCallArg(s)
	}
}

func (s *UnamedCallArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitUnamedCallArg(s)

	default:
		return t.VisitChildren(s)
	}
}

type NamedCallArgContext struct {
	*CallargContext
}

func NewNamedCallArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedCallArgContext {
	var p = new(NamedCallArgContext)

	p.CallargContext = NewEmptyCallargContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CallargContext))

	return p
}

func (s *NamedCallArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedCallArgContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *NamedCallArgContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_ASSIGN, 0)
}

func (s *NamedCallArgContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NamedCallArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterNamedCallArg(s)
	}
}

func (s *NamedCallArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitNamedCallArg(s)
	}
}

func (s *NamedCallArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitNamedCallArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Callarg() (localctx ICallargContext) {
	localctx = NewCallargContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, QuarkParserRULE_callarg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNamedCallArgContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(352)
			p.Realname()
		}
		{
			p.SetState(353)
			p.Match(QuarkParserOP_ASSIGN)
		}
		{
			p.SetState(354)
			p.expr(0)
		}

	case 2:
		localctx = NewUnamedCallArgContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)
			p.expr(0)
		}

	}

	return localctx
}

// IParamarglistContext is an interface to support dynamic dispatch.
type IParamarglistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamarglistContext differentiates from other interfaces.
	IsParamarglistContext()
}

type ParamarglistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamarglistContext() *ParamarglistContext {
	var p = new(ParamarglistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_paramarglist
	return p
}

func (*ParamarglistContext) IsParamarglistContext() {}

func NewParamarglistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamarglistContext {
	var p = new(ParamarglistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_paramarglist

	return p
}

func (s *ParamarglistContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamarglistContext) LANGLE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLANGLE, 0)
}

func (s *ParamarglistContext) AllParamarg() []IParamargContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamargContext)(nil)).Elem())
	var tst = make([]IParamargContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamargContext)
		}
	}

	return tst
}

func (s *ParamarglistContext) Paramarg(i int) IParamargContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamargContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamargContext)
}

func (s *ParamarglistContext) RANGLE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRANGLE, 0)
}

func (s *ParamarglistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *ParamarglistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *ParamarglistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamarglistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamarglistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterParamarglist(s)
	}
}

func (s *ParamarglistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitParamarglist(s)
	}
}

func (s *ParamarglistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitParamarglist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Paramarglist() (localctx IParamarglistContext) {
	localctx = NewParamarglistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, QuarkParserRULE_paramarglist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(QuarkParserLANGLE)
	}
	{
		p.SetState(360)
		p.Paramarg()
	}
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserCOMMA {
		{
			p.SetState(361)
			p.Match(QuarkParserCOMMA)
		}
		{
			p.SetState(362)
			p.Paramarg()
		}

		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(368)
		p.Match(QuarkParserRANGLE)
	}

	return localctx
}

// IConcatContext is an interface to support dynamic dispatch.
type IConcatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcatContext differentiates from other interfaces.
	IsConcatContext()
}

type ConcatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcatContext() *ConcatContext {
	var p = new(ConcatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_concat
	return p
}

func (*ConcatContext) IsConcatContext() {}

func NewConcatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcatContext {
	var p = new(ConcatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_concat

	return p
}

func (s *ConcatContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcatContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserLCURLY, 0)
}

func (s *ConcatContext) AllInnerconcat() []IInnerconcatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInnerconcatContext)(nil)).Elem())
	var tst = make([]IInnerconcatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInnerconcatContext)
		}
	}

	return tst
}

func (s *ConcatContext) Innerconcat(i int) IInnerconcatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInnerconcatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInnerconcatContext)
}

func (s *ConcatContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserRCURLY, 0)
}

func (s *ConcatContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *ConcatContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *ConcatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterConcat(s)
	}
}

func (s *ConcatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitConcat(s)
	}
}

func (s *ConcatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitConcat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Concat() (localctx IConcatContext) {
	localctx = NewConcatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, QuarkParserRULE_concat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(QuarkParserLCURLY)
	}
	{
		p.SetState(371)
		p.Innerconcat()
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == QuarkParserCOMMA {
		{
			p.SetState(372)
			p.Match(QuarkParserCOMMA)
		}
		{
			p.SetState(373)
			p.Innerconcat()
		}

		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(378)
		p.Match(QuarkParserRCURLY)
	}

	return localctx
}

// IInnerconcatContext is an interface to support dynamic dispatch.
type IInnerconcatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInnerconcatContext differentiates from other interfaces.
	IsInnerconcatContext()
}

type InnerconcatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerconcatContext() *InnerconcatContext {
	var p = new(InnerconcatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_innerconcat
	return p
}

func (*InnerconcatContext) IsInnerconcatContext() {}

func NewInnerconcatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerconcatContext {
	var p = new(InnerconcatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_innerconcat

	return p
}

func (s *InnerconcatContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerconcatContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *InnerconcatContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InnerconcatContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserLCURLY, 0)
}

func (s *InnerconcatContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserRCURLY, 0)
}

func (s *InnerconcatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerconcatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerconcatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterInnerconcat(s)
	}
}

func (s *InnerconcatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitInnerconcat(s)
	}
}

func (s *InnerconcatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitInnerconcat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Innerconcat() (localctx IInnerconcatContext) {
	localctx = NewInnerconcatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, QuarkParserRULE_innerconcat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(380)
			p.expr(0)
		}
		{
			p.SetState(381)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(382)
			p.expr(0)
		}
		{
			p.SetState(383)
			p.Match(QuarkParserRCURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(385)
			p.expr(0)
		}

	}

	return localctx
}

// ITypeexprContext is an interface to support dynamic dispatch.
type ITypeexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeexprContext differentiates from other interfaces.
	IsTypeexprContext()
}

type TypeexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeexprContext() *TypeexprContext {
	var p = new(TypeexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_typeexpr
	return p
}

func (*TypeexprContext) IsTypeexprContext() {}

func NewTypeexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeexprContext {
	var p = new(TypeexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_typeexpr

	return p
}

func (s *TypeexprContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeexprContext) CopyFrom(ctx *TypeexprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParameterizedTypeContext struct {
	*TypeexprContext
}

func NewParameterizedTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParameterizedTypeContext {
	var p = new(ParameterizedTypeContext)

	p.TypeexprContext = NewEmptyTypeexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeexprContext))

	return p
}

func (s *ParameterizedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterizedTypeContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *ParameterizedTypeContext) Paramarglist() IParamarglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamarglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamarglistContext)
}

func (s *ParameterizedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterParameterizedType(s)
	}
}

func (s *ParameterizedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitParameterizedType(s)
	}
}

func (s *ParameterizedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitParameterizedType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayTypeContext struct {
	*TypeexprContext
}

func NewArrayTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	p.TypeexprContext = NewEmptyTypeexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeexprContext))

	return p
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *ArrayTypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLBRACE, 0)
}

func (s *ArrayTypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRBRACE, 0)
}

func (s *ArrayTypeContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ArrayTypeContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayTypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *ArrayTypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompleteTypeContext struct {
	*TypeexprContext
}

func NewCompleteTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompleteTypeContext {
	var p = new(CompleteTypeContext)

	p.TypeexprContext = NewEmptyTypeexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeexprContext))

	return p
}

func (s *CompleteTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompleteTypeContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *CompleteTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterCompleteType(s)
	}
}

func (s *CompleteTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitCompleteType(s)
	}
}

func (s *CompleteTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitCompleteType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Typeexpr() (localctx ITypeexprContext) {
	return p.typeexpr(0)
}

func (p *QuarkParser) typeexpr(_p int) (localctx ITypeexprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTypeexprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeexprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, QuarkParserRULE_typeexpr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewCompleteTypeContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(389)
		p.Name()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(406)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
			case 1:
				localctx = NewParameterizedTypeContext(p, NewTypeexprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_typeexpr)
				p.SetState(391)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(392)
					p.Paramarglist()
				}

			case 2:
				localctx = NewArrayTypeContext(p, NewTypeexprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_typeexpr)
				p.SetState(393)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(394)
					p.Match(QuarkParserLBRACE)
				}
				p.SetState(403)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
					{
						p.SetState(395)
						p.expr(0)
					}
					p.SetState(400)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == QuarkParserCOMMA {
						{
							p.SetState(396)
							p.Match(QuarkParserCOMMA)
						}
						{
							p.SetState(397)
							p.expr(0)
						}

						p.SetState(402)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(405)
					p.Match(QuarkParserRBRACE)
				}

			}

		}
		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}

	return localctx
}

// IParamargContext is an interface to support dynamic dispatch.
type IParamargContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamargContext differentiates from other interfaces.
	IsParamargContext()
}

type ParamargContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamargContext() *ParamargContext {
	var p = new(ParamargContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_paramarg
	return p
}

func (*ParamargContext) IsParamargContext() {}

func NewParamargContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamargContext {
	var p = new(ParamargContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_paramarg

	return p
}

func (s *ParamargContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamargContext) KW_TYPE() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_TYPE, 0)
}

func (s *ParamargContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *ParamargContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParamargContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *ParamargContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_ASSIGN, 0)
}

func (s *ParamargContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamargContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamargContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterParamarg(s)
	}
}

func (s *ParamargContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitParamarg(s)
	}
}

func (s *ParamargContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitParamarg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Paramarg() (localctx IParamargContext) {
	localctx = NewParamargContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, QuarkParserRULE_paramarg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(411)
			p.Match(QuarkParserKW_TYPE)
		}
		{
			p.SetState(412)
			p.typeexpr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(413)
			p.expr(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(414)
			p.Realname()
		}
		{
			p.SetState(415)
			p.Match(QuarkParserOP_ASSIGN)
		}
		p.SetState(419)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case QuarkParserLPAREN, QuarkParserLBRACE, QuarkParserLCURLY, QuarkParserOP_COMPLIMENT, QuarkParserOP_LNOT, QuarkParserKW_IF, QuarkParserKW_MATCH, QuarkParserKW_LAMBDA, QuarkParserKW_NEW, QuarkParserKW_SIGNAL, QuarkParserINTEGRAL, QuarkParserREAL_NAME:
			{
				p.SetState(416)
				p.expr(0)
			}

		case QuarkParserKW_TYPE:
			{
				p.SetState(417)
				p.Match(QuarkParserKW_TYPE)
			}
			{
				p.SetState(418)
				p.typeexpr(0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IBranchContext is an interface to support dynamic dispatch.
type IBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBranchContext differentiates from other interfaces.
	IsBranchContext()
}

type BranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchContext() *BranchContext {
	var p = new(BranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_branch
	return p
}

func (*BranchContext) IsBranchContext() {}

func NewBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchContext {
	var p = new(BranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_branch

	return p
}

func (s *BranchContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchContext) CopyFrom(ctx *BranchContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MatchBranchContext struct {
	*BranchContext
}

func NewMatchBranchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchBranchContext {
	var p = new(MatchBranchContext)

	p.BranchContext = NewEmptyBranchContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BranchContext))

	return p
}

func (s *MatchBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchBranchContext) KW_MATCH() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_MATCH, 0)
}

func (s *MatchBranchContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MatchBranchContext) AllLCURLY() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserLCURLY)
}

func (s *MatchBranchContext) LCURLY(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserLCURLY, i)
}

func (s *MatchBranchContext) AllRCURLY() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserRCURLY)
}

func (s *MatchBranchContext) RCURLY(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserRCURLY, i)
}

func (s *MatchBranchContext) AllKW_CASE() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserKW_CASE)
}

func (s *MatchBranchContext) KW_CASE(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_CASE, i)
}

func (s *MatchBranchContext) AllPattern() []IPatternContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPatternContext)(nil)).Elem())
	var tst = make([]IPatternContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPatternContext)
		}
	}

	return tst
}

func (s *MatchBranchContext) Pattern(i int) IPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *MatchBranchContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *MatchBranchContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *MatchBranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterMatchBranch(s)
	}
}

func (s *MatchBranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitMatchBranch(s)
	}
}

func (s *MatchBranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitMatchBranch(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfBranchContext struct {
	*BranchContext
}

func NewIfBranchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfBranchContext {
	var p = new(IfBranchContext)

	p.BranchContext = NewEmptyBranchContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BranchContext))

	return p
}

func (s *IfBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBranchContext) KW_IF() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_IF, 0)
}

func (s *IfBranchContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IfBranchContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfBranchContext) AllLCURLY() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserLCURLY)
}

func (s *IfBranchContext) LCURLY(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserLCURLY, i)
}

func (s *IfBranchContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfBranchContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfBranchContext) AllRCURLY() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserRCURLY)
}

func (s *IfBranchContext) RCURLY(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserRCURLY, i)
}

func (s *IfBranchContext) AllKW_ELIF() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserKW_ELIF)
}

func (s *IfBranchContext) KW_ELIF(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_ELIF, i)
}

func (s *IfBranchContext) KW_ELSE() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_ELSE, 0)
}

func (s *IfBranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterIfBranch(s)
	}
}

func (s *IfBranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitIfBranch(s)
	}
}

func (s *IfBranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitIfBranch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Branch() (localctx IBranchContext) {
	localctx = NewBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, QuarkParserRULE_branch)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(461)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QuarkParserKW_IF:
		localctx = NewIfBranchContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(423)
			p.Match(QuarkParserKW_IF)
		}
		{
			p.SetState(424)
			p.expr(0)
		}
		{
			p.SetState(425)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(426)
			p.Block()
		}
		{
			p.SetState(427)
			p.Match(QuarkParserRCURLY)
		}
		p.SetState(436)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(428)
					p.Match(QuarkParserKW_ELIF)
				}
				{
					p.SetState(429)
					p.expr(0)
				}
				{
					p.SetState(430)
					p.Match(QuarkParserLCURLY)
				}
				{
					p.SetState(431)
					p.Block()
				}
				{
					p.SetState(432)
					p.Match(QuarkParserRCURLY)
				}

			}
			p.SetState(438)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
		}
		p.SetState(444)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(439)
				p.Match(QuarkParserKW_ELSE)
			}
			{
				p.SetState(440)
				p.Match(QuarkParserLCURLY)
			}
			{
				p.SetState(441)
				p.Block()
			}
			{
				p.SetState(442)
				p.Match(QuarkParserRCURLY)
			}

		}

	case QuarkParserKW_MATCH:
		localctx = NewMatchBranchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(446)
			p.Match(QuarkParserKW_MATCH)
		}
		{
			p.SetState(447)
			p.expr(0)
		}
		{
			p.SetState(448)
			p.Match(QuarkParserLCURLY)
		}
		p.SetState(455)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QuarkParserKW_CASE {
			{
				p.SetState(449)
				p.Match(QuarkParserKW_CASE)
			}
			{
				p.SetState(450)
				p.Pattern()
			}
			{
				p.SetState(451)
				p.Match(QuarkParserLCURLY)
			}
			{
				p.SetState(452)
				p.Block()
			}
			{
				p.SetState(453)
				p.Match(QuarkParserRCURLY)
			}

			p.SetState(457)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(459)
			p.Match(QuarkParserRCURLY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_pattern
	return p
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) CopyFrom(ctx *PatternContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParamerterizedTypePatternContext struct {
	*PatternContext
}

func NewParamerterizedTypePatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamerterizedTypePatternContext {
	var p = new(ParamerterizedTypePatternContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *ParamerterizedTypePatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamerterizedTypePatternContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *ParamerterizedTypePatternContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLBRACE, 0)
}

func (s *ParamerterizedTypePatternContext) AllPattern() []IPatternContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPatternContext)(nil)).Elem())
	var tst = make([]IPatternContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPatternContext)
		}
	}

	return tst
}

func (s *ParamerterizedTypePatternContext) Pattern(i int) IPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *ParamerterizedTypePatternContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRBRACE, 0)
}

func (s *ParamerterizedTypePatternContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *ParamerterizedTypePatternContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *ParamerterizedTypePatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterParamerterizedTypePattern(s)
	}
}

func (s *ParamerterizedTypePatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitParamerterizedTypePattern(s)
	}
}

func (s *ParamerterizedTypePatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitParamerterizedTypePattern(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralPatternContext struct {
	*PatternContext
}

func NewLiteralPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralPatternContext {
	var p = new(LiteralPatternContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *LiteralPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralPatternContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterLiteralPattern(s)
	}
}

func (s *LiteralPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitLiteralPattern(s)
	}
}

func (s *LiteralPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitLiteralPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayPatternContext struct {
	*PatternContext
}

func NewArrayPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayPatternContext {
	var p = new(ArrayPatternContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *ArrayPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayPatternContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLBRACE, 0)
}

func (s *ArrayPatternContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRBRACE, 0)
}

func (s *ArrayPatternContext) AllPattern() []IPatternContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPatternContext)(nil)).Elem())
	var tst = make([]IPatternContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPatternContext)
		}
	}

	return tst
}

func (s *ArrayPatternContext) Pattern(i int) IPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *ArrayPatternContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *ArrayPatternContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *ArrayPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterArrayPattern(s)
	}
}

func (s *ArrayPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitArrayPattern(s)
	}
}

func (s *ArrayPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitArrayPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructPatternContext struct {
	*PatternContext
}

func NewStructPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructPatternContext {
	var p = new(StructPatternContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *StructPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructPatternContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserLCURLY, 0)
}

func (s *StructPatternContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserRCURLY, 0)
}

func (s *StructPatternContext) AllPattern() []IPatternContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPatternContext)(nil)).Elem())
	var tst = make([]IPatternContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPatternContext)
		}
	}

	return tst
}

func (s *StructPatternContext) Pattern(i int) IPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *StructPatternContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *StructPatternContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *StructPatternContext) AllTypeexpr() []ITypeexprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeexprContext)(nil)).Elem())
	var tst = make([]ITypeexprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeexprContext)
		}
	}

	return tst
}

func (s *StructPatternContext) Typeexpr(i int) ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *StructPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterStructPattern(s)
	}
}

func (s *StructPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitStructPattern(s)
	}
}

func (s *StructPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitStructPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtomicPatternContext struct {
	*PatternContext
}

func NewAtomicPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomicPatternContext {
	var p = new(AtomicPatternContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *AtomicPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomicPatternContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *AtomicPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterAtomicPattern(s)
	}
}

func (s *AtomicPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitAtomicPattern(s)
	}
}

func (s *AtomicPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitAtomicPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Pattern() (localctx IPatternContext) {
	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, QuarkParserRULE_pattern)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAtomicPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(463)
			p.Realname()
		}

	case 2:
		localctx = NewParamerterizedTypePatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(464)
			p.Realname()
		}
		{
			p.SetState(465)
			p.Match(QuarkParserLBRACE)
		}
		{
			p.SetState(466)
			p.Pattern()
		}
		p.SetState(471)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(467)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(468)
				p.Pattern()
			}

			p.SetState(473)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(474)
			p.Match(QuarkParserRBRACE)
		}

	case 3:
		localctx = NewArrayPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(476)
			p.Match(QuarkParserLBRACE)
		}
		p.SetState(485)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserLBRACE || _la == QuarkParserLCURLY || _la == QuarkParserINTEGRAL || _la == QuarkParserREAL_NAME {
			{
				p.SetState(477)
				p.Pattern()
			}
			p.SetState(482)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == QuarkParserCOMMA {
				{
					p.SetState(478)
					p.Match(QuarkParserCOMMA)
				}
				{
					p.SetState(479)
					p.Pattern()
				}

				p.SetState(484)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(487)
			p.Match(QuarkParserRBRACE)
		}

	case 4:
		localctx = NewLiteralPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(488)
			p.Literal()
		}

	case 5:
		localctx = NewStructPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(489)
			p.Match(QuarkParserLCURLY)
		}

		p.SetState(491)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(490)
				p.typeexpr(0)
			}

		}
		{
			p.SetState(493)
			p.Pattern()
		}

		p.SetState(502)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(495)
				p.Match(QuarkParserCOMMA)
			}
			p.SetState(497)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(496)
					p.typeexpr(0)
				}

			}
			{
				p.SetState(499)
				p.Pattern()
			}

			p.SetState(504)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(505)
			p.Match(QuarkParserRCURLY)
		}

	}

	return localctx
}

// IParameterlistContext is an interface to support dynamic dispatch.
type IParameterlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterlistContext differentiates from other interfaces.
	IsParameterlistContext()
}

type ParameterlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterlistContext() *ParameterlistContext {
	var p = new(ParameterlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_parameterlist
	return p
}

func (*ParameterlistContext) IsParameterlistContext() {}

func NewParameterlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterlistContext {
	var p = new(ParameterlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_parameterlist

	return p
}

func (s *ParameterlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterlistContext) LANGLE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLANGLE, 0)
}

func (s *ParameterlistContext) AllParameterdef() []IParameterdefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterdefContext)(nil)).Elem())
	var tst = make([]IParameterdefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterdefContext)
		}
	}

	return tst
}

func (s *ParameterlistContext) Parameterdef(i int) IParameterdefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterdefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterdefContext)
}

func (s *ParameterlistContext) RANGLE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRANGLE, 0)
}

func (s *ParameterlistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *ParameterlistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *ParameterlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterParameterlist(s)
	}
}

func (s *ParameterlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitParameterlist(s)
	}
}

func (s *ParameterlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitParameterlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Parameterlist() (localctx IParameterlistContext) {
	localctx = NewParameterlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, QuarkParserRULE_parameterlist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(509)
		p.Match(QuarkParserLANGLE)
	}
	{
		p.SetState(510)
		p.Parameterdef()
	}
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserCOMMA {
		{
			p.SetState(511)
			p.Match(QuarkParserCOMMA)
		}
		{
			p.SetState(512)
			p.Parameterdef()
		}

		p.SetState(517)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(518)
		p.Match(QuarkParserRANGLE)
	}

	return localctx
}

// IParameterdefContext is an interface to support dynamic dispatch.
type IParameterdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterdefContext differentiates from other interfaces.
	IsParameterdefContext()
}

type ParameterdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterdefContext() *ParameterdefContext {
	var p = new(ParameterdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_parameterdef
	return p
}

func (*ParameterdefContext) IsParameterdefContext() {}

func NewParameterdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterdefContext {
	var p = new(ParameterdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_parameterdef

	return p
}

func (s *ParameterdefContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterdefContext) CopyFrom(ctx *ParameterdefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ParameterdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValueParameterContext struct {
	*ParameterdefContext
}

func NewValueParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueParameterContext {
	var p = new(ValueParameterContext)

	p.ParameterdefContext = NewEmptyParameterdefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParameterdefContext))

	return p
}

func (s *ValueParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueParameterContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *ValueParameterContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *ValueParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterValueParameter(s)
	}
}

func (s *ValueParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitValueParameter(s)
	}
}

func (s *ValueParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitValueParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeParameterContext struct {
	*ParameterdefContext
}

func NewTypeParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParameterContext {
	var p = new(TypeParameterContext)

	p.ParameterdefContext = NewEmptyParameterdefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParameterdefContext))

	return p
}

func (s *TypeParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParameterContext) KW_TYPE() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_TYPE, 0)
}

func (s *TypeParameterContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *TypeParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterTypeParameter(s)
	}
}

func (s *TypeParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitTypeParameter(s)
	}
}

func (s *TypeParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitTypeParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Parameterdef() (localctx IParameterdefContext) {
	localctx = NewParameterdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, QuarkParserRULE_parameterdef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(525)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QuarkParserKW_TYPE:
		localctx = NewTypeParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(520)
			p.Match(QuarkParserKW_TYPE)
		}
		{
			p.SetState(521)
			p.typeexpr(0)
		}

	case QuarkParserREAL_NAME:
		localctx = NewValueParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(522)
			p.typeexpr(0)
		}
		{
			p.SetState(523)
			p.Realname()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReturnlistContext is an interface to support dynamic dispatch.
type IReturnlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnlistContext differentiates from other interfaces.
	IsReturnlistContext()
}

type ReturnlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnlistContext() *ReturnlistContext {
	var p = new(ReturnlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_returnlist
	return p
}

func (*ReturnlistContext) IsReturnlistContext() {}

func NewReturnlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnlistContext {
	var p = new(ReturnlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_returnlist

	return p
}

func (s *ReturnlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnlistContext) CopyFrom(ctx *ReturnlistContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ReturnlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NamedReturnContext struct {
	*ReturnlistContext
}

func NewNamedReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedReturnContext {
	var p = new(NamedReturnContext)

	p.ReturnlistContext = NewEmptyReturnlistContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ReturnlistContext))

	return p
}

func (s *NamedReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedReturnContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserLPAREN, 0)
}

func (s *NamedReturnContext) AllTypeexpr() []ITypeexprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeexprContext)(nil)).Elem())
	var tst = make([]ITypeexprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeexprContext)
		}
	}

	return tst
}

func (s *NamedReturnContext) Typeexpr(i int) ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *NamedReturnContext) AllRealname() []IRealnameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRealnameContext)(nil)).Elem())
	var tst = make([]IRealnameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRealnameContext)
		}
	}

	return tst
}

func (s *NamedReturnContext) Realname(i int) IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *NamedReturnContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
}

func (s *NamedReturnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *NamedReturnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *NamedReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterNamedReturn(s)
	}
}

func (s *NamedReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitNamedReturn(s)
	}
}

func (s *NamedReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitNamedReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleReturnContext struct {
	*ReturnlistContext
}

func NewSingleReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleReturnContext {
	var p = new(SingleReturnContext)

	p.ReturnlistContext = NewEmptyReturnlistContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ReturnlistContext))

	return p
}

func (s *SingleReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleReturnContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *SingleReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterSingleReturn(s)
	}
}

func (s *SingleReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitSingleReturn(s)
	}
}

func (s *SingleReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitSingleReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Returnlist() (localctx IReturnlistContext) {
	localctx = NewReturnlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, QuarkParserRULE_returnlist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(542)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QuarkParserREAL_NAME:
		localctx = NewSingleReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(527)
			p.typeexpr(0)
		}

	case QuarkParserLPAREN:
		localctx = NewNamedReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(528)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(529)
			p.typeexpr(0)
		}
		{
			p.SetState(530)
			p.Realname()
		}
		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(531)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(532)
				p.typeexpr(0)
			}
			{
				p.SetState(533)
				p.Realname()
			}

			p.SetState(539)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(540)
			p.Match(QuarkParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArgumentdefContext is an interface to support dynamic dispatch.
type IArgumentdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentdefContext differentiates from other interfaces.
	IsArgumentdefContext()
}

type ArgumentdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentdefContext() *ArgumentdefContext {
	var p = new(ArgumentdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_argumentdef
	return p
}

func (*ArgumentdefContext) IsArgumentdefContext() {}

func NewArgumentdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentdefContext {
	var p = new(ArgumentdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_argumentdef

	return p
}

func (s *ArgumentdefContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentdefContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *ArgumentdefContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *ArgumentdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterArgumentdef(s)
	}
}

func (s *ArgumentdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitArgumentdef(s)
	}
}

func (s *ArgumentdefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitArgumentdef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Argumentdef() (localctx IArgumentdefContext) {
	localctx = NewArgumentdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, QuarkParserRULE_argumentdef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(544)
		p.typeexpr(0)
	}
	{
		p.SetState(545)
		p.Realname()
	}

	return localctx
}

// IArgumentlistContext is an interface to support dynamic dispatch.
type IArgumentlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentlistContext differentiates from other interfaces.
	IsArgumentlistContext()
}

type ArgumentlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentlistContext() *ArgumentlistContext {
	var p = new(ArgumentlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_argumentlist
	return p
}

func (*ArgumentlistContext) IsArgumentlistContext() {}

func NewArgumentlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentlistContext {
	var p = new(ArgumentlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_argumentlist

	return p
}

func (s *ArgumentlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentlistContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserLPAREN, 0)
}

func (s *ArgumentlistContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
}

func (s *ArgumentlistContext) AllArgumentdef() []IArgumentdefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentdefContext)(nil)).Elem())
	var tst = make([]IArgumentdefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentdefContext)
		}
	}

	return tst
}

func (s *ArgumentlistContext) Argumentdef(i int) IArgumentdefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentdefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentdefContext)
}

func (s *ArgumentlistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *ArgumentlistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *ArgumentlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterArgumentlist(s)
	}
}

func (s *ArgumentlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitArgumentlist(s)
	}
}

func (s *ArgumentlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitArgumentlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Argumentlist() (localctx IArgumentlistContext) {
	localctx = NewArgumentlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, QuarkParserRULE_argumentlist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(547)
		p.Match(QuarkParserLPAREN)
	}
	p.SetState(556)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserREAL_NAME {
		{
			p.SetState(548)
			p.Argumentdef()
		}
		p.SetState(553)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(549)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(550)
				p.Argumentdef()
			}

			p.SetState(555)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(558)
		p.Match(QuarkParserRPAREN)
	}

	return localctx
}

// IStructdeclContext is an interface to support dynamic dispatch.
type IStructdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructdeclContext differentiates from other interfaces.
	IsStructdeclContext()
}

type StructdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructdeclContext() *StructdeclContext {
	var p = new(StructdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_structdecl
	return p
}

func (*StructdeclContext) IsStructdeclContext() {}

func NewStructdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructdeclContext {
	var p = new(StructdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_structdecl

	return p
}

func (s *StructdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructdeclContext) KW_STRUCT() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_STRUCT, 0)
}

func (s *StructdeclContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *StructdeclContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserLCURLY, 0)
}

func (s *StructdeclContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserRCURLY, 0)
}

func (s *StructdeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *StructdeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *StructdeclContext) Parameterlist() IParameterlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterlistContext)
}

func (s *StructdeclContext) KW_HAS() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_HAS, 0)
}

func (s *StructdeclContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *StructdeclContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *StructdeclContext) AllFielddecl() []IFielddeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFielddeclContext)(nil)).Elem())
	var tst = make([]IFielddeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFielddeclContext)
		}
	}

	return tst
}

func (s *StructdeclContext) Fielddecl(i int) IFielddeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFielddeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFielddeclContext)
}

func (s *StructdeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *StructdeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
}

func (s *StructdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterStructdecl(s)
	}
}

func (s *StructdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitStructdecl(s)
	}
}

func (s *StructdeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitStructdecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Structdecl() (localctx IStructdeclContext) {
	localctx = NewStructdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, QuarkParserRULE_structdecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(560)
			p.Annotation()
		}

		p.SetState(565)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(566)
		p.Match(QuarkParserKW_STRUCT)
	}
	{
		p.SetState(567)
		p.Realname()
	}
	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserLANGLE {
		{
			p.SetState(568)
			p.Parameterlist()
		}

	}
	p.SetState(580)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserKW_HAS {
		{
			p.SetState(571)
			p.Match(QuarkParserKW_HAS)
		}
		{
			p.SetState(572)
			p.Name()
		}
		p.SetState(577)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(573)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(574)
				p.Name()
			}

			p.SetState(579)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(582)
		p.Match(QuarkParserLCURLY)
	}
	p.SetState(586)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(QuarkParserKW_FUTURE-72))|(1<<(QuarkParserREAL_NAME-72))|(1<<(QuarkParserANNOTATION_START-72)))) != 0 {
		{
			p.SetState(583)
			p.Fielddecl()
		}

		p.SetState(588)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(589)
		p.Match(QuarkParserRCURLY)
	}

	return localctx
}

// IFielddeclContext is an interface to support dynamic dispatch.
type IFielddeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFielddeclContext differentiates from other interfaces.
	IsFielddeclContext()
}

type FielddeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFielddeclContext() *FielddeclContext {
	var p = new(FielddeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_fielddecl
	return p
}

func (*FielddeclContext) IsFielddeclContext() {}

func NewFielddeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FielddeclContext {
	var p = new(FielddeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_fielddecl

	return p
}

func (s *FielddeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FielddeclContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *FielddeclContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *FielddeclContext) SEMI() antlr.TerminalNode {
	return s.GetToken(QuarkParserSEMI, 0)
}

func (s *FielddeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *FielddeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *FielddeclContext) KW_FUTURE() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_FUTURE, 0)
}

func (s *FielddeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FielddeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FielddeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterFielddecl(s)
	}
}

func (s *FielddeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitFielddecl(s)
	}
}

func (s *FielddeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitFielddecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Fielddecl() (localctx IFielddeclContext) {
	localctx = NewFielddeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, QuarkParserRULE_fielddecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(594)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(591)
			p.Annotation()
		}

		p.SetState(596)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(598)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserKW_FUTURE {
		{
			p.SetState(597)
			p.Match(QuarkParserKW_FUTURE)
		}

	}
	{
		p.SetState(600)
		p.typeexpr(0)
	}
	{
		p.SetState(601)
		p.Realname()
	}
	{
		p.SetState(602)
		p.Match(QuarkParserSEMI)
	}

	return localctx
}

// IFuncdeclContext is an interface to support dynamic dispatch.
type IFuncdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncdeclContext differentiates from other interfaces.
	IsFuncdeclContext()
}

type FuncdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncdeclContext() *FuncdeclContext {
	var p = new(FuncdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_funcdecl
	return p
}

func (*FuncdeclContext) IsFuncdeclContext() {}

func NewFuncdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncdeclContext {
	var p = new(FuncdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_funcdecl

	return p
}

func (s *FuncdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncdeclContext) KW_DEF() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_DEF, 0)
}

func (s *FuncdeclContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *FuncdeclContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserLCURLY, 0)
}

func (s *FuncdeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncdeclContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserRCURLY, 0)
}

func (s *FuncdeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *FuncdeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *FuncdeclContext) Parameterlist() IParameterlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterlistContext)
}

func (s *FuncdeclContext) Argumentlist() IArgumentlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentlistContext)
}

func (s *FuncdeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(QuarkParserCOLON, 0)
}

func (s *FuncdeclContext) Returnlist() IReturnlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnlistContext)
}

func (s *FuncdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterFuncdecl(s)
	}
}

func (s *FuncdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitFuncdecl(s)
	}
}

func (s *FuncdeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitFuncdecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Funcdecl() (localctx IFuncdeclContext) {
	localctx = NewFuncdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, QuarkParserRULE_funcdecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(607)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(604)
			p.Annotation()
		}

		p.SetState(609)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(610)
		p.Match(QuarkParserKW_DEF)
	}
	{
		p.SetState(611)
		p.Realname()
	}
	p.SetState(613)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserLANGLE {
		{
			p.SetState(612)
			p.Parameterlist()
		}

	}
	p.SetState(616)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserLPAREN {
		{
			p.SetState(615)
			p.Argumentlist()
		}

	}
	p.SetState(620)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserCOLON {
		{
			p.SetState(618)
			p.Match(QuarkParserCOLON)
		}
		{
			p.SetState(619)
			p.Returnlist()
		}

	}
	{
		p.SetState(622)
		p.Match(QuarkParserLCURLY)
	}
	{
		p.SetState(623)
		p.Block()
	}
	{
		p.SetState(624)
		p.Match(QuarkParserRCURLY)
	}

	return localctx
}

// IModuledeclContext is an interface to support dynamic dispatch.
type IModuledeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuledeclContext differentiates from other interfaces.
	IsModuledeclContext()
}

type ModuledeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuledeclContext() *ModuledeclContext {
	var p = new(ModuledeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_moduledecl
	return p
}

func (*ModuledeclContext) IsModuledeclContext() {}

func NewModuledeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuledeclContext {
	var p = new(ModuledeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_moduledecl

	return p
}

func (s *ModuledeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuledeclContext) KW_MODULE() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_MODULE, 0)
}

func (s *ModuledeclContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *ModuledeclContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserLCURLY, 0)
}

func (s *ModuledeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ModuledeclContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserRCURLY, 0)
}

func (s *ModuledeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *ModuledeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *ModuledeclContext) Parameterlist() IParameterlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterlistContext)
}

func (s *ModuledeclContext) Argumentlist() IArgumentlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentlistContext)
}

func (s *ModuledeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(QuarkParserCOLON, 0)
}

func (s *ModuledeclContext) Returnlist() IReturnlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnlistContext)
}

func (s *ModuledeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuledeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuledeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterModuledecl(s)
	}
}

func (s *ModuledeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitModuledecl(s)
	}
}

func (s *ModuledeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitModuledecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Moduledecl() (localctx IModuledeclContext) {
	localctx = NewModuledeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, QuarkParserRULE_moduledecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(629)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(626)
			p.Annotation()
		}

		p.SetState(631)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(632)
		p.Match(QuarkParserKW_MODULE)
	}
	{
		p.SetState(633)
		p.Realname()
	}
	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserLANGLE {
		{
			p.SetState(634)
			p.Parameterlist()
		}

	}
	p.SetState(638)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserLPAREN {
		{
			p.SetState(637)
			p.Argumentlist()
		}

	}
	p.SetState(642)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserCOLON {
		{
			p.SetState(640)
			p.Match(QuarkParserCOLON)
		}
		{
			p.SetState(641)
			p.Returnlist()
		}

	}
	{
		p.SetState(644)
		p.Match(QuarkParserLCURLY)
	}
	{
		p.SetState(645)
		p.Block()
	}
	{
		p.SetState(646)
		p.Match(QuarkParserRCURLY)
	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) ANNOTATION_START() antlr.TerminalNode {
	return s.GetToken(QuarkParserANNOTATION_START, 0)
}

func (s *AnnotationContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (s *AnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, QuarkParserRULE_annotation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(648)
		p.Match(QuarkParserANNOTATION_START)
	}
	{
		p.SetState(649)
		p.Realname()
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INTEGRAL() antlr.TerminalNode {
	return s.GetToken(QuarkParserINTEGRAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, QuarkParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(651)
		p.Match(QuarkParserINTEGRAL)
	}

	return localctx
}

func (p *QuarkParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *AssignableContext = nil
		if localctx != nil {
			t = localctx.(*AssignableContext)
		}
		return p.Assignable_Sempred(t, predIndex)

	case 10:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 16:
		var t *TypeexprContext = nil
		if localctx != nil {
			t = localctx.(*TypeexprContext)
		}
		return p.Typeexpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *QuarkParser) Assignable_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *QuarkParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *QuarkParser) Typeexpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
