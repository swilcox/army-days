package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"
)

type DateEntry struct {
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

type DateEntries struct {
	Config  map[string]interface{} `json:"config"`
	Entries []DateEntry            `json:"entries"`
}

func ReadJsonFile(fileName string) DateEntries {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var entryData DateEntries
	json_err := json.Unmarshal(byteValue, &entryData)
	if json_err != nil {
		fmt.Println(json_err)
		os.Exit(1)
	}
	return entryData
}

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	entryData := ReadJsonFile(homeDir + string(os.PathSeparator) + ".armydays.json")
	fmt.Println("-------[ Recounting the Days ]-------")

	now := time.Now()

	useArmyButtDays := false
	if entryData.Config["useArmyButtDays"] == true {
		useArmyButtDays = true
	}
	for _, entry := range entryData.Entries {
		days := entry.Date.Sub(now).Hours() / 24
		tagline := "day"
		if useArmyButtDays && int(days) == int(math.Round(days)) {
			tagline = "and a butt " + tagline
		}
		intDays := int(math.Round(days))
		if !useArmyButtDays {
			intDays = int(days) + 1
		}
		if intDays > 1 {
			tagline += "s"
		}
		if days < 0.0 && days > -1.0 {
			fmt.Println("Today is " + entry.Title + ".")
		} else if intDays < 0 {
			if entryData.Config["showCompleted"] == true {
				fmt.Println("Completed: " + entry.Title + ".")
			}
		} else {
			fmt.Printf("%d %s until %s.\n", intDays, tagline, entry.Title)
		}
	}
	fmt.Println("-------------------------------------")
}
