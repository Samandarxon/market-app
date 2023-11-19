package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Samandarxon/market_app/helpers"
	"github.com/Samandarxon/market_app/models"
)

func (c *Handler) Client(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		c.CreateClient(w, r)

	case "GET":
		fmt.Println("#############################3")
		var values = r.URL.Query()
		fmt.Println(r.Method, values["id"])
		if _, ok := values["id"]; ok {
			c.GetByIDClient(w, r)
		} else {
			c.GetListClient(w, r)
		}
	case "PUT":
		c.UpdateClient(w, r)

	case "DELETE":
		var values = r.URL.Query()
		if _, ok := values["id"]; ok {
			c.DeleteClient(w, r)
		} else {
			c.DeleteAllClient(w, r)
		}

	}
}

func (c *Handler) CreateClient(w http.ResponseWriter, r *http.Request) {

	var createClient models.CreateClient

	err := json.NewDecoder(r.Body).Decode(&createClient)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	resp, err := c.storage.Client().Create(&createClient)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, resp)
}

func (c *Handler) GetByIDClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SAlom men ishladim,,,")
	id := r.URL.Query().Get("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(w, http.StatusBadRequest, "ID in not uuid")
		return
	}
	fmt.Println("**************************************", id)
	resp, err := c.storage.Client().GetByID(&models.PrimaryKeyClientId{Id: id})

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

func (c *Handler) GetListClient(w http.ResponseWriter, r *http.Request) {

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

	resp, err := c.storage.Client().GetList(&models.GetListClientRequest{
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

func (c *Handler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	var updateClient models.UpdateClient

	err := json.NewDecoder(r.Body).Decode(&updateClient)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println(updateClient)

	rowsAffected, err := c.storage.Client().Update(&updateClient)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(w, http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := c.storage.Client().GetByID(&models.PrimaryKeyClientId{Id: updateClient.Id})

	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")

	fmt.Println(id)
	resp, err := c.storage.Client().Delete(&models.PrimaryKeyClientId{Id: id})

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

func (c *Handler) DeleteAllClient(w http.ResponseWriter, r *http.Request) {
	resp, err := c.storage.Client().DeleteAll()
	var respStr string

	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	respStr = fmt.Sprintf("There were %d items in the database, all deleted", resp)
	handleResponse(w, http.StatusOK, respStr)
}
