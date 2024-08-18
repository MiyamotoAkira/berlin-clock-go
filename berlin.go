package berlin_clock

import (
	"fmt"
	"strconv"
	"strings"
)

func GetClockString(time string) string {
	var sb strings.Builder
	second := time[len(time)-1 : len(time)]
	secondConverted, err := strconv.Atoi(second)
	if err != nil {
		fmt.Print("error")
	}

	if secondConverted%2 == 0 {
		sb.WriteString("O")
	} else {
		sb.WriteString("Y")
	}

	sb.WriteString(`
OOOO
OOOO
OOOOOOOOOOO
OOOO`)

	return sb.String()
}
