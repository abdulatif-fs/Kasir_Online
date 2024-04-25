package controller

import (
	"kasir_online/database"
	"kasir_online/query"
	"kasir_online/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) {
	var (
		result gin.H
	)

	menu, err := query.GetMenu(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": menu,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertMenu(c *gin.Context) {
	var menu structs.Menu

	err := c.ShouldBindJSON(&menu)
	if err != nil {
		panic(err)
	}

	err = query.InsertMenu(database.DbConnection, menu)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Menu",
	})
}

func GetKasir(c *gin.Context) {
	var (
		result gin.H
	)

	kasir, err := query.GetKasir(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": kasir,
		}
	}

	c.JSON(http.StatusOK, result)
}
