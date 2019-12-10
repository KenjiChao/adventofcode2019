package main

import (
	"crypto/sha1"
	"fmt"
	util "github.com/adventofcode"
	"log"
	"math"
	"sort"
	"strconv"
)

type Point struct {
	x             int
	y             int
	isAsteroid    bool
	numOfAsteroid int
}

type Diff struct {
	dx    int
	dy    int
	round int
	point Point
}

type ByUpAndClockwise []Diff

func (c ByUpAndClockwise) Len() int      { return len(c) }
func (c ByUpAndClockwise) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c ByUpAndClockwise) Less(i, j int) bool {
	if c[i].round != c[j].round {
		return c[i].round < c[j].round
	}

	if c[i].Degree() != c[j].Degree() {
		return c[i].Degree() < c[j].Degree()
	}
	return c[i].Length() < c[j].Length()
}

func (d Diff) Degree() float64 {
	if d.dx == 0 && d.dy == 0 {
		log.Fatal("(0, 0) is not allowed")
	}

	if d.dx == 0 {
		if d.dy > 0 {
			return 180
		} else {
			return 0
		}
	}
	if d.dy == 0 {
		if d.dx > 0 {
			return 90
		} else {
			return 270
		}
	}

	angle := math.Atan2(math.Abs(float64(d.dx)), math.Abs(float64(d.dy))) * 180 / math.Pi
	if d.dx > 0 && d.dy > 0 {
		angle = 180 - angle
	}
	if d.dx < 0 && d.dy > 0 {
		angle += 180
	}
	if d.dx < 0 && d.dy < 0 {
		angle = 360 - angle
	}

	return angle
}

func (d Diff) Length() float64 {
	return math.Pow(float64(d.dx), 2) + math.Pow(float64(d.dy), 2)
}

func main() {
	lines := util.ReadLines("day10/input.txt")

	points := toAsteroids(lines)
	a := MaxAsteroidsDetected(points)
	fmt.Println("Part 1:", a.numOfAsteroid)
	v := VaporizedAsteroid(points, a, 199)
	fmt.Println(v)
	fmt.Println("Part 2:", v.x*100+v.y)
}

func VaporizedAsteroid(points [][]Point, origin Point, i int) Point {
	diffs := make([]Diff, 0)
	hashSet := make(map[string]int)

	for i := range points {
		for j := range points[i] {
			if points[i][j] != origin && points[i][j].isAsteroid {
				dx := points[i][j].x - origin.x
				dy := points[i][j].y - origin.y
				gcd := gcd(dx, dy)
				hash := hash(dx/gcd, dy/gcd)
				diffs = append(diffs, Diff{
					dx:    dx,
					dy:    dy,
					round: hashSet[hash],
					point: points[i][j],
				})
				hashSet[hash] = hashSet[hash] + 1
			}
		}
	}
	sort.Sort(ByUpAndClockwise(diffs))
	return diffs[i].point
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

func MaxAsteroidsDetected(asteroids [][]Point) Point {
	max := 0
	var argmax Point
	for i := range asteroids {
		for j := range asteroids[i] {
			detected := asteroidsDetected(asteroids, asteroids[i][j])
			asteroids[i][j].numOfAsteroid = detected
			if detected > max {
				max = detected
				argmax = asteroids[i][j]
			}
		}
	}
	return argmax
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
