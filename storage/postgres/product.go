package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Samandarxon/market_app/helpers"
	"github.com/Samandarxon/market_app/models"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (r productRepo) Create(req *models.CreateProduct) (*models.Product, error) {

	var (
		product      models.Product
		productId, _ = helpers.NewIncrementId(r.db, "product", 8)
		query        = `
		INSERT INTO "product"(
			"id",
			"category_id",
			"title",
			"description",
			"photo",
			"price",
			"updated_at"
		) VALUES ($1 , $2 , $3 ,$4, $5, $6, NOW()) RETURNING *`
	)

	resp := r.db.QueryRow(
		query,
		productId(),
		req.CategoryId,
		req.Title,
		req.Description,
		req.Photo,
		req.Price,
	)

	err := resp.Scan(
		&product.Id,
		&product.CategoryId,
		&product.Title,
		&product.Description,
		&product.Photo,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.PrimaryKeyProductId{Id: product.Id})
}

func (r *productRepo) GetByID(req *models.PrimaryKeyProductId) (*models.Product, error) {

	var (
		product models.Product
		query   = `
			SELECT
				"id",
				"category_id",
				"title",
				"description",
				"photo",
				"price",
				"created_at",
				"updated_at"	
			FROM "product"
			WHERE "id" = $1
		`
	)

	fmt.Println("**************************************", req.Id)
	err := r.db.QueryRow(query, req.Id).Scan(
		&product.Id,
		&product.CategoryId,
		&product.Title,
		&product.Description,
		&product.Photo,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) GetList(req *models.GetListProductRequest) (*models.GetListProductResponse, error) {
	var (
		resp   models.GetListProductResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		sort   = " ORDER BY created_at DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if len(req.CategoryId) > 0 || len(req.Title) > 0 {
		where += " AND title ILIKE '%" + req.Title + "%'" +
			" AND category_id::TEXT ILIKE '%" + req.CategoryId + "%'"
	}

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"category_id",
			"title",
			"description",
			"photo",
			"price",
			"created_at",
			"updated_at"	
		FROM "product"
	`

	query += where + sort + offset + limit
	fmt.Println(query)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			product models.Product
		)

		err = rows.Scan(
			&resp.Count,
			&product.Id,
			&product.CategoryId,
			&product.Title,
			&product.Description,
			&product.Photo,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &product)
	}

	return &resp, nil
}

func (c *productRepo) Update(req *models.UpdateProduct) (int64, error) {

	var (
		query = `UPDATE product SET 
					"category_id" = $2,
					"title" = $3,
					"description" = $4,
					"photo" = $5,
					"price" = $6,
					"updated_at" = Now()
				 WHERE id=$1`
	)

	fmt.Println("UP<<<<<<<<<<<<< ", req.Id)
	result, err := c.db.Exec(
		query,
		req.Id,
		req.CategoryId,
		req.Title,
		req.Description,
		req.Photo,
		req.Price,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}

func (c *productRepo) Delete(req *models.PrimaryKeyProductId) (int64, error) {

	var (
		query = `DELETE FROM product WHERE id=$1`
	)

	result, err := c.db.Exec(query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}

func (c *productRepo) DeleteAll() (int64, error) {

	var (
		query = `DELETE FROM product`
	)

	result, err := c.db.Exec(query)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}
