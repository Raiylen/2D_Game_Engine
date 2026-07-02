package events

import (
	"github.com/raiylen/2d_game_engine/engine/ecs"
)

const (
	Collision     = "collision"
	PlayerControl = "playercontrol"
	KbFire        = "kbfire"
)

type CollisionEvent struct {
	A ecs.EntityID
	B ecs.EntityID
}

type PlayerControlEvent struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

type KbFireEvent struct {
}
