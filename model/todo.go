package model

import "time"

type (
	// A TODO expresses ...
	TODO struct {
		ID          int64     `json:"id" validate:"required"`
		Subject     string    `json:"subject" validate:"required"`
		Description string    `json:"descrption" validate:"required"`
		CreatedAt   time.Time `json:"created_at" validate:"required"`
		UpdatedAt   time.Time `json:"updated_at" validate:"required"`
	}

	// A CreateTODORequest expresses ...
	CreateTODORequest struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	// A CreateTODOResponse expresses ...
	CreateTODOResponse struct {
		TODO TODO `json:"todo"`
	}

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
