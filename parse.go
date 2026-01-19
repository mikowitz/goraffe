// ABOUTME: Implements DOT language lexer and parser for Graphviz DOT files.
// ABOUTME: Provides tokenization and parsing of DOT graph descriptions into Graph objects.
package goraffe

import (
	"fmt"
	"strings"
	"unicode"
)

// TokenType represents the type of a lexical token.
type TokenType int

const (
	// TokenEOF indicates end of input
	TokenEOF TokenType = iota
	// TokenIdent represents identifiers and keywords
	TokenIdent
	// TokenString represents quoted strings
	TokenString
	// TokenNumber represents numeric literals
	TokenNumber
	// TokenLBrace represents {
	TokenLBrace
	// TokenRBrace represents }
	TokenRBrace
	// TokenLBracket represents [
	TokenLBracket
	// TokenRBracket represents ]
	TokenRBracket
	// TokenLParen represents (
	TokenLParen
	// TokenRParen represents )
	TokenRParen
	// TokenSemi represents ;
	TokenSemi
	// TokenComma represents ,
	TokenComma
	// TokenColon represents :
	TokenColon
	// TokenEqual represents =
	TokenEqual
	// TokenArrow represents -> or --
	TokenArrow
	// TokenHTML represents HTML strings <...>
	TokenHTML
)

// String returns a string representation of the token type.
func (t TokenType) String() string {
	switch t {
	case TokenEOF:
		return "EOF"
	case TokenIdent:
		return "IDENT"
	case TokenString:
		return "STRING"
	case TokenNumber:
		return "NUMBER"
	case TokenLBrace:
		return "{"
	case TokenRBrace:
		return "}"
	case TokenLBracket:
		return "["
	case TokenRBracket:
		return "]"
	case TokenLParen:
		return "("
	case TokenRParen:
		return ")"
	case TokenSemi:
		return ";"
	case TokenComma:
		return ","
	case TokenColon:
		return ":"
	case TokenEqual:
		return "="
	case TokenArrow:
		return "ARROW"
	case TokenHTML:
		return "HTML"
	default:
		return "UNKNOWN"
	}
}

// Token represents a lexical token with its type, value, and location.
type Token struct {
	Type  TokenType
	Value string
	Line  int
	Col   int
}

// String returns a string representation of the token.
func (t Token) String() string {
	if t.Value == "" {
		return fmt.Sprintf("%s at %d:%d", t.Type, t.Line, t.Col)
	}
	return fmt.Sprintf("%s(%q) at %d:%d", t.Type, t.Value, t.Line, t.Col)
}

// Lexer tokenizes DOT language input.
type Lexer struct {
	input  string
	pos    int
	line   int
	col    int
	peeked *Token
}

// NewLexer creates a new lexer for the given input string.
func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
		pos:   0,
		line:  1,
		col:   1,
	}
}

