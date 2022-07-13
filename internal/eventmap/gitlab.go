package eventmap

import (
	"github.com/bshramin/githooks/internal/events/push"
	"github.com/bshramin/githooks/internal/events/tag"
	"github.com/go-playground/webhooks/v6/gitlab"
)

// EventProjectHandlerMap is a map of event type to project ID to event handler function
var GitlabTagEventMap = map[string]func(p gitlab.TagEventPayload) error{
	"3472737": tag.Inkscape, // https://gitlab.com/inkscape/inkscape
	"6094330": tag.Commento, // https://gitlab.com/commento/commento
}

// EventProjectHandlerMap is a map of event type to project ID to event handler function
var GitlabPushEventMap = map[string]func(p gitlab.PushEventPayload) error{
	"3472737": push.Inkscape, // https://gitlab.com/inkscape/inkscape
	"6094330": push.Commento, // https://gitlab.com/commento/commento
}
