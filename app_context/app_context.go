package app_context

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AppCtx interface {
	GetMgoDB() *mongo.Client
}

type appCtx struct {
	mgoDB *mongo.Client
}

func NewAppCtx(mgoDB *mongo.Client) *appCtx {
	return &appCtx{
		mgoDB: mgoDB,
	}
}

func (ctx *appCtx) GetMgoDB() *mongo.Client {
	return ctx.mgoDB
}
