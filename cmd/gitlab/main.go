package main

import (
	"fmt"

	"net/http"

	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/sirupsen/logrus"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := gitlab.New(gitlab.Options.Secret("thesecret"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Something received")
		logrus.Error("Something received")
		payload, err := hook.Parse(r, gitlab.TagEvents, gitlab.PushEvents)
		if err != nil {
			if err == gitlab.ErrEventNotFound {
				// ok event wasn't one of the ones asked to be parsed
			} else {
				logrus.Error("Could not parse event: ", err)
			}
		}

		switch payload.(type) {

		case gitlab.TagEventPayload:
			tagPayload := payload.(gitlab.TagEventPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", tagPayload)

		case gitlab.PushEventPayload:
			pushPayload := payload.(gitlab.PushEventPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pushPayload)
		}

	})
	http.ListenAndServe(":3000", nil)
}
