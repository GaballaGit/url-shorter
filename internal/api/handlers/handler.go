package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"url-shorter/internal/api/dto"
	"url-shorter/internal/api/services"
)

type ShortlinkHandler struct {
	shortlinkService services.ShortlinkService
}

func NewHandler(slService services.ShortlinkService) *ShortlinkHandler {
	return &ShortlinkHandler{
		shortlinkService: slService,
	}
}

func (sl *ShortlinkHandler) GetUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("url")

	ctx, cancel := context.WithTimeout(context.Background(), (12 * time.Second))
	defer cancel()

	sl.shortlinkService.Get(ctx, id)
}

func (sl *ShortlinkHandler) PostUrlHandler(w http.ResponseWriter, r *http.Request) {
	var newlink dto.ShortlinkDto
	err := json.NewDecoder(r.Body).Decode(&newlink)
	if err != nil {
		w.Write([]byte("error:" + err.Error()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), (12 * time.Second))
	defer cancel()

	sl.shortlinkService.Post(ctx, newlink)
}

func (sl *ShortlinkHandler) DeleteUrlHandler(w http.ResponseWriter, r *http.Response) {
	url := r.Header.Get("url")

	ctx := context.Background()

	sl.shortlinkService.Delete(ctx, url)
}
