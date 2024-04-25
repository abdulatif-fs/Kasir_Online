package main

import (
	"database/sql"
	"fmt"
	"kasir_online/controller"
	db "kasir_online/database"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "asdfghjkl"
	DBNAME   = "kasir"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME,
		// os.Getenv("PGHOST"),
		// os.Getenv("PGPORT"),
		// os.Getenv("PGUSER"),
		// os.Getenv("PGPASSWORD"),
		// os.Getenv("PGDATABASE"),
	)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	db.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()
	router.GET("/", index)
	router.GET("/menu", controller.GetMenu)
	router.POST("/menu", gin.BasicAuth(gin.Accounts{"admin": "admin"}), controller.InsertMenu)
	router.PUT("/menu/:id", gin.BasicAuth(gin.Accounts{"admin": "admin"}), controller.UpdateMenu)

	router.Run("localhost:8080")
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, "SELAMAT DATANG DI KASIR ONLINE")
}
