package main

import (
	"fmt"

	"github.com/arturoeanton/go-struc2fiber/pkg/commons"
	"github.com/arturoeanton/go-struc2fiber/pkg/handlers"
	"github.com/arturoeanton/go-struc2fiber/pkg/model"
	"github.com/arturoeanton/go-struc2fiber/pkg/repositories"
	"github.com/arturoeanton/go-struc2fiber/pkg/validator"
	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	repositories.FlagLog = false

	//dbDriver := commons.Getenv("DB_DRIVER", "sqlite3")
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

	app.Get("/internal_users", handlers.NewHandler[model.InternalUser]().GetAll)
	app.Get("/internal_users/:id", handlers.NewHandler[model.InternalUser]().GetByID)
	app.Post("/internal_users", handlers.NewHandler[model.InternalUser]().FxCreate(GetValidator("schemas/internal_users.yaml")))
	app.Delete("/internal_users/:id", handlers.NewHandler[model.InternalUser]().DeleteByID)
	app.Put("/internal_users/:id", handlers.NewHandler[model.InternalUser]().FxUpdate(GetValidator("schemas/internal_users.yaml")))

	app.Get("/skill", handlers.NewHandler[model.Skill]().GetAll)
	app.Get("/skill/:id", handlers.NewHandler[model.Skill]().GetByID)
	app.Post("/skill", handlers.NewHandler[model.Skill]().FxCreate(GetValidator("schemas/create_skill.yaml")))
	app.Delete("/skill/:id", handlers.NewHandler[model.Skill]().DeleteByID)
	app.Put("/skill/:id", handlers.NewHandler[model.Skill]().FxUpdate(GetValidator("schemas/update_skill.yaml")))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")

}

func GetValidator(path string) *validator.StructValidator {
	val := validator.NewStructValidator()
	err := val.LoadSchemaFromFile(path)

	if err != nil {
		fmt.Println("Error loading schema:", err)
		val = nil
	}
	return val
}
