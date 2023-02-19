package callkaiwaiservice

import (
	"context"
	"net"
	"strconv"

	"github.com/bufbuild/connect-go"
	callkaiwaibbsv1 "github.com/mopeneko/callkaiwai-bbs/back/gen/callkaiwaibbs/v1"
	"golang.org/x/xerrors"
)

func (s *CallKaiwaiService) CreateNewPost(ctx context.Context, req *connect.Request[callkaiwaibbsv1.CreateNewPostRequest]) (*connect.Response[callkaiwaibbsv1.CreateNewPostResponse], error) {
	tx, err := s.db.Tx(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to initialize transaction: %w", err)
	}

	post, err := tx.Post.
		Create().
		SetName(req.Msg.Name).
		SetGender(int(req.Msg.Gender)).
		SetIntroduction(req.Msg.Introduction).
		SetTweetURL(req.Msg.TweetUrl).
		SetTiktokURL(req.Msg.TiktokUrl).
		SetContactURL(req.Msg.ContactUrl).
		SetContactID(req.Msg.ContactId).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to create post: %w", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", req.Peer().Addr)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse peer address: %w", err)
	}

	if _, err := tx.IPAddressLog.
		Create().
		SetIPAddress(addr.IP.String()).
		SetPost(post).
		Save(ctx); err != nil {
		return nil, xerrors.Errorf("failed to create IP Address log: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, xerrors.Errorf("failed to commit transaction: %w", err)
	}

	return connect.NewResponse(&callkaiwaibbsv1.CreateNewPostResponse{
		Id: strconv.Itoa(post.ID),
	}), nil
}