// Next returns the next token from the input.
func (l *Lexer) Next() Token {
	if l.peeked != nil {
		tok := *l.peeked
		l.peeked = nil
		return tok
	}

	l.skipWhitespaceAndComments()

	if l.pos >= len(l.input) {
		return Token{Type: TokenEOF, Line: l.line, Col: l.col}
	}

	startLine := l.line
	startCol := l.col
	ch := l.input[l.pos]

	// Single character tokens
	switch ch {
	case '{':
		l.advance()
		return Token{Type: TokenLBrace, Value: "{", Line: startLine, Col: startCol}
	case '}':
		l.advance()
		return Token{Type: TokenRBrace, Value: "}", Line: startLine, Col: startCol}
	case '[':
		l.advance()
		return Token{Type: TokenLBracket, Value: "[", Line: startLine, Col: startCol}
	case ']':
		l.advance()
		return Token{Type: TokenRBracket, Value: "]", Line: startLine, Col: startCol}
	case '(':
		l.advance()
		return Token{Type: TokenLParen, Value: "(", Line: startLine, Col: startCol}
	case ')':
		l.advance()
		return Token{Type: TokenRParen, Value: ")", Line: startLine, Col: startCol}
	case ';':
		l.advance()
		return Token{Type: TokenSemi, Value: ";", Line: startLine, Col: startCol}
	case ',':
		l.advance()
		return Token{Type: TokenComma, Value: ",", Line: startLine, Col: startCol}
	case ':':
		l.advance()
		return Token{Type: TokenColon, Value: ":", Line: startLine, Col: startCol}
	case '=':
		l.advance()
		return Token{Type: TokenEqual, Value: "=", Line: startLine, Col: startCol}
	}

	// Arrow tokens (-> or --)
	if ch == '-' {
		l.advance()
		if l.pos < len(l.input) {
			next := l.input[l.pos]
			if next == '>' {
				l.advance()
				return Token{Type: TokenArrow, Value: "->", Line: startLine, Col: startCol}
			} else if next == '-' {
				l.advance()
				return Token{Type: TokenArrow, Value: "--", Line: startLine, Col: startCol}
			}
		}
		// Just a minus sign, treat as part of number or identifier
		l.pos--
		l.col--
	}

	// HTML strings
	if ch == '<' {
		return l.scanHTML(startLine, startCol)
	}

	// Quoted strings
	if ch == '"' {
		return l.scanString(startLine, startCol)
	}

	// Numbers (including negative)
	if unicode.IsDigit(rune(ch)) || (ch == '-' && l.pos+1 < len(l.input) && unicode.IsDigit(rune(l.input[l.pos+1]))) {
		return l.scanNumber(startLine, startCol)
	}

	// Identifiers
	if isIdentStart(ch) {
		return l.scanIdent(startLine, startCol)
	}

	// Unknown character
	l.advance()
	return Token{Type: TokenEOF, Value: string(ch), Line: startLine, Col: startCol}
}

// Peek returns the next token without consuming it.
func (l *Lexer) Peek() Token {
	if l.peeked == nil {
		tok := l.Next()
		l.peeked = &tok
	}
	return *l.peeked
}

// advance moves the position forward by one character.
func (l *Lexer) advance() {
	if l.pos < len(l.input) {
		if l.input[l.pos] == '\n' {
			l.line++
			l.col = 1
		} else {
			l.col++
		}
		l.pos++
	}
}

// skipWhitespaceAndComments skips whitespace and comments.
func (l *Lexer) skipWhitespaceAndComments() {
	for l.pos < len(l.input) {
		ch := l.input[l.pos]

		// Skip whitespace
		if unicode.IsSpace(rune(ch)) {
			l.advance()
			continue
		}

		// Skip // comments
		if ch == '/' && l.pos+1 < len(l.input) && l.input[l.pos+1] == '/' {
			l.advance() // skip first /
			l.advance() // skip second /
			for l.pos < len(l.input) && l.input[l.pos] != '\n' {
				l.advance()
			}
			continue
		}

		// Skip /* */ comments
		if ch == '/' && l.pos+1 < len(l.input) && l.input[l.pos+1] == '*' {
			l.advance() // skip /
			l.advance() // skip *
			for l.pos+1 < len(l.input) {
				if l.input[l.pos] == '*' && l.input[l.pos+1] == '/' {
					l.advance() // skip *
					l.advance() // skip /
					break
				}
				l.advance()
			}
			continue
		}

		// Not whitespace or comment
		break
	}
}

// scanString scans a quoted string token.
func (l *Lexer) scanString(startLine, startCol int) Token {
	l.advance() // skip opening quote
	var sb strings.Builder

	for l.pos < len(l.input) {
		ch := l.input[l.pos]
		if ch == '"' {
			l.advance() // skip closing quote
			return Token{Type: TokenString, Value: sb.String(), Line: startLine, Col: startCol}
		}
		if ch == '\\' && l.pos+1 < len(l.input) {
			l.advance()
			next := l.input[l.pos]
			switch next {
			case 'n':
				sb.WriteByte('\n')
			case 't':
				sb.WriteByte('\t')
			case 'r':
				sb.WriteByte('\r')
			case '\\':
				sb.WriteByte('\\')
			case '"':
				sb.WriteByte('"')
			default:
				sb.WriteByte(next)
			}
			l.advance()
		} else {
			sb.WriteByte(ch)
			l.advance()
		}
	}

	// Unterminated string
	return Token{Type: TokenString, Value: sb.String(), Line: startLine, Col: startCol}
}

