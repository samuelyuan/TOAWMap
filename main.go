package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/samuelyuan/TOAWMap/fileio"
	"github.com/samuelyuan/TOAWMap/graphics"
)

func loadMapDataFromFile(filename string) *fileio.TOAWMapData {
	mapFileExtension := filepath.Ext(filename)
	if strings.ToLower(mapFileExtension) == ".json" {
		fmt.Println("Importing map file from json")
		mapData := fileio.ImportTOAWMapDataFromJson(filename)
		return mapData
	} else {
		fmt.Println("Reading map from file")
		mapData, err := fileio.ReadTOAWScenario(filename)
		if err != nil {
			log.Fatal("Failed to read input file: ", err)
		}
		return mapData
	}

	return nil
}

func main() {
	inputPtr := flag.String("input", "", "Input filename")
	outputPtr := flag.String("output", "output.png", "Output filename")
	modePtr := flag.String("mode", "draw", "Output mode")
	flag.Parse()

	inputFilename := *inputPtr
	outputFilename := *outputPtr
	fmt.Println("Input filename: ", inputFilename)
	fmt.Println("Output filename: ", outputFilename)

	mode := *modePtr
	if mode == "draw" {
		mapData := loadMapDataFromFile(inputFilename)
		graphics.DrawMap(mapData, outputFilename)
	} else if mode == "exportjson" {
		mapData := loadMapDataFromFile(inputFilename)
		fmt.Println("Exporting map to", outputFilename)
		fileio.ExportTOAWMapJson(mapData, outputFilename)
	} else {
		log.Fatal("Invalid output mode: " + mode + ". Mode must be in this list [draw, exportjson].")
	}
}
