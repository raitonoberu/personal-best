package main

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
	"github.com/raitonoberu/personal-best/app/handler"
	"github.com/raitonoberu/personal-best/app/middleware"
	"github.com/raitonoberu/personal-best/app/router"
)

// @title Personal Best API
// @version 0.1
// @description neмытьlE yблюdki

// @contact.name   raitonoberu
// @contact.url    http://raitonobe.ru
// @contact.email  raitonoberu@mail.ru

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	db, err := sql.Open("sqlite", ".db/db.sqlite")
	if err != nil {
		panic(err)
	}

	router := router.New()

	h := handler.New(db)
	router.POST("/api/login", h.Login)
	router.POST("/api/register", h.Register)

	router.GET("/api/users/:id", h.GetUser, middleware.Auth)
	router.PATCH("/api/users", h.UpdateUser, middleware.Auth)
	router.DELETE("/api/users", h.DeleteUser, middleware.Auth)

	router.POST("/api/competitions", h.CreateCompetition, middleware.Auth)
	router.GET("/api/competitions", h.ListCompetitions, middleware.Auth)
	router.GET("/api/competitions/:id", h.GetCompetition, middleware.Auth)
	router.PATCH("/api/competitions/:id", h.UpdateCompetition, middleware.Auth)
	router.DELETE("/api/competitions/:id", h.DeleteCompetition, middleware.Auth)

	router.GET("/api/admin/roles", h.ListRoles, middleware.Auth)

	panic(router.Start(":8080"))
}
