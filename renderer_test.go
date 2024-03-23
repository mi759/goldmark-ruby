package ruby

import (
	"bufio"
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

func TestRenderer_All(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc string
		src  string
		give map[string]string
		want string
	}{
		{
			desc: "group ruby",
			src:  "{新世界|しんせかい}",
			give: map[string]string{"新世界": "しんせかい"},
			want: "<ruby>新世界<rt>しんせかい</rt></ruby>",
		},
		{
			desc: "mono ruby",
			src:  "{新世界|しん|せ|かい}",
			give: map[string]string{"新": "しん", "世": "せ", "界": "かい"},
			want: "<ruby>新<rt>しん</rt>世<rt>せ</rt>界<rt>かい</rt></ruby>",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			// setup renderer
			r := goldmark.New().Renderer()
			r.AddOptions(
				renderer.WithNodeRenderers(
					util.Prioritized(&RubyHTMLRenderer{}, 999),
				),
			)

			// setup nodes
			var keys []string
			for k := range tt.give {
				keys = append(keys, k)
			}
			rubyParent := NewRubyParent()
			for _, base := range keys {
				rubyChild := NewRubyChild([]byte(base), []byte(tt.give[base]))
				rubyParent.AppendChild(rubyParent, rubyChild)
			}

			var buf bytes.Buffer
			w := bufio.NewWriter(&buf)

			require.NoError(t, r.Render(w, []byte(tt.src), rubyParent))
			assert.Equal(t, tt.want, buf.String())

		})
	}
}

func TestRenderer_InvalidRubyParent(t *testing.T) {
	t.Parallel()

	var r RubyHTMLRenderer
	_, err := r.RenderRubyParent(bufio.NewWriter(io.Discard), nil /* src */, ast.NewText(), true /* enter */)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected node")
}

func TestRenderer_InvalidRubyChild(t *testing.T) {
	t.Parallel()

	var r RubyHTMLRenderer
	_, err := r.RenderRubyChild(bufio.NewWriter(io.Discard), nil /* src */, ast.NewText(), true /* enter */)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected node")
}
