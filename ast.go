package ruby

import (
	"fmt"

	gast "github.com/yuin/goldmark/ast"
)

// RubyParent node holds RubyChild node(s).
// It itself has no special field.
type RubyParent struct {
	gast.BaseInline
}

// Dump dumps the RubyParent node to stdout for debugging.
func (n *RubyParent) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

var KindRubyParent = gast.NewNodeKind("RubyParent")

// Kind implements Node.Kind.
func (n *RubyParent) Kind() gast.NodeKind {
	return KindRubyParent
}

// NewRubyParent returns a new RubyParent node.
func NewRubyParent() *RubyParent {
	return &RubyParent{}
}

// RubyChild node has the actual information of the ruby.
type RubyChild struct {
	gast.BaseInline
	BaseText []byte // BaseText is the text to be annotated.
	RubyText []byte // RubyText is the annotation.
}

// Dump dumps the RubyChild node to stdout for debugging.
func (n *RubyChild) Dump(source []byte, level int) {
	m := map[string]string{
		"BaseText": fmt.Sprintf("%v", n.BaseText),
		"RubyText": fmt.Sprintf("%v", n.RubyText),
	}
	gast.DumpHelper(n, source, level, m, nil)
}

var KindRubyChild = gast.NewNodeKind("RubyChild")

// Kind implements Node.Kind.
func (n *RubyChild) Kind() gast.NodeKind {
	return KindRubyChild
}

// NewRubyChild returns a new RubyChild node.
func NewRubyChild(baseText []byte, rubyText []byte) *RubyChild {
	return &RubyChild{BaseText: baseText, RubyText: rubyText}
}
