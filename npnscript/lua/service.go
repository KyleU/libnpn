package lua

import (
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

type Service struct {
	l      *lua.LState
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	opts := lua.Options{
		SkipOpenLibs:        true,
		IncludeGoStackTrace: true,
		MinimizeStackMemory: false,
	}
	l := lua.NewState(opts)
	return &Service{l: l, logger: logger}
}

func (s *Service) Set(k string, v interface{}) {
	// TODO s.l.SetGlobal(k, "")
}

func (s *Service) Call(code string) (interface{}, error) {
	defer s.l.Close()
	err := s.l.DoString(code)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
