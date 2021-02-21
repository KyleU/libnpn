module github.com/kyleu/libnpn/npnscript

go 1.16

require (
	github.com/kyleu/libnpn/npncore v0.0.1 // npn
	github.com/robertkrimen/otto v0.0.0-20200922221731-ef014fd054ac
	github.com/yuin/gopher-lua v0.0.0-20200816102855-ee81675732da
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
)

replace github.com/kyleu/libnpn/npncore => ../npncore
