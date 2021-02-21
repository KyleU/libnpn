module github.com/kyleu/libnpn/npnservice

go 1.16

require (
	github.com/kyleu/libnpn/npncore v0.0.1 // npn
	github.com/kyleu/libnpn/npnuser v0.0.1 // npn
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
)

replace github.com/kyleu/libnpn/npncore => ../npncore
replace github.com/kyleu/libnpn/npnuser => ../npnuser
