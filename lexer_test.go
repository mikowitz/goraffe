package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer_SimpleTokens(t *testing.T) {
	asrt := assert.New(t)

	input := "{ } [ ] ( ) ; , : ="
	lexer := NewLexer(input)

	expected := []TokenType{
		TokenLBrace, TokenRBrace,
		TokenLBracket, TokenRBracket,
		TokenLParen, TokenRParen,
		TokenSemi, TokenComma, TokenColon, TokenEqual,
		TokenEOF,
	}

	for i, expectedType := range expected {
		tok := lexer.Next()
		asrt.Equal(expectedType, tok.Type, "Token %d should be %s, got %s", i, expectedType, tok.Type)
	}
}

func TestLexer_Identifiers(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		input    string
		expected []string
	}{
		{"graph", []string{"graph"}},
		{"digraph", []string{"digraph"}},
		{"subgraph", []string{"subgraph"}},
		{"node", []string{"node"}},
		{"edge", []string{"edge"}},
		{"strict", []string{"strict"}},
		{"Node1 node2 _underscore", []string{"Node1", "node2", "_underscore"}},
		{"abc123 x_y_z", []string{"abc123", "x_y_z"}},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		for i, expected := range tt.expected {
			tok := lexer.Next()
			asrt.Equal(TokenIdent, tok.Type, "Token %d should be IDENT", i)
			asrt.Equal(expected, tok.Value, "Token %d value should be %q, got %q", i, expected, tok.Value)
		}
	}
}

func TestLexer_QuotedStrings(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{`"hello"`, "hello"},
		{`"hello world"`, "hello world"},
		{`"with \"quotes\""`, `with "quotes"`},
		{`"with \\backslash"`, `with \backslash`},
		{`"with\nnewline"`, "with\nnewline"},
		{`"with\ttab"`, "with\ttab"},
		{`""`, ""},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		tok := lexer.Next()
		asrt.Equal(TokenString, tok.Type, "Should be STRING token")
		asrt.Equal(tt.expected, tok.Value, "String value should be %q, got %q", tt.expected, tok.Value)
	}
}

func TestLexer_HTMLStrings(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"<html>", "html"},
		{"<TABLE>", "TABLE"},
		{"<B>", "B"},
		{"<<TABLE><TR><TD>cell</TD></TR></TABLE>>", "<TABLE><TR><TD>cell</TD></TR></TABLE>"},
		{"<>", ""},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		tok := lexer.Next()
		asrt.Equal(TokenHTML, tok.Type, "Should be HTML token")
		asrt.Equal(tt.expected, tok.Value, "HTML value should be %q, got %q", tt.expected, tok.Value)
	}
}

func TestLexer_Arrows(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"->", "->"},
		{"--", "--"},
		{"A -> B", "->"},
		{"A -- B", "--"},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		// Skip to arrow token
		for {
			tok := lexer.Next()
			if tok.Type == TokenArrow {
				asrt.Equal(tt.expected, tok.Value, "Arrow value should be %q, got %q", tt.expected, tok.Value)
				break
			}
			if tok.Type == TokenEOF {
				t.Errorf("Expected arrow token in %q", tt.input)
				break
			}
		}
	}
}

func TestLexer_Comments(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single line comment",
			input:    "A // this is a comment\nB",
			expected: []string{"A", "B"},
		},
		{
			name:     "multi-line comment",
			input:    "A /* this is\na comment */ B",
			expected: []string{"A", "B"},
		},
		{
			name:     "multiple comments",
			input:    "// comment1\nA // comment2\n/* comment3 */ B",
			expected: []string{"A", "B"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexer := NewLexer(tt.input)
			for i, expected := range tt.expected {
				tok := lexer.Next()
				asrt.Equal(TokenIdent, tok.Type, "Token %d should be IDENT", i)
				asrt.Equal(expected, tok.Value, "Token %d value should be %q, got %q", i, expected, tok.Value)
			}
		})
	}
}

