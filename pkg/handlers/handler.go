package handlers

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/arturoeanton/go-struc2fiber/pkg/repositories"
	"github.com/arturoeanton/go-struc2fiber/pkg/services"
	"github.com/arturoeanton/go-struc2fiber/pkg/validator"
	fiber "github.com/gofiber/fiber/v2"
)

type IHandler[T any] interface {
	Name() string
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	FxCreate(vals ...*validator.StructValidator) func(c *fiber.Ctx) error
	DeleteByID(c *fiber.Ctx) error
	FxUpdate(vals ...*validator.StructValidator) func(c *fiber.Ctx) error
}

type Handler[T any] struct {
	service services.IService[T]
	name    string
}

func NewHandler[T any]() *Handler[T] {

	return &Handler[T]{
		name: "items",
		service: services.NewService[T](
			repositories.NewRepository[T](),
		),
	}
}

func (h *Handler[T]) Name() string {
	return h.name
}

func (h *Handler[T]) GetAll(c *fiber.Ctx) error {
	items, _, err := h.service.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": "Failed to get " + h.Name(),
		})
	}
	return c.Status(http.StatusOK).JSON(items)
}

func (h *Handler[T]) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": "Failed to get " + h.Name(),
		})
	}
	return c.Status(http.StatusOK).JSON(item)
}

func (h *Handler[T]) FxCreate(vals ...*validator.StructValidator) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		item := new(T)
		if err := c.BodyParser(item); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": "Failed to get " + h.Name(),
			})
		}
		if len(vals) > 0 && vals[0] != nil {
			flagValid, errors := vals[0].ValidateStruct(item)
			if !flagValid {
				return c.Status(http.StatusBadRequest).JSON(map[string]any{
					"error":  "Validation failed",
					"fields": errors,
				})
			}
		}
		id, err := h.service.Create(item)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": "Failed to get " + h.Name(),
			})
		}
		reflect.ValueOf(item).Elem().FieldByName("ID").SetInt(id)
		return c.Status(http.StatusOK).JSON(id)
	}
}

func (h *Handler[T]) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	rowAffected, err := h.service.Delete(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": "Failed to get " + h.Name(),
		})
	}
	return c.Status(http.StatusOK).JSON(map[string]int64{"rows_affected": rowAffected})
}

func (h *Handler[T]) FxUpdate(vals ...*validator.StructValidator) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(map[string]string{
				"error": "Invalid ID",
			})
		}
		_, err = h.service.GetByID(id)
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(map[string]string{
				"error": h.Name() + " not found",
			})
		}
		item := new(T)
		if err := c.BodyParser(item); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": "Failed to get " + h.Name(),
			})
		}

		reflect.ValueOf(item).Elem().FieldByName("ID").SetInt(idInt)

		if len(vals) > 0 && vals[0] != nil {
			flagValid, errors := vals[0].ValidateStruct(item)
			if !flagValid {
				return c.Status(http.StatusBadRequest).JSON(map[string]any{
					"error":  "Validation failed",
					"fields": errors,
				})
			}
		}

		rowAffected, err := h.service.Update(item)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": "Failed to get " + h.Name(),
			})
		}
		return c.Status(http.StatusNoContent).JSON(rowAffected)
	}
}