// scanHTML scans an HTML string token (balanced angle brackets).
func (l *Lexer) scanHTML(startLine, startCol int) Token {
	l.advance() // skip opening <
	var sb strings.Builder
	depth := 1

	for l.pos < len(l.input) && depth > 0 {
		ch := l.input[l.pos]
		if ch == '<' {
			depth++
		} else if ch == '>' {
			depth--
			if depth == 0 {
				l.advance() // skip closing >
				return Token{Type: TokenHTML, Value: sb.String(), Line: startLine, Col: startCol}
			}
		}
		sb.WriteByte(ch)
		l.advance()
	}

	// Unterminated HTML
	return Token{Type: TokenHTML, Value: sb.String(), Line: startLine, Col: startCol}
}

// scanNumber scans a number token.
func (l *Lexer) scanNumber(startLine, startCol int) Token {
	var sb strings.Builder

	// Handle negative sign
	if l.input[l.pos] == '-' {
		sb.WriteByte('-')
		l.advance()
	}

	// Scan digits
	for l.pos < len(l.input) && (unicode.IsDigit(rune(l.input[l.pos])) || l.input[l.pos] == '.') {
		sb.WriteByte(l.input[l.pos])
		l.advance()
	}

	return Token{Type: TokenNumber, Value: sb.String(), Line: startLine, Col: startCol}
}

// scanIdent scans an identifier token.
func (l *Lexer) scanIdent(startLine, startCol int) Token {
	var sb strings.Builder

	for l.pos < len(l.input) && isIdentChar(l.input[l.pos]) {
		sb.WriteByte(l.input[l.pos])
		l.advance()
	}

	return Token{Type: TokenIdent, Value: sb.String(), Line: startLine, Col: startCol}
}

