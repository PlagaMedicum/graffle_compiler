// Code generated from /home/plagamedicum/Documents/yapis/graffle/GraffleParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // GraffleParser

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 69, 573,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 3, 2, 5, 2, 106, 10, 2, 3, 2, 7,
	2, 109, 10, 2, 12, 2, 14, 2, 112, 11, 2, 3, 2, 3, 2, 5, 2, 116, 10, 2,
	3, 3, 5, 3, 119, 10, 3, 3, 3, 3, 3, 6, 3, 123, 10, 3, 13, 3, 14, 3, 124,
	3, 3, 5, 3, 128, 10, 3, 6, 3, 130, 10, 3, 13, 3, 14, 3, 131, 3, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 138, 10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 143, 10, 5, 12, 5,
	14, 5, 146, 11, 5, 3, 5, 5, 5, 149, 10, 5, 3, 6, 3, 6, 5, 6, 153, 10, 6,
	3, 7, 3, 7, 3, 7, 5, 7, 158, 10, 7, 3, 8, 3, 8, 3, 8, 5, 8, 163, 10, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 169, 10, 8, 5, 8, 171, 10, 8, 3, 8, 6, 8,
	174, 10, 8, 13, 8, 14, 8, 175, 3, 8, 5, 8, 179, 10, 8, 3, 8, 3, 8, 5, 8,
	183, 10, 8, 3, 8, 6, 8, 186, 10, 8, 13, 8, 14, 8, 187, 5, 8, 190, 10, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 197, 10, 9, 12, 9, 14, 9, 200, 11,
	9, 3, 9, 5, 9, 203, 10, 9, 3, 9, 3, 9, 5, 9, 207, 10, 9, 3, 9, 6, 9, 210,
	10, 9, 13, 9, 14, 9, 211, 5, 9, 214, 10, 9, 3, 9, 3, 9, 3, 10, 5, 10, 219,
	10, 10, 3, 10, 3, 10, 3, 10, 5, 10, 224, 10, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 230, 10, 10, 5, 10, 232, 10, 10, 3, 10, 6, 10, 235, 10, 10,
	13, 10, 14, 10, 236, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 5, 13, 250, 10, 13, 3, 14, 3, 14, 3, 14, 5, 14,
	255, 10, 14, 3, 15, 3, 15, 3, 15, 5, 15, 260, 10, 15, 3, 16, 3, 16, 3,
	16, 5, 16, 265, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	5, 16, 274, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 280, 10, 16, 3,
	16, 5, 16, 283, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 289, 10, 16,
	5, 16, 291, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 298, 10,
	17, 3, 18, 5, 18, 301, 10, 18, 3, 18, 3, 18, 5, 18, 305, 10, 18, 3, 18,
	3, 18, 5, 18, 309, 10, 18, 3, 18, 3, 18, 5, 18, 313, 10, 18, 3, 18, 5,
	18, 316, 10, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 330, 10, 21, 12, 21, 14, 21, 333, 11,
	21, 5, 21, 335, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 341, 10, 21,
	3, 21, 5, 21, 344, 10, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 359, 10, 24, 12, 24,
	14, 24, 362, 11, 24, 5, 24, 364, 10, 24, 3, 24, 3, 24, 5, 24, 368, 10,
	24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	5, 25, 380, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 395, 10, 26, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 403, 10, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 436, 10, 31, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 7, 32, 450, 10, 32, 12, 32, 14, 32, 453, 11, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 5, 32, 459, 10, 32, 3, 33, 3, 33, 5, 33, 463, 10, 33, 3, 33, 3,
	33, 5, 33, 467, 10, 33, 3, 33, 3, 33, 5, 33, 471, 10, 33, 5, 33, 473, 10,
	33, 3, 34, 3, 34, 5, 34, 477, 10, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 489, 10, 35, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 500, 10, 36, 3, 37,
	3, 37, 5, 37, 504, 10, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 41, 3, 41, 5, 41, 516, 10, 41, 3, 42, 3, 42, 3, 43, 3, 43,
	3, 44, 3, 44, 5, 44, 524, 10, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 5, 45, 532, 10, 45, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47,
	3, 47, 7, 47, 542, 10, 47, 12, 47, 14, 47, 545, 11, 47, 5, 47, 547, 10,
	47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 5, 49, 557,
	10, 49, 3, 50, 5, 50, 560, 10, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 5,
	52, 567, 10, 52, 3, 52, 3, 52, 5, 52, 571, 10, 52, 3, 52, 2, 2, 53, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
	78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 2, 8, 4, 2, 52, 52,
	54, 54, 3, 2, 25, 35, 3, 2, 21, 24, 3, 2, 17, 20, 3, 2, 4, 5, 5, 2, 3,
	3, 6, 6, 65, 65, 2, 624, 2, 105, 3, 2, 2, 2, 4, 118, 3, 2, 2, 2, 6, 137,
	3, 2, 2, 2, 8, 139, 3, 2, 2, 2, 10, 152, 3, 2, 2, 2, 12, 157, 3, 2, 2,
	2, 14, 159, 3, 2, 2, 2, 16, 193, 3, 2, 2, 2, 18, 218, 3, 2, 2, 2, 20, 238,
	3, 2, 2, 2, 22, 241, 3, 2, 2, 2, 24, 249, 3, 2, 2, 2, 26, 251, 3, 2, 2,
	2, 28, 256, 3, 2, 2, 2, 30, 290, 3, 2, 2, 2, 32, 292, 3, 2, 2, 2, 34, 315,
	3, 2, 2, 2, 36, 317, 3, 2, 2, 2, 38, 320, 3, 2, 2, 2, 40, 324, 3, 2, 2,
	2, 42, 345, 3, 2, 2, 2, 44, 348, 3, 2, 2, 2, 46, 353, 3, 2, 2, 2, 48, 379,
	3, 2, 2, 2, 50, 394, 3, 2, 2, 2, 52, 402, 3, 2, 2, 2, 54, 404, 3, 2, 2,
	2, 56, 410, 3, 2, 2, 2, 58, 416, 3, 2, 2, 2, 60, 435, 3, 2, 2, 2, 62, 458,
	3, 2, 2, 2, 64, 472, 3, 2, 2, 2, 66, 476, 3, 2, 2, 2, 68, 488, 3, 2, 2,
	2, 70, 499, 3, 2, 2, 2, 72, 503, 3, 2, 2, 2, 74, 505, 3, 2, 2, 2, 76, 507,
	3, 2, 2, 2, 78, 509, 3, 2, 2, 2, 80, 515, 3, 2, 2, 2, 82, 517, 3, 2, 2,
	2, 84, 519, 3, 2, 2, 2, 86, 523, 3, 2, 2, 2, 88, 531, 3, 2, 2, 2, 90, 533,
	3, 2, 2, 2, 92, 536, 3, 2, 2, 2, 94, 550, 3, 2, 2, 2, 96, 556, 3, 2, 2,
	2, 98, 559, 3, 2, 2, 2, 100, 563, 3, 2, 2, 2, 102, 566, 3, 2, 2, 2, 104,
	106, 7, 54, 2, 2, 105, 104, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 110,
	3, 2, 2, 2, 107, 109, 5, 34, 18, 2, 108, 107, 3, 2, 2, 2, 109, 112, 3,
	2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 113, 3, 2, 2,
	2, 112, 110, 3, 2, 2, 2, 113, 115, 5, 4, 3, 2, 114, 116, 7, 2, 2, 3, 115,
	114, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 3, 3, 2, 2, 2, 117, 119, 7,
	54, 2, 2, 118, 117, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 129, 3, 2, 2,
	2, 120, 127, 5, 6, 4, 2, 121, 123, 9, 2, 2, 2, 122, 121, 3, 2, 2, 2, 123,
	124, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 128,
	3, 2, 2, 2, 126, 128, 7, 2, 2, 3, 127, 122, 3, 2, 2, 2, 127, 126, 3, 2,
	2, 2, 128, 130, 3, 2, 2, 2, 129, 120, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2,
	131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 5, 3, 2, 2, 2, 133, 138,
	5, 14, 8, 2, 134, 138, 5, 16, 9, 2, 135, 138, 5, 22, 12, 2, 136, 138, 5,
	8, 5, 2, 137, 133, 3, 2, 2, 2, 137, 134, 3, 2, 2, 2, 137, 135, 3, 2, 2,
	2, 137, 136, 3, 2, 2, 2, 138, 7, 3, 2, 2, 2, 139, 144, 5, 10, 6, 2, 140,
	141, 7, 52, 2, 2, 141, 143, 5, 10, 6, 2, 142, 140, 3, 2, 2, 2, 143, 146,
	3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 148, 3, 2,
	2, 2, 146, 144, 3, 2, 2, 2, 147, 149, 7, 52, 2, 2, 148, 147, 3, 2, 2, 2,
	148, 149, 3, 2, 2, 2, 149, 9, 3, 2, 2, 2, 150, 153, 5, 12, 7, 2, 151, 153,
	5, 20, 11, 2, 152, 150, 3, 2, 2, 2, 152, 151, 3, 2, 2, 2, 153, 11, 3, 2,
	2, 2, 154, 158, 5, 48, 25, 2, 155, 158, 5, 92, 47, 2, 156, 158, 5, 86,
	44, 2, 157, 154, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 156, 3, 2, 2, 2,
	158, 13, 3, 2, 2, 2, 159, 160, 7, 37, 2, 2, 160, 170, 5, 70, 36, 2, 161,
	163, 7, 53, 2, 2, 162, 161, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 168,
	3, 2, 2, 2, 164, 169, 7, 38, 2, 2, 165, 169, 7, 49, 2, 2, 166, 167, 7,
	38, 2, 2, 167, 169, 7, 49, 2, 2, 168, 164, 3, 2, 2, 2, 168, 165, 3, 2,
	2, 2, 168, 166, 3, 2, 2, 2, 169, 171, 3, 2, 2, 2, 170, 162, 3, 2, 2, 2,
	170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2, 172, 174, 5, 4, 3, 2, 173,
	172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176,
	3, 2, 2, 2, 176, 178, 3, 2, 2, 2, 177, 179, 7, 54, 2, 2, 178, 177, 3, 2,
	2, 2, 178, 179, 3, 2, 2, 2, 179, 189, 3, 2, 2, 2, 180, 182, 7, 39, 2, 2,
	181, 183, 7, 49, 2, 2, 182, 181, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183,
	185, 3, 2, 2, 2, 184, 186, 5, 4, 3, 2, 185, 184, 3, 2, 2, 2, 186, 187,
	3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 190, 3, 2,
	2, 2, 189, 180, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2,
	191, 192, 5, 102, 52, 2, 192, 15, 3, 2, 2, 2, 193, 194, 7, 37, 2, 2, 194,
	198, 5, 96, 49, 2, 195, 197, 5, 18, 10, 2, 196, 195, 3, 2, 2, 2, 197, 200,
	3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 202, 3, 2,
	2, 2, 200, 198, 3, 2, 2, 2, 201, 203, 7, 54, 2, 2, 202, 201, 3, 2, 2, 2,
	202, 203, 3, 2, 2, 2, 203, 213, 3, 2, 2, 2, 204, 206, 7, 41, 2, 2, 205,
	207, 7, 49, 2, 2, 206, 205, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 209,
	3, 2, 2, 2, 208, 210, 5, 4, 3, 2, 209, 208, 3, 2, 2, 2, 210, 211, 3, 2,
	2, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 214, 3, 2, 2, 2,
	213, 204, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215,
	216, 5, 102, 52, 2, 216, 17, 3, 2, 2, 2, 217, 219, 7, 54, 2, 2, 218, 217,
	3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 7, 40,
	2, 2, 221, 231, 5, 96, 49, 2, 222, 224, 7, 53, 2, 2, 223, 222, 3, 2, 2,
	2, 223, 224, 3, 2, 2, 2, 224, 229, 3, 2, 2, 2, 225, 230, 7, 38, 2, 2, 226,
	230, 7, 49, 2, 2, 227, 228, 7, 38, 2, 2, 228, 230, 7, 49, 2, 2, 229, 225,
	3, 2, 2, 2, 229, 226, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 232, 3, 2,
	2, 2, 231, 223, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 234, 3, 2, 2, 2,
	233, 235, 5, 4, 3, 2, 234, 233, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236,
	234, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 19, 3, 2, 2, 2, 238, 239, 5,
	24, 13, 2, 239, 240, 5, 8, 5, 2, 240, 21, 3, 2, 2, 2, 241, 242, 5, 24,
	13, 2, 242, 243, 5, 4, 3, 2, 243, 244, 5, 102, 52, 2, 244, 23, 3, 2, 2,
	2, 245, 250, 5, 26, 14, 2, 246, 250, 5, 28, 15, 2, 247, 250, 5, 30, 16,
	2, 248, 250, 5, 32, 17, 2, 249, 245, 3, 2, 2, 2, 249, 246, 3, 2, 2, 2,
	249, 247, 3, 2, 2, 2, 249, 248, 3, 2, 2, 2, 250, 25, 3, 2, 2, 2, 251, 252,
	7, 42, 2, 2, 252, 254, 5, 70, 36, 2, 253, 255, 7, 49, 2, 2, 254, 253, 3,
	2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 27, 3, 2, 2, 2, 256, 257, 7, 43, 2,
	2, 257, 259, 5, 70, 36, 2, 258, 260, 7, 49, 2, 2, 259, 258, 3, 2, 2, 2,
	259, 260, 3, 2, 2, 2, 260, 29, 3, 2, 2, 2, 261, 262, 7, 44, 2, 2, 262,
	264, 5, 70, 36, 2, 263, 265, 7, 49, 2, 2, 264, 263, 3, 2, 2, 2, 264, 265,
	3, 2, 2, 2, 265, 291, 3, 2, 2, 2, 266, 267, 7, 44, 2, 2, 267, 268, 5, 12,
	7, 2, 268, 269, 7, 53, 2, 2, 269, 270, 5, 70, 36, 2, 270, 271, 7, 53, 2,
	2, 271, 273, 5, 12, 7, 2, 272, 274, 7, 49, 2, 2, 273, 272, 3, 2, 2, 2,
	273, 274, 3, 2, 2, 2, 274, 291, 3, 2, 2, 2, 275, 276, 7, 44, 2, 2, 276,
	277, 5, 98, 50, 2, 277, 279, 7, 45, 2, 2, 278, 280, 7, 46, 2, 2, 279, 278,
	3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 282, 3, 2, 2, 2, 281, 283, 7, 47,
	2, 2, 282, 281, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2,
	284, 285, 5, 66, 34, 2, 285, 286, 7, 48, 2, 2, 286, 288, 5, 66, 34, 2,
	287, 289, 7, 49, 2, 2, 288, 287, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289,
	291, 3, 2, 2, 2, 290, 261, 3, 2, 2, 2, 290, 266, 3, 2, 2, 2, 290, 275,
	3, 2, 2, 2, 291, 31, 3, 2, 2, 2, 292, 293, 7, 47, 2, 2, 293, 294, 5, 66,
	34, 2, 294, 295, 7, 48, 2, 2, 295, 297, 5, 66, 34, 2, 296, 298, 7, 49,
	2, 2, 297, 296, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 33, 3, 2, 2, 2,
	299, 301, 7, 54, 2, 2, 300, 299, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301,
	302, 3, 2, 2, 2, 302, 316, 5, 36, 19, 2, 303, 305, 7, 54, 2, 2, 304, 303,
	3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 316, 5, 38,
	20, 2, 307, 309, 7, 54, 2, 2, 308, 307, 3, 2, 2, 2, 308, 309, 3, 2, 2,
	2, 309, 310, 3, 2, 2, 2, 310, 316, 5, 42, 22, 2, 311, 313, 7, 54, 2, 2,
	312, 311, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314,
	316, 5, 44, 23, 2, 315, 300, 3, 2, 2, 2, 315, 304, 3, 2, 2, 2, 315, 308,
	3, 2, 2, 2, 315, 312, 3, 2, 2, 2, 316, 35, 3, 2, 2, 2, 317, 318, 5, 40,
	21, 2, 318, 319, 5, 8, 5, 2, 319, 37, 3, 2, 2, 2, 320, 321, 5, 40, 21,
	2, 321, 322, 5, 4, 3, 2, 322, 323, 5, 102, 52, 2, 323, 39, 3, 2, 2, 2,
	324, 325, 7, 63, 2, 2, 325, 334, 7, 7, 2, 2, 326, 331, 5, 98, 50, 2, 327,
	328, 7, 53, 2, 2, 328, 330, 5, 98, 50, 2, 329, 327, 3, 2, 2, 2, 330, 333,
	3, 2, 2, 2, 331, 329, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 335, 3, 2,
	2, 2, 333, 331, 3, 2, 2, 2, 334, 326, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2,
	335, 336, 3, 2, 2, 2, 336, 337, 7, 8, 2, 2, 337, 338, 7, 16, 2, 2, 338,
	343, 5, 96, 49, 2, 339, 341, 7, 53, 2, 2, 340, 339, 3, 2, 2, 2, 340, 341,
	3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 342, 344, 7, 62, 2, 2, 343, 340, 3, 2,
	2, 2, 343, 344, 3, 2, 2, 2, 344, 41, 3, 2, 2, 2, 345, 346, 5, 46, 24, 2,
	346, 347, 5, 8, 5, 2, 347, 43, 3, 2, 2, 2, 348, 349, 5, 46, 24, 2, 349,
	350, 7, 54, 2, 2, 350, 351, 5, 4, 3, 2, 351, 352, 5, 102, 52, 2, 352, 45,
	3, 2, 2, 2, 353, 354, 7, 63, 2, 2, 354, 363, 7, 7, 2, 2, 355, 360, 5, 98,
	50, 2, 356, 357, 7, 53, 2, 2, 357, 359, 5, 98, 50, 2, 358, 356, 3, 2, 2,
	2, 359, 362, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361,
	364, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 363, 355, 3, 2, 2, 2, 363, 364,
	3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365, 367, 7, 8, 2, 2, 366, 368, 7, 16,
	2, 2, 367, 366, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 47, 3, 2, 2, 2,
	369, 370, 7, 63, 2, 2, 370, 371, 7, 16, 2, 2, 371, 380, 5, 98, 50, 2, 372,
	373, 7, 63, 2, 2, 373, 374, 7, 16, 2, 2, 374, 380, 5, 66, 34, 2, 375, 380,
	5, 50, 26, 2, 376, 380, 5, 60, 31, 2, 377, 380, 5, 62, 32, 2, 378, 380,
	5, 64, 33, 2, 379, 369, 3, 2, 2, 2, 379, 372, 3, 2, 2, 2, 379, 375, 3,
	2, 2, 2, 379, 376, 3, 2, 2, 2, 379, 377, 3, 2, 2, 2, 379, 378, 3, 2, 2,
	2, 380, 49, 3, 2, 2, 2, 381, 382, 7, 63, 2, 2, 382, 383, 7, 16, 2, 2, 383,
	384, 7, 61, 2, 2, 384, 385, 7, 7, 2, 2, 385, 386, 5, 96, 49, 2, 386, 387,
	7, 8, 2, 2, 387, 395, 3, 2, 2, 2, 388, 389, 7, 63, 2, 2, 389, 390, 7, 16,
	2, 2, 390, 391, 5, 96, 49, 2, 391, 392, 5, 52, 27, 2, 392, 393, 5, 96,
	49, 2, 393, 395, 3, 2, 2, 2, 394, 381, 3, 2, 2, 2, 394, 388, 3, 2, 2, 2,
	395, 51, 3, 2, 2, 2, 396, 403, 7, 13, 2, 2, 397, 403, 5, 54, 28, 2, 398,
	403, 7, 14, 2, 2, 399, 403, 5, 56, 29, 2, 400, 403, 7, 15, 2, 2, 401, 403,
	5, 58, 30, 2, 402, 396, 3, 2, 2, 2, 402, 397, 3, 2, 2, 2, 402, 398, 3,
	2, 2, 2, 402, 399, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2, 402, 401, 3, 2, 2,
	2, 403, 53, 3, 2, 2, 2, 404, 405, 7, 22, 2, 2, 405, 406, 7, 11, 2, 2, 406,
	407, 7, 65, 2, 2, 407, 408, 7, 12, 2, 2, 408, 409, 7, 13, 2, 2, 409, 55,
	3, 2, 2, 2, 410, 411, 7, 14, 2, 2, 411, 412, 7, 11, 2, 2, 412, 413, 7,
	65, 2, 2, 413, 414, 7, 12, 2, 2, 414, 415, 7, 22, 2, 2, 415, 57, 3, 2,
	2, 2, 416, 417, 7, 22, 2, 2, 417, 418, 7, 11, 2, 2, 418, 419, 7, 65, 2,
	2, 419, 420, 7, 12, 2, 2, 420, 421, 7, 22, 2, 2, 421, 59, 3, 2, 2, 2, 422,
	423, 7, 63, 2, 2, 423, 424, 7, 16, 2, 2, 424, 425, 7, 7, 2, 2, 425, 426,
	7, 60, 2, 2, 426, 427, 7, 8, 2, 2, 427, 436, 5, 96, 49, 2, 428, 429, 7,
	63, 2, 2, 429, 430, 7, 16, 2, 2, 430, 436, 5, 96, 49, 2, 431, 432, 7, 63,
	2, 2, 432, 433, 5, 84, 43, 2, 433, 434, 5, 96, 49, 2, 434, 436, 3, 2, 2,
	2, 435, 422, 3, 2, 2, 2, 435, 428, 3, 2, 2, 2, 435, 431, 3, 2, 2, 2, 436,
	61, 3, 2, 2, 2, 437, 438, 7, 63, 2, 2, 438, 439, 7, 16, 2, 2, 439, 440,
	7, 59, 2, 2, 440, 441, 7, 7, 2, 2, 441, 442, 5, 96, 49, 2, 442, 443, 7,
	8, 2, 2, 443, 459, 3, 2, 2, 2, 444, 445, 7, 63, 2, 2, 445, 446, 7, 16,
	2, 2, 446, 451, 5, 96, 49, 2, 447, 448, 7, 53, 2, 2, 448, 450, 5, 96, 49,
	2, 449, 447, 3, 2, 2, 2, 450, 453, 3, 2, 2, 2, 451, 449, 3, 2, 2, 2, 451,
	452, 3, 2, 2, 2, 452, 459, 3, 2, 2, 2, 453, 451, 3, 2, 2, 2, 454, 455,
	7, 63, 2, 2, 455, 456, 5, 84, 43, 2, 456, 457, 5, 96, 49, 2, 457, 459,
	3, 2, 2, 2, 458, 437, 3, 2, 2, 2, 458, 444, 3, 2, 2, 2, 458, 454, 3, 2,
	2, 2, 459, 63, 3, 2, 2, 2, 460, 462, 5, 60, 31, 2, 461, 463, 5, 94, 48,
	2, 462, 461, 3, 2, 2, 2, 462, 463, 3, 2, 2, 2, 463, 473, 3, 2, 2, 2, 464,
	466, 5, 50, 26, 2, 465, 467, 5, 94, 48, 2, 466, 465, 3, 2, 2, 2, 466, 467,
	3, 2, 2, 2, 467, 473, 3, 2, 2, 2, 468, 470, 5, 62, 32, 2, 469, 471, 5,
	94, 48, 2, 470, 469, 3, 2, 2, 2, 470, 471, 3, 2, 2, 2, 471, 473, 3, 2,
	2, 2, 472, 460, 3, 2, 2, 2, 472, 464, 3, 2, 2, 2, 472, 468, 3, 2, 2, 2,
	473, 65, 3, 2, 2, 2, 474, 477, 5, 98, 50, 2, 475, 477, 5, 68, 35, 2, 476,
	474, 3, 2, 2, 2, 476, 475, 3, 2, 2, 2, 477, 67, 3, 2, 2, 2, 478, 489, 5,
	92, 47, 2, 479, 489, 5, 100, 51, 2, 480, 481, 7, 7, 2, 2, 481, 482, 5,
	70, 36, 2, 482, 483, 7, 8, 2, 2, 483, 489, 3, 2, 2, 2, 484, 485, 7, 7,
	2, 2, 485, 486, 5, 78, 40, 2, 486, 487, 7, 8, 2, 2, 487, 489, 3, 2, 2,
	2, 488, 478, 3, 2, 2, 2, 488, 479, 3, 2, 2, 2, 488, 480, 3, 2, 2, 2, 488,
	484, 3, 2, 2, 2, 489, 69, 3, 2, 2, 2, 490, 491, 5, 66, 34, 2, 491, 492,
	5, 74, 38, 2, 492, 493, 5, 72, 37, 2, 493, 500, 3, 2, 2, 2, 494, 495, 5,
	76, 39, 2, 495, 496, 5, 72, 37, 2, 496, 500, 3, 2, 2, 2, 497, 500, 5, 78,
	40, 2, 498, 500, 5, 66, 34, 2, 499, 490, 3, 2, 2, 2, 499, 494, 3, 2, 2,
	2, 499, 497, 3, 2, 2, 2, 499, 498, 3, 2, 2, 2, 500, 71, 3, 2, 2, 2, 501,
	504, 5, 66, 34, 2, 502, 504, 5, 70, 36, 2, 503, 501, 3, 2, 2, 2, 503, 502,
	3, 2, 2, 2, 504, 73, 3, 2, 2, 2, 505, 506, 9, 3, 2, 2, 506, 75, 3, 2, 2,
	2, 507, 508, 7, 36, 2, 2, 508, 77, 3, 2, 2, 2, 509, 510, 5, 66, 34, 2,
	510, 511, 5, 82, 42, 2, 511, 512, 5, 80, 41, 2, 512, 79, 3, 2, 2, 2, 513,
	516, 5, 66, 34, 2, 514, 516, 5, 78, 40, 2, 515, 513, 3, 2, 2, 2, 515, 514,
	3, 2, 2, 2, 516, 81, 3, 2, 2, 2, 517, 518, 9, 4, 2, 2, 518, 83, 3, 2, 2,
	2, 519, 520, 9, 5, 2, 2, 520, 85, 3, 2, 2, 2, 521, 524, 5, 88, 45, 2, 522,
	524, 5, 90, 46, 2, 523, 521, 3, 2, 2, 2, 523, 522, 3, 2, 2, 2, 524, 87,
	3, 2, 2, 2, 525, 526, 7, 57, 2, 2, 526, 532, 5, 98, 50, 2, 527, 528, 7,
	57, 2, 2, 528, 532, 5, 96, 49, 2, 529, 530, 7, 57, 2, 2, 530, 532, 5, 92,
	47, 2, 531, 525, 3, 2, 2, 2, 531, 527, 3, 2, 2, 2, 531, 529, 3, 2, 2, 2,
	532, 89, 3, 2, 2, 2, 533, 534, 7, 58, 2, 2, 534, 535, 7, 63, 2, 2, 535,
	91, 3, 2, 2, 2, 536, 537, 7, 63, 2, 2, 537, 546, 7, 7, 2, 2, 538, 543,
	5, 96, 49, 2, 539, 540, 7, 53, 2, 2, 540, 542, 5, 96, 49, 2, 541, 539,
	3, 2, 2, 2, 542, 545, 3, 2, 2, 2, 543, 541, 3, 2, 2, 2, 543, 544, 3, 2,
	2, 2, 544, 547, 3, 2, 2, 2, 545, 543, 3, 2, 2, 2, 546, 538, 3, 2, 2, 2,
	546, 547, 3, 2, 2, 2, 547, 548, 3, 2, 2, 2, 548, 549, 7, 8, 2, 2, 549,
	93, 3, 2, 2, 2, 550, 551, 9, 6, 2, 2, 551, 95, 3, 2, 2, 2, 552, 557, 5,
	70, 36, 2, 553, 557, 5, 78, 40, 2, 554, 557, 5, 66, 34, 2, 555, 557, 5,
	100, 51, 2, 556, 552, 3, 2, 2, 2, 556, 553, 3, 2, 2, 2, 556, 554, 3, 2,
	2, 2, 556, 555, 3, 2, 2, 2, 557, 97, 3, 2, 2, 2, 558, 560, 7, 22, 2, 2,
	559, 558, 3, 2, 2, 2, 559, 560, 3, 2, 2, 2, 560, 561, 3, 2, 2, 2, 561,
	562, 7, 63, 2, 2, 562, 99, 3, 2, 2, 2, 563, 564, 9, 7, 2, 2, 564, 101,
	3, 2, 2, 2, 565, 567, 7, 54, 2, 2, 566, 565, 3, 2, 2, 2, 566, 567, 3, 2,
	2, 2, 567, 568, 3, 2, 2, 2, 568, 570, 7, 56, 2, 2, 569, 571, 7, 54, 2,
	2, 570, 569, 3, 2, 2, 2, 570, 571, 3, 2, 2, 2, 571, 103, 3, 2, 2, 2, 77,
	105, 110, 115, 118, 124, 127, 131, 137, 144, 148, 152, 157, 162, 168, 170,
	175, 178, 182, 187, 189, 198, 202, 206, 211, 213, 218, 223, 229, 231, 236,
	249, 254, 259, 264, 273, 279, 282, 288, 290, 297, 300, 304, 308, 312, 315,
	331, 334, 340, 343, 360, 363, 367, 379, 394, 402, 435, 451, 458, 462, 466,
	470, 472, 476, 488, 499, 503, 515, 523, 531, 543, 546, 556, 559, 566, 570,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "'->'", "'<-'",
	"'--'", "'='", "'+='", "'-='", "'*='", "'/='", "'+'", "'-'", "'*'", "'/'",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "'.'", "','",
}
var symbolicNames = []string{
	"", "STRING", "LABEL", "ML_LABEL", "BOOL", "L_PAREN", "R_PAREN", "L_CURLY",
	"R_CURLY", "L_BRACKET", "R_BRACKET", "OR_ARC_LR", "OR_ARC_RL", "UNOR_ARC",
	"ASSIGN", "ADD_ASSIGN", "SUB_ASSIGN", "MULT_ASSIGN", "DIV_ASSIGN", "ADD",
	"SUB", "MULT", "DIV", "NEQ", "EQUALS", "LESS_THAN", "GR_THAN", "LESS_THAN_E",
	"GR_THAN_E", "AND", "OR", "XOR", "NOR", "NAND", "NOT", "IF", "THEN", "ELSE",
	"IS", "DEFAULT", "WHILE", "UNTIL", "FOR", "IN", "RANGE", "FROM", "TO",
	"DO", "SKIP_ITERATION", "BREAK", "ACT_DELIM", "ARG_DELIM", "NEWLINE", "BLOCK_BEGIN",
	"BLOCK_END", "PRINTER", "KEY_INPUT", "G_N", "V_N", "E_N", "WHERE", "ID",
	"WS", "NUMBER", "FLOAT", "INT", "LINE_COMMENT", "M_LINE_COMMENT",
}

