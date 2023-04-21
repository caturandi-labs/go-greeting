package main

import (
	"strings"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {

	headText := strings.Repeat("###", 23)

	header := pterm.DefaultHeader.WithFullWidth(true).WithBackgroundStyle(pterm.NewStyle(pterm.BgGreen)).WithTextStyle(pterm.NewStyle(pterm.FgBlack))
	pterm.DefaultCenter.Println(header.Sprint(headText))

	area, _ := pterm.DefaultArea.Start()
	str, _ := pterm.DefaultBigText.
		WithLetters(
			putils.LettersFromStringWithStyle("SELAMAT ", pterm.NewStyle(pterm.FgGreen)),
			putils.LettersFromStringWithStyle("HARI RAYA", pterm.NewStyle(pterm.FgGreen))).Srender()

	str = pterm.DefaultCenter.Sprint(str)
	area.Update(str)
	area.Stop()

	area, _ = pterm.DefaultArea.Start()
	str, _ = pterm.DefaultBigText.
		WithLetters(
			putils.LettersFromStringWithStyle("IDUL FITRI", pterm.NewStyle(pterm.FgLightYellow))).Srender()

	str = pterm.DefaultCenter.Sprint(str)
	area.Update(str)
	area.Stop()

	area, _ = pterm.DefaultArea.Start()
	str, _ = pterm.DefaultBigText.
		WithLetters(
			putils.LettersFromStringWithStyle("1 SYAWAL 1444H", pterm.NewStyle(pterm.FgLightGreen))).Srender()

	str = pterm.DefaultCenter.Sprint(str)
	area.Update(str)
	area.Stop()

	footer := pterm.DefaultHeader.WithFullWidth(true).WithBackgroundStyle(pterm.NewStyle(pterm.BgGreen)).WithTextStyle(pterm.NewStyle(pterm.FgBlack))
	pterm.DefaultCenter.Println(footer.Sprint("MOHON MAAF LAHIR & BATHIN - CATUR ANDI PAMUNGKAS"))

}
