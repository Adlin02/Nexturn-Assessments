// Exercise: Climate Insights Tracker

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Location struct to represent climate data
type Location struct {
	Name        string
	Temperature float64
	Precipitation float64
}

func main() {

	locations := []Location{
		{"Tokyo", 27.5, 140.0},
		{"Sydney", 22.0, 70.0},
		{"London", 18.0, 200.0},
		{"Paris", 21.0, 110.0},
		{"Cape Town", 25.0, 300.0},
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nClimate Insights Tracker")
		fmt.Println("1. View All Locations")
		fmt.Println("2. Locations with Maximum and Minimum Temperatures")
		fmt.Println("3. Compute Average Precipitation")
		fmt.Println("4. Filter Locations by Precipitation")
		fmt.Println("5. Locate by Name")
		fmt.Println("6. Exit")
		fmt.Print("Select an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			showLocations(locations)
		case "2":
			hot, cold := temperatureExtremes(locations)
			fmt.Printf("Hottest Location: %s (%.2f°C)\n", hot.Name, hot.Temperature)
			fmt.Printf("Coldest Location: %s (%.2f°C)\n", cold.Name, cold.Temperature)
		case "3":
			avgPrecipitation := computeAveragePrecipitation(locations)
			fmt.Printf("Average Precipitation: %.2f mm\n", avgPrecipitation)
		case "4":
			fmt.Print("Enter Precipitation Threshold: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			threshold, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid input. Please provide a valid number.")
				continue
			}
			filterByPrecipitation(locations, threshold)
		case "5":
			fmt.Print("Enter Location Name: ")
			locName, _ := reader.ReadString('\n')
			locName = strings.TrimSpace(locName)
			findLocation(locations, locName)
		case "6":
			fmt.Println("Exiting Climate Insights Tracker. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please select a valid choice.")
		}
	}
}

// Displays all locations and their climate data
func showLocations(locations []Location) {
	fmt.Println("\nLocation Climate Data:")
	for _, loc := range locations {
		fmt.Printf("- %s: Temperature = %.2f°C, Precipitation = %.2f mm\n", loc.Name, loc.Temperature, loc.Precipitation)
	}
}

// Identifies locations with maximum and minimum temperatures
func temperatureExtremes(locations []Location) (max, min Location) {
	max, min = locations[0], locations[0]
	for _, loc := range locations {
		if loc.Temperature > max.Temperature {
			max = loc
		}
		if loc.Temperature < min.Temperature {
			min = loc
		}
	}
	return
}

// Computes the average precipitation across all locations
func computeAveragePrecipitation(locations []Location) float64 {
	totalPrecipitation := 0.0
	for _, loc := range locations {
		totalPrecipitation += loc.Precipitation
	}
	return totalPrecipitation / float64(len(locations))
}

// Filters locations with precipitation exceeding the given threshold
func filterByPrecipitation(locations []Location, threshold float64) {
	fmt.Printf("\nLocations with Precipitation Over %.2f mm:\n", threshold)
	found := false
	for _, loc := range locations {
		if loc.Precipitation > threshold {
			fmt.Printf("- %s: %.2f mm\n", loc.Name, loc.Precipitation)
			found = true
		}
	}
	if !found {
		fmt.Println("No locations found exceeding the given threshold.")
	}
}

// Searches for a location by name
func findLocation(locations []Location, locName string) {
	for _, loc := range locations {
		if strings.EqualFold(loc.Name, locName) {
			fmt.Printf("\nLocation Found: %s\nTemperature: %.2f°C\nPrecipitation: %.2f mm\n", loc.Name, loc.Temperature, loc.Precipitation)
			return
		}
	}
	fmt.Println("Location not found. Please check the name and try again.")
}