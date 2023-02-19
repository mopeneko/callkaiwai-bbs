package service

import (
	"context"

	"github.com/mopeneko/callkaiwai-bbs/back/ent"
	"golang.org/x/xerrors"
)

type PostService struct {
	db *ent.Client
}

func NewPostService(db *ent.Client) *PostService {
	return &PostService{
		db: db,
	}
}

func (s *PostService) CreateNewPost(ctx context.Context, post *ent.Post, srcIPAddr string) (*ent.Post, error) {
	tx, err := s.db.Tx(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to initialize transaction: %w", err)
	}

	obj, err := tx.Post.
		Create().
		SetName(post.Name).
		SetGender(int(post.Gender)).
		SetIntroduction(post.Introduction).
		SetTweetURL(post.TweetURL).
		SetTiktokURL(post.TiktokURL).
		SetContactURL(post.ContactURL).
		SetContactID(post.ContactID).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to create post: %w", err)
	}

	if _, err := tx.IPAddressLog.
		Create().
		SetIPAddress(srcIPAddr).
		SetPost(obj).
		Save(ctx); err != nil {
		return nil, xerrors.Errorf("failed to create IP Address log: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, xerrors.Errorf("failed to commit transaction: %w", err)
	}

	return obj, nil
}
