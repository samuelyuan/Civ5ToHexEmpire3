package fileio

import (
  "log"
  "os"
)

func ConvertCiv5MapToHE3Map(inputFilename string, outputFilename string) {
  civ5MapData, err := ReadCiv5MapFile(inputFilename)
  if err != nil {
    log.Fatal("Failed to load CIV5 Map: ", err)
  }
  civ5MapHeight := len(civ5MapData.MapTiles)
  civ5MapWidth := len(civ5MapData.MapTiles[0])

  // Hex Empire 3 map dimensions are inverted
  tileMap := make([][]*HE3MapTile, int(civ5MapWidth))
  for i := 0; i < civ5MapWidth; i++ {
    tileMap[i] = make([]*HE3MapTile, int(civ5MapHeight))
  }

  for i := 0; i < civ5MapHeight; i++ {
		for j := 0; j < civ5MapWidth; j++ {
      tile := HE3MapTile{}

      civ5MapTile := civ5MapData.MapTiles[i][j]
      elevation := civ5MapTile.Elevation
			terrainString := civ5MapData.TerrainList[civ5MapTile.TerrainType]
      featureString := ""
      if int(civ5MapTile.FeatureTerrainType) < len(civ5MapData.FeatureTerrainList) {
        featureString = civ5MapData.FeatureTerrainList[civ5MapTile.FeatureTerrainType]
      }

      if terrainString == "TERRAIN_COAST" || terrainString == "TERRAIN_OCEAN" {
        tile.Height = 0.0
        tile.IsSea = true
        tile.IsMountain = false
      } else if elevation == 2 { // mountain
        tile.Height = 1.0
        tile.IsSea = false
        tile.IsMountain = true
      } else {
        tile.Height = 0.2
        tile.IsSea = false
        tile.IsMountain = false
      }

      switch terrainString {
      case "TERRAIN_GRASS":
        tile.TileType = Grass
      case "TERRAIN_PLAINS":
        tile.TileType = Farmland
      case "TERRAIN_DESERT":
        tile.TileType = Sand
      case "TERRAIN_TUNDRA", "TERRAIN_SNOW":
        tile.TileType = Snow // No option for tundra
      case "TERRAIN_COAST", "TERRAIN_OCEAN":
        tile.TileType = Grass // default value, but will still be rendered as water
      default:
        tile.TileType = Grass
      }

      if featureString == "FEATURE_FOREST" || featureString == "FEATURE_JUNGLE" {
        tile.TileType = Forest
      }

      tile.HasRoad = false
      tile.HasFlag = false
      tile.Party = -1
      tile.HasInfantry = false
      tile.Infantry = nil
      tile.HasArtillery = false
      tile.Artillery = nil

      tileMap[j][i] = &tile
    }
  }

  he3MapData := &HE3Map{
    MapTiles:  tileMap,
		MapTitle:  inputFilename,
		MapAuthor: "",
		MapStyle:  MapStyle{},
		Width:     int32(civ5MapWidth),
		Depth:     int32(civ5MapHeight),
  }

  serializedHE3Data := Serialize(he3MapData)
	err = os.WriteFile(outputFilename, []byte(serializedHE3Data), 0644)
	if err != nil {
		log.Fatal("Failed to write to output file", err)
	}
}
