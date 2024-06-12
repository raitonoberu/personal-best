package model

import "time"

type Document struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type GetDocumentResponse struct {
	URL string `json:"url"`
}
