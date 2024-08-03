package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	j := json.NewEncoder(w)

	switch r.Method {
	case "POST":
		result, err := h.Create(ctx, &model.CreateTODORequest{})
		if err != nil {
			w.WriteHeader(400)

			return
		}

		if err := j.Encode(result); err != nil {
			log.Println(err)
			w.WriteHeader(500)

			return
		}

	case "GET":
		result, err := h.Read(ctx, &model.ReadTODORequest{})
		if err != nil {
			w.WriteHeader(400)

			return
		}

		if err := j.Encode(result); err != nil {
			log.Println(err)
			w.WriteHeader(500)

			return
		}
	case "PUT":
		result, err := h.Update(ctx, &model.UpdateTODORequest{})
		if err != nil {
			w.WriteHeader(400)

			return
		}

		if err := j.Encode(result); err != nil {
			log.Println(err)
			w.WriteHeader(500)

			return
		}
	case "DELETE":
		result, err := h.Delete(ctx, &model.DeleteTODORequest{})
		if err != nil {
			w.WriteHeader(400)

			return
		}

		if err := j.Encode(result); err != nil {
			log.Println(err)
			w.WriteHeader(500)

			return
		}
	default:
		w.WriteHeader(400)

		return
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	result, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	if err != nil {
		return nil, err
	}

	return &model.CreateTODOResponse{
		TODO: *result,
	}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
