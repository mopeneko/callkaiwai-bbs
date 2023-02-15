package callkaiwaiservice

import (
	"github.com/mopeneko/callkaiwai-bbs/back/gen/callkaiwaibbs/v1/callkaiwaibbsv1connect"
)

type CallKaiwaiService struct {
}

func NewCallKaiwaiServiceHandler() callkaiwaibbsv1connect.CallKaiwaiBBSServiceHandler {
	return &CallKaiwaiService{}
}
