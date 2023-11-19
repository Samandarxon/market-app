package models

type Product struct {
	Id          string `json:"id"`
	CategoryId  string `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Price       int    `json:"price"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateProduct struct {
	CategoryId  string `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Price       int    `json:"price"`
}

type UpdateProduct struct {
	Id          string `json:"id"`
	CategoryId  string `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Price       int    `json:"price"`
}

type PrimaryKeyProductId struct {
	Id string `json:"id"`
}

type GetListProductRequest struct {
	Offset     int64
	Limit      int64
	Title      string
	CategoryId string

	// Search string
}
type GetListProductResponse struct {
	Count    int64
	Products []*Product
}
