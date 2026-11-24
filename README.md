# LCA in a BST

> Striver SDE Sheet 136 · Day 20 BST — descend while both values sit on the same side; the first node between them is the ancestor. Go, stdlib only, tested.

## Problem

The lowest common ancestor of two values in a binary *search* tree — where, unlike the general-tree version (#120), ordering should replace exploration.

## Approach

From the root, compare both targets against the current node: both smaller → the LCA lies left; both larger → right; otherwise the node *splits* them (one on each side, or equals one of them) and is, by the BST ordering, exactly the lowest common ancestor. No recursion into both children, no backtracking — one monotone descent. The contrast with #120 is the point of having both on the sheet: structure you can exploit (ordering) converts an O(n) two-sided search into an O(h) corridor walk.

**Complexity:** O(h) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
