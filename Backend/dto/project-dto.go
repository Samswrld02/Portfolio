package dto

// reading dto project
type ReadProjectDTO struct {
	ID          int
	Title       string
	Description string
	Link        string
}

// create dto project
type UpdateProjectDTO struct {
	Title       *string `json:"title" validate:"min=3,omitempty"`
	Description *string `json:"description" validate:"omitempty,min=3"`
}
