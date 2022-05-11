package models

import (
	"api/database"
)

func GetOne() string {
	pool := database.Pool()

	var data string
	pool.Get(&data, "SELECT 'It works!' AS name;")
	return data
}
