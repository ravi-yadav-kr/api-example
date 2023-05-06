package main

import (
	"github.com/gin-gonic/gin"

	"api-example/db"
	"api-example/handler"
	"api-example/service"
	"api-example/store"
	"fmt"
)

func main() {
	r := gin.Default()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Dbname)

	dbObj := db.InitialiseDb(psqlInfo)
	if dbObj == nil {
		return
	}

	stg := store.New(dbObj)
	svc := service.New(stg)
	h := handler.New(svc)

	r.POST("/employee", h.Create)
	r.GET("/employee", h.Get)
	r.PUT("/employee", h.Update)
	r.DELETE("/employee", h.Delete)
	r.PATCH("/employee")

	err := r.Run()
	if err != nil {
		return
	}

}