var ruleNames = []string{
	"file", "sequence", "sequence_element", "sequence_line", "one_line_sequence_element",
	"atom_action", "if_stmnt", "if_is_stmnt", "case_stmnt", "one_line_stmnt",
	"mult_line_stmnt", "stmnt", "while_stmnt", "until_stmnt", "for_stmnt",
	"from_to_stmnt", "function_declaration", "one_line_function_declaration",
	"mult_line_function_declaration", "function_declaration_head", "one_line_procedure_declaration",
	"mult_line_procedure_declaration", "procedure_declaration_head", "var_assign",
	"arc_assign", "arc", "or_w_arc_lr", "or_w_arc_rl", "unor_w_arc", "vertice_assign",
	"graph_assign", "labeled_assign", "expr", "integral_expr", "logical_expr",
	"log_expr_operand", "bin_log_operator", "unar_log_operator", "arithm_expr",
	"arithm_expr_operand", "bin_arithm_operator", "arithm_assign_operator",
	"builtin_function_call", "built_func_print", "built_func_input", "function_call",
	"label", "value", "variable", "builtin", "block_end",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GraffleParser struct {
	*antlr.BaseParser
}

func NewGraffleParser(input antlr.TokenStream) *GraffleParser {
	this := new(GraffleParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GraffleParser.g4"

	return this
}

// GraffleParser tokens.
const (
	GraffleParserEOF            = antlr.TokenEOF
	GraffleParserSTRING         = 1
	GraffleParserLABEL          = 2
	GraffleParserML_LABEL       = 3
	GraffleParserBOOL           = 4
	GraffleParserL_PAREN        = 5
	GraffleParserR_PAREN        = 6
	GraffleParserL_CURLY        = 7
	GraffleParserR_CURLY        = 8
	GraffleParserL_BRACKET      = 9
	GraffleParserR_BRACKET      = 10
	GraffleParserOR_ARC_LR      = 11
	GraffleParserOR_ARC_RL      = 12
	GraffleParserUNOR_ARC       = 13
	GraffleParserASSIGN         = 14
	GraffleParserADD_ASSIGN     = 15
	GraffleParserSUB_ASSIGN     = 16
	GraffleParserMULT_ASSIGN    = 17
	GraffleParserDIV_ASSIGN     = 18
	GraffleParserADD            = 19
	GraffleParserSUB            = 20
	GraffleParserMULT           = 21
	GraffleParserDIV            = 22
	GraffleParserNEQ            = 23
	GraffleParserEQUALS         = 24
	GraffleParserLESS_THAN      = 25
	GraffleParserGR_THAN        = 26
	GraffleParserLESS_THAN_E    = 27
	GraffleParserGR_THAN_E      = 28
	GraffleParserAND            = 29
	GraffleParserOR             = 30
	GraffleParserXOR            = 31
	GraffleParserNOR            = 32
	GraffleParserNAND           = 33
	GraffleParserNOT            = 34
	GraffleParserIF             = 35
	GraffleParserTHEN           = 36
	GraffleParserELSE           = 37
	GraffleParserIS             = 38
	GraffleParserDEFAULT        = 39
	GraffleParserWHILE          = 40
	GraffleParserUNTIL          = 41
	GraffleParserFOR            = 42
	GraffleParserIN             = 43
	GraffleParserRANGE          = 44
	GraffleParserFROM           = 45
	GraffleParserTO             = 46
	GraffleParserDO             = 47
	GraffleParserSKIP_ITERATION = 48
	GraffleParserBREAK          = 49
	GraffleParserACT_DELIM      = 50
	GraffleParserARG_DELIM      = 51
	GraffleParserNEWLINE        = 52
	GraffleParserBLOCK_BEGIN    = 53
	GraffleParserBLOCK_END      = 54
	GraffleParserPRINTER        = 55
	GraffleParserKEY_INPUT      = 56
	GraffleParserG_N            = 57
	GraffleParserV_N            = 58
	GraffleParserE_N            = 59
	GraffleParserWHERE          = 60
	GraffleParserID             = 61
	GraffleParserWS             = 62
	GraffleParserNUMBER         = 63
	GraffleParserFLOAT          = 64
	GraffleParserINT            = 65
	GraffleParserLINE_COMMENT   = 66
	GraffleParserM_LINE_COMMENT = 67
)

// GraffleParser rules.
const (
	GraffleParserRULE_file                            = 0
	GraffleParserRULE_sequence                        = 1
	GraffleParserRULE_sequence_element                = 2
	GraffleParserRULE_sequence_line                   = 3
	GraffleParserRULE_one_line_sequence_element       = 4
	GraffleParserRULE_atom_action                     = 5
	GraffleParserRULE_if_stmnt                        = 6
	GraffleParserRULE_if_is_stmnt                     = 7
	GraffleParserRULE_case_stmnt                      = 8
	GraffleParserRULE_one_line_stmnt                  = 9
	GraffleParserRULE_mult_line_stmnt                 = 10
	GraffleParserRULE_stmnt                           = 11
	GraffleParserRULE_while_stmnt                     = 12
	GraffleParserRULE_until_stmnt                     = 13
	GraffleParserRULE_for_stmnt                       = 14
	GraffleParserRULE_from_to_stmnt                   = 15
	GraffleParserRULE_function_declaration            = 16
	GraffleParserRULE_one_line_function_declaration   = 17
	GraffleParserRULE_mult_line_function_declaration  = 18
	GraffleParserRULE_function_declaration_head       = 19
	GraffleParserRULE_one_line_procedure_declaration  = 20
	GraffleParserRULE_mult_line_procedure_declaration = 21
	GraffleParserRULE_procedure_declaration_head      = 22
	GraffleParserRULE_var_assign                      = 23
	GraffleParserRULE_arc_assign                      = 24
	GraffleParserRULE_arc                             = 25
	GraffleParserRULE_or_w_arc_lr                     = 26
	GraffleParserRULE_or_w_arc_rl                     = 27
	GraffleParserRULE_unor_w_arc                      = 28
	GraffleParserRULE_vertice_assign                  = 29
	GraffleParserRULE_graph_assign                    = 30
	GraffleParserRULE_labeled_assign                  = 31
	GraffleParserRULE_expr                            = 32
	GraffleParserRULE_integral_expr                   = 33
	GraffleParserRULE_logical_expr                    = 34
	GraffleParserRULE_log_expr_operand                = 35
	GraffleParserRULE_bin_log_operator                = 36
	GraffleParserRULE_unar_log_operator               = 37
	GraffleParserRULE_arithm_expr                     = 38
	GraffleParserRULE_arithm_expr_operand             = 39
	GraffleParserRULE_bin_arithm_operator             = 40
	GraffleParserRULE_arithm_assign_operator          = 41
	GraffleParserRULE_builtin_function_call           = 42
	GraffleParserRULE_built_func_print                = 43
	GraffleParserRULE_built_func_input                = 44
	GraffleParserRULE_function_call                   = 45
	GraffleParserRULE_label                           = 46
	GraffleParserRULE_value                           = 47
	GraffleParserRULE_variable                        = 48
	GraffleParserRULE_builtin                         = 49
	GraffleParserRULE_block_end                       = 50
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMain_body returns the main_body rule contexts.
	GetMain_body() ISequenceContext

	// SetMain_body sets the main_body rule contexts.
	SetMain_body(ISequenceContext)

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	main_body ISequenceContext
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) GetMain_body() ISequenceContext { return s.main_body }

func (s *FileContext) SetMain_body(v ISequenceContext) { s.main_body = v }

func (s *FileContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *FileContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *FileContext) AllFunction_declaration() []IFunction_declarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunction_declarationContext)(nil)).Elem())
	var tst = make([]IFunction_declarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunction_declarationContext)
		}
	}

	return tst
}

