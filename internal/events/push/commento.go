package push

import (
	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/sirupsen/logrus"
)

func Commento(p gitlab.PushEventPayload) error {
	logrus.Info("Commento push event")

	return nil
}
