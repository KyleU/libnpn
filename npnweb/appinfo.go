package npnweb

import (
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnservice/auth"
	"github.com/kyleu/libnpn/npnservice/user"
	"github.com/sirupsen/logrus"
)

type AppInfo interface {
	Debug() bool
	Files() npncore.FileLoader
	User() user.Service
	Auth() auth.Service
	Logger() *logrus.Logger
	Valid() bool
	Public() bool
}
