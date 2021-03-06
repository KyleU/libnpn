module github.com/kyleu/libnpn/npngraphql
go 1.16

require (
	emperror.dev/errors v0.7.0
	github.com/graphql-go/graphql v0.7.9
	github.com/kyleu/libnpn/npnweb v0.0.1 // npn
	github.com/kyleu/libnpn/npncontroller v0.0.1 // npn
)

replace github.com/kyleu/libnpn/npnconnection => ../npnconnection

replace github.com/kyleu/libnpn/npncore => ../npncore

replace github.com/kyleu/libnpn/npncontroller => ../npncontroller

replace github.com/kyleu/libnpn/npnservice => ../npnservice

replace github.com/kyleu/libnpn/npntemplate => ../npntemplate

replace github.com/kyleu/libnpn/npnuser => ../npnuser

replace github.com/kyleu/libnpn/npnweb => ../npnweb
