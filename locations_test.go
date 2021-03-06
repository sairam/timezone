package timezone

import "testing"

var expectedLocations = []Timezone{
	Timezone{"Africa/Abidjan", "CI", "Ivory Coast"},
}

func TestLocations(t *testing.T) {
	if Locations[0] != expectedLocations[0] {
		t.Errorf("ZonesTest returned %s", Locations[0])
	}

	if len(Locations) != 584 {
		t.Errorf("Length of zones was not 584, actual %d", len(Locations))
	}
}
