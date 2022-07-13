package eventmap

import (
	"github.com/bshramin/githooks/internal/events/push"
	"github.com/bshramin/githooks/internal/events/release"
	"github.com/go-playground/webhooks/v6/github"
)

// EventProjectHandlerMap is a map of event type to project ID to event handler function
// How to get github project ID: https://stackoverflow.com/a/47223479/11553370
var GithubReleaseEventMap = map[string]func(p github.ReleasePayload) error{
	"451068035": release.Goxy,  // https://github.com/bshramin/goxy
	"430428787": release.Cappu, // https://github.com/bshramin/cappu

}

// EventProjectHandlerMap is a map of event type to project ID to event handler function
var GithubPushEventMap = map[string]func(p github.PushPayload) error{
	"451068035": push.Goxy,  // https://github.com/bshramin/goxy
	"430428787": push.Cappu, // https://github.com/bshramin/cappu
}
