package release

import (
	"github.com/go-playground/webhooks/v6/github"
	"github.com/sirupsen/logrus"
)

func Cappu(p github.ReleasePayload) error {
	logrus.Info("Cappu release event")

	return nil
}
