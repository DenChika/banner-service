package delete

import (
	"banner_service/lib/api/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type QueryParams struct {
	Id int `json:"id"`
}

type BannerDeleter interface {
	Delete(id int) error
}

func New(log *slog.Logger, deleter BannerDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.Context().Value(chi.RouteCtxKey).(*QueryParams)
		id := queryParams.Id

		err := deleter.Delete(id)
		if err != nil {
			log.Error("Failed deleting banner", "error", err)
			render.JSON(w, r, response.InternalServerError("Failed deleting banner"))

			return
		}

		log.Info("Banner was deleted successfully")
		render.JSON(w, r, response.OkResponse{
			Status: http.StatusNoContent,
		})
	}
}
