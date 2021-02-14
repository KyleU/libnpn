module github.com/kyleu/libnpn/npntemplate

go 1.15

require (
	cloud.google.com/go v0.65.0 // indirect
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/kyleu/libnpn/npnconnection v0.0.1 // npn
	github.com/kyleu/libnpn/npncore v0.0.1 // npn
	github.com/kyleu/libnpn/npnservice v0.0.1 // npn
	github.com/kyleu/libnpn/npnuser v0.0.1 // npn
	github.com/kyleu/libnpn/npnweb v0.0.1 // npn
	github.com/shiyanhui/hero v0.0.2
)

replace github.com/kyleu/libnpn/npnconnection => ../npnconnection

replace github.com/kyleu/libnpn/npncore => ../npncore

replace github.com/kyleu/libnpn/npnuser => ../npnuser

replace github.com/kyleu/libnpn/npnservice => ../npnservice

replace github.com/kyleu/libnpn/npnweb => ../npnweb
