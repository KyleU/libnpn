// Code generated by hero.
// source: github.com/kyleu/libnpn/npntemplate/html/layout/head.html
// DO NOT EDIT!
package npntemplate

import (
	"bytes"

	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnweb"
	"github.com/shiyanhui/hero"
)

func Head(ctx *npnweb.RequestContext, buffer *bytes.Buffer) {
	buffer.WriteString(`
<title>`)
	hero.EscapeHTML(ctx.Title, buffer)
	buffer.WriteString(`</title>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0, viewport-fit=cover">

<meta property="og:title" content="`)
	hero.EscapeHTML(ctx.Title, buffer)
	buffer.WriteString(`">
<meta property="og:type" content="website">
<meta property="og:locale" content="en_US">

`)
	for _, ss := range npncore.IncludedStylesheets {
		buffer.WriteString(`<link rel="stylesheet" media="screen" href="`)
		hero.EscapeHTML(ss, buffer)
		buffer.WriteString(`">
`)
	}
	buffer.WriteString(`<link rel="stylesheet" media="screen" href="/assets/vendor/`)
	hero.EscapeHTML(npncore.AppKey, buffer)
	buffer.WriteString(`.css">
`)
	for _, scr := range npncore.IncludedScripts {
		buffer.WriteString(`<script src="`)
		hero.EscapeHTML(scr, buffer)
		buffer.WriteString(`"></script>
`)
	}
	buffer.WriteString(`<script src="/assets/vendor/`)
	hero.EscapeHTML(npncore.AppKey, buffer)
	buffer.WriteString(`.js"></script>
`)

}
