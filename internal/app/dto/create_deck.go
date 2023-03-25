package dto

type CreateDeckDTO struct {
	Shuffle     *bool    `json:"shuffle" binding:"required"`
	CardFilters []string `json:"cards"`
}
