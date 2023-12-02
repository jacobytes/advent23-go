package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var wordToNumberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	calibrations := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	total := CalculateCalibrations(calibrations)

	fmt.Println(total)
}

func CalculateCalibrations(calibrationData []string) int {
	total := 0

	for _, calibration := range calibrationData {

		firstCalibration := getFirstNumerical(calibration)

		secondCalibration := getFinalNumerical(calibration)

		result, _ := strconv.Atoi(firstCalibration + secondCalibration)

		total += result
	}

	return total
}

func getFirstNumerical(calibration string) string {

	validCalibrationRegex := `one|two|three|four|five|six|seven|eight|nine|\d`

	re := regexp.MustCompile(validCalibrationRegex)

	match := re.FindString(calibration)

	if isWordWithNumericalValue(match) {
		return getDigitForWord(match)
	}

	return match
}

func getFinalNumerical(calibration string) string {
	validFinalCalibrationRegex := `one|two|three|four|five|six|seven|eight|nine|\d`

	re := regexp.MustCompile(validFinalCalibrationRegex)

	matches := re.FindAllStringSubmatchIndex(calibration, -1)

	finalMatchIndex := matches[len(matches)-1]
	result := calibration[finalMatchIndex[0]:finalMatchIndex[1]]

	if isWordWithNumericalValue(result) {
		return getDigitForWord(result)
	}

	return result
}

func isWordWithNumericalValue(value string) bool {

	_, err := strconv.Atoi(value)

	if err != nil {
		return true
	}

	return false
}

func getDigitForWord(word string) string {
	value, _ := wordToNumberMap[word]
	return value
}
