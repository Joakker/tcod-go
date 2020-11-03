package tcod

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

// BgFlag indicates how a background should be drawn
type BgFlag struct {
	flag C.TCOD_bkgnd_flag_t
}

// Renderer indicates which of the available internal renderers
// should be used
type Renderer struct {
	renderer C.TCOD_renderer_t
}

// Alignment indicates which alignment a console should have
type Alignment struct {
	alignment C.TCOD_alignment_t
}

// BgFlags follow
var (
	BgNone     BgFlag = BgFlag{flag: C.TCOD_BKGND_NONE}
	BgSet      BgFlag = BgFlag{flag: C.TCOD_BKGND_SET}
	BgMultiply BgFlag = BgFlag{flag: C.TCOD_BKGND_MULTIPLY}
	BgLighten  BgFlag = BgFlag{flag: C.TCOD_BKGND_LIGHTEN}
	BgDarken   BgFlag = BgFlag{flag: C.TCOD_BKGND_DARKEN}
	BgScreen   BgFlag = BgFlag{flag: C.TCOD_BKGND_SCREEN}
	BgAdd      BgFlag = BgFlag{flag: C.TCOD_BKGND_ADD}
	BgBurn     BgFlag = BgFlag{flag: C.TCOD_BKGND_BURN}
	BgOverlay  BgFlag = BgFlag{flag: C.TCOD_BKGND_OVERLAY}
	BgDefault  BgFlag = BgFlag{flag: C.TCOD_BKGND_DEFAULT}
)

