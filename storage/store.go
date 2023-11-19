package storage

import "github.com/Samandarxon/market_app/models"

type StorageI interface {
	Category() CategoryRepoI
	Client() ClientRepoI
	Product() ProductRepoI
	Branch() BranchRepoI
	Order() OrderRepoI
	OrderProduct() OrderProductRepoI
}

type CategoryRepoI interface {
	Create(req *models.CreateCategory) (*models.Category, error)
	GetByID(req *models.PrimaryKeyCategoryId) (*models.Category, error)
	GetList(req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(req *models.UpdateCategory) (int64, error)
	Delete(req *models.PrimaryKeyCategoryId) (int64, error)
	DeleteAll() (int64, error)
}
type ProductRepoI interface {
	Create(req *models.CreateProduct) (*models.Product, error)
	GetByID(req *models.PrimaryKeyProductId) (*models.Product, error)
	GetList(req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(req *models.UpdateProduct) (int64, error)
	Delete(req *models.PrimaryKeyProductId) (int64, error)
	DeleteAll() (int64, error)
}

type ClientRepoI interface {
	Create(req *models.CreateClient) (*models.Client, error)
	GetByID(req *models.PrimaryKeyClientId) (*models.Client, error)
	GetList(req *models.GetListClientRequest) (*models.GetListClientResponse, error)
	Update(req *models.UpdateClient) (int64, error)
	Delete(req *models.PrimaryKeyClientId) (int64, error)
	DeleteAll() (int64, error)
}
type BranchRepoI interface {
	Create(req *models.CreateBranch) (*models.Branch, error)
	GetByID(req *models.PrimaryKeyBranchId) (*models.Branch, error)
	GetList(req *models.GetListBranchRequest) (*models.GetListBranchResponse, error)
	Update(req *models.UpdateBranch) (int64, error)
	Delete(req *models.PrimaryKeyBranchId) (int64, error)
	DeleteAll() (int64, error)
}

type OrderRepoI interface {
	Create(req *models.CreateOrder) (*models.Order, error)
	GetByID(req *models.OrderPrimaryKey) (*models.Order, error)
	GetList(req *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Update(req *models.UpdateOrder) (int64, error)
	Delete(req *models.OrderPrimaryKey) error
	StatusUpdate(models.CheckStatus) (models.Order, error)
}

type OrderProductRepoI interface {
	Create(req *models.CreateOrderProduct) (*models.OrderProduct, error)
	GetByID(req *models.OrderProductPrimaryKey) (*models.OrderProduct, error)
	GetList(req *models.GetListOrderProductRequest) (*models.GetListOrderProductResponse, error)
	Update(req *models.UpdateOrderProduct) (int64, error)
	Delete(req *models.OrderProductPrimaryKey) error
}
