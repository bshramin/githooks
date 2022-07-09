package events

import "github.com/bshramin/githooks/internal/handlers/tag"

var EventProjectHandlerMap = map[string]map[string]func() error{
	"tag": {
		"expressinterface": tag.ExpressInterface,
	},
}
