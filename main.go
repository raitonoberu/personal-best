package main

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
	"github.com/raitonoberu/personal-best/app/handler"
	"github.com/raitonoberu/personal-best/app/router"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

func main() {
	dbConn, err := sql.Open("sqlite", ".db/db.sqlite")
	if err != nil {
		panic(err)
	}
	db := sqlc.New(dbConn)

	router := router.New()

	h := handler.New(db)
	router.POST("/api/login", h.Login)
	router.POST("/api/register", h.Register)

	router.GET("/api/users", h.ListUsers)
	router.GET("/api/users/:id", h.GetUser)
	// router.PATCH("/api/users", h.UpdateUser)
	router.DELETE("/api/users", h.DeleteUser)

	// router.POST("/api/competitions", h.CreateCompetition)
	// router.GET("/api/competitions", h.ListCompetitions)
	router.GET("/api/competitions/:id", h.GetCompetition)
	// router.DELETE("/api/competitions/:id", h.DeleteCompetition)

	panic(router.Start(":8080"))
}
