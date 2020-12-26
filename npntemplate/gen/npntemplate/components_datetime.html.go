// Code generated by hero.
// source: github.com/kyleu/libnpn/npntemplate/html/components/datetime.html
// DO NOT EDIT!
package npntemplate

import (
	"bytes"
	"time"

	"github.com/kyleu/libnpn/npncore"
	"github.com/shiyanhui/hero"
)

func DateTime(dt *time.Time, buffer *bytes.Buffer) {
	if dt != nil {
		hero.EscapeHTML(npncore.ToDateString(dt), buffer)
	}

}
