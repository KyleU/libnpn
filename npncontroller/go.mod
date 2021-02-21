module github.com/kyleu/libnpn/npncontroller

go 1.16

require (
	github.com/kyleu/libnpn/npnconnection v0.0.1 // npn
	github.com/kyleu/libnpn/npncore v0.0.1 // npn
	github.com/kyleu/libnpn/npntemplate v0.0.1 // npn
	github.com/kyleu/libnpn/npnweb v0.0.1 // npn
)

replace github.com/kyleu/libnpn/npnconnection => ../npnconnection

replace github.com/kyleu/libnpn/npncore => ../npncore

replace github.com/kyleu/libnpn/npnservice => ../npnservice

replace github.com/kyleu/libnpn/npnuser => ../npnuser

replace github.com/kyleu/libnpn/npntemplate => ../npntemplate

replace github.com/kyleu/libnpn/npnweb => ../npnweb
