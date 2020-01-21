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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 98, 638,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 7, 2, 64, 10, 2, 12, 2, 14,
	2, 67, 11, 2, 3, 2, 7, 2, 70, 10, 2, 12, 2, 14, 2, 73, 11, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 5, 3, 80, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 6,
	4, 99, 10, 4, 13, 4, 14, 4, 100, 3, 4, 3, 4, 3, 4, 5, 4, 106, 10, 4, 3,
	5, 7, 5, 109, 10, 5, 12, 5, 14, 5, 112, 11, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 126, 10, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 147, 10, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 7, 8, 155, 10, 8, 12, 8, 14, 8, 158, 11, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 5, 8, 165, 10, 8, 3, 8, 3, 8, 5, 8, 169, 10, 8, 3,
	8, 3, 8, 5, 8, 173, 10, 8, 5, 8, 175, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5,
	8, 181, 10, 8, 3, 8, 3, 8, 3, 8, 5, 8, 186, 10, 8, 3, 8, 3, 8, 3, 8, 6,
	8, 191, 10, 8, 13, 8, 14, 8, 192, 7, 8, 195, 10, 8, 12, 8, 14, 8, 198,
	11, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 206, 10, 10, 13,
	10, 14, 10, 207, 5, 10, 210, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 225, 10,
	12, 13, 12, 14, 12, 226, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 5, 12, 259, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 273, 10, 12, 12, 12,
	14, 12, 276, 11, 12, 5, 12, 278, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 5, 12, 286, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 320, 10, 12, 3, 12, 3, 12, 5,
	12, 324, 10, 12, 3, 12, 5, 12, 327, 10, 12, 3, 12, 5, 12, 330, 10, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 338, 10, 12, 12, 12, 14,
	12, 341, 11, 12, 3, 12, 3, 12, 7, 12, 345, 10, 12, 12, 12, 14, 12, 348,
	11, 12, 3, 13, 3, 13, 3, 13, 7, 13, 353, 10, 13, 12, 13, 14, 13, 356, 11,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 363, 10, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 6, 15, 369, 10, 15, 13, 15, 14, 15, 370, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 381, 10, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 391, 10, 17, 12, 17, 14,
	17, 394, 11, 17, 3, 17, 3, 17, 7, 17, 398, 10, 17, 12, 17, 14, 17, 401,
	11, 17, 3, 18, 3, 18, 3, 18, 5, 18, 406, 10, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 419, 10, 19,
	12, 19, 14, 19, 422, 11, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19,
	429, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 6, 19, 440, 10, 19, 13, 19, 14, 19, 441, 3, 19, 3, 19, 5, 19, 446,
	10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 454, 10, 20, 12,
	20, 14, 20, 457, 11, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20,
	465, 10, 20, 12, 20, 14, 20, 468, 11, 20, 5, 20, 470, 10, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 5, 20, 476, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20,
	482, 10, 20, 3, 20, 7, 20, 485, 10, 20, 12, 20, 14, 20, 488, 11, 20, 3,
	20, 3, 20, 5, 20, 492, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 498,
	10, 21, 12, 21, 14, 21, 501, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 5, 22, 510, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 7, 23, 520, 10, 23, 12, 23, 14, 23, 523, 11, 23, 3, 23,
	3, 23, 5, 23, 527, 10, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3,
	25, 7, 25, 536, 10, 25, 12, 25, 14, 25, 539, 11, 25, 5, 25, 541, 10, 25,
	3, 25, 3, 25, 3, 26, 7, 26, 546, 10, 26, 12, 26, 14, 26, 549, 11, 26, 3,
	26, 3, 26, 3, 26, 5, 26, 554, 10, 26, 3, 26, 3, 26, 3, 26, 3, 26, 7, 26,
	560, 10, 26, 12, 26, 14, 26, 563, 11, 26, 5, 26, 565, 10, 26, 3, 26, 3,
	26, 7, 26, 569, 10, 26, 12, 26, 14, 26, 572, 11, 26, 3, 26, 3, 26, 3, 27,
	7, 27, 577, 10, 27, 12, 27, 14, 27, 580, 11, 27, 3, 27, 5, 27, 583, 10,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 7, 28, 590, 10, 28, 12, 28, 14,
	28, 593, 11, 28, 3, 28, 3, 28, 3, 28, 5, 28, 598, 10, 28, 3, 28, 5, 28,
	601, 10, 28, 3, 28, 3, 28, 5, 28, 605, 10, 28, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 29, 7, 29, 612, 10, 29, 12, 29, 14, 29, 615, 11, 29, 3, 29, 3, 29,
	3, 29, 5, 29, 620, 10, 29, 3, 29, 5, 29, 623, 10, 29, 3, 29, 3, 29, 5,
	29, 627, 10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31,
	3, 31, 3, 31, 2, 5, 14, 22, 32, 32, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 2, 8, 15, 2, 16, 16, 19, 19, 21, 21, 23, 23, 25, 25, 27, 27, 29,
	29, 32, 32, 35, 35, 38, 38, 40, 40, 42, 42, 44, 44, 5, 2, 22, 22, 24, 24,
	26, 26, 4, 2, 18, 18, 20, 20, 6, 2, 37, 37, 39, 39, 41, 41, 43, 43, 6,
	2, 28, 28, 30, 31, 33, 34, 36, 36, 4, 2, 46, 47, 49, 50, 2, 709, 2, 65,
	3, 2, 2, 2, 4, 79, 3, 2, 2, 2, 6, 105, 3, 2, 2, 2, 8, 110, 3, 2, 2, 2,
	10, 113, 3, 2, 2, 2, 12, 146, 3, 2, 2, 2, 14, 185, 3, 2, 2, 2, 16, 199,
	3, 2, 2, 2, 18, 209, 3, 2, 2, 2, 20, 211, 3, 2, 2, 2, 22, 285, 3, 2, 2,
	2, 24, 349, 3, 2, 2, 2, 26, 362, 3, 2, 2, 2, 28, 364, 3, 2, 2, 2, 30, 380,
	3, 2, 2, 2, 32, 382, 3, 2, 2, 2, 34, 405, 3, 2, 2, 2, 36, 445, 3, 2, 2,
	2, 38, 491, 3, 2, 2, 2, 40, 493, 3, 2, 2, 2, 42, 509, 3, 2, 2, 2, 44, 526,
	3, 2, 2, 2, 46, 528, 3, 2, 2, 2, 48, 531, 3, 2, 2, 2, 50, 547, 3, 2, 2,
	2, 52, 578, 3, 2, 2, 2, 54, 591, 3, 2, 2, 2, 56, 613, 3, 2, 2, 2, 58, 632,
	3, 2, 2, 2, 60, 635, 3, 2, 2, 2, 62, 64, 5, 6, 4, 2, 63, 62, 3, 2, 2, 2,
	64, 67, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 71, 3,
	2, 2, 2, 67, 65, 3, 2, 2, 2, 68, 70, 5, 4, 3, 2, 69, 68, 3, 2, 2, 2, 70,
	73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 74, 3, 2, 2,
	2, 73, 71, 3, 2, 2, 2, 74, 75, 7, 2, 2, 3, 75, 3, 3, 2, 2, 2, 76, 80, 5,
	50, 26, 2, 77, 80, 5, 54, 28, 2, 78, 80, 5, 56, 29, 2, 79, 76, 3, 2, 2,
	2, 79, 77, 3, 2, 2, 2, 79, 78, 3, 2, 2, 2, 80, 5, 3, 2, 2, 2, 81, 82, 7,
	69, 2, 2, 82, 83, 5, 18, 10, 2, 83, 84, 7, 4, 2, 2, 84, 106, 3, 2, 2, 2,
	85, 86, 7, 69, 2, 2, 86, 87, 5, 18, 10, 2, 87, 88, 7, 6, 2, 2, 88, 89,
	7, 22, 2, 2, 89, 90, 7, 4, 2, 2, 90, 106, 3, 2, 2, 2, 91, 92, 7, 69, 2,
	2, 92, 93, 5, 18, 10, 2, 93, 94, 7, 6, 2, 2, 94, 95, 7, 8, 2, 2, 95, 98,
	5, 16, 9, 2, 96, 97, 7, 5, 2, 2, 97, 99, 5, 16, 9, 2, 98, 96, 3, 2, 2,
	2, 99, 100, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101,
	102, 3, 2, 2, 2, 102, 103, 7, 9, 2, 2, 103, 104, 7, 4, 2, 2, 104, 106,
	3, 2, 2, 2, 105, 81, 3, 2, 2, 2, 105, 85, 3, 2, 2, 2, 105, 91, 3, 2, 2,
	2, 106, 7, 3, 2, 2, 2, 107, 109, 5, 12, 7, 2, 108, 107, 3, 2, 2, 2, 109,
	112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 9, 3,
	2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 114, 9, 2, 2, 2, 114, 11, 3, 2, 2,
	2, 115, 116, 5, 14, 8, 2, 116, 117, 5, 10, 6, 2, 117, 118, 5, 22, 12, 2,
	118, 119, 7, 4, 2, 2, 119, 147, 3, 2, 2, 2, 120, 121, 7, 79, 2, 2, 121,
	122, 7, 8, 2, 2, 122, 125, 5, 20, 11, 2, 123, 124, 7, 5, 2, 2, 124, 126,
	5, 20, 11, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 3,
	2, 2, 2, 127, 128, 7, 9, 2, 2, 128, 129, 5, 14, 8, 2, 129, 130, 5, 10,
	6, 2, 130, 131, 5, 22, 12, 2, 131, 132, 7, 4, 2, 2, 132, 147, 3, 2, 2,
	2, 133, 134, 7, 74, 2, 2, 134, 135, 5, 32, 17, 2, 135, 136, 5, 16, 9, 2,
	136, 137, 7, 4, 2, 2, 137, 147, 3, 2, 2, 2, 138, 139, 5, 14, 8, 2, 139,
	140, 7, 4, 2, 2, 140, 147, 3, 2, 2, 2, 141, 147, 5, 36, 19, 2, 142, 143,
	7, 68, 2, 2, 143, 144, 5, 22, 12, 2, 144, 145, 7, 4, 2, 2, 145, 147, 3,
	2, 2, 2, 146, 115, 3, 2, 2, 2, 146, 120, 3, 2, 2, 2, 146, 133, 3, 2, 2,
	2, 146, 138, 3, 2, 2, 2, 146, 141, 3, 2, 2, 2, 146, 142, 3, 2, 2, 2, 147,
	13, 3, 2, 2, 2, 148, 149, 8, 8, 1, 2, 149, 150, 5, 22, 12, 2, 150, 151,
	7, 10, 2, 2, 151, 156, 5, 22, 12, 2, 152, 153, 7, 5, 2, 2, 153, 155, 5,
	22, 12, 2, 154, 152, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2,
	2, 2, 156, 157, 3, 2, 2, 2, 157, 159, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2,
	159, 160, 7, 11, 2, 2, 160, 186, 3, 2, 2, 2, 161, 162, 5, 22, 12, 2, 162,
	164, 7, 10, 2, 2, 163, 165, 5, 22, 12, 2, 164, 163, 3, 2, 2, 2, 164, 165,
	3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 168, 7, 3, 2, 2, 167, 169, 5, 22,
	12, 2, 168, 167, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 174, 3, 2, 2, 2,
	170, 172, 7, 3, 2, 2, 171, 173, 5, 22, 12, 2, 172, 171, 3, 2, 2, 2, 172,
	173, 3, 2, 2, 2, 173, 175, 3, 2, 2, 2, 174, 170, 3, 2, 2, 2, 174, 175,
	3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 177, 7, 11, 2, 2, 177, 186, 3, 2,
	2, 2, 178, 186, 5, 18, 10, 2, 179, 181, 7, 84, 2, 2, 180, 179, 3, 2, 2,
	2, 180, 181, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 183, 5, 32, 17, 2,
	183, 184, 5, 16, 9, 2, 184, 186, 3, 2, 2, 2, 185, 148, 3, 2, 2, 2, 185,
	161, 3, 2, 2, 2, 185, 178, 3, 2, 2, 2, 185, 180, 3, 2, 2, 2, 186, 196,
	3, 2, 2, 2, 187, 190, 12, 3, 2, 2, 188, 189, 7, 5, 2, 2, 189, 191, 5, 14,
	8, 2, 190, 188, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2,
	192, 193, 3, 2, 2, 2, 193, 195, 3, 2, 2, 2, 194, 187, 3, 2, 2, 2, 195,
	198, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 15, 3,
	2, 2, 2, 198, 196, 3, 2, 2, 2, 199, 200, 7, 92, 2, 2, 200, 17, 3, 2, 2,
	2, 201, 210, 5, 16, 9, 2, 202, 205, 7, 92, 2, 2, 203, 204, 7, 6, 2, 2,
	204, 206, 7, 92, 2, 2, 205, 203, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207,
	205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 210, 3, 2, 2, 2, 209, 201,
	3, 2, 2, 2, 209, 202, 3, 2, 2, 2, 210, 19, 3, 2, 2, 2, 211, 212, 5, 18,
	10, 2, 212, 21, 3, 2, 2, 2, 213, 214, 8, 12, 1, 2, 214, 286, 5, 60, 31,
	2, 215, 286, 5, 18, 10, 2, 216, 217, 7, 8, 2, 2, 217, 218, 5, 22, 12, 2,
	218, 219, 7, 9, 2, 2, 219, 286, 3, 2, 2, 2, 220, 221, 7, 8, 2, 2, 221,
	224, 5, 22, 12, 2, 222, 223, 7, 5, 2, 2, 223, 225, 5, 22, 12, 2, 224, 222,
	3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2,
	2, 2, 227, 228, 3, 2, 2, 2, 228, 229, 7, 9, 2, 2, 229, 286, 3, 2, 2, 2,
	230, 231, 7, 14, 2, 2, 231, 232, 5, 24, 13, 2, 232, 233, 7, 15, 2, 2, 233,
	286, 3, 2, 2, 2, 234, 235, 7, 81, 2, 2, 235, 236, 5, 32, 17, 2, 236, 237,
	7, 8, 2, 2, 237, 238, 5, 24, 13, 2, 238, 239, 7, 9, 2, 2, 239, 286, 3,
	2, 2, 2, 240, 241, 7, 82, 2, 2, 241, 242, 5, 32, 17, 2, 242, 243, 7, 8,
	2, 2, 243, 244, 5, 24, 13, 2, 244, 245, 7, 9, 2, 2, 245, 286, 3, 2, 2,
	2, 246, 247, 7, 83, 2, 2, 247, 248, 5, 22, 12, 2, 248, 249, 7, 8, 2, 2,
	249, 250, 5, 24, 13, 2, 250, 251, 7, 9, 2, 2, 251, 286, 3, 2, 2, 2, 252,
	253, 7, 76, 2, 2, 253, 254, 5, 48, 25, 2, 254, 255, 7, 17, 2, 2, 255, 256,
	7, 14, 2, 2, 256, 258, 5, 8, 5, 2, 257, 259, 5, 22, 12, 2, 258, 257, 3,
	2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 261, 7, 15, 2,
	2, 261, 286, 3, 2, 2, 2, 262, 263, 7, 45, 2, 2, 263, 286, 5, 22, 12, 16,
	264, 265, 7, 48, 2, 2, 265, 286, 5, 22, 12, 15, 266, 286, 5, 28, 15, 2,
	267, 286, 5, 36, 19, 2, 268, 277, 7, 10, 2, 2, 269, 274, 5, 22, 12, 2,
	270, 271, 7, 5, 2, 2, 271, 273, 5, 22, 12, 2, 272, 270, 3, 2, 2, 2, 273,
	276, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 278,
	3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 277, 269, 3, 2, 2, 2, 277, 278, 3, 2,
	2, 2, 278, 279, 3, 2, 2, 2, 279, 286, 7, 11, 2, 2, 280, 281, 7, 85, 2,
	2, 281, 282, 7, 8, 2, 2, 282, 283, 5, 20, 11, 2, 283, 284, 7, 9, 2, 2,
	284, 286, 3, 2, 2, 2, 285, 213, 3, 2, 2, 2, 285, 215, 3, 2, 2, 2, 285,
	216, 3, 2, 2, 2, 285, 220, 3, 2, 2, 2, 285, 230, 3, 2, 2, 2, 285, 234,
	3, 2, 2, 2, 285, 240, 3, 2, 2, 2, 285, 246, 3, 2, 2, 2, 285, 252, 3, 2,
	2, 2, 285, 262, 3, 2, 2, 2, 285, 264, 3, 2, 2, 2, 285, 266, 3, 2, 2, 2,
	285, 267, 3, 2, 2, 2, 285, 268, 3, 2, 2, 2, 285, 280, 3, 2, 2, 2, 286,
	346, 3, 2, 2, 2, 287, 288, 12, 13, 2, 2, 288, 289, 9, 3, 2, 2, 289, 345,
	5, 22, 12, 14, 290, 291, 12, 12, 2, 2, 291, 292, 9, 4, 2, 2, 292, 345,
	5, 22, 12, 13, 293, 294, 12, 11, 2, 2, 294, 295, 9, 5, 2, 2, 295, 345,
	5, 22, 12, 12, 296, 297, 12, 10, 2, 2, 297, 298, 9, 6, 2, 2, 298, 345,
	5, 22, 12, 11, 299, 300, 12, 9, 2, 2, 300, 301, 9, 7, 2, 2, 301, 345, 5,
	22, 12, 10, 302, 303, 12, 8, 2, 2, 303, 304, 7, 57, 2, 2, 304, 305, 5,
	22, 12, 2, 305, 306, 7, 58, 2, 2, 306, 307, 5, 22, 12, 9, 307, 345, 3,
	2, 2, 2, 308, 309, 12, 25, 2, 2, 309, 310, 7, 6, 2, 2, 310, 345, 5, 16,
	9, 2, 311, 312, 12, 18, 2, 2, 312, 313, 7, 8, 2, 2, 313, 314, 5, 24, 13,
	2, 314, 315, 7, 9, 2, 2, 315, 345, 3, 2, 2, 2, 316, 317, 12, 6, 2, 2, 317,
	319, 7, 10, 2, 2, 318, 320, 5, 22, 12, 2, 319, 318, 3, 2, 2, 2, 319, 320,
	3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321, 323, 7, 3, 2, 2, 322, 324, 5, 22,
	12, 2, 323, 322, 3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324, 326, 3, 2, 2, 2,
	325, 327, 7, 3, 2, 2, 326, 325, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327,
	329, 3, 2, 2, 2, 328, 330, 5, 22, 12, 2, 329, 328, 3, 2, 2, 2, 329, 330,
	3, 2, 2, 2, 330, 331, 3, 2, 2, 2, 331, 345, 7, 11, 2, 2, 332, 333, 12,
	5, 2, 2, 333, 334, 7, 10, 2, 2, 334, 339, 5, 22, 12, 2, 335, 336, 7, 5,
	2, 2, 336, 338, 5, 22, 12, 2, 337, 335, 3, 2, 2, 2, 338, 341, 3, 2, 2,
	2, 339, 337, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 342, 3, 2, 2, 2, 341,
	339, 3, 2, 2, 2, 342, 343, 7, 11, 2, 2, 343, 345, 3, 2, 2, 2, 344, 287,
	3, 2, 2, 2, 344, 290, 3, 2, 2, 2, 344, 293, 3, 2, 2, 2, 344, 296, 3, 2,
	2, 2, 344, 299, 3, 2, 2, 2, 344, 302, 3, 2, 2, 2, 344, 308, 3, 2, 2, 2,
	344, 311, 3, 2, 2, 2, 344, 316, 3, 2, 2, 2, 344, 332, 3, 2, 2, 2, 345,
	348, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 23, 3,
	2, 2, 2, 348, 346, 3, 2, 2, 2, 349, 354, 5, 26, 14, 2, 350, 351, 7, 5,
	2, 2, 351, 353, 5, 26, 14, 2, 352, 350, 3, 2, 2, 2, 353, 356, 3, 2, 2,
	2, 354, 352, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 25, 3, 2, 2, 2, 356,
	354, 3, 2, 2, 2, 357, 358, 5, 16, 9, 2, 358, 359, 7, 16, 2, 2, 359, 360,
	5, 22, 12, 2, 360, 363, 3, 2, 2, 2, 361, 363, 5, 22, 12, 2, 362, 357, 3,
	2, 2, 2, 362, 361, 3, 2, 2, 2, 363, 27, 3, 2, 2, 2, 364, 365, 7, 14, 2,
	2, 365, 368, 5, 30, 16, 2, 366, 367, 7, 5, 2, 2, 367, 369, 5, 30, 16, 2,
	368, 366, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 368, 3, 2, 2, 2, 370,
	371, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 373, 7, 15, 2, 2, 373, 29,
	3, 2, 2, 2, 374, 375, 5, 22, 12, 2, 375, 376, 7, 14, 2, 2, 376, 377, 5,
	22, 12, 2, 377, 378, 7, 15, 2, 2, 378, 381, 3, 2, 2, 2, 379, 381, 5, 22,
	12, 2, 380, 374, 3, 2, 2, 2, 380, 379, 3, 2, 2, 2, 381, 31, 3, 2, 2, 2,
	382, 383, 8, 17, 1, 2, 383, 384, 5, 18, 10, 2, 384, 399, 3, 2, 2, 2, 385,
	386, 12, 3, 2, 2, 386, 387, 7, 10, 2, 2, 387, 392, 5, 34, 18, 2, 388, 389,
	7, 5, 2, 2, 389, 391, 5, 34, 18, 2, 390, 388, 3, 2, 2, 2, 391, 394, 3,
	2, 2, 2, 392, 390, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2, 393, 395, 3, 2, 2,
	2, 394, 392, 3, 2, 2, 2, 395, 396, 7, 11, 2, 2, 396, 398, 3, 2, 2, 2, 397,
	385, 3, 2, 2, 2, 398, 401, 3, 2, 2, 2, 399, 397, 3, 2, 2, 2, 399, 400,
	3, 2, 2, 2, 400, 33, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2, 402, 403, 7, 65,
	2, 2, 403, 406, 5, 32, 17, 2, 404, 406, 5, 22, 12, 2, 405, 402, 3, 2, 2,
	2, 405, 404, 3, 2, 2, 2, 406, 35, 3, 2, 2, 2, 407, 408, 7, 57, 2, 2, 408,
	409, 5, 22, 12, 2, 409, 410, 7, 14, 2, 2, 410, 411, 5, 8, 5, 2, 411, 420,
	7, 15, 2, 2, 412, 413, 7, 59, 2, 2, 413, 414, 5, 22, 12, 2, 414, 415, 7,
	14, 2, 2, 415, 416, 5, 8, 5, 2, 416, 417, 7, 15, 2, 2, 417, 419, 3, 2,
	2, 2, 418, 412, 3, 2, 2, 2, 419, 422, 3, 2, 2, 2, 420, 418, 3, 2, 2, 2,
	420, 421, 3, 2, 2, 2, 421, 428, 3, 2, 2, 2, 422, 420, 3, 2, 2, 2, 423,
	424, 7, 58, 2, 2, 424, 425, 7, 14, 2, 2, 425, 426, 5, 8, 5, 2, 426, 427,
	7, 15, 2, 2, 427, 429, 3, 2, 2, 2, 428, 423, 3, 2, 2, 2, 428, 429, 3, 2,
	2, 2, 429, 446, 3, 2, 2, 2, 430, 431, 7, 72, 2, 2, 431, 432, 5, 22, 12,
	2, 432, 439, 7, 14, 2, 2, 433, 434, 7, 73, 2, 2, 434, 435, 5, 38, 20, 2,
	435, 436, 7, 14, 2, 2, 436, 437, 5, 8, 5, 2, 437, 438, 7, 15, 2, 2, 438,
	440, 3, 2, 2, 2, 439, 433, 3, 2, 2, 2, 440, 441, 3, 2, 2, 2, 441, 439,
	3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 443, 3, 2, 2, 2, 443, 444, 7, 15,
	2, 2, 444, 446, 3, 2, 2, 2, 445, 407, 3, 2, 2, 2, 445, 430, 3, 2, 2, 2,
	446, 37, 3, 2, 2, 2, 447, 492, 5, 16, 9, 2, 448, 449, 5, 16, 9, 2, 449,
	450, 7, 10, 2, 2, 450, 455, 5, 38, 20, 2, 451, 452, 7, 5, 2, 2, 452, 454,
	5, 38, 20, 2, 453, 451, 3, 2, 2, 2, 454, 457, 3, 2, 2, 2, 455, 453, 3,
	2, 2, 2, 455, 456, 3, 2, 2, 2, 456, 458, 3, 2, 2, 2, 457, 455, 3, 2, 2,
	2, 458, 459, 7, 11, 2, 2, 459, 492, 3, 2, 2, 2, 460, 469, 7, 10, 2, 2,
	461, 466, 5, 38, 20, 2, 462, 463, 7, 5, 2, 2, 463, 465, 5, 38, 20, 2, 464,
	462, 3, 2, 2, 2, 465, 468, 3, 2, 2, 2, 466, 464, 3, 2, 2, 2, 466, 467,
	3, 2, 2, 2, 467, 470, 3, 2, 2, 2, 468, 466, 3, 2, 2, 2, 469, 461, 3, 2,
	2, 2, 469, 470, 3, 2, 2, 2, 470, 471, 3, 2, 2, 2, 471, 492, 7, 11, 2, 2,
	472, 492, 5, 60, 31, 2, 473, 475, 7, 14, 2, 2, 474, 476, 5, 32, 17, 2,
	475, 474, 3, 2, 2, 2, 475, 476, 3, 2, 2, 2, 476, 477, 3, 2, 2, 2, 477,
	478, 5, 38, 20, 2, 478, 486, 3, 2, 2, 2, 479, 481, 7, 5, 2, 2, 480, 482,
	5, 32, 17, 2, 481, 480, 3, 2, 2, 2, 481, 482, 3, 2, 2, 2, 482, 483, 3,
	2, 2, 2, 483, 485, 5, 38, 20, 2, 484, 479, 3, 2, 2, 2, 485, 488, 3, 2,
	2, 2, 486, 484, 3, 2, 2, 2, 486, 487, 3, 2, 2, 2, 487, 489, 3, 2, 2, 2,
	488, 486, 3, 2, 2, 2, 489, 490, 7, 15, 2, 2, 490, 492, 3, 2, 2, 2, 491,
	447, 3, 2, 2, 2, 491, 448, 3, 2, 2, 2, 491, 460, 3, 2, 2, 2, 491, 472,
	3, 2, 2, 2, 491, 473, 3, 2, 2, 2, 492, 39, 3, 2, 2, 2, 493, 494, 7, 10,
	2, 2, 494, 499, 5, 42, 22, 2, 495, 496, 7, 5, 2, 2, 496, 498, 5, 42, 22,
	2, 497, 495, 3, 2, 2, 2, 498, 501, 3, 2, 2, 2, 499, 497, 3, 2, 2, 2, 499,
	500, 3, 2, 2, 2, 500, 502, 3, 2, 2, 2, 501, 499, 3, 2, 2, 2, 502, 503,
	7, 11, 2, 2, 503, 41, 3, 2, 2, 2, 504, 505, 7, 65, 2, 2, 505, 510, 5, 32,
	17, 2, 506, 507, 5, 32, 17, 2, 507, 508, 5, 16, 9, 2, 508, 510, 3, 2, 2,
	2, 509, 504, 3, 2, 2, 2, 509, 506, 3, 2, 2, 2, 510, 43, 3, 2, 2, 2, 511,
	527, 5, 32, 17, 2, 512, 513, 7, 8, 2, 2, 513, 514, 5, 32, 17, 2, 514, 521,
	5, 16, 9, 2, 515, 516, 7, 5, 2, 2, 516, 517, 5, 32, 17, 2, 517, 518, 5,
	16, 9, 2, 518, 520, 3, 2, 2, 2, 519, 515, 3, 2, 2, 2, 520, 523, 3, 2, 2,
	2, 521, 519, 3, 2, 2, 2, 521, 522, 3, 2, 2, 2, 522, 524, 3, 2, 2, 2, 523,
	521, 3, 2, 2, 2, 524, 525, 7, 9, 2, 2, 525, 527, 3, 2, 2, 2, 526, 511,
	3, 2, 2, 2, 526, 512, 3, 2, 2, 2, 527, 45, 3, 2, 2, 2, 528, 529, 5, 32,
	17, 2, 529, 530, 5, 16, 9, 2, 530, 47, 3, 2, 2, 2, 531, 540, 7, 8, 2, 2,
	532, 537, 5, 46, 24, 2, 533, 534, 7, 5, 2, 2, 534, 536, 5, 46, 24, 2, 535,
	533, 3, 2, 2, 2, 536, 539, 3, 2, 2, 2, 537, 535, 3, 2, 2, 2, 537, 538,
	3, 2, 2, 2, 538, 541, 3, 2, 2, 2, 539, 537, 3, 2, 2, 2, 540, 532, 3, 2,
	2, 2, 540, 541, 3, 2, 2, 2, 541, 542, 3, 2, 2, 2, 542, 543, 7, 9, 2, 2,
	543, 49, 3, 2, 2, 2, 544, 546, 5, 58, 30, 2, 545, 544, 3, 2, 2, 2, 546,
	549, 3, 2, 2, 2, 547, 545, 3, 2, 2, 2, 547, 548, 3, 2, 2, 2, 548, 550,
	3, 2, 2, 2, 549, 547, 3, 2, 2, 2, 550, 551, 7, 63, 2, 2, 551, 553, 5, 16,
	9, 2, 552, 554, 5, 40, 21, 2, 553, 552, 3, 2, 2, 2, 553, 554, 3, 2, 2,
	2, 554, 564, 3, 2, 2, 2, 555, 556, 7, 62, 2, 2, 556, 561, 5, 18, 10, 2,
	557, 558, 7, 5, 2, 2, 558, 560, 5, 18, 10, 2, 559, 557, 3, 2, 2, 2, 560,
	563, 3, 2, 2, 2, 561, 559, 3, 2, 2, 2, 561, 562, 3, 2, 2, 2, 562, 565,
	3, 2, 2, 2, 563, 561, 3, 2, 2, 2, 564, 555, 3, 2, 2, 2, 564, 565, 3, 2,
	2, 2, 565, 566, 3, 2, 2, 2, 566, 570, 7, 14, 2, 2, 567, 569, 5, 52, 27,
	2, 568, 567, 3, 2, 2, 2, 569, 572, 3, 2, 2, 2, 570, 568, 3, 2, 2, 2, 570,
	571, 3, 2, 2, 2, 571, 573, 3, 2, 2, 2, 572, 570, 3, 2, 2, 2, 573, 574,
	7, 15, 2, 2, 574, 51, 3, 2, 2, 2, 575, 577, 5, 58, 30, 2, 576, 575, 3,
	2, 2, 2, 577, 580, 3, 2, 2, 2, 578, 576, 3, 2, 2, 2, 578, 579, 3, 2, 2,
	2, 579, 582, 3, 2, 2, 2, 580, 578, 3, 2, 2, 2, 581, 583, 7, 74, 2, 2, 582,
	581, 3, 2, 2, 2, 582, 583, 3, 2, 2, 2, 583, 584, 3, 2, 2, 2, 584, 585,
	5, 32, 17, 2, 585, 586, 5, 16, 9, 2, 586, 587, 7, 4, 2, 2, 587, 53, 3,
	2, 2, 2, 588, 590, 5, 58, 30, 2, 589, 588, 3, 2, 2, 2, 590, 593, 3, 2,
	2, 2, 591, 589, 3, 2, 2, 2, 591, 592, 3, 2, 2, 2, 592, 594, 3, 2, 2, 2,
	593, 591, 3, 2, 2, 2, 594, 595, 7, 56, 2, 2, 595, 597, 5, 16, 9, 2, 596,
	598, 5, 40, 21, 2, 597, 596, 3, 2, 2, 2, 597, 598, 3, 2, 2, 2, 598, 600,
	3, 2, 2, 2, 599, 601, 5, 48, 25, 2, 600, 599, 3, 2, 2, 2, 600, 601, 3,
	2, 2, 2, 601, 604, 3, 2, 2, 2, 602, 603, 7, 3, 2, 2, 603, 605, 5, 44, 23,
	2, 604, 602, 3, 2, 2, 2, 604, 605, 3, 2, 2, 2, 605, 606, 3, 2, 2, 2, 606,
	607, 7, 14, 2, 2, 607, 608, 5, 8, 5, 2, 608, 609, 7, 15, 2, 2, 609, 55,
	3, 2, 2, 2, 610, 612, 5, 58, 30, 2, 611, 610, 3, 2, 2, 2, 612, 615, 3,
	2, 2, 2, 613, 611, 3, 2, 2, 2, 613, 614, 3, 2, 2, 2, 614, 616, 3, 2, 2,
	2, 615, 613, 3, 2, 2, 2, 616, 617, 7, 55, 2, 2, 617, 619, 5, 16, 9, 2,
	618, 620, 5, 40, 21, 2, 619, 618, 3, 2, 2, 2, 619, 620, 3, 2, 2, 2, 620,
	622, 3, 2, 2, 2, 621, 623, 5, 48, 25, 2, 622, 621, 3, 2, 2, 2, 622, 623,
	3, 2, 2, 2, 623, 626, 3, 2, 2, 2, 624, 625, 7, 3, 2, 2, 625, 627, 5, 44,
	23, 2, 626, 624, 3, 2, 2, 2, 626, 627, 3, 2, 2, 2, 627, 628, 3, 2, 2, 2,
	628, 629, 7, 14, 2, 2, 629, 630, 5, 8, 5, 2, 630, 631, 7, 15, 2, 2, 631,
	57, 3, 2, 2, 2, 632, 633, 7, 93, 2, 2, 633, 634, 5, 16, 9, 2, 634, 59,
	3, 2, 2, 2, 635, 636, 7, 91, 2, 2, 636, 61, 3, 2, 2, 2, 72, 65, 71, 79,
	100, 105, 110, 125, 146, 156, 164, 168, 172, 174, 180, 185, 192, 196, 207,
	209, 226, 258, 274, 277, 285, 319, 323, 326, 329, 339, 344, 346, 354, 362,
	370, 380, 392, 399, 405, 420, 428, 441, 445, 455, 466, 469, 475, 481, 486,
	491, 499, 509, 521, 526, 537, 540, 547, 553, 561, 564, 570, 578, 582, 591,
	597, 600, 604, 613, 619, 622, 626,
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
	"realname", "name", "clockexpr", "expr", "callarglist", "callarg", "concat",
	"innerconcat", "typeexpr", "typeparam", "branch", "pattern", "parameterlist",
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
	QuarkParserRULE_concat        = 13
	QuarkParserRULE_innerconcat   = 14
	QuarkParserRULE_typeexpr      = 15
	QuarkParserRULE_typeparam     = 16
	QuarkParserRULE_branch        = 17
	QuarkParserRULE_pattern       = 18
	QuarkParserRULE_parameterlist = 19
	QuarkParserRULE_parameterdef  = 20
	QuarkParserRULE_returnlist    = 21
	QuarkParserRULE_argumentdef   = 22
	QuarkParserRULE_argumentlist  = 23
	QuarkParserRULE_structdecl    = 24
	QuarkParserRULE_fielddecl     = 25
	QuarkParserRULE_funcdecl      = 26
	QuarkParserRULE_moduledecl    = 27
	QuarkParserRULE_annotation    = 28
	QuarkParserRULE_literal       = 29
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
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserKW_IMPORT {
		{
			p.SetState(60)
			p.Importdecl()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(QuarkParserKW_MODULE-53))|(1<<(QuarkParserKW_DEF-53))|(1<<(QuarkParserKW_STRUCT-53)))) != 0) || _la == QuarkParserANNOTATION_START {
		{
			p.SetState(66)
			p.Decl()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Structdecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Funcdecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(76)
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

	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Match(QuarkParserKW_IMPORT)
		}
		{
			p.SetState(80)
			p.Name()
		}
		{
			p.SetState(81)
			p.Match(QuarkParserSEMI)
		}

	case 2:
		localctx = NewWildcardImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Match(QuarkParserKW_IMPORT)
		}
		{
			p.SetState(84)
			p.Name()
		}
		{
			p.SetState(85)
			p.Match(QuarkParserDOT)
		}
		{
			p.SetState(86)
			p.Match(QuarkParserOP_MUL)
		}
		{
			p.SetState(87)
			p.Match(QuarkParserSEMI)
		}

	case 3:
		localctx = NewMultiImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)
			p.Match(QuarkParserKW_IMPORT)
		}
		{
			p.SetState(90)
			p.Name()
		}
		{
			p.SetState(91)
			p.Match(QuarkParserDOT)
		}
		{
			p.SetState(92)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(93)
			p.Realname()
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QuarkParserCOMMA {
			{
				p.SetState(94)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(95)
				p.Realname()
			}

			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(100)
			p.Match(QuarkParserRPAREN)
		}
		{
			p.SetState(101)
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
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(105)
				p.Stmt()
			}

		}
		p.SetState(110)
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
		p.SetState(111)
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

	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.assignable(0)
		}
		{
			p.SetState(114)
			p.Assignment()
		}
		{
			p.SetState(115)
			p.expr(0)
		}
		{
			p.SetState(116)
			p.Match(QuarkParserSEMI)
		}

	case 2:
		localctx = NewRegAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(QuarkParserKW_REG)
		}
		{
			p.SetState(119)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(120)

			var _x = p.Clockexpr()

			localctx.(*RegAssignStmtContext).clk = _x
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserCOMMA {
			{
				p.SetState(121)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(122)

				var _x = p.Clockexpr()

				localctx.(*RegAssignStmtContext).rst = _x
			}

		}
		{
			p.SetState(125)
			p.Match(QuarkParserRPAREN)
		}
		{
			p.SetState(126)
			p.assignable(0)
		}
		{
			p.SetState(127)
			p.Assignment()
		}
		{
			p.SetState(128)
			p.expr(0)
		}
		{
			p.SetState(129)
			p.Match(QuarkParserSEMI)
		}

	case 3:
		localctx = NewFutureStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
			p.Match(QuarkParserKW_FUTURE)
		}
		{
			p.SetState(132)
			p.typeexpr(0)
		}
		{
			p.SetState(133)
			p.Realname()
		}
		{
			p.SetState(134)
			p.Match(QuarkParserSEMI)
		}

	case 4:
		localctx = NewDeclarationStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.assignable(0)
		}
		{
			p.SetState(137)
			p.Match(QuarkParserSEMI)
		}

	case 5:
		localctx = NewBranchStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)
			p.Branch()
		}

	case 6:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(140)
			p.Match(QuarkParserKW_RETURN)
		}
		{
			p.SetState(141)
			p.expr(0)
		}
		{
			p.SetState(142)
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
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArrayIndexAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(147)
			p.expr(0)
		}
		{
			p.SetState(148)
			p.Match(QuarkParserLBRACE)
		}
		{
			p.SetState(149)
			p.expr(0)
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(150)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(151)
				p.expr(0)
			}

			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(157)
			p.Match(QuarkParserRBRACE)
		}

	case 2:
		localctx = NewArraySliceAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(159)
			p.expr(0)
		}
		{
			p.SetState(160)
			p.Match(QuarkParserLBRACE)
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_OPEN-79))|(1<<(QuarkParserKW_CLOSE-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
			{
				p.SetState(161)

				var _x = p.expr(0)

				localctx.(*ArraySliceAssignmentContext).msb = _x
			}

		}
		{
			p.SetState(164)
			p.Match(QuarkParserCOLON)
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_OPEN-79))|(1<<(QuarkParserKW_CLOSE-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
			{
				p.SetState(165)

				var _x = p.expr(0)

				localctx.(*ArraySliceAssignmentContext).lsb = _x
			}

		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserCOLON {
			{
				p.SetState(168)
				p.Match(QuarkParserCOLON)
			}
			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_OPEN-79))|(1<<(QuarkParserKW_CLOSE-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
				{
					p.SetState(169)

					var _x = p.expr(0)

					localctx.(*ArraySliceAssignmentContext).step = _x
				}

			}

		}
		{
			p.SetState(174)
			p.Match(QuarkParserRBRACE)
		}

	case 3:
		localctx = NewValueAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(176)
			p.Name()
		}

	case 4:
		localctx = NewVariableDefinitionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserKW_MUT {
			{
				p.SetState(177)
				p.Match(QuarkParserKW_MUT)
			}

		}
		{
			p.SetState(180)
			p.typeexpr(0)
		}
		{
			p.SetState(181)
			p.Realname()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(194)
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
			p.SetState(185)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(186)
						p.Match(QuarkParserCOMMA)
					}
					{
						p.SetState(187)
						p.assignable(0)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(190)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
			}

		}
		p.SetState(196)
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
		p.SetState(197)
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

	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnqualifiedNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Realname()
		}

	case 2:
		localctx = NewQualifiedNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(200)
			p.Match(QuarkParserREAL_NAME)
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(201)
					p.Match(QuarkParserDOT)
				}
				{
					p.SetState(202)
					p.Match(QuarkParserREAL_NAME)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(205)
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
		p.SetState(209)
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

type OpenExprContext struct {
	*ExprContext
}

func NewOpenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpenExprContext {
	var p = new(OpenExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OpenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenExprContext) KW_OPEN() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_OPEN, 0)
}

