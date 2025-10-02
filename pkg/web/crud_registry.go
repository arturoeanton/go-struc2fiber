package web

import (
	"fmt"
	"reflect"

	"github.com/arturoeanton/go-struc2fiber/pkg/handlers"
	"github.com/arturoeanton/go-struc2fiber/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterCRUD[T any](app *fiber.App, resourceName string, model T, vals ...string) {
	modelType := reflect.TypeOf(model)

	// Auto-genera todas las rutas CRUD
	handler := handlers.NewHandler[T]()
	validator1 := validator.NewStructValidator()
	validator1.LoadSchemaFromFile(vals[0])

	var validator2 *validator.StructValidator
	validator2 = validator1
	if len(vals) > 1 {
		validator2 = validator.NewStructValidator()
		validator2.LoadSchemaFromFile(vals[1])
	}

	app.Get("/"+resourceName, handler.GetAll)
	app.Get("/"+resourceName+"/:id", handler.GetByID)
	app.Post("/"+resourceName, handler.FxCreate(validator1))
	app.Put("/"+resourceName+"/:id", handler.FxUpdate(validator2))
	app.Delete("/"+resourceName+"/:id", handler.DeleteByID)

	fmt.Printf("Registered CRUD routes for %s at /%s\n", modelType.Name(), resourceName)
}
