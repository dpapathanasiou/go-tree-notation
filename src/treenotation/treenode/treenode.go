package treenode

// TreeNode defines the core building block
type TreeNode struct {
	Parent   *TreeNode
	Children []*TreeNode
	Line     string
}

const (
	// NodeBreakSymbol delimits nodes (TODO: find an os-independent definition)
	NodeBreakSymbol = "\n"

	// EdgeSymbol is used to indicate the parent/child relationship between nodes
	EdgeSymbol = " "

	// WordBreakSymbol delimits words in a string
	WordBreakSymbol = " "
)
