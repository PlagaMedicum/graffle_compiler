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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 68, 622,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 3, 2, 3, 2, 3, 2, 5, 2, 124, 10, 2, 3, 3, 5, 3, 127, 10, 3, 3, 3, 7,
	3, 130, 10, 3, 12, 3, 14, 3, 133, 11, 3, 3, 4, 5, 4, 136, 10, 4, 3, 4,
	3, 4, 6, 4, 140, 10, 4, 13, 4, 14, 4, 141, 3, 4, 5, 4, 145, 10, 4, 6, 4,
	147, 10, 4, 13, 4, 14, 4, 148, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 155, 10, 5,
	3, 6, 3, 6, 3, 6, 7, 6, 160, 10, 6, 12, 6, 14, 6, 163, 11, 6, 3, 6, 5,
	6, 166, 10, 6, 3, 7, 3, 7, 5, 7, 170, 10, 7, 3, 8, 3, 8, 3, 8, 5, 8, 175,
	10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 180, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9,
	186, 10, 9, 5, 9, 188, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 193, 10, 9, 3, 9,
	3, 9, 3, 10, 5, 10, 198, 10, 10, 3, 10, 3, 10, 5, 10, 202, 10, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 209, 10, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 5, 11, 215, 10, 11, 5, 11, 217, 10, 11, 3, 11, 3, 11, 3, 11, 5, 11,
	222, 10, 11, 3, 12, 3, 12, 3, 12, 7, 12, 227, 10, 12, 12, 12, 14, 12, 230,
	11, 12, 3, 12, 5, 12, 233, 10, 12, 3, 12, 3, 12, 3, 13, 5, 13, 238, 10,
	13, 3, 13, 3, 13, 3, 13, 5, 13, 243, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	5, 13, 249, 10, 13, 5, 13, 251, 10, 13, 3, 13, 3, 13, 3, 14, 5, 14, 256,
	10, 14, 3, 14, 3, 14, 5, 14, 260, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 275,
	10, 17, 3, 18, 3, 18, 3, 18, 5, 18, 280, 10, 18, 3, 19, 3, 19, 3, 19, 5,
	19, 285, 10, 19, 3, 20, 3, 20, 3, 20, 5, 20, 290, 10, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 299, 10, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 5, 20, 305, 10, 20, 3, 20, 5, 20, 308, 10, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 5, 20, 314, 10, 20, 5, 20, 316, 10, 20, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 5, 21, 323, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22,
	329, 10, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 7, 25, 343, 10, 25, 12, 25, 14, 25, 346, 11, 25,
	5, 25, 348, 10, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 354, 10, 25, 3,
	25, 5, 25, 357, 10, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 372, 10, 28, 12, 28, 14,
	28, 375, 11, 28, 5, 28, 377, 10, 28, 3, 28, 3, 28, 5, 28, 381, 10, 28,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 5, 29, 397, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 421, 10, 33,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 429, 10, 34, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 38, 3, 38, 6, 38, 457, 10, 38, 13, 38, 14, 38, 458, 5, 38,
	461, 10, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 7, 39, 468, 10, 39, 12,
	39, 14, 39, 471, 11, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 6, 39, 478,
	10, 39, 13, 39, 14, 39, 479, 5, 39, 482, 10, 39, 3, 40, 3, 40, 5, 40, 486,
	10, 40, 3, 40, 3, 40, 5, 40, 490, 10, 40, 3, 40, 3, 40, 5, 40, 494, 10,
	40, 5, 40, 496, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 510, 10, 41, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 522, 10, 42,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 533,
	10, 43, 3, 44, 3, 44, 5, 44, 537, 10, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3,
	47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 5, 48, 549, 10, 48, 3, 49, 3, 49,
	3, 50, 3, 50, 3, 51, 3, 51, 5, 51, 557, 10, 51, 3, 52, 3, 52, 3, 52, 3,
	52, 3, 52, 3, 52, 5, 52, 565, 10, 52, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54,
	3, 54, 3, 54, 3, 54, 7, 54, 575, 10, 54, 12, 54, 14, 54, 578, 11, 54, 5,
	54, 580, 10, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 5, 56,
	589, 10, 56, 3, 57, 5, 57, 592, 10, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3,
	58, 5, 58, 599, 10, 58, 3, 59, 5, 59, 602, 10, 59, 3, 59, 3, 59, 3, 59,
	5, 59, 607, 10, 59, 3, 60, 5, 60, 610, 10, 60, 3, 60, 3, 60, 7, 60, 614,
	10, 60, 12, 60, 14, 60, 617, 11, 60, 3, 60, 5, 60, 620, 10, 60, 3, 60,
	2, 2, 61, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
	106, 108, 110, 112, 114, 116, 118, 2, 8, 4, 2, 44, 44, 65, 65, 3, 2, 16,
	26, 3, 2, 61, 64, 3, 2, 12, 15, 3, 2, 5, 6, 3, 2, 65, 66, 2, 673, 2, 120,
	3, 2, 2, 2, 4, 131, 3, 2, 2, 2, 6, 135, 3, 2, 2, 2, 8, 154, 3, 2, 2, 2,
	10, 156, 3, 2, 2, 2, 12, 169, 3, 2, 2, 2, 14, 174, 3, 2, 2, 2, 16, 176,
	3, 2, 2, 2, 18, 197, 3, 2, 2, 2, 20, 205, 3, 2, 2, 2, 22, 223, 3, 2, 2,
	2, 24, 237, 3, 2, 2, 2, 26, 255, 3, 2, 2, 2, 28, 263, 3, 2, 2, 2, 30, 266,
	3, 2, 2, 2, 32, 274, 3, 2, 2, 2, 34, 276, 3, 2, 2, 2, 36, 281, 3, 2, 2,
	2, 38, 315, 3, 2, 2, 2, 40, 317, 3, 2, 2, 2, 42, 328, 3, 2, 2, 2, 44, 330,
	3, 2, 2, 2, 46, 333, 3, 2, 2, 2, 48, 337, 3, 2, 2, 2, 50, 358, 3, 2, 2,
	2, 52, 361, 3, 2, 2, 2, 54, 366, 3, 2, 2, 2, 56, 396, 3, 2, 2, 2, 58, 398,
	3, 2, 2, 2, 60, 402, 3, 2, 2, 2, 62, 407, 3, 2, 2, 2, 64, 420, 3, 2, 2,
	2, 66, 428, 3, 2, 2, 2, 68, 430, 3, 2, 2, 2, 70, 436, 3, 2, 2, 2, 72, 442,
	3, 2, 2, 2, 74, 460, 3, 2, 2, 2, 76, 481, 3, 2, 2, 2, 78, 495, 3, 2, 2,
	2, 80, 509, 3, 2, 2, 2, 82, 521, 3, 2, 2, 2, 84, 532, 3, 2, 2, 2, 86, 536,
	3, 2, 2, 2, 88, 538, 3, 2, 2, 2, 90, 540, 3, 2, 2, 2, 92, 542, 3, 2, 2,
	2, 94, 548, 3, 2, 2, 2, 96, 550, 3, 2, 2, 2, 98, 552, 3, 2, 2, 2, 100,
	556, 3, 2, 2, 2, 102, 564, 3, 2, 2, 2, 104, 566, 3, 2, 2, 2, 106, 569,
	3, 2, 2, 2, 108, 583, 3, 2, 2, 2, 110, 588, 3, 2, 2, 2, 112, 591, 3, 2,
	2, 2, 114, 598, 3, 2, 2, 2, 116, 601, 3, 2, 2, 2, 118, 609, 3, 2, 2, 2,
	120, 121, 5, 4, 3, 2, 121, 123, 5, 6, 4, 2, 122, 124, 7, 2, 2, 3, 123,
	122, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 3, 3, 2, 2, 2, 125, 127, 7,
	44, 2, 2, 126, 125, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 3, 2, 2,
	2, 128, 130, 5, 42, 22, 2, 129, 126, 3, 2, 2, 2, 130, 133, 3, 2, 2, 2,
	131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 5, 3, 2, 2, 2, 133, 131,
	3, 2, 2, 2, 134, 136, 7, 44, 2, 2, 135, 134, 3, 2, 2, 2, 135, 136, 3, 2,
	2, 2, 136, 146, 3, 2, 2, 2, 137, 144, 5, 8, 5, 2, 138, 140, 9, 2, 2, 2,
	139, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141,
	142, 3, 2, 2, 2, 142, 145, 3, 2, 2, 2, 143, 145, 7, 2, 2, 3, 144, 139,
	3, 2, 2, 2, 144, 143, 3, 2, 2, 2, 145, 147, 3, 2, 2, 2, 146, 137, 3, 2,
	2, 2, 147, 148, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2,
	149, 7, 3, 2, 2, 2, 150, 155, 5, 16, 9, 2, 151, 155, 5, 22, 12, 2, 152,
	155, 5, 30, 16, 2, 153, 155, 5, 10, 6, 2, 154, 150, 3, 2, 2, 2, 154, 151,
	3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 154, 153, 3, 2, 2, 2, 155, 9, 3, 2, 2,
	2, 156, 161, 5, 12, 7, 2, 157, 158, 7, 65, 2, 2, 158, 160, 5, 12, 7, 2,
	159, 157, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 161,
	162, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 164, 166,
	7, 65, 2, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 11, 3, 2,
	2, 2, 167, 170, 5, 14, 8, 2, 168, 170, 5, 28, 15, 2, 169, 167, 3, 2, 2,
	2, 169, 168, 3, 2, 2, 2, 170, 13, 3, 2, 2, 2, 171, 175, 5, 56, 29, 2, 172,
	175, 5, 106, 54, 2, 173, 175, 5, 100, 51, 2, 174, 171, 3, 2, 2, 2, 174,
	172, 3, 2, 2, 2, 174, 173, 3, 2, 2, 2, 175, 15, 3, 2, 2, 2, 176, 177, 7,
	28, 2, 2, 177, 187, 5, 84, 43, 2, 178, 180, 7, 66, 2, 2, 179, 178, 3, 2,
	2, 2, 179, 180, 3, 2, 2, 2, 180, 185, 3, 2, 2, 2, 181, 186, 7, 29, 2, 2,
	182, 186, 7, 41, 2, 2, 183, 184, 7, 29, 2, 2, 184, 186, 7, 41, 2, 2, 185,
	181, 3, 2, 2, 2, 185, 182, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 186, 188,
	3, 2, 2, 2, 187, 179, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 189, 3, 2,
	2, 2, 189, 192, 5, 6, 4, 2, 190, 193, 5, 18, 10, 2, 191, 193, 5, 20, 11,
	2, 192, 190, 3, 2, 2, 2, 192, 191, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193,
	194, 3, 2, 2, 2, 194, 195, 5, 118, 60, 2, 195, 17, 3, 2, 2, 2, 196, 198,
	7, 44, 2, 2, 197, 196, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 199, 3, 2,
	2, 2, 199, 201, 7, 30, 2, 2, 200, 202, 7, 41, 2, 2, 201, 200, 3, 2, 2,
	2, 201, 202, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 204, 5, 6, 4, 2, 204,
	19, 3, 2, 2, 2, 205, 206, 7, 31, 2, 2, 206, 216, 5, 84, 43, 2, 207, 209,
	7, 66, 2, 2, 208, 207, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 214, 3, 2,
	2, 2, 210, 215, 7, 29, 2, 2, 211, 215, 7, 41, 2, 2, 212, 213, 7, 29, 2,
	2, 213, 215, 7, 41, 2, 2, 214, 210, 3, 2, 2, 2, 214, 211, 3, 2, 2, 2, 214,
	212, 3, 2, 2, 2, 215, 217, 3, 2, 2, 2, 216, 208, 3, 2, 2, 2, 216, 217,
	3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 221, 5, 6, 4, 2, 219, 222, 5, 18,
	10, 2, 220, 222, 5, 20, 11, 2, 221, 219, 3, 2, 2, 2, 221, 220, 3, 2, 2,
	2, 221, 222, 3, 2, 2, 2, 222, 21, 3, 2, 2, 2, 223, 224, 7, 28, 2, 2, 224,
	228, 5, 110, 56, 2, 225, 227, 5, 24, 13, 2, 226, 225, 3, 2, 2, 2, 227,
	230, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 232,
	3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 231, 233, 5, 26, 14, 2, 232, 231, 3,
	2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 235, 5, 118,
	60, 2, 235, 23, 3, 2, 2, 2, 236, 238, 7, 44, 2, 2, 237, 236, 3, 2, 2, 2,
	237, 238, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 240, 7, 32, 2, 2, 240,
	250, 5, 110, 56, 2, 241, 243, 7, 66, 2, 2, 242, 241, 3, 2, 2, 2, 242, 243,
	3, 2, 2, 2, 243, 248, 3, 2, 2, 2, 244, 249, 7, 29, 2, 2, 245, 249, 7, 41,
	2, 2, 246, 247, 7, 29, 2, 2, 247, 249, 7, 41, 2, 2, 248, 244, 3, 2, 2,
	2, 248, 245, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 249, 251, 3, 2, 2, 2, 250,
	242, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253,
	5, 6, 4, 2, 253, 25, 3, 2, 2, 2, 254, 256, 7, 44, 2, 2, 255, 254, 3, 2,
	2, 2, 255, 256, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 259, 7, 33, 2, 2,
	258, 260, 7, 41, 2, 2, 259, 258, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260,
	261, 3, 2, 2, 2, 261, 262, 5, 6, 4, 2, 262, 27, 3, 2, 2, 2, 263, 264, 5,
	32, 17, 2, 264, 265, 5, 10, 6, 2, 265, 29, 3, 2, 2, 2, 266, 267, 5, 32,
	17, 2, 267, 268, 5, 6, 4, 2, 268, 269, 5, 118, 60, 2, 269, 31, 3, 2, 2,
	2, 270, 275, 5, 34, 18, 2, 271, 275, 5, 36, 19, 2, 272, 275, 5, 38, 20,
	2, 273, 275, 5, 40, 21, 2, 274, 270, 3, 2, 2, 2, 274, 271, 3, 2, 2, 2,
	274, 272, 3, 2, 2, 2, 274, 273, 3, 2, 2, 2, 275, 33, 3, 2, 2, 2, 276, 277,
	7, 34, 2, 2, 277, 279, 5, 84, 43, 2, 278, 280, 7, 41, 2, 2, 279, 278, 3,
	2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 35, 3, 2, 2, 2, 281, 282, 7, 35, 2,
	2, 282, 284, 5, 84, 43, 2, 283, 285, 7, 41, 2, 2, 284, 283, 3, 2, 2, 2,
	284, 285, 3, 2, 2, 2, 285, 37, 3, 2, 2, 2, 286, 287, 7, 36, 2, 2, 287,
	289, 5, 84, 43, 2, 288, 290, 7, 41, 2, 2, 289, 288, 3, 2, 2, 2, 289, 290,
	3, 2, 2, 2, 290, 316, 3, 2, 2, 2, 291, 292, 7, 36, 2, 2, 292, 293, 5, 14,
	8, 2, 293, 294, 7, 66, 2, 2, 294, 295, 5, 84, 43, 2, 295, 296, 7, 66, 2,
	2, 296, 298, 5, 14, 8, 2, 297, 299, 7, 41, 2, 2, 298, 297, 3, 2, 2, 2,
	298, 299, 3, 2, 2, 2, 299, 316, 3, 2, 2, 2, 300, 301, 7, 36, 2, 2, 301,
	302, 5, 112, 57, 2, 302, 304, 7, 37, 2, 2, 303, 305, 7, 38, 2, 2, 304,
	303, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 307, 3, 2, 2, 2, 306, 308,
	7, 39, 2, 2, 307, 306, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 309, 3, 2,
	2, 2, 309, 310, 5, 80, 41, 2, 310, 311, 7, 40, 2, 2, 311, 313, 5, 80, 41,
	2, 312, 314, 7, 41, 2, 2, 313, 312, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314,
	316, 3, 2, 2, 2, 315, 286, 3, 2, 2, 2, 315, 291, 3, 2, 2, 2, 315, 300,
	3, 2, 2, 2, 316, 39, 3, 2, 2, 2, 317, 318, 7, 39, 2, 2, 318, 319, 5, 80,
	41, 2, 319, 320, 7, 40, 2, 2, 320, 322, 5, 80, 41, 2, 321, 323, 7, 41,
	2, 2, 322, 321, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 41, 3, 2, 2, 2,
	324, 329, 5, 44, 23, 2, 325, 329, 5, 46, 24, 2, 326, 329, 5, 50, 26, 2,
	327, 329, 5, 52, 27, 2, 328, 324, 3, 2, 2, 2, 328, 325, 3, 2, 2, 2, 328,
	326, 3, 2, 2, 2, 328, 327, 3, 2, 2, 2, 329, 43, 3, 2, 2, 2, 330, 331, 5,
	48, 25, 2, 331, 332, 5, 10, 6, 2, 332, 45, 3, 2, 2, 2, 333, 334, 5, 48,
	25, 2, 334, 335, 5, 6, 4, 2, 335, 336, 5, 118, 60, 2, 336, 47, 3, 2, 2,
	2, 337, 338, 7, 53, 2, 2, 338, 347, 7, 55, 2, 2, 339, 344, 5, 112, 57,
	2, 340, 341, 7, 66, 2, 2, 341, 343, 5, 112, 57, 2, 342, 340, 3, 2, 2, 2,
	343, 346, 3, 2, 2, 2, 344, 342, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345,
	348, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 347, 339, 3, 2, 2, 2, 347, 348,
	3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 350, 7, 56, 2, 2, 350, 351, 7, 11,
	2, 2, 351, 356, 5, 110, 56, 2, 352, 354, 7, 66, 2, 2, 353, 352, 3, 2, 2,
	2, 353, 354, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 357, 7, 52, 2, 2, 356,
	353, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 49, 3, 2, 2, 2, 358, 359, 5,
	54, 28, 2, 359, 360, 5, 10, 6, 2, 360, 51, 3, 2, 2, 2, 361, 362, 5, 54,
	28, 2, 362, 363, 7, 44, 2, 2, 363, 364, 5, 6, 4, 2, 364, 365, 5, 118, 60,
	2, 365, 53, 3, 2, 2, 2, 366, 367, 7, 53, 2, 2, 367, 376, 7, 55, 2, 2, 368,
	373, 5, 112, 57, 2, 369, 370, 7, 66, 2, 2, 370, 372, 5, 112, 57, 2, 371,
	369, 3, 2, 2, 2, 372, 375, 3, 2, 2, 2, 373, 371, 3, 2, 2, 2, 373, 374,
	3, 2, 2, 2, 374, 377, 3, 2, 2, 2, 375, 373, 3, 2, 2, 2, 376, 368, 3, 2,
	2, 2, 376, 377, 3, 2, 2, 2, 377, 378, 3, 2, 2, 2, 378, 380, 7, 56, 2, 2,
	379, 381, 7, 11, 2, 2, 380, 379, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381,
	55, 3, 2, 2, 2, 382, 397, 5, 58, 30, 2, 383, 397, 5, 62, 32, 2, 384, 397,
	5, 74, 38, 2, 385, 386, 7, 53, 2, 2, 386, 387, 7, 11, 2, 2, 387, 397, 5,
	112, 57, 2, 388, 389, 7, 53, 2, 2, 389, 390, 7, 11, 2, 2, 390, 397, 5,
	110, 56, 2, 391, 392, 7, 53, 2, 2, 392, 393, 5, 98, 50, 2, 393, 394, 5,
	110, 56, 2, 394, 397, 3, 2, 2, 2, 395, 397, 5, 78, 40, 2, 396, 382, 3,
	2, 2, 2, 396, 383, 3, 2, 2, 2, 396, 384, 3, 2, 2, 2, 396, 385, 3, 2, 2,
	2, 396, 388, 3, 2, 2, 2, 396, 391, 3, 2, 2, 2, 396, 395, 3, 2, 2, 2, 397,
	57, 3, 2, 2, 2, 398, 399, 7, 53, 2, 2, 399, 400, 7, 11, 2, 2, 400, 401,
	5, 60, 31, 2, 401, 59, 3, 2, 2, 2, 402, 403, 7, 50, 2, 2, 403, 404, 7,
	55, 2, 2, 404, 405, 5, 110, 56, 2, 405, 406, 7, 56, 2, 2, 406, 61, 3, 2,
	2, 2, 407, 408, 7, 53, 2, 2, 408, 409, 7, 11, 2, 2, 409, 410, 5, 64, 33,
	2, 410, 63, 3, 2, 2, 2, 411, 412, 5, 110, 56, 2, 412, 413, 5, 66, 34, 2,
	413, 414, 5, 110, 56, 2, 414, 421, 3, 2, 2, 2, 415, 416, 7, 51, 2, 2, 416,
	417, 7, 55, 2, 2, 417, 418, 5, 110, 56, 2, 418, 419, 7, 56, 2, 2, 419,
	421, 3, 2, 2, 2, 420, 411, 3, 2, 2, 2, 420, 415, 3, 2, 2, 2, 421, 65, 3,
	2, 2, 2, 422, 429, 7, 8, 2, 2, 423, 429, 5, 68, 35, 2, 424, 429, 7, 9,
	2, 2, 425, 429, 5, 70, 36, 2, 426, 429, 7, 10, 2, 2, 427, 429, 5, 72, 37,
	2, 428, 422, 3, 2, 2, 2, 428, 423, 3, 2, 2, 2, 428, 424, 3, 2, 2, 2, 428,
	425, 3, 2, 2, 2, 428, 426, 3, 2, 2, 2, 428, 427, 3, 2, 2, 2, 429, 67, 3,
	2, 2, 2, 430, 431, 7, 62, 2, 2, 431, 432, 7, 59, 2, 2, 432, 433, 5, 116,
	59, 2, 433, 434, 7, 60, 2, 2, 434, 435, 7, 8, 2, 2, 435, 69, 3, 2, 2, 2,
	436, 437, 7, 9, 2, 2, 437, 438, 7, 59, 2, 2, 438, 439, 5, 116, 59, 2, 439,
	440, 7, 60, 2, 2, 440, 441, 7, 62, 2, 2, 441, 71, 3, 2, 2, 2, 442, 443,
	7, 62, 2, 2, 443, 444, 7, 59, 2, 2, 444, 445, 5, 116, 59, 2, 445, 446,
	7, 60, 2, 2, 446, 447, 7, 62, 2, 2, 447, 73, 3, 2, 2, 2, 448, 449, 7, 53,
	2, 2, 449, 450, 7, 11, 2, 2, 450, 461, 5, 76, 39, 2, 451, 452, 7, 53, 2,
	2, 452, 453, 5, 98, 50, 2, 453, 456, 5, 110, 56, 2, 454, 455, 7, 66, 2,
	2, 455, 457, 5, 110, 56, 2, 456, 454, 3, 2, 2, 2, 457, 458, 3, 2, 2, 2,
	458, 456, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 461, 3, 2, 2, 2, 460,
	448, 3, 2, 2, 2, 460, 451, 3, 2, 2, 2, 461, 75, 3, 2, 2, 2, 462, 463, 7,
	49, 2, 2, 463, 464, 7, 55, 2, 2, 464, 469, 5, 110, 56, 2, 465, 466, 7,
	66, 2, 2, 466, 468, 5, 110, 56, 2, 467, 465, 3, 2, 2, 2, 468, 471, 3, 2,
	2, 2, 469, 467, 3, 2, 2, 2, 469, 470, 3, 2, 2, 2, 470, 472, 3, 2, 2, 2,
	471, 469, 3, 2, 2, 2, 472, 473, 7, 56, 2, 2, 473, 482, 3, 2, 2, 2, 474,
	477, 5, 110, 56, 2, 475, 476, 7, 66, 2, 2, 476, 478, 5, 110, 56, 2, 477,
	475, 3, 2, 2, 2, 478, 479, 3, 2, 2, 2, 479, 477, 3, 2, 2, 2, 479, 480,
	3, 2, 2, 2, 480, 482, 3, 2, 2, 2, 481, 462, 3, 2, 2, 2, 481, 474, 3, 2,
	2, 2, 482, 77, 3, 2, 2, 2, 483, 485, 5, 58, 30, 2, 484, 486, 5, 108, 55,
	2, 485, 484, 3, 2, 2, 2, 485, 486, 3, 2, 2, 2, 486, 496, 3, 2, 2, 2, 487,
	489, 5, 62, 32, 2, 488, 490, 5, 108, 55, 2, 489, 488, 3, 2, 2, 2, 489,
	490, 3, 2, 2, 2, 490, 496, 3, 2, 2, 2, 491, 493, 5, 74, 38, 2, 492, 494,
	5, 108, 55, 2, 493, 492, 3, 2, 2, 2, 493, 494, 3, 2, 2, 2, 494, 496, 3,
	2, 2, 2, 495, 483, 3, 2, 2, 2, 495, 487, 3, 2, 2, 2, 495, 491, 3, 2, 2,
	2, 496, 79, 3, 2, 2, 2, 497, 510, 5, 112, 57, 2, 498, 510, 5, 82, 42, 2,
	499, 510, 5, 114, 58, 2, 500, 510, 5, 60, 31, 2, 501, 502, 7, 55, 2, 2,
	502, 503, 5, 64, 33, 2, 503, 504, 7, 56, 2, 2, 504, 510, 3, 2, 2, 2, 505,
	506, 7, 55, 2, 2, 506, 507, 5, 76, 39, 2, 507, 508, 7, 56, 2, 2, 508, 510,
	3, 2, 2, 2, 509, 497, 3, 2, 2, 2, 509, 498, 3, 2, 2, 2, 509, 499, 3, 2,
	2, 2, 509, 500, 3, 2, 2, 2, 509, 501, 3, 2, 2, 2, 509, 505, 3, 2, 2, 2,
	510, 81, 3, 2, 2, 2, 511, 522, 5, 106, 54, 2, 512, 522, 5, 114, 58, 2,
	513, 514, 7, 55, 2, 2, 514, 515, 5, 84, 43, 2, 515, 516, 7, 56, 2, 2, 516,
	522, 3, 2, 2, 2, 517, 518, 7, 55, 2, 2, 518, 519, 5, 92, 47, 2, 519, 520,
	7, 56, 2, 2, 520, 522, 3, 2, 2, 2, 521, 511, 3, 2, 2, 2, 521, 512, 3, 2,
	2, 2, 521, 513, 3, 2, 2, 2, 521, 517, 3, 2, 2, 2, 522, 83, 3, 2, 2, 2,
	523, 524, 5, 80, 41, 2, 524, 525, 5, 88, 45, 2, 525, 526, 5, 86, 44, 2,
	526, 533, 3, 2, 2, 2, 527, 528, 5, 90, 46, 2, 528, 529, 5, 86, 44, 2, 529,
	533, 3, 2, 2, 2, 530, 533, 5, 92, 47, 2, 531, 533, 5, 80, 41, 2, 532, 523,
	3, 2, 2, 2, 532, 527, 3, 2, 2, 2, 532, 530, 3, 2, 2, 2, 532, 531, 3, 2,
	2, 2, 533, 85, 3, 2, 2, 2, 534, 537, 5, 80, 41, 2, 535, 537, 5, 84, 43,
	2, 536, 534, 3, 2, 2, 2, 536, 535, 3, 2, 2, 2, 537, 87, 3, 2, 2, 2, 538,
	539, 9, 3, 2, 2, 539, 89, 3, 2, 2, 2, 540, 541, 7, 27, 2, 2, 541, 91, 3,
	2, 2, 2, 542, 543, 5, 80, 41, 2, 543, 544, 5, 96, 49, 2, 544, 545, 5, 94,
	48, 2, 545, 93, 3, 2, 2, 2, 546, 549, 5, 80, 41, 2, 547, 549, 5, 92, 47,
	2, 548, 546, 3, 2, 2, 2, 548, 547, 3, 2, 2, 2, 549, 95, 3, 2, 2, 2, 550,
	551, 9, 4, 2, 2, 551, 97, 3, 2, 2, 2, 552, 553, 9, 5, 2, 2, 553, 99, 3,
	2, 2, 2, 554, 557, 5, 102, 52, 2, 555, 557, 5, 104, 53, 2, 556, 554, 3,
	2, 2, 2, 556, 555, 3, 2, 2, 2, 557, 101, 3, 2, 2, 2, 558, 559, 7, 47, 2,
	2, 559, 565, 5, 112, 57, 2, 560, 561, 7, 47, 2, 2, 561, 565, 5, 110, 56,
	2, 562, 563, 7, 47, 2, 2, 563, 565, 5, 106, 54, 2, 564, 558, 3, 2, 2, 2,
	564, 560, 3, 2, 2, 2, 564, 562, 3, 2, 2, 2, 565, 103, 3, 2, 2, 2, 566,
	567, 7, 48, 2, 2, 567, 568, 7, 53, 2, 2, 568, 105, 3, 2, 2, 2, 569, 570,
	7, 53, 2, 2, 570, 579, 7, 55, 2, 2, 571, 576, 5, 110, 56, 2, 572, 573,
	7, 66, 2, 2, 573, 575, 5, 110, 56, 2, 574, 572, 3, 2, 2, 2, 575, 578, 3,
	2, 2, 2, 576, 574, 3, 2, 2, 2, 576, 577, 3, 2, 2, 2, 577, 580, 3, 2, 2,
	2, 578, 576, 3, 2, 2, 2, 579, 571, 3, 2, 2, 2, 579, 580, 3, 2, 2, 2, 580,
	581, 3, 2, 2, 2, 581, 582, 7, 56, 2, 2, 582, 107, 3, 2, 2, 2, 583, 584,
	9, 6, 2, 2, 584, 109, 3, 2, 2, 2, 585, 589, 5, 92, 47, 2, 586, 589, 5,
	84, 43, 2, 587, 589, 5, 80, 41, 2, 588, 585, 3, 2, 2, 2, 588, 586, 3, 2,
	2, 2, 588, 587, 3, 2, 2, 2, 589, 111, 3, 2, 2, 2, 590, 592, 7, 62, 2, 2,
	591, 590, 3, 2, 2, 2, 591, 592, 3, 2, 2, 2, 592, 593, 3, 2, 2, 2, 593,
	594, 7, 53, 2, 2, 594, 113, 3, 2, 2, 2, 595, 599, 5, 116, 59, 2, 596, 599,
	7, 4, 2, 2, 597, 599, 7, 7, 2, 2, 598, 595, 3, 2, 2, 2, 598, 596, 3, 2,
	2, 2, 598, 597, 3, 2, 2, 2, 599, 115, 3, 2, 2, 2, 600, 602, 7, 62, 2, 2,
	601, 600, 3, 2, 2, 2, 601, 602, 3, 2, 2, 2, 602, 603, 3, 2, 2, 2, 603,
	606, 7, 3, 2, 2, 604, 605, 9, 7, 2, 2, 605, 607, 7, 3, 2, 2, 606, 604,
	3, 2, 2, 2, 606, 607, 3, 2, 2, 2, 607, 117, 3, 2, 2, 2, 608, 610, 7, 44,
	2, 2, 609, 608, 3, 2, 2, 2, 609, 610, 3, 2, 2, 2, 610, 611, 3, 2, 2, 2,
	611, 615, 7, 46, 2, 2, 612, 614, 7, 65, 2, 2, 613, 612, 3, 2, 2, 2, 614,
	617, 3, 2, 2, 2, 615, 613, 3, 2, 2, 2, 615, 616, 3, 2, 2, 2, 616, 619,
	3, 2, 2, 2, 617, 615, 3, 2, 2, 2, 618, 620, 7, 44, 2, 2, 619, 618, 3, 2,
	2, 2, 619, 620, 3, 2, 2, 2, 620, 119, 3, 2, 2, 2, 79, 123, 126, 131, 135,
	141, 144, 148, 154, 161, 165, 169, 174, 179, 185, 187, 192, 197, 201, 208,
	214, 216, 221, 228, 232, 237, 242, 248, 250, 255, 259, 274, 279, 284, 289,
	298, 304, 307, 313, 315, 322, 328, 344, 347, 353, 356, 373, 376, 380, 396,
	420, 428, 458, 460, 469, 479, 481, 485, 489, 493, 495, 509, 521, 532, 536,
	548, 556, 564, 576, 579, 588, 591, 598, 601, 606, 609, 615, 619,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "'->'", "'<-'", "'--'", "'='", "'+='", "'-='",
	"'*='", "'/='", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "'+'",
	"'-'", "'*'", "'/'", "'.'", "','",
}
var symbolicNames = []string{
	"", "INT", "STRING", "LABEL", "ML_LABEL", "BOOL", "OR_EDGE_LR", "OR_EDGE_RL",
	"UNOR_EDGE", "ASSIGN", "ADD_ASSIGN", "SUB_ASSIGN", "MULT_ASSIGN", "DIV_ASSIGN",
	"NEQ", "EQUALS", "LESS_THAN", "GR_THAN", "LESS_THAN_E", "GR_THAN_E", "AND",
	"OR", "XOR", "NOR", "NAND", "NOT", "IF", "THEN", "ELSE", "ELIF", "IS",
	"DEFAULT", "WHILE", "UNTIL", "FOR", "IN", "RANGE", "FROM", "TO", "DO",
	"SKIP_ITERATION", "BREAK", "NEWLINE", "BLOCK_BEGIN", "BLOCK_END", "PRINTER",
	"KEY_INPUT", "G_N", "V_N", "E_N", "WHERE", "ID", "WS", "L_PAREN", "R_PAREN",
	"L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ADD", "SUB", "MULT", "DIV",
	"ACT_DELIM", "ARG_DELIM", "LINE_COMMENT", "M_LINE_COMMENT",
}

