package berlin_clock

import (
	"fmt"
	"strconv"
	"strings"
)

func GetClockString(time string) string {
	var sb strings.Builder
	secondConverted, minuteConverted, hourConverted := extractTimeComponents(time)

	var secondMarker string
	if secondConverted%2 == 0 {
		secondMarker = "Y"
	} else {
		secondMarker = "O"
	}
	sb.WriteString(secondMarker)
	sb.WriteString("\n")

	hoursRest := hourConverted / 5
	for i := 1; i < 5; i++ {
		var topHourSingleMarker string
		if i <= hoursRest {
			topHourSingleMarker = "R"
		} else {
			topHourSingleMarker = "O"
		}
		sb.WriteString(topHourSingleMarker)
	}
	sb.WriteString("\n")

	hoursRemainder := hourConverted % 5
	for i := 1; i < 5; i++ {
		var bottomHourSingleMarker string
		if i <= hoursRemainder {
			bottomHourSingleMarker = "R"
		} else {
			bottomHourSingleMarker = "O"
		}
		sb.WriteString(bottomHourSingleMarker)
	}

	sb.WriteString("\n")

	minutesRest := minuteConverted / 5
	for i := 1; i < 12; i++ {
		var topMinuteSingleMarker string
		if i <= minutesRest {
			if i%3 == 0 {
				topMinuteSingleMarker = "R"
			} else {
				topMinuteSingleMarker = "Y"
			}
		} else {
			topMinuteSingleMarker = "O"
		}
		sb.WriteString(topMinuteSingleMarker)
	}

	sb.WriteString("\n")

	minutesRemainder := minuteConverted % 5
	for i := 1; i < 5; i++ {
		var bottomMinuteSingleMarker string
		if i <= minutesRemainder {
			bottomMinuteSingleMarker = "Y"
		} else {
			bottomMinuteSingleMarker = "O"
		}
		sb.WriteString(bottomMinuteSingleMarker)
	}

	return sb.String()
}

func extractTimeComponents(time string) (int, int, int) {
	second := time[7:8]
	secondConverted, err := strconv.Atoi(second)
	if err != nil {
		fmt.Print("error")
	}
	minute := time[3:5]
	minuteConverted, err := strconv.Atoi(minute)
	hour := time[0:2]
	hourConverted, err := strconv.Atoi(hour)
	return secondConverted, minuteConverted, hourConverted
}
