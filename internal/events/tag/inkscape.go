package tag

import (
	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/sirupsen/logrus"
)

func Inkscape(p gitlab.TagEventPayload) error {
	logrus.Info("Inkscape tag event")

	return nil
}
