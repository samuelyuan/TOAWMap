package graphics

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/fogleman/gg"
	"github.com/samuelyuan/TOAWMap/fileio"
)

const (
	radius = 10.0
)

func getImagePosition(i int, j int) (float64, float64) {
	angle := math.Pi / 6
	x := radius + float64(j)*radius*(1+math.Sin(angle))
	y := (radius * 1.5) + float64(i)*(2*radius*math.Cos(angle))
	if j%2 == 1 {
		y -= radius * math.Cos(angle)
	}
	return x, y
}

func IsTileEmpty(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 != 0
}

func IsTileSand(tileData *fileio.TileData) bool {
	// Index 1 = arid, 2 = sandy, 3 = r_sandy, 4 = badlands
	return tileData.Data[38]&0x10 == 0 &&
		(tileData.Data[1] != 0 || tileData.Data[2] != 0 || tileData.Data[3] != 0 || tileData.Data[4] != 0)
}

func IsTileHills(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 == 0 && tileData.Data[5] != 0
}

func IsTileMountains(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 == 0 && tileData.Data[6] != 0
}

func IsTileImpassable(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 == 0 && tileData.Data[7] != 0
}

func IsTileMarsh(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 == 0 && tileData.Data[8] != 0
}

func IsTileFloodedMarsh(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 == 0 && tileData.Data[9] != 0
}

func IsTileShallowWater(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 == 0 && tileData.Data[10] != 0
}

func IsTileDeepWater(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 == 0 && tileData.Data[11] != 0
}

func IsTileUrban(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 == 0 &&
		(tileData.Data[14] != 0 || tileData.Data[15] != 0 || tileData.Data[16] != 0 || tileData.Data[17] != 0)
}

func IsTileForest(tileData *fileio.TileData) bool {
	// Index 26 = c_forest, 27 = d_forest, 28 = m_forest, 29 = t_forest
	return tileData.Data[38]&0x10 == 0 &&
		(tileData.Data[26] != 0 || tileData.Data[27] != 0 || tileData.Data[28] != 0 || tileData.Data[29] != 0)
}

func DoesTileHaveRiver(tileData *fileio.TileData) bool {
	// Index 21 = dry river, 22 = river
	return tileData.Data[38]&0x10 == 0 && tileData.Data[22] != 0
}

func DoesTileHaveMajorRiver(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 == 0 && tileData.Data[23] != 0
}

func DoesTileHaveRailroad(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 == 0 && tileData.Data[33] != 0
}

func DoesTileHaveRoad(tileData *fileio.TileData) bool {
	return tileData.Data[38]&0x10 == 0 && tileData.Data[31] != 0
}

func drawTiles(dc *gg.Context, allTileData [][]*fileio.TileData, mapHeight int, mapWidth int) {
	for i := 0; i < mapHeight; i++ {
		for j := 0; j < mapWidth; j++ {
			x, y := getImagePosition(i, j)
			dc.DrawRegularPolygon(6, x, y, radius, 0)

			tileData := allTileData[i][j]

			if IsTileEmpty(tileData) {
				dc.SetRGB255(0, 0, 0)
			} else if IsTileImpassable(tileData) {
				dc.SetRGB255(67, 65, 68)
			} else if IsTileDeepWater(tileData) {
				dc.SetRGB255(21, 43, 116)
			} else if IsTileShallowWater(tileData) {
				dc.SetRGB255(64, 93, 166)
			} else if IsTileForest(tileData) {
				dc.SetRGB255(78, 116, 53)
			} else if IsTileMountains(tileData) {
				dc.SetRGB255(169, 154, 133)
			} else if IsTileHills(tileData) {
				dc.SetRGB255(149, 132, 58)
			} else if IsTileSand(tileData) {
				dc.SetRGB255(189, 159, 86)
			} else if IsTileFloodedMarsh(tileData) {
				dc.SetRGB255(137, 172, 139)
			} else if IsTileMarsh(tileData) {
				dc.SetRGB255(122, 148, 71)
			} else {
				// Grass tile as default
				dc.SetRGB255(146, 155, 59)
			}
			dc.Fill()

			if IsTileUrban(tileData) {
				dc.DrawRectangle(x-(radius/5), y-(radius/5), radius/2, radius/2)
				dc.SetRGB255(255, 255, 255)
				dc.Fill()
			}
		}
	}
}

func drawTileRoutes(dc *gg.Context, routeData byte, x float64, y float64) {
	// Bit mapping: 1=North, 2=Northeast, 4=Southeast, 8=South, 16=Southwest, 32=Northwest
	for n := 0; n < 6; n++ {
		if ((routeData >> n) & 1) != 0 {
			drawRoute(dc, x, y, n)
		}
	}
}

func drawRoute(dc *gg.Context, x float64, y float64, directionIndex int) {
	angle := (math.Pi / 2) - float64(directionIndex)*(math.Pi/3)
	edgeX := x + radius*math.Cos(angle)
	edgeY := y - radius*math.Sin(angle)
	dc.DrawLine(x, y, edgeX, edgeY)
	dc.Stroke()
}

