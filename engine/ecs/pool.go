package ecs

// pool is the type-erased view of a typedPool[T]. The registry holds a
// collection of these without knowing what T is for each one.
type pool interface {
	has(e EntityID) bool
	remove(e EntityID)
}

type typedPool[T any] struct {
	dense         []T
	denseEntities []EntityID
	sparse        []int
}

func NewTypedPool[T any]() *typedPool[T] {
	return &typedPool[T]{
		dense:         []T{},
		denseEntities: []EntityID{},
		sparse:        []int{},
	}
}

func (s *typedPool[T]) grow(e EntityID) {
	if int(e) < len(s.sparse) {
		return
	}
	newLen := int(e) + 1
	newCap := len(s.sparse) * 2
	if newLen > newCap {
		newCap = newLen
	}
	old := s.sparse
	s.sparse = make([]int, newCap)
	for i := range s.sparse {
		s.sparse[i] = -1
	}
	copy(s.sparse, old)
}

func (s *typedPool[T]) has(e EntityID) bool {
	return int(e) < len(s.sparse) && s.sparse[e] != -1
}

func (s *typedPool[T]) set(e EntityID, component T) {
	s.grow(e)
	if idx := s.sparse[e]; idx != -1 {
		s.dense[idx] = component
		return
	}
	s.sparse[e] = len(s.dense)
	s.dense = append(s.dense, component)
	s.denseEntities = append(s.denseEntities, e)
}

func (s *typedPool[T]) get(e EntityID) (*T, bool) {
	if !s.has(e) {
		return nil, false
	}
	idx := s.sparse[e]
	return &s.dense[idx], true
}

func (s *typedPool[T]) remove(e EntityID) {
	if !s.has(e) {
		return
	}
	idx := s.sparse[e]

	lastIdx := len(s.dense) - 1
	lastEntity := s.denseEntities[lastIdx]

	s.dense[idx] = s.dense[lastIdx]
	s.denseEntities[idx] = lastEntity
	s.sparse[lastEntity] = idx

	s.dense = s.dense[:lastIdx]
	s.denseEntities = s.denseEntities[:lastIdx]

	s.sparse[e] = -1
}
