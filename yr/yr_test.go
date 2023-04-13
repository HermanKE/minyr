package yr

import (
	"testing"
	"strings"
)

func TestCelsiusToFahrenheitString(t *testing.T) {
	type test struct {
		input string
		want string
	}

	tests := []test {
		{input: "6", want: "42.8"},
		{input: "0", want: "32.0"},
	}

	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitString(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}

}

func TestCelsiusToFahrenheitLine(t *testing.T) {
	type test struct {
		input string
		want string
	}

	tests := []test {
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
	}

	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitLine(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}

}

func TestLineAmountOfCelsiusFile(t *testing.T) {
	expected := 16756
	filename := "../kjevik-temp-celsius-20220318-20230318.csv"
	lineCount, err := CountLines(filename)
	if err != nil {
		t.Errorf("error counting lines in %s: %v", filename, err)
	}
	if lineCount != expected {
		t.Errorf("unexpected number of lines in %s, expected %d but got %d", filename, expected, lineCount)
	}
}

func TestCalculateAverageFourthElementInCelsiusFile(t *testing.T) {
	filePath := "../kjevik-temp-celsius-20220318-20230318.csv"
	expectedAverage := 8.56

	average, err := CalculateAverageFourthElement(filePath)
	if err != nil {
		t.Errorf("error calculating average: %v", err)
	}

	if average != expectedAverage {
		t.Errorf("average %v does not match expected value %v", average, expectedAverage)
	}
}

func TestReadLastLineContains(t *testing.T) {
    filePath := "../table.csv"
    expectedString := "Data er gyldig per 20.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);endringen er gjort av Herman Erlingsen"

    lastLine, err := ReadLastLine(filePath)
    if err != nil {
        t.Errorf("ReadLastLine returned an error: %v", err)
    }

    if !strings.Contains(lastLine, expectedString) {
        t.Errorf("last line doesn't contain the expected string")
    }
}