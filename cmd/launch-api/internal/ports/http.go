package ports

import (
	"math"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/edanko/nx/cmd/launch-api/internal/app"
	"github.com/edanko/nx/cmd/launch-api/internal/app/commands"
	"github.com/edanko/nx/cmd/launch-api/internal/app/queries"
	httperr "github.com/edanko/nx/pkg/http/errors"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(application app.Application) HTTPServer {
	return HTTPServer{
		app: application,
	}
}

func (h HTTPServer) GetKind(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	appKind, err := h.app.Queries.GetKind.Handle(r.Context(), queries.GetKindRequest{
		ID: id,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	kindResp := GetKindResponse{
		Id:          appKind.ID,
		Name:        appKind.Name,
		Description: appKind.Description,
		Status:      KindStatus(appKind.Status),
	}
	render.Respond(w, r, kindResp)
}

func (h HTTPServer) ListKinds(w http.ResponseWriter, r *http.Request, params ListKindsParams) {
	var limitInt int
	var after uuid.UUID
	if params.After != nil {
		after = uuid.MustParse(*params.After)
	}
	if params.Limit != nil {
		limitInt = *params.Limit
	}

	limit := int(math.Max(float64(limitInt), 20))

	// offset := (page * limit) - limit

	kindsResp := ListKindsResponse{
		// Page:  page,
		After: after.String(),
		Limit: limit,
		// Total: totalKinds,
	}

	// if totalKinds < 1 || offset > (totalKinds-1) {
	// 	render.Respond(w, r, kindsResp)
	// 	return
	// }

	appKinds, err := h.app.Queries.ListKinds.Handle(r.Context(), queries.ListKindsRequest{
		Status: (*string)(params.Status),
		Limit:  &limit,
		After:  &after,
		// Offset: &offset,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	kindsResp.Kinds = appKindsToResponse(appKinds)

	render.Respond(w, r, kindsResp)
}

func (h HTTPServer) CreateKind(w http.ResponseWriter, r *http.Request) {
	postKind := CreateKindRequest{}
	if err := render.Decode(r, &postKind); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	cmd := commands.CreateKind{
		ID:          uuid.New(),
		Name:        postKind.Name,
		Description: postKind.Description,
		Status:      postKind.Status,
	}

	err := h.app.Commands.Send(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.Header().Set("content-location", "/kinds/"+cmd.ID.String())
	w.WriteHeader(http.StatusNoContent)
}

func (h HTTPServer) DeleteKind(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	err := h.app.Commands.Send(r.Context(), commands.DeleteKind{
		ID: id,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
}

func (h HTTPServer) MakeKindDraft(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	err := h.app.Commands.Send(r.Context(), commands.MakeKindDraft{
		ID: id,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h HTTPServer) MakeKindPublished(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	err := h.app.Commands.Send(r.Context(), commands.MakeKindPublished{
		ID: id,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h HTTPServer) ChangeKindDescription(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	postKind := ChangeKindDescription{}
	if err := render.Decode(r, &postKind); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	cmd := commands.ChangeKindDescription{
		ID:          id,
		Description: postKind.Description,
	}

	err := h.app.Commands.Send(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h HTTPServer) ChangeKindName(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	postKind := ChangeKindName{}
	if err := render.Decode(r, &postKind); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	cmd := commands.ChangeKindName{
		ID:   id,
		Name: postKind.Name,
	}

	err := h.app.Commands.Send(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func appKindsToResponse(appKinds []queries.KindModel) []Kind {
	kinds := make([]Kind, 0, len(appKinds))
	for _, km := range appKinds {
		k := Kind{
			Id:          km.ID,
			Name:        km.Name,
			Description: km.Description,
			Status:      KindStatus(km.Status),
		}

		kinds = append(kinds, k)
	}

	return kinds
}
