package tdmarkup

import (
	"strings"
	"time"

	"github.com/tada-team/tdproto"
)

var ops = "*/_~`<>&"

var opInlines = map[rune]tdproto.MarkupType{
	'*': tdproto.Bold,
	'/': tdproto.Italic,
	'_': tdproto.Underscore,
	'~': tdproto.Strike,
	'`': tdproto.Code,
}

var (
	opCodeBlock  = []rune("```")
	opQuoteBlock = []rune("> ")
)

func contains(s string, typ tdproto.MarkupType) bool {
	for s := NewMarkupScanner(s); s.Rest() > 0; {
		_, e := s.Scan(nil)
		if doContains(e, typ) {
			return true
		}
	}
	return false
}

func doContains(e *tdproto.MarkupEntity, substring tdproto.MarkupType) bool {
	if e == nil {
		return false
	}
	if e.Type == substring {
		return true
	}
	for _, child := range e.Childs {
		if doContains(&child, substring) {
			return true
		}
	}
	return false
}

func ContainsTime(s string) bool { return contains(s, tdproto.Time) }

func ParseString(text string, links tdproto.MessageLinks) (string, []tdproto.MarkupEntity) {
	text = strings.ReplaceAll(text, "\r", "")
	if len(links) == 0 && !strings.ContainsAny(text, ops) {
		return text, nil
	}

	var b strings.Builder
	b.Grow(len(text))

	var entities []tdproto.MarkupEntity
	for s := NewMarkupScanner(text); s.Rest() > 0; {
		t, e := s.Scan(links)
		if e != nil {
			entities = append(entities, *e)
		}
		b.WriteString(t)
	}

	return b.String(), entities
}

type MarkupScanner struct {
	*Scanner
	internal bool
}

func NewMarkupScanner(text string) *MarkupScanner {
	return &MarkupScanner{Scanner: NewScanner(text)}
}

// TODO: markdown link
func (s *MarkupScanner) Scan(links tdproto.MessageLinks) (string, *tdproto.MarkupEntity) {
	if isEOF(s.Next()) {
		return "", nil
	}

	// links
	for _, l := range links {
		t, e := s.scanLink(l)
		if e != nil {
			return t, e
		}
	}

	// dates (before html tags!)
	t, e := s.scanTime()
	if e != nil {
		return t, e
	}

	// quotes (before html tags!)
	if s.Position() == 0 || isEOL(s.Current()) {
		t, e := s.scanQuote()
		if e != nil {
			e.Childs = s.scanChilds(t[len(opQuoteBlock):])
			return t, e
		}
		if t != "" {
			return t, nil
		}
	}

	// html tags
	t, e = s.scanUnsafe()
	if e != nil {
		return t, e
	}

	// code block
	t, e = s.scanBlock(opCodeBlock, opCodeBlock, tdproto.CodeBlock)
	if e != nil {
		return t, e
	}
	if t != "" {
		return t, nil
	}

	// inlines
	if typ, ok := opInlines[s.Next()]; ok {
		allowWhitespaceAround := typ == tdproto.Code
		t, e := s.scanInline(s.Next(), typ, allowWhitespaceAround)
		if e != nil {
			if typ == tdproto.Italic && isPath(t) {
				return t, nil
			}
			if typ != tdproto.Code {
				e.Childs = s.scanChilds(t[1 : len(t)-1])
			}
			return t, e
		}
		if t != "" {
			return t, nil
		}
	}

	// unparsed
	return string(s.TakeNext()), nil
}

func (s *MarkupScanner) scanChilds(text string) (res []tdproto.MarkupEntity) {
	if len(text) < 3 {
		return
	}
	scanner := NewMarkupScanner(text)
	scanner.internal = true
	for scanner.Rest() > 0 {
		t, e := scanner.Scan(nil)
		if e != nil {
			if res == nil {
				res = make([]tdproto.MarkupEntity, 0)
			}
			res = append(res, *e)
		}
		if t == "" {
			break
		}
	}
	return
}

var dateLayouts = []string{
	"2006-01-02T15:04:05.000000-0700",
	"2006-01-02T15:04:05.000000Z",
}

func (s *MarkupScanner) scanTime() (string, *tdproto.MarkupEntity) {
	if s.Next() != '<' {
		return "", nil
	}

	start := s.Position()
	s.TakeNext()
	strDt := s.ScanUntil([]rune(">"))
	if len(strDt) < 18 {
		s.Rewind(start)
		return "", nil
	}

	strDt = strDt[:len(strDt)-1]
	for _, layout := range dateLayouts {
		_, err := time.Parse(layout, strDt)
		if err != nil {
			continue
		}
		return "<" + strDt + ">", &tdproto.MarkupEntity{
			Open:        start,
			OpenLength:  1,
			Close:       s.Position() - 1,
			CloseLength: 1,
			Type:        tdproto.Time,
			Time:        strDt,
		}
	}

	s.Rewind(start)
	return "", nil

}

