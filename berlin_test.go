package berlin_clock_test

import (
	berlin "berlin_clock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:00:00")

	expected := `Y
OOOO
OOOO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestOneSecondPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:00:01")

	expected := `O
OOOO
OOOO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestTwoSecondsPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:00:02")

	expected := `Y
OOOO
OOOO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestThreeSecondsPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:00:03")

	expected := `O
OOOO
OOOO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestOneMinutePastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:01:00")

	expected := `Y
OOOO
OOOO
OOOOOOOOOOO
YOOO`

	assert.Equal(t, expected, actual)
}

func TestTwoMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:02:00")

	expected := `Y
OOOO
OOOO
OOOOOOOOOOO
YYOO`

	assert.Equal(t, expected, actual)
}

func TestThreeMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:03:00")

	expected := `Y
OOOO
OOOO
OOOOOOOOOOO
YYYO`

	assert.Equal(t, expected, actual)
}

func TestFourMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:04:00")

	expected := `Y
OOOO
OOOO
OOOOOOOOOOO
YYYY`

	assert.Equal(t, expected, actual)
}

func TestFiveMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:05:00")

	expected := `Y
OOOO
OOOO
YOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestSixMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:06:00")

	expected := `Y
OOOO
OOOO
YOOOOOOOOOO
YOOO`

	assert.Equal(t, expected, actual)
}

func TestSevenMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:07:00")

	expected := `Y
OOOO
OOOO
YOOOOOOOOOO
YYOO`

	assert.Equal(t, expected, actual)
}

func TestTenMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:10:00")

	expected := `Y
OOOO
OOOO
YYOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestFifteenMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:15:00")

	expected := `Y
OOOO
OOOO
YYROOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestThirtyMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:30:00")

	expected := `Y
OOOO
OOOO
YYRYYROOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestFiftyFiveMinutesPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("00:55:00")

	expected := `Y
OOOO
OOOO
YYRYYRYYRYY
OOOO`

	assert.Equal(t, expected, actual)
}

func TestOneHourPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("01:00:00")

	expected := `Y
OOOO
ROOO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestTwoHoursPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("02:00:00")

	expected := `Y
OOOO
RROO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestThreeHoursPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("03:00:00")

	expected := `Y
OOOO
RRRO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestFourHoursPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("04:00:00")

	expected := `Y
OOOO
RRRR
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestFiveHoursPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("05:00:00")

	expected := `Y
ROOO
OOOO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestTenHoursPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("10:00:00")

	expected := `Y
RROO
OOOO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestTwentyThreeHoursPastMidnight(t *testing.T) {
	actual := berlin.GetClockString("23:00:00")

	expected := `Y
RRRR
RRRO
OOOOOOOOOOO
OOOO`

	assert.Equal(t, expected, actual)
}

func TestExampleFromReadme(t *testing.T) {
	actual := berlin.GetClockString("12:56:01")

	expected := `O
RROO
RROO
YYRYYRYYRYY
YOOO`

	assert.Equal(t, expected, actual)
}
