package dto

import "url-shorter/internal/api/domain"

type ShortlinkDto struct {
	Shortlink string `json:"shortlink"`
	Url       string `json:"url"`
}

func (sl *ShortlinkDto) ToDomain() *domain.ShortlinkDomain {
	return &domain.ShortlinkDomain{
		Shortlink: sl.Shortlink,
		Url:       sl.Url,
	}
}
