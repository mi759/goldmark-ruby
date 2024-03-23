package ruby

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type RubyExtender struct {
}

// Ruby is an extension that allow you to use various ruby expressions.
var Ruby = &RubyExtender{}

func (e *RubyExtender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(&RubyParser{}, 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&RubyHTMLRenderer{}, 500),
	))
}
