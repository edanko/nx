package ports

import (
	"math"
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
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

func (h HTTPServer) ListApplicants(w http.ResponseWriter, r *http.Request, params ListApplicantsParams) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) CreateApplicant(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) DeleteApplicant(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) GetApplicant(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ChangeApplicantDescription(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ChangeApplicantName(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) MakeApplicantDraft(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) MakeApplicantPublished(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ListLaunches(w http.ResponseWriter, r *http.Request, params ListLaunchesParams) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) CreateLaunch(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) DeleteLaunch(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) GetLaunch(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ChangeLaunchDescription(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ChangeLaunchName(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) MakeLaunchDraft(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) MakeLaunchPublished(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ListOrders(w http.ResponseWriter, r *http.Request, params ListOrdersParams) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) DeleteOrder(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) GetOrder(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ChangeOrderDescription(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ChangeOrderName(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) MakeOrderDraft(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) MakeOrderPublished(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ListUsers(w http.ResponseWriter, r *http.Request, params ListUsersParams) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) DeleteUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) GetUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ChangeUserDescription(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ChangeUserName(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) MakeUserDraft(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) MakeUserPublished(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// TODO implement me
	panic("implement me")
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

	if params.Limit != nil {
		limitInt = *params.Limit
	}
	limit := int(math.Max(float64(limitInt), 10))

	appKinds, nextCursor, err := h.app.Queries.ListKinds.Handle(
		r.Context(),
		queries.ListKindsRequest{
			Status: (*string)(params.Status),
			Limit:  &limit,
			Cursor: params.Cursor,
		},
	)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	kindsResp := ListKindsResponse{
		NextCursor: nextCursor,
		Kinds:      appKindsToResponse(appKinds),
	}
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

	err := h.app.CommandBus.Send(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.Header().Set("content-location", "/kinds/"+cmd.ID.String())
	w.WriteHeader(http.StatusNoContent)
}

func (h HTTPServer) DeleteKind(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	err := h.app.CommandBus.Send(r.Context(), commands.DeleteKind{
		ID: id,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
}

func (h HTTPServer) MakeKindDraft(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	err := h.app.CommandBus.Send(r.Context(), commands.MakeKindDraft{
		ID: id,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h HTTPServer) MakeKindPublished(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	err := h.app.CommandBus.Send(r.Context(), commands.MakeKindPublished{
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

	err := h.app.CommandBus.Send(r.Context(), cmd)
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

	err := h.app.CommandBus.Send(r.Context(), cmd)
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
