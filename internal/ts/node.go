package ts

import (
	gots "github.com/smacker/go-tree-sitter"
)

// Node is a wrapper for go-tree-sitter Node.
type Node struct {
	n    *gots.Node
	code []byte
}

func MakeRootNode(code_bytes []byte) Node {
	parser := gots.NewParser()
	parser.SetLanguage(GetLanguage())

	tree := parser.Parse(nil, code_bytes)
	root := tree.RootNode()
	tsnode := root.Child(0)
	if tsnode == nil {
		panic("TODO")
	}

	parser.Close()

	return Node{n: tsnode, code: code_bytes}
}

func (n Node) Content() string {
	return n.n.Content(n.code)
}

func (n Node) HasError() bool {
	return n.n.HasError()
}

func (n Node) IsMissing() bool {
	return n.n.IsMissing()
}

func (n Node) Type() string {
	return n.n.Type()
}

func (n Node) ChildCount() uint32 {
	return n.n.ChildCount()
}

func (n Node) LastChild() Node {
	return n.Child(int(n.ChildCount() - 1))
}

func (n Node) Child(idx int) Node {
	tsn := n.n.Child(idx)
	if tsn == nil {
		panic("can't get child")
	}

	return Node{n: tsn, code: n.code}
}

func (n Node) LineNum() uint32 {
	return n.n.StartPoint().Row + 1
}

func (n Node) Column() uint32 {
	return n.n.StartPoint().Column + 1
}

func (n Node) HasNextSibling() bool {
	tsn := n.n.NextSibling()

	if tsn == nil {
		return false
	}

	return true
}

func (n Node) NextSibling() Node {
	return Node{n: n.n.NextSibling(), code: n.code}
}

func (n Node) PrevSibling() Node {
	return Node{n: n.n.PrevSibling(), code: n.code}
}
