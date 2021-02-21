module github.com/kyleu/libnpn/npndatabase

go 1.16

require (
	emperror.dev/errors v0.7.0
	github.com/jackc/pgx/v4 v4.7.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/kyleu/libnpn/npncore v0.0.1 // npn
)

replace github.com/kyleu/libnpn/npncore => ../npncore
