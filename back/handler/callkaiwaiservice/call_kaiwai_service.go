package callkaiwaiservice

import (
	"github.com/mopeneko/callkaiwai-bbs/back/ent"
	"github.com/mopeneko/callkaiwai-bbs/back/gen/callkaiwaibbs/v1/callkaiwaibbsv1connect"
)

type CallKaiwaiService struct {
	db *ent.Client
}

func NewCallKaiwaiServiceHandler(db *ent.Client) callkaiwaibbsv1connect.CallKaiwaiBBSServiceHandler {
	return &CallKaiwaiService{
		db: db,
	}
}
