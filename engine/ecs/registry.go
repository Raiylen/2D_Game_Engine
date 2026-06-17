package ecs

type Registry struct {
	pools map[ComponentID]pool
}

func NewRegistry() *Registry {
	return &Registry{pools: map[ComponentID]pool{}}
}

func (r *Registry) RemoveEntity(e EntityID) {
	for _, p := range r.pools {
		p.remove(e)
	}
}

// RegisterPool creates a new typedPool[T] and adds it to the registry,
// keyed by T's ComponentID.
//
// NOT required before using T -- Pool[T] (and everything built on it,
// like AddComponent/GetComponent) lazily creates the pool on first use,
// mirroring EnTT's assure<T>(). Call this only if you want a pool to
// exist ahead of time, e.g. to pre-reserve capacity before it's touched.
func RegisterPool[T any](r *Registry) *typedPool[T] {
	pool := NewTypedPool[T]()
	r.pools[ComponentIDOf[T]()] = pool
	return pool
}

// Pool retrieves the typed pool for T, deriving the ComponentID from T
// itself. If no pool exists yet for T, one is created and stored.
func GetPool[T any](r *Registry) *typedPool[T] {
	id := ComponentIDOf[T]()
	if p, exists := r.pools[id]; exists {
		return p.(*typedPool[T])
	}
	pool := NewTypedPool[T]()
	r.pools[id] = pool
	return pool
}

func AddComponent[T any](r *Registry, e EntityID, c T) {
	GetPool[T](r).set(e, c)
}

func GetComponent[T any](r *Registry, e EntityID) (*T, bool) {
	return GetPool[T](r).get(e)
}

func HasComponent[T any](r *Registry, e EntityID) bool {
	return GetPool[T](r).has(e)
}

func RemoveComponent[T any](r *Registry, e EntityID) {
	GetPool[T](r).remove(e)
}
