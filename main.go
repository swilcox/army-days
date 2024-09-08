package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/swilcox/army-days/days"
)

var Version = "0.4.2"

func main() {
	parser := argparse.NewParser("army-days", "Calculate days until events")
	filename := parser.String("f", "file", &argparse.Options{
		Required: false,
		Help:     "Input file (JSON or YAML)",
		Default:  "days.yaml",
	})
	versionFlag := parser.Flag("v", "version", &argparse.Options{
		Help: "version information",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	if *versionFlag {
		fmt.Println(Version)
		os.Exit(0)
	}

	configData := days.ProcessFile(*filename)
	outputData := days.ProcessEntries(configData)
	days.DisplayOutputData(outputData)
}