func (s *FileContext) Function_declaration(i int) IFunction_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_declarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunction_declarationContext)
}

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(GraffleParserEOF, 0)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraffleParserRULE_file)

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
	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(102)
			p.Match(GraffleParserNEWLINE)
		}

	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(105)
				p.Function_declaration()
			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	{
		p.SetState(111)

		var _x = p.Sequence()

		localctx.(*FileContext).main_body = _x
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(112)
			p.Match(GraffleParserEOF)
		}

	}

	return localctx
}

// ISequenceContext is an interface to support dynamic dispatch.
type ISequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequenceContext differentiates from other interfaces.
	IsSequenceContext()
}

type SequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceContext() *SequenceContext {
	var p = new(SequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_sequence
	return p
}

func (*SequenceContext) IsSequenceContext() {}

func NewSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceContext {
	var p = new(SequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_sequence

	return p
}

func (s *SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserNEWLINE)
}

func (s *SequenceContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, i)
}

func (s *SequenceContext) AllSequence_element() []ISequence_elementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISequence_elementContext)(nil)).Elem())
	var tst = make([]ISequence_elementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISequence_elementContext)
		}
	}

	return tst
}

func (s *SequenceContext) Sequence_element(i int) ISequence_elementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_elementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISequence_elementContext)
}

func (s *SequenceContext) AllEOF() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserEOF)
}

func (s *SequenceContext) EOF(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserEOF, i)
}

func (s *SequenceContext) AllACT_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserACT_DELIM)
}

func (s *SequenceContext) ACT_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserACT_DELIM, i)
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitSequence(s)
	}
}

func (s *SequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Sequence() (localctx ISequenceContext) {
	localctx = NewSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraffleParserRULE_sequence)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserNEWLINE {
		{
			p.SetState(115)
			p.Match(GraffleParserNEWLINE)
		}

	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(118)
				p.Sequence_element()
			}
			p.SetState(125)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case GraffleParserACT_DELIM, GraffleParserNEWLINE:
				p.SetState(120)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(119)
							_la = p.GetTokenStream().LA(1)

							if !(_la == GraffleParserACT_DELIM || _la == GraffleParserNEWLINE) {
								p.GetErrorHandler().RecoverInline(p)
							} else {
								p.GetErrorHandler().ReportMatch(p)
								p.Consume()
							}
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(122)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
				}

			case GraffleParserEOF:
				{
					p.SetState(124)
					p.Match(GraffleParserEOF)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// ISequence_elementContext is an interface to support dynamic dispatch.
type ISequence_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequence_elementContext differentiates from other interfaces.
	IsSequence_elementContext()
}

type Sequence_elementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequence_elementContext() *Sequence_elementContext {
	var p = new(Sequence_elementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_sequence_element
	return p
}

func (*Sequence_elementContext) IsSequence_elementContext() {}

func NewSequence_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sequence_elementContext {
	var p = new(Sequence_elementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_sequence_element

	return p
}

func (s *Sequence_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Sequence_elementContext) If_stmnt() IIf_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_stmntContext)
}

func (s *Sequence_elementContext) If_is_stmnt() IIf_is_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_is_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_is_stmntContext)
}

func (s *Sequence_elementContext) Mult_line_stmnt() IMult_line_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMult_line_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMult_line_stmntContext)
}

func (s *Sequence_elementContext) Sequence_line() ISequence_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequence_lineContext)
}

func (s *Sequence_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sequence_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sequence_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterSequence_element(s)
	}
}

func (s *Sequence_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitSequence_element(s)
	}
}

func (s *Sequence_elementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitSequence_element(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Sequence_element() (localctx ISequence_elementContext) {
	localctx = NewSequence_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GraffleParserRULE_sequence_element)

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

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.If_stmnt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.If_is_stmnt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.Mult_line_stmnt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(134)
			p.Sequence_line()
		}

	}

	return localctx
}

// ISequence_lineContext is an interface to support dynamic dispatch.
type ISequence_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequence_lineContext differentiates from other interfaces.
	IsSequence_lineContext()
}

type Sequence_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequence_lineContext() *Sequence_lineContext {
	var p = new(Sequence_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_sequence_line
	return p
}

func (*Sequence_lineContext) IsSequence_lineContext() {}

func NewSequence_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sequence_lineContext {
	var p = new(Sequence_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_sequence_line

	return p
}

func (s *Sequence_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Sequence_lineContext) AllOne_line_sequence_element() []IOne_line_sequence_elementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOne_line_sequence_elementContext)(nil)).Elem())
	var tst = make([]IOne_line_sequence_elementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOne_line_sequence_elementContext)
		}
	}

	return tst
}

func (s *Sequence_lineContext) One_line_sequence_element(i int) IOne_line_sequence_elementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOne_line_sequence_elementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOne_line_sequence_elementContext)
}

func (s *Sequence_lineContext) AllACT_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserACT_DELIM)
}

func (s *Sequence_lineContext) ACT_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserACT_DELIM, i)
}

func (s *Sequence_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sequence_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sequence_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterSequence_line(s)
	}
}

func (s *Sequence_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitSequence_line(s)
	}
}

func (s *Sequence_lineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitSequence_line(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Sequence_line() (localctx ISequence_lineContext) {
	localctx = NewSequence_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GraffleParserRULE_sequence_line)

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
	{
		p.SetState(137)
		p.One_line_sequence_element()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(138)
				p.Match(GraffleParserACT_DELIM)
			}
			{
				p.SetState(139)
				p.One_line_sequence_element()
			}

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(145)
			p.Match(GraffleParserACT_DELIM)
		}

	}

	return localctx
}

// IOne_line_sequence_elementContext is an interface to support dynamic dispatch.
type IOne_line_sequence_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOne_line_sequence_elementContext differentiates from other interfaces.
	IsOne_line_sequence_elementContext()
}

type One_line_sequence_elementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOne_line_sequence_elementContext() *One_line_sequence_elementContext {
	var p = new(One_line_sequence_elementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_one_line_sequence_element
	return p
}

func (*One_line_sequence_elementContext) IsOne_line_sequence_elementContext() {}

func NewOne_line_sequence_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *One_line_sequence_elementContext {
	var p = new(One_line_sequence_elementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_one_line_sequence_element

	return p
}

func (s *One_line_sequence_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *One_line_sequence_elementContext) Atom_action() IAtom_actionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtom_actionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtom_actionContext)
}

func (s *One_line_sequence_elementContext) One_line_stmnt() IOne_line_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOne_line_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOne_line_stmntContext)
}

func (s *One_line_sequence_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *One_line_sequence_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *One_line_sequence_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOne_line_sequence_element(s)
	}
}

func (s *One_line_sequence_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOne_line_sequence_element(s)
	}
}

func (s *One_line_sequence_elementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOne_line_sequence_element(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) One_line_sequence_element() (localctx IOne_line_sequence_elementContext) {
	localctx = NewOne_line_sequence_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GraffleParserRULE_one_line_sequence_element)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserPRINTER, GraffleParserKEY_INPUT, GraffleParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Atom_action()
		}

	case GraffleParserWHILE, GraffleParserUNTIL, GraffleParserFOR, GraffleParserFROM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.One_line_stmnt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAtom_actionContext is an interface to support dynamic dispatch.
type IAtom_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtom_actionContext differentiates from other interfaces.
	IsAtom_actionContext()
}

type Atom_actionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtom_actionContext() *Atom_actionContext {
	var p = new(Atom_actionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_atom_action
	return p
}

func (*Atom_actionContext) IsAtom_actionContext() {}

func NewAtom_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Atom_actionContext {
	var p = new(Atom_actionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_atom_action

	return p
}

func (s *Atom_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Atom_actionContext) Var_assign() IVar_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_assignContext)
}

func (s *Atom_actionContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Atom_actionContext) Builtin_function_call() IBuiltin_function_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltin_function_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltin_function_callContext)
}

func (s *Atom_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atom_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Atom_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterAtom_action(s)
	}
}

func (s *Atom_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitAtom_action(s)
	}
}

func (s *Atom_actionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitAtom_action(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Atom_action() (localctx IAtom_actionContext) {
	localctx = NewAtom_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GraffleParserRULE_atom_action)

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

	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Var_assign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Function_call()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.Builtin_function_call()
		}

	}

	return localctx
}

// IIf_stmntContext is an interface to support dynamic dispatch.
type IIf_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() ILogical_exprContext

	// SetCond sets the cond rule contexts.
	SetCond(ILogical_exprContext)

	// IsIf_stmntContext differentiates from other interfaces.
	IsIf_stmntContext()
}

type If_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   ILogical_exprContext
}

func NewEmptyIf_stmntContext() *If_stmntContext {
	var p = new(If_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_if_stmnt
	return p
}

func (*If_stmntContext) IsIf_stmntContext() {}

func NewIf_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmntContext {
	var p = new(If_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_if_stmnt

	return p
}

func (s *If_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmntContext) GetCond() ILogical_exprContext { return s.cond }

func (s *If_stmntContext) SetCond(v ILogical_exprContext) { s.cond = v }

func (s *If_stmntContext) IF() antlr.TerminalNode {
	return s.GetToken(GraffleParserIF, 0)
}

func (s *If_stmntContext) Block_end() IBlock_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endContext)
}

func (s *If_stmntContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *If_stmntContext) AllSequence() []ISequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISequenceContext)(nil)).Elem())
	var tst = make([]ISequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISequenceContext)
		}
	}

	return tst
}

func (s *If_stmntContext) Sequence(i int) ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *If_stmntContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *If_stmntContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GraffleParserELSE, 0)
}

func (s *If_stmntContext) THEN() antlr.TerminalNode {
	return s.GetToken(GraffleParserTHEN, 0)
}

func (s *If_stmntContext) AllDO() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserDO)
}

func (s *If_stmntContext) DO(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, i)
}

func (s *If_stmntContext) ARG_DELIM() antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, 0)
}

