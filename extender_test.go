package ruby

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark"
)

func TestExtend(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc string
		give string
		want string
	}{
		{
			desc: "default",
			give: "ああ、{素晴|すば}らしき{新世界|しん|せ|かい}！",
			want: "<p>ああ、<ruby>素晴<rt>すば</rt></ruby>らしき<ruby>新<rt>しん</rt>世<rt>せ</rt>界<rt>かい</rt></ruby>！</p>\n",
		},
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(
			Ruby,
		),
	)

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer

			require.NoError(t, markdown.Convert([]byte(tt.give), &buf))
			require.Equal(t, tt.want, buf.String())
		})
	}

}
