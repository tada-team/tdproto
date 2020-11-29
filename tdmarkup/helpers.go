package tdmarkup

func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' }

func isEOL(ch rune) bool { return ch == '\r' || ch == '\n' }

func isEOF(ch rune) bool { return ch == EOF }

func isTrailingPunctuation(ch rune) bool {
	return ch == '.' || ch == ',' || ch == ':' || ch == ';' || ch == '!' || ch == '?'
}
