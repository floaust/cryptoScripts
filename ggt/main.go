package main

import (
	"github.com/pterm/pterm"
)

func main() {
	ggt(1012, 137)
}

func ggt(n int, a int) {
	var s1, s2, s3, t1, t2, t3 int
	var output1, output2 string
	for i := 0; i < n; i++ {
		if n%a == 0 {
			break
		}
		if i == 0 {
			t2 = n / a
			s2 = 1
			aTmp := a
			nTmp := n
			n = a
			a = nTmp % aTmp

			output1 += pterm.Sprintln(nTmp, "=", t2, "*", aTmp, "+", a)
			output2 += pterm.Sprintln(a, "=", s2, "* n -", t2, "* a;")
		} else if i == 1 {
			t1 = t2
			s1 = s2
			multiplicand := n / a
			s2 = multiplicand * s1
			t2 = multiplicand*t1 + 1
			aTmp := a
			nTmp := n
			n = a
			a = nTmp % aTmp

			output1 += pterm.Sprintln(nTmp, "=", multiplicand, "*", aTmp, "+", a)
			output2 += pterm.Sprintln(a, "= -", s2, "* n +", t2, "* a;")
		} else {
			multiplicand := n / a
			s3 = multiplicand*s2 + s1
			t3 = multiplicand*t2 + t1
			aTmp := a
			nTmp := n
			n = a
			a = nTmp % aTmp

			if i%2 == 1 {
				output1 += pterm.Sprintln(nTmp, "=", multiplicand, "*", aTmp, "+", a)
				output2 += pterm.Sprintln(a, "= -", s3, "* n +", t3, "* a;")
			} else {
				output1 += pterm.Sprintln(nTmp, "=", multiplicand, "*", aTmp, "+", a)
				output2 += pterm.Sprintln(a, "=", s3, "* n -", t3, "* a;")
			}

			t1 = t2
			t2 = t3
			s1 = s2
			s2 = s3
		}
	}

	panels := pterm.Panels{
		{{Data: output1}, {Data: output2}},
	}

	_ = pterm.DefaultPanel.WithPanels(panels).WithPadding(5).Render()
}
