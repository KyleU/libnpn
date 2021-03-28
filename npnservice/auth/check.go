package auth

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/gofrs/uuid"
)

var (
	// List of allowed auth providers
	RequiredAuthProviders = []string{}
	// List of allowed user email domains
	RequiredAuthDomains = []string{}
)

// Checks if the provided user can sign in to the app
func Check(s Service, userID uuid.UUID, logger *logrus.Logger) bool {
	if (len(RequiredAuthProviders) == 0) && (len(RequiredAuthDomains) == 0) {
		return true
	}
	auths := s.GetByUserID(userID, nil)
	if len(auths) == 0 {
		logger.Debug(fmt.Sprintf("unable to authorize user [%v]: no auths", userID))
		return false
	}
	for _, p := range RequiredAuthProviders {
		matched := false
		for _, a := range auths {
			if a.Provider.Key == p {
				matched = true
			}
		}
		if (!matched) && (len(RequiredAuthProviders) > 0) {
			logger.Debug(fmt.Sprintf("unable to authorize user [%v]: no matching providers", userID))
			return false
		}
	}
	for _, d := range RequiredAuthDomains {
		matched := false
		for _, a := range auths {
			if strings.HasSuffix(a.Email, d) {
				matched = true
			}
		}
		if (!matched) && (len(RequiredAuthDomains) > 0) {
			logger.Debug(fmt.Sprintf("unable to authorize user [%v]: no matching domains", userID))
			return false
		}
	}
	return true
}