var ruleNames = []string{
	"file", "functions_block", "sequence", "sequence_element", "sequence_line",
	"one_line_sequence_element", "atom_action", "if_stmnt", "else_stmnt", "elif_stmnt",
	"if_is_stmnt", "case_stmnt", "default_stmnt", "one_line_stmnt", "mult_line_stmnt",
	"cycle_stmnt", "while_stmnt", "until_stmnt", "for_stmnt", "from_to_stmnt",
	"function_declaration", "one_line_function_declaration", "mult_line_function_declaration",
	"function_declaration_head", "one_line_procedure_declaration", "mult_line_procedure_declaration",
	"procedure_declaration_head", "var_assign", "vertice_assign", "vertice_type",
	"edge_assign", "edge_type", "edge", "or_w_edge_lr", "or_w_edge_rl", "unor_w_edge",
	"graph_assign", "graph_type", "labeled_assign", "expr", "integral_expr",
	"logical_expr", "log_expr_operand", "bin_log_operator", "unar_log_operator",
	"arithm_expr", "arithm_expr_operand", "bin_arithm_operator", "arithm_assign_operator",
	"builtin_function_call", "built_func_print", "built_func_input", "function_call",
	"label", "value", "variable", "builtin_type", "number_type", "block_end",
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
	GraffleParserINT            = 1
	GraffleParserSTRING         = 2
	GraffleParserLABEL          = 3
	GraffleParserML_LABEL       = 4
	GraffleParserBOOL           = 5
	GraffleParserOR_EDGE_LR     = 6
	GraffleParserOR_EDGE_RL     = 7
	GraffleParserUNOR_EDGE      = 8
	GraffleParserASSIGN         = 9
	GraffleParserADD_ASSIGN     = 10
	GraffleParserSUB_ASSIGN     = 11
	GraffleParserMULT_ASSIGN    = 12
	GraffleParserDIV_ASSIGN     = 13
	GraffleParserNEQ            = 14
	GraffleParserEQUALS         = 15
	GraffleParserLESS_THAN      = 16
	GraffleParserGR_THAN        = 17
	GraffleParserLESS_THAN_E    = 18
	GraffleParserGR_THAN_E      = 19
	GraffleParserAND            = 20
	GraffleParserOR             = 21
	GraffleParserXOR            = 22
	GraffleParserNOR            = 23
	GraffleParserNAND           = 24
	GraffleParserNOT            = 25
	GraffleParserIF             = 26
	GraffleParserTHEN           = 27
	GraffleParserELSE           = 28
	GraffleParserELIF           = 29
	GraffleParserIS             = 30
	GraffleParserDEFAULT        = 31
	GraffleParserWHILE          = 32
	GraffleParserUNTIL          = 33
	GraffleParserFOR            = 34
	GraffleParserIN             = 35
	GraffleParserRANGE          = 36
	GraffleParserFROM           = 37
	GraffleParserTO             = 38
	GraffleParserDO             = 39
	GraffleParserSKIP_ITERATION = 40
	GraffleParserBREAK          = 41
	GraffleParserNEWLINE        = 42
	GraffleParserBLOCK_BEGIN    = 43
	GraffleParserBLOCK_END      = 44
	GraffleParserPRINTER        = 45
	GraffleParserKEY_INPUT      = 46
	GraffleParserG_N            = 47
	GraffleParserV_N            = 48
	GraffleParserE_N            = 49
	GraffleParserWHERE          = 50
	GraffleParserID             = 51
	GraffleParserWS             = 52
	GraffleParserL_PAREN        = 53
	GraffleParserR_PAREN        = 54
	GraffleParserL_CURLY        = 55
	GraffleParserR_CURLY        = 56
	GraffleParserL_BRACKET      = 57
	GraffleParserR_BRACKET      = 58
	GraffleParserADD            = 59
	GraffleParserSUB            = 60
	GraffleParserMULT           = 61
	GraffleParserDIV            = 62
	GraffleParserACT_DELIM      = 63
	GraffleParserARG_DELIM      = 64
	GraffleParserLINE_COMMENT   = 65
	GraffleParserM_LINE_COMMENT = 66
)