func (s *OpenExprContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *OpenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserLPAREN, 0)
}

func (s *OpenExprContext) Callarglist() ICallarglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallarglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallarglistContext)
}

func (s *OpenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
}

func (s *OpenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterOpenExpr(s)
	}
}

func (s *OpenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitOpenExpr(s)
	}
}

func (s *OpenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitOpenExpr(s)

	default:
		return t.VisitChildren(s)
	}
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

type CloseExprContext struct {
	*ExprContext
}

func NewCloseExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CloseExprContext {
	var p = new(CloseExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CloseExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CloseExprContext) KW_CLOSE() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_CLOSE, 0)
}

func (s *CloseExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CloseExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserLPAREN, 0)
}

func (s *CloseExprContext) Callarglist() ICallarglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallarglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallarglistContext)
}

func (s *CloseExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
}

func (s *CloseExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterCloseExpr(s)
	}
}

func (s *CloseExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitCloseExpr(s)
	}
}

func (s *CloseExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitCloseExpr(s)

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

func (s *VarExprContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
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

type FieldExprContext struct {
	*ExprContext
}

func NewFieldExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldExprContext {
	var p = new(FieldExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FieldExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FieldExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(QuarkParserDOT, 0)
}

func (s *FieldExprContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *FieldExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterFieldExpr(s)
	}
}

func (s *FieldExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitFieldExpr(s)
	}
}

func (s *FieldExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitFieldExpr(s)

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
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(212)
			p.Literal()
		}

	case 2:
		localctx = NewVarExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(213)
			p.Name()
		}

	case 3:
		localctx = NewParensExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(214)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(215)
			p.expr(0)
		}
		{
			p.SetState(216)
			p.Match(QuarkParserRPAREN)
		}

	case 4:
		localctx = NewTupleExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(218)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(219)
			p.expr(0)
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QuarkParserCOMMA {
			{
				p.SetState(220)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(221)
				p.expr(0)
			}

			p.SetState(224)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(226)
			p.Match(QuarkParserRPAREN)
		}

	case 5:
		localctx = NewConstructorExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(228)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(229)
			p.Callarglist()
		}
		{
			p.SetState(230)
			p.Match(QuarkParserRCURLY)
		}

	case 6:
		localctx = NewNewModuleExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(232)
			p.Match(QuarkParserKW_NEW)
		}
		{
			p.SetState(233)
			p.typeexpr(0)
		}
		{
			p.SetState(234)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(235)
			p.Callarglist()
		}
		{
			p.SetState(236)
			p.Match(QuarkParserRPAREN)
		}

	case 7:
		localctx = NewOpenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Match(QuarkParserKW_OPEN)
		}
		{
			p.SetState(239)
			p.typeexpr(0)
		}
		{
			p.SetState(240)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(241)
			p.Callarglist()
		}
		{
			p.SetState(242)
			p.Match(QuarkParserRPAREN)
		}

	case 8:
		localctx = NewCloseExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(244)
			p.Match(QuarkParserKW_CLOSE)
		}
		{
			p.SetState(245)
			p.expr(0)
		}
		{
			p.SetState(246)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(247)
			p.Callarglist()
		}
		{
			p.SetState(248)
			p.Match(QuarkParserRPAREN)
		}

	case 9:
		localctx = NewLambdaExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(250)
			p.Match(QuarkParserKW_LAMBDA)
		}
		{
			p.SetState(251)
			p.Argumentlist()
		}
		{
			p.SetState(252)
			p.Match(QuarkParserOP_ARROW)
		}
		{
			p.SetState(253)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(254)
			p.Block()
		}
		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_OPEN-79))|(1<<(QuarkParserKW_CLOSE-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
			{
				p.SetState(255)
				p.expr(0)
			}

		}
		{
			p.SetState(258)
			p.Match(QuarkParserRCURLY)
		}

	case 10:
		localctx = NewComplimentExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(260)
			p.Match(QuarkParserOP_COMPLIMENT)
		}
		{
			p.SetState(261)
			p.expr(14)
		}

	case 11:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(262)
			p.Match(QuarkParserOP_LNOT)
		}
		{
			p.SetState(263)
			p.expr(13)
		}

	case 12:
		localctx = NewConcatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(264)
			p.Concat()
		}

	case 13:
		localctx = NewBranchExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(265)
			p.Branch()
		}

	case 14:
		localctx = NewArrayLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(266)
			p.Match(QuarkParserLBRACE)
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_OPEN-79))|(1<<(QuarkParserKW_CLOSE-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
			{
				p.SetState(267)
				p.expr(0)
			}
			p.SetState(272)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == QuarkParserCOMMA {
				{
					p.SetState(268)
					p.Match(QuarkParserCOMMA)
				}
				{
					p.SetState(269)
					p.expr(0)
				}

				p.SetState(274)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(277)
			p.Match(QuarkParserRBRACE)
		}

	case 15:
		localctx = NewClockToExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(278)
			p.Match(QuarkParserKW_SIGNAL)
		}
		{
			p.SetState(279)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(280)
			p.Clockexpr()
		}
		{
			p.SetState(281)
			p.Match(QuarkParserRPAREN)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(342)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivModExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(285)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(286)

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
					p.SetState(287)
					p.expr(12)
				}

			case 2:
				localctx = NewAddSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(289)

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
					p.SetState(290)
					p.expr(11)
				}

			case 3:
				localctx = NewShiftExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(291)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(292)

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
					p.SetState(293)
					p.expr(10)
				}

			case 4:
				localctx = NewBitwiseBinopExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(294)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(295)

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
					p.SetState(296)
					p.expr(9)
				}

			case 5:
				localctx = NewLogicalBinopExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(297)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(298)

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
					p.SetState(299)
					p.expr(8)
				}

			case 6:
				localctx = NewTernaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(300)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(301)
					p.Match(QuarkParserKW_IF)
				}
				{
					p.SetState(302)
					p.expr(0)
				}
				{
					p.SetState(303)
					p.Match(QuarkParserKW_ELSE)
				}
				{
					p.SetState(304)
					p.expr(7)
				}

			case 7:
				localctx = NewFieldExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(306)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
				}
				{
					p.SetState(307)
					p.Match(QuarkParserDOT)
				}
				{
					p.SetState(308)
					p.Realname()
				}

			case 8:
				localctx = NewFunctionCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(309)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(310)
					p.Match(QuarkParserLPAREN)
				}
				{
					p.SetState(311)
					p.Callarglist()
				}
				{
					p.SetState(312)
					p.Match(QuarkParserRPAREN)
				}

			case 9:
				localctx = NewSliceExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(314)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(315)
					p.Match(QuarkParserLBRACE)
				}
				p.SetState(317)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_OPEN-79))|(1<<(QuarkParserKW_CLOSE-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
					{
						p.SetState(316)

						var _x = p.expr(0)

						localctx.(*SliceExprContext).msb = _x
					}

				}
				{
					p.SetState(319)
					p.Match(QuarkParserCOLON)
				}
				p.SetState(321)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(320)

						var _x = p.expr(0)

						localctx.(*SliceExprContext).lsb = _x
					}

				}

				p.SetState(324)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == QuarkParserCOLON {
					{
						p.SetState(323)
						p.Match(QuarkParserCOLON)
					}

				}
				p.SetState(327)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QuarkParserOP_COMPLIMENT-43))|(1<<(QuarkParserOP_LNOT-43))|(1<<(QuarkParserKW_IF-43))|(1<<(QuarkParserKW_MATCH-43))|(1<<(QuarkParserKW_LAMBDA-43)))) != 0) || (((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(QuarkParserKW_NEW-79))|(1<<(QuarkParserKW_OPEN-79))|(1<<(QuarkParserKW_CLOSE-79))|(1<<(QuarkParserKW_SIGNAL-79))|(1<<(QuarkParserINTEGRAL-79))|(1<<(QuarkParserREAL_NAME-79)))) != 0) {
					{
						p.SetState(326)

						var _x = p.expr(0)

						localctx.(*SliceExprContext).step = _x
					}

				}

				{
					p.SetState(329)
					p.Match(QuarkParserRBRACE)
				}

			case 10:
				localctx = NewArrayIndexExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(330)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(331)
					p.Match(QuarkParserLBRACE)
				}
				{
					p.SetState(332)
					p.expr(0)
				}
				p.SetState(337)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == QuarkParserCOMMA {
					{
						p.SetState(333)
						p.Match(QuarkParserCOMMA)
					}
					{
						p.SetState(334)
						p.expr(0)
					}

					p.SetState(339)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(340)
					p.Match(QuarkParserRBRACE)
				}

			}

		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
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
		p.SetState(347)
		p.Callarg()
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserCOMMA {
		{
			p.SetState(348)
			p.Match(QuarkParserCOMMA)
		}
		{
			p.SetState(349)
			p.Callarg()
		}

		p.SetState(354)
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

	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNamedCallArgContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)
			p.Realname()
		}
		{
			p.SetState(356)
			p.Match(QuarkParserOP_ASSIGN)
		}
		{
			p.SetState(357)
			p.expr(0)
		}

	case 2:
		localctx = NewUnamedCallArgContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)
			p.expr(0)
		}

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
	p.EnterRule(localctx, 26, QuarkParserRULE_concat)
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
		p.SetState(362)
		p.Match(QuarkParserLCURLY)
	}
	{
		p.SetState(363)
		p.Innerconcat()
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == QuarkParserCOMMA {
		{
			p.SetState(364)
			p.Match(QuarkParserCOMMA)
		}
		{
			p.SetState(365)
			p.Innerconcat()
		}

		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(370)
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
	p.EnterRule(localctx, 28, QuarkParserRULE_innerconcat)

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

	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(372)
			p.expr(0)
		}
		{
			p.SetState(373)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(374)
			p.expr(0)
		}
		{
			p.SetState(375)
			p.Match(QuarkParserRCURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(377)
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

func (s *ParameterizedTypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLBRACE, 0)
}

func (s *ParameterizedTypeContext) AllTypeparam() []ITypeparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeparamContext)(nil)).Elem())
	var tst = make([]ITypeparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeparamContext)
		}
	}

	return tst
}

