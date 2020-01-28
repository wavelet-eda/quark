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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 98, 692,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 3, 2, 7, 2, 70, 10, 2, 12, 2, 14, 2, 73, 11, 2, 3, 2, 7, 2, 76,
	10, 2, 12, 2, 14, 2, 79, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 87, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 6, 4, 106, 10, 4, 13, 4, 14,
	4, 107, 3, 4, 3, 4, 3, 4, 5, 4, 113, 10, 4, 3, 5, 7, 5, 116, 10, 5, 12,
	5, 14, 5, 119, 11, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 5, 7, 130, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 136, 10, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 157, 10, 7, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 7, 8, 165, 10, 8, 12, 8, 14, 8, 168, 11, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 5, 8, 175, 10, 8, 3, 8, 3, 8, 5, 8, 179, 10, 8, 3, 8,
	3, 8, 5, 8, 183, 10, 8, 5, 8, 185, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8,
	191, 10, 8, 3, 8, 3, 8, 3, 8, 5, 8, 196, 10, 8, 3, 8, 3, 8, 3, 8, 6, 8,
	201, 10, 8, 13, 8, 14, 8, 202, 7, 8, 205, 10, 8, 12, 8, 14, 8, 208, 11,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 216, 10, 10, 13, 10,
	14, 10, 217, 5, 10, 220, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 235, 10, 12,
	13, 12, 14, 12, 236, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 254, 10, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 261, 10, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12,
	275, 10, 12, 12, 12, 14, 12, 278, 11, 12, 5, 12, 280, 10, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 288, 10, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 319, 10, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 328, 10, 12, 3, 12, 3, 12, 5, 12,
	332, 10, 12, 3, 12, 5, 12, 335, 10, 12, 3, 12, 5, 12, 338, 10, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 346, 10, 12, 12, 12, 14, 12,
	349, 11, 12, 3, 12, 3, 12, 7, 12, 353, 10, 12, 12, 12, 14, 12, 356, 11,
	12, 3, 13, 3, 13, 3, 13, 7, 13, 361, 10, 13, 12, 13, 14, 13, 364, 11, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 371, 10, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 7, 15, 377, 10, 15, 12, 15, 14, 15, 380, 11, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 6, 16, 388, 10, 16, 13, 16, 14, 16, 389, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 400, 10, 17,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7,
	18, 412, 10, 18, 12, 18, 14, 18, 415, 11, 18, 5, 18, 417, 10, 18, 3, 18,
	7, 18, 420, 10, 18, 12, 18, 14, 18, 423, 11, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 433, 10, 19, 5, 19, 435, 10, 19,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 7, 20, 448, 10, 20, 12, 20, 14, 20, 451, 11, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 5, 20, 458, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 6, 20, 469, 10, 20, 13, 20, 14, 20, 470, 3, 20,
	3, 20, 5, 20, 475, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7,
	21, 483, 10, 21, 12, 21, 14, 21, 486, 11, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 7, 21, 494, 10, 21, 12, 21, 14, 21, 497, 11, 21, 5, 21, 499,
	10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 505, 10, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 5, 21, 511, 10, 21, 3, 21, 7, 21, 514, 10, 21, 12, 21, 14, 21,
	517, 11, 21, 3, 21, 3, 21, 5, 21, 521, 10, 21, 3, 22, 3, 22, 3, 22, 3,
	22, 7, 22, 527, 10, 22, 12, 22, 14, 22, 530, 11, 22, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 539, 10, 23, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 549, 10, 24, 12, 24, 14, 24, 552,
	11, 24, 3, 24, 3, 24, 5, 24, 556, 10, 24, 3, 25, 5, 25, 559, 10, 25, 3,
	25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 7, 26, 568, 10, 26, 12, 26,
	14, 26, 571, 11, 26, 5, 26, 573, 10, 26, 3, 26, 3, 26, 3, 27, 7, 27, 578,
	10, 27, 12, 27, 14, 27, 581, 11, 27, 3, 27, 3, 27, 3, 27, 5, 27, 586, 10,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 7, 27, 592, 10, 27, 12, 27, 14, 27, 595,
	11, 27, 5, 27, 597, 10, 27, 3, 27, 3, 27, 3, 27, 7, 27, 602, 10, 27, 12,
	27, 14, 27, 605, 11, 27, 3, 27, 3, 27, 3, 28, 7, 28, 610, 10, 28, 12, 28,
	14, 28, 613, 11, 28, 3, 28, 5, 28, 616, 10, 28, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 29, 7, 29, 623, 10, 29, 12, 29, 14, 29, 626, 11, 29, 3, 29, 3, 29,
	3, 29, 5, 29, 631, 10, 29, 3, 29, 3, 29, 3, 29, 5, 29, 636, 10, 29, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 7, 31, 644, 10, 31, 12, 31, 14,
	31, 647, 11, 31, 3, 31, 3, 31, 3, 31, 5, 31, 652, 10, 31, 3, 31, 5, 31,
	655, 10, 31, 3, 31, 3, 31, 5, 31, 659, 10, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 32, 7, 32, 666, 10, 32, 12, 32, 14, 32, 669, 11, 32, 3, 32, 3, 32,
	3, 32, 5, 32, 674, 10, 32, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 680, 10,
	32, 12, 32, 14, 32, 683, 11, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3,
	34, 3, 34, 3, 34, 2, 5, 14, 22, 34, 35, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 60, 62, 64, 66, 2, 10, 15, 2, 16, 16, 18, 18, 20, 20, 22, 22, 24,
	24, 26, 26, 28, 28, 31, 31, 34, 34, 37, 37, 39, 39, 41, 41, 43, 43, 5,
	2, 21, 21, 23, 23, 25, 25, 4, 2, 17, 17, 19, 19, 6, 2, 36, 36, 38, 38,
	40, 40, 42, 42, 4, 2, 44, 45, 52, 53, 6, 2, 27, 27, 29, 30, 32, 33, 35,
	35, 4, 2, 47, 48, 50, 51, 3, 2, 90, 91, 2, 772, 2, 71, 3, 2, 2, 2, 4, 86,
	3, 2, 2, 2, 6, 112, 3, 2, 2, 2, 8, 117, 3, 2, 2, 2, 10, 120, 3, 2, 2, 2,
	12, 156, 3, 2, 2, 2, 14, 195, 3, 2, 2, 2, 16, 209, 3, 2, 2, 2, 18, 219,
	3, 2, 2, 2, 20, 221, 3, 2, 2, 2, 22, 287, 3, 2, 2, 2, 24, 357, 3, 2, 2,
	2, 26, 370, 3, 2, 2, 2, 28, 372, 3, 2, 2, 2, 30, 383, 3, 2, 2, 2, 32, 399,
	3, 2, 2, 2, 34, 401, 3, 2, 2, 2, 36, 434, 3, 2, 2, 2, 38, 474, 3, 2, 2,
	2, 40, 520, 3, 2, 2, 2, 42, 522, 3, 2, 2, 2, 44, 538, 3, 2, 2, 2, 46, 555,
	3, 2, 2, 2, 48, 558, 3, 2, 2, 2, 50, 563, 3, 2, 2, 2, 52, 579, 3, 2, 2,
	2, 54, 611, 3, 2, 2, 2, 56, 624, 3, 2, 2, 2, 58, 637, 3, 2, 2, 2, 60, 645,
	3, 2, 2, 2, 62, 667, 3, 2, 2, 2, 64, 686, 3, 2, 2, 2, 66, 689, 3, 2, 2,
	2, 68, 70, 5, 6, 4, 2, 69, 68, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69,
	3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 77, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2,
	74, 76, 5, 4, 3, 2, 75, 74, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3,
	2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 80, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80,
	81, 7, 2, 2, 3, 81, 3, 3, 2, 2, 2, 82, 87, 5, 52, 27, 2, 83, 87, 5, 58,
	30, 2, 84, 87, 5, 60, 31, 2, 85, 87, 5, 62, 32, 2, 86, 82, 3, 2, 2, 2,
	86, 83, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 85, 3, 2, 2, 2, 87, 5, 3, 2,
	2, 2, 88, 89, 7, 70, 2, 2, 89, 90, 5, 18, 10, 2, 90, 91, 7, 4, 2, 2, 91,
	113, 3, 2, 2, 2, 92, 93, 7, 70, 2, 2, 93, 94, 5, 18, 10, 2, 94, 95, 7,
	6, 2, 2, 95, 96, 7, 21, 2, 2, 96, 97, 7, 4, 2, 2, 97, 113, 3, 2, 2, 2,
	98, 99, 7, 70, 2, 2, 99, 100, 5, 18, 10, 2, 100, 101, 7, 6, 2, 2, 101,
	102, 7, 8, 2, 2, 102, 105, 5, 16, 9, 2, 103, 104, 7, 5, 2, 2, 104, 106,
	5, 16, 9, 2, 105, 103, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 105, 3, 2,
	2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 7, 9, 2, 2,
	110, 111, 7, 4, 2, 2, 111, 113, 3, 2, 2, 2, 112, 88, 3, 2, 2, 2, 112, 92,
	3, 2, 2, 2, 112, 98, 3, 2, 2, 2, 113, 7, 3, 2, 2, 2, 114, 116, 5, 12, 7,
	2, 115, 114, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117,
	118, 3, 2, 2, 2, 118, 9, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 121, 9,
	2, 2, 2, 121, 11, 3, 2, 2, 2, 122, 123, 5, 14, 8, 2, 123, 124, 5, 10, 6,
	2, 124, 125, 5, 22, 12, 2, 125, 126, 7, 4, 2, 2, 126, 157, 3, 2, 2, 2,
	127, 129, 7, 80, 2, 2, 128, 130, 5, 28, 15, 2, 129, 128, 3, 2, 2, 2, 129,
	130, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 7, 8, 2, 2, 132, 135,
	5, 20, 11, 2, 133, 134, 7, 5, 2, 2, 134, 136, 5, 20, 11, 2, 135, 133, 3,
	2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 138, 7, 9, 2,
	2, 138, 139, 5, 14, 8, 2, 139, 140, 5, 10, 6, 2, 140, 141, 5, 22, 12, 2,
	141, 142, 7, 4, 2, 2, 142, 157, 3, 2, 2, 2, 143, 144, 7, 75, 2, 2, 144,
	145, 5, 34, 18, 2, 145, 146, 5, 16, 9, 2, 146, 147, 7, 4, 2, 2, 147, 157,
	3, 2, 2, 2, 148, 149, 5, 14, 8, 2, 149, 150, 7, 4, 2, 2, 150, 157, 3, 2,
	2, 2, 151, 157, 5, 38, 20, 2, 152, 153, 7, 69, 2, 2, 153, 154, 5, 22, 12,
	2, 154, 155, 7, 4, 2, 2, 155, 157, 3, 2, 2, 2, 156, 122, 3, 2, 2, 2, 156,
	127, 3, 2, 2, 2, 156, 143, 3, 2, 2, 2, 156, 148, 3, 2, 2, 2, 156, 151,
	3, 2, 2, 2, 156, 152, 3, 2, 2, 2, 157, 13, 3, 2, 2, 2, 158, 159, 8, 8,
	1, 2, 159, 160, 5, 22, 12, 2, 160, 161, 7, 10, 2, 2, 161, 166, 5, 22, 12,
	2, 162, 163, 7, 5, 2, 2, 163, 165, 5, 22, 12, 2, 164, 162, 3, 2, 2, 2,
	165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167,
	169, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 170, 7, 11, 2, 2, 170, 196,
	3, 2, 2, 2, 171, 172, 5, 22, 12, 2, 172, 174, 7, 10, 2, 2, 173, 175, 5,
	22, 12, 2, 174, 173, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 3, 2,
	2, 2, 176, 178, 7, 3, 2, 2, 177, 179, 5, 22, 12, 2, 178, 177, 3, 2, 2,
	2, 178, 179, 3, 2, 2, 2, 179, 184, 3, 2, 2, 2, 180, 182, 7, 3, 2, 2, 181,
	183, 5, 22, 12, 2, 182, 181, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 185,
	3, 2, 2, 2, 184, 180, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 3, 2,
	2, 2, 186, 187, 7, 11, 2, 2, 187, 196, 3, 2, 2, 2, 188, 196, 5, 18, 10,
	2, 189, 191, 7, 85, 2, 2, 190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191,
	192, 3, 2, 2, 2, 192, 193, 5, 34, 18, 2, 193, 194, 5, 16, 9, 2, 194, 196,
	3, 2, 2, 2, 195, 158, 3, 2, 2, 2, 195, 171, 3, 2, 2, 2, 195, 188, 3, 2,
	2, 2, 195, 190, 3, 2, 2, 2, 196, 206, 3, 2, 2, 2, 197, 200, 12, 3, 2, 2,
	198, 199, 7, 5, 2, 2, 199, 201, 5, 14, 8, 2, 200, 198, 3, 2, 2, 2, 201,
	202, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 205,
	3, 2, 2, 2, 204, 197, 3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2,
	2, 2, 206, 207, 3, 2, 2, 2, 207, 15, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2,
	209, 210, 7, 92, 2, 2, 210, 17, 3, 2, 2, 2, 211, 220, 5, 16, 9, 2, 212,
	215, 7, 92, 2, 2, 213, 214, 7, 6, 2, 2, 214, 216, 7, 92, 2, 2, 215, 213,
	3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 217, 218, 3, 2,
	2, 2, 218, 220, 3, 2, 2, 2, 219, 211, 3, 2, 2, 2, 219, 212, 3, 2, 2, 2,
	220, 19, 3, 2, 2, 2, 221, 222, 5, 18, 10, 2, 222, 21, 3, 2, 2, 2, 223,
	224, 8, 12, 1, 2, 224, 288, 5, 66, 34, 2, 225, 288, 5, 16, 9, 2, 226, 227,
	7, 8, 2, 2, 227, 228, 5, 22, 12, 2, 228, 229, 7, 9, 2, 2, 229, 288, 3,
	2, 2, 2, 230, 231, 7, 8, 2, 2, 231, 234, 5, 22, 12, 2, 232, 233, 7, 5,
	2, 2, 233, 235, 5, 22, 12, 2, 234, 232, 3, 2, 2, 2, 235, 236, 3, 2, 2,
	2, 236, 234, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238,
	239, 7, 9, 2, 2, 239, 288, 3, 2, 2, 2, 240, 241, 5, 34, 18, 2, 241, 242,
	7, 12, 2, 2, 242, 243, 5, 24, 13, 2, 243, 244, 7, 13, 2, 2, 244, 288, 3,
	2, 2, 2, 245, 246, 7, 82, 2, 2, 246, 247, 5, 34, 18, 2, 247, 248, 7, 8,
	2, 2, 248, 249, 5, 24, 13, 2, 249, 250, 7, 9, 2, 2, 250, 288, 3, 2, 2,
	2, 251, 253, 7, 77, 2, 2, 252, 254, 5, 42, 22, 2, 253, 252, 3, 2, 2, 2,
	253, 254, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 256, 5, 50, 26, 2, 256,
	257, 7, 15, 2, 2, 257, 258, 7, 12, 2, 2, 258, 260, 5, 8, 5, 2, 259, 261,
	5, 22, 12, 2, 260, 259, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 262, 3,
	2, 2, 2, 262, 263, 7, 13, 2, 2, 263, 288, 3, 2, 2, 2, 264, 265, 7, 46,
	2, 2, 265, 288, 5, 22, 12, 17, 266, 267, 7, 49, 2, 2, 267, 288, 5, 22,
	12, 16, 268, 288, 5, 30, 16, 2, 269, 288, 5, 38, 20, 2, 270, 279, 7, 10,
	2, 2, 271, 276, 5, 22, 12, 2, 272, 273, 7, 5, 2, 2, 273, 275, 5, 22, 12,
	2, 274, 272, 3, 2, 2, 2, 275, 278, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 276,
	277, 3, 2, 2, 2, 277, 280, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 279, 271,
	3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 288, 7, 11,
	2, 2, 282, 283, 7, 86, 2, 2, 283, 284, 7, 8, 2, 2, 284, 285, 5, 20, 11,
	2, 285, 286, 7, 9, 2, 2, 286, 288, 3, 2, 2, 2, 287, 223, 3, 2, 2, 2, 287,
	225, 3, 2, 2, 2, 287, 226, 3, 2, 2, 2, 287, 230, 3, 2, 2, 2, 287, 240,
	3, 2, 2, 2, 287, 245, 3, 2, 2, 2, 287, 251, 3, 2, 2, 2, 287, 264, 3, 2,
	2, 2, 287, 266, 3, 2, 2, 2, 287, 268, 3, 2, 2, 2, 287, 269, 3, 2, 2, 2,
	287, 270, 3, 2, 2, 2, 287, 282, 3, 2, 2, 2, 288, 354, 3, 2, 2, 2, 289,
	290, 12, 14, 2, 2, 290, 291, 9, 3, 2, 2, 291, 353, 5, 22, 12, 15, 292,
	293, 12, 13, 2, 2, 293, 294, 9, 4, 2, 2, 294, 353, 5, 22, 12, 14, 295,
	296, 12, 12, 2, 2, 296, 297, 9, 5, 2, 2, 297, 353, 5, 22, 12, 13, 298,
	299, 12, 11, 2, 2, 299, 300, 9, 6, 2, 2, 300, 353, 5, 22, 12, 12, 301,
	302, 12, 10, 2, 2, 302, 303, 9, 7, 2, 2, 303, 353, 5, 22, 12, 11, 304,
	305, 12, 9, 2, 2, 305, 306, 9, 8, 2, 2, 306, 353, 5, 22, 12, 10, 307, 308,
	12, 8, 2, 2, 308, 309, 7, 58, 2, 2, 309, 310, 5, 22, 12, 2, 310, 311, 7,
	59, 2, 2, 311, 312, 5, 22, 12, 9, 312, 353, 3, 2, 2, 2, 313, 314, 12, 24,
	2, 2, 314, 315, 7, 6, 2, 2, 315, 353, 5, 16, 9, 2, 316, 318, 12, 19, 2,
	2, 317, 319, 5, 28, 15, 2, 318, 317, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2,
	319, 320, 3, 2, 2, 2, 320, 321, 7, 8, 2, 2, 321, 322, 5, 24, 13, 2, 322,
	323, 7, 9, 2, 2, 323, 353, 3, 2, 2, 2, 324, 325, 12, 6, 2, 2, 325, 327,
	7, 10, 2, 2, 326, 328, 5, 22, 12, 2, 327, 326, 3, 2, 2, 2, 327, 328, 3,
	2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 331, 7, 3, 2, 2, 330, 332, 5, 22, 12,
	2, 331, 330, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 334, 3, 2, 2, 2, 333,
	335, 7, 3, 2, 2, 334, 333, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 337,
	3, 2, 2, 2, 336, 338, 5, 22, 12, 2, 337, 336, 3, 2, 2, 2, 337, 338, 3,
	2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 353, 7, 11, 2, 2, 340, 341, 12, 5,
	2, 2, 341, 342, 7, 10, 2, 2, 342, 347, 5, 22, 12, 2, 343, 344, 7, 5, 2,
	2, 344, 346, 5, 22, 12, 2, 345, 343, 3, 2, 2, 2, 346, 349, 3, 2, 2, 2,
	347, 345, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 350, 3, 2, 2, 2, 349,
	347, 3, 2, 2, 2, 350, 351, 7, 11, 2, 2, 351, 353, 3, 2, 2, 2, 352, 289,
	3, 2, 2, 2, 352, 292, 3, 2, 2, 2, 352, 295, 3, 2, 2, 2, 352, 298, 3, 2,
	2, 2, 352, 301, 3, 2, 2, 2, 352, 304, 3, 2, 2, 2, 352, 307, 3, 2, 2, 2,
	352, 313, 3, 2, 2, 2, 352, 316, 3, 2, 2, 2, 352, 324, 3, 2, 2, 2, 352,
	340, 3, 2, 2, 2, 353, 356, 3, 2, 2, 2, 354, 352, 3, 2, 2, 2, 354, 355,
	3, 2, 2, 2, 355, 23, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 357, 362, 5, 26,
	14, 2, 358, 359, 7, 5, 2, 2, 359, 361, 5, 26, 14, 2, 360, 358, 3, 2, 2,
	2, 361, 364, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363,
	25, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 365, 366, 5, 16, 9, 2, 366, 367,
	7, 16, 2, 2, 367, 368, 5, 22, 12, 2, 368, 371, 3, 2, 2, 2, 369, 371, 5,
	22, 12, 2, 370, 365, 3, 2, 2, 2, 370, 369, 3, 2, 2, 2, 371, 27, 3, 2, 2,
	2, 372, 373, 7, 14, 2, 2, 373, 378, 5, 36, 19, 2, 374, 375, 7, 5, 2, 2,
	375, 377, 5, 36, 19, 2, 376, 374, 3, 2, 2, 2, 377, 380, 3, 2, 2, 2, 378,
	376, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 381, 3, 2, 2, 2, 380, 378,
	3, 2, 2, 2, 381, 382, 7, 9, 2, 2, 382, 29, 3, 2, 2, 2, 383, 384, 7, 12,
	2, 2, 384, 387, 5, 32, 17, 2, 385, 386, 7, 5, 2, 2, 386, 388, 5, 32, 17,
	2, 387, 385, 3, 2, 2, 2, 388, 389, 3, 2, 2, 2, 389, 387, 3, 2, 2, 2, 389,
	390, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 392, 7, 13, 2, 2, 392, 31,
	3, 2, 2, 2, 393, 394, 5, 22, 12, 2, 394, 395, 7, 12, 2, 2, 395, 396, 5,
	22, 12, 2, 396, 397, 7, 13, 2, 2, 397, 400, 3, 2, 2, 2, 398, 400, 5, 22,
	12, 2, 399, 393, 3, 2, 2, 2, 399, 398, 3, 2, 2, 2, 400, 33, 3, 2, 2, 2,
	401, 402, 8, 18, 1, 2, 402, 403, 5, 18, 10, 2, 403, 421, 3, 2, 2, 2, 404,
	405, 12, 4, 2, 2, 405, 420, 5, 28, 15, 2, 406, 407, 12, 3, 2, 2, 407, 416,
	7, 10, 2, 2, 408, 413, 5, 22, 12, 2, 409, 410, 7, 5, 2, 2, 410, 412, 5,
	22, 12, 2, 411, 409, 3, 2, 2, 2, 412, 415, 3, 2, 2, 2, 413, 411, 3, 2,
	2, 2, 413, 414, 3, 2, 2, 2, 414, 417, 3, 2, 2, 2, 415, 413, 3, 2, 2, 2,
	416, 408, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418,
	420, 7, 11, 2, 2, 419, 404, 3, 2, 2, 2, 419, 406, 3, 2, 2, 2, 420, 423,
	3, 2, 2, 2, 421, 419, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 35, 3, 2,
	2, 2, 423, 421, 3, 2, 2, 2, 424, 425, 7, 66, 2, 2, 425, 435, 5, 34, 18,
	2, 426, 435, 5, 22, 12, 2, 427, 428, 5, 16, 9, 2, 428, 432, 7, 16, 2, 2,
	429, 433, 5, 22, 12, 2, 430, 431, 7, 66, 2, 2, 431, 433, 5, 34, 18, 2,
	432, 429, 3, 2, 2, 2, 432, 430, 3, 2, 2, 2, 433, 435, 3, 2, 2, 2, 434,
	424, 3, 2, 2, 2, 434, 426, 3, 2, 2, 2, 434, 427, 3, 2, 2, 2, 435, 37, 3,
	2, 2, 2, 436, 437, 7, 58, 2, 2, 437, 438, 5, 22, 12, 2, 438, 439, 7, 12,
	2, 2, 439, 440, 5, 8, 5, 2, 440, 449, 7, 13, 2, 2, 441, 442, 7, 60, 2,
	2, 442, 443, 5, 22, 12, 2, 443, 444, 7, 12, 2, 2, 444, 445, 5, 8, 5, 2,
	445, 446, 7, 13, 2, 2, 446, 448, 3, 2, 2, 2, 447, 441, 3, 2, 2, 2, 448,
	451, 3, 2, 2, 2, 449, 447, 3, 2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 457,
	3, 2, 2, 2, 451, 449, 3, 2, 2, 2, 452, 453, 7, 59, 2, 2, 453, 454, 7, 12,
	2, 2, 454, 455, 5, 8, 5, 2, 455, 456, 7, 13, 2, 2, 456, 458, 3, 2, 2, 2,
	457, 452, 3, 2, 2, 2, 457, 458, 3, 2, 2, 2, 458, 475, 3, 2, 2, 2, 459,
	460, 7, 73, 2, 2, 460, 461, 5, 22, 12, 2, 461, 468, 7, 12, 2, 2, 462, 463,
	7, 74, 2, 2, 463, 464, 5, 40, 21, 2, 464, 465, 7, 12, 2, 2, 465, 466, 5,
	8, 5, 2, 466, 467, 7, 13, 2, 2, 467, 469, 3, 2, 2, 2, 468, 462, 3, 2, 2,
	2, 469, 470, 3, 2, 2, 2, 470, 468, 3, 2, 2, 2, 470, 471, 3, 2, 2, 2, 471,
	472, 3, 2, 2, 2, 472, 473, 7, 13, 2, 2, 473, 475, 3, 2, 2, 2, 474, 436,
	3, 2, 2, 2, 474, 459, 3, 2, 2, 2, 475, 39, 3, 2, 2, 2, 476, 521, 5, 16,
	9, 2, 477, 478, 5, 16, 9, 2, 478, 479, 7, 10, 2, 2, 479, 484, 5, 40, 21,
	2, 480, 481, 7, 5, 2, 2, 481, 483, 5, 40, 21, 2, 482, 480, 3, 2, 2, 2,
	483, 486, 3, 2, 2, 2, 484, 482, 3, 2, 2, 2, 484, 485, 3, 2, 2, 2, 485,
	487, 3, 2, 2, 2, 486, 484, 3, 2, 2, 2, 487, 488, 7, 11, 2, 2, 488, 521,
	3, 2, 2, 2, 489, 498, 7, 10, 2, 2, 490, 495, 5, 40, 21, 2, 491, 492, 7,
	5, 2, 2, 492, 494, 5, 40, 21, 2, 493, 491, 3, 2, 2, 2, 494, 497, 3, 2,
	2, 2, 495, 493, 3, 2, 2, 2, 495, 496, 3, 2, 2, 2, 496, 499, 3, 2, 2, 2,
	497, 495, 3, 2, 2, 2, 498, 490, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2, 499,
	500, 3, 2, 2, 2, 500, 521, 7, 11, 2, 2, 501, 521, 5, 66, 34, 2, 502, 504,
	7, 12, 2, 2, 503, 505, 5, 34, 18, 2, 504, 503, 3, 2, 2, 2, 504, 505, 3,
	2, 2, 2, 505, 506, 3, 2, 2, 2, 506, 507, 5, 40, 21, 2, 507, 515, 3, 2,
	2, 2, 508, 510, 7, 5, 2, 2, 509, 511, 5, 34, 18, 2, 510, 509, 3, 2, 2,
	2, 510, 511, 3, 2, 2, 2, 511, 512, 3, 2, 2, 2, 512, 514, 5, 40, 21, 2,
	513, 508, 3, 2, 2, 2, 514, 517, 3, 2, 2, 2, 515, 513, 3, 2, 2, 2, 515,
	516, 3, 2, 2, 2, 516, 518, 3, 2, 2, 2, 517, 515, 3, 2, 2, 2, 518, 519,
	7, 13, 2, 2, 519, 521, 3, 2, 2, 2, 520, 476, 3, 2, 2, 2, 520, 477, 3, 2,
	2, 2, 520, 489, 3, 2, 2, 2, 520, 501, 3, 2, 2, 2, 520, 502, 3, 2, 2, 2,
	521, 41, 3, 2, 2, 2, 522, 523, 7, 14, 2, 2, 523, 528, 5, 44, 23, 2, 524,
	525, 7, 5, 2, 2, 525, 527, 5, 44, 23, 2, 526, 524, 3, 2, 2, 2, 527, 530,
	3, 2, 2, 2, 528, 526, 3, 2, 2, 2, 528, 529, 3, 2, 2, 2, 529, 531, 3, 2,
	2, 2, 530, 528, 3, 2, 2, 2, 531, 532, 7, 9, 2, 2, 532, 43, 3, 2, 2, 2,
	533, 534, 7, 66, 2, 2, 534, 539, 5, 34, 18, 2, 535, 536, 5, 34, 18, 2,
	536, 537, 5, 16, 9, 2, 537, 539, 3, 2, 2, 2, 538, 533, 3, 2, 2, 2, 538,
	535, 3, 2, 2, 2, 539, 45, 3, 2, 2, 2, 540, 556, 5, 34, 18, 2, 541, 542,
	7, 8, 2, 2, 542, 543, 5, 34, 18, 2, 543, 550, 5, 16, 9, 2, 544, 545, 7,
	5, 2, 2, 545, 546, 5, 34, 18, 2, 546, 547, 5, 16, 9, 2, 547, 549, 3, 2,
	2, 2, 548, 544, 3, 2, 2, 2, 549, 552, 3, 2, 2, 2, 550, 548, 3, 2, 2, 2,
	550, 551, 3, 2, 2, 2, 551, 553, 3, 2, 2, 2, 552, 550, 3, 2, 2, 2, 553,
	554, 7, 9, 2, 2, 554, 556, 3, 2, 2, 2, 555, 540, 3, 2, 2, 2, 555, 541,
	3, 2, 2, 2, 556, 47, 3, 2, 2, 2, 557, 559, 7, 75, 2, 2, 558, 557, 3, 2,
	2, 2, 558, 559, 3, 2, 2, 2, 559, 560, 3, 2, 2, 2, 560, 561, 5, 34, 18,
	2, 561, 562, 5, 16, 9, 2, 562, 49, 3, 2, 2, 2, 563, 572, 7, 8, 2, 2, 564,
	569, 5, 48, 25, 2, 565, 566, 7, 5, 2, 2, 566, 568, 5, 48, 25, 2, 567, 565,
	3, 2, 2, 2, 568, 571, 3, 2, 2, 2, 569, 567, 3, 2, 2, 2, 569, 570, 3, 2,
	2, 2, 570, 573, 3, 2, 2, 2, 571, 569, 3, 2, 2, 2, 572, 564, 3, 2, 2, 2,
	572, 573, 3, 2, 2, 2, 573, 574, 3, 2, 2, 2, 574, 575, 7, 9, 2, 2, 575,
	51, 3, 2, 2, 2, 576, 578, 5, 64, 33, 2, 577, 576, 3, 2, 2, 2, 578, 581,
	3, 2, 2, 2, 579, 577, 3, 2, 2, 2, 579, 580, 3, 2, 2, 2, 580, 582, 3, 2,
	2, 2, 581, 579, 3, 2, 2, 2, 582, 583, 7, 64, 2, 2, 583, 585, 5, 16, 9,
	2, 584, 586, 5, 42, 22, 2, 585, 584, 3, 2, 2, 2, 585, 586, 3, 2, 2, 2,
	586, 596, 3, 2, 2, 2, 587, 588, 7, 63, 2, 2, 588, 593, 5, 18, 10, 2, 589,
	590, 7, 5, 2, 2, 590, 592, 5, 18, 10, 2, 591, 589, 3, 2, 2, 2, 592, 595,
	3, 2, 2, 2, 593, 591, 3, 2, 2, 2, 593, 594, 3, 2, 2, 2, 594, 597, 3, 2,
	2, 2, 595, 593, 3, 2, 2, 2, 596, 587, 3, 2, 2, 2, 596, 597, 3, 2, 2, 2,
	597, 598, 3, 2, 2, 2, 598, 603, 7, 12, 2, 2, 599, 602, 5, 54, 28, 2, 600,
	602, 5, 58, 30, 2, 601, 599, 3, 2, 2, 2, 601, 600, 3, 2, 2, 2, 602, 605,
	3, 2, 2, 2, 603, 601, 3, 2, 2, 2, 603, 604, 3, 2, 2, 2, 604, 606, 3, 2,
	2, 2, 605, 603, 3, 2, 2, 2, 606, 607, 7, 13, 2, 2, 607, 53, 3, 2, 2, 2,
	608, 610, 5, 64, 33, 2, 609, 608, 3, 2, 2, 2, 610, 613, 3, 2, 2, 2, 611,
	609, 3, 2, 2, 2, 611, 612, 3, 2, 2, 2, 612, 615, 3, 2, 2, 2, 613, 611,
	3, 2, 2, 2, 614, 616, 7, 75, 2, 2, 615, 614, 3, 2, 2, 2, 615, 616, 3, 2,
	2, 2, 616, 617, 3, 2, 2, 2, 617, 618, 5, 34, 18, 2, 618, 619, 5, 16, 9,
	2, 619, 620, 7, 4, 2, 2, 620, 55, 3, 2, 2, 2, 621, 623, 5, 64, 33, 2, 622,
	621, 3, 2, 2, 2, 623, 626, 3, 2, 2, 2, 624, 622, 3, 2, 2, 2, 624, 625,
	3, 2, 2, 2, 625, 627, 3, 2, 2, 2, 626, 624, 3, 2, 2, 2, 627, 628, 7, 57,
	2, 2, 628, 630, 5, 16, 9, 2, 629, 631, 5, 42, 22, 2, 630, 629, 3, 2, 2,
	2, 630, 631, 3, 2, 2, 2, 631, 632, 3, 2, 2, 2, 632, 635, 5, 50, 26, 2,
	633, 634, 7, 3, 2, 2, 634, 636, 5, 46, 24, 2, 635, 633, 3, 2, 2, 2, 635,
	636, 3, 2, 2, 2, 636, 57, 3, 2, 2, 2, 637, 638, 5, 56, 29, 2, 638, 639,
	7, 12, 2, 2, 639, 640, 5, 8, 5, 2, 640, 641, 7, 13, 2, 2, 641, 59, 3, 2,
	2, 2, 642, 644, 5, 64, 33, 2, 643, 642, 3, 2, 2, 2, 644, 647, 3, 2, 2,
	2, 645, 643, 3, 2, 2, 2, 645, 646, 3, 2, 2, 2, 646, 648, 3, 2, 2, 2, 647,
	645, 3, 2, 2, 2, 648, 649, 7, 56, 2, 2, 649, 651, 5, 16, 9, 2, 650, 652,
	5, 42, 22, 2, 651, 650, 3, 2, 2, 2, 651, 652, 3, 2, 2, 2, 652, 654, 3,
	2, 2, 2, 653, 655, 5, 50, 26, 2, 654, 653, 3, 2, 2, 2, 654, 655, 3, 2,
	2, 2, 655, 658, 3, 2, 2, 2, 656, 657, 7, 3, 2, 2, 657, 659, 5, 46, 24,
	2, 658, 656, 3, 2, 2, 2, 658, 659, 3, 2, 2, 2, 659, 660, 3, 2, 2, 2, 660,
	661, 7, 12, 2, 2, 661, 662, 5, 8, 5, 2, 662, 663, 7, 13, 2, 2, 663, 61,
	3, 2, 2, 2, 664, 666, 5, 64, 33, 2, 665, 664, 3, 2, 2, 2, 666, 669, 3,
	2, 2, 2, 667, 665, 3, 2, 2, 2, 667, 668, 3, 2, 2, 2, 668, 670, 3, 2, 2,
	2, 669, 667, 3, 2, 2, 2, 670, 671, 7, 87, 2, 2, 671, 673, 5, 16, 9, 2,
	672, 674, 5, 42, 22, 2, 673, 672, 3, 2, 2, 2, 673, 674, 3, 2, 2, 2, 674,
	675, 3, 2, 2, 2, 675, 681, 7, 12, 2, 2, 676, 677, 5, 56, 29, 2, 677, 678,
	7, 4, 2, 2, 678, 680, 3, 2, 2, 2, 679, 676, 3, 2, 2, 2, 680, 683, 3, 2,
	2, 2, 681, 679, 3, 2, 2, 2, 681, 682, 3, 2, 2, 2, 682, 684, 3, 2, 2, 2,
	683, 681, 3, 2, 2, 2, 684, 685, 7, 13, 2, 2, 685, 63, 3, 2, 2, 2, 686,
	687, 7, 93, 2, 2, 687, 688, 5, 16, 9, 2, 688, 65, 3, 2, 2, 2, 689, 690,
	9, 9, 2, 2, 690, 67, 3, 2, 2, 2, 83, 71, 77, 86, 107, 112, 117, 129, 135,
	156, 166, 174, 178, 182, 184, 190, 195, 202, 206, 217, 219, 236, 253, 260,
	276, 279, 287, 318, 327, 331, 334, 337, 347, 352, 354, 362, 370, 378, 389,
	399, 413, 416, 419, 421, 432, 434, 449, 457, 470, 474, 484, 495, 498, 504,
	510, 515, 520, 528, 538, 550, 555, 558, 569, 572, 579, 585, 593, 596, 601,
	603, 611, 615, 624, 630, 635, 645, 651, 654, 658, 667, 673, 681,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "';'", "','", "'.'", "'_'", "'('", "')'", "'['", "']'", "'{'",
	"'}'", "'!('", "'=>'", "'='", "'+'", "'+='", "'-'", "'-='", "'*'", "'*='",
	"'/'", "'/='", "'%'", "'%='", "'&'", "'&='", "'~&'", "'|'", "'|='", "'~|'",
	"'^'", "'^='", "'~^'", "'<<'", "'<<='", "'>>'", "'>>='", "'<<<'", "'<<<='",
	"'>>>'", "'>>>='", "'<'", "'>'", "'~'", "'and'", "'or'", "'not'", "'implies'",
	"'equates'", "'<='", "'>='", "'=='", "'!='", "'module'", "'def'", "'if'",
	"'else'", "'elif'", "'for'", "'in'", "'has'", "'struct'", "'abstract'",
	"'type'", "'enum'", "'value'", "'return'", "'import'", "'public'", "'private'",
	"'match'", "'case'", "'future'", "'logic'", "'lambda'", "'clock'", "'reset'",
	"'reg'", "'var'", "'new'", "'open'", "'close'", "'mut'", "'signal'", "'trait'",
	"'//'", "'/*'", "", "", "", "'@'", "", "", "", "'*/'",
}
var symbolicNames = []string{
	"", "COLON", "SEMI", "COMMA", "DOT", "UNDERSCORE", "LPAREN", "RPAREN",
	"LBRACE", "RBRACE", "LCURLY", "RCURLY", "PARAM_OPEN", "OP_ARROW", "OP_ASSIGN",
	"OP_ADD", "OP_ADD_ASSIGN", "OP_SUB", "OP_SUB_ASSIGN", "OP_MUL", "OP_MUL_ASSIGN",
	"OP_DIV", "OP_DIV_ASSIGN", "OP_MOD", "OP_MOD_ASSIGN", "OP_BAND", "OP_BAND_ASSIGN",
	"OP_BNAND", "OP_BOR", "OP_BOR_ASSIGN", "OP_BNOR", "OP_XOR", "OP_XOR_ASSIGN",
	"OP_XNOR", "OP_LEFT_SHIFT", "OP_LEFT_SHIFT_ASSIGN", "OP_RIGHT_SHIFT", "OP_RIGHT_SHIFT_ASSIGN",
	"OP_ARITH_LEFT_SHIFT", "OP_ARITH_LEFT_SHIFT_ASSIGN", "OP_ARITH_RIGHT_SHFIT",
	"OP_ARITH_RIGHT_SHIFT_ASSIGN", "LANGLE", "RANGLE", "OP_COMPLIMENT", "OP_LAND",
	"OP_LOR", "OP_LNOT", "OP_IMPLICATION", "OP_EQUIVALENCE", "OP_LTE", "OP_GTE",
	"OP_EQ", "OP_NEQ", "KW_MODULE", "KW_DEF", "KW_IF", "KW_ELSE", "KW_ELIF",
	"KW_FOR", "KW_IN", "KW_HAS", "KW_STRUCT", "KW_ABSTRACT", "KW_TYPE", "KW_ENUM",
	"KW_VALUE", "KW_RETURN", "KW_IMPORT", "KW_PUBLIC", "KW_PRIVATE", "KW_MATCH",
	"KW_CASE", "KW_FUTURE", "KW_LOGIC", "KW_LAMBDA", "KW_CLOCK", "KW_RESET",
	"KW_REG", "KW_VAR", "KW_NEW", "KW_OPEN", "KW_CLOSE", "KW_MUT", "KW_SIGNAL",
	"KW_TRAIT", "COMMENT_START", "BLOCK_COMMENT_START", "DECIMAL", "INTEGRAL",
	"REAL_NAME", "ANNOTATION_START", "WS", "NEW_LINE", "ANYCHAR", "BLOCK_COMMENT_END",
	"BLOCK_COMMENT_CHAR",
}

