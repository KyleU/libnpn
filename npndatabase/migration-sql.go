package npndatabase

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/kyleu/libnpn/npncore"

	"emperror.dev/errors"
	"golang.org/x/text/language"
)

type MigrationFile struct {
	Title string
	F     func(*strings.Builder)
}

type MigrationFiles []*MigrationFile

// Set this in your application's startup if needed, these build your schema from scratch, called by DBWipe
var InitialSchemaMigrations = MigrationFiles{}

// Set this in your application's startup, it's used as the store for all migrations, call by Migrate
var DatabaseMigrations = MigrationFiles{}

func exec(file *MigrationFile, s *Service, logger *logrus.Logger) (string, error) {
	sb := &strings.Builder{}
	file.F(sb)
	sql := sb.String()
	startNanos := npncore.TimerStart()
	_, err := s.Exec(sql, nil, -1)
	if err != nil {
		return "", errors.Wrap(err, "cannot execute ["+file.Title+"]")
	}
	ms := npncore.MicrosToMillis(language.AmericanEnglish, npncore.TimerEnd(startNanos))
	logger.Debug(fmt.Sprintf("ran query [%s] in [%v]", file.Title, ms))
	return sql, nil
}
