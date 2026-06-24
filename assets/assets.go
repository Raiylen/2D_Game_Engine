package assets

import (
	"fmt"
	img "github.com/veandco/go-sdl2/img"
	sdl "github.com/veandco/go-sdl2/sdl"
)

type AssetManager struct {
	renderer *sdl.Renderer
	cache    map[string]*sdl.Texture
}

func NewAssetManager(renderer *sdl.Renderer) *AssetManager {
	return &AssetManager{
		renderer: renderer,
		cache:    map[string]*sdl.Texture{},
	}
}

func (t *AssetManager) AddTexture(id string, path string) error {
	tex, err := img.LoadTexture(t.renderer, path)
	if err != nil {
		return fmt.Errorf("Error loading texture: %v", err)
	}
	t.cache[id] = tex
	return nil
}

func (t *AssetManager) GetTexture(id string) (*sdl.Texture, error) {
	tex, exists := t.cache[id]
	if !exists {
		return nil, fmt.Errorf("Texture for AssetID '%s' not found", id)
	}
	return tex, nil
}

func (t *AssetManager) Close() {
	for _, tex := range t.cache {
		tex.Destroy()
	}
	t.cache = map[string]*sdl.Texture{}
}
