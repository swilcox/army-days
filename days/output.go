package days

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gookit/color"
)

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func ColorOutputData(od []OutputData, w io.Writer) {
	outputLines := []string{}
	longestLine := 5
	for _, outputEntry := range od {
		dayNumString := fmt.Sprintf("%d", Abs(int(outputEntry.Days)))
		if float32(int(outputEntry.Days)) != outputEntry.Days {
			dayNumString = dayNumString + " and a butt"
		}
		dayString := ""
		switch {
		case outputEntry.Days == 1:
			dayString = fmt.Sprintf("%s day until", dayNumString)
		case outputEntry.Days == 0:
			dayString = "Today is"
		case outputEntry.Days > 1 || outputEntry.Days == 0.5:
			dayString = fmt.Sprintf("%s days until", dayNumString)
		case outputEntry.Days < -1 || outputEntry.Days == -0.5:
			dayString = fmt.Sprintf("%s days since", dayNumString)
		case outputEntry.Days == -1:
			dayString = fmt.Sprintf("%s day since", dayNumString)
		}
		line := fmt.Sprintf("%s %s.", dayString, outputEntry.Title)
		if len(line) > longestLine {
			longestLine = len(line)
		}
		outputLines = append(outputLines, line)
	}
	color1 := color.S256(255, 234)
	color2 := color.S256(255, 232)
	header := color.S256(227, 232).AddOpts(color.OpUnderscore) // underlined yellow
	output := fmt.Sprintf("\n%-*s\n", longestLine, " Days")
	fmt.Fprint(w, header.Sprintf(output))
	for i, outputLine := range outputLines {
		output = fmt.Sprintf("%-*s", longestLine, outputLine)
		if i%2 == 0 {
			fmt.Fprint(w, color1.Sprintf(output))
		} else {
			fmt.Fprint(w, color2.Sprintf(output))
		}
		fmt.Fprintln(w)
	}
}

func JsonOutputData(od []OutputData, w io.Writer) {
	b, err := json.Marshal(od)
	if err != nil {
		fmt.Fprintln(w, "Error converting output to JSON: ", err)
		return
	}
	fmt.Fprintln(w, string(b))
}

func DisplayOutputData(od []OutputData) {
	o, _ := os.Stdout.Stat()
	if (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		ColorOutputData(od, os.Stdout)
	} else {
		JsonOutputData(od, os.Stdout)
	}
}
