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

	topHourMarker := getTopHourMarker(hourConverted)

	sb.WriteString(topHourMarker)
	sb.WriteString("\n")

	bottomHourMarker := getBottomHourMarker(hourConverted)

	sb.WriteString(bottomHourMarker)
	sb.WriteString("\n")

	topMinuteMarker := getTopMinuteMarker(minuteConverted)

	sb.WriteString(topMinuteMarker)
	sb.WriteString("\n")

	bottomMinuteMarker := getBottomMinuteMarker(minuteConverted)

	sb.WriteString(bottomMinuteMarker)

	return sb.String()
}

func getBottomMinuteMarker(minuteConverted int) string {
	minutesRemainder := minuteConverted % 5
	var bottomMinuteMarkerBuilder strings.Builder
	for i := 1; i < 5; i++ {
		var bottomMinuteSingleMarker string
		if i <= minutesRemainder {
			bottomMinuteSingleMarker = "Y"
		} else {
			bottomMinuteSingleMarker = "O"
		}
		bottomMinuteMarkerBuilder.WriteString(bottomMinuteSingleMarker)
	}
	return bottomMinuteMarkerBuilder.String()

}

func getTopMinuteMarker(minuteConverted int) string {
	minutesRest := minuteConverted / 5
	var topMinuteMarkerBuilder strings.Builder
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
		topMinuteMarkerBuilder.WriteString(topMinuteSingleMarker)
	}
	return topMinuteMarkerBuilder.String()

}

func getBottomHourMarker(hourConverted int) string {
	hoursRemainder := hourConverted % 5
	var bottomHourMarkerBuilder strings.Builder
	for i := 1; i < 5; i++ {
		var bottomHourSingleMarker string
		if i <= hoursRemainder {
			bottomHourSingleMarker = "R"
		} else {
			bottomHourSingleMarker = "O"
		}
		bottomHourMarkerBuilder.WriteString(bottomHourSingleMarker)
	}
	return bottomHourMarkerBuilder.String()
}

func getTopHourMarker(hourConverted int) string {
	hoursRest := hourConverted / 5
	var topHourMarkerBuilder strings.Builder
	for i := 1; i < 5; i++ {
		var topHourSingleMarker string
		if i <= hoursRest {
			topHourSingleMarker = "R"
		} else {
			topHourSingleMarker = "O"
		}
		topHourMarkerBuilder.WriteString(topHourSingleMarker)
	}
	return topHourMarkerBuilder.String()
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
