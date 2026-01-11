package goraffe

import "strings"

// escapeDOTString escapes special characters in a string for DOT output.
// It escapes:
//   - Backslashes: \ → \\
//   - Double quotes: " → \"
//   - Newlines: \n (actual newline character) → \n (literal backslash-n in output)
//
// This function should be used for all string values in DOT output (labels, attribute values, etc.)
func escapeDOTString(s string) string {
	// Replace backslashes first (must be done before quotes to avoid double-escaping)
	s = strings.ReplaceAll(s, `\`, `\\`)
	// Replace double quotes
	s = strings.ReplaceAll(s, `"`, `\"`)
	// Replace actual newline characters with literal \n
	s = strings.ReplaceAll(s, "\n", `\n`)
	return s
}

// quoteDOTID quotes and escapes a DOT identifier (node ID or graph name).
// Per DOT specification and for safety, we always quote identifiers.
// The string is escaped using escapeDOTString before being quoted.
func quoteDOTID(s string) string {
	return `"` + escapeDOTString(s) + `"`
}
