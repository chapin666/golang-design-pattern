package composite

import (
	"testing"
)

func TestTree(t *testing.T) {
	root := NewConcreateCompany("Root")
	root.add(NewLeaf("叶子节点1"))
	root.add(NewLeaf("叶子节点2"))

	second := NewConcreateCompany("第二层")
	root.add(second)

	root.display(1)
}
