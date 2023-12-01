package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	calibrations := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	total := 0

	for _, calibration := range calibrations {

		firstCalibration := getFirstNumerical(calibration)

		reversed := reverseString(calibration)

		secondCalibration := getFirstNumerical(reversed)

		result, _ := strconv.Atoi(firstCalibration + secondCalibration)

		total += result
	}

	fmt.Println(total)
}

func getFirstNumerical(calibration string) string {

	pattern := `\d`

	re := regexp.MustCompile(pattern)

	match := re.FindString(calibration)

	return match
}

func reverseString(value string) string {
	runes := []rune(value)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
