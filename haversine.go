package main

import (
	"fmt"
	"math"
	"sort"
)

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func haversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371
	dLat := toRadians(lat2 - lat1)
	dLon := toRadians(lon2 - lon1)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(toRadians(lat1))*math.Cos(toRadians(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c

	return distance
}

type Restaurant struct {
	Name      string
	Latitude  float64
	Longitude float64
	Distance  float64
}

func findClosestRestaurants(clientLat, clientLon float64, restaurants []Restaurant) []Restaurant {
	for i := range restaurants {
		restaurants[i].Distance = haversineDistance(clientLat, clientLon, restaurants[i].Latitude, restaurants[i].Longitude)
	}

	sort.Slice(restaurants, func(i, j int) bool {
		return restaurants[i].Distance < restaurants[j].Distance
	})

	return restaurants
}

func main() {
	clientLat := -23.550520
	clientLon := -46.633308

	restaurants := []Restaurant{
		{Name: "Restaurante do João", Latitude: -23.5489, Longitude: -46.6388},
		{Name: "Lanchonete da Maria", Latitude: -23.5510, Longitude: -46.6340},
		{Name: "Cafeteria Chique", Latitude: -23.5530, Longitude: -46.6400},
	}

	closestRestaurants := findClosestRestaurants(clientLat, clientLon, restaurants)

	for _, restaurant := range closestRestaurants {
		fmt.Printf("Nome: %s, Distância %.2f km\n", restaurant.Name, restaurant.Distance)
	}
}
