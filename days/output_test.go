package days

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/gookit/color"
)

func TestJsonOutput(t *testing.T) {
	testOutputData := []OutputData{
		{Title: "item 1", Date: time.Date(2024, 1, 17, 0, 0, 0, 0, time.UTC), Days: 3},
	}
	expectedOutput := `[{"title":"item 1","date":"2024-01-17T00:00:00Z","days":3}]` + "\n"
	buf := new(bytes.Buffer)
	JsonOutputData(testOutputData, buf)
	if got := buf.String(); got != expectedOutput {
		t.Errorf("\nExpected: %v\nGot: %v\n", expectedOutput, got)
	}
}

func TestColorOutput(t *testing.T) {
	color.ForceOpenColor()
	testOutputData := []OutputData{
		{Title: "item -2", Date: time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC), Days: -5},
		{Title: "item -1", Date: time.Date(2024, 1, 13, 0, 0, 0, 0, time.UTC), Days: -1},
		{Title: "item 0", Date: time.Date(2024, 1, 14, 0, 0, 0, 0, time.UTC), Days: 0},
		{Title: "item 0.5", Date: time.Date(2024, 1, 14, 0, 0, 0, 0, time.UTC), Days: 0.5},
		{Title: "item 1", Date: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC), Days: 1},
		{Title: "item 2", Date: time.Date(2024, 1, 17, 0, 0, 0, 0, time.UTC), Days: 2.5},
		{Title: "item 3", Date: time.Date(2024, 3, 10, 0, 0, 0, 0, time.UTC), Days: 42},
	}
	expectedOutput, err := os.ReadFile("./test_data/output.txt")
	if err != nil {
		t.Errorf("error getting expected output.txt file.: %v", err)
	}
	buf := new(bytes.Buffer)
	ColorOutputData(testOutputData, buf)
	if got := buf.Bytes(); string(got) != string(expectedOutput) {
		_ = os.WriteFile("actual_output.txt", buf.Bytes(), 0644)
		t.Errorf("\nExpected: %v\nGot:\n%v\n", expectedOutput, got)
	}
}
