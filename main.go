package main

import (
	"github.com/arturoeanton/go-struc2fiber/pkg/commons"
	"github.com/arturoeanton/go-struc2fiber/pkg/model"
	"github.com/arturoeanton/go-struc2fiber/pkg/repositories"
	"github.com/arturoeanton/go-struc2fiber/pkg/web"
	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	repositories.FlagLog = false

	dbSource := commons.Getenv("DB_SOURCE", "db.sqlite")
	db, err := gorm.Open(sqlite.Open(dbSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	repositories.DB = db

	app := fiber.New()

	web.RegisterCRUD(app, "internal_users", model.InternalUser{}, "schemas/internal_users.yaml")
	web.RegisterCRUD(app, "skill", model.Skill{}, "schemas/create_skill.yaml", "schemas/update_skill.yaml")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")

}
