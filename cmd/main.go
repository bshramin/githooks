package main

import (
	"net/http"

	"github.com/bshramin/githooks/internal/handlers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const appName = "githooks"

func main() {
	configInit(appName)

	http.HandleFunc("/webhooks/gitlab", handlers.GitlabHandler)
	http.HandleFunc("/webhooks/github", handlers.GithubHandler)

	port := viper.GetString("port")
	logrus.Info("Listening on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	logrus.Error(err)
}
