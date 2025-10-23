# TOAW File Format Documentation

This document describes the file format used by The Operational Art of War (TOAW) scenario files (.sce).

## TOAW Game Version Notes

The Operational Art of War IV (TOAW4) uses gzip to compress the files. The file will begin with the gzip header, where the magic number is 0x1f8b and the compression method is 08 for DEFLATE.

TOAW3 and earlier games use PKWare Compression Library to compress each of the blocks. The header will begin with "TOAC".

## Map Header

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

## Tile Data

In the decompressed blocks array, this is the block with index 1, which would be the 2nd element starting from index 0.

In TOAW4, the maximum map size is 700x700 and each tile is 48 bytes. In TOAW3 and earlier games, the maximum map size is either 300x300 or 100x100, and each tile is 47 bytes.

The tile data is stored as a byte array where each tile contains terrain and feature information.

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

## Location Data

Location data contains information about named locations on the map (cities, towns, etc.).

| Type | Size | Description |
| ---- | ---- | ----------- |
| int32 | 4 bytes | X coordinate |
| int32 | 4 bytes | Y coordinate |
| byte[28] | 28 bytes | Location name |

## Unit Data

Unit data contains information about military units on the map. Each unit record is 392 bytes.

| Type | Size | Description |
| ---- | ---- | ----------- |
| byte[20] | 20 bytes | Unit name |
| uint32[60] | 240 bytes | Unknown block 1 |
| byte[48] | 48 bytes | Unknown block 2 |
| byte[4] | 4 bytes | Unknown data at offset 0x134 |
| uint32 | 4 bytes | Unit color and type (bit-packed) |
| uint32 | 4 bytes | Unknown data at offset 0x13c |
| uint32 | 4 bytes | Unknown data at offset 0x140 |
| uint32 | 4 bytes | Proficiency |
| uint32 | 4 bytes | Readiness |
| uint32 | 4 bytes | Supply level |
| uint32 | 4 bytes | Unknown data at offset 0x150 |
| uint32 | 4 bytes | Other unit index on same tile |
| int32 | 4 bytes | X coordinate |
| int32 | 4 bytes | Y coordinate |
| uint32 | 4 bytes | Unknown data at offset 0x160 |
| uint32 | 4 bytes | Unknown data at offset 0x164 |
| uint32 | 4 bytes | Unknown data at offset 0x168 |
| uint32 | 4 bytes | Unknown data at offset 0x16c |
| uint32 | 4 bytes | Unknown data at offset 0x170 |
| uint32 | 4 bytes | Unknown data at offset 0x174 |
| uint32 | 4 bytes | Unit index |
| byte[12] | 12 bytes | Unknown block 4 |

## Team Name Data

Team name data contains information about the two teams/players in the scenario.

| Type | Size | Description |
| ---- | ---- | ----------- |
| byte[17] | 17 bytes | Country name |
| byte[35] | 35 bytes | Force name |
| uint32 | 4 bytes | Proficiency |
| uint32 | 4 bytes | Supply level |
| uint32 | 4 bytes | Country flag ID |

## Data Block Structure

The scenario file contains multiple compressed data blocks:

| Block Index | Description |
| ----------- | ----------- |
| 0 | Unknown |
| 1 | Tile data (map terrain) |
| 2 | Unit data |
| 3 | Unknown |
| 4 | Team name data |
| 5-9 | Unknown |
| 10 | Location data (TOAW3 and earlier) |
| 11 | Location data (TOAW4) |

## Map Data Structure

The complete map data structure contains:

| Field | Type | Description |
| ----- | ---- | ----------- |
| Version | int | Game version |
| AllLocationData | []LocationData | Array of location data |
| AllTeamNameData | []*TeamNameData | Array of team data |
| AllTileData | [][]*TileData | 2D array of tile data |
| AllUnitData | []*UnitData | Array of unit data |
| MapWidth | int | Map width in tiles |
| MapHeight | int | Map height in tiles |
