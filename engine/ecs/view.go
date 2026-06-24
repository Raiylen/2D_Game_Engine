package ecs

type View1[A any] struct {
	poolA *typedPool[A]
}

func NewView1[A any](w *World) View1[A] {
	return View1[A]{
		poolA: GetPool[A](w.Registry),
	}
}

func (v View1[A]) Each(f func(e EntityID, a *A)) {
	for i, e := range v.poolA.denseEntities {
		f(e, &v.poolA.dense[i])
	}
}

type View2[A, B any] struct {
	poolA *typedPool[A]
	poolB *typedPool[B]
}

func NewView2[A, B any](w *World) View2[A, B] {
	return View2[A, B]{
		poolA: GetPool[A](w.Registry),
		poolB: GetPool[B](w.Registry),
	}
}

func (v View2[A, B]) Each(f func(e EntityID, a *A, b *B)) {
	if len(v.poolA.denseEntities) <= len(v.poolB.denseEntities) {
		for i, e := range v.poolA.denseEntities {
			if !v.poolB.has(e) {
				continue
			}
			Bidx := v.poolB.sparse[e]
			f(e, &v.poolA.dense[i], &v.poolB.dense[Bidx])
		}
	} else {
		for i, e := range v.poolB.denseEntities {
			if !v.poolA.has(e) {
				continue
			}
			Aidx := v.poolA.sparse[e]
			f(e, &v.poolA.dense[Aidx], &v.poolB.dense[i])
		}
	}
}

type View3[A, B, C any] struct {
	poolA *typedPool[A]
	poolB *typedPool[B]
	poolC *typedPool[C]
}

func NewView3[A, B, C any](w *World) View3[A, B, C] {
	return View3[A, B, C]{
		poolA: GetPool[A](w.Registry),
		poolB: GetPool[B](w.Registry),
		poolC: GetPool[C](w.Registry),
	}
}

func (v View3[A, B, C]) Each(f func(e EntityID, a *A, b *B, c *C)) {
	aSize := len(v.poolA.denseEntities)
	bSize := len(v.poolB.denseEntities)
	cSize := len(v.poolC.denseEntities)

	if aSize <= bSize && aSize <= cSize {
		for i, e := range v.poolA.denseEntities {
			if !v.poolB.has(e) || !v.poolC.has(e) {
				continue
			}
			Bidx := v.poolB.sparse[e]
			Cidx := v.poolC.sparse[e]
			f(e, &v.poolA.dense[i], &v.poolB.dense[Bidx], &v.poolC.dense[Cidx])
		}
	} else if bSize <= aSize && bSize <= cSize {
		for i, e := range v.poolB.denseEntities {
			if !v.poolA.has(e) || !v.poolC.has(e) {
				continue
			}
			Aidx := v.poolA.sparse[e]
			Cidx := v.poolC.sparse[e]
			f(e, &v.poolA.dense[Aidx], &v.poolB.dense[i], &v.poolC.dense[Cidx])
		}
	} else {
		for i, e := range v.poolC.denseEntities {
			if !v.poolA.has(e) || !v.poolB.has(e) {
				continue
			}
			Aidx := v.poolA.sparse[e]
			Bidx := v.poolB.sparse[e]
			f(e, &v.poolA.dense[Aidx], &v.poolB.dense[Bidx], &v.poolC.dense[i])
		}
	}
}
