package ecs

type ComponentID uint32

var nextComponentID ComponentID
var componentIDs = map[any]ComponentID{}

type typeKey[T any] struct{}

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
