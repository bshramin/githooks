package push

import (
	"github.com/go-playground/webhooks/v6/github"
	"github.com/sirupsen/logrus"
)

func Cappu(p github.PushPayload) error {
	logrus.Info("Cappu push event")

	return nil
}