func (s *If_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterIf_stmnt(s)
	}
}

func (s *If_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitIf_stmnt(s)
	}
}

func (s *If_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitIf_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) If_stmnt() (localctx IIf_stmntContext) {
	localctx = NewIf_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GraffleParserRULE_if_stmnt)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(GraffleParserIF)
	}
	{
		p.SetState(158)

		var _x = p.Logical_expr()

		localctx.(*If_stmntContext).cond = _x
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(GraffleParserTHEN-36))|(1<<(GraffleParserDO-36))|(1<<(GraffleParserARG_DELIM-36)))) != 0 {
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserARG_DELIM {
			{
				p.SetState(159)
				p.Match(GraffleParserARG_DELIM)
			}

		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(162)
				p.Match(GraffleParserTHEN)
			}

		case 2:
			{
				p.SetState(163)
				p.Match(GraffleParserDO)
			}

		case 3:
			{
				p.SetState(164)
				p.Match(GraffleParserTHEN)
			}
			{
				p.SetState(165)
				p.Match(GraffleParserDO)
			}

		}

	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(170)
				p.Sequence()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(175)
			p.Match(GraffleParserNEWLINE)
		}

	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserELSE {
		{
			p.SetState(178)
			p.Match(GraffleParserELSE)
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(179)
				p.Match(GraffleParserDO)
			}

		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(182)
					p.Sequence()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(185)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(189)
		p.Block_end()
	}

	return localctx
}

// IIf_is_stmntContext is an interface to support dynamic dispatch.
type IIf_is_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_is_stmntContext differentiates from other interfaces.
	IsIf_is_stmntContext()
}

type If_is_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_is_stmntContext() *If_is_stmntContext {
	var p = new(If_is_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_if_is_stmnt
	return p
}

func (*If_is_stmntContext) IsIf_is_stmntContext() {}

func NewIf_is_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_is_stmntContext {
	var p = new(If_is_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_if_is_stmnt

	return p
}

func (s *If_is_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *If_is_stmntContext) IF() antlr.TerminalNode {
	return s.GetToken(GraffleParserIF, 0)
}

func (s *If_is_stmntContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *If_is_stmntContext) Block_end() IBlock_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endContext)
}

func (s *If_is_stmntContext) AllCase_stmnt() []ICase_stmntContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICase_stmntContext)(nil)).Elem())
	var tst = make([]ICase_stmntContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICase_stmntContext)
		}
	}

	return tst
}

func (s *If_is_stmntContext) Case_stmnt(i int) ICase_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICase_stmntContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICase_stmntContext)
}

func (s *If_is_stmntContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *If_is_stmntContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(GraffleParserDEFAULT, 0)
}

func (s *If_is_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *If_is_stmntContext) AllSequence() []ISequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISequenceContext)(nil)).Elem())
	var tst = make([]ISequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISequenceContext)
		}
	}

	return tst
}

func (s *If_is_stmntContext) Sequence(i int) ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *If_is_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_is_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_is_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterIf_is_stmnt(s)
	}
}

func (s *If_is_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitIf_is_stmnt(s)
	}
}

func (s *If_is_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitIf_is_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) If_is_stmnt() (localctx IIf_is_stmntContext) {
	localctx = NewIf_is_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GraffleParserRULE_if_is_stmnt)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(GraffleParserIF)
	}
	{
		p.SetState(192)
		p.Value()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(193)
				p.Case_stmnt()
			}

		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(199)
			p.Match(GraffleParserNEWLINE)
		}

	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDEFAULT {
		{
			p.SetState(202)
			p.Match(GraffleParserDEFAULT)
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(203)
				p.Match(GraffleParserDO)
			}

		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(206)
					p.Sequence()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(213)
		p.Block_end()
	}

	return localctx
}

// ICase_stmntContext is an interface to support dynamic dispatch.
type ICase_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCase_stmntContext differentiates from other interfaces.
	IsCase_stmntContext()
}

type Case_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCase_stmntContext() *Case_stmntContext {
	var p = new(Case_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_case_stmnt
	return p
}

func (*Case_stmntContext) IsCase_stmntContext() {}

func NewCase_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_stmntContext {
	var p = new(Case_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_case_stmnt

	return p
}

func (s *Case_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_stmntContext) IS() antlr.TerminalNode {
	return s.GetToken(GraffleParserIS, 0)
}

func (s *Case_stmntContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Case_stmntContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *Case_stmntContext) AllSequence() []ISequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISequenceContext)(nil)).Elem())
	var tst = make([]ISequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISequenceContext)
		}
	}

	return tst
}

func (s *Case_stmntContext) Sequence(i int) ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Case_stmntContext) THEN() antlr.TerminalNode {
	return s.GetToken(GraffleParserTHEN, 0)
}

func (s *Case_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *Case_stmntContext) ARG_DELIM() antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, 0)
}

func (s *Case_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterCase_stmnt(s)
	}
}

func (s *Case_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitCase_stmnt(s)
	}
}

func (s *Case_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitCase_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Case_stmnt() (localctx ICase_stmntContext) {
	localctx = NewCase_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GraffleParserRULE_case_stmnt)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserNEWLINE {
		{
			p.SetState(215)
			p.Match(GraffleParserNEWLINE)
		}

	}
	{
		p.SetState(218)
		p.Match(GraffleParserIS)
	}
	{
		p.SetState(219)
		p.Value()
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(GraffleParserTHEN-36))|(1<<(GraffleParserDO-36))|(1<<(GraffleParserARG_DELIM-36)))) != 0 {
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserARG_DELIM {
			{
				p.SetState(220)
				p.Match(GraffleParserARG_DELIM)
			}

		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(223)
				p.Match(GraffleParserTHEN)
			}

		case 2:
			{
				p.SetState(224)
				p.Match(GraffleParserDO)
			}

		case 3:
			{
				p.SetState(225)
				p.Match(GraffleParserTHEN)
			}
			{
				p.SetState(226)
				p.Match(GraffleParserDO)
			}

		}

	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(231)
				p.Sequence()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}

	return localctx
}

// IOne_line_stmntContext is an interface to support dynamic dispatch.
type IOne_line_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOne_line_stmntContext differentiates from other interfaces.
	IsOne_line_stmntContext()
}

type One_line_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOne_line_stmntContext() *One_line_stmntContext {
	var p = new(One_line_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_one_line_stmnt
	return p
}

func (*One_line_stmntContext) IsOne_line_stmntContext() {}

func NewOne_line_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *One_line_stmntContext {
	var p = new(One_line_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_one_line_stmnt

	return p
}

func (s *One_line_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *One_line_stmntContext) Stmnt() IStmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmntContext)
}

func (s *One_line_stmntContext) Sequence_line() ISequence_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequence_lineContext)
}

func (s *One_line_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *One_line_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *One_line_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOne_line_stmnt(s)
	}
}

func (s *One_line_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOne_line_stmnt(s)
	}
}

func (s *One_line_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOne_line_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) One_line_stmnt() (localctx IOne_line_stmntContext) {
	localctx = NewOne_line_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GraffleParserRULE_one_line_stmnt)

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
		p.SetState(236)
		p.Stmnt()
	}
	{
		p.SetState(237)
		p.Sequence_line()
	}

	return localctx
}

// IMult_line_stmntContext is an interface to support dynamic dispatch.
type IMult_line_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMult_line_stmntContext differentiates from other interfaces.
	IsMult_line_stmntContext()
}

type Mult_line_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMult_line_stmntContext() *Mult_line_stmntContext {
	var p = new(Mult_line_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_mult_line_stmnt
	return p
}

func (*Mult_line_stmntContext) IsMult_line_stmntContext() {}

func NewMult_line_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mult_line_stmntContext {
	var p = new(Mult_line_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_mult_line_stmnt

	return p
}

func (s *Mult_line_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *Mult_line_stmntContext) Stmnt() IStmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmntContext)
}

func (s *Mult_line_stmntContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Mult_line_stmntContext) Block_end() IBlock_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endContext)
}

func (s *Mult_line_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mult_line_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mult_line_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterMult_line_stmnt(s)
	}
}

func (s *Mult_line_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitMult_line_stmnt(s)
	}
}

func (s *Mult_line_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitMult_line_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Mult_line_stmnt() (localctx IMult_line_stmntContext) {
	localctx = NewMult_line_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GraffleParserRULE_mult_line_stmnt)

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
		p.SetState(239)
		p.Stmnt()
	}
	{
		p.SetState(240)
		p.Sequence()
	}
	{
		p.SetState(241)
		p.Block_end()
	}

	return localctx
}

// IStmntContext is an interface to support dynamic dispatch.
type IStmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmntContext differentiates from other interfaces.
	IsStmntContext()
}

type StmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmntContext() *StmntContext {
	var p = new(StmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_stmnt
	return p
}

func (*StmntContext) IsStmntContext() {}

func NewStmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmntContext {
	var p = new(StmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_stmnt

	return p
}

func (s *StmntContext) GetParser() antlr.Parser { return s.parser }

func (s *StmntContext) While_stmnt() IWhile_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_stmntContext)
}

func (s *StmntContext) Until_stmnt() IUntil_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUntil_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUntil_stmntContext)
}

func (s *StmntContext) For_stmnt() IFor_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_stmntContext)
}

func (s *StmntContext) From_to_stmnt() IFrom_to_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFrom_to_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFrom_to_stmntContext)
}

func (s *StmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterStmnt(s)
	}
}

func (s *StmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitStmnt(s)
	}
}

func (s *StmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitStmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Stmnt() (localctx IStmntContext) {
	localctx = NewStmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GraffleParserRULE_stmnt)

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

	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserWHILE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)
			p.While_stmnt()
		}

	case GraffleParserUNTIL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
			p.Until_stmnt()
		}

	case GraffleParserFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(245)
			p.For_stmnt()
		}

	case GraffleParserFROM:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(246)
			p.From_to_stmnt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWhile_stmntContext is an interface to support dynamic dispatch.
type IWhile_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() ILogical_exprContext

	// SetCond sets the cond rule contexts.
	SetCond(ILogical_exprContext)

	// IsWhile_stmntContext differentiates from other interfaces.
	IsWhile_stmntContext()
}

type While_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   ILogical_exprContext
}

func NewEmptyWhile_stmntContext() *While_stmntContext {
	var p = new(While_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_while_stmnt
	return p
}

func (*While_stmntContext) IsWhile_stmntContext() {}

func NewWhile_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmntContext {
	var p = new(While_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_while_stmnt

	return p
}

func (s *While_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmntContext) GetCond() ILogical_exprContext { return s.cond }

func (s *While_stmntContext) SetCond(v ILogical_exprContext) { s.cond = v }

func (s *While_stmntContext) WHILE() antlr.TerminalNode {
	return s.GetToken(GraffleParserWHILE, 0)
}

func (s *While_stmntContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *While_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *While_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterWhile_stmnt(s)
	}
}

func (s *While_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitWhile_stmnt(s)
	}
}

func (s *While_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitWhile_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) While_stmnt() (localctx IWhile_stmntContext) {
	localctx = NewWhile_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GraffleParserRULE_while_stmnt)
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
		p.SetState(249)
		p.Match(GraffleParserWHILE)
	}
	{
		p.SetState(250)

		var _x = p.Logical_expr()

		localctx.(*While_stmntContext).cond = _x
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDO {
		{
			p.SetState(251)
			p.Match(GraffleParserDO)
		}

	}

	return localctx
}

// IUntil_stmntContext is an interface to support dynamic dispatch.
type IUntil_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() ILogical_exprContext

	// SetCond sets the cond rule contexts.
	SetCond(ILogical_exprContext)

	// IsUntil_stmntContext differentiates from other interfaces.
	IsUntil_stmntContext()
}

type Until_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   ILogical_exprContext
}

func NewEmptyUntil_stmntContext() *Until_stmntContext {
	var p = new(Until_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_until_stmnt
	return p
}

func (*Until_stmntContext) IsUntil_stmntContext() {}

func NewUntil_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Until_stmntContext {
	var p = new(Until_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_until_stmnt

	return p
}

func (s *Until_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *Until_stmntContext) GetCond() ILogical_exprContext { return s.cond }

func (s *Until_stmntContext) SetCond(v ILogical_exprContext) { s.cond = v }

func (s *Until_stmntContext) UNTIL() antlr.TerminalNode {
	return s.GetToken(GraffleParserUNTIL, 0)
}

func (s *Until_stmntContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *Until_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *Until_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Until_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Until_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterUntil_stmnt(s)
	}
}

func (s *Until_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitUntil_stmnt(s)
	}
}

func (s *Until_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitUntil_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Until_stmnt() (localctx IUntil_stmntContext) {
	localctx = NewUntil_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GraffleParserRULE_until_stmnt)
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
		p.SetState(254)
		p.Match(GraffleParserUNTIL)
	}
	{
		p.SetState(255)

		var _x = p.Logical_expr()

		localctx.(*Until_stmntContext).cond = _x
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDO {
		{
			p.SetState(256)
			p.Match(GraffleParserDO)
		}

	}

	return localctx
}

// IFor_stmntContext is an interface to support dynamic dispatch.
type IFor_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_stmntContext differentiates from other interfaces.
	IsFor_stmntContext()
}

type For_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmntContext() *For_stmntContext {
	var p = new(For_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_for_stmnt
	return p
}

func (*For_stmntContext) IsFor_stmntContext() {}

func NewFor_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmntContext {
	var p = new(For_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_for_stmnt

	return p
}

func (s *For_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmntContext) CopyFrom(ctx *For_stmntContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *For_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForLogicalContext struct {
	*For_stmntContext
	cond ILogical_exprContext
}

func NewForLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForLogicalContext {
	var p = new(ForLogicalContext)

	p.For_stmntContext = NewEmptyFor_stmntContext()
	p.parser = parser
	p.CopyFrom(ctx.(*For_stmntContext))

	return p
}

func (s *ForLogicalContext) GetCond() ILogical_exprContext { return s.cond }

func (s *ForLogicalContext) SetCond(v ILogical_exprContext) { s.cond = v }

func (s *ForLogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLogicalContext) FOR() antlr.TerminalNode {
	return s.GetToken(GraffleParserFOR, 0)
}

func (s *ForLogicalContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *ForLogicalContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *ForLogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterForLogical(s)
	}
}

func (s *ForLogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitForLogical(s)
	}
}

func (s *ForLogicalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitForLogical(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForVarContext struct {
	*For_stmntContext
	pre_act  IAtom_actionContext
	cond     ILogical_exprContext
	post_act IAtom_actionContext
}

func NewForVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForVarContext {
	var p = new(ForVarContext)

	p.For_stmntContext = NewEmptyFor_stmntContext()
	p.parser = parser
	p.CopyFrom(ctx.(*For_stmntContext))

	return p
}

func (s *ForVarContext) GetPre_act() IAtom_actionContext { return s.pre_act }

func (s *ForVarContext) GetCond() ILogical_exprContext { return s.cond }

func (s *ForVarContext) GetPost_act() IAtom_actionContext { return s.post_act }

func (s *ForVarContext) SetPre_act(v IAtom_actionContext) { s.pre_act = v }

func (s *ForVarContext) SetCond(v ILogical_exprContext) { s.cond = v }

func (s *ForVarContext) SetPost_act(v IAtom_actionContext) { s.post_act = v }

func (s *ForVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForVarContext) FOR() antlr.TerminalNode {
	return s.GetToken(GraffleParserFOR, 0)
}

func (s *ForVarContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *ForVarContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
}

func (s *ForVarContext) AllAtom_action() []IAtom_actionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtom_actionContext)(nil)).Elem())
	var tst = make([]IAtom_actionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtom_actionContext)
		}
	}

	return tst
}

func (s *ForVarContext) Atom_action(i int) IAtom_actionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtom_actionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtom_actionContext)
}

func (s *ForVarContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *ForVarContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *ForVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterForVar(s)
	}
}

func (s *ForVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitForVar(s)
	}
}

