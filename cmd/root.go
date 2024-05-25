package cmd

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
	"github.com/raitonoberu/personal-best/app/handler"
	"github.com/raitonoberu/personal-best/app/middleware"
	"github.com/raitonoberu/personal-best/app/router"
	"github.com/raitonoberu/personal-best/cmd/admin"
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:  "personal-best",
	Usage: "run HTTP server",
	Action: func(ctx *cli.Context) error {
		db, err := sql.Open("sqlite", ".db/db.sqlite")
		if err != nil {
			return err
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

		router.POST("/api/competitions/:id/registrations", h.RegisterForCompetition, middleware.Auth)
		router.GET("/api/competitions/:id/registrations", h.ListCompetitionRegistrations, middleware.Auth)
		router.DELETE("/api/competitions/:id/registrations", h.UnregisterForCompetition, middleware.Auth)

		router.GET("/api/roles", h.ListRoles, middleware.Auth)

		router.POST("/api/admin/users", h.AdminCreateUser, middleware.Auth)
		router.PATCH("/api/admin/users/:id", h.AdminUpdateUser, middleware.Auth)

		return router.Start(":8080")
	},
	Commands: []*cli.Command{
		admin.Create,
	},
}
