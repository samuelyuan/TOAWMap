package fileio

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/samuelyuan/TOAWMap/blast"
)

type TOAWMapHeader struct {
	Header                  [16]byte
	UnknownInt1             uint32
	MapTitle                [264]byte
	Version                 uint32
	UnknownInt3             uint32
	MapDescription          [8192]byte
	EndMessageTeam1Victory1 [8192]byte
	EndMessageTeam1Victory2 [8192]byte
	EndMessageDraw1         [8192]byte
	EndMessageTeam2Victory  [8192]byte
	UnknownBlock1           [8192]byte
	EndMessageDraw2         [8192]byte
	UnknownBlock2           [8192]byte
	TeamGoesFirst           uint32
	TeamNameBlock3          [36]byte
}

type LocationData struct {
	X    int32
	Y    int32
	Name [28]byte
}

type TileData struct {
	Data []byte // Size is hardcoded to 47 bytes
}

type UnitData struct {
	Name                     [20]byte
	UnknownBlock1            [60]uint32
	UnknownBlock2            [48]byte
	Unknown_x134             [4]byte
	UnitColorAndType         uint32 // this can't be split up into separate bytes due to the way the bits are stored
	Unknown_x13c             uint32
	Unknown_x140             uint32
	Proficiency              uint32
	Readiness                uint32
	SupplyLevel              uint32
	Unknown_x150             uint32
	OtherUnitIndexOnSameTile uint32 // 1000 or 4000 (depends on max number of units) doesn't refer to any unit
	X                        int32
	Y                        int32
	Unknown_x160             uint32
	Unknown_x164             uint32
	Unknown_x168             uint32
	Unknown_x16c             uint32
	Unknown_x170             uint32
	Unknown_x174             uint32
	UnitIndex                uint32
	UnknownBlock4            [12]byte
}

type TeamNameData struct {
	CountryName   [17]byte
	ForceName     [35]byte
	Proficiency   uint32
	SupplyLevel   uint32
	CountryFlagId uint32
}

type TOAWMapData struct {
	Version            int
	DecompressedBlocks [][]byte
	AllLocationData    []LocationData
	AllTeamNameData    []*TeamNameData
	AllTileData        [][]*TileData
	AllUnitData        []*UnitData
	MapWidth           int
	MapHeight          int
}

func dumpData(inputData []byte, outputFilename string) {
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		log.Fatal("Failed to create file: ", err)
	}
	defer outputFile.Close()

	numBytesWritten, err := outputFile.Write(inputData)
	if err != nil {
		log.Fatal("Failed to write data: ", err)
	}
	fmt.Println("Saved", numBytesWritten, "bytes to", outputFilename)
}

