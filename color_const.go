package tcod

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

// Color represents a colour with it's RGB or HSV values
type Color struct {
	color C.TCOD_color_t
}

var (

	// DesaturatedRed is color desaturated red
	DesaturatedRed Color = Color{color: C.TCOD_desaturated_red}
	// DesaturatedFlame is color desaturated flame
	DesaturatedFlame Color = Color{color: C.TCOD_desaturated_flame}
	// DesaturatedOrange is color desaturated orange
	DesaturatedOrange Color = Color{color: C.TCOD_desaturated_orange}
	// DesaturatedAmber is color desaturated amber
	DesaturatedAmber Color = Color{color: C.TCOD_desaturated_amber}
	// DesaturatedYellow is color desaturated yellow
	DesaturatedYellow Color = Color{color: C.TCOD_desaturated_yellow}
	// DesaturatedLime is color desaturated lime
	DesaturatedLime Color = Color{color: C.TCOD_desaturated_lime}
	// DesaturatedChartreuse is color desaturated chartreuse
	DesaturatedChartreuse Color = Color{color: C.TCOD_desaturated_chartreuse}
	// DesaturatedGreen is color desaturated green
	DesaturatedGreen Color = Color{color: C.TCOD_desaturated_green}
	// DesaturatedSea is color desaturated sea
	DesaturatedSea Color = Color{color: C.TCOD_desaturated_sea}
	// DesaturatedTurquoise is color desaturated turquoise
	DesaturatedTurquoise Color = Color{color: C.TCOD_desaturated_turquoise}
	// DesaturatedCyan is color desaturated cyan
	DesaturatedCyan Color = Color{color: C.TCOD_desaturated_cyan}
	// DesaturatedSky is color desaturated sky
	DesaturatedSky Color = Color{color: C.TCOD_desaturated_sky}
	// DesaturatedAzure is color desaturated azure
	DesaturatedAzure Color = Color{color: C.TCOD_desaturated_azure}
	// DesaturatedBlue is color desaturated blue
	DesaturatedBlue Color = Color{color: C.TCOD_desaturated_blue}
	// DesaturatedHan is color desaturated han
	DesaturatedHan Color = Color{color: C.TCOD_desaturated_han}
	// DesaturatedViolet is color desaturated violet
	DesaturatedViolet Color = Color{color: C.TCOD_desaturated_violet}
	// DesaturatedPurple is color desaturated purple
	DesaturatedPurple Color = Color{color: C.TCOD_desaturated_purple}
	// DesaturatedFuchsia is color desaturated fuchsia
	DesaturatedFuchsia Color = Color{color: C.TCOD_desaturated_fuchsia}
	// DesaturatedMagenta is color desaturated magenta
	DesaturatedMagenta Color = Color{color: C.TCOD_desaturated_magenta}
	// DesaturatedPink is color desaturated pink
	DesaturatedPink Color = Color{color: C.TCOD_desaturated_pink}
	// DesaturatedCrimson is color desaturated crimson
	DesaturatedCrimson Color = Color{color: C.TCOD_desaturated_crimson}
	// LightestRed is color lightest red
	LightestRed Color = Color{color: C.TCOD_lightest_red}
	// LightestFlame is color lightest flame
	LightestFlame Color = Color{color: C.TCOD_lightest_flame}
	// LightestOrange is color lightest orange
	LightestOrange Color = Color{color: C.TCOD_lightest_orange}
	// LightestAmber is color lightest amber
	LightestAmber Color = Color{color: C.TCOD_lightest_amber}
	// LightestYellow is color lightest yellow
	LightestYellow Color = Color{color: C.TCOD_lightest_yellow}
	// LightestLime is color lightest lime
	LightestLime Color = Color{color: C.TCOD_lightest_lime}
	// LightestChartreuse is color lightest chartreuse
	LightestChartreuse Color = Color{color: C.TCOD_lightest_chartreuse}
	// LightestGreen is color lightest green
	LightestGreen Color = Color{color: C.TCOD_lightest_green}
	// LightestSea is color lightest sea
	LightestSea Color = Color{color: C.TCOD_lightest_sea}
	// LightestTurquoise is color lightest turquoise
	LightestTurquoise Color = Color{color: C.TCOD_lightest_turquoise}
	// LightestCyan is color lightest cyan
	LightestCyan Color = Color{color: C.TCOD_lightest_cyan}
	// LightestSky is color lightest sky
	LightestSky Color = Color{color: C.TCOD_lightest_sky}
	// LightestAzure is color lightest azure
	LightestAzure Color = Color{color: C.TCOD_lightest_azure}
	// LightestBlue is color lightest blue
	LightestBlue Color = Color{color: C.TCOD_lightest_blue}
	// LightestHan is color lightest han
	LightestHan Color = Color{color: C.TCOD_lightest_han}
	// LightestViolet is color lightest violet
	LightestViolet Color = Color{color: C.TCOD_lightest_violet}
	// LightestPurple is color lightest purple
	LightestPurple Color = Color{color: C.TCOD_lightest_purple}
	// LightestFuchsia is color lightest fuchsia
	LightestFuchsia Color = Color{color: C.TCOD_lightest_fuchsia}
	// LightestMagenta is color lightest magenta
	LightestMagenta Color = Color{color: C.TCOD_lightest_magenta}
	// LightestPink is color lightest pink
	LightestPink Color = Color{color: C.TCOD_lightest_pink}
	// LightestCrimson is color lightest crimson
	LightestCrimson Color = Color{color: C.TCOD_lightest_crimson}
	// LighterRed is color lighter red
	LighterRed Color = Color{color: C.TCOD_lighter_red}
	// LighterFlame is color lighter flame
	LighterFlame Color = Color{color: C.TCOD_lighter_flame}
	// LighterOrange is color lighter orange
	LighterOrange Color = Color{color: C.TCOD_lighter_orange}
	// LighterAmber is color lighter amber
	LighterAmber Color = Color{color: C.TCOD_lighter_amber}
	// LighterYellow is color lighter yellow
	LighterYellow Color = Color{color: C.TCOD_lighter_yellow}
	// LighterLime is color lighter lime
	LighterLime Color = Color{color: C.TCOD_lighter_lime}
	// LighterChartreuse is color lighter chartreuse
	LighterChartreuse Color = Color{color: C.TCOD_lighter_chartreuse}
	// LighterGreen is color lighter green
	LighterGreen Color = Color{color: C.TCOD_lighter_green}
	// LighterSea is color lighter sea
	LighterSea Color = Color{color: C.TCOD_lighter_sea}
	// LighterTurquoise is color lighter turquoise
	LighterTurquoise Color = Color{color: C.TCOD_lighter_turquoise}
	// LighterCyan is color lighter cyan
	LighterCyan Color = Color{color: C.TCOD_lighter_cyan}
	// LighterSky is color lighter sky
	LighterSky Color = Color{color: C.TCOD_lighter_sky}
	// LighterAzure is color lighter azure
	LighterAzure Color = Color{color: C.TCOD_lighter_azure}
	// LighterBlue is color lighter blue
	LighterBlue Color = Color{color: C.TCOD_lighter_blue}
	// LighterHan is color lighter han
	LighterHan Color = Color{color: C.TCOD_lighter_han}
	// LighterViolet is color lighter violet
	LighterViolet Color = Color{color: C.TCOD_lighter_violet}
	// LighterPurple is color lighter purple
	LighterPurple Color = Color{color: C.TCOD_lighter_purple}
	// LighterFuchsia is color lighter fuchsia
	LighterFuchsia Color = Color{color: C.TCOD_lighter_fuchsia}
	// LighterMagenta is color lighter magenta
	LighterMagenta Color = Color{color: C.TCOD_lighter_magenta}
	// LighterPink is color lighter pink
	LighterPink Color = Color{color: C.TCOD_lighter_pink}
	// LighterCrimson is color lighter crimson
	LighterCrimson Color = Color{color: C.TCOD_lighter_crimson}
	// LightRed is color light red
	LightRed Color = Color{color: C.TCOD_light_red}
	// LightFlame is color light flame
	LightFlame Color = Color{color: C.TCOD_light_flame}
	// LightOrange is color light orange
	LightOrange Color = Color{color: C.TCOD_light_orange}
	// LightAmber is color light amber
	LightAmber Color = Color{color: C.TCOD_light_amber}
	// LightYellow is color light yellow
	LightYellow Color = Color{color: C.TCOD_light_yellow}
	// LightLime is color light lime
	LightLime Color = Color{color: C.TCOD_light_lime}
	// LightChartreuse is color light chartreuse
	LightChartreuse Color = Color{color: C.TCOD_light_chartreuse}
	// LightGreen is color light green
	LightGreen Color = Color{color: C.TCOD_light_green}
	// LightSea is color light sea
	LightSea Color = Color{color: C.TCOD_light_sea}
	// LightTurquoise is color light turquoise
	LightTurquoise Color = Color{color: C.TCOD_light_turquoise}
	// LightCyan is color light cyan
	LightCyan Color = Color{color: C.TCOD_light_cyan}
	// LightSky is color light sky
	LightSky Color = Color{color: C.TCOD_light_sky}
	// LightAzure is color light azure
	LightAzure Color = Color{color: C.TCOD_light_azure}
	// LightBlue is color light blue
	LightBlue Color = Color{color: C.TCOD_light_blue}
	// LightHan is color light han
	LightHan Color = Color{color: C.TCOD_light_han}
	// LightViolet is color light violet
	LightViolet Color = Color{color: C.TCOD_light_violet}
	// LightPurple is color light purple
	LightPurple Color = Color{color: C.TCOD_light_purple}
	// LightFuchsia is color light fuchsia
	LightFuchsia Color = Color{color: C.TCOD_light_fuchsia}
	// LightMagenta is color light magenta
	LightMagenta Color = Color{color: C.TCOD_light_magenta}
	// LightPink is color light pink
	LightPink Color = Color{color: C.TCOD_light_pink}
	// LightCrimson is color light crimson
	LightCrimson Color = Color{color: C.TCOD_light_crimson}
	// Red is color - red
	Red Color = Color{color: C.TCOD_red}
	// Flame is color - flame
	Flame Color = Color{color: C.TCOD_flame}
	// Orange is color - orange
	Orange Color = Color{color: C.TCOD_orange}
	// Amber is color - amber
	Amber Color = Color{color: C.TCOD_amber}
	// Yellow is color - yellow
	Yellow Color = Color{color: C.TCOD_yellow}
	// Lime is color - lime
	Lime Color = Color{color: C.TCOD_lime}
	// Chartreuse is color - chartreuse
	Chartreuse Color = Color{color: C.TCOD_chartreuse}
	// Green is color - green
	Green Color = Color{color: C.TCOD_green}
	// Sea is color - sea
	Sea Color = Color{color: C.TCOD_sea}
	// Turquoise is color - turquoise
	Turquoise Color = Color{color: C.TCOD_turquoise}
	// Cyan is color - cyan
	Cyan Color = Color{color: C.TCOD_cyan}
	// Sky is color - sky
	Sky Color = Color{color: C.TCOD_sky}
	// Azure is color - azure
	Azure Color = Color{color: C.TCOD_azure}
	// Blue is color - blue
	Blue Color = Color{color: C.TCOD_blue}
	// Han is color - han
	Han Color = Color{color: C.TCOD_han}
	// Violet is color - violet
	Violet Color = Color{color: C.TCOD_violet}
	// Purple is color - purple
	Purple Color = Color{color: C.TCOD_purple}
	// Fuchsia is color - fuchsia
	Fuchsia Color = Color{color: C.TCOD_fuchsia}
	// Magenta is color - magenta
	Magenta Color = Color{color: C.TCOD_magenta}
	// Pink is color - pink
	Pink Color = Color{color: C.TCOD_pink}
	// Crimson is color - crimson
	Crimson Color = Color{color: C.TCOD_crimson}
	// DarkRed is color dark red
	DarkRed Color = Color{color: C.TCOD_dark_red}
	// DarkFlame is color dark flame
	DarkFlame Color = Color{color: C.TCOD_dark_flame}
	// DarkOrange is color dark orange
	DarkOrange Color = Color{color: C.TCOD_dark_orange}
	// DarkAmber is color dark amber
	DarkAmber Color = Color{color: C.TCOD_dark_amber}
	// DarkYellow is color dark yellow
	DarkYellow Color = Color{color: C.TCOD_dark_yellow}
	// DarkLime is color dark lime
	DarkLime Color = Color{color: C.TCOD_dark_lime}
	// DarkChartreuse is color dark chartreuse
	DarkChartreuse Color = Color{color: C.TCOD_dark_chartreuse}
	// DarkGreen is color dark green
	DarkGreen Color = Color{color: C.TCOD_dark_green}
	// DarkSea is color dark sea
	DarkSea Color = Color{color: C.TCOD_dark_sea}
	// DarkTurquoise is color dark turquoise
	DarkTurquoise Color = Color{color: C.TCOD_dark_turquoise}
	// DarkCyan is color dark cyan
	DarkCyan Color = Color{color: C.TCOD_dark_cyan}
	// DarkSky is color dark sky
	DarkSky Color = Color{color: C.TCOD_dark_sky}
	// DarkAzure is color dark azure
	DarkAzure Color = Color{color: C.TCOD_dark_azure}
	// DarkBlue is color dark blue
	DarkBlue Color = Color{color: C.TCOD_dark_blue}
	// DarkHan is color dark han
	DarkHan Color = Color{color: C.TCOD_dark_han}
	// DarkViolet is color dark violet
	DarkViolet Color = Color{color: C.TCOD_dark_violet}
	// DarkPurple is color dark purple
	DarkPurple Color = Color{color: C.TCOD_dark_purple}
	// DarkFuchsia is color dark fuchsia
	DarkFuchsia Color = Color{color: C.TCOD_dark_fuchsia}
	// DarkMagenta is color dark magenta
	DarkMagenta Color = Color{color: C.TCOD_dark_magenta}
	// DarkPink is color dark pink
	DarkPink Color = Color{color: C.TCOD_dark_pink}
	// DarkCrimson is color dark crimson
	DarkCrimson Color = Color{color: C.TCOD_dark_crimson}
	// DarkerRed is color darker red
	DarkerRed Color = Color{color: C.TCOD_darker_red}
	// DarkerFlame is color darker flame
	DarkerFlame Color = Color{color: C.TCOD_darker_flame}
	// DarkerOrange is color darker orange
	DarkerOrange Color = Color{color: C.TCOD_darker_orange}
	// DarkerAmber is color darker amber
	DarkerAmber Color = Color{color: C.TCOD_darker_amber}
	// DarkerYellow is color darker yellow
	DarkerYellow Color = Color{color: C.TCOD_darker_yellow}
	// DarkerLime is color darker lime
	DarkerLime Color = Color{color: C.TCOD_darker_lime}
	// DarkerChartreuse is color darker chartreuse
	DarkerChartreuse Color = Color{color: C.TCOD_darker_chartreuse}
	// DarkerGreen is color darker green
	DarkerGreen Color = Color{color: C.TCOD_darker_green}
	// DarkerSea is color darker sea
	DarkerSea Color = Color{color: C.TCOD_darker_sea}
	// DarkerTurquoise is color darker turquoise
	DarkerTurquoise Color = Color{color: C.TCOD_darker_turquoise}
	// DarkerCyan is color darker cyan
	DarkerCyan Color = Color{color: C.TCOD_darker_cyan}
	// DarkerSky is color darker sky
	DarkerSky Color = Color{color: C.TCOD_darker_sky}
	// DarkerAzure is color darker azure
	DarkerAzure Color = Color{color: C.TCOD_darker_azure}
	// DarkerBlue is color darker blue
	DarkerBlue Color = Color{color: C.TCOD_darker_blue}
	// DarkerHan is color darker han
	DarkerHan Color = Color{color: C.TCOD_darker_han}
	// DarkerViolet is color darker violet
	DarkerViolet Color = Color{color: C.TCOD_darker_violet}
	// DarkerPurple is color darker purple
	DarkerPurple Color = Color{color: C.TCOD_darker_purple}
	// DarkerFuchsia is color darker fuchsia
	DarkerFuchsia Color = Color{color: C.TCOD_darker_fuchsia}
	// DarkerMagenta is color darker magenta
	DarkerMagenta Color = Color{color: C.TCOD_darker_magenta}
	// DarkerPink is color darker pink
	DarkerPink Color = Color{color: C.TCOD_darker_pink}
	// DarkerCrimson is color darker crimson
	DarkerCrimson Color = Color{color: C.TCOD_darker_crimson}
	// DarkestRed is color darkest red
	DarkestRed Color = Color{color: C.TCOD_darkest_red}
	// DarkestFlame is color darkest flame
	DarkestFlame Color = Color{color: C.TCOD_darkest_flame}
	// DarkestOrange is color darkest orange
	DarkestOrange Color = Color{color: C.TCOD_darkest_orange}
	// DarkestAmber is color darkest amber
	DarkestAmber Color = Color{color: C.TCOD_darkest_amber}
	// DarkestYellow is color darkest yellow
	DarkestYellow Color = Color{color: C.TCOD_darkest_yellow}
	// DarkestLime is color darkest lime
	DarkestLime Color = Color{color: C.TCOD_darkest_lime}
	// DarkestChartreuse is color darkest chartreuse
	DarkestChartreuse Color = Color{color: C.TCOD_darkest_chartreuse}
	// DarkestGreen is color darkest green
	DarkestGreen Color = Color{color: C.TCOD_darkest_green}
	// DarkestSea is color darkest sea
	DarkestSea Color = Color{color: C.TCOD_darkest_sea}
	// DarkestTurquoise is color darkest turquoise
	DarkestTurquoise Color = Color{color: C.TCOD_darkest_turquoise}
	// DarkestCyan is color darkest cyan
	DarkestCyan Color = Color{color: C.TCOD_darkest_cyan}
	// DarkestSky is color darkest sky
	DarkestSky Color = Color{color: C.TCOD_darkest_sky}
	// DarkestAzure is color darkest azure
	DarkestAzure Color = Color{color: C.TCOD_darkest_azure}
	// DarkestBlue is color darkest blue
	DarkestBlue Color = Color{color: C.TCOD_darkest_blue}
	// DarkestHan is color darkest han
	DarkestHan Color = Color{color: C.TCOD_darkest_han}
	// DarkestViolet is color darkest violet
	DarkestViolet Color = Color{color: C.TCOD_darkest_violet}
	// DarkestPurple is color darkest purple
	DarkestPurple Color = Color{color: C.TCOD_darkest_purple}
	// DarkestFuchsia is color darkest fuchsia
	DarkestFuchsia Color = Color{color: C.TCOD_darkest_fuchsia}
	// DarkestMagenta is color darkest magenta
	DarkestMagenta Color = Color{color: C.TCOD_darkest_magenta}
	// DarkestPink is color darkest pink
	DarkestPink Color = Color{color: C.TCOD_darkest_pink}
	// DarkestCrimson is color darkest crimson
	DarkestCrimson Color = Color{color: C.TCOD_darkest_crimson}

	// Brass is color brass
	Brass Color = Color{color: C.TCOD_brass}
	// Copper is color copper
	Copper Color = Color{color: C.TCOD_copper}
	// Gold is color gold
	Gold Color = Color{color: C.TCOD_gold}
	// Silver is color silver
	Silver Color = Color{color: C.TCOD_silver}
	// Celadon is color celadon
	Celadon Color = Color{color: C.TCOD_celadon}
	// Peach is color peach
	Peach Color = Color{color: C.TCOD_peach}
	// Grey is color grey
	Grey Color = Color{color: C.TCOD_grey}
	// Sepia is color sepia
	Sepia Color = Color{color: C.TCOD_sepia}
	// Black is color black
	Black Color = Color{color: C.TCOD_black}
	// White is color white
	White Color = Color{color: C.TCOD_white}
)
