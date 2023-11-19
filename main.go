package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Samandarxon/market_app/helpers"
)

func main() {
	dbl, err := helpers.NewIncrementId(&sql.DB{}, "product", 7)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < 100; i++ {
		fmt.Println(dbl())
	}

}
