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
	IsFixed    bool
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

type Player struct {
	UpVelocity    float64
	DownVelocity  float64
	RightVelocity float64
	LeftVelocity  float64
	FacingX       float64
	FacingY       float64
}

type Camera struct {
	X, Y      float64
	Width     int
	Height    int
	MapWidth  int
	MapHeight int
}

type Emitter struct {
	DX, DY         float64
	Frequency      float64
	ProjSpeed      float64
	ProjDuration   float64
	ProjPercDamage int
	Timer          float64
	IsFriendly     bool
}

type Projectile struct {
	Duration   float64
	PercDamage int
	IsFriendly bool
	Timer      float64
}

type Health struct {
	HealthPerc int
}
