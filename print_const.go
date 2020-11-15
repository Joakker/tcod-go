package tcod

// #include <libtcod.h>
import "C"

var (
	// CharHLine is char wall HLine
	CharHLine byte = C.TCOD_CHAR_HLINE
	// CharVLine is char wall VLine
	CharVLine byte = C.TCOD_CHAR_VLINE
	// CharNE is char wall NE
	CharNE byte = C.TCOD_CHAR_NE
	// CharNW is char wall NW
	CharNW byte = C.TCOD_CHAR_NW
	// CharSE is char wall SE
	CharSE byte = C.TCOD_CHAR_SE
	// CharSW is char wall SW
	CharSW byte = C.TCOD_CHAR_SW
	// CharTeeW is char wall TeeW
	CharTeeW byte = C.TCOD_CHAR_TEEW
	// CharTeeE is char wall TeeE
	CharTeeE byte = C.TCOD_CHAR_TEEE
	// CharTeeN is char wall TeeN
	CharTeeN byte = C.TCOD_CHAR_TEEN
	// CharTeeS is char wall TeeS
	CharTeeS byte = C.TCOD_CHAR_TEES

	// CharDHLine is char double wall HLine
	CharDHLine byte = C.TCOD_CHAR_DHLINE
	// CharDVLine is char double wall VLine
	CharDVLine byte = C.TCOD_CHAR_DVLINE
	// CharDNE is char double wall NE
	CharDNE byte = C.TCOD_CHAR_DNE
	// CharDNW is char double wall NW
	CharDNW byte = C.TCOD_CHAR_DNW
	// CharDSE is char double wall SE
	CharDSE byte = C.TCOD_CHAR_DSE
	// CharDSW is char double wall SW
	CharDSW byte = C.TCOD_CHAR_DSW
	// CharDTeeW is char double wall TeeW
	CharDTeeW byte = C.TCOD_CHAR_DTEEW
	// CharDTeeE is char double wall TeeE
	CharDTeeE byte = C.TCOD_CHAR_DTEEE
	// CharDTeeN is char double wall TeeN
	CharDTeeN byte = C.TCOD_CHAR_DTEEN
	// CharDTeeS is char double wall TeeS
	CharDTeeS byte = C.TCOD_CHAR_DTEES

	// CharCheckBoxSet is char CheckBox Set
	CharCheckBoxSet byte = C.TCOD_CHAR_CHECKBOX_SET
	// CharCheckBoxUnset is char CheckBox Unset
	CharCheckBoxUnset byte = C.TCOD_CHAR_CHECKBOX_UNSET
	// CharRadioSet is char Radio Set
	CharRadioSet byte = C.TCOD_CHAR_RADIO_SET
	// CharRadioUnset is char Radio Unset
	CharRadioUnset byte = C.TCOD_CHAR_RADIO_UNSET

	// CharSubpNW char subpixel NW
	CharSubpNW byte = C.TCOD_CHAR_SUBP_NW
	// CharSubpNE char subpixel NE
	CharSubpNE byte = C.TCOD_CHAR_SUBP_NE
	// CharSubpSW char subpixel SW
	CharSubpSW byte = C.TCOD_CHAR_SUBP_SW
	// CharSubpSE char subpixel SE
	CharSubpSE byte = C.TCOD_CHAR_SUBP_SE
	// CharSubpDiag char subpixel Diag
	CharSubpDiag byte = C.TCOD_CHAR_SUBP_DIAG
	// CharSubpN char subpixel N
	CharSubpN byte = C.TCOD_CHAR_SUBP_N
	// CharSubpE char subpixel E
	CharSubpE byte = C.TCOD_CHAR_SUBP_E

	// CharBlock1 char block 1
	CharBlock1 byte = C.TCOD_CHAR_BLOCK1
	// CharBlock2 char block 2
	CharBlock2 byte = C.TCOD_CHAR_BLOCK2
	// CharBlock3 char block 3
	CharBlock3 byte = C.TCOD_CHAR_BLOCK3

	// CharArrowN char arrow N
	CharArrowN byte = C.TCOD_CHAR_ARROW_N
	// CharArrowE char arrow E
	CharArrowE byte = C.TCOD_CHAR_ARROW_E
	// CharArrowS char arrow S
	CharArrowS byte = C.TCOD_CHAR_ARROW_S
	// CharArrowW char arrow W
	CharArrowW byte = C.TCOD_CHAR_ARROW_W

	// CharArrowNTN is char arrow no tail N
	CharArrowNTN byte = C.TCOD_CHAR_ARROW2_N
	// CharArrowNTE is char arrow no tail E
	CharArrowNTE byte = C.TCOD_CHAR_ARROW2_E
	// CharArrowNTS is char arrow no tail S
	CharArrowNTS byte = C.TCOD_CHAR_ARROW2_S
	// CharArrowNTW is char arrow no tail W
	CharArrowNTW byte = C.TCOD_CHAR_ARROW2_W

	// CharCross is char Cross
	CharCross byte = C.TCOD_CHAR_CROSS
	// CharHeart is char Heart
	CharHeart byte = C.TCOD_CHAR_HEART
	// CharDiamond is char Diamond
	CharDiamond byte = C.TCOD_CHAR_DIAMOND
	// CharClub is char Club
	CharClub byte = C.TCOD_CHAR_CLUB
	// CharSpade is char Spade
	CharSpade byte = C.TCOD_CHAR_SPADE
	// CharMale is char Male
	CharMale byte = C.TCOD_CHAR_MALE
	// CharFemale is char Female
	CharFemale byte = C.TCOD_CHAR_FEMALE
	// CharLight is char Light
	CharLight byte = C.TCOD_CHAR_LIGHT
	// CharPilcrow is char Pilcrow
	CharPilcrow byte = C.TCOD_CHAR_PILCROW
	// CharSection is char Section
	CharSection byte = C.TCOD_CHAR_SECTION
	// CharPound is char Pound
	CharPound byte = C.TCOD_CHAR_POUND
	// CharMultiplication is char Multiplication
	CharMultiplication byte = C.TCOD_CHAR_MULTIPLICATION
	// CharFunction is char Function
	CharFunction byte = C.TCOD_CHAR_FUNCTION
	// CharReserved is char Reserved
	CharReserved byte = C.TCOD_CHAR_RESERVED
	// CharHalf is char Half
	CharHalf byte = C.TCOD_CHAR_HALF
	// CharCent is char Cent
	CharCent byte = C.TCOD_CHAR_CENT
	// CharYen is char Yen
	CharYen byte = C.TCOD_CHAR_YEN
	// CharCurrency is char Currency
	CharCurrency byte = C.TCOD_CHAR_CURRENCY
	// CharDivision is char Division
	CharDivision byte = C.TCOD_CHAR_DIVISION
	// CharGrade is char Grade
	CharGrade byte = C.TCOD_CHAR_GRADE
	// CharUmlaut is char Umlaut
	CharUmlaut byte = C.TCOD_CHAR_UMLAUT

	// CharPow1 Char Pow 1
	CharPow1 byte = C.TCOD_CHAR_POW1
	// CharPow2 Char Pow 2
	CharPow2 byte = C.TCOD_CHAR_POW2
	// CharPow3 Char Pow 3
	CharPow3 byte = C.TCOD_CHAR_POW3

	// CharDArrowH Char DArrowH
	CharDArrowH byte = C.TCOD_CHAR_DARROW_H
	// CharDArrowV Char DArrowV
	CharDArrowV byte = C.TCOD_CHAR_DARROW_V

	// CharSmilie is char Smilie -
	CharSmilie byte = C.TCOD_CHAR_SMILIE
	// CharBullet is char Bullet -
	CharBullet byte = C.TCOD_CHAR_BULLET

	// CharSmilieInv is char Smilie Inv
	CharSmilieInv byte = C.TCOD_CHAR_SMILIE_INV
	// CharBulletInv is char Bullet Inv
	CharBulletInv byte = C.TCOD_CHAR_BULLET_INV

	// CharNote Char Note
	CharNote byte = C.TCOD_CHAR_NOTE
	// CharNoteDouble Char NoteDouble
	CharNoteDouble byte = C.TCOD_CHAR_NOTE_DOUBLE
	// CharOneQuarter Char OneQuarter
	CharOneQuarter byte = C.TCOD_CHAR_ONE_QUARTER
	// CharBulletSquare Char BulletSquare
	CharBulletSquare byte = C.TCOD_CHAR_BULLET_SQUARE
)
