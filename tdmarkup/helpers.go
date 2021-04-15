package tdmarkup

func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' }

func isEOL(ch rune) bool { return ch == '\r' || ch == '\n' }

func isEOF(ch rune) bool { return ch == EOF }

func isTrailingPunctuation(ch rune) bool {
	return ch == '.' || ch == ',' || ch == ':' || ch == ';' || ch == '!' || ch == '?'
}

func isPath(s string) bool {
	if len(s) < 3 || s[0] != '/' || s[len(s) -1] != '/' {
		return false
	}
	hasSlash := false
	for _, ch := range s[1:len(s) -2] {
		if ch == '/' {
			hasSlash = true
		}
		if isWhitespace(ch) {
			return false
		}
	}
	return hasSlash
}
