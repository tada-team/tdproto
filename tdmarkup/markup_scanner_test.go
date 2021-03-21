package tdmarkup

import (
	"regexp"
	"testing"
)

func TestToPlainNoQuotes(t *testing.T) {
	for k, v := range map[string]struct {
		raw   string
		plain string
	}{
		"noop": {
			raw:   "123",
			plain: "123",
		},
		"quote": {
			raw:   "> 123\n456",
			plain: "456",
		},
	} {
		t.Run(k, func(t *testing.T) {
			s, entities := ParseString(v.raw, nil)
			if s != v.raw {
				t.Fatalf("invalid ParseString text.\n got: '%s'\nwant: '%s'", s, v.raw)
			}
			if text := ToPlain(s, entities, &ToPlainOpts{
				DisableQuotes: true,
			}); text != v.plain {
				t.Errorf(
					"invalid ToPlainNoQuotes() output\n raw: '%s'\n got: '%s'\nwant: '%s'\n ent: %+v",
					v.raw, text, v.plain, entities,
				)
			}
		})
	}
}
func TestParse(t *testing.T) {
	for _, v := range MarkupTestCases {
		t.Run(v.Title, func(t *testing.T) {
			s, entities := ParseString(v.Raw, v.Links)
			if s != v.Raw {
				t.Fatalf("invalid ParseString text.\n got: '%s'\nwant: '%s'", s, v.Raw)
			}

			if v.Html != "" {
				if text := ToHtml(s, entities); text != v.Html {
					t.Errorf(
						"invalid ToHtml() output\n raw: '%s'\n got: '%s'\nwant: '%s'\n ent: %+v",
						v.Raw, text, v.Html, entities,
					)
				}
			}

			if v.Plain != "" {
				if text := ToPlain(s, entities, nil); text != v.Plain {
					t.Errorf(
						"invalid ToPlain() output\n raw: '%s'\n got: '%s'\nwant: '%s'\n ent: %+v",
						v.Raw, text, v.Plain, entities,
					)
				}
			}
		})
	}
}

func TestMustTime(t *testing.T) {
	mustTime("2020-10-27T12:24:09.781121Z")
	mustTime("2000-01-02T17:15:00.000000Z")
	mustTime("2006-01-02T15:04:05.000000-0700")
}

func TestContainsTime(t *testing.T) {
	for k, v := range map[string]struct {
		text     string
		contains bool
	}{
		"base format": {
			text:     "<2020-10-27T12:24:09.781121Z>",
			contains: true,
		},
		"tada format": {
			text:     "<2020-10-02T15:04:05.000000-0700>",
			contains: true,
		},
		"into chars": {
			text:     "*<2000-01-02T17:15:00.000000Z>*",
			contains: true,
		},
		"simple string": {
			text:     "asdasdasdasd",
			contains: false,
		},
		"simple string with quotes": {
			text:     "<asdasdasdasd>",
			contains: false,
		},
		"between words": {
			text:     "hello <2020-10-02T15:04:05.000000-0700> world",
			contains: true,
		},
	} {
		t.Run(k, func(t *testing.T) {
			if contains := ContainsTime(v.text); contains != v.contains {
				t.Error("bad ContainsTime result. want:", v.contains, "got:", contains, "text:", v.text)
			}
		})
	}
}

// BenchmarkParse/regex-12                 	  107926	     11367 ns/op	    1129 B/op	      34 allocs/op
//
// BenchmarkParse/markupScanner-12         	  238516	      4882 ns/op	    2600 B/op	      71 allocs/op
// => strings.Builder.Grow
// BenchmarkParse/markupScanner-12         	  323514	      3734 ns/op	    2584 B/op	      68 allocs/op
func BenchmarkParse(b *testing.B) {
	raw := "*bold* /italic/ 1234567 _3_ `code`"
	plain := "bold italic 1234567 3 code"

	b.Run("markupScanner", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s, entities := ParseString(raw, nil)
			res := ToPlain(s, entities, nil)
			if res != plain {
				b.Fatalf("invalid plain, got: %s, want: %s", res, plain)
			}
		}
	})

	b.Run("regex", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			res := clearFormat(raw)
			if res != plain {
				b.Fatalf("invalid plain, got: %s, want: %s", res, plain)
			}
		}
	})
}

func mask(s string) *regexp.Regexp {
	return regexp.MustCompile(`(^|\s)` + s + `([^\s` + s + `].*?[^\s` + s + `]|[^\s` + s + `])` + s + `(\s|$|\?|:|\.|,|!)`)
}

var formatRules = []*regexp.Regexp{
	regexp.MustCompile("(?ms)(^|\\s)```(.+?)```" + `(\s|$|\?|:|\.|,|!)`),
	mask("`"),
	mask(`\*`),
	mask("/"),
	mask("_"),
	mask("~"),
}

func clearFormat(text string) string {
	for _, regex := range formatRules {
		text = regex.ReplaceAllStringFunc(text, func(m string) string {
			parts := regex.FindStringSubmatch(m)
			return parts[1] + parts[2] + parts[3]
		})
	}
	return text
}
