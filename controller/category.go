package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Samandarxon/market_app/helpers"
	"github.com/Samandarxon/market_app/models"
)

func (c *Handler) Category(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		c.CreateCategory(w, r)

	case "GET":
		var values = r.URL.Query()
		if _, ok := values["id"]; ok {
			c.GetByIDCategory(w, r)
		} else {
			c.GetListCategory(w, r)
		}
	case "PUT":
		c.UpdateCategory(w, r)

	case "DELETE":
		var values = r.URL.Query()
		if _, ok := values["id"]; ok {
			c.DeleteCategory(w, r)
		} else {
			c.DeleteAllCategory(w, r)
		}

	}
}

func (c *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	var createCategory models.CreateCategory

	err := json.NewDecoder(r.Body).Decode(&createCategory)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	if createCategory.ParentId != "" {
		if !helpers.IsValidUUID(createCategory.ParentId) {
			handleResponse(w, http.StatusBadRequest, "parent id is not uuid")
			return
		}
	}

	resp, err := c.storage.Category().Create(&createCategory)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, resp)
}

func (c *Handler) GetByIDCategory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(w, http.StatusBadRequest, "ID in not uuid")
		return
	}

	resp, err := c.storage.Category().GetByID(&models.PrimaryKeyCategoryId{Id: id})

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

func (c *Handler) GetListCategory(w http.ResponseWriter, r *http.Request) {

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

	search := r.URL.Query().Get("search")
	if err != nil {
		handleResponse(w, http.StatusBadRequest, "invalid query search")
		return
	}

	resp, err := c.storage.Category().GetList(&models.GetListCategoryRequest{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var updateCategory models.UpdateCategory

	err := json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println(updateCategory)

	if updateCategory.ParentId != "" {
		if !helpers.IsValidUUID(updateCategory.ParentId) {
			handleResponse(w, http.StatusBadRequest, "parent id is not uuid")
			return
		}
	}

	rowsAffected, err := c.storage.Category().Update(&updateCategory)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(w, http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := c.storage.Category().GetByID(&models.PrimaryKeyCategoryId{Id: updateCategory.Id})

	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")

	fmt.Println(id)
	resp, err := c.storage.Category().Delete(&models.PrimaryKeyCategoryId{Id: id})

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

func (c *Handler) DeleteAllCategory(w http.ResponseWriter, r *http.Request) {
	resp, err := c.storage.Category().DeleteAll()
	var respStr string

	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	respStr = fmt.Sprintf("There were %d items in the database, all deleted", resp)
	handleResponse(w, http.StatusOK, respStr)
}
