module github.com/kyleu/libnpn/npnservice-db

go 1.15

require (
	github.com/kyleu/libnpn/npncore v0.0.1 // npn
	github.com/kyleu/libnpn/npndatabase v0.0.1 // npn
	github.com/kyleu/libnpn/npnservice v0.0.1 // npn
	github.com/kyleu/libnpn/npnuser v0.0.1 // npn
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
)

replace github.com/kyleu/libnpn/npncore => ../npncore
replace github.com/kyleu/libnpn/npndatabase => ../npndatabase
replace github.com/kyleu/libnpn/npnservice => ../npnservice
replace github.com/kyleu/libnpn/npnuser => ../npnuser
