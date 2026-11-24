package main

import "testing"

func TestLCA(t *testing.T) {
	root := tn(6, tn(2, tn(0, nil, nil), tn(4, tn(3, nil, nil), tn(5, nil, nil))), tn(8, tn(7, nil, nil), tn(9, nil, nil)))
	cases := []struct {
		a, b, want int
	}{
		{2, 8, 6},
		{3, 5, 4},
		{0, 5, 2},
		{2, 4, 2}, // ancestor-descendant
		{7, 9, 8},
		{3, 3, 3},
	}
	for _, c := range cases {
		if got := LCA(root, c.a, c.b); got == nil || got.Val != c.want {
			t.Errorf("LCA(%d,%d) = %v, want %d", c.a, c.b, got, c.want)
		}
	}
	if LCA(nil, 1, 2) != nil {
		t.Error("empty tree should give nil")
	}
}