// GraffleParser rules.
const (
	GraffleParserRULE_file                            = 0
	GraffleParserRULE_functions_block                 = 1
	GraffleParserRULE_sequence                        = 2
	GraffleParserRULE_sequence_element                = 3
	GraffleParserRULE_sequence_line                   = 4
	GraffleParserRULE_one_line_sequence_element       = 5
	GraffleParserRULE_atom_action                     = 6
	GraffleParserRULE_if_stmnt                        = 7
	GraffleParserRULE_else_stmnt                      = 8
	GraffleParserRULE_elif_stmnt                      = 9
	GraffleParserRULE_if_is_stmnt                     = 10
	GraffleParserRULE_case_stmnt                      = 11
	GraffleParserRULE_default_stmnt                   = 12
	GraffleParserRULE_one_line_stmnt                  = 13
	GraffleParserRULE_mult_line_stmnt                 = 14
	GraffleParserRULE_cycle_stmnt                     = 15
	GraffleParserRULE_while_stmnt                     = 16
	GraffleParserRULE_until_stmnt                     = 17
	GraffleParserRULE_for_stmnt                       = 18
	GraffleParserRULE_from_to_stmnt                   = 19
	GraffleParserRULE_function_declaration            = 20
	GraffleParserRULE_one_line_function_declaration   = 21
	GraffleParserRULE_mult_line_function_declaration  = 22
	GraffleParserRULE_function_declaration_head       = 23
	GraffleParserRULE_one_line_procedure_declaration  = 24
	GraffleParserRULE_mult_line_procedure_declaration = 25
	GraffleParserRULE_procedure_declaration_head      = 26
	GraffleParserRULE_var_assign                      = 27
	GraffleParserRULE_vertice_assign                  = 28
	GraffleParserRULE_vertice_type                    = 29
	GraffleParserRULE_edge_assign                     = 30
	GraffleParserRULE_edge_type                       = 31
	GraffleParserRULE_edge                            = 32
	GraffleParserRULE_or_w_edge_lr                    = 33
	GraffleParserRULE_or_w_edge_rl                    = 34
	GraffleParserRULE_unor_w_edge                     = 35
	GraffleParserRULE_graph_assign                    = 36
	GraffleParserRULE_graph_type                      = 37
	GraffleParserRULE_labeled_assign                  = 38
	GraffleParserRULE_expr                            = 39
	GraffleParserRULE_integral_expr                   = 40
	GraffleParserRULE_logical_expr                    = 41
	GraffleParserRULE_log_expr_operand                = 42
	GraffleParserRULE_bin_log_operator                = 43
	GraffleParserRULE_unar_log_operator               = 44
	GraffleParserRULE_arithm_expr                     = 45
	GraffleParserRULE_arithm_expr_operand             = 46
	GraffleParserRULE_bin_arithm_operator             = 47
	GraffleParserRULE_arithm_assign_operator          = 48
	GraffleParserRULE_builtin_function_call           = 49
	GraffleParserRULE_built_func_print                = 50
	GraffleParserRULE_built_func_input                = 51
	GraffleParserRULE_function_call                   = 52
	GraffleParserRULE_label                           = 53
	GraffleParserRULE_value                           = 54
	GraffleParserRULE_variable                        = 55
	GraffleParserRULE_builtin_type                    = 56
	GraffleParserRULE_number_type                     = 57
	GraffleParserRULE_block_end                       = 58
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *FileContext) Functions_block() IFunctions_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctions_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctions_blockContext)
}

