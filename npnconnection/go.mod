module github.com/kyleu/libnpn/npnconnection

go 1.15

require (
	emperror.dev/errors v0.7.0
	github.com/kyleu/libnpn/npncore v0.0.1 // npn
	github.com/kyleu/libnpn/npnuser v0.0.1 // npn
	github.com/gorilla/websocket v1.4.1
	logur.dev/logur v0.16.2
)

replace github.com/kyleu/libnpn/npncore => ../npncore

replace github.com/kyleu/libnpn/npnuser => ../npnuser
