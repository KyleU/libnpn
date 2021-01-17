module github.com/kyleu/libnpn/npnservice-fs

go 1.15

require (
	github.com/kyleu/libnpn/npncore v0.0.1 // npn
	github.com/kyleu/libnpn/npnservice v0.0.1 // npn
	github.com/kyleu/libnpn/npnuser v0.0.1 // npn
)

replace github.com/kyleu/libnpn/npncore => ../npncore
replace github.com/kyleu/libnpn/npnservice => ../npnservice
replace github.com/kyleu/libnpn/npnuser => ../npnuser
