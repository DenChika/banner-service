package get

import (
	"banner_service/internal/dto"
	"banner_service/lib/api/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type QueryParams struct {
	TagId           int  `json:"tag_id"`
	FeatureId       int  `json:"feature_id"`
	UseLastRevision bool `json:"use_last_revision"`
}

type UserBannerGetter interface {
	Get(banner *dto.GetUserBanner) (any, error)
}

func New(log *slog.Logger, getter UserBannerGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.Context().Value(chi.RouteCtxKey).(*QueryParams)

		getUserBannerDto := &dto.GetUserBanner{
			FeatureId:       queryParams.FeatureId,
			TagId:           queryParams.TagId,
			UseLastRevision: queryParams.UseLastRevision,
		}

		banner, err := getter.Get(getUserBannerDto)
		if err != nil {
			log.Error("Failed getting user banner", "error", err)
			render.JSON(w, r, response.InternalServerError("Failed getting user banner"))

			return
		}
		if banner == nil {
			log.Error("User banner was not found", "error", err)
			render.JSON(w, r, response.NotFound())

			return
		}

		log.Info("User banner was gotten successfully", slog.Any("resp", banner))
		render.JSON(w, r, response.Ok(banner))
	}
}
