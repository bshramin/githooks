package push

import (
	"github.com/go-playground/webhooks/v6/github"
	"github.com/sirupsen/logrus"
)

func Goxy(p github.PushPayload) error {
	logrus.Info("Goxy push event")

	return nil
}
