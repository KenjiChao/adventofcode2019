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
	NumberOfPixels   = Width * Height
	BlackColor       = 0
	TransparentColor = 2
)

type Pixels [Height][Width]int
type Layer struct {
	*Pixels
	numberOfDigits map[int]int
}

func NewLayer() *Layer {
	return &Layer{
		Pixels:         new(Pixels),
		numberOfDigits: make(map[int]int),
	}
}

func (l *Layer) numberOfDigit(digit int) int {
	if val, ok := l.numberOfDigits[digit]; ok {
		fmt.Println("Use cached value!")
		return val
	}
	num := 0
	for _, row := range l.Pixels {
		for _, pixel := range row {
			if pixel == digit {
				num++
			}
		}
	}
	l.numberOfDigits[digit] = num
	return num
}

func (l Layer) String() string {
	var sb strings.Builder
	for _, row := range l.Pixels {
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

type Layers []Layer

func (layers Layers) decode() string {
	decodedLayer := NewLayer()
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			for k := 0; k < len(layers); k++ {
				if layers[k].Pixels[i][j] != TransparentColor {
					decodedLayer.Pixels[i][j] = layers[k].Pixels[i][j]
					break
				}
			}
		}
	}

	return decodedLayer.String()
}

func main() {
	lines := util.ReadLines("day8/input.txt")

	layers := toLayers(lines[0])
	layer := fewestZeoDigitLayer(layers)
	fmt.Printf("First time: %v\n", layer.numberOfDigit(0))
	fmt.Printf("Second time: %v\n", layer.numberOfDigit(0))
	fmt.Printf("Part 1: %v\n", layer.numberOfDigit(1)*layer.numberOfDigit(2))
	fmt.Printf("Part 2:\n%v\n", layers.decode())
}

func fewestZeoDigitLayer(layers Layers) Layer {
	fewestZeroDigits := NumberOfPixels
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
		layerIndex := i / NumberOfPixels
		if layerIndex == len(layers) {
			layers = append(layers, *NewLayer())
		}
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}

		widthIndex := (i % NumberOfPixels) % Width
		heightIndex := (i % NumberOfPixels) / Width
		layers[layerIndex].Pixels[heightIndex][widthIndex] = digit
	}
	return layers
}
