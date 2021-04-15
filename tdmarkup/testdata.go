package tdmarkup

import "github.com/tada-team/tdproto"

// Data for testing other implementation
var MarkupTestCases = []struct {
	Title string
	Raw   string
	Html  string
	Plain string
	Links []tdproto.MessageLink
}{
	{
		Title: "noop",
		Raw:   "123",
		Html:  "123",
		Plain: "123",
	},
	{
		Title: "empty bold",
		Raw:   "**",
		Html:  "**",
		Plain: "**",
	},
	{
		Title: "one bold",
		Raw:   "*1*",
		Html:  "<b>1</b>",
		Plain: "1",
	},
	{
		Title: "bold",
		Raw:   "123 *456* 789",
		Html:  "123 <b>456</b> 789",
		Plain: "123 456 789",
	},
	{
		Title: "bold is inline only",
		Raw:   "123 *45\n6* 789",
		Html:  "123 *45\n6* 789",
		Plain: "123 *45\n6* 789",
	},
	{
		Title: "bold ru",
		Raw:   "привет, *ромашки!*",
		Html:  "привет, <b>ромашки!</b>",
		Plain: "привет, ромашки!",
	},
	{
		Title: "bold only",
		Raw:   "*123*",
		Html:  "<b>123</b>",
		Plain: "123",
	},
	{
		Title: "bold and punctuation",
		Raw:   "*123*!",
		Html:  "<b>123</b>!",
		Plain: "123!",
	},
	{
		Title: "not a bold",
		Raw:   "*123\n456*",
		Html:  "*123\n456*",
		Plain: "*123\n456*",
	},
	{
		Title: "bold eof",
		Raw:   "123 *456*",
		Html:  "123 <b>456</b>",
		Plain: "123 456",
	},
	{
		Title: "bold eol",
		Raw:   "123 *456*\n*789*",
		Html:  "123 <b>456</b>\n<b>789</b>",
		Plain: "123 456\n789",
	},
	{
		Title: "double bold",
		Raw:   "123 *456* 789 *10*",
		Html:  "123 <b>456</b> 789 <b>10</b>",
		Plain: "123 456 789 10",
	},
	{
		Title: "code only",
		Raw:   "`123`",
		Html:  "<code>123</code>",
		Plain: "123",
	},
	{
		Title: "no markup inside code",
		Raw:   "`*123*`",
		Html:  "<code>*123*</code>",
		Plain: "*123*",
	},
	{
		Title: "no markup inside code 2",
		Raw:   "`123 ~456~ `",
		Html:  "<code>123 ~456~ </code>",
		Plain: "123 ~456~ ",
	},
	{
		Title: "no markup inside code 3",
		Raw:   "example: ` 123 ~456~ `",
		Html:  "example: <code> 123 ~456~ </code>",
		Plain: "example:  123 ~456~ ",
	},
	{
		Title: "italic",
		Raw:   "123 /456/ 789",
		Html:  "123 <i>456</i> 789",
		Plain: "123 456 789",
	},
	{
		Title: "not italic",
		Raw:   "123 / 456/ 789",
		Html:  "123 / 456/ 789",
		Plain: "123 / 456/ 789",
	},
	{
		Title: "not italic 2",
		Raw:   "123 / 456 / 789",
		Html:  "123 / 456 / 789",
		Plain: "123 / 456 / 789",
	},
	{
		Title: "not italic 3",
		Raw:   "123/ 456 /789",
		Html:  "123/ 456 /789",
		Plain: "123/ 456 /789",
	},
	{
		Title: "path",
		Raw:   "GET /path/to/ here",
		Html:  "GET /path/to/ here",
		Plain: "GET /path/to/ here",
	},
	{
		Title: "not italic 4 without links",
		Raw:   "https://ya.ru/",
		Html:  "https://ya.ru/",
		Plain: "https://ya.ru/",
	},
	{
		Title: "not italic 4 with links",
		Raw:   "https://ya.ru/",
		Html:  `<a href="https://ya.ru">ya.ru</a>`,
		Plain: "ya.ru",
		Links: []tdproto.MessageLink{
			{
				Pattern: "https://ya.ru/",
				Url:     "https://ya.ru",
				Text:    "ya.ru",
			},
		},
	},
	{
		Title: "bold italic",
		Raw:   "*/123/*",
		Html:  "<b><i>123</i></b>",
		Plain: "123",
	},
	{
		Title: "code block",
		Raw:   "```code```",
		Html:  "<pre>code</pre>",
		Plain: "code",
	},
	{
		Title: "code block multiline",
		Raw:   "```code\nnl```",
		Html:  "<pre>code\nnl</pre>",
		Plain: "code\nnl",
	},
	{
		Title: "code block with newlines",
		Raw:   "```\ncode\nline\n```",
		Html:  "<pre>code\nline</pre>",
		Plain: "code\nline",
	},
	{
		Title: "code block with newlines with spaces",
		Raw:   "``` \ncode\nline\n```",
		Html:  "<pre>code\nline</pre>",
		Plain: "code\nline",
	},
	{
		Title: "code incomplete",
		Raw:   "```code\nnl``",
		Html:  "```code\nnl``",
		Plain: "```code\nnl``",
	},
	{
		Title: "bold italic with spaces",
		Raw:   "123 /4 *5* 6/ 789",
		Html:  "123 <i>4 <b>5</b> 6</i> 789",
		Plain: "123 4 5 6 789",
	},
	{
		Title: "quote",
		Raw:   "> 123\n456",
		Html:  "<blockquote>123</blockquote>\n456",
		Plain: "> 123\n456",
	},
	{
		Title: "quote only",
		Raw:   "> 123",
		Html:  "<blockquote>123</blockquote>\n",
		Plain: "> 123",
	},
	{
		Title: "bold in quote",
		Raw:   "> 12 *3*\n456",
		Html:  "<blockquote>12 <b>3</b></blockquote>\n456",
		Plain: "> 12 3\n456",
	},
	{
		Title: "quote in quote",
		Raw:   "> 3 > 4",
		Html:  "<blockquote>3 &gt; 4</blockquote>\n",
		Plain: "> 3 > 4",
	},
	{
		Title: "not a quote",
		Raw:   "3 > 4",
		Html:  "3 &gt; 4",
		Plain: "3 > 4",
	},
	{
		Title: "quote from new line",
		Raw:   "\n> b",
		Html:  "\n<blockquote>b</blockquote>\n",
	},
	{
		Title: "multiquotes",
		Raw: `> 1
2
> 3
4
> 5
6`,
		Html: `<blockquote>1</blockquote>
2
<blockquote>3</blockquote>
4
<blockquote>5</blockquote>
6`,
	},
	{
		Title: "link only",
		Raw:   "123 https://ya.ru 456",
		Html:  `123 <a href="https://ya.ru">ya.ru</a> 456`,
		Plain: `123 ya.ru 456`,
		Links: []tdproto.MessageLink{
			{
				Pattern: "https://ya.ru",
				Url:     "https://ya.ru",
				Text:    "ya.ru",
			},
		},
	},
	{
		Title: "mailto",
		Raw:   "123 xxx@gmail.com 456",
		Html:  `123 <a href="mailto:xxx@gmail.com">xxx@gmail.com</a> 456`,
		Plain: `123 xxx@gmail.com 456`,
		Links: []tdproto.MessageLink{
			{
				Pattern: "xxx@gmail.com",
				Url:     "mailto:xxx@gmail.com",
				Text:    "xxx@gmail.com",
			},
		},
	},
	{
		Title: "link only with trailing slash",
		Raw:   "123 https://ya.ru/? 456",
		Html:  `123 <a href="https://ya.ru">ya.ru</a> 456`,
		Plain: "123 ya.ru 456",
		Links: []tdproto.MessageLink{
			{
				Pattern: "https://ya.ru/?",
				Url:     "https://ya.ru",
				Text:    "ya.ru",
			},
		},
	},
	{
		Title: "manage.py",
		Raw:   "123 manage.py 456",
		Html:  `123 <a href="http://manage.py">manage.py</a> 456`,
		Plain: "123 manage.py 456",
		Links: []tdproto.MessageLink{
			{
				Pattern: "manage.py",
				Url:     "http://manage.py",
				Text:    "manage.py",
			},
		},
	},
	{
		Title: "tag",
		Raw:   "<br>",
		Html:  "&lt;br&gt;",
		Plain: "<br>",
	},
	{
		Title: "lt",
		Raw:   "<<<<<<",
		Html:  "&lt;&lt;&lt;&lt;&lt;&lt;",
		Plain: "<<<<<<",
	},
	{
		Title: "bolditalic tag",
		Raw:   "<b><i>123</i></b>",
		Html:  "&lt;b&gt;&lt;i&gt;123&lt;/i&gt;&lt;/b&gt;",
		Plain: "<b><i>123</i></b>",
	},
	{
		Title: "bold italic without markup",
		Raw:   "<b>1 <i>2</i> 3</b>",
		Html:  "&lt;b&gt;1 &lt;i&gt;2&lt;/i&gt; 3&lt;/b&gt;",
		Plain: "<b>1 <i>2</i> 3</b>",
	},
	{
		Title: "tag after markup",
		Raw:   "/*6688*/ 3<br>",
		Html:  "<i><b>6688</b></i> 3&lt;br&gt;",
		Plain: "6688 3<br>",
	},
	{
		Title: "bold italic with markup",
		Raw:   "<b>1 <i>2</i> /*6688*/ 3</b>",
		Html:  "&lt;b&gt;1 &lt;i&gt;2&lt;/i&gt; <i><b>6688</b></i> 3&lt;/b&gt;",
		Plain: "<b>1 <i>2</i> 6688 3</b>",
	},
	{
		Title: "nested markup no spaces",
		Raw:   `_1-*2-~34~-6*-7_`,
		Html:  `<u>1-*2-~34~-6*-7</u>`,
		Plain: `1-*2-~34~-6*-7`,
	},
	{
		Title: "date",
		Raw:   "<2000-01-02T10:15:00.000000-0700>",
		Html:  "<time>02.01.00 10:15</time>",
		Plain: "02.01.00 10:15",
	},
	{
		Title: "date in text",
		Raw:   "Тестовая задача с крайним сроком <2020-11-12T13:00:45.795000Z>",
		Plain: "Тестовая задача с крайним сроком 12.11.20 13:00",
	},
	{
		Title: "clear-all",
		Raw:   "_x_ ~y~ *z* /a/ `b` ```c```",
		Plain: "x y z a b c",
	},
	{
		Title: "username",
		Raw:   "aaa @user1_name1 @user2_name2",
		Plain: "aaa @user1_name1 @user2_name2",
	},
	{
		Title: "bold-username",
		Raw:   "aaa *@user1_name1 @user2_name2*",
		Plain: "aaa @user1_name1 @user2_name2",
	},
	{
		Title: "punctuation",
		Raw:   "_123_?",
		Plain: "123?",
	},
	{
		Title: "double",
		Raw:   "_123__ f",
		Html:  "_123__ f",
		Plain: "_123__ f",
	},
	{
		Title: "double-double",
		Raw:   "_123 _ f",
		Html:  "_123 _ f",
		Plain: "_123 _ f",
	},
	{
		Title: "emoji",
		Raw:   "_hey_ *hop😂* _лалалей_",
		Html:  "<u>hey</u> <b>hop😂</b> <u>лалалей</u>",
		Plain: "hey hop😂 лалалей",
	},
	{
		Title: "multi pre",
		Raw:   "``` 1 ```2``` ```3",
		Html:  "<pre>1</pre>2<pre></pre>3",
		Plain:   "123",
	},
}
