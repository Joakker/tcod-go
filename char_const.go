package tcod

// #include <libtcod.h>
import "C"

// rune enum
var (
	// CharSmiley is Char 0x263A
	CharSmiley rune = 0x263A
	// CharSmileyInv is Char 0x263B
	CharSmileyInv rune = 0x263B
	// CharHeart is Char 0x2665
	CharHeart rune = 0x2665
	// CharDiamond is Char 0x2666
	CharDiamond rune = 0x2666
	// CharClub is Char 0x2663
	CharClub rune = 0x2663
	// CharSpade is Char 0x2660
	CharSpade rune = 0x2660
	// CharBullet is Char 0x2022
	CharBullet rune = 0x2022
	// CharBulletActive is Char 0x25D8
	CharBulletActive rune = 0x25D8
	// CharBulletInv is Char 0x25CB
	CharBulletInv rune = 0x25CB
	// CharBulletInvActive is Char 0x25D9
	CharBulletInvActive rune = 0x25D9
	// CharMale is Char 0x2642
	CharMale rune = 0x2642
	// CharFemale is Char 0x2640
	CharFemale rune = 0x2640
	// CharNote is Char 0x266A
	CharNote rune = 0x266A
	// CharDoubleNote is Char 0x266B
	CharDoubleNote rune = 0x266B
	// CharStar is Char 0x263C
	CharStar rune = 0x263C
	// CharTriangleArrowR is Char 0x25BA
	CharTriangleArrowR rune = 0x25BA
	// CharTriangleArrowL is Char 0x25C4
	CharTriangleArrowL rune = 0x25C4
	// CharDoubleArrowV is Char 0x2195
	CharDoubleArrowV rune = 0x2195
	// CharDoubleExclamation is Char 0x203C
	CharDoubleExclamation rune = 0x203C
	// CharPilcrow is Char 0x00B6
	CharPilcrow rune = 0x00B6
	// CharLamp is Char 0x00A7
	CharLamp rune = 0x00A7
	// CharUnderBlock is Char 0x25AC
	CharUnderBlock rune = 0x25AC
	// CharDoubleArrowVertUnderscore is Char 0x21A8
	CharDoubleArrowVertUnderscore rune = 0x21A8
	// CharArrowUp is Char 0x2191
	CharArrowUp rune = 0x2191
	// CharArrowDown is Char 0x2193
	CharArrowDown rune = 0x2193
	// CharArrowRight is Char 0x2192
	CharArrowRight rune = 0x2192
	// CharArrowLeft is Char 0x2190
	CharArrowLeft rune = 0x2190
	// CharRightAngle is Char 0x221F
	CharRightAngle rune = 0x221F
	// CharDoubleArrowH is Char 0x2194
	CharDoubleArrowH rune = 0x2194
	// CharTriangleArrowU is Char 0x25B2
	CharTriangleArrowU rune = 0x25B2
	// CharTriangleArrowD is Char 0x25BC
	CharTriangleArrowD rune = 0x25BC
	// CharHouse is Char 0x2302
	CharHouse rune = 0x2302
	// CharUpperCedilla UpperCedilla
	CharUpperCedilla rune = 128
	// CharULowermlaut ULowermlaut
	CharULowermlaut rune = 129
	// CharELowerAcute ELowerAcute
	CharELowerAcute rune = 130
	// CharALowerCaret ALowerCaret
	CharALowerCaret rune = 131
	// CharALowerUmlaut ALowerUmlaut
	CharALowerUmlaut rune = 132
	// CharALowerGrave ALowerGrave
	CharALowerGrave rune = 133
	// CharALowerCircle ALowerCircle
	CharALowerCircle rune = 134
	// CharLowerCedilla LowerCedilla
	CharLowerCedilla rune = 135
	// CharELowerCaret ELowerCaret
	CharELowerCaret rune = 136
	// CharELowerUmlaut ELowerUmlaut
	CharELowerUmlaut rune = 137
	// CharELowerGrave ELowerGrave
	CharELowerGrave rune = 138
	// CharILowerUmlaut ILowerUmlaut
	CharILowerUmlaut rune = 139
	// CharILowerCaret ILowerCaret
	CharILowerCaret rune = 140
	// CharILowerGrave ILowerGrave
	CharILowerGrave rune = 141
	// CharAUpperUmlaut AUpperUmlaut
	CharAUpperUmlaut rune = 142
	// CharAUpperCircle AUpperCircle
	CharAUpperCircle rune = 143
	// CharEUpperAcute EUpperAcute
	CharEUpperAcute rune = 144
	// CharAELower AELower
	CharAELower rune = 145
	// CharAEUpper AEUpper
	CharAEUpper rune = 146
	// CharOLowerCaret OLowerCaret
	CharOLowerCaret rune = 147
	// CharOLowerUmlaut OLowerUmlaut
	CharOLowerUmlaut rune = 148
	// CharOLowerGrave OLowerGrave
	CharOLowerGrave rune = 149
	// CharULowerCaret ULowerCaret
	CharULowerCaret rune = 150
	// CharULowerGrave ULowerGrave
	CharULowerGrave rune = 151
	// CharYLowerUmlaut YLowerUmlaut
	CharYLowerUmlaut rune = 152
	// CharOUpperUmlaut OUpperUmlaut
	CharOUpperUmlaut rune = 153
	// CharUUpperUmlaut UUpperUmlaut
	CharUUpperUmlaut rune = 154
	// CharCent is Char 0x00A2
	CharCent rune = 0x00A2
	// CharPound is Char 0x00A3
	CharPound rune = 0x00A3
	// CharYen is Char 0x00A5
	CharYen rune = 0x00A5
	// CharPoint is Char 0x20A7
	CharPoint rune = 0x20A7
	// CharFunction is Char 0x0192
	CharFunction rune = 0x0192
	// CharALowerAcute ALowerAcute
	CharALowerAcute rune = 160
	// CharILowerAccute ILowerAccute
	CharILowerAccute rune = 161
	// CharOLowerAccute OLowerAccute
	CharOLowerAccute rune = 162
	// CharULowerAccute ULowerAccute
	CharULowerAccute rune = 163
	// CharNTildeLower NTildeLower
	CharNTildeLower rune = 164
	// CharNTildeUpper NTildeUpper
	CharNTildeUpper rune = 165
	// CharExpA is Char 0x00AA
	CharExpA rune = 0x00AA
	// CharExpO is Char 0x00BA
	CharExpO rune = 0x00BA
	// CharInvQuestion is Char 0x00BF
	CharInvQuestion rune = 0x00BF
	// CharLogicalNot is Char 0x2310
	CharLogicalNot rune = 0x2310
	// CharLogicalNotInv is Char 0x00AC
	CharLogicalNotInv rune = 0x00AC
	// CharHalf is Char 0x00BD
	CharHalf rune = 0x00BD
	// CharQuarter is Char 0x00BC
	CharQuarter rune = 0x00BC
	// CharInvExclamation is Char 0x00A1
	CharInvExclamation rune = 0x00A1
	// CharDoubleArrowL is Char 0x00AB
	CharDoubleArrowL rune = 0x00AB
	// CharDoubleArrowR is Char 0x00BB
	CharDoubleArrowR rune = 0x00BB
	// CharBlock1 is Char 0x2591
	CharBlock1 rune = 0x2591
	// CharBlock2 is Char 0x2592
	CharBlock2 rune = 0x2592
	// CharBlock3 is Char 0x2593
	CharBlock3 rune = 0x2593
	// CharWallV is Char 0x2502
	CharWallV rune = 0x2502
	// CharTeeW is Char 0x2524
	CharTeeW rune = 0x2524
	// CharTeeHDoubleW is Char 0x2561
	CharTeeHDoubleW rune = 0x2561
	// CharTeeVDoubleW is Char 0x2562
	CharTeeVDoubleW rune = 0x2562
	// CharCornerVDoubleNE is Char 0x2556
	CharCornerVDoubleNE rune = 0x2556
	// CharCornerHDoubleNE is Char 0x2555
	CharCornerHDoubleNE rune = 0x2555
	// CharTeeDoubleW is Char 0x2563
	CharTeeDoubleW rune = 0x2563
	// CharWallDoubleV is Char 0x2551
	CharWallDoubleV rune = 0x2551
	// CharCornerDoubleNE is Char 0x2557
	CharCornerDoubleNE rune = 0x2557
	// CharCornerDoubleSE is Char 0x255D
	CharCornerDoubleSE rune = 0x255D
	// CharCornerVDoubleSE is Char 0x255C
	CharCornerVDoubleSE rune = 0x255C
	// CharCornerHDoubleSE is Char 0x255B
	CharCornerHDoubleSE rune = 0x255B
	// CharCornerNE is Char 0x2510
	CharCornerNE rune = 0x2510
	// CharCornerSW is Char 0x2514
	CharCornerSW rune = 0x2514
	// CharTeeN is Char 0x2534
	CharTeeN rune = 0x2534
	// CharTeeS is Char 0x252C
	CharTeeS rune = 0x252C
	// CharTeeE is Char 0x251C
	CharTeeE rune = 0x251C
	// CharWallH is Char 0x2500
	CharWallH rune = 0x2500
	// CharWallCross is Char 0x253C
	CharWallCross rune = 0x253C
	// CharTeeHDoubleE is Char 0x255E
	CharTeeHDoubleE rune = 0x255E
	// CharTeeVDoubleE is Char 0x255F
	CharTeeVDoubleE rune = 0x255F
	// CharCornerDoubleSW is Char 0x255A
	CharCornerDoubleSW rune = 0x255A
	// CharCornerDoubleNW is Char 0x2554
	CharCornerDoubleNW rune = 0x2554
	// CharTeeDoubleN is Char 0x2569
	CharTeeDoubleN rune = 0x2569
	// CharTeeDoubleS is Char 0x2566
	CharTeeDoubleS rune = 0x2566
	// CharTeeDoubleE is Char 0x2560
	CharTeeDoubleE rune = 0x2560
	// CharWallDoubleH is Char 0x2550
	CharWallDoubleH rune = 0x2550
	// CharWallDoubleCross is Char 0x256C
	CharWallDoubleCross rune = 0x256C
	// CharTeeHDoubleN is Char 0x2567
	CharTeeHDoubleN rune = 0x2567
	// CharTeeVDoubleN is Char 0x2568
	CharTeeVDoubleN rune = 0x2568
	// CharTeeHDoubleS is Char 0x2564
	CharTeeHDoubleS rune = 0x2564
	// CharTeeVDoubleS is Char 0x2565
	CharTeeVDoubleS rune = 0x2565
	// CharCornerVDoubleSW is Char 0x2559
	CharCornerVDoubleSW rune = 0x2559
	// CharCornerHdoubleSW is Char 0x2558
	CharCornerHdoubleSW rune = 0x2558
	// CharCornerHDoubleNW is Char 0x2552
	CharCornerHDoubleNW rune = 0x2552
	// CharCornerVDoubleNW is Char 0x2553
	CharCornerVDoubleNW rune = 0x2553
	// CharWallVDoubleCross is Char 0x256B
	CharWallVDoubleCross rune = 0x256B
	// CharWallHDoubleCross is Char 0x256A
	CharWallHDoubleCross rune = 0x256A
	// CharCornerSE is Char 0x2518
	CharCornerSE rune = 0x2518
	// CharCornerNW is Char 0x250C
	CharCornerNW rune = 0x250C
	// CharBlock4 is Char 0x2588
	CharBlock4 rune = 0x2588
	// CharSubpDown is Char 0x2584
	CharSubpDown rune = 0x2584
	// CharSubpLeft is Char 0x258C
	CharSubpLeft rune = 0x258C
	// CharSubpRight is Char 0x2590
	CharSubpRight rune = 0x2590
	// CharSubpUp is Char 0x2580
	CharSubpUp rune = 0x2580
	// CharAlpha is Char 0x03B1
	CharAlpha rune = 0x03B1
	// CharBeta is Char 0x00DF
	CharBeta rune = 0x00DF
	// CharGamma is Char 0x0393
	CharGamma rune = 0x0393
	// CharPi is Char 0x03C0
	CharPi rune = 0x03C0
	// CharSigmaUpper is Char 0x03A3
	CharSigmaUpper rune = 0x03A3
	// CharSigmaLower is Char 0x03C3
	CharSigmaLower rune = 0x03C3
	// CharMi is Char 0x00B5
	CharMi rune = 0x00B5
	// CharTau is Char 0x03C4
	CharTau rune = 0x03C4
	// CharPhiUpper is Char 0x03A6
	CharPhiUpper rune = 0x03A6
	// CharTheta is Char 0x0398
	CharTheta rune = 0x0398
	// CharOmega is Char 0x03A9
	CharOmega rune = 0x03A9
	// CharDelta is Char 0x03B4
	CharDelta rune = 0x03B4
	// CharInfinite is Char 0x221E
	CharInfinite rune = 0x221E
	// CharPhiLower is Char 0x03C6
	CharPhiLower rune = 0x03C6
	// CharEpsilon is Char 0x03B5
	CharEpsilon rune = 0x03B5
	// CharArc is Char 0x2229
	CharArc rune = 0x2229
	// CharList is Char 0x2261
	CharList rune = 0x2261
	// CharPlusMinus is Char 0x00B1
	CharPlusMinus rune = 0x00B1
	// CharGreaterEqual is Char 0x2265
	CharGreaterEqual rune = 0x2265
	// CharLowerEqual is Char 0x2264
	CharLowerEqual rune = 0x2264
	// CharIntegralUp is Char 0x2320
	CharIntegralUp rune = 0x2320
	// CharIntegralDown is Char 0x2321
	CharIntegralDown rune = 0x2321
	// CharDivision is Char 0x00F7
	CharDivision rune = 0x00F7
	// CharAlmostEqual is Char 0x2248
	CharAlmostEqual rune = 0x2248
	// CharDegree is Char 0x00B0
	CharDegree rune = 0x00B0
	// CharDot1 is Char 0x2219
	CharDot1 rune = 0x2219
	// CharDot2 is Char 0x00B7
	CharDot2 rune = 0x00B7
	// CharRoot is Char 0x221A
	CharRoot rune = 0x221A
	// CharExpN is Char 0x207F
	CharExpN rune = 0x207F
	// CharExp2 is Char 0x00B2
	CharExp2 rune = 0x00B2
	// CharSquare is Char 0x25A0
	CharSquare rune = 0x25A0
	// CharSquareInv is Char 0x00A0
	CharSquareInv rune = 0x00A0
)
