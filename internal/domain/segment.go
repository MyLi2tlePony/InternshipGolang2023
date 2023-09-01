package domain

import "time"

type Segment struct {
	Name      string     `json:"name"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type CreateLinkCSVRequest struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type CreateLinkCSVResponse struct {
	FileName int `json:"fileName"`
}
