package treenode

// TreeNode defines the core building block: a string value, and list of children
type TreeNode struct {
	Value    string
	Children []*TreeNode
}

const (
	// NodeBreakSymbol delimits nodes (lines)
	NodeBreakSymbol = "\n"

	// IgnoreableSymbol is for Windows, so that NodeBreakSymbol does not have to change based on os type
	IgnoreableSymbol = "\r"

	// EdgeSymbol is used to indicate the parent/child relationship between nodes
	EdgeSymbol = " "

	// WordBreakSymbol delimits words in a string
	WordBreakSymbol = " "
)
