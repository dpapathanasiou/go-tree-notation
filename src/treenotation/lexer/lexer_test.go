package lexer

import (
	"testing"
	"treenotation/treenode"
)

const (
	TestInputOne = `html
  header
    css my stylesheet
    js my javascript
  body
    div おはようございます
    div good morning`
)

func TestNodeLexer(t *testing.T) {
	testOneNodeValues := []string{
		"html",
		"  header",
		"    css my stylesheet",
		"    js my javascript",
		"  body",
		"    div おはようございます",
		"    div good morning",
	}

	l := NewLexer(TestInputOne)
	tree := l.GetTreeNode()
	parsedNodeValues := Traverse(tree, nil)

	if len(testOneNodeValues) != len(parsedNodeValues) {
		t.Fatalf("node count test failed: expected=%d, actual=%d",
			len(testOneNodeValues), len(parsedNodeValues))
	}

	for i, testOutput := range testOneNodeValues {
		node := parsedNodeValues[i]
		if node != testOutput {
			t.Fatalf("node tests[%d] failed: expected=%s, actual=%s",
				i, testOutput, node)
		}
	}
}

func walker(tree *treenode.TreeNode, nodes [][]string) [][]string {
	var children []string
	for _, child := range tree.Children {
		children = append(children, child.Value)
	}

	nodes = append(nodes, children)
	for _, child := range tree.Children {
		nodes = walker(child, nodes)
	}
	return nodes
}

func TestParentChildLexer(t *testing.T) {
	testOneParentChildValues := [][]string{
		{
			"  header",
			"  body",
		},
		{
			"    css my stylesheet",
			"    js my javascript",
		},
		{},
		{},
		{
			"    div おはようございます",
			"    div good morning",
		},
		{},
		{},
	}

	l := NewLexer(TestInputOne)
	actualParentChildValues := walker(l.GetTreeNode(), nil)

	for i, testChildren := range testOneParentChildValues {
		for j, testChildValue := range testChildren {
			if testChildValue != actualParentChildValues[i][j] {
				t.Fatalf("parent/child tests[%d,%d] failed: expected=%s, actual=%s",
					i, j, testChildValue, actualParentChildValues[i][j])
			}
		}
	}
}
