package main

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
	"github.com/raitonoberu/personal-best/app/handler"
	"github.com/raitonoberu/personal-best/app/middleware"
	"github.com/raitonoberu/personal-best/app/router"
)

func main() {
	db, err := sql.Open("sqlite", ".db/db.sqlite")
	if err != nil {
		panic(err)
	}

	router := router.New()

	h := handler.New(db)
	router.POST("/api/login", h.Login)
	router.POST("/api/register", h.Register)

	// router.GET("/api/users", h.ListUsers)
	router.GET("/api/users/:id", h.GetUser)
	router.PATCH("/api/users", h.UpdateUser, middleware.MustAuth)
	router.DELETE("/api/users", h.DeleteUser, middleware.MustAuth)

	router.POST("/api/competitions", h.CreateCompetition, middleware.MustAuth)
	router.GET("/api/competitions", h.ListCompetitions)
	router.GET("/api/competitions/:id", h.GetCompetition)
	router.PATCH("/api/competitions/:id", h.UpdateCompetition, middleware.MustAuth)
	router.DELETE("/api/competitions/:id", h.DeleteCompetition, middleware.MustAuth)

	panic(router.Start(":8080"))
}
