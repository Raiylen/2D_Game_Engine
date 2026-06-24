package ecs

// pool is the type-erased view of a typedPool[T]. It exists for one
// reason: Go maps require a single fixed value type, so
// map[ComponentID]*typedPool[Position] could never also hold a
// *typedPool[Velocity] -- they're unrelated concrete types. pool is the
// one type both can satisfy, which is what lets Registry hold pools for
// every component type in a single map.
//
// Deliberately minimal: only operations that are useful WITHOUT knowing
// T live here (has/remove). Anything that needs T (get, set, reading
// dense) stays on typedPool[T] itself and is reached through the
// generic GetPool[T] function instead.
type pool interface {
	has(e EntityID) bool
	remove(e EntityID)
	swap(i, j int)
	at(i int) EntityID
	index(e EntityID) int
	size() int
}

// typedPool[T] is a sparse set: a flat, contiguous "dense" slice
// holding the actual component data (cache-friendly, directly
// iterable), a parallel slice mapping dense index -> EntityID, and a
// "sparse" slice mapping EntityID -> dense index (or -1 if absent).
// This is the same structure EnTT's storage<T> uses internally.
type typedPool[T any] struct {
	dense         []T        // the actual component values, tightly packed
	denseEntities []EntityID // mirrors dense, but each index is the entityID
	sparse        []int      // sparse[entityID] -> index into dense, or -1
}

func NewTypedPool[T any]() *typedPool[T] {
	return &typedPool[T]{
		dense:         []T{},
		denseEntities: []EntityID{},
		sparse:        []int{},
	}
}

// grow expands sparse so index e is valid, initializing new slots to -1
// (meaning "no component").
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

// set inserts a new component or overwrites an existing one -- callers
// (AddComponent, SetComponent) don't need to distinguish the two cases.
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

// get returns a pointer directly into the dense slice -- no copy, no
// boxing. This is the payoff of keeping typedPool generic: callers get
// back a real *T they can mutate in place.
func (s *typedPool[T]) get(e EntityID) (*T, bool) {
	if !s.has(e) {
		return nil, false
	}
	idx := s.sparse[e]
	return &s.dense[idx], true
}

// remove is a classic sparse-set removal: swap the removed element with
// the last element in dense (so dense stays contiguous with no holes),
// fix up the sparse entry for whichever entity got moved, then shrink.
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

// swap exchanges the dense-array contents at indices i and j, keeping
// dense, denseEntities, and sparse consistent. This is the primitive
// group bookkeeping used to maintain its owned/unowned partition
func (s *typedPool[T]) swap(i, j int) {
	s.dense[i], s.dense[j] = s.dense[j], s.dense[i]
	s.denseEntities[i], s.denseEntities[j] = s.denseEntities[j], s.denseEntities[i]
	s.sparse[s.denseEntities[i]] = i
	s.sparse[s.denseEntities[j]] = j
}

// has checks a pool's dense slice for an entity e
func (s *typedPool[T]) has(e EntityID) bool {
	return int(e) < len(s.sparse) && s.sparse[e] != -1
}

// at returns the entity stored at a pool's dense position i
func (s *typedPool[T]) at(i int) EntityID {
	return s.denseEntities[i]
}

// index returns a pool's e current dense position or -1
func (s *typedPool[T]) index(e EntityID) int {
	if !s.has(e) {
		return -1
	}
	return s.sparse[e]
}

// size returns how many entities a pool dense slice currently holds
func (s *typedPool[T]) size() int {
	return len(s.dense)
}
