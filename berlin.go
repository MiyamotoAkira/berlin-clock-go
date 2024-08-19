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
		var topHourMarker string
		if i <= hoursRest {
			topHourMarker = "R"
		} else {
			topHourMarker = "O"
		}
		sb.WriteString(topHourMarker)
	}
	sb.WriteString("\n")

	hoursRemainder := hourConverted % 5
	for i := 1; i < 5; i++ {
		var bottomHourMarker string
		if i <= hoursRemainder {
			bottomHourMarker = "R"
		} else {
			bottomHourMarker = "O"
		}
		sb.WriteString(bottomHourMarker)
	}

	sb.WriteString("\n")

	minutesRest := minuteConverted / 5
	for i := 1; i < 12; i++ {
		var topMinuteMarker string
		if i <= minutesRest {
			if i%3 == 0 {
				topMinuteMarker = "R"
			} else {
				topMinuteMarker = "Y"
			}
		} else {
			topMinuteMarker = "O"
		}
		sb.WriteString(topMinuteMarker)
	}

	sb.WriteString("\n")

	minutesRemainder := minuteConverted % 5
	for i := 1; i < 5; i++ {
		var bottomMinuteMarker string
		if i <= minutesRemainder {
			bottomMinuteMarker = "Y"
		} else {
			bottomMinuteMarker = "O"
		}
		sb.WriteString(bottomMinuteMarker)
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
