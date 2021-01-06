package lexer

import (
	"testing"
)

func TestInputLexer(t *testing.T) {
	testInput := `html
  body
    div おはようございます
    div good morning`

	testNodeValues := []string{
		"html",
		"body",
		"div おはようございます",
		"div good morning",
	}
	expected := len(testNodeValues)

	l := NewLexer(testInput)
	parsedNodeValues := Traverse(l.GetTreeNode(), nil)

	if expected != len(parsedNodeValues) {
		t.Fatalf("node count test failed: expected=%d, actual=%d", expected, len(parsedNodeValues))
	}

	for i, testOutput := range testNodeValues {
		node := parsedNodeValues[i]
		if node != testOutput {
			t.Fatalf("tests[%d] failed: expected=%s, actual=%s", i, testOutput, node)
		}
	}

}
