package ecs

// System is the contract every gameplay system implements (see
// movementSystem in the systems package for an example). Kept in its
// own file rather than folded into world.go, even though World is its
// only consumer -- it's a distinct concept, the same reasoning that
// keeps Registry and typedPool in their own files despite World using
// both.
type System interface {
	Update(w *World, dt float64)
}

type Renderer interface {
	Render(w *World)
}

type EventHandler interface {
	RegisterHandlers(w *World)
}
