// Code generated by hero.
// source: github.com/kyleu/npn/libnpn/npntemplate/html/components/uuid.html
// DO NOT EDIT!
package npntemplate

import (
	"bytes"

	"github.com/gofrs/uuid"
	"github.com/shiyanhui/hero"
)

func UUID(u *uuid.UUID, buffer *bytes.Buffer) {
	if u != nil {
		buffer.WriteString(`<span title="`)
		hero.EscapeHTML(u.String(), buffer)
		buffer.WriteString(`">`)
		hero.EscapeHTML(u.String()[0:8], buffer)
		buffer.WriteString(`</span>`)
	}

}
