package save

import (
	"banner_service/internal/dto"
	"banner_service/lib/api/response"
	"errors"
	"github.com/go-chi/render"
	"io"
	"log/slog"
	"net/http"
)

type RequestBody struct {
	TagIds    []int                  `json:"tag_ids"`
	FeatureId int                    `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}

type BannerSaver interface {
	Save(banner *dto.PostBanner) (int, error)
}

func New(log *slog.Logger, saver BannerSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RequestBody

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("Request body is empty")
			render.JSON(w, r, response.BadRequest("Empty request"))

			return
		}
		if err != nil {
			log.Error("Failed decoding body", "error", err)
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

		id, err := saver.Save(postBannerDto)
		if err != nil {
			log.Error("Failed saving banner", "error", err)
			render.JSON(w, r, response.InternalServerError("Failed saving banner"))

			return
		}

		log.Info("Saved banner with id", slog.Int("id", id))
		render.JSON(w, r, response.Ok(
			map[string]interface{}{
				"id": id,
			}))
	}
}
