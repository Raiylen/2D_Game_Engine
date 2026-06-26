package events

import (
	"github.com/raiylen/2d_game_engine/engine/ecs"
)

const (
	Collision = "collision"
	KbControl = "kbcontrol"
)

type CollisionEvent struct {
	A ecs.EntityID
	B ecs.EntityID
}

type KbControlEvent struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}
