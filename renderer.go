package ruby

import (
	"fmt"

	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type RubyHTMLRenderer struct{}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
// It tells goldmark which render function to call when it encounters ruby nodes in the AST.
func (r *RubyHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindRubyParent, r.RenderRubyParent)
	reg.Register(KindRubyChild, r.RenderRubyChild)
}

// RenderRubyParent renders the provided RubyParent node.
func (r *RubyHTMLRenderer) RenderRubyParent(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	n, ok := node.(*RubyParent)
	if !ok {
		return gast.WalkStop, fmt.Errorf("unexpected node %T, expected *ruby.RubyParent", n)
	}

	if entering {
		_, _ = w.WriteString("<ruby>")
	} else {
		_, _ = w.WriteString("</ruby>")
	}

	return gast.WalkContinue, nil
}

// RenderRubyParent renders the provided RubyChild node.
func (r *RubyHTMLRenderer) RenderRubyChild(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	n, ok := node.(*RubyChild)
	if !ok {
		return gast.WalkStop, fmt.Errorf("unexpected node %T, expected *ruby.RubyChild", n)
	}

	if entering {
		_, _ = w.WriteString(string(n.BaseText))

	} else {
		_, _ = w.WriteString("<rt>" + string(n.RubyText) + "</rt>")
	}

	return gast.WalkContinue, nil
}
