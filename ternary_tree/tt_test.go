package ternary_tree

import "testing"

func Test_TT(t *testing.T) {
	tt := NewTernaryTree()

	tt.Insert("fuck")
	tt.Insert("fuck you")
	tt.Insert("fu")
	tt.Insert("insert")
	tt.Insert("hello")
	tt.Insert("hi")
	tt.Insert("be")

	t.Log(tt.HasPrefix("fuck"))

	t.Log(tt.Traverse())

	t.Log(tt.SearchPrefix("fuck"))

	t.Log(tt.SearchPrefix("in"))

	t.Log(tt.SearchPrefix("helloworld"))
}
