package ecs

type View2[A, B any] struct {
	poolA *typedPool[A]
	poolB *typedPool[B]
}

func NewView2[A, B any](a *typedPool[A], b *typedPool[B]) View2[A, B] {
	return View2[A, B]{
		poolA: a,
		poolB: b,
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

func Query2[A, B any](w *World) View2[A, B] {
	return NewView2(GetPool[A](w.registry), GetPool[B](w.registry))
}
