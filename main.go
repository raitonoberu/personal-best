package main

import (
	"fmt"
	"os"

	"github.com/raitonoberu/personal-best/cmd"
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
	err := cmd.App.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal:", err.Error())
		os.Exit(1)
	}
}
