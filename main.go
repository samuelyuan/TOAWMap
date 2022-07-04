package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
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

type GroupColor struct {
	OuterColor color.RGBA
	InnerColor color.RGBA
}

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

func initGroupColorMap() map[int]GroupColor {
	groupColorMap := make(map[int]GroupColor)
	groupColorMap[0] = GroupColor{
		OuterColor: color.RGBA{0, 107, 189, 255}, // blue
		InnerColor: color.RGBA{222, 0, 41, 255},  // red
	}
	groupColorMap[5] = GroupColor{
		OuterColor: color.RGBA{165, 132, 66, 255},  // brown
		InnerColor: color.RGBA{247, 247, 247, 255}, // white
	}
	groupColorMap[6] = GroupColor{
		OuterColor: color.RGBA{165, 132, 66, 255}, // brown
		InnerColor: color.RGBA{0, 0, 255, 255},    // blue
	}
	groupColorMap[7] = GroupColor{
		OuterColor: color.RGBA{165, 132, 66, 255}, // brown
		InnerColor: color.RGBA{0, 132, 0, 255},    // green
	}
	groupColorMap[8] = GroupColor{
		OuterColor: color.RGBA{165, 132, 66, 255}, // brown
		InnerColor: color.RGBA{0, 0, 107, 255},    // dark blue
	}
	groupColorMap[9] = GroupColor{
		OuterColor: color.RGBA{165, 132, 66, 255}, // brown
		InnerColor: color.RGBA{0, 0, 0, 255},      // black
	}
	groupColorMap[10] = GroupColor{
		OuterColor: color.RGBA{173, 173, 173, 255}, // gray
		InnerColor: color.RGBA{239, 239, 239, 255}, // white
	}
	groupColorMap[11] = GroupColor{
		OuterColor: color.RGBA{173, 173, 173, 255}, // gray
		InnerColor: color.RGBA{0, 0, 0, 255},       // black
	}
	groupColorMap[12] = GroupColor{
		OuterColor: color.RGBA{173, 173, 173, 255}, // gray
		InnerColor: color.RGBA{231, 239, 247, 255}, // light blue
	}
	groupColorMap[13] = GroupColor{
		OuterColor: color.RGBA{173, 173, 173, 255}, // gray
		InnerColor: color.RGBA{148, 165, 148, 255}, // green
	}
	groupColorMap[14] = GroupColor{
		OuterColor: color.RGBA{173, 173, 173, 255}, // gray
		InnerColor: color.RGBA{140, 8, 0, 255},     // red
	}
	groupColorMap[15] = GroupColor{
		OuterColor: color.RGBA{24, 140, 24, 255},   // green
		InnerColor: color.RGBA{198, 214, 181, 255}, // light green
	}
	groupColorMap[19] = GroupColor{
		OuterColor: color.RGBA{0, 132, 0, 255}, // green
		InnerColor: color.RGBA{0, 0, 255, 255}, // blue
	}
	groupColorMap[20] = GroupColor{
		OuterColor: color.RGBA{128, 0, 0, 255},   // maroon
		InnerColor: color.RGBA{115, 99, 82, 255}, // brown
	}
	groupColorMap[21] = GroupColor{
		OuterColor: color.RGBA{128, 0, 0, 255}, // maroon
		InnerColor: color.RGBA{0, 0, 0, 255},   // black
	}
	groupColorMap[22] = GroupColor{
		OuterColor: color.RGBA{128, 0, 0, 255}, // maroon
		InnerColor: color.RGBA{99, 0, 0, 255},  // darker red
	}
	groupColorMap[23] = GroupColor{
		OuterColor: color.RGBA{128, 0, 0, 255},   // maroon
		InnerColor: color.RGBA{99, 123, 66, 255}, // green
	}
	groupColorMap[24] = GroupColor{
		OuterColor: color.RGBA{128, 0, 0, 255}, // maroon
		InnerColor: color.RGBA{255, 0, 0, 255}, // red
	}
	groupColorMap[30] = GroupColor{
		OuterColor: color.RGBA{247, 247, 247, 255}, // white
		InnerColor: color.RGBA{123, 198, 255, 255}, // light blue
	}
	groupColorMap[31] = GroupColor{
		OuterColor: color.RGBA{247, 247, 247, 255}, // white
		InnerColor: color.RGBA{255, 255, 0, 255},   // yellow
	}
	groupColorMap[32] = GroupColor{
		OuterColor: color.RGBA{247, 247, 247, 255}, // white
		InnerColor: color.RGBA{255, 247, 0, 255},   // yellow
	}
	groupColorMap[33] = GroupColor{
		OuterColor: color.RGBA{247, 247, 247, 255}, // white
		InnerColor: color.RGBA{156, 189, 148, 255}, // green
	}
	groupColorMap[34] = GroupColor{
		OuterColor: color.RGBA{247, 247, 247, 255}, // white
		InnerColor: color.RGBA{189, 189, 189, 255}, // gray
	}
	groupColorMap[35] = GroupColor{
		OuterColor: color.RGBA{239, 222, 0, 255},  // yellow
		InnerColor: color.RGBA{255, 247, 99, 255}, // light yellow
	}
	groupColorMap[36] = GroupColor{
		OuterColor: color.RGBA{239, 222, 0, 255},   // yellow
		InnerColor: color.RGBA{165, 214, 148, 255}, // green
	}
	groupColorMap[37] = GroupColor{
		OuterColor: color.RGBA{239, 222, 0, 255},   // yellow
		InnerColor: color.RGBA{112, 194, 240, 255}, // light blue
	}
	groupColorMap[38] = GroupColor{
		OuterColor: color.RGBA{239, 222, 0, 255},   // yellow
		InnerColor: color.RGBA{206, 189, 156, 255}, // beige
	}
	groupColorMap[39] = GroupColor{
		OuterColor: color.RGBA{239, 222, 0, 255},   // yellow
		InnerColor: color.RGBA{247, 247, 247, 255}, // white
	}
	groupColorMap[40] = GroupColor{
		OuterColor: color.RGBA{198, 181, 132, 255}, // tan
		InnerColor: color.RGBA{239, 239, 65, 255},  // yellow
	}
	groupColorMap[41] = GroupColor{
		OuterColor: color.RGBA{198, 181, 132, 255}, // tan
		InnerColor: color.RGBA{0, 0, 255, 255},     // blue
	}
	groupColorMap[42] = GroupColor{
		OuterColor: color.RGBA{198, 181, 132, 255}, // tan
		InnerColor: color.RGBA{123, 198, 255, 255}, // light blue
	}
	groupColorMap[43] = GroupColor{
		OuterColor: color.RGBA{198, 181, 132, 255}, // tan
		InnerColor: color.RGBA{255, 0, 0, 255},     // red
	}
	groupColorMap[44] = GroupColor{
		OuterColor: color.RGBA{198, 181, 132, 255}, // tan
		InnerColor: color.RGBA{255, 255, 0, 255},   // yellow
	}
	groupColorMap[45] = GroupColor{
		OuterColor: color.RGBA{148, 165, 66, 255},  // light green
		InnerColor: color.RGBA{189, 189, 189, 255}, // gray
	}
	groupColorMap[49] = GroupColor{
		OuterColor: color.RGBA{148, 165, 66, 255},  // light green
		InnerColor: color.RGBA{247, 247, 247, 255}, // white
	}
	groupColorMap[50] = GroupColor{
		OuterColor: color.RGBA{132, 140, 66, 255},  // dark green
		InnerColor: color.RGBA{247, 239, 115, 255}, // yellow
	}
	groupColorMap[51] = GroupColor{
		OuterColor: color.RGBA{132, 140, 66, 255}, // dark green
		InnerColor: color.RGBA{255, 0, 0, 255},    // red
	}
	groupColorMap[52] = GroupColor{
		OuterColor: color.RGBA{132, 140, 66, 255}, // dark green
		InnerColor: color.RGBA{0, 0, 255, 255},    // blue
	}
	groupColorMap[53] = GroupColor{
		OuterColor: color.RGBA{132, 140, 66, 255},  // dark green
		InnerColor: color.RGBA{123, 198, 255, 255}, // light blue
	}
	groupColorMap[54] = GroupColor{
		OuterColor: color.RGBA{132, 140, 66, 255}, // dark green
		InnerColor: color.RGBA{0, 0, 132, 255},    // dark blue
	}
	groupColorMap[55] = GroupColor{
		OuterColor: color.RGBA{82, 156, 255, 255}, // light blue
		InnerColor: color.RGBA{189, 24, 24, 255},  // red
	}
	groupColorMap[56] = GroupColor{
		OuterColor: color.RGBA{82, 156, 255, 255},  // light blue
		InnerColor: color.RGBA{239, 239, 239, 255}, // white
	}
	groupColorMap[57] = GroupColor{
		OuterColor: color.RGBA{82, 156, 255, 255}, // light blue
		InnerColor: color.RGBA{33, 123, 214, 255}, // blue
	}
	groupColorMap[58] = GroupColor{
		OuterColor: color.RGBA{82, 156, 255, 255},  // light blue
		InnerColor: color.RGBA{189, 222, 247, 255}, // light blue
	}
	groupColorMap[59] = GroupColor{
		OuterColor: color.RGBA{82, 156, 255, 255}, // light blue
		InnerColor: color.RGBA{8, 107, 181, 255},  // dark blue
	}
	groupColorMap[60] = GroupColor{
		OuterColor: color.RGBA{189, 206, 189, 255}, // mint
		InnerColor: color.RGBA{107, 107, 107, 255}, // gray
	}
	groupColorMap[61] = GroupColor{
		OuterColor: color.RGBA{189, 206, 189, 255}, // mint
		InnerColor: color.RGBA{247, 247, 247, 255}, // white
	}
	groupColorMap[62] = GroupColor{
		OuterColor: color.RGBA{189, 206, 189, 255}, // mint
		InnerColor: color.RGBA{156, 173, 66, 255},  // green
	}
	groupColorMap[63] = GroupColor{
		OuterColor: color.RGBA{189, 206, 189, 255}, // mint
		InnerColor: color.RGBA{255, 255, 156, 255}, // light yellow
	}
	groupColorMap[64] = GroupColor{
		OuterColor: color.RGBA{189, 206, 189, 255}, // mint
		InnerColor: color.RGBA{247, 189, 107, 255}, // orange
	}
	groupColorMap[65] = GroupColor{
		OuterColor: color.RGBA{156, 123, 41, 255}, // light brown
		InnerColor: color.RGBA{165, 99, 24, 255},  // brown
	}
	groupColorMap[66] = GroupColor{
		OuterColor: color.RGBA{156, 123, 41, 255}, // light brown
		InnerColor: color.RGBA{82, 132, 206, 255}, // blue
	}
	groupColorMap[67] = GroupColor{
		OuterColor: color.RGBA{156, 123, 41, 255}, // light brown
		InnerColor: color.RGBA{33, 148, 123, 255}, // turquoise
	}
	groupColorMap[68] = GroupColor{
		OuterColor: color.RGBA{156, 123, 41, 255},  // light brown
		InnerColor: color.RGBA{222, 214, 173, 255}, // beige
	}
	groupColorMap[69] = GroupColor{
		OuterColor: color.RGBA{156, 123, 41, 255},  // light brown
		InnerColor: color.RGBA{206, 189, 148, 255}, // light beige
	}
	groupColorMap[70] = GroupColor{
		OuterColor: color.RGBA{148, 165, 148, 255}, // light gray
		InnerColor: color.RGBA{247, 247, 247, 255}, // white
	}
	groupColorMap[71] = GroupColor{
		OuterColor: color.RGBA{148, 165, 148, 255}, // light gray
		InnerColor: color.RGBA{222, 206, 173, 255}, // beige
	}
	groupColorMap[72] = GroupColor{
		// yellow outline
		OuterColor: color.RGBA{148, 165, 148, 255}, // light gray
		// original color is black, but to distinguish with index 73, use yellow outline color
		InnerColor: color.RGBA{247, 239, 165, 255},
		// InnerColor: color.RGBA{8, 16, 8, 255},  // black
	}
	groupColorMap[73] = GroupColor{
		// white outline
		OuterColor: color.RGBA{148, 165, 148, 255}, // light gray
		InnerColor: color.RGBA{8, 16, 8, 255},      // black
	}
	groupColorMap[74] = GroupColor{
		// white outline
		OuterColor: color.RGBA{148, 165, 148, 255}, // light gray
		InnerColor: color.RGBA{198, 231, 231, 255}, // light blue
	}
	groupColorMap[77] = GroupColor{
		OuterColor: color.RGBA{107, 181, 90, 255},  // green
		InnerColor: color.RGBA{231, 231, 123, 255}, // yellow
	}
	groupColorMap[90] = GroupColor{
		OuterColor: color.RGBA{198, 24, 24, 255}, // red
		InnerColor: color.RGBA{156, 16, 16, 255}, // darker red
	}
	groupColorMap[91] = GroupColor{
		OuterColor: color.RGBA{198, 24, 24, 255},   // red
		InnerColor: color.RGBA{231, 231, 231, 255}, // gray
	}
	groupColorMap[92] = GroupColor{
		OuterColor: color.RGBA{198, 24, 24, 255},  // red
		InnerColor: color.RGBA{165, 132, 49, 255}, // brown
	}
	groupColorMap[93] = GroupColor{
		OuterColor: color.RGBA{198, 24, 24, 255}, // red
		InnerColor: color.RGBA{222, 0, 41, 255},  // lighter red
	}
	groupColorMap[94] = GroupColor{
		OuterColor: color.RGBA{198, 24, 24, 255}, // red
		InnerColor: color.RGBA{16, 16, 16, 255},  // black
	}
	groupColorMap[95] = GroupColor{
		OuterColor: color.RGBA{198, 181, 132, 255}, // beige
		InnerColor: color.RGBA{148, 198, 140, 255}, // green
	}
	groupColorMap[96] = GroupColor{
		OuterColor: color.RGBA{198, 181, 132, 255}, // beige
		InnerColor: color.RGBA{30, 128, 200, 255},  // blue
	}
	groupColorMap[97] = GroupColor{
		OuterColor: color.RGBA{198, 181, 132, 255}, // beige
		InnerColor: color.RGBA{198, 222, 239, 255}, // light blue
	}
	groupColorMap[98] = GroupColor{
		OuterColor: color.RGBA{198, 181, 132, 255}, // beige
		InnerColor: color.RGBA{231, 222, 148, 255}, // yellow
	}
	groupColorMap[99] = GroupColor{
		OuterColor: color.RGBA{198, 181, 132, 255}, // tan
		InnerColor: color.RGBA{239, 222, 165, 255}, // light tan
	}
	groupColorMap[100] = GroupColor{
		OuterColor: color.RGBA{0, 0, 0, 255}, // black
		InnerColor: color.RGBA{0, 0, 0, 255}, // black
	}
	groupColorMap[101] = GroupColor{
		OuterColor: color.RGBA{0, 0, 0, 255},    // black
		InnerColor: color.RGBA{222, 0, 41, 255}, // red
	}
	groupColorMap[102] = GroupColor{
		OuterColor: color.RGBA{0, 0, 0, 255},     // black
		InnerColor: color.RGBA{107, 90, 74, 255}, // brown
	}
	groupColorMap[103] = GroupColor{
		OuterColor: color.RGBA{0, 0, 0, 255},     // black
		InnerColor: color.RGBA{90, 107, 90, 255}, // green
	}
	groupColorMap[104] = GroupColor{
		OuterColor: color.RGBA{0, 0, 0, 255},      // black
		InnerColor: color.RGBA{16, 107, 148, 255}, // blue
	}
	groupColorMap[105] = GroupColor{
		OuterColor: color.RGBA{115, 115, 115, 255}, // gray
		InnerColor: color.RGBA{148, 165, 148, 255}, // green
	}
	groupColorMap[106] = GroupColor{
		OuterColor: color.RGBA{115, 115, 115, 255}, // gray
		InnerColor: color.RGBA{24, 24, 24, 255},    // black
	}
	groupColorMap[107] = GroupColor{
		OuterColor: color.RGBA{115, 115, 115, 255}, // gray
		InnerColor: color.RGBA{148, 181, 66, 255},  // lighter green
	}
	groupColorMap[108] = GroupColor{
		OuterColor: color.RGBA{115, 115, 115, 255}, // gray
		InnerColor: color.RGBA{247, 247, 239, 255}, // white
	}
	groupColorMap[109] = GroupColor{
		OuterColor: color.RGBA{115, 115, 115, 255}, // gray
		InnerColor: color.RGBA{231, 189, 123, 255}, // tan
	}
	return groupColorMap
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
		fmt.Println("Group", teamId, ", color:", groupColorMap[teamId], ", unit count:", len(unitList), ", units:", unitList)
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

func drawMap(mapData *fileio.TOAWMapData, outputFilename string) {
	maxImageWidth, maxImageHeight := getImagePosition(mapData.MapHeight, mapData.MapWidth)
	dc := gg.NewContext(int(maxImageWidth), int(maxImageHeight))
	fmt.Println("Map height: ", mapData.MapHeight, ", width: ", mapData.MapWidth)

	drawTiles(dc, mapData.AllTileData, int(mapData.MapHeight), int(mapData.MapWidth))
	drawRiversAndRoads(dc, mapData.AllTileData, int(mapData.MapHeight), int(mapData.MapWidth))
	drawUnits(dc, mapData)
	drawLocations(dc, mapData)

	dc.SavePNG(outputFilename)
	fmt.Println("Saved image to", outputFilename)
}

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

	drawMap(mapData, outputFilename)
}
