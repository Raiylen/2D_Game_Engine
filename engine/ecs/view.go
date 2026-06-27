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

func (v View1[A]) Len() int {
	return len(v.poolA.dense)
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
			Bidx := v.poolB.sparse[e.Index()]
			f(e, &v.poolA.dense[i], &v.poolB.dense[Bidx])
		}
	} else {
		for i, e := range v.poolB.denseEntities {
			if !v.poolA.has(e) {
				continue
			}
			Aidx := v.poolA.sparse[e.Index()]
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
			Bidx := v.poolB.sparse[e.Index()]
			Cidx := v.poolC.sparse[e.Index()]
			f(e, &v.poolA.dense[i], &v.poolB.dense[Bidx], &v.poolC.dense[Cidx])
		}
	} else if bSize <= aSize && bSize <= cSize {
		for i, e := range v.poolB.denseEntities {
			if !v.poolA.has(e) || !v.poolC.has(e) {
				continue
			}
			Aidx := v.poolA.sparse[e.Index()]
			Cidx := v.poolC.sparse[e.Index()]
			f(e, &v.poolA.dense[Aidx], &v.poolB.dense[i], &v.poolC.dense[Cidx])
		}
	} else {
		for i, e := range v.poolC.denseEntities {
			if !v.poolA.has(e) || !v.poolB.has(e) {
				continue
			}
			Aidx := v.poolA.sparse[e.Index()]
			Bidx := v.poolB.sparse[e.Index()]
			f(e, &v.poolA.dense[Aidx], &v.poolB.dense[Bidx], &v.poolC.dense[i])
		}
	}
}

type View4[A, B, C, D any] struct {
	poolA *typedPool[A]
	poolB *typedPool[B]
	poolC *typedPool[C]
	poolD *typedPool[D]
}

func NewView4[A, B, C, D any](w *World) View4[A, B, C, D] {
	return View4[A, B, C, D]{
		poolA: GetPool[A](w.Registry),
		poolB: GetPool[B](w.Registry),
		poolC: GetPool[C](w.Registry),
		poolD: GetPool[D](w.Registry),
	}
}

func (v View4[A, B, C, D]) Each(f func(e EntityID, a *A, b *B, c *C, d *D)) {
	aSize := len(v.poolA.denseEntities)
	bSize := len(v.poolB.denseEntities)
	cSize := len(v.poolC.denseEntities)
	dSize := len(v.poolD.denseEntities)

	if aSize <= bSize && aSize <= cSize && aSize <= dSize {
		for i, e := range v.poolA.denseEntities {
			if !v.poolB.has(e) || !v.poolC.has(e) || !v.poolD.has(e) {
				continue
			}
			Bidx := v.poolB.sparse[e.Index()]
			Cidx := v.poolC.sparse[e.Index()]
			Didx := v.poolD.sparse[e.Index()]
			f(e, &v.poolA.dense[i], &v.poolB.dense[Bidx], &v.poolC.dense[Cidx], &v.poolD.dense[Didx])
		}
	} else if bSize <= aSize && bSize <= cSize && bSize <= dSize {
		for i, e := range v.poolB.denseEntities {
			if !v.poolA.has(e) || !v.poolC.has(e) || !v.poolD.has(e) {
				continue
			}
			Aidx := v.poolA.sparse[e.Index()]
			Cidx := v.poolC.sparse[e.Index()]
			Didx := v.poolD.sparse[e.Index()]
			f(e, &v.poolA.dense[Aidx], &v.poolB.dense[i], &v.poolC.dense[Cidx], &v.poolD.dense[Didx])
		}
	} else if cSize <= dSize && cSize <= aSize && cSize <= bSize {
		for i, e := range v.poolC.denseEntities {
			if !v.poolA.has(e) || !v.poolB.has(e) || !v.poolD.has(e) {
				continue
			}
			Aidx := v.poolA.sparse[e.Index()]
			Bidx := v.poolB.sparse[e.Index()]
			Didx := v.poolD.sparse[e.Index()]
			f(e, &v.poolA.dense[Aidx], &v.poolB.dense[Bidx], &v.poolC.dense[i], &v.poolD.dense[Didx])
		}
	} else {
		for i, e := range v.poolD.denseEntities {
			if !v.poolA.has(e) || !v.poolB.has(e) || !v.poolC.has(e) {
				continue
			}
			Aidx := v.poolA.sparse[e.Index()]
			Bidx := v.poolB.sparse[e.Index()]
			Cidx := v.poolC.sparse[e.Index()]
			f(e, &v.poolA.dense[Aidx], &v.poolB.dense[Bidx], &v.poolC.dense[Cidx], &v.poolD.dense[i])
		}
	}
}
