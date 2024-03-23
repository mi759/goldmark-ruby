# goldmark-ruby
[![Go Reference](https://pkg.go.dev/badge/github.com/mi759/goldmark-ruby.svg)](https://pkg.go.dev/github.com/mi759/goldmark-ruby)

goldmark-ruby is an extension for the [goldmark](https://github.com/yuin/goldmark) that parses [DenDenMarkdown](https://github.com/denshoch/DenDenMarkdown) style [rubys](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby) (like `{漢字|かんじ}`).

This extension has nothing to do with the original authors of DenDenMarkdown.

## Installation
```
go get github.com/mi759/goldmark-ruby
```

## Usage
```go
import (
	"bytes"
	"fmt"

	ruby "github.com/mi759/goldmark-ruby"
	"github.com/yuin/goldmark"
)

func main() {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			ruby.Ruby,
		),
	)
	source := `ああ、{素晴|すば}らしき{新世界|しん|せ|かい}！`
	var buf bytes.Buffer
	err := markdown.Convert([]byte(source), &buf)
	if err != nil {
		panic(err)
	}
	fmt.Print(buf.String())
}
```

## License
MIT
