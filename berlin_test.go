package berlin_clock_test

import (
	berlin "berlin_clock"
	"testing"
)

func TestMidnight(t *testing.T) {

	actual := berlin.GetClockString("00:00:00")

	expected := `O
OOOO
OOOO
OOOOOOOOOOO
OOOO`

	if actual != expected {
		t.Errorf("Wrong clock. Expected\n%s\nbut got\n%s\ninstead", expected, actual)
	}

}
