package dispatcher

import (
	"bidchain/fabric/context"
)

type ICommandDispatcher interface {
	SendHttpRequest(ctx *context.HttpCommandContext)
}