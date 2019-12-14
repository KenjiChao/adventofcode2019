package main

import (
	"container/list"
	"fmt"
	util "github.com/adventofcode"
	"math"
	"strconv"
	"strings"
)

type Reaction struct {
	n    int
	need map[string]int
}

const Src = "ORE"
const Dst = "FUEL"

func main() {
	lines := util.ReadLines("day14/input.txt")

	reaction := make(map[string]Reaction)
	indegree := make(map[string]int)

	for _, line := range lines {
		io := strings.Split(line, " => ")
		n, out := expand(io[1])

		r := Reaction{
			n:    n,
			need: make(map[string]int),
		}

		for _, s := range strings.Split(io[0], ", ") {
			nin, in := expand(s)
			r.need[in] = nin

			indegree[in]++
		}

		reaction[out] = r
	}

	fmt.Println(cost(1, indegree, reaction))
}

func cost(nFuel int, indegree map[string]int, reaction map[string]Reaction) int {
	indegree[Dst] = 0
	req := map[string]int{Dst: nFuel}
	q := list.New()
	for k, v := range indegree {
		if v == 0 {
			q.PushBack(k)
		}
	}

	for cLen := q.Len(); cLen != 0; {
		for i := 0; i < cLen; i++ {
			element := q.Front()
			chemical := element.Value.(string)
			if chemical == Src {
				return req[Src]
			}

			r := reaction[chemical]
			times := int(math.Ceil(float64(req[chemical]) / float64(r.n)))
			for name, count := range r.need {
				req[name] += times * count
				indegree[name]--
				if indegree[name] == 0 {
					q.PushBack(name)
				}
			}
			q.Remove(element)
		}
	}

	return req["ORE"]

}

func expand(s string) (int, string) {
	arr := strings.Split(s, " ")
	n, _ := strconv.Atoi(arr[0])
	return n, arr[1]
}