func (s *FileContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Functions_block()
	}
	{
		p.SetState(119)
		p.Sequence()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(120)
			p.Match(GraffleParserEOF)
		}

	}

	return localctx
}

// IFunctions_blockContext is an interface to support dynamic dispatch.
type IFunctions_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctions_blockContext differentiates from other interfaces.
	IsFunctions_blockContext()
}

type Functions_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctions_blockContext() *Functions_blockContext {
	var p = new(Functions_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_functions_block
	return p
}

func (*Functions_blockContext) IsFunctions_blockContext() {}

func NewFunctions_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Functions_blockContext {
	var p = new(Functions_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_functions_block

	return p
}

func (s *Functions_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Functions_blockContext) AllFunction_declaration() []IFunction_declarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunction_declarationContext)(nil)).Elem())
	var tst = make([]IFunction_declarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunction_declarationContext)
		}
	}

	return tst
}

func (s *Functions_blockContext) Function_declaration(i int) IFunction_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_declarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunction_declarationContext)
}

func (s *Functions_blockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserNEWLINE)
}

func (s *Functions_blockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, i)
}

func (s *Functions_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Functions_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Functions_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterFunctions_block(s)
	}
}

func (s *Functions_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitFunctions_block(s)
	}
}

func (s *Functions_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitFunctions_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Functions_block() (localctx IFunctions_blockContext) {
	localctx = NewFunctions_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraffleParserRULE_functions_block)
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
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == GraffleParserNEWLINE {
				{
					p.SetState(123)
					p.Match(GraffleParserNEWLINE)
				}

			}
			{
				p.SetState(126)
				p.Function_declaration()
			}

		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 4, GraffleParserRULE_sequence)
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
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserNEWLINE {
		{
			p.SetState(132)
			p.Match(GraffleParserNEWLINE)
		}

	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(GraffleParserIF-26))|(1<<(GraffleParserWHILE-26))|(1<<(GraffleParserUNTIL-26))|(1<<(GraffleParserFOR-26))|(1<<(GraffleParserFROM-26))|(1<<(GraffleParserPRINTER-26))|(1<<(GraffleParserKEY_INPUT-26))|(1<<(GraffleParserID-26)))) != 0) {
		{
			p.SetState(135)
			p.Sequence_element()
		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GraffleParserNEWLINE, GraffleParserACT_DELIM:
			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(136)
						_la = p.GetTokenStream().LA(1)

						if !(_la == GraffleParserNEWLINE || _la == GraffleParserACT_DELIM) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(139)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
			}

		case GraffleParserEOF:
			{
				p.SetState(141)
				p.Match(GraffleParserEOF)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 6, GraffleParserRULE_sequence_element)

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

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.If_stmnt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.If_is_stmnt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(150)
			p.Mult_line_stmnt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(151)
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
	p.EnterRule(localctx, 8, GraffleParserRULE_sequence_line)

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
		p.SetState(154)
		p.One_line_sequence_element()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(155)
				p.Match(GraffleParserACT_DELIM)
			}
			{
				p.SetState(156)
				p.One_line_sequence_element()
			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(162)
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
	p.EnterRule(localctx, 10, GraffleParserRULE_one_line_sequence_element)

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

	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserPRINTER, GraffleParserKEY_INPUT, GraffleParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Atom_action()
		}

	case GraffleParserWHILE, GraffleParserUNTIL, GraffleParserFOR, GraffleParserFROM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
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
	p.EnterRule(localctx, 12, GraffleParserRULE_atom_action)

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

	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Var_assign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.Function_call()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(171)
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

	// IsIf_stmntContext differentiates from other interfaces.
	IsIf_stmntContext()
}

