package fileio

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type TOAWMapJson struct {
	GameName   string
	FileFormat string
	MapData    *TOAWMapData
}

func ImportTOAWMapDataFromJson(inputFilename string) *TOAWMapData {
	jsonFile, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal("Failed to open json file", err)
	}
	defer jsonFile.Close()

	jsonContents, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var mapJson *TOAWMapJson
	json.Unmarshal(jsonContents, &mapJson)

	if mapJson == nil {
		log.Fatal("The json data in " + inputFilename + " is missing or incorrect")
	}

	return mapJson.MapData
}

func ExportTOAWMapJson(mapData *TOAWMapData, outputFilename string) {
	polytopiaJson := &TOAWMapJson{
		GameName:   "TOAW",
		FileFormat: "TOAW map scenario",
		MapData:    mapData,
	}

	file, err := json.MarshalIndent(polytopiaJson, "", " ")
	if err != nil {
		log.Fatal("Failed to marshal data: ", err)
	}

	err = ioutil.WriteFile(outputFilename, file, 0644)
	if err != nil {
		log.Fatal("Error writing to ", outputFilename)
	}
}