func (s *ParameterizedTypeContext) Typeparam(i int) ITypeparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeparamContext)
}

func (s *ParameterizedTypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRBRACE, 0)
}

func (s *ParameterizedTypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserCOMMA)
}

func (s *ParameterizedTypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserCOMMA, i)
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
	_startState := 30
	p.EnterRecursionRule(localctx, 30, QuarkParserRULE_typeexpr, _p)
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
		p.SetState(381)
		p.Name()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParameterizedTypeContext(p, NewTypeexprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_typeexpr)
			p.SetState(383)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(384)
				p.Match(QuarkParserLBRACE)
			}
			{
				p.SetState(385)
				p.Typeparam()
			}
			p.SetState(390)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == QuarkParserCOMMA {
				{
					p.SetState(386)
					p.Match(QuarkParserCOMMA)
				}
				{
					p.SetState(387)
					p.Typeparam()
				}

				p.SetState(392)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(393)
				p.Match(QuarkParserRBRACE)
			}

		}
		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeparamContext is an interface to support dynamic dispatch.
type ITypeparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeparamContext differentiates from other interfaces.
	IsTypeparamContext()
}

type TypeparamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeparamContext() *TypeparamContext {
	var p = new(TypeparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_typeparam
	return p
}

func (*TypeparamContext) IsTypeparamContext() {}

func NewTypeparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeparamContext {
	var p = new(TypeparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_typeparam

	return p
}

func (s *TypeparamContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeparamContext) KW_TYPE() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_TYPE, 0)
}

