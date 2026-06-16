package domain

import "url-shorter/internal/api/dbservices"

type ShortlinkDomain struct {
	Shortlink string
	Url       string
}

func (sl *ShortlinkDomain) ToDB() *dbservices.InsertLinkParams {
	return &dbservices.InsertLinkParams{
		Shortlink: sl.Shortlink,
		Url:       sl.Url,
	}
}