func ReadTOAWScenario(filename string) (*TOAWMapData, error) {
	inputFile, err := os.Open(filename)
	defer inputFile.Close()
	if err != nil {
		log.Fatal("Failed to load map: ", err)
		return nil, err
	}

	// If the game is a TOAW4 scenario, it will be compressed using Gzip
	// The older games will have TOAC in the header
	var gzipHeader [3]byte
	_, err = io.ReadFull(inputFile, gzipHeader[:])
	if err != nil {
		return nil, err
	}
	// Reset file pointer
	inputFile.Seek(0, 0)
	// Gzip header: The magic number is 0x1f8band the compression method is 08 for DEFLATE
	if gzipHeader[0] == 0x1f && gzipHeader[1] == 0x8b && gzipHeader[2] == 0x08 {
		return ReadTOAW4Scenario(inputFile)
	}

	fi, err := inputFile.Stat()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fileLength := fi.Size()
	streamReader := io.NewSectionReader(inputFile, int64(0), fileLength)

	mapHeader := TOAWMapHeader{}
	if err := binary.Read(streamReader, binary.LittleEndian, &mapHeader); err != nil {
		return nil, err
	}

	fmt.Println("Version:", mapHeader.Version)
	fmt.Println("Map title:", string(mapHeader.MapTitle[:]))
	fmt.Println("Map description:", string(bytes.Trim(mapHeader.MapDescription[:], "\x00")))
	fmt.Println("EndMessageTeam1Victory1:", string(bytes.Trim(mapHeader.EndMessageTeam1Victory1[:], "\x00")))
	fmt.Println("EndMessageTeam1Victory2:", string(bytes.Trim(mapHeader.EndMessageTeam1Victory2[:], "\x00")))
	fmt.Println("EndMessageDraw1:", string(bytes.Trim(mapHeader.EndMessageDraw1[:], "\x00")))
	fmt.Println("EndMessageTeam2Victory:", string(bytes.Trim(mapHeader.EndMessageTeam2Victory[:], "\x00")))
	fmt.Println("EndMessageDraw2:", string(bytes.Trim(mapHeader.EndMessageDraw2[:], "\x00")))
	fmt.Println("Team goes first:", mapHeader.TeamGoesFirst)

	totalBlocks := 12
	// Later version has an additional block
	if mapHeader.Version >= 0x79 {
		totalBlocks = 13
	}

	decompressedBlocks := make([][]byte, totalBlocks)
	for i := 0; i < totalBlocks; i++ {
		blockSize := uint32(0)
		if err := binary.Read(streamReader, binary.LittleEndian, &blockSize); err != nil {
			return nil, err
		}
		fmt.Println("Block", i, "size:", blockSize)

		blockData := make([]byte, blockSize)
		if err := binary.Read(streamReader, binary.LittleEndian, &blockData); err != nil {
			return nil, err
		}

		// The blocks stored in this file are compressed using PKWare Compression Library
		// Decompress data
		b := bytes.NewReader(blockData)
		r, err := blast.NewReader(b)
		if err != nil {
			return nil, err
		}

		decompressedData, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}
		// dumpData(decompressedData, fmt.Sprintf("block%v.txt", i))
		decompressedBlocks[i] = decompressedData
		r.Close()
	}

	unknownData1 := make([]byte, 256)
	if mapHeader.Version == 0x42 { // TOAW1
		unknownData1 = make([]byte, 232)
	}
	if err := binary.Read(streamReader, binary.LittleEndian, &unknownData1); err != nil {
		return nil, err
	}
	fmt.Println("unknownData1:", unknownData1)

	blockSize := uint32(0)
	if err := binary.Read(streamReader, binary.LittleEndian, &blockSize); err != nil {
		return nil, err
	}
	fmt.Println("Last block size:", blockSize)

	// This block is also compressed, but the result is zero bytes
	blockData := make([]byte, blockSize)
	if err := binary.Read(streamReader, binary.LittleEndian, &blockData); err != nil {
		return nil, err
	}

	version := int(mapHeader.Version)
	mapWidth := 1 + int(binary.LittleEndian.Uint32(unknownData1[0:4]))
	mapHeight := 1 + int(binary.LittleEndian.Uint32(unknownData1[4:8]))

	locationBlockIndex := 10
	if version >= 0x79 {
		locationBlockIndex = 11
	}
	locationBlock := decompressedBlocks[locationBlockIndex]

	mapData := &TOAWMapData{
		Version:            version,
		DecompressedBlocks: decompressedBlocks,
		AllTileData:        GetTileData(decompressedBlocks[1], mapHeight, mapWidth),
		AllLocationData:    GetLocationData(locationBlock),
		AllUnitData:        GetUnitData(decompressedBlocks),
		AllTeamNameData:    GetTeamNameData(decompressedBlocks),
		MapWidth:           mapWidth,
		MapHeight:          mapHeight,
	}
	return mapData, nil
}

func ReadTOAW4Scenario(compressedFile *os.File) (*TOAWMapData, error) {
	decompressedFile, err := gzip.NewReader(compressedFile)
	if err != nil {
		return nil, err
	}
	defer decompressedFile.Close()

	decompressedFileContents, err := ioutil.ReadAll(decompressedFile)
	if err != nil {
		return nil, err
	}

	fmt.Println("Decompressed contents size: ", len(decompressedFileContents))
	streamReader := io.NewSectionReader(bytes.NewReader(decompressedFileContents), int64(0), int64(len(decompressedFileContents)))

	mapHeader := TOAWMapHeader{}
	if err := binary.Read(streamReader, binary.LittleEndian, &mapHeader); err != nil {
		return nil, err
	}

	fmt.Println("Version:", mapHeader.Version)
	fmt.Println("Map title:", string(mapHeader.MapTitle[:]))
	fmt.Println("Map description:", string(bytes.Trim(mapHeader.MapDescription[:], "\x00")))
	fmt.Println("EndMessageTeam1Victory1:", string(bytes.Trim(mapHeader.EndMessageTeam1Victory1[:], "\x00")))
	fmt.Println("EndMessageTeam1Victory2:", string(bytes.Trim(mapHeader.EndMessageTeam1Victory2[:], "\x00")))
	fmt.Println("EndMessageDraw1:", string(bytes.Trim(mapHeader.EndMessageDraw1[:], "\x00")))
	fmt.Println("EndMessageTeam2Victory:", string(bytes.Trim(mapHeader.EndMessageTeam2Victory[:], "\x00")))
	fmt.Println("EndMessageDraw2:", string(bytes.Trim(mapHeader.EndMessageDraw2[:], "\x00")))
	fmt.Println("Team goes first:", mapHeader.TeamGoesFirst)

	unknownData1 := make([]byte, 448)
	if err := binary.Read(streamReader, binary.LittleEndian, &unknownData1); err != nil {
		return nil, err
	}
	fmt.Println("unknownData1:", unknownData1)

	// This is similar to TOAW3
	decompressedBlock0 := make([]byte, 696)
	if err := binary.Read(streamReader, binary.LittleEndian, &decompressedBlock0); err != nil {
		return nil, err
	}

	// This block contains the map data
	decompressedBlock1 := make([]byte, 48*700*700)
	if err := binary.Read(streamReader, binary.LittleEndian, &decompressedBlock1); err != nil {
		return nil, err
	}

	version := int(mapHeader.Version)
	mapWidth := 1 + int(binary.LittleEndian.Uint32(unknownData1[132:136]))
	mapHeight := 1 + int(binary.LittleEndian.Uint32(unknownData1[136:140]))

	fmt.Println("Map width: ", mapWidth, ", map height: ", mapHeight)

	// This block seems to be in the same location
	// 144000 bytes long
	// 36 bytes per location * 4000 locations maximum
	locationBlock := decompressedFileContents[64859168:65003168]

	mapData := &TOAWMapData{
		Version:            version,
		DecompressedBlocks: [][]byte{},
		AllTileData:        GetTileData(decompressedBlock1, mapHeight, mapWidth),
		AllLocationData:    GetLocationData(locationBlock),
		AllUnitData:        []*UnitData{},
		AllTeamNameData:    []*TeamNameData{},
		MapWidth:           mapWidth,
		MapHeight:          mapHeight,
	}
	return mapData, nil
}