var ruleNames = []string{
	"quarkpackage", "decl", "importdecl", "block", "assignment", "stmt", "assignable",
	"realname", "name", "clockexpr", "expr", "callarglist", "callarg", "paramarglist",
	"concat", "innerconcat", "typeexpr", "paramarg", "branch", "pattern", "parameterlist",
	"parameterdef", "returnlist", "argumentdef", "argumentlist", "structdecl",
	"fielddecl", "funcsig", "funcdecl", "moduledecl", "traitdecl", "annotation",
	"literal",
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
	QuarkParserLCURLY                      = 10
	QuarkParserRCURLY                      = 11
	QuarkParserPARAM_OPEN                  = 12
	QuarkParserOP_ARROW                    = 13
	QuarkParserOP_ASSIGN                   = 14
	QuarkParserOP_ADD                      = 15
	QuarkParserOP_ADD_ASSIGN               = 16
	QuarkParserOP_SUB                      = 17
	QuarkParserOP_SUB_ASSIGN               = 18
	QuarkParserOP_MUL                      = 19
	QuarkParserOP_MUL_ASSIGN               = 20
	QuarkParserOP_DIV                      = 21
	QuarkParserOP_DIV_ASSIGN               = 22
	QuarkParserOP_MOD                      = 23
	QuarkParserOP_MOD_ASSIGN               = 24
	QuarkParserOP_BAND                     = 25
	QuarkParserOP_BAND_ASSIGN              = 26
	QuarkParserOP_BNAND                    = 27
	QuarkParserOP_BOR                      = 28
	QuarkParserOP_BOR_ASSIGN               = 29
	QuarkParserOP_BNOR                     = 30
	QuarkParserOP_XOR                      = 31
	QuarkParserOP_XOR_ASSIGN               = 32
	QuarkParserOP_XNOR                     = 33
	QuarkParserOP_LEFT_SHIFT               = 34
	QuarkParserOP_LEFT_SHIFT_ASSIGN        = 35
	QuarkParserOP_RIGHT_SHIFT              = 36
	QuarkParserOP_RIGHT_SHIFT_ASSIGN       = 37
	QuarkParserOP_ARITH_LEFT_SHIFT         = 38
	QuarkParserOP_ARITH_LEFT_SHIFT_ASSIGN  = 39
	QuarkParserOP_ARITH_RIGHT_SHFIT        = 40
	QuarkParserOP_ARITH_RIGHT_SHIFT_ASSIGN = 41
	QuarkParserLANGLE                      = 42
	QuarkParserRANGLE                      = 43
	QuarkParserOP_COMPLIMENT               = 44
	QuarkParserOP_LAND                     = 45
	QuarkParserOP_LOR                      = 46
	QuarkParserOP_LNOT                     = 47
	QuarkParserOP_IMPLICATION              = 48
	QuarkParserOP_EQUIVALENCE              = 49
	QuarkParserOP_LTE                      = 50
	QuarkParserOP_GTE                      = 51
	QuarkParserOP_EQ                       = 52
	QuarkParserOP_NEQ                      = 53
	QuarkParserKW_MODULE                   = 54
	QuarkParserKW_DEF                      = 55
	QuarkParserKW_IF                       = 56
	QuarkParserKW_ELSE                     = 57
	QuarkParserKW_ELIF                     = 58
	QuarkParserKW_FOR                      = 59
	QuarkParserKW_IN                       = 60
	QuarkParserKW_HAS                      = 61
	QuarkParserKW_STRUCT                   = 62
	QuarkParserKW_ABSTRACT                 = 63
	QuarkParserKW_TYPE                     = 64
	QuarkParserKW_ENUM                     = 65
	QuarkParserKW_VALUE                    = 66
	QuarkParserKW_RETURN                   = 67
	QuarkParserKW_IMPORT                   = 68
	QuarkParserKW_PUBLIC                   = 69
	QuarkParserKW_PRIVATE                  = 70
	QuarkParserKW_MATCH                    = 71
	QuarkParserKW_CASE                     = 72
	QuarkParserKW_FUTURE                   = 73
	QuarkParserKW_LOGIC                    = 74
	QuarkParserKW_LAMBDA                   = 75
	QuarkParserKW_CLOCK                    = 76
	QuarkParserKW_RESET                    = 77
	QuarkParserKW_REG                      = 78
	QuarkParserKW_VAR                      = 79
	QuarkParserKW_NEW                      = 80
	QuarkParserKW_OPEN                     = 81
	QuarkParserKW_CLOSE                    = 82
	QuarkParserKW_MUT                      = 83
	QuarkParserKW_SIGNAL                   = 84
	QuarkParserKW_TRAIT                    = 85
	QuarkParserCOMMENT_START               = 86
	QuarkParserBLOCK_COMMENT_START         = 87
	QuarkParserDECIMAL                     = 88
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
	QuarkParserRULE_funcsig       = 27
	QuarkParserRULE_funcdecl      = 28
	QuarkParserRULE_moduledecl    = 29
	QuarkParserRULE_traitdecl     = 30
	QuarkParserRULE_annotation    = 31
	QuarkParserRULE_literal       = 32
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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserKW_IMPORT {
		{
			p.SetState(66)
			p.Importdecl()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(QuarkParserKW_MODULE-54))|(1<<(QuarkParserKW_DEF-54))|(1<<(QuarkParserKW_STRUCT-54))|(1<<(QuarkParserKW_TRAIT-54)))) != 0) || _la == QuarkParserANNOTATION_START {
		{
			p.SetState(72)
			p.Decl()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
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

func (s *DeclContext) Traitdecl() ITraitdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITraitdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITraitdeclContext)
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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Structdecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Funcdecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.Moduledecl()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.Traitdecl()
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

	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(QuarkParserKW_IMPORT)
		}
		{
			p.SetState(87)
			p.Name()
		}
		{
			p.SetState(88)
			p.Match(QuarkParserSEMI)
		}

	case 2:
		localctx = NewWildcardImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(QuarkParserKW_IMPORT)
		}
		{
			p.SetState(91)
			p.Name()
		}
		{
			p.SetState(92)
			p.Match(QuarkParserDOT)
		}
		{
			p.SetState(93)
			p.Match(QuarkParserOP_MUL)
		}
		{
			p.SetState(94)
			p.Match(QuarkParserSEMI)
		}

	case 3:
		localctx = NewMultiImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.Match(QuarkParserKW_IMPORT)
		}
		{
			p.SetState(97)
			p.Name()
		}
		{
			p.SetState(98)
			p.Match(QuarkParserDOT)
		}
		{
			p.SetState(99)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(100)
			p.Realname()
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QuarkParserCOMMA {
			{
				p.SetState(101)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(102)
				p.Realname()
			}

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(107)
			p.Match(QuarkParserRPAREN)
		}
		{
			p.SetState(108)
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
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(112)
				p.Stmt()
			}

		}
		p.SetState(117)
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
		p.SetState(118)
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

