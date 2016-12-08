package timezone

import (
	"errors"
	"fmt"
	"time"
)

// ValidLocation returns true if the given string can be found as a
// timezone location in the timezone.Locations array.
func ValidLocation(s string) bool {
	for _, zone := range Locations {
		if zone.Location == s {
			return true
		}
	}
	return false
}

func GroupLocationByOffset() map[int][]Timezone {
	var z = make(map[int][]Timezone)
	for _, loc := range Locations {
		l, err := time.LoadLocation(loc.Location)
		if err != nil {
			fmt.Println("bad location", loc.Location)
			continue
		}
		_, offset := time.Now().In(l).Zone()
		z[offset] = append(z[offset], loc)
	}
	return z

}

func LocationsFromOffset(offset int) ([]Timezone, error) {
	var z []Timezone
	for _, loc := range Locations {
		l, err := time.LoadLocation(loc.Location)
		if err != nil {
			continue
		}
		_, zoneOffset := time.Now().In(l).Zone()
		if offset == zoneOffset {
			z = append(z, loc)
		}
	}
	return z, nil
}

// Offset returns the abbreviated name of the zone of l (such as "CET")
// and its offset in seconds east of UTC. The location should be valid IANA
// timezone location.
func Offset(loc string) (zone string, offset int, err error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return zone, offset, err
	}
	zone, offset = time.Now().In(l).Zone()
	return zone, offset, nil
}

// Country returns all timezones with given country name.
// If none is found, returns an error.
func Country(c string) ([]Timezone, error) {
	var z []Timezone
	for _, zone := range Locations {
		if zone.Country == c {
			z = append(z, zone)
		}
	}
	if len(z) == 0 {
		return z, errors.New("no timezones found")
	}
	return z, nil
}

// Code returns all timezones with given country code.
// If none is found, returns an error.
func Code(c string) ([]Timezone, error) {
	var z []Timezone
	for _, zone := range Locations {
		if zone.Code == c {
			z = append(z, zone)
		}
	}
	if len(z) == 0 {
		return z, errors.New("no timezones found")
	}
	return z, nil
}
