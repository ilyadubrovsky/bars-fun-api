package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ilyadubrovsky/bars"
	"log"
	"net/http"
)

const (
	api    = "/api"
	grades = "/grades"
)

func main() {
	router := gin.Default()
	apiGroup := router.Group(api)
	{
		apiGroup.POST(grades, handleGrades)
	}

	router.Run(":8080")
}

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleGrades(c *gin.Context) {
	var data UserData
	if err := c.BindJSON(&data); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	client := bars.NewClient()
	log.Printf("authorization user: %s : %s\n", data.Username, data.Password)
	if err := client.Authorization(context.TODO(), data.Username, data.Password); err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	jsonBytes, err := client.GetProgressTable()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Data(http.StatusOK, "application/json", jsonBytes)
}