func (s *ForVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitForVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForRangeContext struct {
	*For_stmntContext
	from IExprContext
	to   IExprContext
}

func NewForRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeContext {
	var p = new(ForRangeContext)

	p.For_stmntContext = NewEmptyFor_stmntContext()
	p.parser = parser
	p.CopyFrom(ctx.(*For_stmntContext))

	return p
}

func (s *ForRangeContext) GetFrom() IExprContext { return s.from }

func (s *ForRangeContext) GetTo() IExprContext { return s.to }

func (s *ForRangeContext) SetFrom(v IExprContext) { s.from = v }

func (s *ForRangeContext) SetTo(v IExprContext) { s.to = v }

func (s *ForRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeContext) FOR() antlr.TerminalNode {
	return s.GetToken(GraffleParserFOR, 0)
}

func (s *ForRangeContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ForRangeContext) IN() antlr.TerminalNode {
	return s.GetToken(GraffleParserIN, 0)
}

func (s *ForRangeContext) TO() antlr.TerminalNode {
	return s.GetToken(GraffleParserTO, 0)
}

func (s *ForRangeContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ForRangeContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForRangeContext) RANGE() antlr.TerminalNode {
	return s.GetToken(GraffleParserRANGE, 0)
}

func (s *ForRangeContext) FROM() antlr.TerminalNode {
	return s.GetToken(GraffleParserFROM, 0)
}

func (s *ForRangeContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *ForRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterForRange(s)
	}
}

func (s *ForRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitForRange(s)
	}
}

func (s *ForRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitForRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) For_stmnt() (localctx IFor_stmntContext) {
	localctx = NewFor_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GraffleParserRULE_for_stmnt)
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

	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForLogicalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Match(GraffleParserFOR)
		}
		{
			p.SetState(260)

			var _x = p.Logical_expr()

			localctx.(*ForLogicalContext).cond = _x
		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(261)
				p.Match(GraffleParserDO)
			}

		}

	case 2:
		localctx = NewForVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.Match(GraffleParserFOR)
		}
		{
			p.SetState(265)

			var _x = p.Atom_action()

			localctx.(*ForVarContext).pre_act = _x
		}
		{
			p.SetState(266)
			p.Match(GraffleParserARG_DELIM)
		}
		{
			p.SetState(267)

			var _x = p.Logical_expr()

			localctx.(*ForVarContext).cond = _x
		}
		{
			p.SetState(268)
			p.Match(GraffleParserARG_DELIM)
		}
		{
			p.SetState(269)

			var _x = p.Atom_action()

			localctx.(*ForVarContext).post_act = _x
		}
		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(270)
				p.Match(GraffleParserDO)
			}

		}

	case 3:
		localctx = NewForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(273)
			p.Match(GraffleParserFOR)
		}
		{
			p.SetState(274)
			p.Variable()
		}
		{
			p.SetState(275)
			p.Match(GraffleParserIN)
		}
		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserRANGE {
			{
				p.SetState(276)
				p.Match(GraffleParserRANGE)
			}

		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserFROM {
			{
				p.SetState(279)
				p.Match(GraffleParserFROM)
			}

		}
		{
			p.SetState(282)

			var _x = p.Expr()

			localctx.(*ForRangeContext).from = _x
		}
		{
			p.SetState(283)
			p.Match(GraffleParserTO)
		}
		{
			p.SetState(284)

			var _x = p.Expr()

			localctx.(*ForRangeContext).to = _x
		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(285)
				p.Match(GraffleParserDO)
			}

		}

	}

	return localctx
}

// IFrom_to_stmntContext is an interface to support dynamic dispatch.
type IFrom_to_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFrom returns the from rule contexts.
	GetFrom() IExprContext

	// GetTo returns the to rule contexts.
	GetTo() IExprContext

	// SetFrom sets the from rule contexts.
	SetFrom(IExprContext)

	// SetTo sets the to rule contexts.
	SetTo(IExprContext)

	// IsFrom_to_stmntContext differentiates from other interfaces.
	IsFrom_to_stmntContext()
}

type From_to_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	from   IExprContext
	to     IExprContext
}

func NewEmptyFrom_to_stmntContext() *From_to_stmntContext {
	var p = new(From_to_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_from_to_stmnt
	return p
}

func (*From_to_stmntContext) IsFrom_to_stmntContext() {}

func NewFrom_to_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_to_stmntContext {
	var p = new(From_to_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_from_to_stmnt

	return p
}

func (s *From_to_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *From_to_stmntContext) GetFrom() IExprContext { return s.from }

func (s *From_to_stmntContext) GetTo() IExprContext { return s.to }

func (s *From_to_stmntContext) SetFrom(v IExprContext) { s.from = v }

func (s *From_to_stmntContext) SetTo(v IExprContext) { s.to = v }

func (s *From_to_stmntContext) FROM() antlr.TerminalNode {
	return s.GetToken(GraffleParserFROM, 0)
}

func (s *From_to_stmntContext) TO() antlr.TerminalNode {
	return s.GetToken(GraffleParserTO, 0)
}

func (s *From_to_stmntContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *From_to_stmntContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *From_to_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *From_to_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_to_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_to_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterFrom_to_stmnt(s)
	}
}

func (s *From_to_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitFrom_to_stmnt(s)
	}
}

func (s *From_to_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitFrom_to_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) From_to_stmnt() (localctx IFrom_to_stmntContext) {
	localctx = NewFrom_to_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GraffleParserRULE_from_to_stmnt)
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
		p.SetState(290)
		p.Match(GraffleParserFROM)
	}
	{
		p.SetState(291)

		var _x = p.Expr()

		localctx.(*From_to_stmntContext).from = _x
	}
	{
		p.SetState(292)
		p.Match(GraffleParserTO)
	}
	{
		p.SetState(293)

		var _x = p.Expr()

		localctx.(*From_to_stmntContext).to = _x
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDO {
		{
			p.SetState(294)
			p.Match(GraffleParserDO)
		}

	}

	return localctx
}

// IFunction_declarationContext is an interface to support dynamic dispatch.
type IFunction_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_declarationContext differentiates from other interfaces.
	IsFunction_declarationContext()
}

type Function_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_declarationContext() *Function_declarationContext {
	var p = new(Function_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_function_declaration
	return p
}

func (*Function_declarationContext) IsFunction_declarationContext() {}

func NewFunction_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declarationContext {
	var p = new(Function_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_function_declaration

	return p
}

func (s *Function_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declarationContext) One_line_function_declaration() IOne_line_function_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOne_line_function_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOne_line_function_declarationContext)
}

func (s *Function_declarationContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *Function_declarationContext) Mult_line_function_declaration() IMult_line_function_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMult_line_function_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMult_line_function_declarationContext)
}

func (s *Function_declarationContext) One_line_procedure_declaration() IOne_line_procedure_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOne_line_procedure_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOne_line_procedure_declarationContext)
}

func (s *Function_declarationContext) Mult_line_procedure_declaration() IMult_line_procedure_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMult_line_procedure_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMult_line_procedure_declarationContext)
}

func (s *Function_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterFunction_declaration(s)
	}
}

func (s *Function_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitFunction_declaration(s)
	}
}

func (s *Function_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitFunction_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Function_declaration() (localctx IFunction_declarationContext) {
	localctx = NewFunction_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GraffleParserRULE_function_declaration)
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

	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserNEWLINE {
			{
				p.SetState(297)
				p.Match(GraffleParserNEWLINE)
			}

		}
		{
			p.SetState(300)
			p.One_line_function_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserNEWLINE {
			{
				p.SetState(301)
				p.Match(GraffleParserNEWLINE)
			}

		}
		{
			p.SetState(304)
			p.Mult_line_function_declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserNEWLINE {
			{
				p.SetState(305)
				p.Match(GraffleParserNEWLINE)
			}

		}
		{
			p.SetState(308)
			p.One_line_procedure_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserNEWLINE {
			{
				p.SetState(309)
				p.Match(GraffleParserNEWLINE)
			}

		}
		{
			p.SetState(312)
			p.Mult_line_procedure_declaration()
		}

	}

	return localctx
}

// IOne_line_function_declarationContext is an interface to support dynamic dispatch.
type IOne_line_function_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOne_line_function_declarationContext differentiates from other interfaces.
	IsOne_line_function_declarationContext()
}

type One_line_function_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOne_line_function_declarationContext() *One_line_function_declarationContext {
	var p = new(One_line_function_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_one_line_function_declaration
	return p
}

func (*One_line_function_declarationContext) IsOne_line_function_declarationContext() {}

func NewOne_line_function_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *One_line_function_declarationContext {
	var p = new(One_line_function_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_one_line_function_declaration

	return p
}

func (s *One_line_function_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *One_line_function_declarationContext) Function_declaration_head() IFunction_declaration_headContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_declaration_headContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_declaration_headContext)
}

func (s *One_line_function_declarationContext) Sequence_line() ISequence_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequence_lineContext)
}

func (s *One_line_function_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *One_line_function_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *One_line_function_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOne_line_function_declaration(s)
	}
}

func (s *One_line_function_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOne_line_function_declaration(s)
	}
}

func (s *One_line_function_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOne_line_function_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) One_line_function_declaration() (localctx IOne_line_function_declarationContext) {
	localctx = NewOne_line_function_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GraffleParserRULE_one_line_function_declaration)

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
		p.SetState(315)
		p.Function_declaration_head()
	}
	{
		p.SetState(316)
		p.Sequence_line()
	}

	return localctx
}

// IMult_line_function_declarationContext is an interface to support dynamic dispatch.
type IMult_line_function_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMult_line_function_declarationContext differentiates from other interfaces.
	IsMult_line_function_declarationContext()
}

type Mult_line_function_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMult_line_function_declarationContext() *Mult_line_function_declarationContext {
	var p = new(Mult_line_function_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_mult_line_function_declaration
	return p
}

func (*Mult_line_function_declarationContext) IsMult_line_function_declarationContext() {}

func NewMult_line_function_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mult_line_function_declarationContext {
	var p = new(Mult_line_function_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_mult_line_function_declaration

	return p
}

func (s *Mult_line_function_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Mult_line_function_declarationContext) Function_declaration_head() IFunction_declaration_headContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_declaration_headContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_declaration_headContext)
}

func (s *Mult_line_function_declarationContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Mult_line_function_declarationContext) Block_end() IBlock_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endContext)
}

func (s *Mult_line_function_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mult_line_function_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mult_line_function_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterMult_line_function_declaration(s)
	}
}

func (s *Mult_line_function_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitMult_line_function_declaration(s)
	}
}

func (s *Mult_line_function_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitMult_line_function_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Mult_line_function_declaration() (localctx IMult_line_function_declarationContext) {
	localctx = NewMult_line_function_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GraffleParserRULE_mult_line_function_declaration)

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
		p.SetState(318)
		p.Function_declaration_head()
	}
	{
		p.SetState(319)
		p.Sequence()
	}
	{
		p.SetState(320)
		p.Block_end()
	}

	return localctx
}

// IFunction_declaration_headContext is an interface to support dynamic dispatch.
type IFunction_declaration_headContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpd1 returns the opd1 rule contexts.
	GetOpd1() IVariableContext

	// GetOpd2 returns the opd2 rule contexts.
	GetOpd2() IVariableContext

	// GetReturn_val returns the return_val rule contexts.
	GetReturn_val() IValueContext

	// SetOpd1 sets the opd1 rule contexts.
	SetOpd1(IVariableContext)

	// SetOpd2 sets the opd2 rule contexts.
	SetOpd2(IVariableContext)

	// SetReturn_val sets the return_val rule contexts.
	SetReturn_val(IValueContext)

	// IsFunction_declaration_headContext differentiates from other interfaces.
	IsFunction_declaration_headContext()
}

type Function_declaration_headContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	opd1       IVariableContext
	opd2       IVariableContext
	return_val IValueContext
}

func NewEmptyFunction_declaration_headContext() *Function_declaration_headContext {
	var p = new(Function_declaration_headContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_function_declaration_head
	return p
}

func (*Function_declaration_headContext) IsFunction_declaration_headContext() {}

func NewFunction_declaration_headContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declaration_headContext {
	var p = new(Function_declaration_headContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_function_declaration_head

	return p
}

func (s *Function_declaration_headContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declaration_headContext) GetOpd1() IVariableContext { return s.opd1 }

func (s *Function_declaration_headContext) GetOpd2() IVariableContext { return s.opd2 }

func (s *Function_declaration_headContext) GetReturn_val() IValueContext { return s.return_val }

func (s *Function_declaration_headContext) SetOpd1(v IVariableContext) { s.opd1 = v }

func (s *Function_declaration_headContext) SetOpd2(v IVariableContext) { s.opd2 = v }

func (s *Function_declaration_headContext) SetReturn_val(v IValueContext) { s.return_val = v }

func (s *Function_declaration_headContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Function_declaration_headContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Function_declaration_headContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Function_declaration_headContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Function_declaration_headContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Function_declaration_headContext) WHERE() antlr.TerminalNode {
	return s.GetToken(GraffleParserWHERE, 0)
}

func (s *Function_declaration_headContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *Function_declaration_headContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Function_declaration_headContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *Function_declaration_headContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
}

func (s *Function_declaration_headContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declaration_headContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declaration_headContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterFunction_declaration_head(s)
	}
}

func (s *Function_declaration_headContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitFunction_declaration_head(s)
	}
}

func (s *Function_declaration_headContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitFunction_declaration_head(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Function_declaration_head() (localctx IFunction_declaration_headContext) {
	localctx = NewFunction_declaration_headContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GraffleParserRULE_function_declaration_head)
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
		p.SetState(322)
		p.Match(GraffleParserID)
	}
	{
		p.SetState(323)
		p.Match(GraffleParserL_PAREN)
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserSUB || _la == GraffleParserID {
		{
			p.SetState(324)

			var _x = p.Variable()

			localctx.(*Function_declaration_headContext).opd1 = _x
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraffleParserARG_DELIM {
			{
				p.SetState(325)
				p.Match(GraffleParserARG_DELIM)
			}
			{
				p.SetState(326)

				var _x = p.Variable()

				localctx.(*Function_declaration_headContext).opd2 = _x
			}

			p.SetState(331)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(334)
		p.Match(GraffleParserR_PAREN)
	}
	{
		p.SetState(335)
		p.Match(GraffleParserASSIGN)
	}
	{
		p.SetState(336)

		var _x = p.Value()

		localctx.(*Function_declaration_headContext).return_val = _x
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserARG_DELIM || _la == GraffleParserWHERE {
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserARG_DELIM {
			{
				p.SetState(337)
				p.Match(GraffleParserARG_DELIM)
			}

		}
		{
			p.SetState(340)
			p.Match(GraffleParserWHERE)
		}

	}

	return localctx
}

// IOne_line_procedure_declarationContext is an interface to support dynamic dispatch.
type IOne_line_procedure_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOne_line_procedure_declarationContext differentiates from other interfaces.
	IsOne_line_procedure_declarationContext()
}

type One_line_procedure_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOne_line_procedure_declarationContext() *One_line_procedure_declarationContext {
	var p = new(One_line_procedure_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_one_line_procedure_declaration
	return p
}

func (*One_line_procedure_declarationContext) IsOne_line_procedure_declarationContext() {}

func NewOne_line_procedure_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *One_line_procedure_declarationContext {
	var p = new(One_line_procedure_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_one_line_procedure_declaration

	return p
}

func (s *One_line_procedure_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *One_line_procedure_declarationContext) Procedure_declaration_head() IProcedure_declaration_headContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedure_declaration_headContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedure_declaration_headContext)
}

func (s *One_line_procedure_declarationContext) Sequence_line() ISequence_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequence_lineContext)
}

func (s *One_line_procedure_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *One_line_procedure_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *One_line_procedure_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOne_line_procedure_declaration(s)
	}
}

func (s *One_line_procedure_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOne_line_procedure_declaration(s)
	}
}

func (s *One_line_procedure_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOne_line_procedure_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) One_line_procedure_declaration() (localctx IOne_line_procedure_declarationContext) {
	localctx = NewOne_line_procedure_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GraffleParserRULE_one_line_procedure_declaration)

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
		p.SetState(343)
		p.Procedure_declaration_head()
	}
	{
		p.SetState(344)
		p.Sequence_line()
	}

	return localctx
}

