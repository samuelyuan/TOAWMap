package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/samuelyuan/TOAWMap/fileio"
	"github.com/samuelyuan/TOAWMap/graphics"
)

func main() {
	inputPtr := flag.String("input", "", "Input filename")
	outputPtr := flag.String("output", "output.png", "Output filename")
	flag.Parse()

	inputFilename := *inputPtr
	outputFilename := *outputPtr
	fmt.Println("Input filename: ", inputFilename)
	fmt.Println("Output filename: ", outputFilename)

	mapData, err := fileio.ReadTOAWScenario(inputFilename)
	if err != nil {
		log.Fatal("Failed to read input file: ", err)
	}

	graphics.DrawMap(mapData, outputFilename)
}
