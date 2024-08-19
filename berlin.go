package berlin_clock

import (
	"fmt"
	"strconv"
	"strings"
)

const RED_MARKER = "R"
const YELLOW_MARKER = "Y"
const OFF_MARKER = "O"

func GetClockString(time string) string {
	var sb strings.Builder
	secondConverted, minuteConverted, hourConverted := extractTimeComponents(time)

	sb.WriteString(getSecondMarker(secondConverted))
	sb.WriteString("\n")

	sb.WriteString(getTopHourMarker(hourConverted))
	sb.WriteString("\n")

	sb.WriteString(getBottomHourMarker(hourConverted))
	sb.WriteString("\n")

	sb.WriteString(getTopMinuteMarker(minuteConverted))
	sb.WriteString("\n")

	sb.WriteString(getBottomMinuteMarker(minuteConverted))

	return sb.String()
}

func getSecondMarker(secondConverted int) string {
	if secondConverted%2 == 0 {
		return YELLOW_MARKER
	} else {
		return OFF_MARKER
	}
}

func getBottomMinuteMarker(minuteConverted int) string {
	minutesRemainder := minuteConverted % 5
	var bottomMinuteMarkerBuilder strings.Builder
	for i := 1; i < 5; i++ {
		var bottomMinuteSingleMarker string
		if i <= minutesRemainder {
			bottomMinuteSingleMarker = YELLOW_MARKER
		} else {
			bottomMinuteSingleMarker = OFF_MARKER
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
				topMinuteSingleMarker = RED_MARKER
			} else {
				topMinuteSingleMarker = YELLOW_MARKER
			}
		} else {
			topMinuteSingleMarker = OFF_MARKER
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
			bottomHourSingleMarker = RED_MARKER
		} else {
			bottomHourSingleMarker = OFF_MARKER
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
			topHourSingleMarker = RED_MARKER
		} else {
			topHourSingleMarker = OFF_MARKER
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
