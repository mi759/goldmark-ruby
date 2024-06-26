package ruby

import (
	"bytes"
	"strings"
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
			want: "<p>ああ、<ruby>素晴<rt>すば</rt></ruby>らしき<ruby>新<rt>しん</rt>世<rt>せ</rt>界<rt>かい</rt></ruby>！</p>",
		},
		{
			desc: "no multilne",
			give: `ああ、{素晴|
すば\}らしき{新世界
|しん|せ|かい}！`,
			want: `<p>ああ、{素晴|
すば}らしき{新世界
|しん|せ|かい}！</p>`,
		},
		{
			desc: "no formatting",
			give: "ああ、{*素晴*|_すば_}らしき{**新世界**|しん|せ|かい}！",
			want: "<p>ああ、<ruby>*素晴*<rt>_すば_</rt></ruby>らしき<ruby>**新世界**<rt>しん|せ|かい</rt></ruby>！</p>",
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
			require.Equal(t, strings.TrimSpace(tt.want), strings.TrimSpace(buf.String()))
		})
	}

}
