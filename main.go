package main

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/winge005/scmd/commands"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		pterm.DefaultBigText.WithLetters(
			pterm.NewLettersFromStringWithStyle("S", pterm.NewStyle(pterm.FgCyan)),
			pterm.NewLettersFromStringWithStyle("cmd", pterm.NewStyle(pterm.FgLightMagenta))).
			Render()
		pterm.DefaultBasicText.Println(pterm.Red("Simon's") + pterm.LightWhite("Command Line") + pterm.Blue("tool!"))
		pterm.DefaultBasicText.Println(pterm.LightGreen("Run scmd with -help for help"))
		pterm.Error.Println("No arguments provided!!")
	}
	pterm.Info.Println("Argument provided " + pterm.LightWhite(argsWithoutProg[0]))

	switch argsWithoutProg[0] {
	case "help":
		fmt.Println(commands.HelpGetInfo())
		break
	}
}
