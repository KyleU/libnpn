// Code generated by hero.
// source: github.com/kyleu/npn/libnpn/npntemplate/html/components/pagination.html
// DO NOT EDIT!
package npntemplate

import (
	"bytes"
	"net/url"

	"github.com/kyleu/libnpn/npncore"
	"github.com/shiyanhui/hero"
)

func Pagination(section string, count int, params *npncore.Params, req *url.URL, buffer *bytes.Buffer) {
	if params != nil {
		if params.HasNextPage(count) {
			buffer.WriteString(`
    <div class="right">
      <a href="?`)
			hero.EscapeHTML(params.NextPage().ToQueryString(req), buffer)
			buffer.WriteString(`">Next page</a>
    </div>
  `)
		}
		if params.HasPreviousPage() {
			buffer.WriteString(`
    <div class="left">
      <a href="?`)
			hero.EscapeHTML(params.PreviousPage().ToQueryString(req), buffer)
			buffer.WriteString(`">Previous page</a>
    </div>
  `)
		}
	}

}
