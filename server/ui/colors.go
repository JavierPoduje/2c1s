package ui

import gloss "github.com/charmbracelet/lipgloss"

func AliveCellColor() gloss.Color {
	return gloss.Color("#40A02B")
}

func DeadCellColor() gloss.Color {
	return gloss.Color("#494D64")
}

func TogglerAliveCellColor() gloss.Color {
	return blueColor()
}

func TogglerDeadCellColor() gloss.Color {
	return redColor()
}

func foregroundColor() gloss.Color {
	return gloss.Color("#CAD3F5")
}
func secondaryForegroundColor() gloss.Color {
	return gloss.Color("#5c5f77")
}

func blackColor() gloss.Color {
	return gloss.Color("#494D64")
}

func redColor() gloss.Color {
	return gloss.Color("#ED8796")
}

func greenColor() gloss.Color {
	return gloss.Color("#A6DA95")
}

func yellowColor() gloss.Color {
	return gloss.Color("#EED49F")
}

func blueColor() gloss.Color {
	return gloss.Color("#8AADF4")
}

func magentaColor() gloss.Color {
	return gloss.Color("#F5BDE6")
}

func cyanColor() gloss.Color {
	return gloss.Color("#8BD5CA")
}

func whiteColor() gloss.Color {
	return gloss.Color("#B8C0E0")
}
