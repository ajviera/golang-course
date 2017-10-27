package tree

import "sort"

// Tree expose
type Tree struct {
	value       int
	left, right *Tree
}

// Add expose
func (t *Tree) Add(value int) *Tree {
	if t == nil {
		t = new(Tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = t.left.Add(value)
	} else {
		t.right = t.right.Add(value)
	}
	return t
}

// PrintInOrder expose
func (t *Tree) PrintInOrder() []int {
	var slice []int
	result := t.GetElements(slice)
	sort.Ints(result)
	return result
}

// GetElements expose
func (t *Tree) GetElements(slice []int) []int {
	if t == nil {
		return slice
	}
	slice = t.left.GetElements(slice)
	slice = t.right.GetElements(slice)
	slice = append(slice, t.value)
	return slice
}
