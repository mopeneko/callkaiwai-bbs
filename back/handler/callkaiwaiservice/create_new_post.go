package callkaiwaiservice

import (
	"context"
	"net"
	"strconv"

	"github.com/bufbuild/connect-go"
	"github.com/mopeneko/callkaiwai-bbs/back/ent"
	callkaiwaibbsv1 "github.com/mopeneko/callkaiwai-bbs/back/gen/callkaiwaibbs/v1"
	"golang.org/x/xerrors"
)

func (s *CallKaiwaiService) CreateNewPost(ctx context.Context, req *connect.Request[callkaiwaibbsv1.CreateNewPostRequest]) (*connect.Response[callkaiwaibbsv1.CreateNewPostResponse], error) {
	post := &ent.Post{
		Name:         req.Msg.Name,
		Gender:       int(req.Msg.Gender),
		Introduction: req.Msg.Introduction,
		TweetURL:     req.Msg.TweetUrl,
		TiktokURL:    req.Msg.TiktokUrl,
		ContactURL:   req.Msg.ContactUrl,
		ContactID:    req.Msg.ContactId,
	}

	addr, err := net.ResolveTCPAddr("tcp", req.Peer().Addr)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse peer address: %w", err)
	}

	obj, err := s.PostService.CreateNewPost(ctx, post, addr.IP.String())
	if err != nil {
		return nil, xerrors.Errorf("failed to create new post: %w", err)
	}

	return connect.NewResponse(&callkaiwaibbsv1.CreateNewPostResponse{
		Id: strconv.Itoa(obj.ID),
	}), nil
}
