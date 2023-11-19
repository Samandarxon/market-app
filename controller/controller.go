package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Samandarxon/market_app/config"
	"github.com/Samandarxon/market_app/storage"
)

type Handler struct {
	cfg     *config.Config
	storage storage.StorageI
}

type Respons struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data,omitempty"`
}

func NewController(cfg *config.Config, strg storage.StorageI) *Handler {
	return &Handler{cfg: cfg, storage: strg}
}

func handleResponse(w http.ResponseWriter, status int, data interface{}) {

	var description string
	switch code := status; {
	case code < 400:
		description = "success"
		fmt.Println(data == 0)
		if data == 0 {
			description = "failed"
		}
		// sam, _ := json.MarshalIndent(Respons{
		// 	Status:      status,
		// 	Description: description,
		// 	Data:        data,
		// }, "", " ")
		// log.Println("Status ", status, string(sam))

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(Respons{
			Status:      status,
			Description: description,
			Data:        data,
		})
	default:
		description = "error"
		log.Println(config.Error, "erro while:", Respons{
			Status:      status,
			Description: description,
			Data:        data,
		})

		if code == 500 {
			description = "Internal Server Error"
		}

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(Respons{
			Status:      status,
			Description: description,
			Data:        data,
		})
	}
}

func getIntegerOrDefaultValue(value string, defaultValue int64) (int64, error) {

	if len(value) <= 0 {
		return defaultValue, nil
	}

	number, err := strconv.Atoi(value)
	return int64(number), err
}
