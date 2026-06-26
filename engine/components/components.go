package components

import (
// "github.com/veandco/go-sdl2/sdl"
)

type Transform struct {
	X, Y     float64
	Scale    float64
	Rotation float64
}

type RigidBody struct {
	DX, DY float64
}

type Sprite struct {
	Width      int
	Height     int
	SourcePosX int
	SourcePosY int
	AssetID    string
	Layer      int
}

type Animation struct {
	NumFrames    int
	CurrentFrame int
	FrameRate    int
	StartTime    int
	Loop         bool
}

type BoxCollider struct {
	Width   int
	Height  int
	OffsetX int
	OffsetY int
}

type KeyboardControl struct {
	UpVelocity    int
	DownVelocity  int
	RightVelocity int
	LeftVelocity  int
}
