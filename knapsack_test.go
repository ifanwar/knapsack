package knapsack

import (
	"context"
	"testing"
)

func TestGreedy(t *testing.T) {
	ctx := context.Background()
	data := map[string]int64{"a": 10, "b": 40, "c": 30, "d": 50}
	res := Greedy(ctx, data, 60)
	t.Log("Greedy result:", res)
}

func TestStandard(t *testing.T) {
	ctx := context.Background()
	data := map[string]int64{"a": 10, "b": 40, "c": 30, "d": 50}
	res := Standard(ctx, data, 60)
	t.Log("Standard result:", res)
}
