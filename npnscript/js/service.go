package js

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/robertkrimen/otto"
)

type Service struct {
	ot     *otto.Otto
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	ot := otto.New()
	return &Service{ot: ot, logger: logger}
}

func (s *Service) Set(k string, v interface{}) {
	err := s.ot.Set(k, v)
	if err != nil {
		s.logger.Warn(fmt.Sprintf("unable to set [%v] of type [%v]: %+v", k, v, err))
	}
}

func (s *Service) Call(code string) (interface{}, error) {
	ret, err := s.ot.Run(code)
	if err != nil {
		return nil, err
	}

	return ret.Export()
}