// Character constants follow
var (
	CharHLine byte = C.TCOD_CHAR_HLINE
	CharVLine byte = C.TCOD_CHAR_VLINE
	CharNE    byte = C.TCOD_CHAR_NE
	CharNW    byte = C.TCOD_CHAR_NW
	CharSE    byte = C.TCOD_CHAR_SE
	CharSW    byte = C.TCOD_CHAR_SW
	CharTeeW  byte = C.TCOD_CHAR_TEEW
	CharTeeE  byte = C.TCOD_CHAR_TEEE
	CharTeeN  byte = C.TCOD_CHAR_TEEN
	CharTeeS  byte = C.TCOD_CHAR_TEES

	CharDHLine byte = C.TCOD_CHAR_DHLINE
	CharDVLine byte = C.TCOD_CHAR_DVLINE
	CharDNE    byte = C.TCOD_CHAR_DNE
	CharDNW    byte = C.TCOD_CHAR_DNW
	CharDSE    byte = C.TCOD_CHAR_DSE
	CharDSW    byte = C.TCOD_CHAR_DSW
	CharDTeeW  byte = C.TCOD_CHAR_DTEEW
	CharDTeeE  byte = C.TCOD_CHAR_DTEEE
	CharDTeeN  byte = C.TCOD_CHAR_DTEEN
	CharDTeeS  byte = C.TCOD_CHAR_DTEES

	CharCheckBoxSet   byte = C.TCOD_CHAR_CHECKBOX_SET
	CharCheckBoxUnset byte = C.TCOD_CHAR_CHECKBOX_UNSET
	CharRadioSet      byte = C.TCOD_CHAR_RADIO_SET
	CharRadioUnset    byte = C.TCOD_CHAR_RADIO_UNSET

	CharSubpNW   byte = C.TCOD_CHAR_SUBP_NW
	CharSubpNE   byte = C.TCOD_CHAR_SUBP_NE
	CharSubpSW   byte = C.TCOD_CHAR_SUBP_SW
	CharSubpSE   byte = C.TCOD_CHAR_SUBP_SE
	CharSubpDiag byte = C.TCOD_CHAR_SUBP_DIAG
	CharSubpN    byte = C.TCOD_CHAR_SUBP_N
	CharSubpE    byte = C.TCOD_CHAR_SUBP_E

	CharBlock1 byte = C.TCOD_CHAR_BLOCK1
	CharBlock2 byte = C.TCOD_CHAR_BLOCK2
	CharBlock3 byte = C.TCOD_CHAR_BLOCK3

	CharArrowN byte = C.TCOD_CHAR_ARROW_N
	CharArrowE byte = C.TCOD_CHAR_ARROW_E
	CharArrowS byte = C.TCOD_CHAR_ARROW_S
	CharArrowW byte = C.TCOD_CHAR_ARROW_W

	CharArrowNTN byte = C.TCOD_CHAR_ARROW2_N
	CharArrowNTE byte = C.TCOD_CHAR_ARROW2_E
	CharArrowNTS byte = C.TCOD_CHAR_ARROW2_S
	CharArrowNTW byte = C.TCOD_CHAR_ARROW2_W

	CharCross          byte = C.TCOD_CHAR_CROSS
	CharHeart          byte = C.TCOD_CHAR_HEART
	CharDiamond        byte = C.TCOD_CHAR_DIAMOND
	CharClub           byte = C.TCOD_CHAR_CLUB
	CharSpade          byte = C.TCOD_CHAR_SPADE
	CharMale           byte = C.TCOD_CHAR_MALE
	CharFemale         byte = C.TCOD_CHAR_FEMALE
	CharLight          byte = C.TCOD_CHAR_LIGHT
	CharPilcrow        byte = C.TCOD_CHAR_PILCROW
	CharSection        byte = C.TCOD_CHAR_SECTION
	CharPound          byte = C.TCOD_CHAR_POUND
	CharMultiplication byte = C.TCOD_CHAR_MULTIPLICATION
	CharFunction       byte = C.TCOD_CHAR_FUNCTION
	CharReserved       byte = C.TCOD_CHAR_RESERVED
	CharHalf           byte = C.TCOD_CHAR_HALF
	CharCent           byte = C.TCOD_CHAR_CENT
	CharYen            byte = C.TCOD_CHAR_YEN
	CharCurrency       byte = C.TCOD_CHAR_CURRENCY
	CharDivision       byte = C.TCOD_CHAR_DIVISION
	CharGrade          byte = C.TCOD_CHAR_GRADE
	CharUmlaut         byte = C.TCOD_CHAR_UMLAUT

	CharPow1 byte = C.TCOD_CHAR_POW1
	CharPow2 byte = C.TCOD_CHAR_POW2
	CharPow3 byte = C.TCOD_CHAR_POW3

	CharDArrowH byte = C.TCOD_CHAR_DARROW_H
	CharDArrowV byte = C.TCOD_CHAR_DARROW_V

	CharSmilie byte = C.TCOD_CHAR_SMILIE
	CharBullet byte = C.TCOD_CHAR_BULLET

	CharSmilieInv byte = C.TCOD_CHAR_SMILIE_INV
	CharBulletInv byte = C.TCOD_CHAR_BULLET_INV

	CharNote         byte = C.TCOD_CHAR_NOTE
	CharNoteDouble   byte = C.TCOD_CHAR_NOTE_DOUBLE
	CharOneQuarter   byte = C.TCOD_CHAR_ONE_QUARTER
	CharBulletSquare byte = C.TCOD_CHAR_BULLET_SQUARE
)

// Renderers follow
var (
	RenderOpenGL  Renderer = Renderer{renderer: C.TCOD_RENDERER_OPENGL}
	RenderOpenGL2 Renderer = Renderer{renderer: C.TCOD_RENDERER_OPENGL2}
	RenderSDL     Renderer = Renderer{renderer: C.TCOD_RENDERER_SDL}
	RenderSDL2    Renderer = Renderer{renderer: C.TCOD_RENDERER_SDL2}
)

var (
	AlignLeft   Alignment = Alignment{alignment: C.TCOD_LEFT}
	AlignRight  Alignment = Alignment{alignment: C.TCOD_RIGHT}
	AlignCenter Alignment = Alignment{alignment: C.TCOD_CENTER}
)
