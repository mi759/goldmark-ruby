package ruby

import (
	"fmt"

	gast "github.com/yuin/goldmark/ast"
)

type RubyParent struct {
	gast.BaseInline
}

func (n *RubyParent) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

var KindRubyParent = gast.NewNodeKind("RubyParent")

func (n *RubyParent) Kind() gast.NodeKind {
	return KindRubyParent
}

func NewRubyParent() *RubyParent {
	return &RubyParent{}
}

type RubyChild struct {
	gast.BaseInline
	BaseText []byte
	RubyText []byte
}

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
