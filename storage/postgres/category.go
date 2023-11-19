package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Samandarxon/market_app/helpers"
	"github.com/Samandarxon/market_app/models"
)

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r categoryRepo) Create(req *models.CreateCategory) (*models.Category, error) {
	var (
		category models.Category
		parentId sql.NullString
		query    = `
		INSERT INTO "category"(
			"title",
			"image",
			"parent_id",
			"updated_at"
		) VALUES ($1 , $2 , $3 , NOW()) RETURNING *`
	)

	resp := r.db.QueryRow(
		query,
		req.Title,
		req.Image,
		helpers.NewNullString(req.ParentId),
	)

	err := resp.Scan(
		&category.Id,
		&category.Title,
		&category.Image,
		&parentId,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	category.ParentId = parentId.String

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.PrimaryKeyCategoryId{Id: category.Id})
}

func (r *categoryRepo) GetByID(req *models.PrimaryKeyCategoryId) (*models.Category, error) {

	var (
		category models.Category
		query    = `
			SELECT
				"id",
				"title",
				"image",
				COALESCE(CAST("parent_id" AS VARCHAR), ''),
				"created_at",
				"updated_at"	
			FROM "category"
			WHERE "id" = $1
		`
	)

	err := r.db.QueryRow(query, req.Id).Scan(
		&category.Id,
		&category.Title,
		&category.Image,
		&category.ParentId,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepo) GetList(req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {
	var (
		resp   models.GetListCategoryResponse
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

	if len(req.Search) > 0 {
		where += " AND title ILIKE" + " '%" + req.Search + "%'"
	}

	// if len(req.Query) > 0 {
	// 	where += req.Query
	// }

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"title",
			"image",
			"parent_id",
			"created_at",
			"updated_at"
		FROM "category"
	`

	query += where + sort + offset + limit
	fmt.Println(query)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			category models.Category
			parentID sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&category.Id,
			&category.Title,
			&category.Image,
			&parentID,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		category.ParentId = parentID.String
		resp.Categories = append(resp.Categories, &category)
	}

	return &resp, nil
}

func (c *categoryRepo) Update(req *models.UpdateCategory) (int64, error) {

	var (
		query = `UPDATE category SET 
					"title" = $2,
					"image" = $3,
					"parent_id" = $4,
					"updated_at" = Now()
				 WHERE id=$1`
	)

	result, err := c.db.Exec(
		query,
		req.Id,
		req.Title,
		req.Image,
		helpers.NewNullString(req.ParentId),
	)
	fmt.Println(err, "#######################################")
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}

func (c *categoryRepo) Delete(req *models.PrimaryKeyCategoryId) (int64, error) {

	var (
		query = `DELETE FROM category WHERE id=$1`
	)

	result, err := c.db.Exec(query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}

func (c *categoryRepo) DeleteAll() (int64, error) {

	var (
		query = `DELETE FROM category;`
	)

	result, err := c.db.Exec(query)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}
