package models

type Category struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	ParentId  string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateCategory struct {
	Title    string `json:"title"`
	Image    string `json:"image"`
	ParentId string `json:"parent_id"`
}

type UpdateCategory struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	ParentId string `json:"parent_id"`
}

type PrimaryKeyCategoryId struct {
	Id string `json:"id"`
}

type GetListCategoryRequest struct {
	Offset int64
	Limit  int64
	Search string
}
type GetListCategoryResponse struct {
	Count      int64
	Categories []*Category
}
