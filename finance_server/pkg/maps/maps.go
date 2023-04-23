package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"runtime"

	"googlemaps.github.io/maps"
)

func GoogleMaps() {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyCRqlyMOvXhAMHkRQeYI6LK2l6_UKUR21s"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := &maps.GeocodingRequest{
		Address: "1600 Amphitheatre Parkway, Mountain View, CA",
	}

	resp, err := c.Geocode(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	location := resp[0].Geometry.Location
	lat := location.Lat
	lng := location.Lng

	// Construct a URL for the Google Maps window
	q := url.Values{}
	q.Set("api", "1")
	q.Set("query", fmt.Sprintf("%f,%f", lat, lng))
	url := fmt.Sprintf("https://www.google.com/maps/search/?%s", q.Encode())

	// Open the URL in a browser
	err = openURL(url)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

func openURL(url string) error {
	var cmd string
	switch {
	case isWindows():
		cmd = "cmd /c start"
	case isMac():
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	return exec.Command(cmd, url).Run()
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func isMac() bool {
	return runtime.GOOS == "darwin"
}

func main() {
	GoogleMaps()
}
