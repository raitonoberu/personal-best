package admin

import (
	"database/sql"
	"fmt"

	"github.com/raitonoberu/personal-best/db/sqlc"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
)

var Create = &cli.Command{
	Name:     "create",
	Usage:    "create admin user",
	Category: "admin",
	Action: func(ctx *cli.Context) error {
		db, err := sql.Open("sqlite", ".db/db.sqlite")
		if err != nil {
			return err
		}
		queries := sqlc.New(db)

		var email, password string
		fmt.Print("Enter email: ")
		fmt.Scanln(&email)
		fmt.Print("Enter password: ")
		fmt.Scanln(&password)

		passwordHash, err := bcrypt.GenerateFromPassword(
			[]byte(password), bcrypt.DefaultCost,
		)
		if err != nil {
			return err
		}

		_, err = queries.CreateUser(ctx.Context, sqlc.CreateUserParams{
			RoleID:     1,
			Email:      email,
			Password:   string(passwordHash),
			FirstName:  "Админ",
			MiddleName: "Админов",
			LastName:   "Админович",
		})
		return err
	},
}
