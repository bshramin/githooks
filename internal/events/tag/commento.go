package tag

import (
	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/sirupsen/logrus"
)

func Commento(p gitlab.TagEventPayload) error {
	logrus.Info("Commento tag event")

	return nil
}
