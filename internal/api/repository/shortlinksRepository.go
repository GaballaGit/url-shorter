package repository

import (
	"context"

	"url-shorter/internal/api/dbservices"
	"url-shorter/internal/api/domain"
)

type ShortlinkRepository interface {
	Get(ctx context.Context, shortlink string)
	Post(ctx context.Context, args domain.ShortlinkDomain)
	Delete(ctx context.Context, url string)
}

type ShortlinkRepo struct {
	db dbservices.Queries
}

func NewShortlinkRepo(db dbservices.Queries) *ShortlinkRepo {
	return &ShortlinkRepo{
		db: db,
	}
}

var _ ShortlinkRepository = (*ShortlinkRepo)(nil)

func (r *ShortlinkRepo) Get(ctx context.Context, shortlink string) {
	r.db.GetUrl(ctx, shortlink)
}

func (r *ShortlinkRepo) Post(ctx context.Context, args domain.ShortlinkDomain) {
	// TODO: conv to db model
	r.db.InsertLink(ctx, *args.ToDB())
}

func (r *ShortlinkRepo) Delete(ctx context.Context, shortlink string) {
	r.db.DeleteUrl(ctx, shortlink)
}