func (s *TypeparamContext) Typeexpr() ITypeexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeexprContext)
}

func (s *TypeparamContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TypeparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterTypeparam(s)
	}
}

func (s *TypeparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitTypeparam(s)
	}
}

func (s *TypeparamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitTypeparam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Typeparam() (localctx ITypeparamContext) {
	localctx = NewTypeparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, QuarkParserRULE_typeparam)

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

	p.SetState(403)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QuarkParserKW_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(400)
			p.Match(QuarkParserKW_TYPE)
		}
		{
			p.SetState(401)
			p.typeexpr(0)
		}

	case QuarkParserLPAREN, QuarkParserLBRACE, QuarkParserLCURLY, QuarkParserOP_COMPLIMENT, QuarkParserOP_LNOT, QuarkParserKW_IF, QuarkParserKW_MATCH, QuarkParserKW_LAMBDA, QuarkParserKW_NEW, QuarkParserKW_OPEN, QuarkParserKW_CLOSE, QuarkParserKW_SIGNAL, QuarkParserINTEGRAL, QuarkParserREAL_NAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(402)
			p.expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 34, QuarkParserRULE_branch)
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

	p.SetState(443)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QuarkParserKW_IF:
		localctx = NewIfBranchContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.Match(QuarkParserKW_IF)
		}
		{
			p.SetState(406)
			p.expr(0)
		}
		{
			p.SetState(407)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(408)
			p.Block()
		}
		{
			p.SetState(409)
			p.Match(QuarkParserRCURLY)
		}
		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(410)
					p.Match(QuarkParserKW_ELIF)
				}
				{
					p.SetState(411)
					p.expr(0)
				}
				{
					p.SetState(412)
					p.Match(QuarkParserLCURLY)
				}
				{
					p.SetState(413)
					p.Block()
				}
				{
					p.SetState(414)
					p.Match(QuarkParserRCURLY)
				}

			}
			p.SetState(420)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
		}
		p.SetState(426)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(421)
				p.Match(QuarkParserKW_ELSE)
			}
			{
				p.SetState(422)
				p.Match(QuarkParserLCURLY)
			}
			{
				p.SetState(423)
				p.Block()
			}
			{
				p.SetState(424)
				p.Match(QuarkParserRCURLY)
			}

		}

	case QuarkParserKW_MATCH:
		localctx = NewMatchBranchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(428)
			p.Match(QuarkParserKW_MATCH)
		}
		{
			p.SetState(429)
			p.expr(0)
		}
		{
			p.SetState(430)
			p.Match(QuarkParserLCURLY)
		}
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QuarkParserKW_CASE {
			{
				p.SetState(431)
				p.Match(QuarkParserKW_CASE)
			}
			{
				p.SetState(432)
				p.Pattern()
			}
			{
				p.SetState(433)
				p.Match(QuarkParserLCURLY)
			}
			{
				p.SetState(434)
				p.Block()
			}
			{
				p.SetState(435)
				p.Match(QuarkParserRCURLY)
			}

			p.SetState(439)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(441)
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
	p.EnterRule(localctx, 36, QuarkParserRULE_pattern)
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

	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAtomicPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(445)
			p.Realname()
		}

	case 2:
		localctx = NewParamerterizedTypePatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(446)
			p.Realname()
		}
		{
			p.SetState(447)
			p.Match(QuarkParserLBRACE)
		}
		{
			p.SetState(448)
			p.Pattern()
		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(449)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(450)
				p.Pattern()
			}

			p.SetState(455)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(456)
			p.Match(QuarkParserRBRACE)
		}

	case 3:
		localctx = NewArrayPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(458)
			p.Match(QuarkParserLBRACE)
		}
		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserLBRACE || _la == QuarkParserLCURLY || _la == QuarkParserINTEGRAL || _la == QuarkParserREAL_NAME {
			{
				p.SetState(459)
				p.Pattern()
			}
			p.SetState(464)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == QuarkParserCOMMA {
				{
					p.SetState(460)
					p.Match(QuarkParserCOMMA)
				}
				{
					p.SetState(461)
					p.Pattern()
				}

				p.SetState(466)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(469)
			p.Match(QuarkParserRBRACE)
		}

	case 4:
		localctx = NewLiteralPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(470)
			p.Literal()
		}

	case 5:
		localctx = NewStructPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(471)
			p.Match(QuarkParserLCURLY)
		}

		p.SetState(473)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(472)
				p.typeexpr(0)
			}

		}
		{
			p.SetState(475)
			p.Pattern()
		}

		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(477)
				p.Match(QuarkParserCOMMA)
			}
			p.SetState(479)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(478)
					p.typeexpr(0)
				}

			}
			{
				p.SetState(481)
				p.Pattern()
			}

			p.SetState(486)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(487)
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

