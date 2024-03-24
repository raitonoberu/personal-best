package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/raitonoberu/personal-best/app/controller"
	"github.com/raitonoberu/personal-best/app/middleware"
	"github.com/raitonoberu/personal-best/db"
)

func main() {
	if err := db.Init(".db/db.sqlite"); err != nil {
		panic(err)
	}

	f := fiber.New()

	f.Use(logger.New())
	f.Use(middleware.Auth)

	f.Get("/login", controller.LoginPage)
	f.Post("/login", controller.Login)
	f.Get("/register", controller.RegisterPage)
	f.Post("/register", controller.Register)
	f.Get("/logout", controller.Logout)

	f.Get("/profile", controller.GetProfile, middleware.MustAuth)
	f.Patch("/profile", controller.UpdateProfile, middleware.MustAuth)
	f.Delete("/profile", controller.DeleteProfile, middleware.MustAuth)

	f.Get("/", controller.ListCompetitions)
	f.Get("/competitions", controller.ListCompetitions)
	f.Post("/competitions", controller.NewCompetition, middleware.MustAuth)
	f.Get("/competitions/:id", controller.ViewCompetition)
	f.Patch("/competitions/:id", controller.UpdateCompetition, middleware.MustAuth)
	f.Delete("/competitions/:id", controller.DeleteCompetition, middleware.MustAuth)

	f.Post("/competitions/:comp/players", controller.AddPlayer, middleware.MustAuth)
	f.Patch("/competitions/:comp/players/:id", controller.UpdatePlayer, middleware.MustAuth)
	f.Delete("/competitions/:comp/players/:id", controller.DeletePlayer, middleware.MustAuth)

	f.Post("/competitions/:comp/games", controller.AddGame, middleware.MustAuth)

	f.Static("/static", "./static")

	panic(f.Listen(":8080"))
}