// IMult_line_procedure_declarationContext is an interface to support dynamic dispatch.
type IMult_line_procedure_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMult_line_procedure_declarationContext differentiates from other interfaces.
	IsMult_line_procedure_declarationContext()
}

type Mult_line_procedure_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMult_line_procedure_declarationContext() *Mult_line_procedure_declarationContext {
	var p = new(Mult_line_procedure_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_mult_line_procedure_declaration
	return p
}

func (*Mult_line_procedure_declarationContext) IsMult_line_procedure_declarationContext() {}

func NewMult_line_procedure_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mult_line_procedure_declarationContext {
	var p = new(Mult_line_procedure_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_mult_line_procedure_declaration

	return p
}

func (s *Mult_line_procedure_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Mult_line_procedure_declarationContext) Procedure_declaration_head() IProcedure_declaration_headContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedure_declaration_headContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedure_declaration_headContext)
}

func (s *Mult_line_procedure_declarationContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *Mult_line_procedure_declarationContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Mult_line_procedure_declarationContext) Block_end() IBlock_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endContext)
}

func (s *Mult_line_procedure_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mult_line_procedure_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mult_line_procedure_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterMult_line_procedure_declaration(s)
	}
}

func (s *Mult_line_procedure_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitMult_line_procedure_declaration(s)
	}
}

func (s *Mult_line_procedure_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitMult_line_procedure_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Mult_line_procedure_declaration() (localctx IMult_line_procedure_declarationContext) {
	localctx = NewMult_line_procedure_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GraffleParserRULE_mult_line_procedure_declaration)

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
		p.SetState(346)
		p.Procedure_declaration_head()
	}
	{
		p.SetState(347)
		p.Match(GraffleParserNEWLINE)
	}
	{
		p.SetState(348)
		p.Sequence()
	}
	{
		p.SetState(349)
		p.Block_end()
	}

	return localctx
}

// IProcedure_declaration_headContext is an interface to support dynamic dispatch.
type IProcedure_declaration_headContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpd1 returns the opd1 rule contexts.
	GetOpd1() IVariableContext

	// GetOpd2 returns the opd2 rule contexts.
	GetOpd2() IVariableContext

	// SetOpd1 sets the opd1 rule contexts.
	SetOpd1(IVariableContext)

	// SetOpd2 sets the opd2 rule contexts.
	SetOpd2(IVariableContext)

	// IsProcedure_declaration_headContext differentiates from other interfaces.
	IsProcedure_declaration_headContext()
}

type Procedure_declaration_headContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	opd1   IVariableContext
	opd2   IVariableContext
}

func NewEmptyProcedure_declaration_headContext() *Procedure_declaration_headContext {
	var p = new(Procedure_declaration_headContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_procedure_declaration_head
	return p
}

func (*Procedure_declaration_headContext) IsProcedure_declaration_headContext() {}

func NewProcedure_declaration_headContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Procedure_declaration_headContext {
	var p = new(Procedure_declaration_headContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_procedure_declaration_head

	return p
}

func (s *Procedure_declaration_headContext) GetParser() antlr.Parser { return s.parser }

func (s *Procedure_declaration_headContext) GetOpd1() IVariableContext { return s.opd1 }

func (s *Procedure_declaration_headContext) GetOpd2() IVariableContext { return s.opd2 }

func (s *Procedure_declaration_headContext) SetOpd1(v IVariableContext) { s.opd1 = v }

func (s *Procedure_declaration_headContext) SetOpd2(v IVariableContext) { s.opd2 = v }

func (s *Procedure_declaration_headContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Procedure_declaration_headContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Procedure_declaration_headContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Procedure_declaration_headContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Procedure_declaration_headContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *Procedure_declaration_headContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Procedure_declaration_headContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *Procedure_declaration_headContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
}

func (s *Procedure_declaration_headContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Procedure_declaration_headContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Procedure_declaration_headContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterProcedure_declaration_head(s)
	}
}

func (s *Procedure_declaration_headContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitProcedure_declaration_head(s)
	}
}

func (s *Procedure_declaration_headContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitProcedure_declaration_head(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Procedure_declaration_head() (localctx IProcedure_declaration_headContext) {
	localctx = NewProcedure_declaration_headContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GraffleParserRULE_procedure_declaration_head)
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
		p.SetState(351)
		p.Match(GraffleParserID)
	}
	{
		p.SetState(352)
		p.Match(GraffleParserL_PAREN)
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserSUB || _la == GraffleParserID {
		{
			p.SetState(353)

			var _x = p.Variable()

			localctx.(*Procedure_declaration_headContext).opd1 = _x
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraffleParserARG_DELIM {
			{
				p.SetState(354)
				p.Match(GraffleParserARG_DELIM)
			}
			{
				p.SetState(355)

				var _x = p.Variable()

				localctx.(*Procedure_declaration_headContext).opd2 = _x
			}

			p.SetState(360)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(363)
		p.Match(GraffleParserR_PAREN)
	}
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserASSIGN {
		{
			p.SetState(364)
			p.Match(GraffleParserASSIGN)
		}

	}

	return localctx
}

// IVar_assignContext is an interface to support dynamic dispatch.
type IVar_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVal returns the val rule contexts.
	GetVal() IVariableContext

	// SetVal sets the val rule contexts.
	SetVal(IVariableContext)

	// IsVar_assignContext differentiates from other interfaces.
	IsVar_assignContext()
}

type Var_assignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	val    IVariableContext
}

func NewEmptyVar_assignContext() *Var_assignContext {
	var p = new(Var_assignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_var_assign
	return p
}

func (*Var_assignContext) IsVar_assignContext() {}

func NewVar_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_assignContext {
	var p = new(Var_assignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_var_assign

	return p
}

func (s *Var_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_assignContext) GetVal() IVariableContext { return s.val }

func (s *Var_assignContext) SetVal(v IVariableContext) { s.val = v }

func (s *Var_assignContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Var_assignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Var_assignContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Var_assignContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Var_assignContext) Arc_assign() IArc_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArc_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArc_assignContext)
}

func (s *Var_assignContext) Vertice_assign() IVertice_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertice_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertice_assignContext)
}

func (s *Var_assignContext) Graph_assign() IGraph_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraph_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGraph_assignContext)
}

func (s *Var_assignContext) Labeled_assign() ILabeled_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabeled_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabeled_assignContext)
}

func (s *Var_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterVar_assign(s)
	}
}

func (s *Var_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitVar_assign(s)
	}
}

func (s *Var_assignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitVar_assign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Var_assign() (localctx IVar_assignContext) {
	localctx = NewVar_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GraffleParserRULE_var_assign)

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

	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(367)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(368)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(369)

			var _x = p.Variable()

			localctx.(*Var_assignContext).val = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(370)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(371)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(372)
			p.Expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.Arc_assign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(374)
			p.Vertice_assign()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(375)
			p.Graph_assign()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(376)
			p.Labeled_assign()
		}

	}

	return localctx
}

// IArc_assignContext is an interface to support dynamic dispatch.
type IArc_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArc_assignContext differentiates from other interfaces.
	IsArc_assignContext()
}

type Arc_assignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArc_assignContext() *Arc_assignContext {
	var p = new(Arc_assignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_arc_assign
	return p
}

func (*Arc_assignContext) IsArc_assignContext() {}

func NewArc_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arc_assignContext {
	var p = new(Arc_assignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_arc_assign

	return p
}

func (s *Arc_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Arc_assignContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Arc_assignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Arc_assignContext) E_N() antlr.TerminalNode {
	return s.GetToken(GraffleParserE_N, 0)
}

func (s *Arc_assignContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Arc_assignContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Arc_assignContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Arc_assignContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Arc_assignContext) Arc() IArcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArcContext)
}

func (s *Arc_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arc_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arc_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterArc_assign(s)
	}
}

func (s *Arc_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitArc_assign(s)
	}
}

func (s *Arc_assignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitArc_assign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Arc_assign() (localctx IArc_assignContext) {
	localctx = NewArc_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GraffleParserRULE_arc_assign)

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

	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(379)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(380)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(381)
			p.Match(GraffleParserE_N)
		}
		{
			p.SetState(382)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(383)
			p.Value()
		}
		{
			p.SetState(384)
			p.Match(GraffleParserR_PAREN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(386)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(387)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(388)
			p.Value()
		}
		{
			p.SetState(389)
			p.Arc()
		}
		{
			p.SetState(390)
			p.Value()
		}

	}

	return localctx
}

// IArcContext is an interface to support dynamic dispatch.
type IArcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArcContext differentiates from other interfaces.
	IsArcContext()
}

type ArcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArcContext() *ArcContext {
	var p = new(ArcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_arc
	return p
}

func (*ArcContext) IsArcContext() {}

func NewArcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArcContext {
	var p = new(ArcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_arc

	return p
}

func (s *ArcContext) GetParser() antlr.Parser { return s.parser }

func (s *ArcContext) OR_ARC_LR() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_ARC_LR, 0)
}

func (s *ArcContext) Or_w_arc_lr() IOr_w_arc_lrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_w_arc_lrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_w_arc_lrContext)
}

func (s *ArcContext) OR_ARC_RL() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_ARC_RL, 0)
}

func (s *ArcContext) Or_w_arc_rl() IOr_w_arc_rlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_w_arc_rlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_w_arc_rlContext)
}

func (s *ArcContext) UNOR_ARC() antlr.TerminalNode {
	return s.GetToken(GraffleParserUNOR_ARC, 0)
}

func (s *ArcContext) Unor_w_arc() IUnor_w_arcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnor_w_arcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnor_w_arcContext)
}

func (s *ArcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterArc(s)
	}
}

func (s *ArcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitArc(s)
	}
}

func (s *ArcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitArc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Arc() (localctx IArcContext) {
	localctx = NewArcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GraffleParserRULE_arc)

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

	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(394)
			p.Match(GraffleParserOR_ARC_LR)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(395)
			p.Or_w_arc_lr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(396)
			p.Match(GraffleParserOR_ARC_RL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(397)
			p.Or_w_arc_rl()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(398)
			p.Match(GraffleParserUNOR_ARC)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(399)
			p.Unor_w_arc()
		}

	}

	return localctx
}

// IOr_w_arc_lrContext is an interface to support dynamic dispatch.
type IOr_w_arc_lrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWeight returns the weight token.
	GetWeight() antlr.Token

	// SetWeight sets the weight token.
	SetWeight(antlr.Token)

	// IsOr_w_arc_lrContext differentiates from other interfaces.
	IsOr_w_arc_lrContext()
}

type Or_w_arc_lrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	weight antlr.Token
}

func NewEmptyOr_w_arc_lrContext() *Or_w_arc_lrContext {
	var p = new(Or_w_arc_lrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_or_w_arc_lr
	return p
}

func (*Or_w_arc_lrContext) IsOr_w_arc_lrContext() {}

func NewOr_w_arc_lrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_w_arc_lrContext {
	var p = new(Or_w_arc_lrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_or_w_arc_lr

	return p
}

func (s *Or_w_arc_lrContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_w_arc_lrContext) GetWeight() antlr.Token { return s.weight }

func (s *Or_w_arc_lrContext) SetWeight(v antlr.Token) { s.weight = v }

func (s *Or_w_arc_lrContext) SUB() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, 0)
}

func (s *Or_w_arc_lrContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_BRACKET, 0)
}

func (s *Or_w_arc_lrContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_BRACKET, 0)
}

func (s *Or_w_arc_lrContext) OR_ARC_LR() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_ARC_LR, 0)
}

func (s *Or_w_arc_lrContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GraffleParserNUMBER, 0)
}

func (s *Or_w_arc_lrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_w_arc_lrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_w_arc_lrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOr_w_arc_lr(s)
	}
}

func (s *Or_w_arc_lrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOr_w_arc_lr(s)
	}
}

func (s *Or_w_arc_lrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOr_w_arc_lr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Or_w_arc_lr() (localctx IOr_w_arc_lrContext) {
	localctx = NewOr_w_arc_lrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GraffleParserRULE_or_w_arc_lr)

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
		p.SetState(402)
		p.Match(GraffleParserSUB)
	}
	{
		p.SetState(403)
		p.Match(GraffleParserL_BRACKET)
	}
	{
		p.SetState(404)

		var _m = p.Match(GraffleParserNUMBER)

		localctx.(*Or_w_arc_lrContext).weight = _m
	}
	{
		p.SetState(405)
		p.Match(GraffleParserR_BRACKET)
	}
	{
		p.SetState(406)
		p.Match(GraffleParserOR_ARC_LR)
	}

	return localctx
}

// IOr_w_arc_rlContext is an interface to support dynamic dispatch.
type IOr_w_arc_rlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWeight returns the weight token.
	GetWeight() antlr.Token

	// SetWeight sets the weight token.
	SetWeight(antlr.Token)

	// IsOr_w_arc_rlContext differentiates from other interfaces.
	IsOr_w_arc_rlContext()
}

type Or_w_arc_rlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	weight antlr.Token
}

func NewEmptyOr_w_arc_rlContext() *Or_w_arc_rlContext {
	var p = new(Or_w_arc_rlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_or_w_arc_rl
	return p
}

func (*Or_w_arc_rlContext) IsOr_w_arc_rlContext() {}

func NewOr_w_arc_rlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_w_arc_rlContext {
	var p = new(Or_w_arc_rlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_or_w_arc_rl

	return p
}

func (s *Or_w_arc_rlContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_w_arc_rlContext) GetWeight() antlr.Token { return s.weight }

func (s *Or_w_arc_rlContext) SetWeight(v antlr.Token) { s.weight = v }

func (s *Or_w_arc_rlContext) OR_ARC_RL() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_ARC_RL, 0)
}

func (s *Or_w_arc_rlContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_BRACKET, 0)
}

func (s *Or_w_arc_rlContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_BRACKET, 0)
}

func (s *Or_w_arc_rlContext) SUB() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, 0)
}

func (s *Or_w_arc_rlContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GraffleParserNUMBER, 0)
}

func (s *Or_w_arc_rlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_w_arc_rlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_w_arc_rlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOr_w_arc_rl(s)
	}
}

func (s *Or_w_arc_rlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOr_w_arc_rl(s)
	}
}

func (s *Or_w_arc_rlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOr_w_arc_rl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Or_w_arc_rl() (localctx IOr_w_arc_rlContext) {
	localctx = NewOr_w_arc_rlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GraffleParserRULE_or_w_arc_rl)

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
		p.SetState(408)
		p.Match(GraffleParserOR_ARC_RL)
	}
	{
		p.SetState(409)
		p.Match(GraffleParserL_BRACKET)
	}
	{
		p.SetState(410)

		var _m = p.Match(GraffleParserNUMBER)

		localctx.(*Or_w_arc_rlContext).weight = _m
	}
	{
		p.SetState(411)
		p.Match(GraffleParserR_BRACKET)
	}
	{
		p.SetState(412)
		p.Match(GraffleParserSUB)
	}

	return localctx
}

// IUnor_w_arcContext is an interface to support dynamic dispatch.
type IUnor_w_arcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWeight returns the weight token.
	GetWeight() antlr.Token

	// SetWeight sets the weight token.
	SetWeight(antlr.Token)

	// IsUnor_w_arcContext differentiates from other interfaces.
	IsUnor_w_arcContext()
}

type Unor_w_arcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	weight antlr.Token
}

func NewEmptyUnor_w_arcContext() *Unor_w_arcContext {
	var p = new(Unor_w_arcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_unor_w_arc
	return p
}

func (*Unor_w_arcContext) IsUnor_w_arcContext() {}

func NewUnor_w_arcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unor_w_arcContext {
	var p = new(Unor_w_arcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_unor_w_arc

	return p
}

func (s *Unor_w_arcContext) GetParser() antlr.Parser { return s.parser }

func (s *Unor_w_arcContext) GetWeight() antlr.Token { return s.weight }

func (s *Unor_w_arcContext) SetWeight(v antlr.Token) { s.weight = v }

func (s *Unor_w_arcContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserSUB)
}

func (s *Unor_w_arcContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, i)
}

