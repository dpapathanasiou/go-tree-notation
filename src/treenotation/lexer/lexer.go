package lexer

import (
	"strings"
	"treenotation/treenode"
)

// Lexer defines the struct for inspecting inputs, in order to produce TreeNode results
type Lexer struct {
	tree *treenode.TreeNode
}

// NewLexer constructs a new instance of the Lexer for the given string, normalizing it regardless of os origin
func NewLexer(input string) *Lexer {
	replacer := strings.NewReplacer(treenode.IgnoreableSymbol, "")
	nodes := strings.Split(replacer.Replace(input), treenode.NodeBreakSymbol)
	l := &Lexer{tree: generateTreeNodes(nodes)}
	return l
}

// GetTreeNode returns the parsed TreeNode
func (l *Lexer) GetTreeNode() *treenode.TreeNode {
	return l.tree
}

// Traverse walks all the TreeNodes and returns a list of their node string values
func Traverse(tree *treenode.TreeNode, nodes []string) []string {
	nodes = append(nodes, tree.Value)
	for _, child := range tree.Children {
		nodes = Traverse(child, nodes)
	}
	return nodes
}

// getIndentation counts the number of EdgeSymbol bytes found on the left side of this node string
func getIndentation(node string) int {
	indent := 0
	for _, word := range strings.Split(node, treenode.EdgeSymbol) {
		if word != "" {
			break
		}
		indent++
	}
	return indent
}

// generateTreeNodes converts the list node of strings into the TreeNode, from the root down
func generateTreeNodes(nodes []string) *treenode.TreeNode {
	var tree *treenode.TreeNode
	// track the parent/child relations with a stack:
	// indent (new child) -> push
	// outdent or peer    <- pop
	stack := []*treenode.TreeNode{}
	for _, node := range nodes {
		if tree == nil {
			tree = &treenode.TreeNode{Value: node}
			stack = append(stack, tree)
		} else {
			treeIndent := getIndentation(tree.Value)
			nodeIndent := getIndentation(node)
			if nodeIndent > treeIndent {
				// new child
				child := &treenode.TreeNode{Value: node, Parent: tree}
				tree.Children = append(tree.Children, child)
				stack = append(stack, child)
			} else if nodeIndent == treeIndent {
				// new peer child
				for len(stack) > 0 && nodeIndent == getIndentation(tree.Value) {
					tree = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
				child := &treenode.TreeNode{Value: node, Parent: tree}
				tree.Children = append(tree.Children, child)
				stack = append(stack, child)
			} else if nodeIndent < treeIndent {
				// new parent (child of prior parent)
				for len(stack) > 0 && nodeIndent < getIndentation(tree.Value) {
					tree = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
				child := &treenode.TreeNode{Value: node, Parent: tree}
				tree.Children = append(tree.Children, child)
				stack = append(stack, child)
			}
			tree = stack[len(stack)-1]
		}
	}

	// return the root TreeNode
	for tree.Parent != nil {
		tree = tree.Parent
	}
	return tree
}
