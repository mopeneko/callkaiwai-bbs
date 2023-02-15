package callkaiwaiservice

import (
	"context"

	"github.com/bufbuild/connect-go"
	callkaiwaibbsv1 "github.com/mopeneko/callkaiwai-bbs/back/gen/callkaiwaibbs/v1"
)

func (s *CallKaiwaiService) Ping(ctx context.Context, _ *connect.Request[callkaiwaibbsv1.PingRequest]) (*connect.Response[callkaiwaibbsv1.PingResponse], error) {
	return connect.NewResponse(&callkaiwaibbsv1.PingResponse{
		Text: "Pong",
	}), nil
}
