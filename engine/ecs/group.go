package ecs

// Group maintains a slice of component pools that should maintain
// entity alignment for use by systems for faster iteration. count is
// the number of entities currently belonging to the group -- meaning
// they have an entry in every pool the group owns -- and also serves
// as the partition boundary within each owned pool's dense array:
// indices [0, count) hold exactly those entities, at matching indices
// across every pool in the group.
type Group struct {
	pools []pool
	count int
}

func NewGroup(pools ...pool) *Group {
	g := &Group{pools: pools}
	g.refresh()
	return g
}

func (g *Group) qualifies(e EntityID) bool {
	for _, p := range g.pools {
		if !p.has(e) {
			return false
		}
	}
	return true
}

// admit swaps e into the boundary slot (index count) for every pool
// in the group, then advances count. This assumes qualifies(e) has
// already returned true -- admit does NOT check membership itself,
// and will index out of bounds if called on an entity missing from
// any owned pool.
func (g *Group) admit(e EntityID) {
	for _, p := range g.pools {
		if idx := p.index(e); idx != g.count {
			p.swap(idx, g.count)
		}
	}
	g.count++
}

// refresh walks the base pool's dense array starting from the current
// count boundary, and for each entity found that qualifies for the
// group (has an entry in every owned pool), admits it -- swapping it
// into the count position across every pool in the group, not just
// base.
func (g *Group) refresh() {
	if len(g.pools) == 0 {
		return
	}
	base := g.pools[0]
	for i := g.count; i < base.size(); i++ {
		e := base.at(i)
		if g.qualifies(e) {
			g.admit(e)
		}
	}
}

func (g *Group) evict(e EntityID) {
	last := g.count - 1
	for _, p := range g.pools {
		if idx := p.index(e); idx != last {
			p.swap(idx, last)
		}
	}
	g.count--
}

func (g *Group) isMember(e EntityID) bool {
	idx := g.pools[0].index(e)
	return idx != -1 && idx < g.count
}

func (g *Group) onAdd(e EntityID) {
	if g.qualifies(e) && !g.isMember(e) {
		g.admit(e)
	}
}

// onRemove must be called BEFORE an entity is removed
func (g *Group) onRemove(e EntityID) {
	if g.isMember(e) {
		g.evict(e)
	}
}

type GroupView2[A, B any] struct {
	group *Group
	poolA *typedPool[A]
	poolB *typedPool[B]
}

func NewGroupView2[A, B any](w *World) GroupView2[A, B] {
	idA := ComponentIDOf[A]()
	idB := ComponentIDOf[B]()

	gA, existsA := w.Registry.groupOwner[idA]
	gB, existsB := w.Registry.groupOwner[idB]

	if !existsA || !existsB || gA != gB {
		panic("ecs: no group owns this component pair -- use NewView2 instead")
	}

	return GroupView2[A, B]{
		group: gA,
		poolA: GetPool[A](w.Registry),
		poolB: GetPool[B](w.Registry),
	}

}

func (v GroupView2[A, B]) Each(f func(e EntityID, a *A, b *B)) {
	for i := 0; i < v.group.count; i++ {
		f(v.group.pools[0].at(i), &v.poolA.dense[i], &v.poolB.dense[i])
	}
}

type GroupView3[A, B, C any] struct {
	group *Group
	poolA *typedPool[A]
	poolB *typedPool[B]
	poolC *typedPool[C]
}

func NewGroupView3[A, B, C any](w *World) GroupView3[A, B, C] {
	idA := ComponentIDOf[A]()
	idB := ComponentIDOf[B]()
	idC := ComponentIDOf[C]()

	gA, existsA := w.Registry.groupOwner[idA]
	gB, existsB := w.Registry.groupOwner[idB]
	gC, existsC := w.Registry.groupOwner[idC]

	if !existsA || !existsB || !existsC || gA != gB || gA != gC {
		panic("ecs: no group owns this component triple -- use NewView3 instead")

	}

	return GroupView3[A, B, C]{
		group: gA,
		poolA: GetPool[A](w.Registry),
		poolB: GetPool[B](w.Registry),
		poolC: GetPool[C](w.Registry),
	}
}

func (v GroupView3[A, B, C]) Each(f func(e EntityID, a *A, b *B, c *C)) {
	for i := 0; i < v.group.count; i++ {
		f(v.group.pools[0].at(i), &v.poolA.dense[i], &v.poolB.dense[i], &v.poolC.dense[i])
	}
}
