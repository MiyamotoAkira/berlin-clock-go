package berlin_clock

import (
	"fmt"
	"strconv"
	"strings"
)

func GetClockString(time string) string {
	var sb strings.Builder
	second := time[7:8]
	secondConverted, err := strconv.Atoi(second)
	if err != nil {
		fmt.Print("error")
	}
	minute := time[3:5]
	minuteConverted, err := strconv.Atoi(minute)

	if secondConverted%2 == 0 {
		sb.WriteString("O")
	} else {
		sb.WriteString("Y")
	}

	sb.WriteString(`
OOOO
OOOO
`)

	minutesRest := minuteConverted / 5
	for i := 1; i < 12; i++ {
		if i <= minutesRest {
			if i%3 == 0 {
				sb.WriteString("R")
			} else {
				sb.WriteString("Y")
			}
		} else {
			sb.WriteString("O")
		}
	}

	sb.WriteString("\n")

	minutesRemainder := minuteConverted % 5
	for i := 1; i < 5; i++ {
		if i <= minutesRemainder {
			sb.WriteString("Y")
		} else {
			sb.WriteString("O")
		}
	}

	return sb.String()
}
