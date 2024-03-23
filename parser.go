package ruby

import (
	"bytes"
	"unicode/utf8"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type RubyParser struct {
}

var (
	_open  = []byte("{")
	_pipe  = []byte{'|'}
	_close = []byte("}")
)

func (s *RubyParser) Trigger() []byte {
	return _open
}

// Parse parses a markdown ruby.
func (p *RubyParser) Parse(_ ast.Node, block text.Reader, _ parser.Context) ast.Node {
	line, seg := block.PeekLine()

	// pipe and stop must be on the same line
	firstPipe := bytes.Index(line, _pipe)
	stop := bytes.Index(line, _close)
	if firstPipe < 0 || stop < 0 {
		return nil
	}

	baseSeg := text.NewSegment(seg.Start+len(_open), seg.Start+firstPipe)
	baseBytes := block.Value(baseSeg)
	rubySeg := text.NewSegment(seg.Start+firstPipe+1, seg.Start+stop)
	rubyBytes := block.Value(rubySeg)

	// split ruby text at pipes, to determine if it is mono ruby
	var rubys [][]byte
	tmpRubyBytes := rubyBytes
	for {
		idx := bytes.Index(tmpRubyBytes, _pipe)
		if idx < 0 {
			// append the last one
			rubys = append(rubys, tmpRubyBytes)
			break
		}
		rubys = append(rubys, tmpRubyBytes[:idx])
		tmpRubyBytes = tmpRubyBytes[idx+1:] // move to next
	}

	// valid mono ruby
	if utf8.RuneCount(baseBytes) == len(rubys) {
		baseRunes := []rune(string(baseBytes))
		rubyParent := NewRubyParent()
		for i, ruby := range rubys {
			baseRuneBytes := []byte(string(baseRunes[i])) // a rune of base text
			rubyChild := NewRubyChild(baseRuneBytes, ruby)
			rubyParent.AppendChild(rubyParent, rubyChild)
		}
		block.Advance(stop + 1)
		return rubyParent
	}

	// invalid mono ruby or group ruby
	rubyParent := NewRubyParent()
	rubyChild := NewRubyChild(baseBytes, rubyBytes)
	rubyParent.AppendChild(rubyParent, rubyChild)

	block.Advance(stop + 1)
	return rubyParent
}
