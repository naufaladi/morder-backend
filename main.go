package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	Id    string  `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

type item_allergens struct {
	ItemId     string `json:"itemId"`
	AllergenId string `json:"allergenId"`
}

type allergen struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

var items = []item{
	{Id: "1", Title: "Hamburger", Price: 13.50},
	{Id: "2", Title: "Spaghetti", Price: 8.75},
	{Id: "3", Title: "Pizza", Price: 15.20},
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)

	router.Run("localhost:5555")
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}
