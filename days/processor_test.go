package days

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestCalculateDays(t *testing.T) {
	now := time.Date(2023, 5, 1, 10, 0, 0, 0, time.UTC)
	tests := []struct {
		name            string
		entryDate       string
		useArmyButtDays bool
		expected        float32
	}{
		{"Future date", "2023-05-10", false, 9},
		{"Past date", "2023-04-20", false, -11},
		{"Same day", "2023-05-01", false, 0},
		{"Army butt day before noon", "2023-05-10", true, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entryDate, _ := time.Parse(dateFormat, tt.entryDate)
			result := CalculateDays(entryDate, now, tt.useArmyButtDays)
			if result != tt.expected {
				t.Errorf("calculateDays() = %v, want %v", result, tt.expected)
			}
		})
	}

	// Test Army butt day after noon
	nowAfterNoon := time.Date(2023, 5, 1, 14, 0, 0, 0, time.UTC)
	entryDate, _ := time.Parse(dateFormat, "2023-05-10")
	result := CalculateDays(entryDate, nowAfterNoon, true)
	expected := float32(8.5)
	if result != expected {
		t.Errorf("calculateDays() with Army butt day after noon = %v, want %v", result, expected)
	}
}

func TestDateParser(t *testing.T) {
	tests := []struct {
		name         string
		dateString   string
		expectedDate time.Time
		expectedErr  error
	}{
		{"short date", "2023-01-17", time.Date(2023, 1, 17, 0, 0, 0, 0, time.UTC), nil},
		{
			"long date",
			"2023-01-17T10:12:00+00:00",
			time.Date(
				2023, 1, 17, 10, 12, 0, 0, time.FixedZone("", 0),
			),
			nil,
		},
		{"z date", "2023-01-17T10:12:00Z", time.Date(2023, 1, 17, 10, 12, 0, 0, time.UTC), nil},
		{"bad date 1", "2023-01-17T", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), errors.New("error")},
		{"bad date 2", "2023/01/17", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), errors.New("error")},
		{"bad date 3", "2023-01-17 12:31:00+00:00", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), errors.New("error")}, // note: only RFC3339 date/times in full form
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseDateString(tt.dateString)
			if result != tt.expectedDate || (tt.expectedErr == nil && err != nil) || (tt.expectedErr != nil && err == nil) {
				t.Errorf("calculateDays() = %v (%v), want %v, %v", result, err, tt.expectedDate, tt.expectedErr)
			}
		})
	}

}

func TestProcessFile(t *testing.T) {
	// now := time.Date(2023, 5, 1, 10, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		fileName string
	}{
		{"days1", "./test_data/days1.yaml"},
		{"days2", "./test_data/days2.yaml"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configData := ProcessFile(tt.fileName)
			// add more to the test here...
			fmt.Println(configData)
		})
	}
}

func TestProcessEntries(t *testing.T) {
	configData := ConfigData{
		Config: Config{false, false},
		Entries: []Entry{
			{"item 1", "2100-01-17"},
			{"item 2", "2099-01-17"},
			{"item 3", "2000-01-01"},
		},
	}
	testOutputData := ProcessEntries(configData)
	if len(testOutputData) != 2 {
		t.Error("Expected 2 entries")
	}
	if testOutputData[0].Title != "item 2" {
		t.Error("Sorting of entries has failed")
	}
}
