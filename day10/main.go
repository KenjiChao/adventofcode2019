package main

import (
	"crypto/sha1"
	"fmt"
	util "github.com/adventofcode"
	"math"
	"strconv"
)

type Asteroid struct {
	x int
	y int
}

func main() {
	lines := util.ReadLines("day10/input.txt")

	asteroids := toAsteroids(lines)
	fmt.Println("Part 1:", MaxAsteroidsDetected(asteroids))
}

func toAsteroids(lines []string) []Asteroid {
	asteroids := make([]Asteroid, 0)

	for i := range lines {
		for j, r := range lines[i] {
			if r == '#' {
				asteroids = append(asteroids, Asteroid{x: j, y: i})
			}
		}
	}

	return asteroids
}

func MaxAsteroidsDetected(asteroids []Asteroid) int {
	max := 0
	for _, a := range asteroids {
		detected := asteroidsDetected(asteroids, a)
		if detected > max {
			max = detected
		}
	}
	return max
}

func asteroidsDetected(asteroids []Asteroid, a Asteroid) int {
	hashSet := make(map[string]bool)
	for i := range asteroids {
		if asteroids[i] != a {
			dx := asteroids[i].x - a.x
			dy := asteroids[i].y - a.y
			gcd := gcd(dx, dy)
			hashSet[hash(dx/gcd, dy/gcd)] = true
		}
	}
	return len(hashSet)
}

func hash(a, b int) string {
	h := sha1.New()
	h.Write([]byte(strconv.Itoa(a)))
	h.Write([]byte(strconv.Itoa(b)))
	return string(h.Sum(nil))
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return int(math.Abs(float64(a)))
}
