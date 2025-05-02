# Knapsack Go Library

This library provides two methods to solve the knapsack problem:

- `Greedy`: A fast approximation
- `Standard`: Dynamic programming solution

### Installation

```bash
go get github.com/ifanwar/knapsack

res := knapsack.Greedy(ctx, map[string]int64{"a": 10, "b": 20}, 25)
// or
res := knapsack.Standard(ctx, map[string]float64{"x": 5.5, "y": 9.9}, 10)
