package config

type cfg struct {
	// Size of usable chunk on screen
	ChunkSize int
	// Has to be power of 2, but not greater than ChunkSize
	ZoomFactor int
}

var GlobalConfig cfg = cfg{
	ChunkSize:  256,
	ZoomFactor: 4,
}