func GetLocationData(locationBlock []byte) []LocationData {
	streamReader := io.NewSectionReader(bytes.NewReader(locationBlock), int64(0), int64(len(locationBlock)))
	locationDataSize := 36
	numLocations := len(locationBlock) / locationDataSize
	allLocationData := make([]LocationData, numLocations)
	if err := binary.Read(streamReader, binary.LittleEndian, &allLocationData); err != nil {
		log.Fatal("Failed to read location data:", err)
	}
	return allLocationData
}

func GetTileData(mapBlock []byte, mapHeight int, mapWidth int) [][]*TileData {
	allTileData := make([][]*TileData, mapHeight)
	for i := 0; i < len(allTileData); i++ {
		allTileData[i] = make([]*TileData, mapWidth)
	}

	// The file format keeps the map block a constant size, but the unused data is set to zero bytes
	var tileDataSize, columnDataSize int
	if len(mapBlock) == 48*700*700 {
		// Only for TOAW4 maps
		// The maximum map dimensions were expanded to 700x700
		tileDataSize = 48
		columnDataSize = tileDataSize * 700
	} else if len(mapBlock) == 47*300*300 {
		// The maximum map dimensions for TOAW3 or earlier games is 300x300, but most maps will never reach that size
		tileDataSize = 47
		columnDataSize = tileDataSize * 300
	} else if len(mapBlock) == 47*100*100 {
		// The maximum map dimensions is assumed to be 100x100
		tileDataSize = 47
		columnDataSize = tileDataSize * 100
	} else {
		log.Fatal(fmt.Sprintf("Failed to read tile data block of length %v", len(mapBlock)))
	}
	for x := 0; x < mapWidth; x++ {
		columnStart := x * columnDataSize
		columnEnd := (x + 1) * columnDataSize
		columnData := mapBlock[columnStart:columnEnd]
		streamReader := io.NewSectionReader(bytes.NewReader(columnData), int64(0), int64(len(columnData)))

		for y := 0; y < mapHeight; y++ {
			tileData := make([]byte, tileDataSize)
			if err := binary.Read(streamReader, binary.LittleEndian, &tileData); err != nil {
				log.Fatal("Failed to read tile data:", err)
			}

			allTileData[y][x] = &TileData{
				Data: tileData[:],
			}
		}
	}
	return allTileData
}

func GetTeamNameData(decompressedBlocks [][]byte) []*TeamNameData {
	teamBlockIndex := 4
	teamBlock := decompressedBlocks[teamBlockIndex]

	// There are only 2 teams because this is a 1v1 game
	maxTeams := 2
	allTeamNameData := make([]*TeamNameData, maxTeams)

	streamReader := io.NewSectionReader(bytes.NewReader(teamBlock), int64(0), int64(len(teamBlock)))
	for i := 0; i < maxTeams; i++ {
		teamNameData := TeamNameData{}
		if err := binary.Read(streamReader, binary.LittleEndian, &teamNameData); err != nil {
			log.Fatal("Failed to read team name data:", err)
		}
		allTeamNameData[i] = &teamNameData

		fmt.Println("Team", i, ":", teamNameData)
	}
	return allTeamNameData
}

func GetUnitData(decompressedBlocks [][]byte) []*UnitData {
	unitBlockIndex := 2
	unitBlock := decompressedBlocks[unitBlockIndex]

	// Maximum number of units is 4000, but it can be less in some files
	maxUnits := len(unitBlock) / 392
	fmt.Println("Maximum units: ", maxUnits)
	allUnitData := make([]*UnitData, maxUnits)

	streamReader := io.NewSectionReader(bytes.NewReader(unitBlock), int64(0), int64(len(unitBlock)))

	for i := 0; i < maxUnits; i++ {
		unitData := UnitData{}
		if err := binary.Read(streamReader, binary.LittleEndian, &unitData); err != nil {
			log.Fatal("Failed to read unit data:", err)
		}

		allUnitData[i] = &unitData
	}
	return allUnitData
}
