package ecs

// ComponentID is a small, dense integer used to key pools in the
// Registry's map. It exists so the registry can store many different
// component types in one map without needing reflect.TypeOf -- see
// typeKey below for how that's achieved.
type ComponentID uint32

var nextComponentID ComponentID

// componentIDs maps a per-type marker to the ID assigned for that type.
// Keying by typeKey[T]{} (not by T's name as a string, and not via
// reflect) is what guarantees ComponentIDOf[Position]() always returns
// the same value no matter where or how many times it's called.
var componentIDs = map[any]ComponentID{}

type typeKey[T any] struct{}

// ComponentIDOf returns the stable ID for T, assigning a new one the
// first time T is seen and returning the cached one on every call after
// that. Safe to call repeatedly and from anywhere -- callers never need
// to track IDs themselves (contrast with an earlier version of this
// function that just incremented a counter with no way to verify a
// type hadn't already been assigned one).
func ComponentIDOf[T any]() ComponentID {
	key := typeKey[T]{}
	if id, exists := componentIDs[key]; exists {
		return id
	}
	id := nextComponentID
	nextComponentID++
	componentIDs[key] = id
	return id
}
