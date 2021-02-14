module github.com/kyleu/libnpn/npnweb

go 1.15

require (
	emperror.dev/errors v0.7.0
	github.com/kyleu/libnpn/npncore v0.0.1 // npn
	github.com/kyleu/libnpn/npnservice v0.0.1 // npn
	github.com/kyleu/libnpn/npnuser v0.0.1 // npn
	github.com/gorilla/mux v1.7.4
	github.com/mitchellh/mapstructure v1.1.2
	github.com/gorilla/sessions v1.2.0
)

replace github.com/kyleu/libnpn/npncore => ../npncore

replace github.com/kyleu/libnpn/npnservice => ../npnservice

replace github.com/kyleu/libnpn/npnuser => ../npnuser