// isIdentStart returns true if the character can start an identifier.
func isIdentStart(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

// isIdentChar returns true if the character can be part of an identifier.
func isIdentChar(ch byte) bool {
	return isIdentStart(ch) || (ch >= '0' && ch <= '9')
}

// Parser parses DOT language input into a Graph.
type Parser struct {
	lexer   *Lexer
	current Token
}

// newParser creates a new parser for the given input string.
func newParser(input string) *Parser {
	p := &Parser{
		lexer: NewLexer(input),
	}
	p.advance() // Load first token
	return p
}

// advance moves to the next token.
func (p *Parser) advance() {
	p.current = p.lexer.Next()
}

// expect consumes a token of the expected type or returns an error.
func (p *Parser) expect(expected TokenType) error {
	if p.current.Type != expected {
		return fmt.Errorf("expected %s, got %s at %d:%d",
			expected, p.current.Type, p.current.Line, p.current.Col)
	}
	p.advance()
	return nil
}

// match returns true if the current token matches the given type.
func (p *Parser) match(tokenType TokenType) bool {
	return p.current.Type == tokenType
}

// matchKeyword returns true if the current token is an identifier with the given value.
func (p *Parser) matchKeyword(keyword string) bool {
	return p.current.Type == TokenIdent && p.current.Value == keyword
}

// parseGraph parses a complete DOT graph.
// Syntax: [strict] (graph|digraph) [ID] { stmt_list }
func (p *Parser) parseGraph() (*Graph, error) {
	// Check for 'strict' keyword
	strict := false
	if p.matchKeyword("strict") {
		strict = true
		p.advance()
	}

	// Expect 'graph' or 'digraph'
	if !p.matchKeyword("graph") && !p.matchKeyword("digraph") {
		return nil, fmt.Errorf("expected 'graph' or 'digraph' at %d:%d", p.current.Line, p.current.Col)
	}

	directed := p.current.Value == "digraph"
	p.advance()

	// Optional graph name
	var name string
	if p.current.Type == TokenIdent || p.current.Type == TokenString {
		name = p.current.Value
		p.advance()
	}

	// Expect opening brace
	if err := p.expect(TokenLBrace); err != nil {
		return nil, err
	}

	// Create graph with options
	opts := []GraphOption{}
	if directed {
		opts = append(opts, Directed)
	} else {
		opts = append(opts, Undirected)
	}
	if strict {
		opts = append(opts, Strict)
	}

	g := NewGraph(opts...)
	g.name = name

	// Parse statements
	if err := p.parseStmtList(g); err != nil {
		return nil, err
	}

	// Expect closing brace
	if err := p.expect(TokenRBrace); err != nil {
		return nil, err
	}

	return g, nil
}

// parseStmtList parses a list of statements until a closing brace.
func (p *Parser) parseStmtList(g *Graph) error {
	for !p.match(TokenRBrace) && !p.match(TokenEOF) {
		if err := p.parseStmt(g); err != nil {
			return err
		}

		// Skip optional semicolon
		if p.match(TokenSemi) {
			p.advance()
		}
	}
	return nil
}

// parseStmt parses a single statement (skeleton implementation).
// Currently just skips unknown statements to allow basic graph parsing.
func (p *Parser) parseStmt(g *Graph) error {
	// For now, just recognize and skip different statement types

	// Check for keywords: node, edge, graph, subgraph
	if p.matchKeyword("node") || p.matchKeyword("edge") || p.matchKeyword("graph") {
		// Skip keyword and following attribute list if present
		p.advance()
		if p.match(TokenLBracket) {
			return p.skipAttrList()
		}
		return nil
	}

	if p.matchKeyword("subgraph") {
		// Skip subgraph (will implement later)
		return p.skipSubgraph()
	}

	// Check for bare '{' (anonymous subgraph)
	if p.match(TokenLBrace) {
		return p.skipSubgraph()
	}

	// Otherwise, skip any node or edge statement
	// This is a placeholder - we'll implement proper parsing later
	return p.skipStatement()
}

// skipAttrList skips over an attribute list [...].
func (p *Parser) skipAttrList() error {
	if err := p.expect(TokenLBracket); err != nil {
		return err
	}

	depth := 1
	for depth > 0 && !p.match(TokenEOF) {
		if p.match(TokenLBracket) {
			depth++
		} else if p.match(TokenRBracket) {
			depth--
		}
		p.advance()
	}

	return nil
}

// skipSubgraph skips over a subgraph statement.
func (p *Parser) skipSubgraph() error {
	// Skip 'subgraph' keyword if present
	if p.matchKeyword("subgraph") {
		p.advance()
		// Skip optional name
		if p.current.Type == TokenIdent || p.current.Type == TokenString {
			p.advance()
		}
	}

	// Expect opening brace
	if err := p.expect(TokenLBrace); err != nil {
		return err
	}

	// Skip until closing brace
	depth := 1
	for depth > 0 && !p.match(TokenEOF) {
		if p.match(TokenLBrace) {
			depth++
		} else if p.match(TokenRBrace) {
			depth--
		}
		p.advance()
	}

	return nil
}

// skipStatement skips tokens until we hit a semicolon or statement boundary.
func (p *Parser) skipStatement() error {
	// Skip tokens until semicolon, closing brace, or EOF
	for !p.match(TokenSemi) && !p.match(TokenRBrace) && !p.match(TokenEOF) {
		// Handle nested braces and brackets
		if p.match(TokenLBrace) {
			if err := p.skipSubgraph(); err != nil {
				return err
			}
			continue
		}
		if p.match(TokenLBracket) {
			if err := p.skipAttrList(); err != nil {
				return err
			}
			continue
		}
		p.advance()
	}
	return nil
}
