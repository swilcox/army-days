package days

import (
	"fmt"
	"os"
	"sort"
	"time"

	"gopkg.in/yaml.v3"
)

const dateFormat = "2006-01-02"

func ProcessFile(filename string) ConfigData {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	var d ConfigData
	err = yaml.Unmarshal(data, &d)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}
	return d
}

func parseDateString(ds string) (time.Time, error) {
	dateTime, err := time.Parse(dateFormat, ds)
	if err != nil {
		dateTime, err = time.Parse(time.RFC3339, ds)
	}
	return dateTime, err
}

func ProcessEntries(d ConfigData) []OutputData {
	sort.Slice(d.Entries, func(i, j int) bool {
		dateI, _ := parseDateString(d.Entries[i].Date)
		dateJ, _ := parseDateString(d.Entries[j].Date)
		return dateI.Before(dateJ)
	})
	now := time.Now()
	results := []OutputData{}
	for _, entry := range d.Entries {
		entryDate, err := parseDateString(entry.Date)
		if err != nil {
			fmt.Printf("Error parsing date for entry '%s': %v\n", entry.Title, err)
			continue
		}
		days := CalculateDays(entryDate, now, d.Config.UseArmyButtDays)
		if days >= 0 || d.Config.ShowCompleted {
			results = append(results, OutputData{entry.Title, entryDate, days})
		}
	}
	return results
}

func CalculateDays(entryDate, now time.Time, useArmyButtDays bool) float32 {
	simplifiedEntry := time.Date(entryDate.Year(), entryDate.Month(), entryDate.Day(), 0, 0, 0, 0, time.UTC)
	simplfiedNow := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	daysInt := int(simplifiedEntry.Sub(simplfiedNow).Hours() / 24)
	days := float32(daysInt)
	if useArmyButtDays && days > 0 && now.Hour() >= 12 {
		days -= .5
	}
	return days
}
