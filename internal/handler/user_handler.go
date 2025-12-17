package handler

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/user-api/internal/models"
	"github.com/user-api/internal/service"
)

type UserHandler struct {
	svc *service.UserService
	val *validator.Validate
}

func NewUserHandler(s *service.UserService) *UserHandler {
	h := &UserHandler{
		svc: s,
		val: validator.New(),
	}
	return h
}

// create user
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	// parse body
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
	}

	// validate
	err = h.val.Struct(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "name and dob are required"})
	}

	// create
	user, err := h.svc.Create(c.Context(), req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(user)
}

// get user
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id_str := c.Params("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	user, err := h.svc.GetByID(c.Context(), int32(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

// list users
func (h *UserHandler) List(c *fiber.Ctx) error {
	users, err := h.svc.List(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// return empty array not null
	if users == nil {
		users = []models.UserWithAge{}
	}

	return c.JSON(users)
}

// update user
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id_str := c.Params("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	var req models.UpdateUserRequest
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
	}

	err = h.val.Struct(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "name and dob are required"})
	}

	user, err := h.svc.Update(c.Context(), int32(id), req)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

// delete user
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id_str := c.Params("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	err = h.svc.Delete(c.Context(), int32(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}
