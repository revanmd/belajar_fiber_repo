package handlers

import (
	"belajar-api/models"
	"belajar-api/response"
	"belajar-api/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserBelajarHandler struct {
	UserBelajarService *services.UserBelajarService
}

// GetAllUserBelajar handles the request to fetch all UserBelajar records
func (h *UserBelajarHandler) GetAllUserBelajar(c *fiber.Ctx) error {
	nama := c.Query("nama")
	users, err := h.UserBelajarService.GetAllUserBelajar()
	if err != nil {
		return response.InternalServerError(c, err, "Cannot retrieve users")
	}
	log.Println(nama)
	return response.SuccessHandler(c, fiber.StatusOK, "Users retrieved successfully", users)
}

// GetUserBelajarByUID handles the request to fetch a UserBelajar by its UID
func (h *UserBelajarHandler) GetUserBelajarByUID(c *fiber.Ctx) error {
	uid := c.Params("uid")
	user, err := h.UserBelajarService.GetUserBelajarByUID(uid)
	if err != nil {
		return response.InternalServerError(c, err, "Cannot retrieve user")
	}
	if user == nil {
		return response.BadRequest(c, nil, "User not found")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "User retrieved successfully", user)
}

// CreateUserBelajar handles the request to create a new UserBelajar
func (h *UserBelajarHandler) CreateUserBelajar(c *fiber.Ctx) error {
	var user models.UserBelajar
	if err := c.BodyParser(&user); err != nil {
		return response.BadRequest(c, err, "Cannot parse JSON")
	}
	if err := h.UserBelajarService.CreateUserBelajar(&user); err != nil {
		return response.InternalServerError(c, err, "Cannot create user")
	}
	return response.SuccessHandler(c, fiber.StatusCreated, "User created successfully", user)
}

// UpdateUserBelajar handles the request to update a UserBelajar by its UID
func (h *UserBelajarHandler) UpdateUserBelajar(c *fiber.Ctx) error {
	uid := c.Params("uid")
	var updatedUser models.UserBelajar
	if err := c.BodyParser(&updatedUser); err != nil {
		return response.BadRequest(c, err, "Cannot parse JSON")
	}
	if err := h.UserBelajarService.UpdateUserBelajar(uid, &updatedUser); err != nil {
		return response.InternalServerError(c, err, "Cannot update user")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "User updated successfully", updatedUser)
}

// DeleteUserBelajar handles the request to delete a UserBelajar by its UID
func (h *UserBelajarHandler) DeleteUserBelajar(c *fiber.Ctx) error {
	uid := c.Params("uid")
	if err := h.UserBelajarService.DeleteUserBelajar(uid); err != nil {
		return response.InternalServerError(c, err, "Cannot delete user")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "User deleted successfully", nil)
}
