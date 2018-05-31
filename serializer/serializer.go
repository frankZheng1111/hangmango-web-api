package serrializer

import "time"

type BaseResource struct {
	TotalCount int `json:"total_count"`
}

type ResourceTimestamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
