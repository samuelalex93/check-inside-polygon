package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func isPointInPolygon(point Point, polygon []Point) bool {
	numVertices := len(polygon)
	if numVertices < 3 {
		return false
	}

	j := numVertices - 1
	oddNodes := false

	for i := 0; i < numVertices; i++ {
		if (polygon[i].Latitude < point.Latitude && polygon[j].Latitude >= point.Latitude ||
			polygon[j].Latitude < point.Latitude && polygon[i].Latitude >= point.Latitude) &&
			(polygon[i].Longitude <= point.Longitude || polygon[j].Longitude <= point.Longitude) {
			if polygon[i].Longitude+(point.Latitude-polygon[i].Latitude)/(polygon[j].Latitude-polygon[i].Latitude)*(polygon[j].Longitude-polygon[i].Longitude) < point.Longitude {
				oddNodes = !oddNodes
			}
		}
		j = i
	}

	return oddNodes
}

func main() {
	targetPolygon := "[[-20.45896903,-54.5857805],[-20.45937607,-54.58571345],[-20.45944391,-54.58624721],[-20.45908713,-54.58631426],[-20.45904441,-54.58632499]]"
	longitude := -54.5860
	latitude := -20.4591
	var valueChild [][]float64
	if err := json.Unmarshal([]byte(targetPolygon), &valueChild); err != nil {
		fmt.Println(err)
	}
	var pointStructs []Point
	for _, p := range valueChild {
		pointStructs = append(pointStructs, Point{
			Latitude:  p[0],
			Longitude: p[1],
		})
	}
	fmt.Println("Latitude ", latitude)
	fmt.Println("Longitude ", longitude)
	fmt.Println("Polygon ", pointStructs)
	pointToCheck := Point{Latitude: latitude, Longitude: longitude}
	if isPointInPolygon(pointToCheck, pointStructs) {
		fmt.Println("Point is inside polygon")
	} else {
		fmt.Println("Point is outside polygon")
	}
}
