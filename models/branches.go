package models

type Branch struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Photo         string `json:"photo"`
	WorkStartHour string `json:"work_start_hour"`
	WorkEndHour   string `json:"work_end_hour"`
	Address       string `json:"address"`
	DeliveryPrice int    `json:"delivery_price"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type CreateBranch struct {
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Photo         string `json:"photo"`
	WorkStartHour string `json:"work_start_hour"`
	WorkEndHour   string `json:"work_end_hour"`
	Address       string `json:"address"`
	DeliveryPrice int    `json:"delivery_price"`
	Status        string `json:"status"`
}

type UpdateBranch struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Photo         string `json:"photo"`
	WorkStartHour string `json:"work_start_hour"`
	WorkEndHour   string `json:"work_end_hour"`
	Address       string `json:"address"`
	// DeliveryPrice int    `json:"delivery_price"`
	Status string `json:"status"`
}

type PrimaryKeyBranchId struct {
	Id string `json:"id"`
}

type GetListBranchRequest struct {
	Offset   int64
	Limit    int64
	DateFrom string
	DateTo   string
	Name     string
}
type GetListBranchResponse struct {
	Count    int64
	Branches []*Branch
}
