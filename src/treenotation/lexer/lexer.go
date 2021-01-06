package lexer

import (
	"sort"
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

// groupByIndent groups the node strings by indent level, returning them in a map whose key is the indention count
func groupByIndent(nodes []string) map[int][]string {
	var m = make(map[int][]string)
	for _, node := range nodes {
		indents := 0
		for _, word := range strings.Split(node, treenode.EdgeSymbol) {
			if word != "" {
				break
			}
			indents++
		}
		m[indents] = append(m[indents], strings.TrimLeft(node, treenode.EdgeSymbol))
	}
	return m
}

// generateTreeNodes uses groupByIndent to convert the list node strings into the treenode, from the root down
func generateTreeNodes(nodes []string) *treenode.TreeNode {
	grouped := groupByIndent(nodes)

	// get the keys (indent levels) and sort them by ascending order
	keys := make([]int, 0, len(grouped))
	for k := range grouped {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var root *treenode.TreeNode
	for _, indent := range keys {
		for _, node := range grouped[indent] {
			root = addToTree(root, node)
		}
	}
	return root
}

// addToTree recursively creates or appends children to the tree
func addToTree(tree *treenode.TreeNode, node string) *treenode.TreeNode {
	if tree == nil {
		return &treenode.TreeNode{Value: node}
	}
	tree.Children = append(tree.Children, addToTree(nil, node))
	return tree
}
