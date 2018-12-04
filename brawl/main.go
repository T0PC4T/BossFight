package brawl

var (
	screenWidth  int
	screenHeight int
	tilesWidth   int
	tilesHeight  int
	tileSize     int
	tileSizeF    float64
)

// InitializeBrawl sets up the the global variables from configurations done in menus
func InitializeBrawl(screenWidthV, screenHeightV, tilesWidthV, tilesHeightV int) {
	screenWidth = screenWidthV
	screenHeight = screenHeightV
	tilesWidth = tilesWidthV
	tilesHeight = tilesHeightV
	tileSize = screenWidth / tilesWidth
	tileSizeF = float64(tileSize)

}