func TestLexer_CompleteGraph(t *testing.T) {
	asrt := assert.New(t)

	input := `digraph G {
		A [shape=box];
		B [label="Node B"];
		A -> B [label="edge"];
	}`

	lexer := NewLexer(input)

	expected := []struct {
		tokenType TokenType
		value     string
	}{
		{TokenIdent, "digraph"},
		{TokenIdent, "G"},
		{TokenLBrace, "{"},
		{TokenIdent, "A"},
		{TokenLBracket, "["},
		{TokenIdent, "shape"},
		{TokenEqual, "="},
		{TokenIdent, "box"},
		{TokenRBracket, "]"},
		{TokenSemi, ";"},
		{TokenIdent, "B"},
		{TokenLBracket, "["},
		{TokenIdent, "label"},
		{TokenEqual, "="},
		{TokenString, "Node B"},
		{TokenRBracket, "]"},
		{TokenSemi, ";"},
		{TokenIdent, "A"},
		{TokenArrow, "->"},
		{TokenIdent, "B"},
		{TokenLBracket, "["},
		{TokenIdent, "label"},
		{TokenEqual, "="},
		{TokenString, "edge"},
		{TokenRBracket, "]"},
		{TokenSemi, ";"},
		{TokenRBrace, "}"},
		{TokenEOF, ""},
	}

	for i, exp := range expected {
		tok := lexer.Next()
		asrt.Equal(exp.tokenType, tok.Type, "Token %d type should be %s, got %s", i, exp.tokenType, tok.Type)
		if exp.value != "" {
			asrt.Equal(exp.value, tok.Value, "Token %d value should be %q, got %q", i, exp.value, tok.Value)
		}
	}
}

func TestLexer_Numbers(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"123", "123"},
		{"0", "0"},
		{"3.14", "3.14"},
		{"-5", "-5"},
		{"-2.5", "-2.5"},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		tok := lexer.Next()
		asrt.Equal(TokenNumber, tok.Type, "Should be NUMBER token")
		asrt.Equal(tt.expected, tok.Value, "Number value should be %q, got %q", tt.expected, tok.Value)
	}
}

func TestLexer_Peek(t *testing.T) {
	asrt := assert.New(t)

	input := "A B C"
	lexer := NewLexer(input)

	// Peek should not consume token
	tok1 := lexer.Peek()
	asrt.Equal(TokenIdent, tok1.Type)
	asrt.Equal("A", tok1.Value)

	// Peeking again should return same token
	tok2 := lexer.Peek()
	asrt.Equal(tok1, tok2)

	// Next should return the peeked token
	tok3 := lexer.Next()
	asrt.Equal(tok1, tok3)

	// Next token should be different
	tok4 := lexer.Next()
	asrt.Equal(TokenIdent, tok4.Type)
	asrt.Equal("B", tok4.Value)
}

func TestLexer_LineAndColumn(t *testing.T) {
	asrt := assert.New(t)

	input := "A\nB\n  C"
	lexer := NewLexer(input)

	tok1 := lexer.Next()
	asrt.Equal(1, tok1.Line, "Token A should be on line 1")
	asrt.Equal(1, tok1.Col, "Token A should be at column 1")

	tok2 := lexer.Next()
	asrt.Equal(2, tok2.Line, "Token B should be on line 2")
	asrt.Equal(1, tok2.Col, "Token B should be at column 1")

	tok3 := lexer.Next()
	asrt.Equal(3, tok3.Line, "Token C should be on line 3")
	asrt.Equal(3, tok3.Col, "Token C should be at column 3")
}

func TestLexer_MixedTokens(t *testing.T) {
	asrt := assert.New(t)

	input := `node1 -> node2 [label="edge label"]`
	lexer := NewLexer(input)

	tokens := []Token{}
	for {
		tok := lexer.Next()
		tokens = append(tokens, tok)
		if tok.Type == TokenEOF {
			break
		}
	}

	asrt.Equal(9, len(tokens), "Should have 9 tokens including EOF")
	asrt.Equal(TokenIdent, tokens[0].Type)
	asrt.Equal("node1", tokens[0].Value)
	asrt.Equal(TokenArrow, tokens[1].Type)
	asrt.Equal("->", tokens[1].Value)
	asrt.Equal(TokenIdent, tokens[2].Type)
	asrt.Equal("node2", tokens[2].Value)
}

func TestLexer_EdgeCases(t *testing.T) {
	asrt := assert.New(t)

	t.Run("empty input", func(t *testing.T) {
		lexer := NewLexer("")
		tok := lexer.Next()
		asrt.Equal(TokenEOF, tok.Type)
	})

	t.Run("only whitespace", func(t *testing.T) {
		lexer := NewLexer("   \n  \t  \n")
		tok := lexer.Next()
		asrt.Equal(TokenEOF, tok.Type)
	})

	t.Run("only comments", func(t *testing.T) {
		lexer := NewLexer("// comment\n/* another */")
		tok := lexer.Next()
		asrt.Equal(TokenEOF, tok.Type)
	})

	t.Run("unterminated string", func(t *testing.T) {
		lexer := NewLexer(`"unterminated`)
		tok := lexer.Next()
		asrt.Equal(TokenString, tok.Type)
		asrt.Equal("unterminated", tok.Value)
	})

	t.Run("unterminated HTML", func(t *testing.T) {
		lexer := NewLexer("<unterminated")
		tok := lexer.Next()
		asrt.Equal(TokenHTML, tok.Type)
	})
}
