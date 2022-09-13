package app_context

import (
	"go-challenege/pkg/tokenprovider"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppCtx interface {
	GetMgoDB() *mongo.Client
	GetATProvider() tokenprovider.Provider
	GetRTProvider() tokenprovider.Provider
}

type appCtx struct {
	mgoDB      *mongo.Client
	atProvider tokenprovider.Provider
	rtProvider tokenprovider.Provider
}

func NewAppCtx(mgoDB *mongo.Client, atProvider,
	rtProvider tokenprovider.Provider) *appCtx {
	return &appCtx{
		mgoDB:      mgoDB,
		atProvider: atProvider,
		rtProvider: rtProvider,
	}
}

func (ctx *appCtx) GetMgoDB() *mongo.Client {
	return ctx.mgoDB
}

func (ctx *appCtx) GetRTProvider() tokenprovider.Provider {
	return ctx.rtProvider
}

func (ctx *appCtx) GetATProvider() tokenprovider.Provider {
	return ctx.atProvider
}
