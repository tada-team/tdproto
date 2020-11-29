package tdmarkup

import (
	"fmt"
	"html"
	"log"
	"strings"
	"time"

	"github.com/tada-team/tdproto"
)

type Rules func(middle []rune, e tdproto.MarkupEntity) ([]rune, []rune, []rune)

const timeFormat = "02.01.06 15:04"

type ToPlainOpts struct {
	DisableQuotes bool
	Location      *time.Location
}

func ToHtml(s string, e []tdproto.MarkupEntity) string {
	return string(conv([]rune(s), e, func(middle []rune, e tdproto.MarkupEntity) ([]rune, []rune, []rune) {
		switch e.Type {
		case tdproto.Bold:
			return []rune("<b>"), middle, []rune("</b>")
		case tdproto.Italic:
			return []rune("<i>"), middle, []rune("</i>")
		case tdproto.Underscore:
			return []rune("<u>"), middle, []rune("</u>")
		case tdproto.Code:
			return []rune("<code>"), middle, []rune("</code>")
		case tdproto.Strike:
			return []rune("<s>"), middle, []rune("</s>")
		case tdproto.CodeBlock:
			return []rune("<pre>"), middle, []rune("</pre>")
		case tdproto.Quote:
			return []rune("<blockquote>"), middle, []rune("</blockquote>\n")
		case tdproto.Link:
			return []rune(fmt.Sprintf(`<a href="%s">`, e.Url)), []rune(e.Repl), []rune("</a>")
		case tdproto.Time:
			return []rune("<time>"), []rune(mustTime(e.Time).Format(timeFormat)), []rune("</time>")
		case tdproto.Unsafe:
			return []rune(""), []rune(html.EscapeString(string(middle))), []rune("")
		default:
			return []rune(""), middle, []rune("")
		}
	}))
}

func ToPlain(s string, e []tdproto.MarkupEntity, opts *ToPlainOpts) string {
	if opts == nil {
		opts = new(ToPlainOpts)
	}

	row := string(conv([]rune(s), e, func(middle []rune, e tdproto.MarkupEntity) ([]rune, []rune, []rune) {
		switch e.Type {
		case tdproto.Quote:
			if opts.DisableQuotes {
				return []rune(""), []rune(""), []rune("")
			}
			return []rune("> "), middle, []rune("\n")
		case tdproto.Link:
			for _, prefix := range [...]string{
				"tadateam://",
				"tel:",
				"mailto:",
			} {
				if strings.HasPrefix(e.Url, prefix) {
					return []rune(""), []rune(e.Repl), []rune("")
				}
			}
			return []rune(""), []rune(e.Url), []rune("")
		case tdproto.Time:
			t := mustTime(e.Time)
			if loc := opts.Location; loc != nil {
				t = t.In(loc)
			}
			return []rune(""), []rune(t.Format(timeFormat)), []rune("")
		default:
			return []rune(""), middle, []rune("")
		}
	}))
	if !opts.DisableQuotes {
		return strings.TrimRight(row, "\n\r")
	}
	return row
}

func conv(runes []rune, entities []tdproto.MarkupEntity, rules Rules) []rune {
	offset := 0
	for _, e := range entities {
		middleDiff := 0
		middle := runes[e.Open+e.OpenLength+offset : e.Close+offset]
		op, rulesMiddle, cl := rules(middle, e)

		if len(e.Childs) > 0 {
			entitiesMiddle := conv(middle, e.Childs, rules)
			middleDiff = len(entitiesMiddle) - len(middle)
			middle = entitiesMiddle
		} else {
			middleDiff = len(rulesMiddle) - len(middle)
			middle = rulesMiddle
		}

		res := make([]rune, 0, len(runes))
		res = append(res, runes[:e.Open+offset]...)
		res = append(res, op...)
		res = append(res, middle...)
		res = append(res, cl...)
		res = append(res, runes[e.Close+e.CloseLength+offset:]...)

		runes = res
		offset += len(op) + len(cl) - e.OpenLength - e.CloseLength + middleDiff
	}
	return runes
}

func mustTime(s string) time.Time {
	dt, err := time.Parse("2006-01-02T15:04:05Z0700", s)
	if err != nil {
		log.Panicln("invalid date:", s, err)
	}
	return dt
}
