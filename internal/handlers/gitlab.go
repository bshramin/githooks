package handlers

import (
	"fmt"
	"net/http"

	"github.com/bshramin/githooks/internal/eventmap"
	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GitlabHandler(w http.ResponseWriter, r *http.Request) {
	secret := viper.GetString("githooksecret")
	hook, _ := gitlab.New(gitlab.Options.Secret(secret))

	payload, err := hook.Parse(r, gitlab.TagEvents, gitlab.PushEvents)
	if err != nil {
		if err == gitlab.ErrEventNotFound {
			logrus.Info("Received event that we don't care about")
		} else {
			logrus.Error("Could not parse event: ", err)
		}
	}

	switch p := payload.(type) {

	case gitlab.TagEventPayload:
		logrus.Info("Received tag event: ", p.Ref)
		logrus.Info(p)
		f, ok := eventmap.GitlabTagEventMap[fmt.Sprintf("%d", p.ProjectID)]
		if ok {
			err := f(p)
			if err != nil {
				logrus.Error("Could not handle event: ", err)
			}
		} else {
			logrus.Info("No tag event handler for project: " + fmt.Sprintf("%d", p.ProjectID))
		}

	case gitlab.PushEventPayload:
		logrus.Info("Received tag event: ", p.Ref)
		logrus.Info(p)
		f, ok := eventmap.GitlabPushEventMap[fmt.Sprintf("%d", p.ProjectID)]
		if ok {
			err := f(p)
			if err != nil {
				logrus.Error("Could not handle event: ", err)
			}
		} else {
			logrus.Info("No push event handler for project: " + fmt.Sprintf("%d", p.ProjectID))
		}

	default:
		logrus.Info("Received event and took no action: ", p)
	}
}
