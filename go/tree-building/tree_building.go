package tree

import (
	"fmt"
	"sort"
)

// Node has ID int and Children []*Node
type Node struct {
	ID       int
	Children []*Node
}

// Record has ID int and Parent int
type Record struct {
	ID     int
	Parent int
}

// Build function receive an array of Record as input and return the root node of tree structure.
// if the input cant build a valid tree, it would return an error.
func Build(records []Record) (*Node, error) {
	var root *Node = nil
	nodes := map[int]*Node{}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, record := range records {
		id, parentID := record.ID, record.Parent
		if i != id {
			return nil, fmt.Errorf("nodes must be continuous")
		}
		if id == parentID && id != 0 {
			return nil, fmt.Errorf("non-root node cannot have itself as parent")
		}
		_, ok := nodes[id]
		if ok {
			return nil, fmt.Errorf("duplicate node")
		}
		node := &Node{ID: id}
		nodes[id] = node
		if id == 0 {
			if parentID != 0 {
				return nil, fmt.Errorf("root node has parent")
			}
			root = node
			continue
		}
		parent, ok := nodes[parentID]
		if !ok {
			parent = &Node{ID: parentID}
			nodes[parentID] = parent
		}
		parent.Children = append(parent.Children, node)
	}

	if root == nil && len(records) != 0 {
		return nil, fmt.Errorf("no root node")
	}

	return root, nil
}