func drawRiversAndRoads(dc *gg.Context, allTileData [][]*fileio.TileData, mapHeight int, mapWidth int) {
	for i := 0; i < mapHeight; i++ {
		for j := 0; j < mapWidth; j++ {
			tileData := allTileData[i][j]

			// The body of water covers the river
			if IsTileDeepWater(tileData) || IsTileShallowWater(tileData) {
				continue
			}

			x, y := getImagePosition(i, j)

			if DoesTileHaveRiver(tileData) {
				dc.SetRGB255(91, 130, 150)
				routeData := tileData.Data[22]
				drawTileRoutes(dc, routeData, x, y)
			}
			if DoesTileHaveMajorRiver(tileData) {
				dc.SetRGB255(57, 82, 148)
				routeData := tileData.Data[23]
				drawTileRoutes(dc, routeData, x, y)
			}
			if DoesTileHaveRoad(tileData) {
				dc.SetRGB255(195, 167, 87)
				routeData := tileData.Data[31]
				drawTileRoutes(dc, routeData, x, y)
			}
			if DoesTileHaveRailroad(tileData) {
				dc.SetRGB255(102, 91, 72)
				routeData := tileData.Data[33]
				drawTileRoutes(dc, routeData, x, y)
			}
		}
	}
}

func drawUnits(dc *gg.Context, mapData *fileio.TOAWMapData) {
	rand.Seed(time.Now().UnixNano())

	groupColorMap := initGroupColorMap()
	unitTeamMap := make(map[int][]string)

	for i := 0; i < len(mapData.AllUnitData); i++ {
		unitData := mapData.AllUnitData[i]
		x := int(unitData.X)
		y := int(unitData.Y)
		if x == 999 || y == 999 {
			continue
		}
		imageX, imageY := getImagePosition(y, x)

		team := int((unitData.UnitColorAndType + ((unitData.UnitColorAndType >> 31) & 0x7f)) >> 7)
		unitType := int(unitData.UnitColorAndType & 0x8000007f)

		if _, ok := groupColorMap[int(team)]; !ok {
			fmt.Println("Generating random color for group", team)
			groupColorMap[int(team)] = GroupColor{
				OuterColor: color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255},
				InnerColor: color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255},
			}
		}
		if _, ok := unitTeamMap[int(team)]; !ok {
			unitTeamMap[int(team)] = make([]string, 0)
		}
		teamName := string(strings.Split(string(unitData.Name[:]), "\x00")[0])
		unitTeamMap[int(team)] = append(unitTeamMap[int(team)], fmt.Sprintf("%v (type: %v)", teamName, unitType))

		outerColor := groupColorMap[int(team)].OuterColor
		dc.DrawRectangle(imageX-(radius*2/3), imageY-(radius*2/3), radius*4/3, radius*4/3)
		dc.SetRGB255(int(outerColor.R), int(outerColor.G), int(outerColor.B))
		dc.Fill()

		innerColor := groupColorMap[int(team)].InnerColor
		dc.DrawRectangle(imageX-(radius*1/3), imageY-(radius*1/3), radius*2/3, radius*2/3)
		dc.SetRGB255(int(innerColor.R), int(innerColor.G), int(innerColor.B))
		dc.Fill()

		// dc.SetRGB255(255, 255, 255)
		// name := string(strings.Split(string(unitData.Name[:]), "\x00")[0])
		// dc.DrawString(name, imageX-(5.0*float64(len(name))/2.0), imageY-(radius*1.5))
	}
	fmt.Println("Team data:")
	keys := make([]int, 0, len(unitTeamMap))
	for k := range unitTeamMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, teamId := range keys {
		unitList := unitTeamMap[teamId]
		groupColor := groupColorMap[teamId]
		fmt.Printf("Group %d: %d units (color: %v)\n", teamId, len(unitList), groupColor)
		for i, unit := range unitList {
			fmt.Printf("  %d. %s\n", i+1, unit)
		}
	}
}

func drawLocations(dc *gg.Context, mapData *fileio.TOAWMapData) {
	for i := 0; i < len(mapData.AllLocationData); i++ {
		location := mapData.AllLocationData[i]
		x := int(location.X)
		y := int(location.Y)
		locationName := string(strings.Split(string(location.Name[:]), "\x00")[0])
		if x == 999 {
			// Empty location data
			continue
		}

		dc.SetRGB255(255, 255, 255)

		imageX, imageY := getImagePosition(y, x)
		dc.DrawString(locationName, imageX-(5.0*float64(len(locationName))/2.0), imageY-(radius*1.25))
	}
}

func DrawMap(mapData *fileio.TOAWMapData, outputFilename string) {
	maxImageWidth, maxImageHeight := getImagePosition(mapData.MapHeight, mapData.MapWidth)
	dc := gg.NewContext(int(maxImageWidth), int(maxImageHeight))
	fmt.Printf("Rendering map: %dx%d tiles\n", mapData.MapWidth, mapData.MapHeight)

	drawTiles(dc, mapData.AllTileData, int(mapData.MapHeight), int(mapData.MapWidth))
	drawRiversAndRoads(dc, mapData.AllTileData, int(mapData.MapHeight), int(mapData.MapWidth))
	drawUnits(dc, mapData)
	drawLocations(dc, mapData)

	dc.SavePNG(outputFilename)
	fmt.Println("Saved image to", outputFilename)
}
