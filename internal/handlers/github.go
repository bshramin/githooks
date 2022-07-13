package handlers

import (
	"fmt"
	"net/http"

	"github.com/bshramin/githooks/internal/eventmap"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GithubHandler(w http.ResponseWriter, r *http.Request) {
	secret := viper.GetString("githooksecret")
	hook, _ := github.New(github.Options.Secret(secret))

	payload, err := hook.Parse(r, github.ReleaseEvent, github.PushEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			logrus.Info("Received event that we don't care about")
		} else {
			logrus.Error("Could not parse event: ", err)
		}
	}

	switch p := payload.(type) {

	case github.ReleasePayload:
		logrus.Info("Received tag event: ", p.Release.TagName)
		logrus.Info(p)
		f, ok := eventmap.GithubReleaseEventMap[fmt.Sprintf("%d", p.Repository.ID)]
		if ok {
			err := f(p)
			if err != nil {
				logrus.Error("Could not handle event: ", err)
			}
		} else {
			logrus.Info("No tag event handler for project: " + fmt.Sprintf("%d", p.Repository.ID))
		}

	case github.PushPayload:
		logrus.Info("Received tag event: ", p.Ref)
		logrus.Info(p)
		f, ok := eventmap.GithubPushEventMap[fmt.Sprintf("%d", p.Repository.ID)]
		if ok {
			err := f(p)
			if err != nil {
				logrus.Error("Could not handle event: ", err)
			}
		} else {
			logrus.Info("No push event handler for project: " + fmt.Sprintf("%d", p.Repository.ID))
		}

	default:
		logrus.Info("Received event and took no action: ", p)
	}
}
