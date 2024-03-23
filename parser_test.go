package ruby

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func TestParser(t *testing.T) {
	t.Parallel()

	tests := []struct {
		give            string
		wantNumOfChilds int
		wantBaseTexts   []string
		wantRubyTexts   []string
	}{
		{give: "{新世界|しんせかい}", wantNumOfChilds: 1, wantBaseTexts: []string{"新世界"}, wantRubyTexts: []string{"しんせかい"}},
		{give: "{新世界|しん|せ|かい}", wantNumOfChilds: 3, wantBaseTexts: []string{"新", "世", "界"}, wantRubyTexts: []string{"しん", "せ", "かい"}},
		{give: "{新世界|し|ん|せ|か|い}", wantNumOfChilds: 1, wantBaseTexts: []string{"新世界"}, wantRubyTexts: []string{"し|ん|せ|か|い"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.give, func(t *testing.T) {
			r := text.NewReader([]byte(tt.give))

			var p RubyParser
			got := p.Parse(nil /* parent */, r, parser.NewContext())
			require.NotNil(t, got, "expected Node, got nil")

			_, ok := got.(*RubyParent)
			assert.True(t, ok, "expected Node, got %T", got)

			if assert.Equal(t, tt.wantNumOfChilds, got.ChildCount(), "children mismatch") {
				child := got.FirstChild()
				for i := 0; i < tt.wantNumOfChilds; i++ {
					if rubyChild, ok := child.(*RubyChild); assert.True(t, ok, "expected RubyChild, got %T", child) {
						assert.Equal(t,
							tt.wantBaseTexts[i],
							string(rubyChild.BaseText),
							"ruby text mismatch, expected %s, got %s",
							tt.wantBaseTexts[i],
							rubyChild.BaseText,
						)
						assert.Equal(t,
							tt.wantRubyTexts[i],
							string(rubyChild.RubyText),
							"ruby text mismatch, expected %s, got %s",
							tt.wantRubyTexts[i],
							rubyChild.RubyText,
						)
					}
					child = child.NextSibling()
				}
			}
		})
	}
}
