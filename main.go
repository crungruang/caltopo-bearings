package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
)

type Feature struct {
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type Geometry struct {
	Coordinates []interface{} `json:"coordinates"`
	Type        string        `json:"type"`
}

type Properties struct {
	Title string `json:"title"`
	Class string `json:"class"`
}

type FeatureCollection struct {
	Features []Feature `json:"features"`
}

func calculateBearingAndDistance(features []Feature) {
	for i := 0; i < len(features); i++ {
		for j := 0; j < len(features); j++ {
			if i == j {
				continue
			}

			fromFeature := features[i]
			toFeature := features[j]

			if fromFeature.Geometry.Type == "Point" && fromFeature.Properties.Class == "Marker" &&
				toFeature.Geometry.Type == "Point" && toFeature.Properties.Class == "Marker" {

				fromCoordinates := fromFeature.Geometry.Coordinates
				toCoordinates := toFeature.Geometry.Coordinates

				if len(fromCoordinates) >= 2 && len(toCoordinates) >= 2 {
					lat1 := fromCoordinates[1].(float64)
					lon1 := fromCoordinates[0].(float64)
					lat2 := toCoordinates[1].(float64)
					lon2 := toCoordinates[0].(float64)

					point1 := orb.Point{lon1, lat1}
					point2 := orb.Point{lon2, lat2}

					bearing := geo.Bearing(point1, point2)
					distance := geo.Distance(point1, point2)

					fmt.Printf("From %s to %s:\n", fromFeature.Properties.Title, toFeature.Properties.Title)
					fmt.Printf("Bearing: %.2f degrees\n", bearing)
					fmt.Printf("Distance: %.2f meters\n", distance)
					fmt.Println()
				}
			}
		}
	}
}

func main() {
	// Prompt the user for the JSON file name
	fmt.Print("Enter the name of the JSON file: ")
	var fileName string
	_, err := fmt.Scan(&fileName)
	if err != nil {
		fmt.Printf("Error reading the input: %v\n", err)
		return
	}

	// Read JSON data from the user-specified file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading the JSON file: %v\n", err)
		return
	}

	var featureCollection FeatureCollection

	err = json.Unmarshal(data, &featureCollection)
	if err != nil {
		fmt.Printf("Error parsing the JSON: %v\n", err)
		return
	}

	// Find sequences of "Point" features with "Marker" class and calculate bearing and distance
	var sequence []Feature
	for _, feature := range featureCollection.Features {
		if feature.Geometry.Type == "Point" && feature.Properties.Class == "Marker" {
			sequence = append(sequence, feature)
		}
	}

	// Calculate bearing and distance for all possible permutations of coordinates
	calculateBearingAndDistance(sequence)
}