func (s *RegAssignStmtContext) Paramarglist() IParamarglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamarglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamarglistContext)
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

	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.assignable(0)
		}
		{
			p.SetState(121)
			p.Assignment()
		}
		{
			p.SetState(122)
			p.expr(0)
		}
		{
			p.SetState(123)
			p.Match(QuarkParserSEMI)
		}

	case 2:
		localctx = NewRegAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Match(QuarkParserKW_REG)
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserPARAM_OPEN {
			{
				p.SetState(126)
				p.Paramarglist()
			}

		}
		{
			p.SetState(129)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(130)

			var _x = p.Clockexpr()

			localctx.(*RegAssignStmtContext).clk = _x
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserCOMMA {
			{
				p.SetState(131)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(132)

				var _x = p.Clockexpr()

				localctx.(*RegAssignStmtContext).rst = _x
			}

		}
		{
			p.SetState(135)
			p.Match(QuarkParserRPAREN)
		}
		{
			p.SetState(136)
			p.assignable(0)
		}
		{
			p.SetState(137)
			p.Assignment()
		}
		{
			p.SetState(138)
			p.expr(0)
		}
		{
			p.SetState(139)
			p.Match(QuarkParserSEMI)
		}

	case 3:
		localctx = NewFutureStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(141)
			p.Match(QuarkParserKW_FUTURE)
		}
		{
			p.SetState(142)
			p.typeexpr(0)
		}
		{
			p.SetState(143)
			p.Realname()
		}
		{
			p.SetState(144)
			p.Match(QuarkParserSEMI)
		}

	case 4:
		localctx = NewDeclarationStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(146)
			p.assignable(0)
		}
		{
			p.SetState(147)
			p.Match(QuarkParserSEMI)
		}

	case 5:
		localctx = NewBranchStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(149)
			p.Branch()
		}

	case 6:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(150)
			p.Match(QuarkParserKW_RETURN)
		}
		{
			p.SetState(151)
			p.expr(0)
		}
		{
			p.SetState(152)
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
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArrayIndexAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(157)
			p.expr(0)
		}
		{
			p.SetState(158)
			p.Match(QuarkParserLBRACE)
		}
		{
			p.SetState(159)
			p.expr(0)
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(160)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(161)
				p.expr(0)
			}

			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(167)
			p.Match(QuarkParserRBRACE)
		}

	case 2:
		localctx = NewArraySliceAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(169)
			p.expr(0)
		}
		{
			p.SetState(170)
			p.Match(QuarkParserLBRACE)
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(QuarkParserOP_COMPLIMENT-44))|(1<<(QuarkParserOP_LNOT-44))|(1<<(QuarkParserKW_IF-44))|(1<<(QuarkParserKW_MATCH-44))|(1<<(QuarkParserKW_LAMBDA-44)))) != 0) || (((_la-80)&-(0x1f+1)) == 0 && ((1<<uint((_la-80)))&((1<<(QuarkParserKW_NEW-80))|(1<<(QuarkParserKW_SIGNAL-80))|(1<<(QuarkParserDECIMAL-80))|(1<<(QuarkParserINTEGRAL-80))|(1<<(QuarkParserREAL_NAME-80)))) != 0) {
			{
				p.SetState(171)

				var _x = p.expr(0)

				localctx.(*ArraySliceAssignmentContext).msb = _x
			}

		}
		{
			p.SetState(174)
			p.Match(QuarkParserCOLON)
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(QuarkParserOP_COMPLIMENT-44))|(1<<(QuarkParserOP_LNOT-44))|(1<<(QuarkParserKW_IF-44))|(1<<(QuarkParserKW_MATCH-44))|(1<<(QuarkParserKW_LAMBDA-44)))) != 0) || (((_la-80)&-(0x1f+1)) == 0 && ((1<<uint((_la-80)))&((1<<(QuarkParserKW_NEW-80))|(1<<(QuarkParserKW_SIGNAL-80))|(1<<(QuarkParserDECIMAL-80))|(1<<(QuarkParserINTEGRAL-80))|(1<<(QuarkParserREAL_NAME-80)))) != 0) {
			{
				p.SetState(175)

				var _x = p.expr(0)

				localctx.(*ArraySliceAssignmentContext).lsb = _x
			}

		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserCOLON {
			{
				p.SetState(178)
				p.Match(QuarkParserCOLON)
			}
			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(QuarkParserOP_COMPLIMENT-44))|(1<<(QuarkParserOP_LNOT-44))|(1<<(QuarkParserKW_IF-44))|(1<<(QuarkParserKW_MATCH-44))|(1<<(QuarkParserKW_LAMBDA-44)))) != 0) || (((_la-80)&-(0x1f+1)) == 0 && ((1<<uint((_la-80)))&((1<<(QuarkParserKW_NEW-80))|(1<<(QuarkParserKW_SIGNAL-80))|(1<<(QuarkParserDECIMAL-80))|(1<<(QuarkParserINTEGRAL-80))|(1<<(QuarkParserREAL_NAME-80)))) != 0) {
				{
					p.SetState(179)

					var _x = p.expr(0)

					localctx.(*ArraySliceAssignmentContext).step = _x
				}

			}

		}
		{
			p.SetState(184)
			p.Match(QuarkParserRBRACE)
		}

	case 3:
		localctx = NewValueAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(186)
			p.Name()
		}

	case 4:
		localctx = NewVariableDefinitionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserKW_MUT {
			{
				p.SetState(187)
				p.Match(QuarkParserKW_MUT)
			}

		}
		{
			p.SetState(190)
			p.typeexpr(0)
		}
		{
			p.SetState(191)
			p.Realname()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTupleDestructerContext(p, NewAssignableContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_assignable)
			p.SetState(195)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(196)
						p.Match(QuarkParserCOMMA)
					}
					{
						p.SetState(197)
						p.assignable(0)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(200)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
			}

		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
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
		p.SetState(207)
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

	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnqualifiedNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.Realname()
		}

	case 2:
		localctx = NewQualifiedNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.Match(QuarkParserREAL_NAME)
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(211)
					p.Match(QuarkParserDOT)
				}
				{
					p.SetState(212)
					p.Match(QuarkParserREAL_NAME)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(215)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
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
		p.SetState(219)
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

type CompareExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewCompareExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExprContext {
	var p = new(CompareExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CompareExprContext) GetOp() antlr.Token { return s.op }

func (s *CompareExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CompareExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompareExprContext) LANGLE() antlr.TerminalNode {
	return s.GetToken(QuarkParserLANGLE, 0)
}

func (s *CompareExprContext) RANGLE() antlr.TerminalNode {
	return s.GetToken(QuarkParserRANGLE, 0)
}

func (s *CompareExprContext) OP_LTE() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_LTE, 0)
}

