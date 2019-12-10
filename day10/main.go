package main

import (
	"crypto/sha1"
	"fmt"
	util "github.com/adventofcode"
	"math"
	"strconv"
)

type Point struct {
	x          int
	y          int
	isAsteroid bool
}

func main() {
	lines := util.ReadLines("day10/input.txt")

	points := toAsteroids(lines)
	fmt.Println("Part 1:", MaxAsteroidsDetected(points))
}

func toAsteroids(lines []string) [][]Point {
	rows := len(lines)
	cols := len(lines[0])
	asteroids := make([][]Point, rows)
	for i := range asteroids {
		asteroids[i] = make([]Point, cols)
	}

	for i := range lines {
		for j, r := range lines[i] {
			asteroids[i][j] = Point{
				x:          j,
				y:          i,
				isAsteroid: r == '#',
			}
		}
	}

	return asteroids
}

func MaxAsteroidsDetected(asteroids [][]Point) int {
	max := 0
	for i := range asteroids {
		for j := range asteroids[i] {
			detected := asteroidsDetected(asteroids, asteroids[i][j])
			if detected > max {
				max = detected
			}
		}
	}
	return max
}

func asteroidsDetected(asteroids [][]Point, a Point) int {
	if !a.isAsteroid {
		return 0
	}
	hashSet := make(map[string]bool)
	for i := range asteroids {
		for j := range asteroids[i] {
			if asteroids[i][j] != a && asteroids[i][j].isAsteroid {
				dx := asteroids[i][j].x - a.x
				dy := asteroids[i][j].y - a.y
				gcd := gcd(dx, dy)
				hashSet[hash(dx/gcd, dy/gcd)] = true
			}
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
