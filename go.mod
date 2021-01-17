module github.com/kyleu/libnpn

go 1.14

replace github.com/kyleu/libnpn/npncontroller => ./npncontroller

replace github.com/kyleu/libnpn/npnconnection => ./npnconnection

replace github.com/kyleu/libnpn/npncore => ./npncore

replace github.com/kyleu/libnpn/npndatabase => ./npndatabase

replace github.com/kyleu/libnpn/npnqueue => ./npnqueue

replace github.com/kyleu/libnpn/npnscript => ./npnscript

replace github.com/kyleu/libnpn/npnservice => ./npnservice

replace github.com/kyleu/libnpn/npnservice-db => ./npnservice-db

replace github.com/kyleu/libnpn/npnservice-fs => ./npnservice-fs

replace github.com/kyleu/libnpn/npntemplate => ./npntemplate

replace github.com/kyleu/libnpn/npnuser => ./npnuser

replace github.com/kyleu/libnpn/npnweb => ./npnweb

require (
	github.com/kyleu/libnpn/npnconnection v0.0.1 // npn
	github.com/kyleu/libnpn/npncontroller v0.0.1 // npn
	github.com/kyleu/libnpn/npncore v0.0.1 // npn
	github.com/kyleu/libnpn/npndatabase v0.0.1 // npn
	github.com/kyleu/libnpn/npnqueue v0.0.1 // npn
	github.com/kyleu/libnpn/npnscript v0.0.1 // npn
	github.com/kyleu/libnpn/npnservice v0.0.1 // npn
	github.com/kyleu/libnpn/npnservice-db v0.0.1 // npn
	github.com/kyleu/libnpn/npnservice-fs v0.0.1 // npn
	github.com/kyleu/libnpn/npntemplate v0.0.1 // npn
	github.com/kyleu/libnpn/npnuser v0.0.1 // npn
	github.com/kyleu/libnpn/npnweb v0.0.1 // npn
	logur.dev/logur v0.16.2
)