func (s *CompareExprContext) OP_GTE() antlr.TerminalNode {
	return s.GetToken(QuarkParserOP_GTE, 0)
}

func (s *CompareExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterCompareExpr(s)
	}
}

func (s *CompareExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitCompareExpr(s)
	}
}

func (s *CompareExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitCompareExpr(s)

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
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(222)
			p.Literal()
		}

	case 2:
		localctx = NewVarExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(223)
			p.Realname()
		}

	case 3:
		localctx = NewParensExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(224)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(225)
			p.expr(0)
		}
		{
			p.SetState(226)
			p.Match(QuarkParserRPAREN)
		}

	case 4:
		localctx = NewTupleExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(228)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(229)
			p.expr(0)
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QuarkParserCOMMA {
			{
				p.SetState(230)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(231)
				p.expr(0)
			}

			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(236)
			p.Match(QuarkParserRPAREN)
		}

	case 5:
		localctx = NewConstructorExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.typeexpr(0)
		}
		{
			p.SetState(239)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(240)
			p.Callarglist()
		}
		{
			p.SetState(241)
			p.Match(QuarkParserRCURLY)
		}

	case 6:
		localctx = NewNewModuleExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(243)
			p.Match(QuarkParserKW_NEW)
		}
		{
			p.SetState(244)
			p.typeexpr(0)
		}
		{
			p.SetState(245)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(246)
			p.Callarglist()
		}
		{
			p.SetState(247)
			p.Match(QuarkParserRPAREN)
		}

	case 7:
		localctx = NewLambdaExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(249)
			p.Match(QuarkParserKW_LAMBDA)
		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserPARAM_OPEN {
			{
				p.SetState(250)
				p.Parameterlist()
			}

		}
		{
			p.SetState(253)
			p.Argumentlist()
		}
		{
			p.SetState(254)
			p.Match(QuarkParserOP_ARROW)
		}
		{
			p.SetState(255)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(256)
			p.Block()
		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(QuarkParserOP_COMPLIMENT-44))|(1<<(QuarkParserOP_LNOT-44))|(1<<(QuarkParserKW_IF-44))|(1<<(QuarkParserKW_MATCH-44))|(1<<(QuarkParserKW_LAMBDA-44)))) != 0) || (((_la-80)&-(0x1f+1)) == 0 && ((1<<uint((_la-80)))&((1<<(QuarkParserKW_NEW-80))|(1<<(QuarkParserKW_SIGNAL-80))|(1<<(QuarkParserDECIMAL-80))|(1<<(QuarkParserINTEGRAL-80))|(1<<(QuarkParserREAL_NAME-80)))) != 0) {
			{
				p.SetState(257)
				p.expr(0)
			}

		}
		{
			p.SetState(260)
			p.Match(QuarkParserRCURLY)
		}

	case 8:
		localctx = NewComplimentExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(262)
			p.Match(QuarkParserOP_COMPLIMENT)
		}
		{
			p.SetState(263)
			p.expr(15)
		}

	case 9:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(264)
			p.Match(QuarkParserOP_LNOT)
		}
		{
			p.SetState(265)
			p.expr(14)
		}

	case 10:
		localctx = NewConcatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(266)
			p.Concat()
		}

	case 11:
		localctx = NewBranchExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(267)
			p.Branch()
		}

	case 12:
		localctx = NewArrayLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(268)
			p.Match(QuarkParserLBRACE)
		}
		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(QuarkParserOP_COMPLIMENT-44))|(1<<(QuarkParserOP_LNOT-44))|(1<<(QuarkParserKW_IF-44))|(1<<(QuarkParserKW_MATCH-44))|(1<<(QuarkParserKW_LAMBDA-44)))) != 0) || (((_la-80)&-(0x1f+1)) == 0 && ((1<<uint((_la-80)))&((1<<(QuarkParserKW_NEW-80))|(1<<(QuarkParserKW_SIGNAL-80))|(1<<(QuarkParserDECIMAL-80))|(1<<(QuarkParserINTEGRAL-80))|(1<<(QuarkParserREAL_NAME-80)))) != 0) {
			{
				p.SetState(269)
				p.expr(0)
			}
			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == QuarkParserCOMMA {
				{
					p.SetState(270)
					p.Match(QuarkParserCOMMA)
				}
				{
					p.SetState(271)
					p.expr(0)
				}

				p.SetState(276)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(279)
			p.Match(QuarkParserRBRACE)
		}

	case 13:
		localctx = NewClockToExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(280)
			p.Match(QuarkParserKW_SIGNAL)
		}
		{
			p.SetState(281)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(282)
			p.Clockexpr()
		}
		{
			p.SetState(283)
			p.Match(QuarkParserRPAREN)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(350)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivModExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(288)

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
					p.SetState(289)
					p.expr(13)
				}

			case 2:
				localctx = NewAddSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(291)

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
					p.SetState(292)
					p.expr(12)
				}

			case 3:
				localctx = NewShiftExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(294)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ShiftExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(QuarkParserOP_LEFT_SHIFT-34))|(1<<(QuarkParserOP_RIGHT_SHIFT-34))|(1<<(QuarkParserOP_ARITH_LEFT_SHIFT-34))|(1<<(QuarkParserOP_ARITH_RIGHT_SHFIT-34)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ShiftExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(295)
					p.expr(11)
				}

			case 4:
				localctx = NewCompareExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(297)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(QuarkParserLANGLE-42))|(1<<(QuarkParserRANGLE-42))|(1<<(QuarkParserOP_LTE-42))|(1<<(QuarkParserOP_GTE-42)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(298)
					p.expr(10)
				}

			case 5:
				localctx = NewBitwiseBinopExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(299)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(300)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BitwiseBinopExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-25)&-(0x1f+1)) == 0 && ((1<<uint((_la-25)))&((1<<(QuarkParserOP_BAND-25))|(1<<(QuarkParserOP_BNAND-25))|(1<<(QuarkParserOP_BOR-25))|(1<<(QuarkParserOP_BNOR-25))|(1<<(QuarkParserOP_XOR-25))|(1<<(QuarkParserOP_XNOR-25)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BitwiseBinopExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(301)
					p.expr(9)
				}

			case 6:
				localctx = NewLogicalBinopExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(302)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(303)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalBinopExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(QuarkParserOP_LAND-45))|(1<<(QuarkParserOP_LOR-45))|(1<<(QuarkParserOP_IMPLICATION-45))|(1<<(QuarkParserOP_EQUIVALENCE-45)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalBinopExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(304)
					p.expr(8)
				}

			case 7:
				localctx = NewTernaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(306)
					p.Match(QuarkParserKW_IF)
				}
				{
					p.SetState(307)
					p.expr(0)
				}
				{
					p.SetState(308)
					p.Match(QuarkParserKW_ELSE)
				}
				{
					p.SetState(309)
					p.expr(7)
				}

			case 8:
				localctx = NewSelectorExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(311)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(312)
					p.Match(QuarkParserDOT)
				}
				{
					p.SetState(313)
					p.Realname()
				}

			case 9:
				localctx = NewFunctionCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(314)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				p.SetState(316)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == QuarkParserPARAM_OPEN {
					{
						p.SetState(315)
						p.Paramarglist()
					}

				}
				{
					p.SetState(318)
					p.Match(QuarkParserLPAREN)
				}
				{
					p.SetState(319)
					p.Callarglist()
				}
				{
					p.SetState(320)
					p.Match(QuarkParserRPAREN)
				}

			case 10:
				localctx = NewSliceExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(322)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(323)
					p.Match(QuarkParserLBRACE)
				}
				p.SetState(325)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(QuarkParserOP_COMPLIMENT-44))|(1<<(QuarkParserOP_LNOT-44))|(1<<(QuarkParserKW_IF-44))|(1<<(QuarkParserKW_MATCH-44))|(1<<(QuarkParserKW_LAMBDA-44)))) != 0) || (((_la-80)&-(0x1f+1)) == 0 && ((1<<uint((_la-80)))&((1<<(QuarkParserKW_NEW-80))|(1<<(QuarkParserKW_SIGNAL-80))|(1<<(QuarkParserDECIMAL-80))|(1<<(QuarkParserINTEGRAL-80))|(1<<(QuarkParserREAL_NAME-80)))) != 0) {
					{
						p.SetState(324)

						var _x = p.expr(0)

						localctx.(*SliceExprContext).msb = _x
					}

				}
				{
					p.SetState(327)
					p.Match(QuarkParserCOLON)
				}
				p.SetState(329)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(328)

						var _x = p.expr(0)

						localctx.(*SliceExprContext).lsb = _x
					}

				}

				p.SetState(332)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == QuarkParserCOLON {
					{
						p.SetState(331)
						p.Match(QuarkParserCOLON)
					}

				}
				p.SetState(335)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(QuarkParserOP_COMPLIMENT-44))|(1<<(QuarkParserOP_LNOT-44))|(1<<(QuarkParserKW_IF-44))|(1<<(QuarkParserKW_MATCH-44))|(1<<(QuarkParserKW_LAMBDA-44)))) != 0) || (((_la-80)&-(0x1f+1)) == 0 && ((1<<uint((_la-80)))&((1<<(QuarkParserKW_NEW-80))|(1<<(QuarkParserKW_SIGNAL-80))|(1<<(QuarkParserDECIMAL-80))|(1<<(QuarkParserINTEGRAL-80))|(1<<(QuarkParserREAL_NAME-80)))) != 0) {
					{
						p.SetState(334)

						var _x = p.expr(0)

						localctx.(*SliceExprContext).step = _x
					}

				}

				{
					p.SetState(337)
					p.Match(QuarkParserRBRACE)
				}

			case 11:
				localctx = NewArrayIndexExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_expr)
				p.SetState(338)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(339)
					p.Match(QuarkParserLBRACE)
				}
				{
					p.SetState(340)
					p.expr(0)
				}
				p.SetState(345)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == QuarkParserCOMMA {
					{
						p.SetState(341)
						p.Match(QuarkParserCOMMA)
					}
					{
						p.SetState(342)
						p.expr(0)
					}

					p.SetState(347)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(348)
					p.Match(QuarkParserRBRACE)
				}

			}

		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
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
		p.SetState(355)
		p.Callarg()
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserCOMMA {
		{
			p.SetState(356)
			p.Match(QuarkParserCOMMA)
		}
		{
			p.SetState(357)
			p.Callarg()
		}

		p.SetState(362)
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

	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNamedCallArgContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(363)
			p.Realname()
		}
		{
			p.SetState(364)
			p.Match(QuarkParserOP_ASSIGN)
		}
		{
			p.SetState(365)
			p.expr(0)
		}

	case 2:
		localctx = NewUnamedCallArgContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(367)
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

