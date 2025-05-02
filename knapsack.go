package knapsack

import (
	"context"
	"sort"
)

type Number interface {
	int64 | float64
}

// Greedy solves knapsack using a greedy strategy: picks the highest value items first
func Greedy[K comparable, V Number](ctx context.Context, dataset map[K]V, target V) []K {
	type kv struct {
		Key   K
		Value V
	}

	var items []kv
	for k, v := range dataset {
		items = append(items, kv{k, v})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})

	var result []K
	var total V
	for _, item := range items {
		if total+item.Value <= target {
			total += item.Value
			result = append(result, item.Key)
		}
	}
	return result
}

// Standard solves 0/1 knapsack problem using dynamic programming
func Standard[K comparable, V Number](ctx context.Context, dataset map[K]V, target V) []K {
	// Convert map to slices for processing
	type item struct {
		Key   K
		Value V
	}
	var items []item
	for k, v := range dataset {
		items = append(items, item{k, v})
	}

	n := len(items)
	if n == 0 {
		return nil
	}

	// Using int64 indexes internally for simplicity
	capacity := int64(target)
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, capacity+1)
	}

	// Convert to int64 for indexing
	toInt64 := func(v V) int64 {
		switch x := any(v).(type) {
		case int64:
			return x
		case float64:
			return int64(x)
		default:
			return 0
		}
	}

	// Build DP table
	for i := 1; i <= n; i++ {
		val := toInt64(items[i-1].Value)
		for w := int64(0); w <= capacity; w++ {
			if val > w {
				dp[i][w] = dp[i-1][w]
			} else {
				with := dp[i-1][w-val] + val
				without := dp[i-1][w]
				if with > without {
					dp[i][w] = with
				} else {
					dp[i][w] = without
				}
			}
		}
	}

	// Backtrack
	var result []K
	w := capacity
	for i := n; i > 0 && w >= 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			result = append(result, items[i-1].Key)
			w -= toInt64(items[i-1].Value)
		}
	}

	return result
}
