package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
}

func main() {
	inputPtr := flag.String("input", "", "Input filename (.sce or .json)")
	outputPtr := flag.String("output", "output.png", "Output filename")
	modePtr := flag.String("mode", "draw", "Output mode: draw or exportjson")
	helpPtr := flag.Bool("help", false, "Show help information")
	flag.Parse()

	// Show help if requested
	if *helpPtr {
		showHelp()
		return
	}

	// Validate required input
	if *inputPtr == "" {
		fmt.Println("Error: Input filename is required")
		fmt.Println("Use -help for usage information")
		os.Exit(1)
	}

	inputFilename := *inputPtr
	outputFilename := *outputPtr
	
	fmt.Printf("TOAWMap - Processing: %s\n", inputFilename)
	fmt.Printf("Output: %s\n", outputFilename)

	mode := *modePtr
	if mode == "draw" {
		fmt.Println("Reading map data...")
		mapData := loadMapDataFromFile(inputFilename)
		fmt.Println("Generating map image...")
		graphics.DrawMap(mapData, outputFilename)
		fmt.Printf("Map saved to %s\n", outputFilename)
	} else if mode == "exportjson" {
		fmt.Println("Reading map data...")
		mapData := loadMapDataFromFile(inputFilename)
		fmt.Printf("Exporting map to %s...\n", outputFilename)
		fileio.ExportTOAWMapJson(mapData, outputFilename)
		fmt.Printf("Map exported to %s\n", outputFilename)
	} else {
		fmt.Printf("Error: Invalid mode '%s'\n", mode)
		fmt.Println("Valid modes: draw, exportjson")
		fmt.Println("Use -help for more information")
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Println("TOAWMap - The Operational Art of War Map Renderer")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  TOAWMap -input=<file> [-output=<file>] [-mode=<mode>]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -input string")
	fmt.Println("        Input filename (.sce or .json) (required)")
	fmt.Println("  -output string")
	fmt.Println("        Output filename (default: output.png)")
	fmt.Println("  -mode string")
	fmt.Println("        Output mode: draw or exportjson (default: draw)")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  TOAWMap -input=scenario.sce -output=map.png")
	fmt.Println("  TOAWMap -input=scenario.sce -mode=exportjson -output=data.json")
	fmt.Println("  TOAWMap -input=map.json -output=rendered.png")
	fmt.Println()
	fmt.Println("Supported games:")
	fmt.Println("  - The Operational Art of War: Century of Warfare")
	fmt.Println("  - The Operational Art of War III")
	fmt.Println("  - The Operational Art of War IV")
}
