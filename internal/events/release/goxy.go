package release

import (
	"github.com/go-playground/webhooks/v6/github"
	"github.com/sirupsen/logrus"
)

func Goxy(p github.ReleasePayload) error {
	logrus.Info("Goxy release event")

	return nil
}
