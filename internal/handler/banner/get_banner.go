package banner

import (
	"banner_service/internal/dto"
	"banner_service/lib/api/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type QueryParams struct {
	FeatureId int `json:"feature_id"`
	TagId     int `json:"tag_id"`
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
}

type BannerGetter interface {
	GetBanner(banner *dto.GetBanner) (*dto.Banner, error)
}

func New(log *slog.Logger, getter BannerGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.Context().Value(chi.RouteCtxKey).(*QueryParams)

		getBannerDto := &dto.GetBanner{
			FeatureId: queryParams.FeatureId,
			TagId:     queryParams.TagId,
			Limit:     queryParams.Limit,
			Offset:    queryParams.Offset,
		}

		banner, err := getter.GetBanner(getBannerDto)
		if err != nil {
			log.Error("Error getting banner", err)
			render.JSON(w, r, response.InternalServerError("Failed getting banner"))

			return
		}
		if banner == nil {
			render.JSON(w, r, response.NotFound())

			return
		}

		render.JSON(w, r, response.Ok(banner))
	}
}