func (s *ParamarglistContext) PARAM_OPEN() antlr.TerminalNode {
	return s.GetToken(QuarkParserPARAM_OPEN, 0)
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

func (s *ParamarglistContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
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
		p.SetState(370)
		p.Match(QuarkParserPARAM_OPEN)
	}
	{
		p.SetState(371)
		p.Paramarg()
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserCOMMA {
		{
			p.SetState(372)
			p.Match(QuarkParserCOMMA)
		}
		{
			p.SetState(373)
			p.Paramarg()
		}

		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(379)
		p.Match(QuarkParserRPAREN)
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
		p.SetState(381)
		p.Match(QuarkParserLCURLY)
	}
	{
		p.SetState(382)
		p.Innerconcat()
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == QuarkParserCOMMA {
		{
			p.SetState(383)
			p.Match(QuarkParserCOMMA)
		}
		{
			p.SetState(384)
			p.Innerconcat()
		}

		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(389)
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

	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			p.expr(0)
		}
		{
			p.SetState(392)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(393)
			p.expr(0)
		}
		{
			p.SetState(394)
			p.Match(QuarkParserRCURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
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
		p.SetState(400)
		p.Name()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(417)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
			case 1:
				localctx = NewParameterizedTypeContext(p, NewTypeexprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_typeexpr)
				p.SetState(402)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(403)
					p.Paramarglist()
				}

			case 2:
				localctx = NewArrayTypeContext(p, NewTypeexprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QuarkParserRULE_typeexpr)
				p.SetState(404)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(405)
					p.Match(QuarkParserLBRACE)
				}
				p.SetState(414)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuarkParserLPAREN)|(1<<QuarkParserLBRACE)|(1<<QuarkParserLCURLY))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(QuarkParserOP_COMPLIMENT-44))|(1<<(QuarkParserOP_LNOT-44))|(1<<(QuarkParserKW_IF-44))|(1<<(QuarkParserKW_MATCH-44))|(1<<(QuarkParserKW_LAMBDA-44)))) != 0) || (((_la-80)&-(0x1f+1)) == 0 && ((1<<uint((_la-80)))&((1<<(QuarkParserKW_NEW-80))|(1<<(QuarkParserKW_SIGNAL-80))|(1<<(QuarkParserDECIMAL-80))|(1<<(QuarkParserINTEGRAL-80))|(1<<(QuarkParserREAL_NAME-80)))) != 0) {
					{
						p.SetState(406)
						p.expr(0)
					}
					p.SetState(411)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == QuarkParserCOMMA {
						{
							p.SetState(407)
							p.Match(QuarkParserCOMMA)
						}
						{
							p.SetState(408)
							p.expr(0)
						}

						p.SetState(413)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(416)
					p.Match(QuarkParserRBRACE)
				}

			}

		}
		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
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

	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(422)
			p.Match(QuarkParserKW_TYPE)
		}
		{
			p.SetState(423)
			p.typeexpr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(424)
			p.expr(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(425)
			p.Realname()
		}
		{
			p.SetState(426)
			p.Match(QuarkParserOP_ASSIGN)
		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case QuarkParserLPAREN, QuarkParserLBRACE, QuarkParserLCURLY, QuarkParserOP_COMPLIMENT, QuarkParserOP_LNOT, QuarkParserKW_IF, QuarkParserKW_MATCH, QuarkParserKW_LAMBDA, QuarkParserKW_NEW, QuarkParserKW_SIGNAL, QuarkParserDECIMAL, QuarkParserINTEGRAL, QuarkParserREAL_NAME:
			{
				p.SetState(427)
				p.expr(0)
			}

		case QuarkParserKW_TYPE:
			{
				p.SetState(428)
				p.Match(QuarkParserKW_TYPE)
			}
			{
				p.SetState(429)
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

	p.SetState(472)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QuarkParserKW_IF:
		localctx = NewIfBranchContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(434)
			p.Match(QuarkParserKW_IF)
		}
		{
			p.SetState(435)
			p.expr(0)
		}
		{
			p.SetState(436)
			p.Match(QuarkParserLCURLY)
		}
		{
			p.SetState(437)
			p.Block()
		}
		{
			p.SetState(438)
			p.Match(QuarkParserRCURLY)
		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(439)
					p.Match(QuarkParserKW_ELIF)
				}
				{
					p.SetState(440)
					p.expr(0)
				}
				{
					p.SetState(441)
					p.Match(QuarkParserLCURLY)
				}
				{
					p.SetState(442)
					p.Block()
				}
				{
					p.SetState(443)
					p.Match(QuarkParserRCURLY)
				}

			}
			p.SetState(449)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())
		}
		p.SetState(455)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(450)
				p.Match(QuarkParserKW_ELSE)
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

		}

	case QuarkParserKW_MATCH:
		localctx = NewMatchBranchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(457)
			p.Match(QuarkParserKW_MATCH)
		}
		{
			p.SetState(458)
			p.expr(0)
		}
		{
			p.SetState(459)
			p.Match(QuarkParserLCURLY)
		}
		p.SetState(466)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QuarkParserKW_CASE {
			{
				p.SetState(460)
				p.Match(QuarkParserKW_CASE)
			}
			{
				p.SetState(461)
				p.Pattern()
			}
			{
				p.SetState(462)
				p.Match(QuarkParserLCURLY)
			}
			{
				p.SetState(463)
				p.Block()
			}
			{
				p.SetState(464)
				p.Match(QuarkParserRCURLY)
			}

			p.SetState(468)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(470)
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

	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAtomicPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(474)
			p.Realname()
		}

	case 2:
		localctx = NewParamerterizedTypePatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(475)
			p.Realname()
		}
		{
			p.SetState(476)
			p.Match(QuarkParserLBRACE)
		}
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
		{
			p.SetState(485)
			p.Match(QuarkParserRBRACE)
		}

	case 3:
		localctx = NewArrayPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(487)
			p.Match(QuarkParserLBRACE)
		}
		p.SetState(496)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QuarkParserLBRACE || _la == QuarkParserLCURLY || (((_la-88)&-(0x1f+1)) == 0 && ((1<<uint((_la-88)))&((1<<(QuarkParserDECIMAL-88))|(1<<(QuarkParserINTEGRAL-88))|(1<<(QuarkParserREAL_NAME-88)))) != 0) {
			{
				p.SetState(488)
				p.Pattern()
			}
			p.SetState(493)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == QuarkParserCOMMA {
				{
					p.SetState(489)
					p.Match(QuarkParserCOMMA)
				}
				{
					p.SetState(490)
					p.Pattern()
				}

				p.SetState(495)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(498)
			p.Match(QuarkParserRBRACE)
		}

	case 4:
		localctx = NewLiteralPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(499)
			p.Literal()
		}

	case 5:
		localctx = NewStructPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(500)
			p.Match(QuarkParserLCURLY)
		}

		p.SetState(502)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(501)
				p.typeexpr(0)
			}

		}
		{
			p.SetState(504)
			p.Pattern()
		}

		p.SetState(513)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(506)
				p.Match(QuarkParserCOMMA)
			}
			p.SetState(508)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(507)
					p.typeexpr(0)
				}

			}
			{
				p.SetState(510)
				p.Pattern()
			}

			p.SetState(515)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(516)
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

