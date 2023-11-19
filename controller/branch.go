package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Samandarxon/market_app/helpers"
	"github.com/Samandarxon/market_app/models"
)

func (c *Handler) Branch(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		c.CreateBranch(w, r)

	case "GET":
		fmt.Println("#############################3")
		var values = r.URL.Query()
		fmt.Println(r.Method, values["id"])
		if _, ok := values["id"]; ok {
			c.GetByIDBranch(w, r)
		} else {
			c.GetListBranch(w, r)
		}
	case "PUT":
		c.UpdateBranch(w, r)

	case "DELETE":
		var values = r.URL.Query()
		if _, ok := values["id"]; ok {
			c.DeleteBranch(w, r)
		} else {
			c.DeleteAllBranch(w, r)
		}

	}
}

func (c *Handler) CreateBranch(w http.ResponseWriter, r *http.Request) {

	var createBranch models.CreateBranch

	err := json.NewDecoder(r.Body).Decode(&createBranch)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	resp, err := c.storage.Branch().Create(&createBranch)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, resp)
}

func (c *Handler) GetByIDBranch(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(w, http.StatusBadRequest, "ID in not uuid")
		return
	}
	fmt.Println("**************************************", id)
	resp, err := c.storage.Branch().GetByID(&models.PrimaryKeyBranchId{Id: id})

	if err == sql.ErrNoRows {
		handleResponse(w, http.StatusBadRequest, "No rows in result set")
		return
	}

	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)

}

func (c *Handler) GetListBranch(w http.ResponseWriter, r *http.Request) {

	limit, err := getIntegerOrDefaultValue(r.URL.Query().Get("limit"), 10)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, "invalid query limit")
		return
	}

	offset, err := getIntegerOrDefaultValue(r.URL.Query().Get("offset"), 0)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, "invalid query offset")
		return
	}

	name := r.URL.Query().Get("name")
	dateFrom := r.URL.Query().Get("from")
	dateTo := r.URL.Query().Get("to")

	if err != nil {
		handleResponse(w, http.StatusBadRequest, "invalid query search")
		return
	}

	resp, err := c.storage.Branch().GetList(&models.GetListBranchRequest{
		Limit:    limit,
		Offset:   offset,
		DateFrom: dateFrom,
		DateTo:   dateTo,
		Name:     name,
	})

	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) UpdateBranch(w http.ResponseWriter, r *http.Request) {
	var updateBranch models.UpdateBranch

	err := json.NewDecoder(r.Body).Decode(&updateBranch)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	fmt.Printf("UPDATE>>>>>> %+v\n", updateBranch)

	rowsAffected, err := c.storage.Branch().Update(&updateBranch)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(w, http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := c.storage.Branch().GetByID(&models.PrimaryKeyBranchId{Id: updateBranch.Id})

	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) DeleteBranch(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")

	fmt.Println(id)
	resp, err := c.storage.Branch().Delete(&models.PrimaryKeyBranchId{Id: id})

	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}
	if resp != 0 {
		handleResponse(w, http.StatusOK, "Item removed")
		return
	}
	handleResponse(w, http.StatusOK, "Data was not deleted")
}

func (c *Handler) DeleteAllBranch(w http.ResponseWriter, r *http.Request) {
	resp, err := c.storage.Branch().DeleteAll()
	var respStr string

	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	respStr = fmt.Sprintf("There were %d items in the database, all deleted", resp)
	handleResponse(w, http.StatusOK, respStr)
}