func (s *ParameterlistContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLBRACE, 0)
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

func (s *ParameterlistContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRBRACE, 0)
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
	p.EnterRule(localctx, 38, QuarkParserRULE_parameterlist)
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
		p.SetState(491)
		p.Match(QuarkParserLBRACE)
	}
	{
		p.SetState(492)
		p.Parameterdef()
	}
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserCOMMA {
		{
			p.SetState(493)
			p.Match(QuarkParserCOMMA)
		}
		{
			p.SetState(494)
			p.Parameterdef()
		}

		p.SetState(499)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(500)
		p.Match(QuarkParserRBRACE)
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
	p.EnterRule(localctx, 40, QuarkParserRULE_parameterdef)

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

	switch p.GetTokenStream().LA(1) {
	case QuarkParserKW_TYPE:
		localctx = NewTypeParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(502)
			p.Match(QuarkParserKW_TYPE)
		}
		{
			p.SetState(503)
			p.typeexpr(0)
		}

	case QuarkParserREAL_NAME:
		localctx = NewValueParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(504)
			p.typeexpr(0)
		}
		{
			p.SetState(505)
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
	p.EnterRule(localctx, 42, QuarkParserRULE_returnlist)
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

	p.SetState(524)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QuarkParserREAL_NAME:
		localctx = NewSingleReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(509)
			p.typeexpr(0)
		}

	case QuarkParserLPAREN:
		localctx = NewNamedReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(510)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(511)
			p.typeexpr(0)
		}
		{
			p.SetState(512)
			p.Realname()
		}
		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(513)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(514)
				p.typeexpr(0)
			}
			{
				p.SetState(515)
				p.Realname()
			}

			p.SetState(521)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(522)
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
	p.EnterRule(localctx, 44, QuarkParserRULE_argumentdef)

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
		p.SetState(526)
		p.typeexpr(0)
	}
	{
		p.SetState(527)
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
	p.EnterRule(localctx, 46, QuarkParserRULE_argumentlist)
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
		p.SetState(529)
		p.Match(QuarkParserLPAREN)
	}
	p.SetState(538)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserREAL_NAME {
		{
			p.SetState(530)
			p.Argumentdef()
		}
		p.SetState(535)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(531)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(532)
				p.Argumentdef()
			}

			p.SetState(537)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(540)
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
	p.EnterRule(localctx, 48, QuarkParserRULE_structdecl)
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
	p.SetState(545)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(542)
			p.Annotation()
		}

		p.SetState(547)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(548)
		p.Match(QuarkParserKW_STRUCT)
	}
	{
		p.SetState(549)
		p.Realname()
	}
	p.SetState(551)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserLBRACE {
		{
			p.SetState(550)
			p.Parameterlist()
		}

	}
	p.SetState(562)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserKW_HAS {
		{
			p.SetState(553)
			p.Match(QuarkParserKW_HAS)
		}
		{
			p.SetState(554)
			p.Name()
		}
		p.SetState(559)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(555)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(556)
				p.Name()
			}

			p.SetState(561)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(564)
		p.Match(QuarkParserLCURLY)
	}
	p.SetState(568)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(QuarkParserKW_FUTURE-72))|(1<<(QuarkParserREAL_NAME-72))|(1<<(QuarkParserANNOTATION_START-72)))) != 0 {
		{
			p.SetState(565)
			p.Fielddecl()
		}

		p.SetState(570)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(571)
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
	p.EnterRule(localctx, 50, QuarkParserRULE_fielddecl)
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
	p.SetState(576)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(573)
			p.Annotation()
		}

		p.SetState(578)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(580)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserKW_FUTURE {
		{
			p.SetState(579)
			p.Match(QuarkParserKW_FUTURE)
		}

	}
	{
		p.SetState(582)
		p.typeexpr(0)
	}
	{
		p.SetState(583)
		p.Realname()
	}
	{
		p.SetState(584)
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
	p.EnterRule(localctx, 52, QuarkParserRULE_funcdecl)
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
	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(586)
			p.Annotation()
		}

		p.SetState(591)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(592)
		p.Match(QuarkParserKW_DEF)
	}
	{
		p.SetState(593)
		p.Realname()
	}
	p.SetState(595)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserLBRACE {
		{
			p.SetState(594)
			p.Parameterlist()
		}

	}
	p.SetState(598)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserLPAREN {
		{
			p.SetState(597)
			p.Argumentlist()
		}

	}
	p.SetState(602)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserCOLON {
		{
			p.SetState(600)
			p.Match(QuarkParserCOLON)
		}
		{
			p.SetState(601)
			p.Returnlist()
		}

	}
	{
		p.SetState(604)
		p.Match(QuarkParserLCURLY)
	}
	{
		p.SetState(605)
		p.Block()
	}
	{
		p.SetState(606)
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
	p.EnterRule(localctx, 54, QuarkParserRULE_moduledecl)
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
	p.SetState(611)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(608)
			p.Annotation()
		}

		p.SetState(613)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(614)
		p.Match(QuarkParserKW_MODULE)
	}
	{
		p.SetState(615)
		p.Realname()
	}
	p.SetState(617)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserLBRACE {
		{
			p.SetState(616)
			p.Parameterlist()
		}

	}
	p.SetState(620)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserLPAREN {
		{
			p.SetState(619)
			p.Argumentlist()
		}

	}
	p.SetState(624)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserCOLON {
		{
			p.SetState(622)
			p.Match(QuarkParserCOLON)
		}
		{
			p.SetState(623)
			p.Returnlist()
		}

	}
	{
		p.SetState(626)
		p.Match(QuarkParserLCURLY)
	}
	{
		p.SetState(627)
		p.Block()
	}
	{
		p.SetState(628)
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
	p.EnterRule(localctx, 56, QuarkParserRULE_annotation)

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
		p.SetState(630)
		p.Match(QuarkParserANNOTATION_START)
	}
	{
		p.SetState(631)
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
	p.EnterRule(localctx, 58, QuarkParserRULE_literal)

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
		p.SetState(633)
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

	case 15:
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
		return p.Precpred(p.GetParserRuleContext(), 23)

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
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
