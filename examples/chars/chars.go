// Package main defines a demo for viewing the character map
// supported by the library, as defined by CP437
// https://en.wikipedia.org/wiki/Code_page_437
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Joakker/tcod-go"
	"github.com/Joakker/tcod-go/tinput"
)

var charNames = map[string]byte{
	"CharSmiley":                    1,
	"CharSmileyInv":                 2,
	"CharHeart":                     3,
	"CharDiamon":                    4,
	"CharClub":                      5,
	"CharSpade":                     6,
	"CharBullet":                    7,
	"CharBulletActive":              8,
	"CharBulletInv":                 9,
	"CharBulletInvActive":           10,
	"CharMale":                      11,
	"CharFemale":                    12,
	"CharNote":                      13,
	"CharDoubleNote":                14,
	"CharStar":                      15,
	"CharTriangleArrowR":            16,
	"CharTriangleArrowL":            17,
	"CharDoubleArrowV":              18,
	"CharDoubleExclamation":         19,
	"CharLamp":                      21,
	"CharUnderBlock":                22,
	"CharDoubleArrowVertUnderscore": 23,
	"CharArrowUp":                   24,
	"CharArrowDown":                 25,
	"CharArrowRight":                26,
	"CharArrowLeft":                 27,
	"CharRightAngle":                28,
	"CharDoubleArrowH":              29,
	"CharTriangleArrowU":            30,
	"CharTriangleArrowD":            31,
	"CharTick":                      39,
	"CharPilcrow":                   94,
	"CharHouse":                     127,
	"CharUpperCedilla":              128,
	"CharULowermlaut":               129,
	"CharELowerAcute":               130,
	"CharALowerCaret":               131,
	"CharALowerUmlaut":              132,
	"CharALowerGrave":               133,
	"CharALowerCircle":              134,
	"CharLowerCedilla":              135,
	"CharELowerCaret":               136,
	"CharELowerUmlaut":              137,
	"CharELowerGrave":               138,
	"CharILowerUmlaut":              139,
	"CharILowerCaret":               140,
	"CharILowerGrave":               141,
	"CharAUpperUmlaut":              142,
	"CharAUpperCircle":              143,
	"CharEUpperAcute":               144,
	"CharAELower":                   145,
	"CharAEUpper":                   146,
	"CharOLowerCaret":               147,
	"CharOLowerUmlaut":              148,
	"CharOLowerGrave":               149,
	"CharULowerCaret":               150,
	"CharULowerGrave":               151,
	"CharYLowerUmlaut":              152,
	"CharOUpperUmlaut":              153,
	"CharUUpperUmlaut":              154,
	"CharCent":                      155,
	"CharPound":                     156,
	"CharYen":                       157,
	"CharPoint":                     158,
	"CharFunction":                  159,
	"CharALowerAcute":               160,
	"CharILowerAccute":              161,
	"CharOLowerAccute":              162,
	"CharULowerAccute":              163,
	"CharNTildeLower":               164,
	"CharNTildeUpper":               165,
	"CharExpA":                      166,
	"CharExpO":                      167,
	"CharInvQuestion":               168,
	"CharLogicalNot":                169,
	"CharLogicalNotInv":             170,
	"CharHalf":                      171,
	"CharQuarter":                   172,
	"CharInvExclamation":            173,
	"CharDoubleArrowL":              174,
	"CharDoubleArrowR":              175,
	"CharBlock1":                    176,
	"CharBlock2":                    177,
	"CharBlock3":                    178,
	"CharWallV":                     179,
	"CharTeeW":                      180,
	"CharTeeHDoubleW":               181,
	"CharTeeVDoubleW":               182,
	"CharCornerVDoubleNE":           183,
	"CharCornerHDoubleNE":           184,
	"CharTeeDoubleW":                185,
	"CharWallDoubleV":               186,
	"CharCornerDoubleNE":            187,
	"CharCornerDoubleSE":            188,
	"CharCornerVDoubleSE":           189,
	"CharCornerHDoubleSE":           190,
	"CharCornerNE":                  191,
	"CharCornerSW":                  192,
	"CharTeeN":                      193,
	"CharTeeS":                      194,
	"CharTeeE":                      195,
	"CharWallH":                     196,
	"CharWallCross":                 197,
	"CharTeeHDoubleE":               198,
	"CharTeeVDoubleE":               199,
	"CharCornerDoubleSW":            200,
	"CharCornerDoubleNW":            201,
	"CharTeeDoubleN":                202,
	"CharTeeDoubleS":                203,
	"CharTeeDoubleE":                204,
	"CharWallDoubleH":               205,
	"CharWallDoubleCross":           206,
	"CharTeeHDoubleN":               207,
	"CharTeeVDoubleN":               208,
	"CharTeeHDoubleS":               209,
	"CharTeeVDoubleS":               210,
	"CharCornerVDoubleSW":           211,
	"CharCornerHdoubleSW":           212,
	"CharCornerHDoubleNW":           213,
	"CharCornerVDoubleNW":           214,
	"CharWallVDoubleCross":          215,
	"CharWallHDoubleCross":          216,
	"CharCornerSE":                  217,
	"CharCornerNW":                  218,
	"CharBlock4":                    219,
	"CharSubpDown":                  220,
	"CharSubpLeft":                  221,
	"CharSubpRight":                 222,
	"CharSubpUp":                    223,
	"CharAlpha":                     224,
	"CharBeta":                      225,
	"CharGamma":                     226,
	"CharPi":                        227,
	"CharSigmaUpper":                228,
	"CharSigmaLower":                229,
	"CharMi":                        230,
	"CharTau":                       231,
	"CharPhiUpper":                  232,
	"CharTheta":                     233,
	"CharOmega":                     234,
	"CharDelta":                     235,
	"CharInfinite":                  236,
	"CharPhiLower":                  237,
	"CharEpsilon":                   238,
	"CharArc":                       239,
	"CharList":                      240,
	"CharPlusMinus":                 241,
	"CharGreaterEqual":              242,
	"CharLowerEqual":                243,
	"CharIntegralUp":                244,
	"CharIntegralDown":              245,
	"CharDivision":                  246,
	"CharAlmostEqual":               247,
	"CharDegree":                    248,
	"CharDot1":                      249,
	"CharDot2":                      250,
	"CharRoot":                      251,
	"CharExpN":                      252,
	"CharExp2":                      253,
	"CharSquare":                    254,
	"CharSquareInv":                 255,
}

