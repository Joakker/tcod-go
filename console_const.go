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

var (
	// Background none
	BgNone BgFlag = BgFlag{flag: C.TCOD_BKGND_NONE}
	// Background set
	BgSet BgFlag = BgFlag{flag: C.TCOD_BKGND_SET}
	// Background multiply
	BgMultiply BgFlag = BgFlag{flag: C.TCOD_BKGND_MULTIPLY}
	// Background lighten
	BgLighten BgFlag = BgFlag{flag: C.TCOD_BKGND_LIGHTEN}
	// Background darken
	BgDarken BgFlag = BgFlag{flag: C.TCOD_BKGND_DARKEN}
	// Background screen
	BgScreen BgFlag = BgFlag{flag: C.TCOD_BKGND_SCREEN}
	// Background add
	BgAdd BgFlag = BgFlag{flag: C.TCOD_BKGND_ADD}
	// Background burn
	BgBurn BgFlag = BgFlag{flag: C.TCOD_BKGND_BURN}
	// Background overlay
	BgOverlay BgFlag = BgFlag{flag: C.TCOD_BKGND_OVERLAY}
	// Background default
	BgDefault BgFlag = BgFlag{flag: C.TCOD_BKGND_DEFAULT}
)

var (
	// Char Wall HLine
	CharHLine byte = C.TCOD_CHAR_HLINE
	// Char Wall VLine
	CharVLine byte = C.TCOD_CHAR_VLINE
	// Char Wall NE
	CharNE byte = C.TCOD_CHAR_NE
	// Char Wall NW
	CharNW byte = C.TCOD_CHAR_NW
	// Char Wall SE
	CharSE byte = C.TCOD_CHAR_SE
	// Char Wall SW
	CharSW byte = C.TCOD_CHAR_SW
	// Char Wall TeeW
	CharTeeW byte = C.TCOD_CHAR_TEEW
	// Char Wall TeeE
	CharTeeE byte = C.TCOD_CHAR_TEEE
	// Char Wall TeeN
	CharTeeN byte = C.TCOD_CHAR_TEEN
	// Char Wall TeeS
	CharTeeS byte = C.TCOD_CHAR_TEES

	// Char Double Wall HLine
	CharDHLine byte = C.TCOD_CHAR_DHLINE
	// Char Double Wall VLine
	CharDVLine byte = C.TCOD_CHAR_DVLINE
	// Char Double Wall NE
	CharDNE byte = C.TCOD_CHAR_DNE
	// Char Double Wall NW
	CharDNW byte = C.TCOD_CHAR_DNW
	// Char Double Wall SE
	CharDSE byte = C.TCOD_CHAR_DSE
	// Char Double Wall SW
	CharDSW byte = C.TCOD_CHAR_DSW
	// Char Double Wall TeeW
	CharDTeeW byte = C.TCOD_CHAR_DTEEW
	// Char Double Wall TeeE
	CharDTeeE byte = C.TCOD_CHAR_DTEEE
	// Char Double Wall TeeN
	CharDTeeN byte = C.TCOD_CHAR_DTEEN
	// Char Double Wall TeeS
	CharDTeeS byte = C.TCOD_CHAR_DTEES

	// Char CheckBox Set
	CharCheckBoxSet byte = C.TCOD_CHAR_CHECKBOX_SET
	// Char CheckBox Unset
	CharCheckBoxUnset byte = C.TCOD_CHAR_CHECKBOX_UNSET
	// Char Radio Set
	CharRadioSet byte = C.TCOD_CHAR_RADIO_SET
	// Char Radio Unset
	CharRadioUnset byte = C.TCOD_CHAR_RADIO_UNSET

	// Char Subpixel NW
	CharSubpNW byte = C.TCOD_CHAR_SUBP_NW
	// Char Subpixel NE
	CharSubpNE byte = C.TCOD_CHAR_SUBP_NE
	// Char Subpixel SW
	CharSubpSW byte = C.TCOD_CHAR_SUBP_SW
	// Char Subpixel SE
	CharSubpSE byte = C.TCOD_CHAR_SUBP_SE
	// Char Subpixel Diag
	CharSubpDiag byte = C.TCOD_CHAR_SUBP_DIAG
	// Char Subpixel N
	CharSubpN byte = C.TCOD_CHAR_SUBP_N
	// Char Subpixel E
	CharSubpE byte = C.TCOD_CHAR_SUBP_E

	// Char Block 1
	CharBlock1 byte = C.TCOD_CHAR_BLOCK1
	// Char Block 2
	CharBlock2 byte = C.TCOD_CHAR_BLOCK2
	// Char Block 3
	CharBlock3 byte = C.TCOD_CHAR_BLOCK3

	// Char Arrow N
	CharArrowN byte = C.TCOD_CHAR_ARROW_N
	// Char Arrow E
	CharArrowE byte = C.TCOD_CHAR_ARROW_E
	// Char Arrow S
	CharArrowS byte = C.TCOD_CHAR_ARROW_S
	// Char Arrow W
	CharArrowW byte = C.TCOD_CHAR_ARROW_W

	// Char Arrow No Tail N
	CharArrowNTN byte = C.TCOD_CHAR_ARROW2_N
	// Char Arrow No Tail E
	CharArrowNTE byte = C.TCOD_CHAR_ARROW2_E
	// Char Arrow No Tail S
	CharArrowNTS byte = C.TCOD_CHAR_ARROW2_S
	// Char Arrow No Tail W
	CharArrowNTW byte = C.TCOD_CHAR_ARROW2_W

	// Char Cross
	CharCross byte = C.TCOD_CHAR_CROSS
	// Char Heart
	CharHeart byte = C.TCOD_CHAR_HEART
	// Char Diamond
	CharDiamond byte = C.TCOD_CHAR_DIAMOND
	// Char Club
	CharClub byte = C.TCOD_CHAR_CLUB
	// Char Spade
	CharSpade byte = C.TCOD_CHAR_SPADE
	// Char Male
	CharMale byte = C.TCOD_CHAR_MALE
	// Char Female
	CharFemale byte = C.TCOD_CHAR_FEMALE
	// Char Light
	CharLight byte = C.TCOD_CHAR_LIGHT
	// Char Pilcrow
	CharPilcrow byte = C.TCOD_CHAR_PILCROW
	// Char Section
	CharSection byte = C.TCOD_CHAR_SECTION
	// Char Pound
	CharPound byte = C.TCOD_CHAR_POUND
	// Char Multiplication
	CharMultiplication byte = C.TCOD_CHAR_MULTIPLICATION
	// Char Function
	CharFunction byte = C.TCOD_CHAR_FUNCTION
	// Char Reserved
	CharReserved byte = C.TCOD_CHAR_RESERVED
	// Char Half
	CharHalf byte = C.TCOD_CHAR_HALF
	// Char Cent
	CharCent byte = C.TCOD_CHAR_CENT
	// Char Yen
	CharYen byte = C.TCOD_CHAR_YEN
	// Char Currency
	CharCurrency byte = C.TCOD_CHAR_CURRENCY
	// Char Division
	CharDivision byte = C.TCOD_CHAR_DIVISION
	// Char Grade
	CharGrade byte = C.TCOD_CHAR_GRADE
	// Char Umlaut
	CharUmlaut byte = C.TCOD_CHAR_UMLAUT

	// Char Pow 1
	CharPow1 byte = C.TCOD_CHAR_POW1
	// Char Pow 2
	CharPow2 byte = C.TCOD_CHAR_POW2
	// Char Pow 3
	CharPow3 byte = C.TCOD_CHAR_POW3

	// Char DArrowH
	CharDArrowH byte = C.TCOD_CHAR_DARROW_H
	// Char DArrowV
	CharDArrowV byte = C.TCOD_CHAR_DARROW_V

	// Char Smilie -
	CharSmilie byte = C.TCOD_CHAR_SMILIE
	// Char Bullet -
	CharBullet byte = C.TCOD_CHAR_BULLET

	// Char Smilie Inv
	CharSmilieInv byte = C.TCOD_CHAR_SMILIE_INV
	// Char Bullet Inv
	CharBulletInv byte = C.TCOD_CHAR_BULLET_INV

	// Char Note
	CharNote byte = C.TCOD_CHAR_NOTE
	// Char NoteDouble
	CharNoteDouble byte = C.TCOD_CHAR_NOTE_DOUBLE
	// Char OneQuarter
	CharOneQuarter byte = C.TCOD_CHAR_ONE_QUARTER
	// Char BulletSquare
	CharBulletSquare byte = C.TCOD_CHAR_BULLET_SQUARE
)

var (
	// Renderer OpenGL
	RenderOpenGL Renderer = Renderer{renderer: C.TCOD_RENDERER_OPENGL}
	// Renderer OpenGL2
	RenderOpenGL2 Renderer = Renderer{renderer: C.TCOD_RENDERER_OPENGL2}
	// Renderer SDL
	RenderSDL Renderer = Renderer{renderer: C.TCOD_RENDERER_SDL}
	// Renderer SDL2
	RenderSDL2 Renderer = Renderer{renderer: C.TCOD_RENDERER_SDL2}
)

var (
	// Alignment Left
	AlignLeft Alignment = Alignment{alignment: C.TCOD_LEFT}
	// Alignment Right
	AlignRight Alignment = Alignment{alignment: C.TCOD_RIGHT}
	// Alignment Center
	AlignCenter Alignment = Alignment{alignment: C.TCOD_CENTER}
)
