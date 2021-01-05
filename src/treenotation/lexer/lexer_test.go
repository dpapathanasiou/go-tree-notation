package lexer

import "testing"

func TestInputLexer(t *testing.T) {
	testInput := `if true
  print Hello world`

	testOutputs := []string{
		"if true",
		"  print Hello world",
	}

	l := NewLexer(testInput)

	for i, testOutput := range testOutputs {
		node, _ := l.NextNode()

		if node != testOutput {
			t.Fatalf("tests[%d] failed: expected=%s, actual=%s",
				i, testOutput, node)
		}

	}
}
