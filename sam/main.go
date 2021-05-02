package main

import (
	"github.com/pterm/pterm"
	"strconv"
	"strings"
)

func main() {
	sam(45, 257, 133)
}

func sam(x int, n int, d int64) {
	pterm.Println(strconv.FormatInt(d, 2))
	binSplit := strings.Split(strconv.FormatInt(d, 2), "")
	var sOrm string
	var output string

	xMod := x

	for i, s := range binSplit {
		if i >= 1 {
			xMod = (xMod * xMod) % n
			sOrm += pterm.Sprintln(i, ": sq ")
			output += pterm.Sprintln(xMod, " mod ", n)
			if s == "1" {
				xMod = (xMod * x) % n
				sOrm += pterm.Sprintln(i, ": mul ")
				output += pterm.Sprintln(xMod, " mod ", n)
			}
		}
	}

	panels := pterm.Panels{
		{{Data: sOrm}, {Data: output}},
	}

	_ = pterm.DefaultPanel.WithPanels(panels).WithPadding(2).Render()
}