type If_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *If_stmntContext) IF() antlr.TerminalNode {
	return s.GetToken(GraffleParserIF, 0)
}

func (s *If_stmntContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *If_stmntContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *If_stmntContext) Block_end() IBlock_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endContext)
}

func (s *If_stmntContext) Else_stmnt() IElse_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElse_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElse_stmntContext)
}

func (s *If_stmntContext) Elif_stmnt() IElif_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElif_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElif_stmntContext)
}

func (s *If_stmntContext) THEN() antlr.TerminalNode {
	return s.GetToken(GraffleParserTHEN, 0)
}

func (s *If_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
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
	p.EnterRule(localctx, 14, GraffleParserRULE_if_stmnt)
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
		p.SetState(174)
		p.Match(GraffleParserIF)
	}
	{
		p.SetState(175)
		p.Logical_expr()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserTHEN || _la == GraffleParserDO || _la == GraffleParserARG_DELIM {
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserARG_DELIM {
			{
				p.SetState(176)
				p.Match(GraffleParserARG_DELIM)
			}

		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(179)
				p.Match(GraffleParserTHEN)
			}

		case 2:
			{
				p.SetState(180)
				p.Match(GraffleParserDO)
			}

		case 3:
			{
				p.SetState(181)
				p.Match(GraffleParserTHEN)
			}
			{
				p.SetState(182)
				p.Match(GraffleParserDO)
			}

		}

	}
	{
		p.SetState(187)
		p.Sequence()
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(188)
			p.Else_stmnt()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(189)
			p.Elif_stmnt()
		}

	}
	{
		p.SetState(192)
		p.Block_end()
	}

	return localctx
}

// IElse_stmntContext is an interface to support dynamic dispatch.
type IElse_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElse_stmntContext differentiates from other interfaces.
	IsElse_stmntContext()
}

type Else_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_stmntContext() *Else_stmntContext {
	var p = new(Else_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_else_stmnt
	return p
}

func (*Else_stmntContext) IsElse_stmntContext() {}

func NewElse_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_stmntContext {
	var p = new(Else_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_else_stmnt

	return p
}

func (s *Else_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_stmntContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GraffleParserELSE, 0)
}

func (s *Else_stmntContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Else_stmntContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *Else_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *Else_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterElse_stmnt(s)
	}
}

func (s *Else_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitElse_stmnt(s)
	}
}

func (s *Else_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitElse_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Else_stmnt() (localctx IElse_stmntContext) {
	localctx = NewElse_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GraffleParserRULE_else_stmnt)
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
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserNEWLINE {
		{
			p.SetState(194)
			p.Match(GraffleParserNEWLINE)
		}

	}
	{
		p.SetState(197)
		p.Match(GraffleParserELSE)
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDO {
		{
			p.SetState(198)
			p.Match(GraffleParserDO)
		}

	}
	{
		p.SetState(201)
		p.Sequence()
	}

	return localctx
}

// IElif_stmntContext is an interface to support dynamic dispatch.
type IElif_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElif_stmntContext differentiates from other interfaces.
	IsElif_stmntContext()
}

type Elif_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElif_stmntContext() *Elif_stmntContext {
	var p = new(Elif_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_elif_stmnt
	return p
}

func (*Elif_stmntContext) IsElif_stmntContext() {}

func NewElif_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Elif_stmntContext {
	var p = new(Elif_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_elif_stmnt

	return p
}

func (s *Elif_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *Elif_stmntContext) ELIF() antlr.TerminalNode {
	return s.GetToken(GraffleParserELIF, 0)
}

func (s *Elif_stmntContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *Elif_stmntContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Elif_stmntContext) Else_stmnt() IElse_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElse_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElse_stmntContext)
}

func (s *Elif_stmntContext) Elif_stmnt() IElif_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElif_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElif_stmntContext)
}

func (s *Elif_stmntContext) THEN() antlr.TerminalNode {
	return s.GetToken(GraffleParserTHEN, 0)
}

func (s *Elif_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *Elif_stmntContext) ARG_DELIM() antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, 0)
}

func (s *Elif_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Elif_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Elif_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterElif_stmnt(s)
	}
}

func (s *Elif_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitElif_stmnt(s)
	}
}

func (s *Elif_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitElif_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Elif_stmnt() (localctx IElif_stmntContext) {
	localctx = NewElif_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GraffleParserRULE_elif_stmnt)
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
		p.SetState(203)
		p.Match(GraffleParserELIF)
	}
	{
		p.SetState(204)
		p.Logical_expr()
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserTHEN || _la == GraffleParserDO || _la == GraffleParserARG_DELIM {
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserARG_DELIM {
			{
				p.SetState(205)
				p.Match(GraffleParserARG_DELIM)
			}

		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(208)
				p.Match(GraffleParserTHEN)
			}

		case 2:
			{
				p.SetState(209)
				p.Match(GraffleParserDO)
			}

		case 3:
			{
				p.SetState(210)
				p.Match(GraffleParserTHEN)
			}
			{
				p.SetState(211)
				p.Match(GraffleParserDO)
			}

		}

	}
	{
		p.SetState(216)
		p.Sequence()
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(217)
			p.Else_stmnt()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(218)
			p.Elif_stmnt()
		}

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

func (s *If_is_stmntContext) Default_stmnt() IDefault_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefault_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefault_stmntContext)
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
	p.EnterRule(localctx, 20, GraffleParserRULE_if_is_stmnt)

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
		p.SetState(221)
		p.Match(GraffleParserIF)
	}
	{
		p.SetState(222)
		p.Value()
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(223)
				p.Case_stmnt()
			}

		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(229)
			p.Default_stmnt()
		}

	}
	{
		p.SetState(232)
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

func (s *Case_stmntContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Case_stmntContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
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
	p.EnterRule(localctx, 22, GraffleParserRULE_case_stmnt)
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
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserNEWLINE {
		{
			p.SetState(234)
			p.Match(GraffleParserNEWLINE)
		}

	}
	{
		p.SetState(237)
		p.Match(GraffleParserIS)
	}
	{
		p.SetState(238)
		p.Value()
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserTHEN || _la == GraffleParserDO || _la == GraffleParserARG_DELIM {
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserARG_DELIM {
			{
				p.SetState(239)
				p.Match(GraffleParserARG_DELIM)
			}

		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(242)
				p.Match(GraffleParserTHEN)
			}

		case 2:
			{
				p.SetState(243)
				p.Match(GraffleParserDO)
			}

		case 3:
			{
				p.SetState(244)
				p.Match(GraffleParserTHEN)
			}
			{
				p.SetState(245)
				p.Match(GraffleParserDO)
			}

		}

	}
	{
		p.SetState(250)
		p.Sequence()
	}

	return localctx
}

// IDefault_stmntContext is an interface to support dynamic dispatch.
type IDefault_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefault_stmntContext differentiates from other interfaces.
	IsDefault_stmntContext()
}

type Default_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_stmntContext() *Default_stmntContext {
	var p = new(Default_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_default_stmnt
	return p
}

func (*Default_stmntContext) IsDefault_stmntContext() {}

func NewDefault_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_stmntContext {
	var p = new(Default_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_default_stmnt

	return p
}

func (s *Default_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_stmntContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(GraffleParserDEFAULT, 0)
}

func (s *Default_stmntContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Default_stmntContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *Default_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *Default_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Default_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterDefault_stmnt(s)
	}
}

func (s *Default_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitDefault_stmnt(s)
	}
}

func (s *Default_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitDefault_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Default_stmnt() (localctx IDefault_stmntContext) {
	localctx = NewDefault_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GraffleParserRULE_default_stmnt)
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
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserNEWLINE {
		{
			p.SetState(252)
			p.Match(GraffleParserNEWLINE)
		}

	}
	{
		p.SetState(255)
		p.Match(GraffleParserDEFAULT)
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
	{
		p.SetState(259)
		p.Sequence()
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

func (s *One_line_stmntContext) Cycle_stmnt() ICycle_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICycle_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICycle_stmntContext)
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
	p.EnterRule(localctx, 26, GraffleParserRULE_one_line_stmnt)

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
		p.SetState(261)
		p.Cycle_stmnt()
	}
	{
		p.SetState(262)
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

func (s *Mult_line_stmntContext) Cycle_stmnt() ICycle_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICycle_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICycle_stmntContext)
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
	p.EnterRule(localctx, 28, GraffleParserRULE_mult_line_stmnt)

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
		p.SetState(264)
		p.Cycle_stmnt()
	}
	{
		p.SetState(265)
		p.Sequence()
	}
	{
		p.SetState(266)
		p.Block_end()
	}

	return localctx
}

// ICycle_stmntContext is an interface to support dynamic dispatch.
type ICycle_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCycle_stmntContext differentiates from other interfaces.
	IsCycle_stmntContext()
}

type Cycle_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCycle_stmntContext() *Cycle_stmntContext {
	var p = new(Cycle_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_cycle_stmnt
	return p
}

func (*Cycle_stmntContext) IsCycle_stmntContext() {}

func NewCycle_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cycle_stmntContext {
	var p = new(Cycle_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_cycle_stmnt

	return p
}

func (s *Cycle_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *Cycle_stmntContext) While_stmnt() IWhile_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_stmntContext)
}

func (s *Cycle_stmntContext) Until_stmnt() IUntil_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUntil_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUntil_stmntContext)
}

func (s *Cycle_stmntContext) For_stmnt() IFor_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_stmntContext)
}

func (s *Cycle_stmntContext) From_to_stmnt() IFrom_to_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFrom_to_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFrom_to_stmntContext)
}

func (s *Cycle_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cycle_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cycle_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterCycle_stmnt(s)
	}
}

func (s *Cycle_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitCycle_stmnt(s)
	}
}

func (s *Cycle_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitCycle_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Cycle_stmnt() (localctx ICycle_stmntContext) {
	localctx = NewCycle_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GraffleParserRULE_cycle_stmnt)

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

	p.SetState(272)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserWHILE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.While_stmnt()
		}

	case GraffleParserUNTIL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(269)
			p.Until_stmnt()
		}

	case GraffleParserFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(270)
			p.For_stmnt()
		}

	case GraffleParserFROM:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(271)
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

	// IsWhile_stmntContext differentiates from other interfaces.
	IsWhile_stmntContext()
}

