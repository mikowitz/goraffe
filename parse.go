// ABOUTME: Implements DOT language lexer and parser for Graphviz DOT files.
// ABOUTME: Provides tokenization and parsing of DOT graph descriptions into Graph objects.
package goraffe

import (
	"fmt"
	"io"
	"os"
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

// parseStmt parses a single statement.
func (p *Parser) parseStmt(g *Graph) error {
	// Check for keywords: node, edge, graph, subgraph
	if p.matchKeyword("node") || p.matchKeyword("edge") || p.matchKeyword("graph") {
		// Default attribute statements
		keyword := p.current.Value
		p.advance()
		if p.match(TokenLBracket) {
			attrs, err := p.parseAttrList()
			if err != nil {
				return err
			}
			// Apply default attributes based on keyword
			return p.applyDefaultAttrs(g, keyword, attrs)
		}
		return nil
	}

	if p.matchKeyword("subgraph") || p.match(TokenLBrace) {
		// Parse the subgraph
		sg, err := p.parseSubgraph(g)
		if err != nil {
			return err
		}

		// Check if this subgraph is followed by an arrow (edge statement)
		if p.match(TokenArrow) {
			// Subgraph is used as edge endpoint
			nodeIDs := make([]string, 0)
			for _, node := range sg.Nodes() {
				nodeIDs = append(nodeIDs, node.ID())
			}
			return p.parseEdgeStmtWithNodes(g, nodeIDs)
		}

		// Just a standalone subgraph declaration
		return nil
	}

	// Try to parse as node or edge statement
	return p.parseNodeOrEdgeStmt(g)
}

// parseID parses an identifier (ID in DOT grammar).
// Can be: identifier, quoted string, number, or HTML string.
func (p *Parser) parseID() (string, error) {
	switch p.current.Type {
	case TokenIdent, TokenString, TokenNumber, TokenHTML:
		id := p.current.Value
		p.advance()
		return id, nil
	default:
		return "", fmt.Errorf("expected ID, got %s at %d:%d", p.current.Type, p.current.Line, p.current.Col)
	}
}

// parseAttrList parses an attribute list [attr=value, attr=value, ...].
// Returns a map of attribute key-value pairs.
func (p *Parser) parseAttrList() (map[string]string, error) {
	attrs := make(map[string]string)

	if err := p.expect(TokenLBracket); err != nil {
		return nil, err
	}

	for !p.match(TokenRBracket) && !p.match(TokenEOF) {
		// Parse attribute name
		if !p.match(TokenIdent) {
			return nil, fmt.Errorf("expected attribute name, got %s at %d:%d", p.current.Type, p.current.Line, p.current.Col)
		}
		name := p.current.Value
		p.advance()

		// Expect '='
		if err := p.expect(TokenEqual); err != nil {
			return nil, err
		}

		// Parse attribute value
		value, err := p.parseID()
		if err != nil {
			return nil, err
		}

		attrs[name] = value

		// Skip optional comma or semicolon
		if p.match(TokenComma) || p.match(TokenSemi) {
			p.advance()
		}
	}

	if err := p.expect(TokenRBracket); err != nil {
		return nil, err
	}

	return attrs, nil
}

// parseNodeOrEdgeStmt parses either a node or edge statement.
// This is tricky because we need lookahead to distinguish them.
func (p *Parser) parseNodeOrEdgeStmt(g *Graph) error {
	// Check if this starts with a subgraph (for edge endpoints)
	if p.matchKeyword("subgraph") || p.match(TokenLBrace) {
		// This must be an edge with subgraph as endpoint
		nodeIDs, err := p.parseEdgeEndpoint(g)
		if err != nil {
			return err
		}

		// Subgraphs can only appear in edge statements
		if !p.match(TokenArrow) {
			return fmt.Errorf("subgraph must be followed by edge operator at %d:%d", p.current.Line, p.current.Col)
		}

		return p.parseEdgeStmtWithNodes(g, nodeIDs)
	}

	// Parse first ID
	id, err := p.parseID()
	if err != nil {
		return err
	}

	// Check if this is an edge statement (next token is arrow)
	if p.match(TokenArrow) {
		return p.parseEdgeStmtWithNodes(g, []string{id})
	}

	// Otherwise, it's a node statement
	return p.parseNodeStmt(g, id)
}

// parseEdgeEndpoint parses an edge endpoint, which can be either:
// - A node ID (returns single-element list)
// - A subgraph (returns all node IDs in the subgraph)
func (p *Parser) parseEdgeEndpoint(g *Graph) ([]string, error) {
	// Check if this is a subgraph
	if p.matchKeyword("subgraph") || p.match(TokenLBrace) {
		// Parse the subgraph
		sg, err := p.parseSubgraph(g)
		if err != nil {
			return nil, err
		}

		// Collect all node IDs from the subgraph
		nodes := sg.Nodes()
		nodeIDs := make([]string, len(nodes))
		for i, node := range nodes {
			nodeIDs[i] = node.ID()
		}

		return nodeIDs, nil
	}

	// Otherwise, parse a single node ID
	id, err := p.parseID()
	if err != nil {
		return nil, err
	}

	return []string{id}, nil
}

// parseNodeStmt parses a node statement: nodeID [attributes].
func (p *Parser) parseNodeStmt(g *Graph, id string) error {
	// Parse optional attributes
	var attrs map[string]string
	if p.match(TokenLBracket) {
		var err error
		attrs, err = p.parseAttrList()
		if err != nil {
			return err
		}
	}

	// Create node with attributes
	nodeOpts := p.mapNodeAttributes(attrs)
	node := NewNode(id, nodeOpts...)
	return g.AddNode(node)
}

// parseEdgeStmt parses an edge statement: nodeID arrow nodeID ... [attributes].
// Handles edge chains: A -> B -> C creates edges A->B and B->C.
// Deprecated: Use parseEdgeStmtWithNodes instead.
func (p *Parser) parseEdgeStmt(g *Graph, firstID string) error {
	return p.parseEdgeStmtWithNodes(g, []string{firstID})
}

// parseEdgeStmtWithNodes parses an edge statement where endpoints can be subgraphs.
// Each endpoint is a list of node IDs (single node or all nodes from a subgraph).
// Creates edges between all combinations of nodes at adjacent endpoints.
func (p *Parser) parseEdgeStmtWithNodes(g *Graph, firstNodes []string) error {
	// Store all endpoints as lists of node IDs
	endpoints := [][]string{firstNodes}

	// Parse edge chain
	for p.match(TokenArrow) {
		p.advance() // consume arrow

		// Parse next endpoint (node or subgraph)
		nodeIDs, err := p.parseEdgeEndpoint(g)
		if err != nil {
			return err
		}
		endpoints = append(endpoints, nodeIDs)
	}

	// Parse optional edge attributes
	var attrs map[string]string
	if p.match(TokenLBracket) {
		var err error
		attrs, err = p.parseAttrList()
		if err != nil {
			return err
		}
	}

	// Create edges between each pair of adjacent endpoints
	edgeOpts := p.mapEdgeAttributes(attrs)
	for i := 0; i < len(endpoints)-1; i++ {
		fromNodes := endpoints[i]
		toNodes := endpoints[i+1]

		// Create edges from all nodes in fromNodes to all nodes in toNodes
		for _, fromID := range fromNodes {
			for _, toID := range toNodes {
				from := NewNode(fromID)
				to := NewNode(toID)
				if _, err := g.AddEdge(from, to, edgeOpts...); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// mapNodeAttributes maps parsed attributes to NodeOption functions.
func (p *Parser) mapNodeAttributes(attrs map[string]string) []NodeOption {
	if attrs == nil {
		return nil
	}

	opts := []NodeOption{}

	// Map known attributes
	if label, ok := attrs["label"]; ok {
		opts = append(opts, WithLabel(label))
	}
	if shape, ok := attrs["shape"]; ok {
		opts = append(opts, withShape(Shape(shape)))
	}
	if color, ok := attrs["color"]; ok {
		opts = append(opts, WithColor(color))
	}
	if fillcolor, ok := attrs["fillcolor"]; ok {
		opts = append(opts, WithFillColor(fillcolor))
	}
	if fontname, ok := attrs["fontname"]; ok {
		opts = append(opts, WithFontName(fontname))
	}
	if fontsize, ok := attrs["fontsize"]; ok {
		// Parse fontsize as float - ignore errors for now
		var size float64
		fmt.Sscanf(fontsize, "%f", &size)
		if size > 0 {
			opts = append(opts, WithFontSize(size))
		}
	}

	// Add custom attributes for unknown ones
	knownAttrs := map[string]bool{
		"label": true, "shape": true, "color": true,
		"fillcolor": true, "fontname": true, "fontsize": true,
	}
	for key, value := range attrs {
		if !knownAttrs[key] {
			opts = append(opts, WithNodeAttribute(key, value))
		}
	}

	return opts
}

// mapEdgeAttributes maps parsed attributes to EdgeOption functions.
func (p *Parser) mapEdgeAttributes(attrs map[string]string) []EdgeOption {
	if attrs == nil {
		return nil
	}

	opts := []EdgeOption{}

	// Map known attributes
	if label, ok := attrs["label"]; ok {
		opts = append(opts, WithEdgeLabel(label))
	}
	if color, ok := attrs["color"]; ok {
		opts = append(opts, WithEdgeColor(color))
	}
	if style, ok := attrs["style"]; ok {
		opts = append(opts, WithEdgeStyle(EdgeStyle(style)))
	}
	if arrowhead, ok := attrs["arrowhead"]; ok {
		opts = append(opts, WithArrowHead(ArrowType(arrowhead)))
	}
	if arrowtail, ok := attrs["arrowtail"]; ok {
		opts = append(opts, WithArrowTail(ArrowType(arrowtail)))
	}
	if weight, ok := attrs["weight"]; ok {
		var w float64
		fmt.Sscanf(weight, "%f", &w)
		if w > 0 {
			opts = append(opts, WithWeight(w))
		}
	}

	// Add custom attributes for unknown ones
	knownAttrs := map[string]bool{
		"label": true, "color": true, "style": true,
		"arrowhead": true, "arrowtail": true, "weight": true,
	}
	for key, value := range attrs {
		if !knownAttrs[key] {
			opts = append(opts, WithEdgeAttribute(key, value))
		}
	}

	return opts
}

// applyDefaultAttrs applies default attributes to the graph.
func (p *Parser) applyDefaultAttrs(g *Graph, keyword string, attrs map[string]string) error {
	switch keyword {
	case "node":
		opts := p.mapNodeAttributes(attrs)
		// Apply to graph's default node attributes
		for _, opt := range opts {
			opt.applyNode(g.DefaultNodeAttrs())
		}
	case "edge":
		opts := p.mapEdgeAttributes(attrs)
		// Apply to graph's default edge attributes
		for _, opt := range opts {
			opt.applyEdge(g.DefaultEdgeAttrs())
		}
	case "graph":
		// Apply graph attributes
		for key, value := range attrs {
			g.Attrs().setCustom(key, value)
		}
	}
	return nil
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

// parseSubgraph parses a subgraph statement.
// Syntax: [subgraph [ID]] { stmt_list }
// Returns the created subgraph.
func (p *Parser) parseSubgraph(g *Graph) (*Subgraph, error) {
	// Parse 'subgraph' keyword if present
	if p.matchKeyword("subgraph") {
		p.advance()
	}

	// Parse optional subgraph name
	var name string
	if p.current.Type == TokenIdent || p.current.Type == TokenString {
		name = p.current.Value
		p.advance()
	}

	// Expect opening brace
	if err := p.expect(TokenLBrace); err != nil {
		return nil, err
	}

	// Parse subgraph contents and create subgraph
	var parseErr error
	sg := g.Subgraph(name, func(s *Subgraph) {
		// Parse statements into this subgraph
		parseErr = p.parseSubgraphStmts(s, g)
	})

	if parseErr != nil {
		return nil, parseErr
	}

	// Expect closing brace
	if err := p.expect(TokenRBrace); err != nil {
		return nil, err
	}

	return sg, nil
}

// parseSubgraphStmts parses a list of statements within a subgraph.
// Nodes and edges are added to the subgraph, while default attributes affect the parent graph.
func (p *Parser) parseSubgraphStmts(sg *Subgraph, g *Graph) error {
	for !p.match(TokenRBrace) && !p.match(TokenEOF) {
		// Skip optional semicolons
		if p.match(TokenSemi) {
			p.advance()
			continue
		}

		// Parse statement in subgraph context
		if err := p.parseSubgraphStmt(sg, g); err != nil {
			return err
		}

		// Skip optional trailing semicolon
		if p.match(TokenSemi) {
			p.advance()
		}
	}
	return nil
}

// parseSubgraphStmt parses a single statement within a subgraph context.
func (p *Parser) parseSubgraphStmt(sg *Subgraph, g *Graph) error {
	// Check for keywords: node, edge, graph
	if p.matchKeyword("node") || p.matchKeyword("edge") || p.matchKeyword("graph") {
		// Default attribute statements - these affect the parent graph
		keyword := p.current.Value
		p.advance()
		if p.match(TokenLBracket) {
			attrs, err := p.parseAttrList()
			if err != nil {
				return err
			}
			// Apply default attributes to parent graph
			return p.applyDefaultAttrs(g, keyword, attrs)
		}
		return nil
	}

	// Check for nested subgraph
	if p.matchKeyword("subgraph") {
		_, err := p.parseNestedSubgraph(sg)
		return err
	}

	// Check for bare '{' (anonymous nested subgraph)
	if p.match(TokenLBrace) {
		_, err := p.parseNestedSubgraph(sg)
		return err
	}

	// Parse node or edge statement into this subgraph
	return p.parseNodeOrEdgeStmtInSubgraph(sg)
}

// parseNestedSubgraph parses a nested subgraph within a parent subgraph.
func (p *Parser) parseNestedSubgraph(parent *Subgraph) (*Subgraph, error) {
	// Parse 'subgraph' keyword if present
	if p.matchKeyword("subgraph") {
		p.advance()
	}

	// Parse optional subgraph name
	var name string
	if p.current.Type == TokenIdent || p.current.Type == TokenString {
		name = p.current.Value
		p.advance()
	}

	// Expect opening brace
	if err := p.expect(TokenLBrace); err != nil {
		return nil, err
	}

	// Parse nested subgraph contents
	var parseErr error
	sg := parent.Subgraph(name, func(s *Subgraph) {
		parseErr = p.parseSubgraphStmts(s, parent.parent)
	})

	if parseErr != nil {
		return nil, parseErr
	}

	// Expect closing brace
	if err := p.expect(TokenRBrace); err != nil {
		return nil, err
	}

	return sg, nil
}

// parseNodeOrEdgeStmtInSubgraph parses a node or edge statement and adds it to the subgraph.
func (p *Parser) parseNodeOrEdgeStmtInSubgraph(sg *Subgraph) error {
	// Parse first ID
	id, err := p.parseID()
	if err != nil {
		return err
	}

	// Check if this is an edge statement (next token is arrow)
	if p.match(TokenArrow) {
		return p.parseEdgeStmtInSubgraph(sg, id)
	}

	// Otherwise, it's a node statement
	return p.parseNodeStmtInSubgraph(sg, id)
}

// parseNodeStmtInSubgraph parses a node statement and adds it to the subgraph.
func (p *Parser) parseNodeStmtInSubgraph(sg *Subgraph, id string) error {
	// Parse optional attributes
	var attrs map[string]string
	if p.match(TokenLBracket) {
		var err error
		attrs, err = p.parseAttrList()
		if err != nil {
			return err
		}
	}

	// Create node with attributes
	nodeOpts := p.mapNodeAttributes(attrs)
	node := NewNode(id, nodeOpts...)
	return sg.AddNode(node)
}

// parseEdgeStmtInSubgraph parses an edge statement and adds it to the subgraph.
func (p *Parser) parseEdgeStmtInSubgraph(sg *Subgraph, firstID string) error {
	nodes := []string{firstID}

	// Parse edge chain
	for p.match(TokenArrow) {
		p.advance() // consume arrow

		// Parse next node ID
		id, err := p.parseID()
		if err != nil {
			return err
		}
		nodes = append(nodes, id)
	}

	// Parse optional edge attributes
	var attrs map[string]string
	if p.match(TokenLBracket) {
		var err error
		attrs, err = p.parseAttrList()
		if err != nil {
			return err
		}
	}

	// Create edges for the chain
	edgeOpts := p.mapEdgeAttributes(attrs)
	for i := 0; i < len(nodes)-1; i++ {
		from := NewNode(nodes[i])
		to := NewNode(nodes[i+1])
		if _, err := sg.AddEdge(from, to, edgeOpts...); err != nil {
			return err
		}
	}

	return nil
}

// skipStatement skips tokens until we hit a semicolon or statement boundary.
func (p *Parser) skipStatement() error {
	// Skip tokens until semicolon, closing brace, or EOF
	for !p.match(TokenSemi) && !p.match(TokenRBrace) && !p.match(TokenEOF) {
		// Handle nested braces and brackets
		if p.match(TokenLBrace) {
			if err := p.skipBraceBlock(); err != nil {
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

// skipBraceBlock skips over a block enclosed in braces { ... }.
func (p *Parser) skipBraceBlock() error {
	if err := p.expect(TokenLBrace); err != nil {
		return err
	}

	depth := 1
	for depth > 0 && !p.match(TokenEOF) {
		if p.match(TokenLBrace) {
			depth++
		} else if p.match(TokenRBrace) {
			depth--
		}
		if depth > 0 {
			p.advance()
		}
	}

	if err := p.expect(TokenRBrace); err != nil {
		return err
	}

	return nil
}

// ParseError represents an error that occurred during parsing with location information.
type ParseError struct {
	Message string // Error message
	Line    int    // Line number where error occurred (1-based)
	Col     int    // Column number where error occurred (1-based)
	Snippet string // Surrounding context from the input
}

// Error implements the error interface.
func (e *ParseError) Error() string {
	if e.Line > 0 && e.Col > 0 {
		return fmt.Sprintf("parse error at %d:%d: %s", e.Line, e.Col, e.Message)
	}
	return fmt.Sprintf("parse error: %s", e.Message)
}

// wrapParseError wraps an error with location information from the current token.
func (p *Parser) wrapParseError(err error) *ParseError {
	if err == nil {
		return nil
	}

	// If it's already a ParseError, return it as-is
	if perr, ok := err.(*ParseError); ok {
		return perr
	}

	// Extract a snippet from the input around the error location
	snippet := p.extractSnippet()

	return &ParseError{
		Message: err.Error(),
		Line:    p.current.Line,
		Col:     p.current.Col,
		Snippet: snippet,
	}
}

// extractSnippet extracts a few lines of context around the current token.
func (p *Parser) extractSnippet() string {
	if p.lexer == nil || p.lexer.input == "" {
		return ""
	}

	lines := strings.Split(p.lexer.input, "\n")
	if len(lines) == 0 || p.current.Line < 1 || p.current.Line > len(lines) {
		return ""
	}

	// Get the current line (1-based indexing)
	lineIdx := p.current.Line - 1
	currentLine := lines[lineIdx]

	// Create a pointer to the error location
	pointer := strings.Repeat(" ", p.current.Col-1) + "^"

	return fmt.Sprintf("%s\n%s", currentLine, pointer)
}

// ParseString parses a DOT graph from a string.
// Returns the parsed Graph or a ParseError if parsing fails.
//
// Example:
//
//	g, err := goraffe.ParseString("digraph { A -> B; }")
//	if err != nil {
//	    log.Fatal(err)
//	}
func ParseString(dot string) (*Graph, error) {
	parser := newParser(dot)
	g, err := parser.parseGraph()
	if err != nil {
		return nil, parser.wrapParseError(err)
	}
	return g, nil
}

// Parse parses a DOT graph from an io.Reader.
// Returns the parsed Graph or a ParseError if parsing fails.
//
// Example:
//
//	file, _ := os.Open("graph.dot")
//	defer file.Close()
//	g, err := goraffe.Parse(file)
//	if err != nil {
//	    log.Fatal(err)
//	}
func Parse(r io.Reader) (*Graph, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, &ParseError{
			Message: fmt.Sprintf("failed to read input: %v", err),
		}
	}
	return ParseString(string(data))
}

// ParseFile parses a DOT graph from a file.
// Returns the parsed Graph or a ParseError if parsing fails.
//
// Example:
//
//	g, err := goraffe.ParseFile("graph.dot")
//	if err != nil {
//	    log.Fatal(err)
//	}
func ParseFile(path string) (*Graph, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, &ParseError{
			Message: fmt.Sprintf("failed to open file: %v", err),
		}
	}
	defer file.Close()

	return Parse(file)
}
