package ecs

// Registry holds one pool per component type, keyed by ComponentID.
// This is the type-erasure boundary in practice: Registry itself never
// mentions Position, Velocity, or any specific component -- adding a
// new component type later requires zero changes to this file.
type Registry struct {
	pools      map[ComponentID]pool
	groupOwner map[ComponentID]*Group
	groups     []*Group
}

func NewRegistry() *Registry {
	return &Registry{
		pools:      map[ComponentID]pool{},
		groupOwner: map[ComponentID]*Group{},
		groups:     []*Group{}, // a list of all registered groups
	}
}

// RemoveEntity erases an entity from every pool in the registry, with
// no per-type code. This is the clearest payoff of the pool interface:
// the loop body doesn't know or care what T each pool holds.
func (r *Registry) RemoveEntity(e EntityID) {
	for _, g := range r.groups {
		g.onRemove(e)
	}
	for _, p := range r.pools {
		p.remove(e)
	}
}

// RegisterPool creates a new typedPool[T] and adds it to the registry,
// keyed by T's ComponentID. -- NOT required before using T. Pool[T]
// (and everything built on it, like AddComponent/GetComponent) lazily
// creates the pool on first use, mirroring EnTT's assure<T>(). Call
// this only if you want a pool to exist ahead of time, e.g. to
// pre-reserve capacity before it's touched.
func RegisterPool[T any](r *Registry) *typedPool[T] {
	pool := NewTypedPool[T]()
	r.pools[ComponentIDOf[T]()] = pool
	return pool
}

func RegisterGroup2[A, B any](r *Registry) *Group {
	idA := ComponentIDOf[A]()
	idB := ComponentIDOf[B]()

	if _, claimed := r.groupOwner[idA]; claimed {
		panic("ecs: component pool already owned by another group")
	}

	if _, claimed := r.groupOwner[idB]; claimed {
		panic("ecs: component pool already owned by another group")
	}

	g := NewGroup(GetPool[A](r), GetPool[B](r))

	r.groupOwner[idA] = g
	r.groupOwner[idB] = g
	r.groups = append(r.groups, g)
	return g
}

func RegisterGroup3[A, B, C any](r *Registry) *Group {
	idA := ComponentIDOf[A]()
	idB := ComponentIDOf[B]()
	idC := ComponentIDOf[C]()

	if _, claimed := r.groupOwner[idA]; claimed {
		panic("ecs: component pool already owned by another group")
	}

	if _, claimed := r.groupOwner[idB]; claimed {
		panic("ecs: component pool already owned by another group")
	}

	if _, claimed := r.groupOwner[idC]; claimed {
		panic("ecs: component pool already owned by another group")
	}

	g := NewGroup(GetPool[A](r), GetPool[B](r), GetPool[C](r))

	r.groupOwner[idA] = g
	r.groupOwner[idB] = g
	r.groupOwner[idC] = g
	r.groups = append(r.groups, g)
	return g
}

// GetPool is the one place the erasure boundary gets crossed back into
// typed code: it looks up the erased `pool` value for T's ComponentID
// and type-asserts it back to *typedPool[T] -- a runtime tag check, not
// reflect. Because it auto-creates a missing pool rather than returning
// nil, every function below it never needs to special-case "not yet
// registered" as distinct from "exists but empty."
//
// This can't be a method on Registry -- Go methods can't carry their
// own type parameters, only the receiver's. T has to live on a free
// function instead.
func GetPool[T any](r *Registry) *typedPool[T] {
	id := ComponentIDOf[T]()
	if p, exists := r.pools[id]; exists {
		return p.(*typedPool[T])
	}
	pool := NewTypedPool[T]()
	r.pools[id] = pool
	return pool
}

// AddComponent / GetComponent / HasComponent / RemoveComponent are the
// generic replacements for what used to be hand-written per-type files
// (e.g. the old position.go/velocity.go, each with their own
// Add/Set/Get/Remove). One definition here now covers every component
// type that exists or ever will.
func AddComponent[T any](r *Registry, e EntityID, c T) {
	GetPool[T](r).set(e, c)
	if g, exists := r.groupOwner[ComponentIDOf[T]()]; exists {
		g.onAdd(e)
	}
}

func GetComponent[T any](r *Registry, e EntityID) (*T, bool) {
	return GetPool[T](r).get(e)
}

func HasComponent[T any](r *Registry, e EntityID) bool {
	return GetPool[T](r).has(e)
}

func RemoveComponent[T any](r *Registry, e EntityID) {
	if g, exists := r.groupOwner[ComponentIDOf[T]()]; exists {
		g.onRemove(e)
	}
	GetPool[T](r).remove(e)
}