func (s *ParameterlistContext) PARAM_OPEN() antlr.TerminalNode {
	return s.GetToken(QuarkParserPARAM_OPEN, 0)
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

func (s *ParameterlistContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QuarkParserRPAREN, 0)
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
		p.SetState(520)
		p.Match(QuarkParserPARAM_OPEN)
	}
	{
		p.SetState(521)
		p.Parameterdef()
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserCOMMA {
		{
			p.SetState(522)
			p.Match(QuarkParserCOMMA)
		}
		{
			p.SetState(523)
			p.Parameterdef()
		}

		p.SetState(528)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(529)
		p.Match(QuarkParserRPAREN)
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

	p.SetState(536)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QuarkParserKW_TYPE:
		localctx = NewTypeParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(531)
			p.Match(QuarkParserKW_TYPE)
		}
		{
			p.SetState(532)
			p.typeexpr(0)
		}

	case QuarkParserREAL_NAME:
		localctx = NewValueParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(533)
			p.typeexpr(0)
		}
		{
			p.SetState(534)
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

	p.SetState(553)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QuarkParserREAL_NAME:
		localctx = NewSingleReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(538)
			p.typeexpr(0)
		}

	case QuarkParserLPAREN:
		localctx = NewNamedReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(539)
			p.Match(QuarkParserLPAREN)
		}
		{
			p.SetState(540)
			p.typeexpr(0)
		}
		{
			p.SetState(541)
			p.Realname()
		}
		p.SetState(548)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(542)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(543)
				p.typeexpr(0)
			}
			{
				p.SetState(544)
				p.Realname()
			}

			p.SetState(550)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(551)
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

