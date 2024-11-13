package graphics

import (
  	"image/color"
)

type GroupColor struct {
	OuterColor color.RGBA
	InnerColor color.RGBA
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
