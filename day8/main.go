package main

import (
	"fmt"
	util "github.com/adventofcode"
	"strconv"
)

const (
	Width  = 25
	Height = 6
	Pixels = Width * Height
)

type Layer [Width][Height]int

func main() {
	lines := util.ReadLines("day8/input.txt")

	layers := toLayers(lines[0])
	layer := fewestZeoDigitLayer(layers)
	fmt.Println(layer)
	fmt.Println("Part 1:", layer.numberOfDigit(1)*layer.numberOfDigit(2))
}

func (l *Layer) numberOfDigit(digit int) int {
	num := 0
	for _, row := range l {
		for _, pixel := range row {
			if pixel == digit {
				num++
			}
		}
	}
	return num
}

func fewestZeoDigitLayer(layers []Layer) Layer {
	fewestZeroDigits := Pixels
	var fewestZeroDigitsLayer Layer
	for _, layer := range layers {
		numberOfZeros := layer.numberOfDigit(0)
		if numberOfZeros < fewestZeroDigits {
			fewestZeroDigits = numberOfZeros
			fewestZeroDigitsLayer = layer
		}
	}
	return fewestZeroDigitsLayer
}

func toLayers(input string) []Layer {
	layers := make([]Layer, 0)
	for i, r := range []rune(input) {
		layerIndex := i / Pixels
		if layerIndex == len(layers) {
			layers = append(layers, *new(Layer))
		}
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}

		widthIndex := (i % Pixels) % Width
		heightIndex := (i % Pixels) / Width
		layers[layerIndex][widthIndex][heightIndex] = digit
	}
	return layers
}