func (s *ArgumentdefContext) KW_FUTURE() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_FUTURE, 0)
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
	p.SetState(556)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserKW_FUTURE {
		{
			p.SetState(555)
			p.Match(QuarkParserKW_FUTURE)
		}

	}
	{
		p.SetState(558)
		p.typeexpr(0)
	}
	{
		p.SetState(559)
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
		p.SetState(561)
		p.Match(QuarkParserLPAREN)
	}
	p.SetState(570)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserKW_FUTURE || _la == QuarkParserREAL_NAME {
		{
			p.SetState(562)
			p.Argumentdef()
		}
		p.SetState(567)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(563)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(564)
				p.Argumentdef()
			}

			p.SetState(569)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(572)
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

func (s *StructdeclContext) AllFuncdecl() []IFuncdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncdeclContext)(nil)).Elem())
	var tst = make([]IFuncdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncdeclContext)
		}
	}

	return tst
}

func (s *StructdeclContext) Funcdecl(i int) IFuncdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncdeclContext)
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
	p.SetState(577)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(574)
			p.Annotation()
		}

		p.SetState(579)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(580)
		p.Match(QuarkParserKW_STRUCT)
	}
	{
		p.SetState(581)
		p.Realname()
	}
	p.SetState(583)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserPARAM_OPEN {
		{
			p.SetState(582)
			p.Parameterlist()
		}

	}
	p.SetState(594)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserKW_HAS {
		{
			p.SetState(585)
			p.Match(QuarkParserKW_HAS)
		}
		{
			p.SetState(586)
			p.Name()
		}
		p.SetState(591)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QuarkParserCOMMA {
			{
				p.SetState(587)
				p.Match(QuarkParserCOMMA)
			}
			{
				p.SetState(588)
				p.Name()
			}

			p.SetState(593)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(596)
		p.Match(QuarkParserLCURLY)
	}
	p.SetState(601)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserKW_DEF || _la == QuarkParserKW_FUTURE || _la == QuarkParserREAL_NAME || _la == QuarkParserANNOTATION_START {
		p.SetState(599)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 67, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(597)
				p.Fielddecl()
			}

		case 2:
			{
				p.SetState(598)
				p.Funcdecl()
			}

		}

		p.SetState(603)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(604)
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
	p.SetState(609)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(606)
			p.Annotation()
		}

		p.SetState(611)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(613)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserKW_FUTURE {
		{
			p.SetState(612)
			p.Match(QuarkParserKW_FUTURE)
		}

	}
	{
		p.SetState(615)
		p.typeexpr(0)
	}
	{
		p.SetState(616)
		p.Realname()
	}
	{
		p.SetState(617)
		p.Match(QuarkParserSEMI)
	}

	return localctx
}

