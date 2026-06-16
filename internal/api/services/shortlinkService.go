package services

import (
	"context"

	"url-shorter/internal/api/dbservices"
	"url-shorter/internal/api/repository"
)

type ShortlinkServicer interface {
	Get(ctx context.Context, shortlink string)
	Post(ctx context.Context, args dbservices.InsertLinkParams)
	Delete(ctx context.Context, url string)
}

type ShortlinkService struct {
	r *repository.ShortlinkRepo
}

func NewShortlinkService(r *repository.ShortlinkRepo) *ShortlinkService {
	return &ShortlinkService{
		r: r,
	}
}

var _ ShortlinkServicer = (*ShortlinkService)(nil)

func (s *ShortlinkService) Get(ctx context.Context, shortlink string) {
	s.r.Get(ctx, shortlink)
}

func (s *ShortlinkService) Post(ctx context.Context, args dbservices.InsertLinkParams) {
	// TODO: use a domain model here instead of db service
	s.r.Post(ctx, args)
}
func (s *ShortlinkService) Delete(ctx context.Context, shortlink string) {
	s.r.Delete(ctx, shortlink)
}
