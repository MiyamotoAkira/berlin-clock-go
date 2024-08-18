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
OOOOOOOOOOO
`)

	if minuteConverted > 0 {
		sb.WriteString("YOOO")
	} else {
		sb.WriteString("OOOO")
	}

	return sb.String()
}
