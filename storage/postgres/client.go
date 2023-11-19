package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Samandarxon/market_app/models"
)

type clientRepo struct {
	db *sql.DB
}

func NewClientRepo(db *sql.DB) *clientRepo {
	return &clientRepo{
		db: db,
	}
}

func (r clientRepo) Create(req *models.CreateClient) (*models.Client, error) {
	var (
		client models.Client
		query  = `
		INSERT INTO "client"(
			"first_name",
			"last_name",
			"phone",
			"photo",
			"date_of_birth",
			"updated_at"
		) VALUES ($1 , $2 , $3 ,$4, $5, NOW()) RETURNING *`
	)

	resp := r.db.QueryRow(
		query,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.Photo,
		req.DateOfBirth,
	)

	err := resp.Scan(
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.Phone,
		&client.Photo,
		&client.DateOfBirth,
		&client.CreatedAt,
		&client.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.PrimaryKeyClientId{Id: client.Id})
}

func (r *clientRepo) GetByID(req *models.PrimaryKeyClientId) (*models.Client, error) {

	var (
		client models.Client
		query  = `
			SELECT
				"id",
				"first_name",
				"last_name",
				"phone",
				"photo",
				"date_of_birth",
				"created_at",
				"updated_at"	
			FROM "client"
			WHERE "id" = $1
		`
	)

	fmt.Println("**************************************", req.Id)
	err := r.db.QueryRow(query, req.Id).Scan(
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.Phone,
		&client.Photo,
		&client.DateOfBirth,
		&client.CreatedAt,
		&client.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *clientRepo) GetList(req *models.GetListClientRequest) (*models.GetListClientResponse, error) {
	var (
		resp   models.GetListClientResponse
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
		where += " AND first_name ILIKE" + " '%" + req.Search + "%'"
	}

	// if len(req.Query) > 0 {
	// 	where += req.Query
	// }

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"first_name",
			"last_name",
			"phone",
			"photo",
			"date_of_birth",
			"created_at",
			"updated_at"	
		FROM "client"
	`

	query += where + sort + offset + limit
	fmt.Println(query)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			client models.Client
		)

		err = rows.Scan(
			&resp.Count,
			&client.Id,
			&client.FirstName,
			&client.LastName,
			&client.Phone,
			&client.Photo,
			&client.DateOfBirth,
			&client.CreatedAt,
			&client.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Clients = append(resp.Clients, &client)
	}

	return &resp, nil
}

func (c *clientRepo) Update(req *models.UpdateClient) (int64, error) {

	var (
		query = `UPDATE client SET 
					"first_name" = $2,
					"last_name" = $3,
					"phone" = $4,
					"photo" = $5,
					"date_of_birth" = $6,
					"updated_at" = Now()
				 WHERE id=$1`
	)

	result, err := c.db.Exec(
		query,
		req.Id,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.Photo,
		req.DateOfBirth,
	)
	fmt.Println(err, "#######################################")
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}

func (c *clientRepo) Delete(req *models.PrimaryKeyClientId) (int64, error) {

	var (
		query = `DELETE FROM client WHERE id=$1`
	)

	result, err := c.db.Exec(query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}

func (c *clientRepo) DeleteAll() (int64, error) {

	var (
		query = `DELETE FROM client`
	)

	result, err := c.db.Exec(query)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}
