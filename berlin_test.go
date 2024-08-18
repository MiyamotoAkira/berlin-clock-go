package berlin_clock_test

import (
	berlin "berlin_clock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:00:00")

	expected := `O
OOOO
OOOO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestOneSecondPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:00:01")

	expected := `Y
OOOO
OOOO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestTwoSecondsPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:00:02")

	expected := `O
OOOO
OOOO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestThreeSecondsPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:00:03")

	expected := `Y
OOOO
OOOO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestOneMinutePastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:01:00")

	expected := `O
OOOO
OOOO
OOOOOOOOOOO
YOOO`

	assert.Equal(t, expected, actual)
}

func TestTwoMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:02:00")

	expected := `O
OOOO
OOOO
OOOOOOOOOOO
YYOO`

	assert.Equal(t, expected, actual)
}
