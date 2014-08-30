package ternary_tree

import "fmt"

type tNode struct {
	valueSet               bool
	splitString            string
	loNode, eqNode, hiNode *tNode
}

// insert a string into ternary tree
func insert(t *tNode, str string, count int) {
	// split the string
	vals := []rune(str[count:])
	if len(vals) <= 0 {
		return
	}

	head := string(vals[0])
	var countAdded bool

	if t.isNil() {
		t.setString(head)
		t.valueSet = true
		count++
		countAdded = true
	}

	switch {
	case head > t.splitString:
		if t.hiNode == nil {
			t.hiNode = newNode()
		}
		insert(t.hiNode, str, count)
	case head < t.splitString:
		if t.loNode == nil {
			t.loNode = newNode()
		}
		insert(t.loNode, str, count)
	default:
		if !countAdded {
			count++
		}
		if count == len(str) {
			return
		}
		if t.eqNode == nil {
			t.eqNode = newNode()
		}
		insert(t.eqNode, str, count)
	}
}

// traverse the tree
// return all string
func traverse(t *tNode, str string) (strs []string) {
	if t == nil || t.isNil() {
		return
	}

	str += t.splitString
	strs = append(strs, str)
	strs = append(strs, traverse(t.eqNode, str)...)

	str = ""
	strs = append(strs, traverse(t.loNode, str)...)

	str = ""
	strs = append(strs, traverse(t.hiNode, str)...)

	return
}

// check if the tree contains such prefix
func hasPrefix(t *tNode, prefix string, count int) bool {
	vals := []rune(prefix[count:])
	if len(vals) <= 0 {
		return true
	}
	head := string(vals[0])

	if t == nil || t.isNil() {
		return false
	}

	switch {
	case head > t.splitString:
		return hasPrefix(t.hiNode, prefix, count)
	case head < t.splitString:
		return hasPrefix(t.loNode, prefix, count)
	default:
		count++
		return hasPrefix(t.eqNode, prefix, count)
	}
}

// search for the strings which have prefix
func search(t *tNode, prefix string, count int) (ss []string) {
	// find prefix
	vals := []rune(prefix[count:])
	if len(vals) <= 0 {
		ss = traverse(t, "")
		for i, v := range ss {
			ss[i] = fmt.Sprintf("%s%s", prefix, v)
		}
		return
	}

	head := string(vals[0])

	if t == nil || t.isNil() {
		return
	}

	switch {
	case head < t.splitString:
		ss = append(ss, search(t.loNode, prefix, count)...)
	case head > t.splitString:
		ss = append(ss, search(t.hiNode, prefix, count)...)
	default:
		count++
		ss = append(ss, search(t.eqNode, prefix, count)...)
	}
	return ss
}

func newNode() *tNode                 { return new(tNode) }
func (t *tNode) isNil() bool          { return !t.valueSet }
func (t *tNode) setString(val string) { t.splitString = val }

// implement ternary search tree
type Tree struct {
	root *tNode
}

func NewTernaryTree() *Tree { return &Tree{root: newNode()} }

func (t *Tree) Insert(val string) {
	insert(t.root, val, 0)
}

func (t *Tree) Traverse() []string {
	return traverse(t.root, "")
}

func (t *Tree) SearchPrefix(prefix string) []string {
	return search(t.root, prefix, 0)
}

func (t *Tree) HasPrefix(prefix string) bool {
	return hasPrefix(t.root, prefix, 0)
}
