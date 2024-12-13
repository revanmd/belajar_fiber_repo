package handlers

import (
	"belajar-api/models"
	"belajar-api/response"
	"belajar-api/services"

	"github.com/gofiber/fiber/v2"
)

type MasterPriorityHandler struct {
	MasterPriorityService *services.MasterPriorityService
}

// GetAllMasterPriorities handles the request to fetch all MasterPriorities
func (h *MasterPriorityHandler) GetAllMasterPriorities(c *fiber.Ctx) error {
	priorities, err := h.MasterPriorityService.GetAllMasterPriorities()
	if err != nil {
		return response.InternalServerError(c, err, "Cannot retrieve priorities")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Priorities retrieved successfully", priorities)
}

// GetMasterPriorityByUID handles the request to fetch a MasterPriority by its UID
func (h *MasterPriorityHandler) GetMasterPriorityByUID(c *fiber.Ctx) error {
	uid := c.Params("uid")
	priority, err := h.MasterPriorityService.GetMasterPriorityByUID(uid)
	if err != nil {
		return response.InternalServerError(c, err, "Cannot retrieve priority")
	}
	if priority == nil {
		return response.BadRequest(c, err, "Priority not found")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Priority retrieved successfully", priority)
}

// CreateMasterPriority handles the request to create a new MasterPriority
func (h *MasterPriorityHandler) CreateMasterPriority(c *fiber.Ctx) error {
	var priority models.MasterPriority
	if err := c.BodyParser(&priority); err != nil {
		return response.BadRequest(c, err, "Cannot parse JSON")
	}
	if err := h.MasterPriorityService.CreateMasterPriority(&priority); err != nil {
		return response.InternalServerError(c, err, "Cannot create priority")
	}
	return response.SuccessHandler(c, fiber.StatusCreated, "Priority created successfully", priority)
}

// UpdateMasterPriority handles the request to update a MasterPriority by its UID
func (h *MasterPriorityHandler) UpdateMasterPriority(c *fiber.Ctx) error {
	uid := c.Params("uid")
	var updatedPriority models.MasterPriority
	if err := c.BodyParser(&updatedPriority); err != nil {
		return response.BadRequest(c, err, "Cannot parse JSON")
	}
	if err := h.MasterPriorityService.UpdateMasterPriority(uid, &updatedPriority); err != nil {
		return response.InternalServerError(c, err, "Cannot update priority")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Priority updated successfully", updatedPriority)
}

// DeleteMasterPriority handles the request to delete a MasterPriority by its UID
func (h *MasterPriorityHandler) DeleteMasterPriority(c *fiber.Ctx) error {
	uid := c.Params("uid")
	if err := h.MasterPriorityService.DeleteMasterPriority(uid); err != nil {
		return response.InternalServerError(c, err, "Cannot delete priority")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Priority deleted successfully", nil)
}
