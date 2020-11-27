package tcod

// #include <libtcod.h>
import "C"

// byte enum
var (
	// CharSmiley Smiley
	CharSmiley byte = 1
	// CharSmileyInv SmileyInv
	CharSmileyInv byte = 2
	// CharHeart Heart
	CharHeart byte = 3
	// CharDiamon Diamon
	CharDiamon byte = 4
	// CharClub Club
	CharClub byte = 5
	// CharSpade Spade
	CharSpade byte = 6
	// CharBullet Bullet
	CharBullet byte = 7
	// CharBulletActive BulletActive
	CharBulletActive byte = 8
	// CharBulletInv BulletInv
	CharBulletInv byte = 9
	// CharBulletInvActive BulletInvActive
	CharBulletInvActive byte = 10
	// CharMale Male
	CharMale byte = 11
	// CharFemale Female
	CharFemale byte = 12
	// CharNote Note
	CharNote byte = 13
	// CharDoubleNote DoubleNote
	CharDoubleNote byte = 14
	// CharStar Star
	CharStar byte = 15
	// CharTriangleArrowR TriangleArrowR
	CharTriangleArrowR byte = 16
	// CharTriangleArrowL TriangleArrowL
	CharTriangleArrowL byte = 17
	// CharDoubleArrowV DoubleArrowV
	CharDoubleArrowV byte = 18
	// CharDoubleExclamation DoubleExclamation
	CharDoubleExclamation byte = 19
	// CharLamp Lamp
	CharLamp byte = 21
	// CharUnderBlock UnderBlock
	CharUnderBlock byte = 22
	// CharDoubleArrowVertUnderscore DoubleArrowVertUnderscore
	CharDoubleArrowVertUnderscore byte = 23
	// CharArrowUp ArrowUp
	CharArrowUp byte = 24
	// CharArrowDown ArrowDown
	CharArrowDown byte = 25
	// CharArrowRight ArrowRight
	CharArrowRight byte = 26
	// CharArrowLeft ArrowLeft
	CharArrowLeft byte = 27
	// CharRightAngle RightAngle
	CharRightAngle byte = 28
	// CharDoubleArrowH DoubleArrowH
	CharDoubleArrowH byte = 29
	// CharTriangleArrowU TriangleArrowU
	CharTriangleArrowU byte = 30
	// CharTriangleArrowD TriangleArrowD
	CharTriangleArrowD byte = 31
	// CharTick Tick
	CharTick byte = 39
	// CharPilcrow Pilcrow
	CharPilcrow byte = 94
	// CharHouse House
	CharHouse byte = 127
	// CharUpperCedilla UpperCedilla
	CharUpperCedilla byte = 128
	// CharULowermlaut ULowermlaut
	CharULowermlaut byte = 129
	// CharELowerAcute ELowerAcute
	CharELowerAcute byte = 130
	// CharALowerCaret ALowerCaret
	CharALowerCaret byte = 131
	// CharALowerUmlaut ALowerUmlaut
	CharALowerUmlaut byte = 132
	// CharALowerGrave ALowerGrave
	CharALowerGrave byte = 133
	// CharALowerCircle ALowerCircle
	CharALowerCircle byte = 134
	// CharLowerCedilla LowerCedilla
	CharLowerCedilla byte = 135
	// CharELowerCaret ELowerCaret
	CharELowerCaret byte = 136
	// CharELowerUmlaut ELowerUmlaut
	CharELowerUmlaut byte = 137
	// CharELowerGrave ELowerGrave
	CharELowerGrave byte = 138
	// CharILowerUmlaut ILowerUmlaut
	CharILowerUmlaut byte = 139
	// CharILowerCaret ILowerCaret
	CharILowerCaret byte = 140
	// CharILowerGrave ILowerGrave
	CharILowerGrave byte = 141
	// CharAUpperUmlaut AUpperUmlaut
	CharAUpperUmlaut byte = 142
	// CharAUpperCircle AUpperCircle
	CharAUpperCircle byte = 143
	// CharEUpperAcute EUpperAcute
	CharEUpperAcute byte = 144
	// CharAELower AELower
	CharAELower byte = 145
	// CharAEUpper AEUpper
	CharAEUpper byte = 146
	// CharOLowerCaret OLowerCaret
	CharOLowerCaret byte = 147
	// CharOLowerUmlaut OLowerUmlaut
	CharOLowerUmlaut byte = 148
	// CharOLowerGrave OLowerGrave
	CharOLowerGrave byte = 149
	// CharULowerCaret ULowerCaret
	CharULowerCaret byte = 150
	// CharULowerGrave ULowerGrave
	CharULowerGrave byte = 151
	// CharYLowerUmlaut YLowerUmlaut
	CharYLowerUmlaut byte = 152
	// CharOUpperUmlaut OUpperUmlaut
	CharOUpperUmlaut byte = 153
	// CharUUpperUmlaut UUpperUmlaut
	CharUUpperUmlaut byte = 154
	// CharCent Cent
	CharCent byte = 155
	// CharPound Pound
	CharPound byte = 156
	// CharYen Yen
	CharYen byte = 157
	// CharPoint Point
	CharPoint byte = 158
	// CharFunction Function
	CharFunction byte = 159
	// CharALowerAcute ALowerAcute
	CharALowerAcute byte = 160
	// CharILowerAccute ILowerAccute
	CharILowerAccute byte = 161
	// CharOLowerAccute OLowerAccute
	CharOLowerAccute byte = 162
	// CharULowerAccute ULowerAccute
	CharULowerAccute byte = 163
	// CharNTildeLower NTildeLower
	CharNTildeLower byte = 164
	// CharNTildeUpper NTildeUpper
	CharNTildeUpper byte = 165
	// CharExpA ExpA
	CharExpA byte = 166
	// CharExpO ExpO
	CharExpO byte = 167
	// CharInvQuestion InvQuestion
	CharInvQuestion byte = 168
	// CharLogicalNot LogicalNot
	CharLogicalNot byte = 169
	// CharLogicalNotInv LogicalNotInv
	CharLogicalNotInv byte = 170
	// CharHalf Half
	CharHalf byte = 171
	// CharQuarter Quarter
	CharQuarter byte = 172
	// CharInvExclamation InvExclamation
	CharInvExclamation byte = 173
	// CharDoubleArrowL DoubleArrowL
	CharDoubleArrowL byte = 174
	// CharDoubleArrowR DoubleArrowR
	CharDoubleArrowR byte = 175
	// CharBlock1 Block1
	CharBlock1 byte = 176
	// CharBlock2 Block2
	CharBlock2 byte = 177
	// CharBlock3 Block3
	CharBlock3 byte = 178
	// CharWallV WallV
	CharWallV byte = 179
	// CharTeeW TeeW
	CharTeeW byte = 180
	// CharTeeHDoubleW TeeHDoubleW
	CharTeeHDoubleW byte = 181
	// CharTeeVDoubleW TeeVDoubleW
	CharTeeVDoubleW byte = 182
	// CharCornerVDoubleNE CornerVDoubleNE
	CharCornerVDoubleNE byte = 183
	// CharCornerHDoubleNE CornerHDoubleNE
	CharCornerHDoubleNE byte = 184
	// CharTeeDoubleW TeeDoubleW
	CharTeeDoubleW byte = 185
	// CharWallDoubleV WallDoubleV
	CharWallDoubleV byte = 186
	// CharCornerDoubleNE CornerDoubleNE
	CharCornerDoubleNE byte = 187
	// CharCornerDoubleSE CornerDoubleSE
	CharCornerDoubleSE byte = 188
	// CharCornerVDoubleSE CornerVDoubleSE
	CharCornerVDoubleSE byte = 189
	// CharCornerHDoubleSE CornerHDoubleSE
	CharCornerHDoubleSE byte = 190
	// CharCornerNE CornerNE
	CharCornerNE byte = 191
	// CharCornerSW CornerSW
	CharCornerSW byte = 192
	// CharTeeN TeeN
	CharTeeN byte = 193
	// CharTeeS TeeS
	CharTeeS byte = 194
	// CharTeeE TeeE
	CharTeeE byte = 195
	// CharWallH WallH
	CharWallH byte = 196
	// CharWallCross WallCross
	CharWallCross byte = 197
	// CharTeeHDoubleE TeeHDoubleE
	CharTeeHDoubleE byte = 198
	// CharTeeVDoubleE TeeVDoubleE
	CharTeeVDoubleE byte = 199
	// CharCornerDoubleSW CornerDoubleSW
	CharCornerDoubleSW byte = 200
	// CharCornerDoubleNW CornerDoubleNW
	CharCornerDoubleNW byte = 201
	// CharTeeDoubleN TeeDoubleN
	CharTeeDoubleN byte = 202
	// CharTeeDoubleS TeeDoubleS
	CharTeeDoubleS byte = 203
	// CharTeeDoubleE TeeDoubleE
	CharTeeDoubleE byte = 204
	// CharWallDoubleH WallDoubleH
	CharWallDoubleH byte = 205
	// CharWallDoubleCross WallDoubleCross
	CharWallDoubleCross byte = 206
	// CharTeeHDoubleN TeeHDoubleN
	CharTeeHDoubleN byte = 207
	// CharTeeVDoubleN TeeVDoubleN
	CharTeeVDoubleN byte = 208
	// CharTeeHDoubleS TeeHDoubleS
	CharTeeHDoubleS byte = 209
	// CharTeeVDoubleS TeeVDoubleS
	CharTeeVDoubleS byte = 210
	// CharCornerVDoubleSW CornerVDoubleSW
	CharCornerVDoubleSW byte = 211
	// CharCornerHdoubleSW CornerHdoubleSW
	CharCornerHdoubleSW byte = 212
	// CharCornerHDoubleNW CornerHDoubleNW
	CharCornerHDoubleNW byte = 213
	// CharCornerVDoubleNW CornerVDoubleNW
	CharCornerVDoubleNW byte = 214
	// CharWallVDoubleCross WallVDoubleCross
	CharWallVDoubleCross byte = 215
	// CharWallHDoubleCross WallHDoubleCross
	CharWallHDoubleCross byte = 216
	// CharCornerSE CornerSE
	CharCornerSE byte = 217
	// CharCornerNW CornerNW
	CharCornerNW byte = 218
	// CharBlock4 Block4
	CharBlock4 byte = 219
	// CharSubpDown SubpDown
	CharSubpDown byte = 220
	// CharSubpLeft SubpLeft
	CharSubpLeft byte = 221
	// CharSubpRight SubpRight
	CharSubpRight byte = 222
	// CharSubpUp SubpUp
	CharSubpUp byte = 223
	// CharAlpha Alpha
	CharAlpha byte = 224
	// CharBeta Beta
	CharBeta byte = 225
	// CharGamma Gamma
	CharGamma byte = 226
	// CharPi Pi
	CharPi byte = 227
	// CharSigmaUpper SigmaUpper
	CharSigmaUpper byte = 228
	// CharSigmaLower SigmaLower
	CharSigmaLower byte = 229
	// CharMi Mi
	CharMi byte = 230
	// CharTau Tau
	CharTau byte = 231
	// CharPhiUpper PhiUpper
	CharPhiUpper byte = 232
	// CharTheta Theta
	CharTheta byte = 233
	// CharOmega Omega
	CharOmega byte = 234
	// CharDelta Delta
	CharDelta byte = 235
	// CharInfinite Infinite
	CharInfinite byte = 236
	// CharPhiLower PhiLower
	CharPhiLower byte = 237
	// CharEpsilon Epsilon
	CharEpsilon byte = 238
	// CharArc Arc
	CharArc byte = 239
	// CharList List
	CharList byte = 240
	// CharPlusMinus PlusMinus
	CharPlusMinus byte = 241
	// CharGreaterEqual GreaterEqual
	CharGreaterEqual byte = 242
	// CharLowerEqual LowerEqual
	CharLowerEqual byte = 243
	// CharIntegralUp IntegralUp
	CharIntegralUp byte = 244
	// CharIntegralDown IntegralDown
	CharIntegralDown byte = 245
	// CharDivision Division
	CharDivision byte = 246
	// CharAlmostEqual AlmostEqual
	CharAlmostEqual byte = 247
	// CharDegree Degree
	CharDegree byte = 248
	// CharDot1 Dot1
	CharDot1 byte = 249
	// CharDot2 Dot2
	CharDot2 byte = 250
	// CharRoot Root
	CharRoot byte = 251
	// CharExpN ExpN
	CharExpN byte = 252
	// CharExp2 Exp2
	CharExp2 byte = 253
	// CharSquare Square
	CharSquare byte = 254
	// CharSquareInv SquareInv
	CharSquareInv byte = 255
)