func (s *MarkupScanner) scanInline(marker rune, typ tdproto.MarkupType, allowWhitespaceAround bool) (string, *tdproto.MarkupEntity) {
	start := s.Position()

	var b strings.Builder
	b.Grow(s.Length() - start)
	b.WriteRune(s.TakeNext())

	if !(start == 0 || isWhitespace(s.Prev()) || isEOL(s.Prev()) || allowWhitespaceAround) {
		s.Rewind(start)
		return "", nil
	}

	if (isWhitespace(s.Next()) || isEOL(s.Next())) && !allowWhitespaceAround {
		s.Rewind(start)
		return "", nil
	}

	e := &tdproto.MarkupEntity{
		Type:       typ,
		Open:       s.Position() - 1,
		OpenLength: 1,
	}

	for s.Rest() > 0 {
		ch := s.TakeNext()
		b.WriteRune(ch)

		prev := s.Prev()
		next := s.Next()

		if ch == marker && s.Since(start) > 2 && next != marker && (allowWhitespaceAround || !isWhitespace(prev)) && prev != ch &&
			(isWhitespace(next) || isEOF(next) || isEOL(next) || isTrailingPunctuation(next)) {
			e.Close = s.Position() - 1
			e.CloseLength = 1
			return b.String(), e
		}

		if isEOL(ch) {
			break
		}
	}

	s.Rewind(start)
	return "", nil
}

func (s *MarkupScanner) scanBlock(op, cl []rune, typ tdproto.MarkupType) (string, *tdproto.MarkupEntity) {
	start := s.Position()

	t := s.ScanUntil(op)
	if t == "" {
		return "", nil
	}

	var b strings.Builder
	b.Grow(s.Length() - start)
	b.WriteString(t)

	e := &tdproto.MarkupEntity{
		Type:       typ,
		Open:       s.Position() - len(op),
		OpenLength: len(op),
	}

	for s.Next() == ' ' {
		e.OpenLength += 1
		b.WriteRune(s.TakeNext())
	}

	for s.Next() == '\n' {
		e.OpenLength += 1
		b.WriteRune(s.TakeNext())
	}

	var tail []rune
	for s.Rest() > 0 {
		t := s.ScanUntil(cl)
		if t == "" {
			ch := s.TakeNext()
			b.WriteRune(ch)
			tail = append(tail, ch)
			continue
		}
		b.WriteString(t)
		e.Close = s.Position() - len(cl)
		e.CloseLength = len(cl)

		for i := len(tail) - 1; i >= 0; i-- {
			ch := tail[i]
			if !(isWhitespace(ch) || isEOL(ch)) {
				break
			}
			e.Close--
			e.CloseLength++
		}
		return b.String(), e
	}

	s.Rewind(start)
	return "", nil
}

func (s *MarkupScanner) scanQuote() (string, *tdproto.MarkupEntity) {
	t := s.ScanUntil(opQuoteBlock)
	if t == "" {
		return "", nil
	}

	var b strings.Builder
	b.Grow(s.Length() - s.Position())
	b.WriteString(t)

	e := &tdproto.MarkupEntity{
		Type:       tdproto.Quote,
		Open:       s.Position() - len(opQuoteBlock),
		OpenLength: len(opQuoteBlock),
	}

	for {
		ch := s.Next()
		if isEOL(ch) || isEOF(ch) {
			e.Close = s.Position()
			if isEOL(ch) {
				e.CloseLength = 1
			}
			return b.String(), e
		}
		b.WriteRune(s.TakeNext())
	}
}

func (s *MarkupScanner) scanLink(l tdproto.MessageLink) (string, *tdproto.MarkupEntity) {
	start := s.Position()

	for _, r := range []rune(l.Pattern) {
		if s.TakeNext() != r {
			s.Rewind(start)
			return "", nil
		}
	}

	return l.Pattern, &tdproto.MarkupEntity{
		Type:  tdproto.Link,
		Url:   l.Url,
		Repl:  l.Text,
		Open:  start,
		Close: s.Position(),
	}
}

func (s *MarkupScanner) scanUnsafe() (string, *tdproto.MarkupEntity) {
	switch s.Next() {
	case '&', '<', '>':
		start := s.Position()
		return string(s.TakeNext()), &tdproto.MarkupEntity{
			Open:  start,
			Close: start + 1,
			Type:  tdproto.Unsafe,
		}
	default:
		return string(s.Next()), nil
	}
}
