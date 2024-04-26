package controller

import (
	"kasir_online/database"
	"kasir_online/query"
	"kasir_online/structs"
	"net/http"
	"strconv"

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

func UpdateMenu(c *gin.Context) {
	var menu structs.Menu
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&menu)
	if err != nil {
		panic(err)
	}

	menu.Id = int(id)
	err = query.UpdateMenu(database.DbConnection, menu)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Succes Updates Menu",
	})
}

func DeleteMenu(c *gin.Context) {
	var menu structs.Menu
	id, _ := strconv.Atoi(c.Param("id"))

	menu.Id = int(id)

	err := query.DeleteMenu(database.DbConnection, menu)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Deleted Menu Success",
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

func InsertKasir(c *gin.Context) {
	var kasir structs.Kasir

	err := c.ShouldBindJSON(&kasir)
	if err != nil {
		panic(err)
	}

	err = query.InsertKasir(database.DbConnection, kasir)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Kasir",
	})
}

func GetTransaksi(c *gin.Context) {
	var (
		result gin.H
	)

	transaksi, err := query.GetTransaksi(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": transaksi,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertTransaksi(c *gin.Context) {
	var transaksi structs.Transaksi

	err := c.ShouldBindJSON(&transaksi)
	if err != nil {
		panic(err)
	}

	err = query.InsertTransaksi(database.DbConnection, transaksi)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Transaksi",
	})
}

func GetDetailTransaksi(c *gin.Context) {
	var (
		result           gin.H
		detail_transaksi structs.Detail_transaksi
	)
	id, _ := strconv.Atoi(c.Param("id"))
	detail_transaksi.Transaksi_id = int(id)

	transaksi, err := query.GetDetailTransaksi(database.DbConnection, detail_transaksi)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": transaksi,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertDetailTransaksi(c *gin.Context) {
	var Detail_transaksi structs.Detail_transaksi
	var menu structs.Menu

	err := c.ShouldBindJSON(&Detail_transaksi)
	if err != nil {
		panic(err)
	}

	err = query.InsertDetailTransaksi(database.DbConnection, Detail_transaksi, menu)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Detail_Transaksi",
	})
}
