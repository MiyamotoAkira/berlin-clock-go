package berlin_clock

import (
	"fmt"
	"strconv"
)

func GetClockString(time string) string {
	second := time[len(time)-1 : len(time)]

	converted, err := strconv.Atoi(second)

	if err != nil {
		fmt.Print("error")
	}

	is_even := converted%2 == 0

	if is_even {
		return `O
OOOO
OOOO
OOOOOOOOOOO
OOOO`
	}
	return `Y
OOOO
OOOO
OOOOOOOOOOO
OOOO`
}
