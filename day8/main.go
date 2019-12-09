package main

import (
	"fmt"
	util "github.com/adventofcode"
	"strconv"
	"strings"
)

const (
	Width            = 25
	Height           = 6
	Pixels           = Width * Height
	BlackColor       = 0
	TransparentColor = 2
)

type Layer [Height][Width]int
type Layers []Layer

func main() {
	lines := util.ReadLines("day8/input.txt")

	layers := toLayers(lines[0])
	layer := fewestZeoDigitLayer(layers)
	fmt.Printf("Part 1: %v\n", layer.numberOfDigit(1)*layer.numberOfDigit(2))
	fmt.Printf("Part 2:\n%v\n", layers.decode())
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

func (l Layer) String() string {
	var sb strings.Builder
	for _, row := range l {
		for _, pixel := range row {
			if pixel == BlackColor {
				sb.WriteString("  ")
			} else {
				sb.WriteRune('◯')
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (layers Layers) decode() string {
	var decodedLayer Layer
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			for k := 0; k < len(layers); k++ {
				if layers[k][i][j] != TransparentColor {
					decodedLayer[i][j] = layers[k][i][j]
					break
				}
			}
		}
	}

	return decodedLayer.String()
}

func fewestZeoDigitLayer(layers Layers) Layer {
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

func toLayers(input string) Layers {
	layers := make(Layers, 0)
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
		layers[layerIndex][heightIndex][widthIndex] = digit
	}
	return layers
}
