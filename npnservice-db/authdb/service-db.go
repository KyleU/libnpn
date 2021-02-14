package authdb

import (
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/kyleu/libnpn/npnservice/auth"

	"github.com/kyleu/libnpn/npndatabase"
	"github.com/kyleu/libnpn/npnservice/user"
)

type ServiceDatabase struct {
	enabled          bool
	enabledProviders auth.Providers
	redir            string
	db               *npndatabase.Service
	logger           *logrus.Logger
	users            user.Service
}

var _ auth.Service = (*ServiceDatabase)(nil)

func NewServiceDatabase(enabled bool, redir string /* actions *action.Service, */, db *npndatabase.Service, logger *logrus.Logger, users user.Service) auth.Service {
	if !strings.HasPrefix(redir, "http") {
		redir = "https://" + redir
	}
	if !strings.HasSuffix(redir, "/") {
		redir += "/"
	}

	svc := &ServiceDatabase{
		enabled: enabled,
		redir:   redir,
		db:      db,
		logger:  logger,
		users:   users,
	}

	for _, p := range auth.AllProviders {
		cfg := auth.GetConfig(svc, p)
		if cfg != nil {
			svc.enabledProviders = append(svc.enabledProviders, p)
		}
	}
	if len(svc.enabledProviders) == 0 {
		svc.enabled = false
	} else {
		logger.Info("auth service started for [" + strings.Join(svc.enabledProviders.Names(), ", ") + "]")
	}

	return svc
}

func (s *ServiceDatabase) Enabled() bool {
	return s.enabled
}

func (s *ServiceDatabase) EnabledProviders() auth.Providers {
	return s.enabledProviders
}

func (s *ServiceDatabase) FullURL(path string) string {
	return s.redir + strings.TrimPrefix(path, "/")
}