func (s *Unor_w_arcContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_BRACKET, 0)
}

func (s *Unor_w_arcContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_BRACKET, 0)
}

func (s *Unor_w_arcContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GraffleParserNUMBER, 0)
}

func (s *Unor_w_arcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unor_w_arcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unor_w_arcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterUnor_w_arc(s)
	}
}

func (s *Unor_w_arcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitUnor_w_arc(s)
	}
}

func (s *Unor_w_arcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitUnor_w_arc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Unor_w_arc() (localctx IUnor_w_arcContext) {
	localctx = NewUnor_w_arcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GraffleParserRULE_unor_w_arc)

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
		p.SetState(414)
		p.Match(GraffleParserSUB)
	}
	{
		p.SetState(415)
		p.Match(GraffleParserL_BRACKET)
	}
	{
		p.SetState(416)

		var _m = p.Match(GraffleParserNUMBER)

		localctx.(*Unor_w_arcContext).weight = _m
	}
	{
		p.SetState(417)
		p.Match(GraffleParserR_BRACKET)
	}
	{
		p.SetState(418)
		p.Match(GraffleParserSUB)
	}

	return localctx
}

// IVertice_assignContext is an interface to support dynamic dispatch.
type IVertice_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVertice_assignContext differentiates from other interfaces.
	IsVertice_assignContext()
}

type Vertice_assignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVertice_assignContext() *Vertice_assignContext {
	var p = new(Vertice_assignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_vertice_assign
	return p
}

func (*Vertice_assignContext) IsVertice_assignContext() {}

func NewVertice_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vertice_assignContext {
	var p = new(Vertice_assignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_vertice_assign

	return p
}

func (s *Vertice_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Vertice_assignContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Vertice_assignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Vertice_assignContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Vertice_assignContext) V_N() antlr.TerminalNode {
	return s.GetToken(GraffleParserV_N, 0)
}

func (s *Vertice_assignContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Vertice_assignContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Vertice_assignContext) Arithm_assign_operator() IArithm_assign_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_assign_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_assign_operatorContext)
}

func (s *Vertice_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vertice_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vertice_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterVertice_assign(s)
	}
}

func (s *Vertice_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitVertice_assign(s)
	}
}

func (s *Vertice_assignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitVertice_assign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Vertice_assign() (localctx IVertice_assignContext) {
	localctx = NewVertice_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GraffleParserRULE_vertice_assign)

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

	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(420)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(421)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(422)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(423)
			p.Match(GraffleParserV_N)
		}
		{
			p.SetState(424)
			p.Match(GraffleParserR_PAREN)
		}
		{
			p.SetState(425)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(426)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(427)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(428)
			p.Value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(429)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(430)
			p.Arithm_assign_operator()
		}
		{
			p.SetState(431)
			p.Value()
		}

	}

	return localctx
}

// IGraph_assignContext is an interface to support dynamic dispatch.
type IGraph_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGraph_assignContext differentiates from other interfaces.
	IsGraph_assignContext()
}

type Graph_assignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraph_assignContext() *Graph_assignContext {
	var p = new(Graph_assignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_graph_assign
	return p
}

func (*Graph_assignContext) IsGraph_assignContext() {}

func NewGraph_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Graph_assignContext {
	var p = new(Graph_assignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_graph_assign

	return p
}

func (s *Graph_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Graph_assignContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Graph_assignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Graph_assignContext) G_N() antlr.TerminalNode {
	return s.GetToken(GraffleParserG_N, 0)
}

func (s *Graph_assignContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Graph_assignContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Graph_assignContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Graph_assignContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Graph_assignContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *Graph_assignContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
}

func (s *Graph_assignContext) Arithm_assign_operator() IArithm_assign_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_assign_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_assign_operatorContext)
}

func (s *Graph_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Graph_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Graph_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterGraph_assign(s)
	}
}

func (s *Graph_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitGraph_assign(s)
	}
}

func (s *Graph_assignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitGraph_assign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Graph_assign() (localctx IGraph_assignContext) {
	localctx = NewGraph_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GraffleParserRULE_graph_assign)

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

	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(436)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(437)
			p.Match(GraffleParserG_N)
		}
		{
			p.SetState(438)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(439)
			p.Value()
		}
		{
			p.SetState(440)
			p.Match(GraffleParserR_PAREN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(442)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(443)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(444)
			p.Value()
		}
		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(445)
					p.Match(GraffleParserARG_DELIM)
				}
				{
					p.SetState(446)
					p.Value()
				}

			}
			p.SetState(451)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(452)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(453)
			p.Arithm_assign_operator()
		}
		{
			p.SetState(454)
			p.Value()
		}

	}

	return localctx
}

// ILabeled_assignContext is an interface to support dynamic dispatch.
type ILabeled_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabeled_assignContext differentiates from other interfaces.
	IsLabeled_assignContext()
}

type Labeled_assignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabeled_assignContext() *Labeled_assignContext {
	var p = new(Labeled_assignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_labeled_assign
	return p
}

func (*Labeled_assignContext) IsLabeled_assignContext() {}

func NewLabeled_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Labeled_assignContext {
	var p = new(Labeled_assignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_labeled_assign

	return p
}

func (s *Labeled_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Labeled_assignContext) Vertice_assign() IVertice_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertice_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertice_assignContext)
}

func (s *Labeled_assignContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *Labeled_assignContext) Arc_assign() IArc_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArc_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArc_assignContext)
}

func (s *Labeled_assignContext) Graph_assign() IGraph_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraph_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGraph_assignContext)
}

func (s *Labeled_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Labeled_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Labeled_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterLabeled_assign(s)
	}
}

func (s *Labeled_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitLabeled_assign(s)
	}
}

func (s *Labeled_assignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitLabeled_assign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Labeled_assign() (localctx ILabeled_assignContext) {
	localctx = NewLabeled_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GraffleParserRULE_labeled_assign)
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

	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(458)
			p.Vertice_assign()
		}
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserLABEL || _la == GraffleParserML_LABEL {
			{
				p.SetState(459)
				p.Label()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(462)
			p.Arc_assign()
		}
		p.SetState(464)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserLABEL || _la == GraffleParserML_LABEL {
			{
				p.SetState(463)
				p.Label()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(466)
			p.Graph_assign()
		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserLABEL || _la == GraffleParserML_LABEL {
			{
				p.SetState(467)
				p.Label()
			}

		}

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
	p.RuleIndex = GraffleParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExprContext) Integral_expr() IIntegral_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegral_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegral_exprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GraffleParserRULE_expr)

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

	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(472)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(473)
			p.Integral_expr()
		}

	}

	return localctx
}

// IIntegral_exprContext is an interface to support dynamic dispatch.
type IIntegral_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegral_exprContext differentiates from other interfaces.
	IsIntegral_exprContext()
}

type Integral_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegral_exprContext() *Integral_exprContext {
	var p = new(Integral_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_integral_expr
	return p
}

func (*Integral_exprContext) IsIntegral_exprContext() {}

func NewIntegral_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integral_exprContext {
	var p = new(Integral_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_integral_expr

	return p
}

func (s *Integral_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Integral_exprContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Integral_exprContext) Builtin() IBuiltinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltinContext)
}

func (s *Integral_exprContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Integral_exprContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *Integral_exprContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Integral_exprContext) Arithm_expr() IArithm_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_exprContext)
}

func (s *Integral_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integral_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integral_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterIntegral_expr(s)
	}
}

func (s *Integral_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitIntegral_expr(s)
	}
}

func (s *Integral_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitIntegral_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Integral_expr() (localctx IIntegral_exprContext) {
	localctx = NewIntegral_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GraffleParserRULE_integral_expr)

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

	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(476)
			p.Function_call()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(477)
			p.Builtin()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(478)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(479)
			p.Logical_expr()
		}
		{
			p.SetState(480)
			p.Match(GraffleParserR_PAREN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(482)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(483)
			p.Arithm_expr()
		}
		{
			p.SetState(484)
			p.Match(GraffleParserR_PAREN)
		}

	}

	return localctx
}

// ILogical_exprContext is an interface to support dynamic dispatch.
type ILogical_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetBin_op returns the bin_op rule contexts.
	GetBin_op() IBin_log_operatorContext

	// GetRight returns the right rule contexts.
	GetRight() ILog_expr_operandContext

	// GetUn_op returns the un_op rule contexts.
	GetUn_op() IUnar_log_operatorContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetBin_op sets the bin_op rule contexts.
	SetBin_op(IBin_log_operatorContext)

	// SetRight sets the right rule contexts.
	SetRight(ILog_expr_operandContext)

	// SetUn_op sets the un_op rule contexts.
	SetUn_op(IUnar_log_operatorContext)

	// IsLogical_exprContext differentiates from other interfaces.
	IsLogical_exprContext()
}

type Logical_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExprContext
	bin_op IBin_log_operatorContext
	right  ILog_expr_operandContext
	un_op  IUnar_log_operatorContext
}

func NewEmptyLogical_exprContext() *Logical_exprContext {
	var p = new(Logical_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_logical_expr
	return p
}

func (*Logical_exprContext) IsLogical_exprContext() {}

func NewLogical_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_exprContext {
	var p = new(Logical_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_logical_expr

	return p
}

func (s *Logical_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_exprContext) GetLeft() IExprContext { return s.left }

func (s *Logical_exprContext) GetBin_op() IBin_log_operatorContext { return s.bin_op }

func (s *Logical_exprContext) GetRight() ILog_expr_operandContext { return s.right }

func (s *Logical_exprContext) GetUn_op() IUnar_log_operatorContext { return s.un_op }

func (s *Logical_exprContext) SetLeft(v IExprContext) { s.left = v }

func (s *Logical_exprContext) SetBin_op(v IBin_log_operatorContext) { s.bin_op = v }

func (s *Logical_exprContext) SetRight(v ILog_expr_operandContext) { s.right = v }

func (s *Logical_exprContext) SetUn_op(v IUnar_log_operatorContext) { s.un_op = v }

func (s *Logical_exprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Logical_exprContext) Bin_log_operator() IBin_log_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBin_log_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBin_log_operatorContext)
}

func (s *Logical_exprContext) Log_expr_operand() ILog_expr_operandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILog_expr_operandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILog_expr_operandContext)
}

func (s *Logical_exprContext) Unar_log_operator() IUnar_log_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnar_log_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnar_log_operatorContext)
}

func (s *Logical_exprContext) Arithm_expr() IArithm_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_exprContext)
}

func (s *Logical_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterLogical_expr(s)
	}
}

func (s *Logical_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitLogical_expr(s)
	}
}

func (s *Logical_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitLogical_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Logical_expr() (localctx ILogical_exprContext) {
	localctx = NewLogical_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GraffleParserRULE_logical_expr)

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

	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(488)

			var _x = p.Expr()

			localctx.(*Logical_exprContext).left = _x
		}
		{
			p.SetState(489)

			var _x = p.Bin_log_operator()

			localctx.(*Logical_exprContext).bin_op = _x
		}
		{
			p.SetState(490)

			var _x = p.Log_expr_operand()

			localctx.(*Logical_exprContext).right = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(492)

			var _x = p.Unar_log_operator()

			localctx.(*Logical_exprContext).un_op = _x
		}
		{
			p.SetState(493)
			p.Log_expr_operand()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(495)
			p.Arithm_expr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(496)
			p.Expr()
		}

	}

	return localctx
}

// ILog_expr_operandContext is an interface to support dynamic dispatch.
type ILog_expr_operandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLog_expr_operandContext differentiates from other interfaces.
	IsLog_expr_operandContext()
}

type Log_expr_operandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLog_expr_operandContext() *Log_expr_operandContext {
	var p = new(Log_expr_operandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_log_expr_operand
	return p
}

func (*Log_expr_operandContext) IsLog_expr_operandContext() {}

func NewLog_expr_operandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Log_expr_operandContext {
	var p = new(Log_expr_operandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_log_expr_operand

	return p
}

func (s *Log_expr_operandContext) GetParser() antlr.Parser { return s.parser }

func (s *Log_expr_operandContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Log_expr_operandContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *Log_expr_operandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Log_expr_operandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Log_expr_operandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterLog_expr_operand(s)
	}
}

func (s *Log_expr_operandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitLog_expr_operand(s)
	}
}

func (s *Log_expr_operandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitLog_expr_operand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Log_expr_operand() (localctx ILog_expr_operandContext) {
	localctx = NewLog_expr_operandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GraffleParserRULE_log_expr_operand)

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

	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(500)
			p.Logical_expr()
		}

	}

	return localctx
}

// IBin_log_operatorContext is an interface to support dynamic dispatch.
type IBin_log_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBin_log_operatorContext differentiates from other interfaces.
	IsBin_log_operatorContext()
}

type Bin_log_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBin_log_operatorContext() *Bin_log_operatorContext {
	var p = new(Bin_log_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_bin_log_operator
	return p
}

func (*Bin_log_operatorContext) IsBin_log_operatorContext() {}

func NewBin_log_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bin_log_operatorContext {
	var p = new(Bin_log_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_bin_log_operator

	return p
}

func (s *Bin_log_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Bin_log_operatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GraffleParserEQUALS, 0)
}

func (s *Bin_log_operatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEQ, 0)
}

func (s *Bin_log_operatorContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(GraffleParserLESS_THAN, 0)
}

func (s *Bin_log_operatorContext) GR_THAN() antlr.TerminalNode {
	return s.GetToken(GraffleParserGR_THAN, 0)
}

func (s *Bin_log_operatorContext) LESS_THAN_E() antlr.TerminalNode {
	return s.GetToken(GraffleParserLESS_THAN_E, 0)
}

func (s *Bin_log_operatorContext) GR_THAN_E() antlr.TerminalNode {
	return s.GetToken(GraffleParserGR_THAN_E, 0)
}

func (s *Bin_log_operatorContext) AND() antlr.TerminalNode {
	return s.GetToken(GraffleParserAND, 0)
}

func (s *Bin_log_operatorContext) NAND() antlr.TerminalNode {
	return s.GetToken(GraffleParserNAND, 0)
}

func (s *Bin_log_operatorContext) OR() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR, 0)
}

func (s *Bin_log_operatorContext) NOR() antlr.TerminalNode {
	return s.GetToken(GraffleParserNOR, 0)
}

func (s *Bin_log_operatorContext) XOR() antlr.TerminalNode {
	return s.GetToken(GraffleParserXOR, 0)
}

func (s *Bin_log_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bin_log_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bin_log_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBin_log_operator(s)
	}
}

func (s *Bin_log_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBin_log_operator(s)
	}
}

func (s *Bin_log_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBin_log_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Bin_log_operator() (localctx IBin_log_operatorContext) {
	localctx = NewBin_log_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GraffleParserRULE_bin_log_operator)
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
		p.SetState(503)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-23)&-(0x1f+1)) == 0 && ((1<<uint((_la-23)))&((1<<(GraffleParserNEQ-23))|(1<<(GraffleParserEQUALS-23))|(1<<(GraffleParserLESS_THAN-23))|(1<<(GraffleParserGR_THAN-23))|(1<<(GraffleParserLESS_THAN_E-23))|(1<<(GraffleParserGR_THAN_E-23))|(1<<(GraffleParserAND-23))|(1<<(GraffleParserOR-23))|(1<<(GraffleParserXOR-23))|(1<<(GraffleParserNOR-23))|(1<<(GraffleParserNAND-23)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnar_log_operatorContext is an interface to support dynamic dispatch.
type IUnar_log_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnar_log_operatorContext differentiates from other interfaces.
	IsUnar_log_operatorContext()
}

type Unar_log_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnar_log_operatorContext() *Unar_log_operatorContext {
	var p = new(Unar_log_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_unar_log_operator
	return p
}

func (*Unar_log_operatorContext) IsUnar_log_operatorContext() {}

func NewUnar_log_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unar_log_operatorContext {
	var p = new(Unar_log_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_unar_log_operator

	return p
}

func (s *Unar_log_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Unar_log_operatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(GraffleParserNOT, 0)
}

func (s *Unar_log_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unar_log_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unar_log_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterUnar_log_operator(s)
	}
}

func (s *Unar_log_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitUnar_log_operator(s)
	}
}

func (s *Unar_log_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitUnar_log_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Unar_log_operator() (localctx IUnar_log_operatorContext) {
	localctx = NewUnar_log_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GraffleParserRULE_unar_log_operator)

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
		p.SetState(505)
		p.Match(GraffleParserNOT)
	}

	return localctx
}

// IArithm_exprContext is an interface to support dynamic dispatch.
type IArithm_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetOp returns the op rule contexts.
	GetOp() IBin_arithm_operatorContext

	// GetRight returns the right rule contexts.
	GetRight() IArithm_expr_operandContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetOp sets the op rule contexts.
	SetOp(IBin_arithm_operatorContext)

	// SetRight sets the right rule contexts.
	SetRight(IArithm_expr_operandContext)

	// IsArithm_exprContext differentiates from other interfaces.
	IsArithm_exprContext()
}

