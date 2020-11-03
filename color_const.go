package tcod

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

// Color represents a colour with it's RGB or HSV values
type Color struct {
	color C.TCOD_color_t
}

var (

	// Color desaturated red
	DesaturatedRed Color = Color{color: C.TCOD_desaturated_red}
	// Color desaturated flame
	DesaturatedFlame Color = Color{color: C.TCOD_desaturated_flame}
	// Color desaturated orange
	DesaturatedOrange Color = Color{color: C.TCOD_desaturated_orange}
	// Color desaturated amber
	DesaturatedAmber Color = Color{color: C.TCOD_desaturated_amber}
	// Color desaturated yellow
	DesaturatedYellow Color = Color{color: C.TCOD_desaturated_yellow}
	// Color desaturated lime
	DesaturatedLime Color = Color{color: C.TCOD_desaturated_lime}
	// Color desaturated chartreuse
	DesaturatedChartreuse Color = Color{color: C.TCOD_desaturated_chartreuse}
	// Color desaturated green
	DesaturatedGreen Color = Color{color: C.TCOD_desaturated_green}
	// Color desaturated sea
	DesaturatedSea Color = Color{color: C.TCOD_desaturated_sea}
	// Color desaturated turquoise
	DesaturatedTurquoise Color = Color{color: C.TCOD_desaturated_turquoise}
	// Color desaturated cyan
	DesaturatedCyan Color = Color{color: C.TCOD_desaturated_cyan}
	// Color desaturated sky
	DesaturatedSky Color = Color{color: C.TCOD_desaturated_sky}
	// Color desaturated azure
	DesaturatedAzure Color = Color{color: C.TCOD_desaturated_azure}
	// Color desaturated blue
	DesaturatedBlue Color = Color{color: C.TCOD_desaturated_blue}
	// Color desaturated han
	DesaturatedHan Color = Color{color: C.TCOD_desaturated_han}
	// Color desaturated violet
	DesaturatedViolet Color = Color{color: C.TCOD_desaturated_violet}
	// Color desaturated purple
	DesaturatedPurple Color = Color{color: C.TCOD_desaturated_purple}
	// Color desaturated fuchsia
	DesaturatedFuchsia Color = Color{color: C.TCOD_desaturated_fuchsia}
	// Color desaturated magenta
	DesaturatedMagenta Color = Color{color: C.TCOD_desaturated_magenta}
	// Color desaturated pink
	DesaturatedPink Color = Color{color: C.TCOD_desaturated_pink}
	// Color desaturated crimson
	DesaturatedCrimson Color = Color{color: C.TCOD_desaturated_crimson}
	// Color lightest red
	LightestRed Color = Color{color: C.TCOD_lightest_red}
	// Color lightest flame
	LightestFlame Color = Color{color: C.TCOD_lightest_flame}
	// Color lightest orange
	LightestOrange Color = Color{color: C.TCOD_lightest_orange}
	// Color lightest amber
	LightestAmber Color = Color{color: C.TCOD_lightest_amber}
	// Color lightest yellow
	LightestYellow Color = Color{color: C.TCOD_lightest_yellow}
	// Color lightest lime
	LightestLime Color = Color{color: C.TCOD_lightest_lime}
	// Color lightest chartreuse
	LightestChartreuse Color = Color{color: C.TCOD_lightest_chartreuse}
	// Color lightest green
	LightestGreen Color = Color{color: C.TCOD_lightest_green}
	// Color lightest sea
	LightestSea Color = Color{color: C.TCOD_lightest_sea}
	// Color lightest turquoise
	LightestTurquoise Color = Color{color: C.TCOD_lightest_turquoise}
	// Color lightest cyan
	LightestCyan Color = Color{color: C.TCOD_lightest_cyan}
	// Color lightest sky
	LightestSky Color = Color{color: C.TCOD_lightest_sky}
	// Color lightest azure
	LightestAzure Color = Color{color: C.TCOD_lightest_azure}
	// Color lightest blue
	LightestBlue Color = Color{color: C.TCOD_lightest_blue}
	// Color lightest han
	LightestHan Color = Color{color: C.TCOD_lightest_han}
	// Color lightest violet
	LightestViolet Color = Color{color: C.TCOD_lightest_violet}
	// Color lightest purple
	LightestPurple Color = Color{color: C.TCOD_lightest_purple}
	// Color lightest fuchsia
	LightestFuchsia Color = Color{color: C.TCOD_lightest_fuchsia}
	// Color lightest magenta
	LightestMagenta Color = Color{color: C.TCOD_lightest_magenta}
	// Color lightest pink
	LightestPink Color = Color{color: C.TCOD_lightest_pink}
	// Color lightest crimson
	LightestCrimson Color = Color{color: C.TCOD_lightest_crimson}
	// Color lighter red
	LighterRed Color = Color{color: C.TCOD_lighter_red}
	// Color lighter flame
	LighterFlame Color = Color{color: C.TCOD_lighter_flame}
	// Color lighter orange
	LighterOrange Color = Color{color: C.TCOD_lighter_orange}
	// Color lighter amber
	LighterAmber Color = Color{color: C.TCOD_lighter_amber}
	// Color lighter yellow
	LighterYellow Color = Color{color: C.TCOD_lighter_yellow}
	// Color lighter lime
	LighterLime Color = Color{color: C.TCOD_lighter_lime}
	// Color lighter chartreuse
	LighterChartreuse Color = Color{color: C.TCOD_lighter_chartreuse}
	// Color lighter green
	LighterGreen Color = Color{color: C.TCOD_lighter_green}
	// Color lighter sea
	LighterSea Color = Color{color: C.TCOD_lighter_sea}
	// Color lighter turquoise
	LighterTurquoise Color = Color{color: C.TCOD_lighter_turquoise}
	// Color lighter cyan
	LighterCyan Color = Color{color: C.TCOD_lighter_cyan}
	// Color lighter sky
	LighterSky Color = Color{color: C.TCOD_lighter_sky}
	// Color lighter azure
	LighterAzure Color = Color{color: C.TCOD_lighter_azure}
	// Color lighter blue
	LighterBlue Color = Color{color: C.TCOD_lighter_blue}
	// Color lighter han
	LighterHan Color = Color{color: C.TCOD_lighter_han}
	// Color lighter violet
	LighterViolet Color = Color{color: C.TCOD_lighter_violet}
	// Color lighter purple
	LighterPurple Color = Color{color: C.TCOD_lighter_purple}
	// Color lighter fuchsia
	LighterFuchsia Color = Color{color: C.TCOD_lighter_fuchsia}
	// Color lighter magenta
	LighterMagenta Color = Color{color: C.TCOD_lighter_magenta}
	// Color lighter pink
	LighterPink Color = Color{color: C.TCOD_lighter_pink}
	// Color lighter crimson
	LighterCrimson Color = Color{color: C.TCOD_lighter_crimson}
	// Color light red
	LightRed Color = Color{color: C.TCOD_light_red}
	// Color light flame
	LightFlame Color = Color{color: C.TCOD_light_flame}
	// Color light orange
	LightOrange Color = Color{color: C.TCOD_light_orange}
	// Color light amber
	LightAmber Color = Color{color: C.TCOD_light_amber}
	// Color light yellow
	LightYellow Color = Color{color: C.TCOD_light_yellow}
	// Color light lime
	LightLime Color = Color{color: C.TCOD_light_lime}
	// Color light chartreuse
	LightChartreuse Color = Color{color: C.TCOD_light_chartreuse}
	// Color light green
	LightGreen Color = Color{color: C.TCOD_light_green}
	// Color light sea
	LightSea Color = Color{color: C.TCOD_light_sea}
	// Color light turquoise
	LightTurquoise Color = Color{color: C.TCOD_light_turquoise}
	// Color light cyan
	LightCyan Color = Color{color: C.TCOD_light_cyan}
	// Color light sky
	LightSky Color = Color{color: C.TCOD_light_sky}
	// Color light azure
	LightAzure Color = Color{color: C.TCOD_light_azure}
	// Color light blue
	LightBlue Color = Color{color: C.TCOD_light_blue}
	// Color light han
	LightHan Color = Color{color: C.TCOD_light_han}
	// Color light violet
	LightViolet Color = Color{color: C.TCOD_light_violet}
	// Color light purple
	LightPurple Color = Color{color: C.TCOD_light_purple}
	// Color light fuchsia
	LightFuchsia Color = Color{color: C.TCOD_light_fuchsia}
	// Color light magenta
	LightMagenta Color = Color{color: C.TCOD_light_magenta}
	// Color light pink
	LightPink Color = Color{color: C.TCOD_light_pink}
	// Color light crimson
	LightCrimson Color = Color{color: C.TCOD_light_crimson}
	// Color - red
	Red Color = Color{color: C.TCOD_red}
	// Color - flame
	Flame Color = Color{color: C.TCOD_flame}
	// Color - orange
	Orange Color = Color{color: C.TCOD_orange}
	// Color - amber
	Amber Color = Color{color: C.TCOD_amber}
	// Color - yellow
	Yellow Color = Color{color: C.TCOD_yellow}
	// Color - lime
	Lime Color = Color{color: C.TCOD_lime}
	// Color - chartreuse
	Chartreuse Color = Color{color: C.TCOD_chartreuse}
	// Color - green
	Green Color = Color{color: C.TCOD_green}
	// Color - sea
	Sea Color = Color{color: C.TCOD_sea}
	// Color - turquoise
	Turquoise Color = Color{color: C.TCOD_turquoise}
	// Color - cyan
	Cyan Color = Color{color: C.TCOD_cyan}
	// Color - sky
	Sky Color = Color{color: C.TCOD_sky}
	// Color - azure
	Azure Color = Color{color: C.TCOD_azure}
	// Color - blue
	Blue Color = Color{color: C.TCOD_blue}
	// Color - han
	Han Color = Color{color: C.TCOD_han}
	// Color - violet
	Violet Color = Color{color: C.TCOD_violet}
	// Color - purple
	Purple Color = Color{color: C.TCOD_purple}
	// Color - fuchsia
	Fuchsia Color = Color{color: C.TCOD_fuchsia}
	// Color - magenta
	Magenta Color = Color{color: C.TCOD_magenta}
	// Color - pink
	Pink Color = Color{color: C.TCOD_pink}
	// Color - crimson
	Crimson Color = Color{color: C.TCOD_crimson}
	// Color dark red
	DarkRed Color = Color{color: C.TCOD_dark_red}
	// Color dark flame
	DarkFlame Color = Color{color: C.TCOD_dark_flame}
	// Color dark orange
	DarkOrange Color = Color{color: C.TCOD_dark_orange}
	// Color dark amber
	DarkAmber Color = Color{color: C.TCOD_dark_amber}
	// Color dark yellow
	DarkYellow Color = Color{color: C.TCOD_dark_yellow}
	// Color dark lime
	DarkLime Color = Color{color: C.TCOD_dark_lime}
	// Color dark chartreuse
	DarkChartreuse Color = Color{color: C.TCOD_dark_chartreuse}
	// Color dark green
	DarkGreen Color = Color{color: C.TCOD_dark_green}
	// Color dark sea
	DarkSea Color = Color{color: C.TCOD_dark_sea}
	// Color dark turquoise
	DarkTurquoise Color = Color{color: C.TCOD_dark_turquoise}
	// Color dark cyan
	DarkCyan Color = Color{color: C.TCOD_dark_cyan}
	// Color dark sky
	DarkSky Color = Color{color: C.TCOD_dark_sky}
	// Color dark azure
	DarkAzure Color = Color{color: C.TCOD_dark_azure}
	// Color dark blue
	DarkBlue Color = Color{color: C.TCOD_dark_blue}
	// Color dark han
	DarkHan Color = Color{color: C.TCOD_dark_han}
	// Color dark violet
	DarkViolet Color = Color{color: C.TCOD_dark_violet}
	// Color dark purple
	DarkPurple Color = Color{color: C.TCOD_dark_purple}
	// Color dark fuchsia
	DarkFuchsia Color = Color{color: C.TCOD_dark_fuchsia}
	// Color dark magenta
	DarkMagenta Color = Color{color: C.TCOD_dark_magenta}
	// Color dark pink
	DarkPink Color = Color{color: C.TCOD_dark_pink}
	// Color dark crimson
	DarkCrimson Color = Color{color: C.TCOD_dark_crimson}
	// Color darker red
	DarkerRed Color = Color{color: C.TCOD_darker_red}
	// Color darker flame
	DarkerFlame Color = Color{color: C.TCOD_darker_flame}
	// Color darker orange
	DarkerOrange Color = Color{color: C.TCOD_darker_orange}
	// Color darker amber
	DarkerAmber Color = Color{color: C.TCOD_darker_amber}
	// Color darker yellow
	DarkerYellow Color = Color{color: C.TCOD_darker_yellow}
	// Color darker lime
	DarkerLime Color = Color{color: C.TCOD_darker_lime}
	// Color darker chartreuse
	DarkerChartreuse Color = Color{color: C.TCOD_darker_chartreuse}
	// Color darker green
	DarkerGreen Color = Color{color: C.TCOD_darker_green}
	// Color darker sea
	DarkerSea Color = Color{color: C.TCOD_darker_sea}
	// Color darker turquoise
	DarkerTurquoise Color = Color{color: C.TCOD_darker_turquoise}
	// Color darker cyan
	DarkerCyan Color = Color{color: C.TCOD_darker_cyan}
	// Color darker sky
	DarkerSky Color = Color{color: C.TCOD_darker_sky}
	// Color darker azure
	DarkerAzure Color = Color{color: C.TCOD_darker_azure}
	// Color darker blue
	DarkerBlue Color = Color{color: C.TCOD_darker_blue}
	// Color darker han
	DarkerHan Color = Color{color: C.TCOD_darker_han}
	// Color darker violet
	DarkerViolet Color = Color{color: C.TCOD_darker_violet}
	// Color darker purple
	DarkerPurple Color = Color{color: C.TCOD_darker_purple}
	// Color darker fuchsia
	DarkerFuchsia Color = Color{color: C.TCOD_darker_fuchsia}
	// Color darker magenta
	DarkerMagenta Color = Color{color: C.TCOD_darker_magenta}
	// Color darker pink
	DarkerPink Color = Color{color: C.TCOD_darker_pink}
	// Color darker crimson
	DarkerCrimson Color = Color{color: C.TCOD_darker_crimson}
	// Color darkest red
	DarkestRed Color = Color{color: C.TCOD_darkest_red}
	// Color darkest flame
	DarkestFlame Color = Color{color: C.TCOD_darkest_flame}
	// Color darkest orange
	DarkestOrange Color = Color{color: C.TCOD_darkest_orange}
	// Color darkest amber
	DarkestAmber Color = Color{color: C.TCOD_darkest_amber}
	// Color darkest yellow
	DarkestYellow Color = Color{color: C.TCOD_darkest_yellow}
	// Color darkest lime
	DarkestLime Color = Color{color: C.TCOD_darkest_lime}
	// Color darkest chartreuse
	DarkestChartreuse Color = Color{color: C.TCOD_darkest_chartreuse}
	// Color darkest green
	DarkestGreen Color = Color{color: C.TCOD_darkest_green}
	// Color darkest sea
	DarkestSea Color = Color{color: C.TCOD_darkest_sea}
	// Color darkest turquoise
	DarkestTurquoise Color = Color{color: C.TCOD_darkest_turquoise}
	// Color darkest cyan
	DarkestCyan Color = Color{color: C.TCOD_darkest_cyan}
	// Color darkest sky
	DarkestSky Color = Color{color: C.TCOD_darkest_sky}
	// Color darkest azure
	DarkestAzure Color = Color{color: C.TCOD_darkest_azure}
	// Color darkest blue
	DarkestBlue Color = Color{color: C.TCOD_darkest_blue}
	// Color darkest han
	DarkestHan Color = Color{color: C.TCOD_darkest_han}
	// Color darkest violet
	DarkestViolet Color = Color{color: C.TCOD_darkest_violet}
	// Color darkest purple
	DarkestPurple Color = Color{color: C.TCOD_darkest_purple}
	// Color darkest fuchsia
	DarkestFuchsia Color = Color{color: C.TCOD_darkest_fuchsia}
	// Color darkest magenta
	DarkestMagenta Color = Color{color: C.TCOD_darkest_magenta}
	// Color darkest pink
	DarkestPink Color = Color{color: C.TCOD_darkest_pink}
	// Color darkest crimson
	DarkestCrimson Color = Color{color: C.TCOD_darkest_crimson}

	// Color brass
	Brass Color = Color{color: C.TCOD_brass}
	// Color copper
	Copper Color = Color{color: C.TCOD_copper}
	// Color gold
	Gold Color = Color{color: C.TCOD_gold}
	// Color silver
	Silver Color = Color{color: C.TCOD_silver}
	// Color celadon
	Celadon Color = Color{color: C.TCOD_celadon}
	// Color peach
	Peach Color = Color{color: C.TCOD_peach}
	// Color grey
	Grey Color = Color{color: C.TCOD_grey}
	// Color sepia
	Sepia Color = Color{color: C.TCOD_sepia}
	// Color black
	Black Color = Color{color: C.TCOD_black}
	// Color white
	White Color = Color{color: C.TCOD_white}
)
