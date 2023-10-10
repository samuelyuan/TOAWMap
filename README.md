# TOAWMap

This program will read scenario files (.sce) and render the maps.

The following games are supported:

* The Operational Art of War: Century of Warfare
* The Operational Art of War III
* The Operational Art of War IV

### Command-Line Usage

```
./TOAWMap.exe -input=[input filename] -output=[output filename]
```

Example
```
./TOAWMap.exe -input=scenario.sce -output=scenario.png
```

<div style="display:inline-block;">
<img src="https://raw.githubusercontent.com/samuelyuan/TOAWMap/master/screenshots/korea50.png" alt="korea50" width="145" height="300" />
<img src="https://raw.githubusercontent.com/samuelyuan/TOAWMap/master/screenshots/manchuria.png" alt="manchuria" width="300" height="300" />
</div>

### About

There have been many maps generated for The Operational Art of War (TOAW) series, but there isn't much information on the file formats and there is no way to parse the data. This project analyzes the different games and you can see how there is a lot of overlap between the various file formats.

### TOAW Game Version Notes

The Operational Art of War IV (TOAW4) uses gzip to compress the files. The file will begin with the gzip header, where the magic number is 0x1f8band the compression method is 08 for DEFLATE.

TOAW3 and earlier games use PKWare Compression Library to compress each of the blocks. The header will begin with "TOAC".

### Map Header

| Type | Size | Description |
| ---- | ---- | ----------- |
| byte[16] | 16 bytes | Header |
| uint32 | 4 bytes | UnknownInt1 |
| byte[264] | 264 bytes | Map title |
| uint32 | 4 bytes | Version |
| uint32 | 4 bytes | UnknownInt3 |
| byte[8192] | 8192 bytes | Map description |
| byte[8192] | 8192 bytes | End message for Team1, 1st victory message |
| byte[8192] | 8192 bytes | End message for Team1, 2nd victory message |
| byte[8192] | 8192 bytes | End message draw 1 |
| byte[8192] | 8192 bytes | End message for Team2, victory message |
| byte[8192] | 8192 bytes | UnknownBlock1 |
| byte[8192] | 8192 bytes | End message draw 2 |
| byte[8192] | 8192 bytes | UnknownBlock2 |
| uint32 | 4 bytes | Team goes first |
| byte[36] | 36 bytes | TeamNameBlock3 |

### Tile Data

In the decompressed blocks array, this is the block with index 1, which would be the 2nd element starting from index 0.

In TOAW4, the maximum map size is 700x700. In TOAW3 and earlier games, the maximum map size is either 300x300 or 100x100.

Each tile is 47 bytes, which is represented as a byte array in the code.

| Index | Description |
| ----- | ----------- |
| 1-4 | Sand tile flag (index 1 = arid, 2 = sandy, 3 = r_sandy, 4 = badlands) |
| 5 | Hills tile flag |
| 6 | Mountains tile flag |
| 7 | Impassable tile flag |
| 8 | Marsh tile flag |
| 9 | Flooded marsh tile flag |
| 10 | Shallow water tile flag |
| 11 | Deep water tile flag |
| 14-17 | Urban tile flag |
| 22 | River tile flag |
| 23 | Major river tile flag |
| 26-29 | Forest tile flag (index 26 = c_forest, 27 = d_forest, 28 = m_forest, 29 = t_forest) |
| 31 | Road tile flag |
| 33 | Railroad tile flag |
| 38 | Tile type |