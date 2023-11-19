package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Samandarxon/market_app/models"
)

func (c *Handler) Product(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		c.CreateProduct(w, r)

	case "GET":
		fmt.Println("#############################3")
		var values = r.URL.Query()
		fmt.Println(r.Method, values["id"])
		if _, ok := values["id"]; ok {
			c.GetByIDProduct(w, r)
		} else {
			c.GetListProduct(w, r)
		}
	case "PUT":
		c.UpdateProduct(w, r)

	case "DELETE":
		var values = r.URL.Query()
		if _, ok := values["id"]; ok {
			c.DeleteProduct(w, r)
		} else {
			c.DeleteAllProduct(w, r)
		}

	}
}

func (c *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var createProduct models.CreateProduct

	err := json.NewDecoder(r.Body).Decode(&createProduct)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	resp, err := c.storage.Product().Create(&createProduct)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, resp)
}

func (c *Handler) GetByIDProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SAlom men ishladim,,,")
	id := r.URL.Query().Get("id")

	// if !helpers.IsValidUUID(id) {
	// 	handleResponse(w, http.StatusBadRequest, "ID in not uuid")
	// 	return
	// }
	fmt.Println("**************************************", id)
	resp, err := c.storage.Product().GetByID(&models.PrimaryKeyProductId{Id: id})

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

func (c *Handler) GetListProduct(w http.ResponseWriter, r *http.Request) {

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

	title := r.URL.Query().Get("title")
	categoryId := r.URL.Query().Get("category_id")
	// search := r.URL.Query().Get("search")

	if err != nil {
		handleResponse(w, http.StatusBadRequest, "invalid query search")
		return
	}

	resp, err := c.storage.Product().GetList(&models.GetListProductRequest{
		Limit:      limit,
		Offset:     offset,
		Title:      title,
		CategoryId: categoryId,
		// Search: search,
	})

	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProduct models.UpdateProduct

	err := json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	fmt.Printf("UPDATE>>>>>> %+v\n", updateProduct)

	rowsAffected, err := c.storage.Product().Update(&updateProduct)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(w, http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := c.storage.Product().GetByID(&models.PrimaryKeyProductId{Id: updateProduct.Id})

	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")

	fmt.Println(id)
	resp, err := c.storage.Product().Delete(&models.PrimaryKeyProductId{Id: id})

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

func (c *Handler) DeleteAllProduct(w http.ResponseWriter, r *http.Request) {
	resp, err := c.storage.Product().DeleteAll()
	var respStr string

	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	respStr = fmt.Sprintf("There were %d items in the database, all deleted", resp)
	handleResponse(w, http.StatusOK, respStr)
}
