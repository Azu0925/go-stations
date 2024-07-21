package model

import "time"

type (
	// A TODO expresses ...
	TODO struct {
		ID          int64     `json:"id",required`
		Subject     string    `json:"subject",required`
		Description string    `json:"descrption",required`
		CreatedAt   time.Time `json:"created_at",required`
		UpdatedAt   time.Time `json:"updated_at",required`
	}

	// A CreateTODORequest expresses ...
	CreateTODORequest struct{}
	// A CreateTODOResponse expresses ...
	CreateTODOResponse struct{}

	// A ReadTODORequest expresses ...
	ReadTODORequest struct{}
	// A ReadTODOResponse expresses ...
	ReadTODOResponse struct{}

	// A UpdateTODORequest expresses ...
	UpdateTODORequest struct{}
	// A UpdateTODOResponse expresses ...
	UpdateTODOResponse struct{}

	// A DeleteTODORequest expresses ...
	DeleteTODORequest struct{}
	// A DeleteTODOResponse expresses ...
	DeleteTODOResponse struct{}
)
