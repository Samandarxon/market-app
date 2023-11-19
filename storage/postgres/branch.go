package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Samandarxon/market_app/models"
)

type branchRepo struct {
	db *sql.DB
}

func NewBranchRepo(db *sql.DB) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (r branchRepo) Create(req *models.CreateBranch) (*models.Branch, error) {

	var (
		branch models.Branch
		query  = `
		INSERT INTO "branch"(
			"name",
			"phone",
			"photo",
			"work_start_hour",
			"work_end_hour",
			"address",
			"status",
			"updated_at"
		) VALUES ($1 , $2 , $3 ,$4, $5, $6, $7 , NOW()) RETURNING *`
	)

	resp := r.db.QueryRow(
		query,
		req.Name,
		req.Phone,
		req.Photo,
		req.WorkStartHour,
		req.WorkEndHour,
		req.Address,
		req.Status,
	)

	err := resp.Scan(
		&branch.Id,
		&branch.Name,
		&branch.Phone,
		&branch.Photo,
		&branch.WorkStartHour,
		&branch.WorkEndHour,
		&branch.Address,
		&branch.DeliveryPrice,
		&branch.Status,
		&branch.CreatedAt,
		&branch.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.PrimaryKeyBranchId{Id: branch.Id})
}

func (r *branchRepo) GetByID(req *models.PrimaryKeyBranchId) (*models.Branch, error) {

	var (
		branch models.Branch
		query  = `
			SELECT
				"id",
				"name",
				"phone",
				"photo",
				"work_start_hour",
				"work_end_hour",
				"address",
				"delivery_price",
				"status",
				"created_at",
				"updated_at"	
			FROM "branch"
			WHERE "id" = $1
		`
	)

	fmt.Println("**************************************", req.Id)
	err := r.db.QueryRow(query, req.Id).Scan(
		&branch.Id,
		&branch.Name,
		&branch.Phone,
		&branch.Photo,
		&branch.WorkStartHour,
		&branch.WorkEndHour,
		&branch.Address,
		&branch.DeliveryPrice,
		&branch.Status,
		&branch.CreatedAt,
		&branch.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &branch, nil
}

func (r *branchRepo) GetList(req *models.GetListBranchRequest) (*models.GetListBranchResponse, error) {
	var (
		resp   models.GetListBranchResponse
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

	if len(req.Name) > 0 {
		where += " AND name ILIKE '%" + req.Name + "%'"
	}
	if len(req.DateTo) > 0 || len(req.DateFrom) > 0 {
		where += " AND created_at BETWEEN '" + req.DateFrom + "' AND '" + req.DateTo + "'"
	}

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"name",
			"phone",
			"photo",
			"work_start_hour",
			"work_end_hour",
			"address",
			"delivery_price",
			"status",
			"created_at",
			"updated_at"	
		FROM "branch"
	`

	query += where + sort + offset + limit
	fmt.Println(query)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			branch models.Branch
		)

		err = rows.Scan(
			&resp.Count,
			&branch.Id,
			&branch.Name,
			&branch.Phone,
			&branch.Photo,
			&branch.WorkStartHour,
			&branch.WorkEndHour,
			&branch.Address,
			&branch.DeliveryPrice,
			&branch.Status,
			&branch.CreatedAt,
			&branch.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Branches = append(resp.Branches, &branch)
	}

	return &resp, nil
}

func (c *branchRepo) Update(req *models.UpdateBranch) (int64, error) {

	var (
		query = `UPDATE branch SET 
					"name" = $2,
					"phone" = $3,
					"photo" = $4,
					"work_start_hour" = $5,
					"work_end_hour" = $6,
					"address" = $7,
					"status" = $8,
					"updated_at" = Now()
				 WHERE id=$1`
	)

	fmt.Println("UP<<<<<<<<<<<<< ", req.Id)
	result, err := c.db.Exec(
		query,
		req.Id,
		req.Name,
		req.Phone,
		req.Photo,
		req.WorkStartHour,
		req.WorkEndHour,
		req.Address,
		req.Status,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}

func (c *branchRepo) Delete(req *models.PrimaryKeyBranchId) (int64, error) {

	var (
		query = `DELETE FROM branch WHERE id=$1`
	)

	result, err := c.db.Exec(query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}

func (c *branchRepo) DeleteAll() (int64, error) {

	var (
		query = `DELETE FROM branch`
	)

	result, err := c.db.Exec(query)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}
