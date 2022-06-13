package state

import (
	"strings"

	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/internals/semantic"
)

type State struct {
	text     []string
	semantic *semantic.Semantic
	parser   *parser.Parser
}

func New(text string) *State {
	t := strings.Split(text, "\n")
	return &State{
		text: t,
	}
}

func (s *State) ProcessExcept(line int) {
	if line >= len(s.text) || line == -1 {
		line = len(s.text)
	}

	text := s.text[:line]
	if line+1 < len(s.text) {
		text = append(text, s.text[line+1:]...)
	}

	l := lexer.New(strings.Join(text, "\n"))
	p := parser.New(l)
	program := p.Parse()
	analyzer := semantic.New()
	analyzer.Analyze(program)
	s.semantic = analyzer
	s.parser = p
}
