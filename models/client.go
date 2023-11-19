package models

type Client struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	Photo       string `json:"photo"`
	DateOfBirth string `json:"date_of_birth"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateClient struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	Photo       string `json:"photo"`
	DateOfBirth string `json:"date_of_birth"`
}

type UpdateClient struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	Photo       string `json:"photo"`
	DateOfBirth string `json:"date_of_birth"`
}

type PrimaryKeyClientId struct {
	Id string `json:"id"`
}

type GetListClientRequest struct {
	Offset int64
	Limit  int64
	Search string
}
type GetListClientResponse struct {
	Count   int64
	Clients []*Client
}
