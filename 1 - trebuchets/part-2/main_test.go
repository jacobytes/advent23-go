package main

import (
	"fmt"
	"testing"
)

func BestCalibrationsPart2(t *testing.T) {
	calibrations := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	total := CalculateCalibrations(calibrations)

	if total != 281 {
		t.Fail()
	}
}

func TestCalibrationPart2EdgeCases(t *testing.T) {
	calibrations := []string{
		"eightone",
	}

	total := CalculateCalibrations(calibrations)

	if total != 18 {
		fmt.Println(total)
		t.Fail()
	}
}
