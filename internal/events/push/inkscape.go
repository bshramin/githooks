package push

import (
	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/sirupsen/logrus"
)

func Inkscape(p gitlab.PushEventPayload) error {
	logrus.Info("Inkscape push event")
	return nil
}
