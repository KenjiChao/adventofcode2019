package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalFuelPart1 := 0
	totalFuelPart2 := 0
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		totalFuelPart1 += RequiredFuel(i)
		totalFuelPart2 += RequiredActualFuel(i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total Fuel Part1: ", totalFuelPart1)
	fmt.Println("Total Fuel Part2: ", totalFuelPart2)
}

func RequiredFuel(mass int) int {
	return int(math.Floor(float64(mass/3)) - 2)
}

func RequiredActualFuel(mass int) (total int) {
	fuel := RequiredFuel(mass)
	for fuel > 0 {
		total += fuel
		fuel = RequiredFuel(fuel)
	}
	return
}
