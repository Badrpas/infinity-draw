package repo

import (
	"github.com/badrpas/infinity-draw/server/pkg/config"
	"math/rand"
)

type Chunk struct {
	x, y, z string
	Data    []byte
}

// Save to DB
func (c *Chunk) Save() error {
	return nil
}

type ChunkRepo struct {
	ActiveChunks []*Chunk
	freeSpot     int
}

func (r *ChunkRepo) GetChunk(x, y, z string) *Chunk {
	for _, chunk := range r.ActiveChunks {
		if chunk == nil {
			continue
		}
		if chunk.x == x && chunk.y == y && chunk.z == z {
			return chunk
		}
	}

	size := config.GlobalConfig.ChunkSize
	randomChunk := &Chunk{
		Data: make([]byte, size*size),
	}

	for y := 0; y < size; y++ {
		if rand.Intn(100) > 90 {
			continue
		}
		max := rand.Intn(size/2) + size/2
		for x := rand.Intn(size / 2); x < max; x++ {
			randomChunk.Data[y*size+x] = 255
		}
	}

	r.ActiveChunks[r.freeSpot] = randomChunk
	r.updateFreeSpot()

	return randomChunk
}

func (r *ChunkRepo) updateFreeSpot() {
	for idx, chunk := range r.ActiveChunks {
		if chunk == nil {
			r.freeSpot = idx
		}
	}
}

func (r *ChunkRepo) Offload(chunk *Chunk) error {
	for idx, c := range r.ActiveChunks {
		if c == nil {
			continue
		}
		if chunk.x == c.x && chunk.y == c.y && chunk.z == c.z {
			r.ActiveChunks[idx] = nil
			r.updateFreeSpot()
		}
	}

	return chunk.Save()
}