type While_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 32, GraffleParserRULE_while_stmnt)
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
		p.SetState(274)
		p.Match(GraffleParserWHILE)
	}
	{
		p.SetState(275)
		p.Logical_expr()
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDO {
		{
			p.SetState(276)
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

	// IsUntil_stmntContext differentiates from other interfaces.
	IsUntil_stmntContext()
}

type Until_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 34, GraffleParserRULE_until_stmnt)
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
		p.SetState(279)
		p.Match(GraffleParserUNTIL)
	}
	{
		p.SetState(280)
		p.Logical_expr()
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDO {
		{
			p.SetState(281)
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
}

func NewForLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForLogicalContext {
	var p = new(ForLogicalContext)

	p.For_stmntContext = NewEmptyFor_stmntContext()
	p.parser = parser
	p.CopyFrom(ctx.(*For_stmntContext))

	return p
}

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
}

func NewForVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForVarContext {
	var p = new(ForVarContext)

	p.For_stmntContext = NewEmptyFor_stmntContext()
	p.parser = parser
	p.CopyFrom(ctx.(*For_stmntContext))

	return p
}

func (s *ForVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForVarContext) FOR() antlr.TerminalNode {
	return s.GetToken(GraffleParserFOR, 0)
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

func (s *ForVarContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *ForVarContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
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
}

func NewForRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeContext {
	var p = new(ForRangeContext)

	p.For_stmntContext = NewEmptyFor_stmntContext()
	p.parser = parser
	p.CopyFrom(ctx.(*For_stmntContext))

	return p
}

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

func (s *ForRangeContext) TO() antlr.TerminalNode {
	return s.GetToken(GraffleParserTO, 0)
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
	p.EnterRule(localctx, 36, GraffleParserRULE_for_stmnt)
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForLogicalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Match(GraffleParserFOR)
		}
		{
			p.SetState(285)
			p.Logical_expr()
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(286)
				p.Match(GraffleParserDO)
			}

		}

	case 2:
		localctx = NewForVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(289)
			p.Match(GraffleParserFOR)
		}
		{
			p.SetState(290)
			p.Atom_action()
		}
		{
			p.SetState(291)
			p.Match(GraffleParserARG_DELIM)
		}
		{
			p.SetState(292)
			p.Logical_expr()
		}
		{
			p.SetState(293)
			p.Match(GraffleParserARG_DELIM)
		}
		{
			p.SetState(294)
			p.Atom_action()
		}
		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(295)
				p.Match(GraffleParserDO)
			}

		}

	case 3:
		localctx = NewForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(298)
			p.Match(GraffleParserFOR)
		}
		{
			p.SetState(299)
			p.Variable()
		}
		{
			p.SetState(300)
			p.Match(GraffleParserIN)
		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserRANGE {
			{
				p.SetState(301)
				p.Match(GraffleParserRANGE)
			}

		}
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserFROM {
			{
				p.SetState(304)
				p.Match(GraffleParserFROM)
			}

		}
		{
			p.SetState(307)
			p.Expr()
		}
		{
			p.SetState(308)
			p.Match(GraffleParserTO)
		}
		{
			p.SetState(309)
			p.Expr()
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(310)
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

	// IsFrom_to_stmntContext differentiates from other interfaces.
	IsFrom_to_stmntContext()
}

type From_to_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *From_to_stmntContext) FROM() antlr.TerminalNode {
	return s.GetToken(GraffleParserFROM, 0)
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

func (s *From_to_stmntContext) TO() antlr.TerminalNode {
	return s.GetToken(GraffleParserTO, 0)
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
	p.EnterRule(localctx, 38, GraffleParserRULE_from_to_stmnt)
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
		p.SetState(315)
		p.Match(GraffleParserFROM)
	}
	{
		p.SetState(316)
		p.Expr()
	}
	{
		p.SetState(317)
		p.Match(GraffleParserTO)
	}
	{
		p.SetState(318)
		p.Expr()
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDO {
		{
			p.SetState(319)
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
	p.EnterRule(localctx, 40, GraffleParserRULE_function_declaration)

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

	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.One_line_function_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)
			p.Mult_line_function_declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.One_line_procedure_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(325)
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
	p.EnterRule(localctx, 42, GraffleParserRULE_one_line_function_declaration)

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
		p.SetState(328)
		p.Function_declaration_head()
	}
	{
		p.SetState(329)
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
	p.EnterRule(localctx, 44, GraffleParserRULE_mult_line_function_declaration)

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
		p.SetState(331)
		p.Function_declaration_head()
	}
	{
		p.SetState(332)
		p.Sequence()
	}
	{
		p.SetState(333)
		p.Block_end()
	}

	return localctx
}

// IFunction_declaration_headContext is an interface to support dynamic dispatch.
type IFunction_declaration_headContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturn_val returns the return_val rule contexts.
	GetReturn_val() IValueContext

	// SetReturn_val sets the return_val rule contexts.
	SetReturn_val(IValueContext)

	// IsFunction_declaration_headContext differentiates from other interfaces.
	IsFunction_declaration_headContext()
}

type Function_declaration_headContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
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

func (s *Function_declaration_headContext) GetReturn_val() IValueContext { return s.return_val }

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

func (s *Function_declaration_headContext) WHERE() antlr.TerminalNode {
	return s.GetToken(GraffleParserWHERE, 0)
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
	p.EnterRule(localctx, 46, GraffleParserRULE_function_declaration_head)
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
		p.SetState(335)
		p.Match(GraffleParserID)
	}
	{
		p.SetState(336)
		p.Match(GraffleParserL_PAREN)
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserID || _la == GraffleParserSUB {
		{
			p.SetState(337)
			p.Variable()
		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraffleParserARG_DELIM {
			{
				p.SetState(338)
				p.Match(GraffleParserARG_DELIM)
			}
			{
				p.SetState(339)
				p.Variable()
			}

			p.SetState(344)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(347)
		p.Match(GraffleParserR_PAREN)
	}
	{
		p.SetState(348)
		p.Match(GraffleParserASSIGN)
	}
	{
		p.SetState(349)

		var _x = p.Value()

		localctx.(*Function_declaration_headContext).return_val = _x
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserWHERE || _la == GraffleParserARG_DELIM {
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserARG_DELIM {
			{
				p.SetState(350)
				p.Match(GraffleParserARG_DELIM)
			}

		}
		{
			p.SetState(353)
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
	p.EnterRule(localctx, 48, GraffleParserRULE_one_line_procedure_declaration)

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
		p.SetState(356)
		p.Procedure_declaration_head()
	}
	{
		p.SetState(357)
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
	p.EnterRule(localctx, 50, GraffleParserRULE_mult_line_procedure_declaration)

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
		p.Procedure_declaration_head()
	}
	{
		p.SetState(360)
		p.Match(GraffleParserNEWLINE)
	}
	{
		p.SetState(361)
		p.Sequence()
	}
	{
		p.SetState(362)
		p.Block_end()
	}

	return localctx
}

// IProcedure_declaration_headContext is an interface to support dynamic dispatch.
type IProcedure_declaration_headContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcedure_declaration_headContext differentiates from other interfaces.
	IsProcedure_declaration_headContext()
}

type Procedure_declaration_headContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *Procedure_declaration_headContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Procedure_declaration_headContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Procedure_declaration_headContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
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

func (s *Procedure_declaration_headContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
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
	p.EnterRule(localctx, 52, GraffleParserRULE_procedure_declaration_head)
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
		p.SetState(364)
		p.Match(GraffleParserID)
	}
	{
		p.SetState(365)
		p.Match(GraffleParserL_PAREN)
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserID || _la == GraffleParserSUB {
		{
			p.SetState(366)
			p.Variable()
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraffleParserARG_DELIM {
			{
				p.SetState(367)
				p.Match(GraffleParserARG_DELIM)
			}
			{
				p.SetState(368)
				p.Variable()
			}

			p.SetState(373)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(376)
		p.Match(GraffleParserR_PAREN)
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserASSIGN {
		{
			p.SetState(377)
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

	// IsVar_assignContext differentiates from other interfaces.
	IsVar_assignContext()
}

type Var_assignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *Var_assignContext) Vertice_assign() IVertice_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertice_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertice_assignContext)
}

func (s *Var_assignContext) Edge_assign() IEdge_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdge_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEdge_assignContext)
}

func (s *Var_assignContext) Graph_assign() IGraph_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraph_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGraph_assignContext)
}

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

func (s *Var_assignContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Var_assignContext) Arithm_assign_operator() IArithm_assign_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_assign_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_assign_operatorContext)
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
	p.EnterRule(localctx, 54, GraffleParserRULE_var_assign)

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

	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(380)
			p.Vertice_assign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(381)
			p.Edge_assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(382)
			p.Graph_assign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(383)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(384)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(385)
			p.Variable()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
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

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(389)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(390)
			p.Arithm_assign_operator()
		}
		{
			p.SetState(391)
			p.Value()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(393)
			p.Labeled_assign()
		}

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

func (s *Vertice_assignContext) Vertice_type() IVertice_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertice_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertice_typeContext)
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
	p.EnterRule(localctx, 56, GraffleParserRULE_vertice_assign)

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
		p.SetState(396)
		p.Match(GraffleParserID)
	}
	{
		p.SetState(397)
		p.Match(GraffleParserASSIGN)
	}
	{
		p.SetState(398)
		p.Vertice_type()
	}

	return localctx
}

// IVertice_typeContext is an interface to support dynamic dispatch.
type IVertice_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVertice_typeContext differentiates from other interfaces.
	IsVertice_typeContext()
}

type Vertice_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVertice_typeContext() *Vertice_typeContext {
	var p = new(Vertice_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_vertice_type
	return p
}

func (*Vertice_typeContext) IsVertice_typeContext() {}

func NewVertice_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vertice_typeContext {
	var p = new(Vertice_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_vertice_type

	return p
}

func (s *Vertice_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Vertice_typeContext) V_N() antlr.TerminalNode {
	return s.GetToken(GraffleParserV_N, 0)
}

func (s *Vertice_typeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Vertice_typeContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Vertice_typeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Vertice_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vertice_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vertice_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterVertice_type(s)
	}
}

func (s *Vertice_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitVertice_type(s)
	}
}

func (s *Vertice_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitVertice_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Vertice_type() (localctx IVertice_typeContext) {
	localctx = NewVertice_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GraffleParserRULE_vertice_type)

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
		p.SetState(400)
		p.Match(GraffleParserV_N)
	}
	{
		p.SetState(401)
		p.Match(GraffleParserL_PAREN)
	}
	{
		p.SetState(402)
		p.Value()
	}
	{
		p.SetState(403)
		p.Match(GraffleParserR_PAREN)
	}

	return localctx
}

// IEdge_assignContext is an interface to support dynamic dispatch.
type IEdge_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdge_assignContext differentiates from other interfaces.
	IsEdge_assignContext()
}

type Edge_assignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdge_assignContext() *Edge_assignContext {
	var p = new(Edge_assignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_edge_assign
	return p
}

func (*Edge_assignContext) IsEdge_assignContext() {}

func NewEdge_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Edge_assignContext {
	var p = new(Edge_assignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_edge_assign

	return p
}

func (s *Edge_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Edge_assignContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Edge_assignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Edge_assignContext) Edge_type() IEdge_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdge_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEdge_typeContext)
}

func (s *Edge_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Edge_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Edge_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterEdge_assign(s)
	}
}

func (s *Edge_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitEdge_assign(s)
	}
}

func (s *Edge_assignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitEdge_assign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Edge_assign() (localctx IEdge_assignContext) {
	localctx = NewEdge_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GraffleParserRULE_edge_assign)

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
		p.SetState(405)
		p.Match(GraffleParserID)
	}
	{
		p.SetState(406)
		p.Match(GraffleParserASSIGN)
	}
	{
		p.SetState(407)
		p.Edge_type()
	}

	return localctx
}

// IEdge_typeContext is an interface to support dynamic dispatch.
type IEdge_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdge_typeContext differentiates from other interfaces.
	IsEdge_typeContext()
}

