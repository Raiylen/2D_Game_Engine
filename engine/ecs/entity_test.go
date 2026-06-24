package ecs

import (
	"github.com/raiylen/2d_game_engine/logger"
	"testing"
)

func TestEntityRecycling(t *testing.T) {
	w := NewWorld()
	w.SetLogger(logger.NewLogger())

	a := w.NewEntity()
	w.DestroyEntity(a)
	w.Update(0)

	b := w.NewEntity()

	if b.Index() != a.Index() {
		t.Errorf("expected index %d, got %d", a.Index(), b.Index())
	}
	if b.Gen() != a.Gen()+1 {
		t.Errorf("expected gen %d, got %d", a.Gen()+1, b.Gen())
	}
	if w.IsAlive(a) {
		t.Error("stale entity should not be alive")
	}
	if !w.IsAlive(b) {
		t.Error("recycled entity should be alive")
	}
}
