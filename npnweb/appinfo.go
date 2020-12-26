package npnweb

import (
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnservice/auth"
	"github.com/kyleu/libnpn/npnservice/user"
	"logur.dev/logur"
)

type AppInfo interface {
	Debug() bool
	Files() npncore.FileLoader
	User() user.Service
	Auth() auth.Service
	Logger() logur.Logger
	Valid() bool
	Public() bool
	Secret() string
}
