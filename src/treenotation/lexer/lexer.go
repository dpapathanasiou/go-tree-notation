package lexer

import (
	"errors"
	"strings"
	"treenotation/treenode"
)

// Lexer defines the struct for inspecting inputs, in order to produce TreeNode results
type Lexer struct {
	nodes       []string
	currentNode int
}

// NewLexer constructs a new instance of the Lexer for the given string, normalizing it regardless of os origin
func NewLexer(input string) *Lexer {
	replacer := strings.NewReplacer(treenode.IgnoreableSymbol, "")
	nodes := strings.Split(replacer.Replace(input), treenode.NodeBreakSymbol)
	l := &Lexer{
		nodes:       nodes,
		currentNode: 0,
	}
	return l
}

// HasNextNode defines whether or not all the node strings have been traversed
func (l *Lexer) HasNextNode() bool {
	return l.currentNode < len(l.nodes)
}

// NextNode produces the next available node string, or an error if the input has been traversed
func (l *Lexer) NextNode() (string, error) {
	if !l.HasNextNode() {
		return "", errors.New("no more node strings")
	}
	node := l.nodes[l.currentNode]
	l.currentNode++
	return node, nil
}
