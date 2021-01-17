module github.com/kyleu/libnpn/npnexport

go 1.15

require (
	emperror.dev/errors v0.7.0
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/johnfercher/maroto v0.27.0
	github.com/kyleu/libnpn/npncore v0.0.0 // npn
)

replace github.com/kyleu/libnpn/npncore => ../npncore