type Edge_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdge_typeContext() *Edge_typeContext {
	var p = new(Edge_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_edge_type
	return p
}

func (*Edge_typeContext) IsEdge_typeContext() {}

func NewEdge_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Edge_typeContext {
	var p = new(Edge_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_edge_type

	return p
}

func (s *Edge_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Edge_typeContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Edge_typeContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Edge_typeContext) Edge() IEdgeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdgeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEdgeContext)
}

func (s *Edge_typeContext) E_N() antlr.TerminalNode {
	return s.GetToken(GraffleParserE_N, 0)
}

func (s *Edge_typeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Edge_typeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Edge_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Edge_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Edge_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterEdge_type(s)
	}
}

func (s *Edge_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitEdge_type(s)
	}
}

func (s *Edge_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitEdge_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Edge_type() (localctx IEdge_typeContext) {
	localctx = NewEdge_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GraffleParserRULE_edge_type)

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

	p.SetState(418)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserINT, GraffleParserSTRING, GraffleParserBOOL, GraffleParserNOT, GraffleParserV_N, GraffleParserID, GraffleParserL_PAREN, GraffleParserSUB:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.Value()
		}
		{
			p.SetState(410)
			p.Edge()
		}
		{
			p.SetState(411)
			p.Value()
		}

	case GraffleParserE_N:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(413)
			p.Match(GraffleParserE_N)
		}
		{
			p.SetState(414)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(415)
			p.Value()
		}
		{
			p.SetState(416)
			p.Match(GraffleParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEdgeContext is an interface to support dynamic dispatch.
type IEdgeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdgeContext differentiates from other interfaces.
	IsEdgeContext()
}

type EdgeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdgeContext() *EdgeContext {
	var p = new(EdgeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_edge
	return p
}

func (*EdgeContext) IsEdgeContext() {}

func NewEdgeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EdgeContext {
	var p = new(EdgeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_edge

	return p
}

func (s *EdgeContext) GetParser() antlr.Parser { return s.parser }

func (s *EdgeContext) OR_EDGE_LR() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_EDGE_LR, 0)
}

func (s *EdgeContext) Or_w_edge_lr() IOr_w_edge_lrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_w_edge_lrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_w_edge_lrContext)
}

func (s *EdgeContext) OR_EDGE_RL() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_EDGE_RL, 0)
}

func (s *EdgeContext) Or_w_edge_rl() IOr_w_edge_rlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_w_edge_rlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_w_edge_rlContext)
}

func (s *EdgeContext) UNOR_EDGE() antlr.TerminalNode {
	return s.GetToken(GraffleParserUNOR_EDGE, 0)
}

func (s *EdgeContext) Unor_w_edge() IUnor_w_edgeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnor_w_edgeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnor_w_edgeContext)
}

func (s *EdgeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EdgeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterEdge(s)
	}
}

func (s *EdgeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitEdge(s)
	}
}

func (s *EdgeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitEdge(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Edge() (localctx IEdgeContext) {
	localctx = NewEdgeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GraffleParserRULE_edge)

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

	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(420)
			p.Match(GraffleParserOR_EDGE_LR)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(421)
			p.Or_w_edge_lr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(422)
			p.Match(GraffleParserOR_EDGE_RL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(423)
			p.Or_w_edge_rl()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(424)
			p.Match(GraffleParserUNOR_EDGE)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(425)
			p.Unor_w_edge()
		}

	}

	return localctx
}

// IOr_w_edge_lrContext is an interface to support dynamic dispatch.
type IOr_w_edge_lrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOr_w_edge_lrContext differentiates from other interfaces.
	IsOr_w_edge_lrContext()
}

type Or_w_edge_lrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOr_w_edge_lrContext() *Or_w_edge_lrContext {
	var p = new(Or_w_edge_lrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_or_w_edge_lr
	return p
}

func (*Or_w_edge_lrContext) IsOr_w_edge_lrContext() {}

func NewOr_w_edge_lrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_w_edge_lrContext {
	var p = new(Or_w_edge_lrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_or_w_edge_lr

	return p
}

func (s *Or_w_edge_lrContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_w_edge_lrContext) SUB() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, 0)
}

func (s *Or_w_edge_lrContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_BRACKET, 0)
}

func (s *Or_w_edge_lrContext) Number_type() INumber_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumber_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumber_typeContext)
}

func (s *Or_w_edge_lrContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_BRACKET, 0)
}

func (s *Or_w_edge_lrContext) OR_EDGE_LR() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_EDGE_LR, 0)
}

func (s *Or_w_edge_lrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_w_edge_lrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_w_edge_lrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOr_w_edge_lr(s)
	}
}

func (s *Or_w_edge_lrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOr_w_edge_lr(s)
	}
}

func (s *Or_w_edge_lrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOr_w_edge_lr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Or_w_edge_lr() (localctx IOr_w_edge_lrContext) {
	localctx = NewOr_w_edge_lrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GraffleParserRULE_or_w_edge_lr)

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
		p.SetState(428)
		p.Match(GraffleParserSUB)
	}
	{
		p.SetState(429)
		p.Match(GraffleParserL_BRACKET)
	}
	{
		p.SetState(430)
		p.Number_type()
	}
	{
		p.SetState(431)
		p.Match(GraffleParserR_BRACKET)
	}
	{
		p.SetState(432)
		p.Match(GraffleParserOR_EDGE_LR)
	}

	return localctx
}

// IOr_w_edge_rlContext is an interface to support dynamic dispatch.
type IOr_w_edge_rlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOr_w_edge_rlContext differentiates from other interfaces.
	IsOr_w_edge_rlContext()
}

type Or_w_edge_rlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOr_w_edge_rlContext() *Or_w_edge_rlContext {
	var p = new(Or_w_edge_rlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_or_w_edge_rl
	return p
}

func (*Or_w_edge_rlContext) IsOr_w_edge_rlContext() {}

func NewOr_w_edge_rlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_w_edge_rlContext {
	var p = new(Or_w_edge_rlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_or_w_edge_rl

	return p
}

func (s *Or_w_edge_rlContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_w_edge_rlContext) OR_EDGE_RL() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_EDGE_RL, 0)
}

func (s *Or_w_edge_rlContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_BRACKET, 0)
}

func (s *Or_w_edge_rlContext) Number_type() INumber_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumber_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumber_typeContext)
}

func (s *Or_w_edge_rlContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_BRACKET, 0)
}

func (s *Or_w_edge_rlContext) SUB() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, 0)
}

func (s *Or_w_edge_rlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_w_edge_rlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_w_edge_rlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOr_w_edge_rl(s)
	}
}

func (s *Or_w_edge_rlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOr_w_edge_rl(s)
	}
}

func (s *Or_w_edge_rlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOr_w_edge_rl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Or_w_edge_rl() (localctx IOr_w_edge_rlContext) {
	localctx = NewOr_w_edge_rlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GraffleParserRULE_or_w_edge_rl)

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
		p.SetState(434)
		p.Match(GraffleParserOR_EDGE_RL)
	}
	{
		p.SetState(435)
		p.Match(GraffleParserL_BRACKET)
	}
	{
		p.SetState(436)
		p.Number_type()
	}
	{
		p.SetState(437)
		p.Match(GraffleParserR_BRACKET)
	}
	{
		p.SetState(438)
		p.Match(GraffleParserSUB)
	}

	return localctx
}

// IUnor_w_edgeContext is an interface to support dynamic dispatch.
type IUnor_w_edgeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnor_w_edgeContext differentiates from other interfaces.
	IsUnor_w_edgeContext()
}

type Unor_w_edgeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnor_w_edgeContext() *Unor_w_edgeContext {
	var p = new(Unor_w_edgeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_unor_w_edge
	return p
}

func (*Unor_w_edgeContext) IsUnor_w_edgeContext() {}

func NewUnor_w_edgeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unor_w_edgeContext {
	var p = new(Unor_w_edgeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_unor_w_edge

	return p
}

func (s *Unor_w_edgeContext) GetParser() antlr.Parser { return s.parser }

func (s *Unor_w_edgeContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserSUB)
}

func (s *Unor_w_edgeContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, i)
}

func (s *Unor_w_edgeContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_BRACKET, 0)
}

func (s *Unor_w_edgeContext) Number_type() INumber_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumber_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumber_typeContext)
}

func (s *Unor_w_edgeContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_BRACKET, 0)
}

func (s *Unor_w_edgeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unor_w_edgeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unor_w_edgeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterUnor_w_edge(s)
	}
}

func (s *Unor_w_edgeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitUnor_w_edge(s)
	}
}

func (s *Unor_w_edgeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitUnor_w_edge(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Unor_w_edge() (localctx IUnor_w_edgeContext) {
	localctx = NewUnor_w_edgeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GraffleParserRULE_unor_w_edge)

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
		p.SetState(440)
		p.Match(GraffleParserSUB)
	}
	{
		p.SetState(441)
		p.Match(GraffleParserL_BRACKET)
	}
	{
		p.SetState(442)
		p.Number_type()
	}
	{
		p.SetState(443)
		p.Match(GraffleParserR_BRACKET)
	}
	{
		p.SetState(444)
		p.Match(GraffleParserSUB)
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

func (s *Graph_assignContext) Graph_type() IGraph_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraph_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGraph_typeContext)
}

func (s *Graph_assignContext) Arithm_assign_operator() IArithm_assign_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_assign_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_assign_operatorContext)
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

func (s *Graph_assignContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *Graph_assignContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
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
	p.EnterRule(localctx, 72, GraffleParserRULE_graph_assign)

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

	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(446)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(447)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(448)
			p.Graph_type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(449)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(450)
			p.Arithm_assign_operator()
		}
		{
			p.SetState(451)
			p.Value()
		}
		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(452)
					p.Match(GraffleParserARG_DELIM)
				}
				{
					p.SetState(453)
					p.Value()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(456)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IGraph_typeContext is an interface to support dynamic dispatch.
type IGraph_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGraph_typeContext differentiates from other interfaces.
	IsGraph_typeContext()
}

type Graph_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraph_typeContext() *Graph_typeContext {
	var p = new(Graph_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_graph_type
	return p
}

func (*Graph_typeContext) IsGraph_typeContext() {}

func NewGraph_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Graph_typeContext {
	var p = new(Graph_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_graph_type

	return p
}

func (s *Graph_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Graph_typeContext) G_N() antlr.TerminalNode {
	return s.GetToken(GraffleParserG_N, 0)
}

func (s *Graph_typeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Graph_typeContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Graph_typeContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Graph_typeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Graph_typeContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *Graph_typeContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
}

func (s *Graph_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Graph_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Graph_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterGraph_type(s)
	}
}

func (s *Graph_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitGraph_type(s)
	}
}