var (
	all   bool
	name  bool
	input bool
)

func init() {
	flag.BoolVar(&all, "all", false, "View all the characters")
	flag.BoolVar(&name, "name", false, "Enter a character's name to see it rendered")
	flag.BoolVar(&input, "input", false, "Display the input string")
	l := flag.Bool("list", false, "List all the non-ASCII characters")

	flag.Parse()

	if *l {
		for k := range charNames {
			fmt.Println(k)
		}
		os.Exit(0)
	}
}

func main() {
	if all {
		root, err := tcod.InitRoot(
			16, 16, "Character Demo",
			false, tcod.RenderSDL2,
		)
		if err != nil {
			log.Fatal(err)
		}

		for !tcod.WindowClosed() {
			i := tinput.NewInput()
			i.Check()
			root.Clear()
			for i := 0; i < 16; i++ {
				for j := 0; j < 16; j++ {
					root.SetChar(i, j, byte(i+16*j))
				}
			}
			tcod.Flush()
		}
	} else if name {
		reader := bufio.NewReader(os.Stdin)
		name, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		name = name[:len(name)-1]
		root, err := tcod.InitRoot(3, 3, name, false, tcod.RenderSDL2)
		if err != nil {
			log.Fatal(err)
		}
		for !tcod.WindowClosed() {
			i := tinput.NewInput()
			i.Check()
			root.Clear()
			root.SetChar(1, 1, charNames[name])
			tcod.Flush()
		}
	} else if input {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = line[:len(line)-1]
		fmt.Println(line)
		root, err := tcod.InitRoot(20, 5, line, false, tcod.RenderSDL2)
		if err != nil {
			log.Fatal(err)
		}
		for !tcod.WindowClosed() {
			i := tinput.NewInput()
			i.Check()
			root.Clear()
			root.Print(1, 1, line)
			tcod.Flush()
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		c, err := reader.ReadByte()
		if err != nil {
			log.Fatal(err)
		}

		root, err := tcod.InitRoot(3, 3, "Character Demo", false, tcod.RenderSDL2)
		if err != nil {
			log.Fatal(err)
		}

		for !tcod.WindowClosed() {
			i := tinput.NewInput()
			i.Check()
			root.Clear()
			root.SetChar(1, 1, c)
			tcod.Flush()
		}
	}
}
