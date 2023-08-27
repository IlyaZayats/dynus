package v1

type NewSlugRequest struct {
	Name string `json:"name" binding:"required"`
}

type RemoveSlugRequest struct {
	Name string `json:"name" binding:"required"`
}

type ActiveUserSlugsRequest struct {
	UserId string `uri:"user_id" binding:"required"`
}

type UpdateUserSlugsRequest struct {
	UserId      string   `uri:"user_id"`
	InsertSlugs []string `json:"insert_slugs"`
	DeleteSlugs []string `json:"delete_slugs"`
}