func (s *Graph_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitGraph_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Graph_type() (localctx IGraph_typeContext) {
	localctx = NewGraph_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GraffleParserRULE_graph_type)
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

	p.SetState(479)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserG_N:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(460)
			p.Match(GraffleParserG_N)
		}
		{
			p.SetState(461)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(462)
			p.Value()
		}
		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraffleParserARG_DELIM {
			{
				p.SetState(463)
				p.Match(GraffleParserARG_DELIM)
			}
			{
				p.SetState(464)
				p.Value()
			}

			p.SetState(469)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(470)
			p.Match(GraffleParserR_PAREN)
		}

	case GraffleParserINT, GraffleParserSTRING, GraffleParserBOOL, GraffleParserNOT, GraffleParserV_N, GraffleParserID, GraffleParserL_PAREN, GraffleParserSUB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(472)
			p.Value()
		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(473)
					p.Match(GraffleParserARG_DELIM)
				}
				{
					p.SetState(474)
					p.Value()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(477)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext())
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *Labeled_assignContext) Edge_assign() IEdge_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdge_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEdge_assignContext)
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
	p.EnterRule(localctx, 76, GraffleParserRULE_labeled_assign)
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

	p.SetState(493)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(481)
			p.Vertice_assign()
		}
		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserLABEL || _la == GraffleParserML_LABEL {
			{
				p.SetState(482)
				p.Label()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(485)
			p.Edge_assign()
		}
		p.SetState(487)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserLABEL || _la == GraffleParserML_LABEL {
			{
				p.SetState(486)
				p.Label()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(489)
			p.Graph_assign()
		}
		p.SetState(491)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserLABEL || _la == GraffleParserML_LABEL {
			{
				p.SetState(490)
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

func (s *ExprContext) Builtin_type() IBuiltin_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltin_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltin_typeContext)
}

func (s *ExprContext) Vertice_type() IVertice_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertice_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertice_typeContext)
}

func (s *ExprContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *ExprContext) Edge_type() IEdge_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdge_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEdge_typeContext)
}

func (s *ExprContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *ExprContext) Graph_type() IGraph_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraph_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGraph_typeContext)
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
	p.EnterRule(localctx, 78, GraffleParserRULE_expr)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(495)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(496)
			p.Integral_expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(497)
			p.Builtin_type()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(498)
			p.Vertice_type()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(499)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(500)
			p.Edge_type()
		}
		{
			p.SetState(501)
			p.Match(GraffleParserR_PAREN)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(503)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(504)
			p.Graph_type()
		}
		{
			p.SetState(505)
			p.Match(GraffleParserR_PAREN)
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

func (s *Integral_exprContext) Builtin_type() IBuiltin_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltin_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltin_typeContext)
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
	p.EnterRule(localctx, 80, GraffleParserRULE_integral_expr)

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

	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(509)
			p.Function_call()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(510)
			p.Builtin_type()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(511)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(512)
			p.Logical_expr()
		}
		{
			p.SetState(513)
			p.Match(GraffleParserR_PAREN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(515)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(516)
			p.Arithm_expr()
		}
		{
			p.SetState(517)
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
	p.EnterRule(localctx, 82, GraffleParserRULE_logical_expr)

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

	p.SetState(530)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(521)

			var _x = p.Expr()

			localctx.(*Logical_exprContext).left = _x
		}
		{
			p.SetState(522)

			var _x = p.Bin_log_operator()

			localctx.(*Logical_exprContext).bin_op = _x
		}
		{
			p.SetState(523)

			var _x = p.Log_expr_operand()

			localctx.(*Logical_exprContext).right = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(525)

			var _x = p.Unar_log_operator()

			localctx.(*Logical_exprContext).un_op = _x
		}
		{
			p.SetState(526)
			p.Log_expr_operand()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(528)
			p.Arithm_expr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(529)
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
	p.EnterRule(localctx, 84, GraffleParserRULE_log_expr_operand)

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

	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(532)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(533)
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
	p.EnterRule(localctx, 86, GraffleParserRULE_bin_log_operator)
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
		p.SetState(536)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraffleParserNEQ)|(1<<GraffleParserEQUALS)|(1<<GraffleParserLESS_THAN)|(1<<GraffleParserGR_THAN)|(1<<GraffleParserLESS_THAN_E)|(1<<GraffleParserGR_THAN_E)|(1<<GraffleParserAND)|(1<<GraffleParserOR)|(1<<GraffleParserXOR)|(1<<GraffleParserNOR)|(1<<GraffleParserNAND))) != 0) {
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
	p.EnterRule(localctx, 88, GraffleParserRULE_unar_log_operator)

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
		p.SetState(538)
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
	p.EnterRule(localctx, 90, GraffleParserRULE_arithm_expr)

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
		p.SetState(540)

		var _x = p.Expr()

		localctx.(*Arithm_exprContext).left = _x
	}
	{
		p.SetState(541)

		var _x = p.Bin_arithm_operator()

		localctx.(*Arithm_exprContext).op = _x
	}
	{
		p.SetState(542)

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
	p.EnterRule(localctx, 92, GraffleParserRULE_arithm_expr_operand)

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

	p.SetState(546)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(544)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(545)
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
	p.EnterRule(localctx, 94, GraffleParserRULE_bin_arithm_operator)
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

		if !(((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(GraffleParserADD-59))|(1<<(GraffleParserSUB-59))|(1<<(GraffleParserMULT-59))|(1<<(GraffleParserDIV-59)))) != 0) {
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
	p.EnterRule(localctx, 96, GraffleParserRULE_arithm_assign_operator)
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
		p.SetState(550)
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
	p.EnterRule(localctx, 98, GraffleParserRULE_builtin_function_call)

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

	switch p.GetTokenStream().LA(1) {
	case GraffleParserPRINTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(552)
			p.Built_func_print()
		}

	case GraffleParserKEY_INPUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(553)
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
	p.EnterRule(localctx, 100, GraffleParserRULE_built_func_print)

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

	p.SetState(562)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(556)
			p.Match(GraffleParserPRINTER)
		}
		{
			p.SetState(557)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(558)
			p.Match(GraffleParserPRINTER)
		}
		{
			p.SetState(559)
			p.Value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(560)
			p.Match(GraffleParserPRINTER)
		}
		{
			p.SetState(561)
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
	p.EnterRule(localctx, 102, GraffleParserRULE_built_func_input)

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
		p.SetState(564)
		p.Match(GraffleParserKEY_INPUT)
	}
	{
		p.SetState(565)
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
	p.EnterRule(localctx, 104, GraffleParserRULE_function_call)
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
		p.SetState(567)
		p.Match(GraffleParserID)
	}
	{
		p.SetState(568)
		p.Match(GraffleParserL_PAREN)
	}
	p.SetState(577)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraffleParserINT)|(1<<GraffleParserSTRING)|(1<<GraffleParserBOOL)|(1<<GraffleParserNOT))) != 0) || (((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(GraffleParserV_N-48))|(1<<(GraffleParserID-48))|(1<<(GraffleParserL_PAREN-48))|(1<<(GraffleParserSUB-48)))) != 0) {
		{
			p.SetState(569)
			p.Value()
		}
		p.SetState(574)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraffleParserARG_DELIM {
			{
				p.SetState(570)
				p.Match(GraffleParserARG_DELIM)
			}
			{
				p.SetState(571)
				p.Value()
			}

			p.SetState(576)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(579)
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
	p.EnterRule(localctx, 106, GraffleParserRULE_label)
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
		p.SetState(581)
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

func (s *ValueContext) Arithm_expr() IArithm_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_exprContext)
}

func (s *ValueContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *ValueContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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
	p.EnterRule(localctx, 108, GraffleParserRULE_value)

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

	p.SetState(586)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(583)
			p.Arithm_expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(584)
			p.Logical_expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(585)
			p.Expr()
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
	p.EnterRule(localctx, 110, GraffleParserRULE_variable)
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

	if _la == GraffleParserSUB {
		{
			p.SetState(588)
			p.Match(GraffleParserSUB)
		}

	}
	{
		p.SetState(591)
		p.Match(GraffleParserID)
	}

	return localctx
}

// IBuiltin_typeContext is an interface to support dynamic dispatch.
type IBuiltin_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltin_typeContext differentiates from other interfaces.
	IsBuiltin_typeContext()
}

type Builtin_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltin_typeContext() *Builtin_typeContext {
	var p = new(Builtin_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_builtin_type
	return p
}

func (*Builtin_typeContext) IsBuiltin_typeContext() {}

func NewBuiltin_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Builtin_typeContext {
	var p = new(Builtin_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_builtin_type

	return p
}

func (s *Builtin_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Builtin_typeContext) Number_type() INumber_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumber_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumber_typeContext)
}

func (s *Builtin_typeContext) STRING() antlr.TerminalNode {
	return s.GetToken(GraffleParserSTRING, 0)
}

func (s *Builtin_typeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GraffleParserBOOL, 0)
}

func (s *Builtin_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Builtin_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Builtin_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBuiltin_type(s)
	}
}

func (s *Builtin_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBuiltin_type(s)
	}
}

func (s *Builtin_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBuiltin_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Builtin_type() (localctx IBuiltin_typeContext) {
	localctx = NewBuiltin_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, GraffleParserRULE_builtin_type)

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

	p.SetState(596)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserINT, GraffleParserSUB:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(593)
			p.Number_type()
		}

	case GraffleParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(594)
			p.Match(GraffleParserSTRING)
		}

	case GraffleParserBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(595)
			p.Match(GraffleParserBOOL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumber_typeContext is an interface to support dynamic dispatch.
type INumber_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumber_typeContext differentiates from other interfaces.
	IsNumber_typeContext()
}

type Number_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_typeContext() *Number_typeContext {
	var p = new(Number_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_number_type
	return p
}

func (*Number_typeContext) IsNumber_typeContext() {}

func NewNumber_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_typeContext {
	var p = new(Number_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_number_type

	return p
}

func (s *Number_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_typeContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserINT)
}

func (s *Number_typeContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserINT, i)
}

func (s *Number_typeContext) SUB() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, 0)
}

func (s *Number_typeContext) ACT_DELIM() antlr.TerminalNode {
	return s.GetToken(GraffleParserACT_DELIM, 0)
}

func (s *Number_typeContext) ARG_DELIM() antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, 0)
}

func (s *Number_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterNumber_type(s)
	}
}

func (s *Number_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitNumber_type(s)
	}
}

func (s *Number_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitNumber_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Number_type() (localctx INumber_typeContext) {
	localctx = NewNumber_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, GraffleParserRULE_number_type)
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
	p.SetState(599)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserSUB {
		{
			p.SetState(598)
			p.Match(GraffleParserSUB)
		}

	}
	{
		p.SetState(601)
		p.Match(GraffleParserINT)
	}
	p.SetState(604)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 73, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(602)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GraffleParserACT_DELIM || _la == GraffleParserARG_DELIM) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(603)
			p.Match(GraffleParserINT)
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

func (s *Block_endContext) AllACT_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserACT_DELIM)
}

func (s *Block_endContext) ACT_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserACT_DELIM, i)
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
	p.EnterRule(localctx, 116, GraffleParserRULE_block_end)
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
	p.SetState(607)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserNEWLINE {
		{
			p.SetState(606)
			p.Match(GraffleParserNEWLINE)
		}

	}
	{
		p.SetState(609)
		p.Match(GraffleParserBLOCK_END)
	}
	p.SetState(613)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 75, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(610)
				p.Match(GraffleParserACT_DELIM)
			}

		}
		p.SetState(615)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 75, p.GetParserRuleContext())
	}
	p.SetState(617)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 76, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(616)
			p.Match(GraffleParserNEWLINE)
		}

	}

	return localctx
}