type Arithm_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExprContext
	op     IBin_arithm_operatorContext
	right  IArithm_expr_operandContext
}

func NewEmptyArithm_exprContext() *Arithm_exprContext {
	var p = new(Arithm_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_arithm_expr
	return p
}

func (*Arithm_exprContext) IsArithm_exprContext() {}

func NewArithm_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithm_exprContext {
	var p = new(Arithm_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_arithm_expr

	return p
}

func (s *Arithm_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithm_exprContext) GetLeft() IExprContext { return s.left }

func (s *Arithm_exprContext) GetOp() IBin_arithm_operatorContext { return s.op }

func (s *Arithm_exprContext) GetRight() IArithm_expr_operandContext { return s.right }

func (s *Arithm_exprContext) SetLeft(v IExprContext) { s.left = v }

func (s *Arithm_exprContext) SetOp(v IBin_arithm_operatorContext) { s.op = v }

func (s *Arithm_exprContext) SetRight(v IArithm_expr_operandContext) { s.right = v }

func (s *Arithm_exprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Arithm_exprContext) Bin_arithm_operator() IBin_arithm_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBin_arithm_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBin_arithm_operatorContext)
}

func (s *Arithm_exprContext) Arithm_expr_operand() IArithm_expr_operandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_expr_operandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_expr_operandContext)
}

func (s *Arithm_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithm_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arithm_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterArithm_expr(s)
	}
}

func (s *Arithm_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitArithm_expr(s)
	}
}

func (s *Arithm_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitArithm_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Arithm_expr() (localctx IArithm_exprContext) {
	localctx = NewArithm_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GraffleParserRULE_arithm_expr)

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
		p.SetState(507)

		var _x = p.Expr()

		localctx.(*Arithm_exprContext).left = _x
	}
	{
		p.SetState(508)

		var _x = p.Bin_arithm_operator()

		localctx.(*Arithm_exprContext).op = _x
	}
	{
		p.SetState(509)

		var _x = p.Arithm_expr_operand()

		localctx.(*Arithm_exprContext).right = _x
	}

	return localctx
}

// IArithm_expr_operandContext is an interface to support dynamic dispatch.
type IArithm_expr_operandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArithm_expr_operandContext differentiates from other interfaces.
	IsArithm_expr_operandContext()
}

type Arithm_expr_operandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithm_expr_operandContext() *Arithm_expr_operandContext {
	var p = new(Arithm_expr_operandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_arithm_expr_operand
	return p
}

func (*Arithm_expr_operandContext) IsArithm_expr_operandContext() {}

func NewArithm_expr_operandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithm_expr_operandContext {
	var p = new(Arithm_expr_operandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_arithm_expr_operand

	return p
}

func (s *Arithm_expr_operandContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithm_expr_operandContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Arithm_expr_operandContext) Arithm_expr() IArithm_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_exprContext)
}

func (s *Arithm_expr_operandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithm_expr_operandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arithm_expr_operandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterArithm_expr_operand(s)
	}
}

func (s *Arithm_expr_operandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitArithm_expr_operand(s)
	}
}

func (s *Arithm_expr_operandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitArithm_expr_operand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Arithm_expr_operand() (localctx IArithm_expr_operandContext) {
	localctx = NewArithm_expr_operandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GraffleParserRULE_arithm_expr_operand)

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

	p.SetState(513)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(511)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(512)
			p.Arithm_expr()
		}

	}

	return localctx
}

// IBin_arithm_operatorContext is an interface to support dynamic dispatch.
type IBin_arithm_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBin_arithm_operatorContext differentiates from other interfaces.
	IsBin_arithm_operatorContext()
}

type Bin_arithm_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBin_arithm_operatorContext() *Bin_arithm_operatorContext {
	var p = new(Bin_arithm_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_bin_arithm_operator
	return p
}

func (*Bin_arithm_operatorContext) IsBin_arithm_operatorContext() {}

func NewBin_arithm_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bin_arithm_operatorContext {
	var p = new(Bin_arithm_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_bin_arithm_operator

	return p
}

func (s *Bin_arithm_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Bin_arithm_operatorContext) MULT() antlr.TerminalNode {
	return s.GetToken(GraffleParserMULT, 0)
}

func (s *Bin_arithm_operatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(GraffleParserDIV, 0)
}

func (s *Bin_arithm_operatorContext) ADD() antlr.TerminalNode {
	return s.GetToken(GraffleParserADD, 0)
}

func (s *Bin_arithm_operatorContext) SUB() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, 0)
}

func (s *Bin_arithm_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bin_arithm_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bin_arithm_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBin_arithm_operator(s)
	}
}

func (s *Bin_arithm_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBin_arithm_operator(s)
	}
}

func (s *Bin_arithm_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBin_arithm_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Bin_arithm_operator() (localctx IBin_arithm_operatorContext) {
	localctx = NewBin_arithm_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GraffleParserRULE_bin_arithm_operator)
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
		p.SetState(515)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraffleParserADD)|(1<<GraffleParserSUB)|(1<<GraffleParserMULT)|(1<<GraffleParserDIV))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArithm_assign_operatorContext is an interface to support dynamic dispatch.
type IArithm_assign_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArithm_assign_operatorContext differentiates from other interfaces.
	IsArithm_assign_operatorContext()
}

type Arithm_assign_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithm_assign_operatorContext() *Arithm_assign_operatorContext {
	var p = new(Arithm_assign_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_arithm_assign_operator
	return p
}

func (*Arithm_assign_operatorContext) IsArithm_assign_operatorContext() {}

func NewArithm_assign_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithm_assign_operatorContext {
	var p = new(Arithm_assign_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_arithm_assign_operator

	return p
}

func (s *Arithm_assign_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithm_assign_operatorContext) ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserADD_ASSIGN, 0)
}

func (s *Arithm_assign_operatorContext) SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB_ASSIGN, 0)
}

func (s *Arithm_assign_operatorContext) MULT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserMULT_ASSIGN, 0)
}

func (s *Arithm_assign_operatorContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserDIV_ASSIGN, 0)
}

func (s *Arithm_assign_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithm_assign_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arithm_assign_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterArithm_assign_operator(s)
	}
}

func (s *Arithm_assign_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitArithm_assign_operator(s)
	}
}

func (s *Arithm_assign_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitArithm_assign_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Arithm_assign_operator() (localctx IArithm_assign_operatorContext) {
	localctx = NewArithm_assign_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GraffleParserRULE_arithm_assign_operator)
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
		p.SetState(517)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraffleParserADD_ASSIGN)|(1<<GraffleParserSUB_ASSIGN)|(1<<GraffleParserMULT_ASSIGN)|(1<<GraffleParserDIV_ASSIGN))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBuiltin_function_callContext is an interface to support dynamic dispatch.
type IBuiltin_function_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltin_function_callContext differentiates from other interfaces.
	IsBuiltin_function_callContext()
}

type Builtin_function_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltin_function_callContext() *Builtin_function_callContext {
	var p = new(Builtin_function_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_builtin_function_call
	return p
}

func (*Builtin_function_callContext) IsBuiltin_function_callContext() {}

func NewBuiltin_function_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Builtin_function_callContext {
	var p = new(Builtin_function_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_builtin_function_call

	return p
}

func (s *Builtin_function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Builtin_function_callContext) Built_func_print() IBuilt_func_printContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuilt_func_printContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuilt_func_printContext)
}

func (s *Builtin_function_callContext) Built_func_input() IBuilt_func_inputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuilt_func_inputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuilt_func_inputContext)
}

func (s *Builtin_function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Builtin_function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Builtin_function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBuiltin_function_call(s)
	}
}

func (s *Builtin_function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBuiltin_function_call(s)
	}
}

func (s *Builtin_function_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBuiltin_function_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Builtin_function_call() (localctx IBuiltin_function_callContext) {
	localctx = NewBuiltin_function_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GraffleParserRULE_builtin_function_call)

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

	p.SetState(521)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserPRINTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(519)
			p.Built_func_print()
		}

	case GraffleParserKEY_INPUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(520)
			p.Built_func_input()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBuilt_func_printContext is an interface to support dynamic dispatch.
type IBuilt_func_printContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuilt_func_printContext differentiates from other interfaces.
	IsBuilt_func_printContext()
}

type Built_func_printContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuilt_func_printContext() *Built_func_printContext {
	var p = new(Built_func_printContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_built_func_print
	return p
}

func (*Built_func_printContext) IsBuilt_func_printContext() {}

func NewBuilt_func_printContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Built_func_printContext {
	var p = new(Built_func_printContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_built_func_print

	return p
}

func (s *Built_func_printContext) GetParser() antlr.Parser { return s.parser }

func (s *Built_func_printContext) PRINTER() antlr.TerminalNode {
	return s.GetToken(GraffleParserPRINTER, 0)
}

func (s *Built_func_printContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Built_func_printContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Built_func_printContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Built_func_printContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Built_func_printContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Built_func_printContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBuilt_func_print(s)
	}
}

func (s *Built_func_printContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBuilt_func_print(s)
	}
}

func (s *Built_func_printContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBuilt_func_print(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Built_func_print() (localctx IBuilt_func_printContext) {
	localctx = NewBuilt_func_printContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, GraffleParserRULE_built_func_print)

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

	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(523)
			p.Match(GraffleParserPRINTER)
		}
		{
			p.SetState(524)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(525)
			p.Match(GraffleParserPRINTER)
		}
		{
			p.SetState(526)
			p.Value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(527)
			p.Match(GraffleParserPRINTER)
		}
		{
			p.SetState(528)
			p.Function_call()
		}

	}

	return localctx
}

// IBuilt_func_inputContext is an interface to support dynamic dispatch.
type IBuilt_func_inputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuilt_func_inputContext differentiates from other interfaces.
	IsBuilt_func_inputContext()
}

type Built_func_inputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuilt_func_inputContext() *Built_func_inputContext {
	var p = new(Built_func_inputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_built_func_input
	return p
}

func (*Built_func_inputContext) IsBuilt_func_inputContext() {}

func NewBuilt_func_inputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Built_func_inputContext {
	var p = new(Built_func_inputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_built_func_input

	return p
}

func (s *Built_func_inputContext) GetParser() antlr.Parser { return s.parser }

func (s *Built_func_inputContext) KEY_INPUT() antlr.TerminalNode {
	return s.GetToken(GraffleParserKEY_INPUT, 0)
}

func (s *Built_func_inputContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Built_func_inputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Built_func_inputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Built_func_inputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBuilt_func_input(s)
	}
}

func (s *Built_func_inputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBuilt_func_input(s)
	}
}

func (s *Built_func_inputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBuilt_func_input(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Built_func_input() (localctx IBuilt_func_inputContext) {
	localctx = NewBuilt_func_inputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, GraffleParserRULE_built_func_input)

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
		p.SetState(531)
		p.Match(GraffleParserKEY_INPUT)
	}
	{
		p.SetState(532)
		p.Match(GraffleParserID)
	}

	return localctx
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_function_call
	return p
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Function_callContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Function_callContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Function_callContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Function_callContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Function_callContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *Function_callContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (s *Function_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitFunction_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, GraffleParserRULE_function_call)
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
		p.SetState(534)
		p.Match(GraffleParserID)
	}
	{
		p.SetState(535)
		p.Match(GraffleParserL_PAREN)
	}
	p.SetState(544)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraffleParserSTRING)|(1<<GraffleParserBOOL)|(1<<GraffleParserL_PAREN)|(1<<GraffleParserSUB))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(GraffleParserNOT-34))|(1<<(GraffleParserID-34))|(1<<(GraffleParserNUMBER-34)))) != 0) {
		{
			p.SetState(536)
			p.Value()
		}
		p.SetState(541)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraffleParserARG_DELIM {
			{
				p.SetState(537)
				p.Match(GraffleParserARG_DELIM)
			}
			{
				p.SetState(538)
				p.Value()
			}

			p.SetState(543)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(546)
		p.Match(GraffleParserR_PAREN)
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) LABEL() antlr.TerminalNode {
	return s.GetToken(GraffleParserLABEL, 0)
}

func (s *LabelContext) ML_LABEL() antlr.TerminalNode {
	return s.GetToken(GraffleParserML_LABEL, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, GraffleParserRULE_label)
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
		p.SetState(548)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GraffleParserLABEL || _la == GraffleParserML_LABEL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *ValueContext) Arithm_expr() IArithm_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_exprContext)
}

func (s *ValueContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueContext) Builtin() IBuiltinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltinContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, GraffleParserRULE_value)

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

	p.SetState(554)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 71, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(550)
			p.Logical_expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(551)
			p.Arithm_expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(552)
			p.Expr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(553)
			p.Builtin()
		}

	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *VariableContext) SUB() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, GraffleParserRULE_variable)
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
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserSUB {
		{
			p.SetState(556)
			p.Match(GraffleParserSUB)
		}

	}
	{
		p.SetState(559)
		p.Match(GraffleParserID)
	}

	return localctx
}

// IBuiltinContext is an interface to support dynamic dispatch.
type IBuiltinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltinContext differentiates from other interfaces.
	IsBuiltinContext()
}

type BuiltinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinContext() *BuiltinContext {
	var p = new(BuiltinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_builtin
	return p
}

func (*BuiltinContext) IsBuiltinContext() {}

func NewBuiltinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinContext {
	var p = new(BuiltinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_builtin

	return p
}

func (s *BuiltinContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltinContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GraffleParserNUMBER, 0)
}

func (s *BuiltinContext) STRING() antlr.TerminalNode {
	return s.GetToken(GraffleParserSTRING, 0)
}

func (s *BuiltinContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GraffleParserBOOL, 0)
}

func (s *BuiltinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBuiltin(s)
	}
}

func (s *BuiltinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBuiltin(s)
	}
}

func (s *BuiltinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBuiltin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Builtin() (localctx IBuiltinContext) {
	localctx = NewBuiltinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, GraffleParserRULE_builtin)
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == GraffleParserSTRING || _la == GraffleParserBOOL || _la == GraffleParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBlock_endContext is an interface to support dynamic dispatch.
type IBlock_endContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlock_endContext differentiates from other interfaces.
	IsBlock_endContext()
}

type Block_endContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_endContext() *Block_endContext {
	var p = new(Block_endContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_block_end
	return p
}

func (*Block_endContext) IsBlock_endContext() {}

func NewBlock_endContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_endContext {
	var p = new(Block_endContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_block_end

	return p
}

func (s *Block_endContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_endContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(GraffleParserBLOCK_END, 0)
}

func (s *Block_endContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserNEWLINE)
}

func (s *Block_endContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, i)
}

func (s *Block_endContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_endContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_endContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBlock_end(s)
	}
}

func (s *Block_endContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBlock_end(s)
	}
}

func (s *Block_endContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBlock_end(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Block_end() (localctx IBlock_endContext) {
	localctx = NewBlock_endContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, GraffleParserRULE_block_end)
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
	p.SetState(564)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserNEWLINE {
		{
			p.SetState(563)
			p.Match(GraffleParserNEWLINE)
		}

	}
	{
		p.SetState(566)
		p.Match(GraffleParserBLOCK_END)
	}
	p.SetState(568)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 74, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(567)
			p.Match(GraffleParserNEWLINE)
		}

	}

	return localctx
}