// IFuncsigContext is an interface to support dynamic dispatch.
type IFuncsigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncsigContext differentiates from other interfaces.
	IsFuncsigContext()
}

type FuncsigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncsigContext() *FuncsigContext {
	var p = new(FuncsigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_funcsig
	return p
}

func (*FuncsigContext) IsFuncsigContext() {}

func NewFuncsigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncsigContext {
	var p = new(FuncsigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_funcsig

	return p
}

func (s *FuncsigContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncsigContext) KW_DEF() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_DEF, 0)
}

func (s *FuncsigContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *FuncsigContext) Argumentlist() IArgumentlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentlistContext)
}

func (s *FuncsigContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *FuncsigContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *FuncsigContext) Parameterlist() IParameterlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterlistContext)
}

func (s *FuncsigContext) COLON() antlr.TerminalNode {
	return s.GetToken(QuarkParserCOLON, 0)
}

func (s *FuncsigContext) Returnlist() IReturnlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnlistContext)
}

func (s *FuncsigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncsigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncsigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterFuncsig(s)
	}
}

func (s *FuncsigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitFuncsig(s)
	}
}

func (s *FuncsigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitFuncsig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Funcsig() (localctx IFuncsigContext) {
	localctx = NewFuncsigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, QuarkParserRULE_funcsig)
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
	p.SetState(622)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(619)
			p.Annotation()
		}

		p.SetState(624)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(625)
		p.Match(QuarkParserKW_DEF)
	}
	{
		p.SetState(626)
		p.Realname()
	}
	p.SetState(628)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserPARAM_OPEN {
		{
			p.SetState(627)
			p.Parameterlist()
		}

	}
	{
		p.SetState(630)
		p.Argumentlist()
	}
	p.SetState(633)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserCOLON {
		{
			p.SetState(631)
			p.Match(QuarkParserCOLON)
		}
		{
			p.SetState(632)
			p.Returnlist()
		}

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

func (s *FuncdeclContext) Funcsig() IFuncsigContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncsigContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncsigContext)
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
	p.EnterRule(localctx, 56, QuarkParserRULE_funcdecl)

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
		p.SetState(635)
		p.Funcsig()
	}
	{
		p.SetState(636)
		p.Match(QuarkParserLCURLY)
	}
	{
		p.SetState(637)
		p.Block()
	}
	{
		p.SetState(638)
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
	p.EnterRule(localctx, 58, QuarkParserRULE_moduledecl)
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
	p.SetState(643)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(640)
			p.Annotation()
		}

		p.SetState(645)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(646)
		p.Match(QuarkParserKW_MODULE)
	}
	{
		p.SetState(647)
		p.Realname()
	}
	p.SetState(649)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserPARAM_OPEN {
		{
			p.SetState(648)
			p.Parameterlist()
		}

	}
	p.SetState(652)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserLPAREN {
		{
			p.SetState(651)
			p.Argumentlist()
		}

	}
	p.SetState(656)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserCOLON {
		{
			p.SetState(654)
			p.Match(QuarkParserCOLON)
		}
		{
			p.SetState(655)
			p.Returnlist()
		}

	}
	{
		p.SetState(658)
		p.Match(QuarkParserLCURLY)
	}
	{
		p.SetState(659)
		p.Block()
	}
	{
		p.SetState(660)
		p.Match(QuarkParserRCURLY)
	}

	return localctx
}

// ITraitdeclContext is an interface to support dynamic dispatch.
type ITraitdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTraitdeclContext differentiates from other interfaces.
	IsTraitdeclContext()
}

type TraitdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTraitdeclContext() *TraitdeclContext {
	var p = new(TraitdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarkParserRULE_traitdecl
	return p
}

func (*TraitdeclContext) IsTraitdeclContext() {}

func NewTraitdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TraitdeclContext {
	var p = new(TraitdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarkParserRULE_traitdecl

	return p
}

func (s *TraitdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TraitdeclContext) KW_TRAIT() antlr.TerminalNode {
	return s.GetToken(QuarkParserKW_TRAIT, 0)
}

func (s *TraitdeclContext) Realname() IRealnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnameContext)
}

func (s *TraitdeclContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserLCURLY, 0)
}

func (s *TraitdeclContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(QuarkParserRCURLY, 0)
}

func (s *TraitdeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *TraitdeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *TraitdeclContext) Parameterlist() IParameterlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterlistContext)
}

func (s *TraitdeclContext) AllFuncsig() []IFuncsigContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncsigContext)(nil)).Elem())
	var tst = make([]IFuncsigContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncsigContext)
		}
	}

	return tst
}

func (s *TraitdeclContext) Funcsig(i int) IFuncsigContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncsigContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncsigContext)
}

func (s *TraitdeclContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(QuarkParserSEMI)
}

func (s *TraitdeclContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(QuarkParserSEMI, i)
}

func (s *TraitdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TraitdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TraitdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.EnterTraitdecl(s)
	}
}

func (s *TraitdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarkParserListener); ok {
		listenerT.ExitTraitdecl(s)
	}
}

func (s *TraitdeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QuarkParserVisitor:
		return t.VisitTraitdecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QuarkParser) Traitdecl() (localctx ITraitdeclContext) {
	localctx = NewTraitdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, QuarkParserRULE_traitdecl)
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
	p.SetState(665)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserANNOTATION_START {
		{
			p.SetState(662)
			p.Annotation()
		}

		p.SetState(667)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(668)
		p.Match(QuarkParserKW_TRAIT)
	}
	{
		p.SetState(669)
		p.Realname()
	}
	p.SetState(671)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuarkParserPARAM_OPEN {
		{
			p.SetState(670)
			p.Parameterlist()
		}

	}
	{
		p.SetState(673)
		p.Match(QuarkParserLCURLY)
	}
	p.SetState(679)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QuarkParserKW_DEF || _la == QuarkParserANNOTATION_START {
		{
			p.SetState(674)
			p.Funcsig()
		}
		{
			p.SetState(675)
			p.Match(QuarkParserSEMI)
		}

		p.SetState(681)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(682)
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
	p.EnterRule(localctx, 62, QuarkParserRULE_annotation)

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
		p.SetState(684)
		p.Match(QuarkParserANNOTATION_START)
	}
	{
		p.SetState(685)
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

func (s *LiteralContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(QuarkParserDECIMAL, 0)
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
	p.EnterRule(localctx, 64, QuarkParserRULE_literal)
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
		p.SetState(687)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QuarkParserDECIMAL || _la == QuarkParserINTEGRAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *QuarkParser) Typeexpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
