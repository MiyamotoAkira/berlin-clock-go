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

func TestThreeMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:03:00")

	expected := `O
OOOO
OOOO
OOOOOOOOOOO
YYYO`

	assert.Equal(t, expected, actual)
}

func TestFourMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:04:00")

	expected := `O
OOOO
OOOO
OOOOOOOOOOO
YYYY`

	assert.Equal(t, expected, actual)
}

func TestFiveMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:05:00")

	expected := `O
OOOO
OOOO
YOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestSixMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:06:00")

	expected := `O
OOOO
OOOO
YOOOOOOOOOO
YOOO`

	assert.Equal(t, expected, actual)
}

func TestSevenMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:07:00")

	expected := `O
OOOO
OOOO
YOOOOOOOOOO
YYOO`

	assert.Equal(t, expected, actual)
}

func TestTenMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:10:00")

	expected := `O
OOOO
OOOO
YYOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}
