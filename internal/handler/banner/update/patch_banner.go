package update

import (
	"banner_service/internal/dto"
	"banner_service/lib/api/response"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"log/slog"
	"net/http"
)

type QueryParams struct {
	Id int `json:"id"`
}

type RequestBody struct {
	TagIds    []int                  `json:"tag_ids,omitempty"`
	FeatureId int                    `json:"feature_id,omitempty"`
	Content   map[string]interface{} `json:"content,omitempty"`
	IsActive  bool                   `json:"is_active,omitempty"`
}

type BannerPatcher interface {
	Patch(id int, banner *dto.PostBanner) error
}

func New(log *slog.Logger, patcher BannerPatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.Context().Value(chi.RouteCtxKey).(*QueryParams)
		id := queryParams.Id

		var req RequestBody

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("Request body is empty")
			render.JSON(w, r, response.BadRequest("Empty request"))

			return
		}
		if err != nil {
			log.Error("Failed to decode body", "error", err)
			render.JSON(w, r, response.BadRequest("Invalid request"))

			return
		}

		log.Info("Request body decoded successfully", slog.Any("req", req))

		postBannerDto := &dto.PostBanner{
			TagIds:    req.TagIds,
			FeatureId: req.FeatureId,
			Content:   req.Content,
			IsActive:  req.IsActive,
		}

		err = patcher.Patch(id, postBannerDto)
		if err != nil {
			log.Error("Failed updating banner", "error", err)
			render.JSON(w, r, response.InternalServerError("Failed updating banner"))

			return
		}

		log.Info("Updated banner with id", slog.Int("id", id), slog.Any("body", req))
		render.JSON(w, r, response.Ok())
	}
}
