package certificate

import (
	"bytes"
)

// Lookup looks up the given path in the certificate tree.
func Lookup(path [][]byte, node Node) []byte {
	if len(path) == 0 {
		switch n := node.(type) {
		case Leaf:
			return n
		default:
			return nil
		}
	}

	n := findLabel(flattenNode(node), path[0])
	if n != nil {
		return Lookup(path[1:], *n)
	}
	return nil
}

// LookupPath returns a path from the given labels.
func LookupPath(p ...string) [][]byte {
	var path [][]byte
	for _, p := range p {
		path = append(path, []byte(p))
	}
	return path
}

func LookupNode(path [][]byte, node Node) *Node {
	if len(path) == 0 {
		return &node
	}

	n := findLabel(flattenNode(node), path[0])
	if n != nil {
		return LookupNode(path[1:], *n)
	}
	return nil
}

func findLabel(nodes []Node, label Label) *Node {
	for _, node := range nodes {
		switch n := node.(type) {
		case Labeled:
			if bytes.Equal(label, n.Label) {
				return &n.Tree
			}
		}
	}
	return nil
}

func flattenNode(node Node) []Node {
	switch n := node.(type) {
	case Empty:
		return nil
	case Fork:
		return append(
			flattenNode(n.LeftTree),
			flattenNode(n.RightTree)...,
		)
	default:
		return []Node{node}
	}
}
