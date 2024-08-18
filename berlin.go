package berlin_clock

import (
	"fmt"
	"strconv"
	"strings"
)

func GetClockString(time string) string {
	second := time[len(time)-1 : len(time)]
	var sb strings.Builder

	converted, err := strconv.Atoi(second)

	if err != nil {
		fmt.Print("error")
	}

	isEven := converted%2 == 0

	if isEven {
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
